package repository

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog/log"

	"github.com/gofrs/uuid"
	"github.com/nuriansyah/lokatra-payment/internal/domain/payment/model"
	"github.com/nuriansyah/lokatra-payment/shared/failure"
)

// ─── Interfaces ──────────────────────────────────────────────────────────────

type InvoiceRepository interface {
	CreateInvoice(ctx context.Context, inv *model.Invoices, items []model.InvoiceLineItems) error
	ResolveInvoiceByID(ctx context.Context, id uuid.UUID) (*model.Invoices, error)
	ResolveInvoicesByFilter(ctx context.Context, filter model.Filter) ([]model.InvoicesFilterResult, error)
	UpdateInvoiceByID(ctx context.Context, id uuid.UUID, inv *model.Invoices) error
	UpdateInvoiceRemainingAmount(ctx context.Context, id uuid.UUID, paymentAmount uuid.UUID, delta interface{}) (int64, error)
}

type InvoiceLineItemRepository interface {
	CreateInvoiceLineItems(ctx context.Context, items []model.InvoiceLineItems) error
	ResolveInvoiceLineItemsByInvoiceID(ctx context.Context, invoiceID uuid.UUID) ([]model.InvoiceLineItems, error)
}

type InvoicePaymentLinkRepository interface {
	CreateInvoicePaymentLinks(ctx context.Context, links []model.InvoicePaymentLinks) error
	ResolveInvoicePaymentLinksByInvoiceID(ctx context.Context, invoiceID uuid.UUID) ([]model.InvoicePaymentLinks, error)
	UpdateInvoicePaymentLinkClickCount(ctx context.Context, id uuid.UUID) error
}

type InvoicePaymentRepository interface {
	CreateInvoicePayment(ctx context.Context, pay *model.InvoicePayments) error
	ResolveInvoicePaymentsByInvoiceID(ctx context.Context, invoiceID uuid.UUID) ([]model.InvoicePayments, error)
}

type InvoiceStatusEventRepository interface {
	CreateInvoiceStatusEvent(ctx context.Context, event *model.InvoiceStatusEvents) error
	ResolveInvoiceStatusEventsByInvoiceID(ctx context.Context, invoiceID uuid.UUID) ([]model.InvoiceStatusEvents, error)
}

// ─── Invoice Implementation ──────────────────────────────────────────────────

func defaultInvoiceSelectFields() string {
	return `"id","invoice_code","merchant_id","customer_id","customer_name","customer_email","customer_phone",
		"subtotal","discount_amount","pph23_amount","pph23_rate","ppn_amount","ppn_rate",
		"total_amount","paid_amount","remaining_amount","currency","status",
		"issued_at","due_at","paid_at","voided_at","void_reason",
		"source_service","source_type","source_id","description","notes","terms",
		"meta_created_at","meta_created_by","meta_updated_at","meta_updated_by","meta_deleted_at","meta_deleted_by"`
}

func (r *RepositoryImpl) CreateInvoice(ctx context.Context, inv *model.Invoices, items []model.InvoiceLineItems) error {
	tx, err := r.db.Write.BeginTxx(ctx, nil)
	if err != nil {
		return failure.InternalError(err)
	}
	defer tx.Rollback()

	query := fmt.Sprintf(`INSERT INTO "invoices" (%s) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13,$14,$15,$16,$17,$18,$19,$20,$21,$22,$23,$24,$25,$26,$27,$28,$29,$30)`,
		`"id","invoice_code","merchant_id","customer_id","customer_name","customer_email","customer_phone",
		 "subtotal","discount_amount","pph23_amount","pph23_rate","ppn_amount","ppn_rate",
		 "total_amount","paid_amount","remaining_amount","currency","status",
		 "issued_at","due_at","paid_at","voided_at","void_reason",
		 "source_service","source_type","source_id","description","notes","terms",
		 "meta_created_at","meta_created_by"`)

	_, err = tx.ExecContext(ctx, query,
		inv.Id, inv.InvoiceCode, inv.MerchantId, inv.CustomerId,
		inv.CustomerName, inv.CustomerEmail, inv.CustomerPhone,
		inv.Subtotal, inv.DiscountAmount, inv.Pph23Amount, inv.Pph23Rate,
		inv.PpnAmount, inv.PpnRate, inv.TotalAmount, inv.PaidAmount, inv.RemainingAmount,
		inv.Currency, inv.Status,
		inv.IssuedAt, inv.DueAt, inv.PaidAt, inv.VoidedAt, inv.VoidReason,
		inv.SourceService, inv.SourceType, inv.SourceId,
		inv.Description, inv.Notes, inv.Terms,
		inv.MetaCreatedAt, inv.MetaCreatedBy,
	)
	if err != nil {
		log.Error().Err(err).Msg("[CreateInvoice] failed")
		return failure.InternalError(err)
	}

	if len(items) > 0 {
		if err := insertInvoiceLineItems(ctx, tx, inv.Id, items); err != nil {
			return err
		}
	}

	return tx.Commit()
}

func insertInvoiceLineItems(ctx context.Context, tx *sqlx.Tx, invoiceID uuid.UUID, items []model.InvoiceLineItems) error {
	query := `INSERT INTO "invoice_line_items"
		("id","invoice_id","line_no","name","description","quantity","unit_price",
		 "discount_percent","discount_amount","ppn_rate","ppn_amount","pph23_rate","pph23_amount",
		 "subtotal","total_amount","currency","sku","category",
		 "meta_created_at","meta_created_by")
		VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13,$14,$15,$16,$17,$18,$19,$20)`

	for _, item := range items {
		_, err := tx.ExecContext(ctx, query,
			item.Id, item.InvoiceId, item.LineNo, item.Name, item.Description,
			item.Quantity, item.UnitPrice, item.DiscountPercent, item.DiscountAmount,
			item.PpnRate, item.PpnAmount, item.Pph23Rate, item.Pph23Amount,
			item.Subtotal, item.TotalAmount, item.Currency, item.Sku, item.Category,
			item.MetaCreatedAt, item.MetaCreatedBy,
		)
		if err != nil {
			log.Error().Err(err).Msg("[insertInvoiceLineItems] failed")
			return failure.InternalError(err)
		}
	}
	return nil
}

func (r *RepositoryImpl) ResolveInvoiceByID(ctx context.Context, id uuid.UUID) (*model.Invoices, error) {
	var inv model.Invoices
	query := fmt.Sprintf(`SELECT %s FROM "invoices" WHERE "id" = $1 AND "meta_deleted_at" IS NULL`, defaultInvoiceSelectFields())
	if err := r.db.Read.GetContext(ctx, &inv, query, id); err != nil {
		if err == sql.ErrNoRows {
			return nil, failure.NotFound("invoice")
		}
		log.Error().Err(err).Msg("[ResolveInvoiceByID] failed")
		return nil, failure.InternalError(err)
	}
	return &inv, nil
}

func (r *RepositoryImpl) ResolveInvoicesByFilter(ctx context.Context, filter model.Filter) (result []model.InvoicesFilterResult, err error) {
	var args []interface{}
	whereClauses := []string{`"meta_deleted_at" IS NULL`}

	for _, f := range filter.FilterFields {
		switch f.Field {
		case "status":
			args = append(args, f.Value)
			whereClauses = append(whereClauses, fmt.Sprintf(`"status" = $%d`, len(args)))
		case "merchant_id":
			args = append(args, f.Value)
			whereClauses = append(whereClauses, fmt.Sprintf(`"merchant_id" = $%d`, len(args)))
		case "customer_id":
			args = append(args, f.Value)
			whereClauses = append(whereClauses, fmt.Sprintf(`"customer_id" = $%d`, len(args)))
		case "invoice_code":
			args = append(args, "%"+fmt.Sprintf("%v", f.Value)+"%")
			whereClauses = append(whereClauses, fmt.Sprintf(`"invoice_code" LIKE $%d`, len(args)))
		}
	}

	whereStr := strings.Join(whereClauses, " AND ")
	offset := 0
	if filter.Pagination.Page > 1 {
		offset = (filter.Pagination.Page - 1) * filter.Pagination.PageSize
	}
	limit := filter.Pagination.PageSize
	if limit <= 0 {
		limit = 20
	}

	args = append(args, limit, offset)
	query := fmt.Sprintf(`SELECT %s FROM "invoices" WHERE %s ORDER BY "meta_created_at" DESC LIMIT $%d OFFSET $%d`,
		defaultInvoiceSelectFields(), whereStr, len(args)-1, len(args))

	var items []model.Invoices
	if err = r.db.Read.SelectContext(ctx, &items, query, args...); err != nil {
		log.Error().Err(err).Msg("[ResolveInvoicesByFilter] failed")
		return nil, failure.InternalError(err)
	}

	result = make([]model.InvoicesFilterResult, len(items))
	for i, item := range items {
		result[i] = model.InvoicesFilterResult{Invoices: item}
	}
	return result, nil
}

func (r *RepositoryImpl) UpdateInvoiceByID(ctx context.Context, id uuid.UUID, inv *model.Invoices) error {
	query := `UPDATE "invoices" SET
		"invoice_code" = $1,"merchant_id" = $2,"customer_id" = $3,
		"customer_name" = $4,"customer_email" = $5,"customer_phone" = $6,
		"subtotal" = $7,"discount_amount" = $8,"pph23_amount" = $9,"pph23_rate" = $10,
		"ppn_amount" = $11,"ppn_rate" = $12,"total_amount" = $13,"paid_amount" = $14,"remaining_amount" = $15,
		"currency" = $16,"status" = $17,
		"issued_at" = $18,"due_at" = $19,"paid_at" = $20,"voided_at" = $21,"void_reason" = $22,
		"description" = $23,"notes" = $24,"terms" = $25,
		"meta_updated_at" = $26,"meta_updated_by" = $27
		WHERE "id" = $28 AND "meta_deleted_at" IS NULL`
	result, err := r.exec(ctx, query, []interface{}{
		inv.InvoiceCode, inv.MerchantId, inv.CustomerId,
		inv.CustomerName, inv.CustomerEmail, inv.CustomerPhone,
		inv.Subtotal, inv.DiscountAmount, inv.Pph23Amount, inv.Pph23Rate,
		inv.PpnAmount, inv.PpnRate, inv.TotalAmount, inv.PaidAmount, inv.RemainingAmount,
		inv.Currency, inv.Status,
		inv.IssuedAt, inv.DueAt, inv.PaidAt, inv.VoidedAt, inv.VoidReason,
		inv.Description, inv.Notes, inv.Terms,
		inv.MetaUpdatedAt, inv.MetaUpdatedBy, id,
	})
	if err != nil {
		log.Error().Err(err).Msg("[UpdateInvoiceByID] failed")
		return failure.InternalError(err)
	}
	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return failure.NotFound("invoice")
	}
	return nil
}

func (r *RepositoryImpl) UpdateInvoiceRemainingAmount(ctx context.Context, id uuid.UUID, _ uuid.UUID, delta interface{}) (int64, error) {
	query := `UPDATE "invoices" SET
		"paid_amount" = "paid_amount" + $2,
		"remaining_amount" = "remaining_amount" - $2,
		"meta_updated_at" = now()
		WHERE "id" = $1 AND "remaining_amount" >= $2 AND "meta_deleted_at" IS NULL`
	result, err := r.exec(ctx, query, []interface{}{id, delta})
	if err != nil {
		log.Error().Err(err).Msg("[UpdateInvoiceRemainingAmount] failed")
		return 0, failure.InternalError(err)
	}
	rowsAffected, _ := result.RowsAffected()
	return rowsAffected, nil
}

// ─── Line Item Implementation ────────────────────────────────────────────────

func (r *RepositoryImpl) CreateInvoiceLineItems(ctx context.Context, items []model.InvoiceLineItems) error {
	if len(items) == 0 {
		return nil
	}
	tx, err := r.db.Write.BeginTxx(ctx, nil)
	if err != nil {
		return failure.InternalError(err)
	}
	defer tx.Rollback()
	if err := insertInvoiceLineItems(ctx, tx, items[0].InvoiceId, items); err != nil {
		return err
	}
	return tx.Commit()
}

func (r *RepositoryImpl) ResolveInvoiceLineItemsByInvoiceID(ctx context.Context, invoiceID uuid.UUID) ([]model.InvoiceLineItems, error) {
	var items []model.InvoiceLineItems
	query := `SELECT "id","invoice_id","line_no","name","description","quantity","unit_price",
		"discount_percent","discount_amount","ppn_rate","ppn_amount","pph23_rate","pph23_amount",
		"subtotal","total_amount","currency","sku","category",
		"meta_created_at","meta_created_by","meta_updated_at","meta_updated_by","meta_deleted_at","meta_deleted_by"
		FROM "invoice_line_items" WHERE "invoice_id" = $1 AND "meta_deleted_at" IS NULL ORDER BY "line_no" ASC`
	if err := r.db.Read.SelectContext(ctx, &items, query, invoiceID); err != nil {
		return nil, failure.InternalError(err)
	}
	return items, nil
}

// ─── Payment Link Implementation ─────────────────────────────────────────────

func (r *RepositoryImpl) CreateInvoicePaymentLinks(ctx context.Context, links []model.InvoicePaymentLinks) error {
	for _, link := range links {
		query := `INSERT INTO "invoice_payment_links"
			("id","invoice_id","provider_code","link_url","link_type","status","click_count","payment_count","expires_at",
			 "meta_created_at","meta_created_by")
			VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11)`
		_, err := r.exec(ctx, query, []interface{}{
			link.Id, link.InvoiceId, link.ProviderCode, link.LinkUrl, link.LinkType,
			link.Status, link.ClickCount, link.PaymentCount, link.ExpiresAt,
			link.MetaCreatedAt, link.MetaCreatedBy,
		})
		if err != nil {
			log.Error().Err(err).Msg("[CreateInvoicePaymentLinks] failed")
			return failure.InternalError(err)
		}
	}
	return nil
}

func (r *RepositoryImpl) ResolveInvoicePaymentLinksByInvoiceID(ctx context.Context, invoiceID uuid.UUID) ([]model.InvoicePaymentLinks, error) {
	var links []model.InvoicePaymentLinks
	query := `SELECT "id","invoice_id","provider_code","link_url","link_type","status","click_count","payment_count","expires_at",
		"meta_created_at","meta_created_by","meta_updated_at","meta_updated_by","meta_deleted_at","meta_deleted_by"
		FROM "invoice_payment_links" WHERE "invoice_id" = $1 AND "meta_deleted_at" IS NULL`
	if err := r.db.Read.SelectContext(ctx, &links, query, invoiceID); err != nil {
		return nil, failure.InternalError(err)
	}
	return links, nil
}

func (r *RepositoryImpl) UpdateInvoicePaymentLinkClickCount(ctx context.Context, id uuid.UUID) error {
	query := `UPDATE "invoice_payment_links" SET "click_count" = "click_count" + 1 WHERE "id" = $1 AND "meta_deleted_at" IS NULL`
	if _, err := r.exec(ctx, query, []interface{}{id}); err != nil {
		return failure.InternalError(err)
	}
	return nil
}

// ─── Invoice Payment Implementation ──────────────────────────────────────────

func (r *RepositoryImpl) CreateInvoicePayment(ctx context.Context, pay *model.InvoicePayments) error {
	query := `INSERT INTO "invoice_payments"
		("id","invoice_id","payment_intent_id","provider_code","amount","currency","status",
		 "meta_created_at","meta_created_by")
		VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9)`
	_, err := r.exec(ctx, query, []interface{}{
		pay.Id, pay.InvoiceId, pay.PaymentIntentId, pay.ProviderCode,
		pay.Amount, pay.Currency, pay.Status,
		pay.MetaCreatedAt, pay.MetaCreatedBy,
	})
	if err != nil {
		log.Error().Err(err).Msg("[CreateInvoicePayment] failed")
		return failure.InternalError(err)
	}
	return nil
}

func (r *RepositoryImpl) ResolveInvoicePaymentsByInvoiceID(ctx context.Context, invoiceID uuid.UUID) ([]model.InvoicePayments, error) {
	var pays []model.InvoicePayments
	query := `SELECT "id","invoice_id","payment_intent_id","provider_code","amount","currency","status",
		"meta_created_at","meta_created_by","meta_updated_at","meta_updated_by"
		FROM "invoice_payments" WHERE "invoice_id" = $1 ORDER BY "meta_created_at" DESC`
	if err := r.db.Read.SelectContext(ctx, &pays, query, invoiceID); err != nil {
		return nil, failure.InternalError(err)
	}
	return pays, nil
}

// ─── Status Event Implementation ─────────────────────────────────────────────

func (r *RepositoryImpl) CreateInvoiceStatusEvent(ctx context.Context, event *model.InvoiceStatusEvents) error {
	query := `INSERT INTO "invoice_status_events"
		("id","invoice_id","event_type","old_status","new_status","reason","actor_id",
		 "meta_created_at","meta_created_by")
		VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9)`
	_, err := r.exec(ctx, query, []interface{}{
		event.Id, event.InvoiceId, event.EventType, event.OldStatus,
		event.NewStatus, event.Reason, event.ActorId,
		event.MetaCreatedAt, event.MetaCreatedBy,
	})
	if err != nil {
		log.Error().Err(err).Msg("[CreateInvoiceStatusEvent] failed")
		return failure.InternalError(err)
	}
	return nil
}

func (r *RepositoryImpl) ResolveInvoiceStatusEventsByInvoiceID(ctx context.Context, invoiceID uuid.UUID) ([]model.InvoiceStatusEvents, error) {
	var events []model.InvoiceStatusEvents
	query := `SELECT "id","invoice_id","event_type","old_status","new_status","reason","actor_id",
		"meta_created_at","meta_created_by"
		FROM "invoice_status_events" WHERE "invoice_id" = $1 ORDER BY "meta_created_at" ASC`
	if err := r.db.Read.SelectContext(ctx, &events, query, invoiceID); err != nil {
		return nil, failure.InternalError(err)
	}
	return events, nil
}

// Compile-time checks
var (
	_ InvoiceRepository          = (*RepositoryImpl)(nil)
	_ InvoiceLineItemRepository  = (*RepositoryImpl)(nil)
	_ InvoicePaymentLinkRepository = (*RepositoryImpl)(nil)
	_ InvoicePaymentRepository   = (*RepositoryImpl)(nil)
	_ InvoiceStatusEventRepository = (*RepositoryImpl)(nil)
)
