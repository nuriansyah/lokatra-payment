package service

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/gofrs/uuid"
	"github.com/guregu/null"
	"github.com/rs/zerolog"

	pg "github.com/nuriansyah/lokatra-payment/external/paymentgateway"
	paymentmodel "github.com/nuriansyah/lokatra-payment/internal/domain/payment/model"
	"github.com/nuriansyah/lokatra-payment/internal/domain/payment/model/dto"
	"github.com/nuriansyah/lokatra-payment/shared"
	"github.com/nuriansyah/lokatra-payment/shared/failure"
	"github.com/nuriansyah/lokatra-payment/shared/nuuid"
	"github.com/shopspring/decimal"
)

func (s *ServiceImpl) log() *zerolog.Logger {
	return &s.logger
}

// ─── Create Invoice ──────────────────────────────────────────────────────────

func (s *ServiceImpl) CreateInvoice(ctx context.Context, req dto.CreateInvoiceRequest) (dto.InvoiceDetailResponse, error) {
	currency := strings.ToUpper(strings.TrimSpace(req.Currency))
	pph23Rate := decimal.Zero
	ppnRate := decimal.NewFromFloat(0.11)
	if req.Pph23Rate != nil {
		pph23Rate = *req.Pph23Rate
	}
	if req.PpnRate != nil {
		ppnRate = *req.PpnRate
	}

	invoiceID := uuid.Must(uuid.NewV7())
	lineItems := make([]paymentmodel.InvoiceLineItems, 0, len(req.LineItems))
	var subtotal, totalDiscount, totalPpn, totalPph23 decimal.Decimal
	now := time.Now().UTC()

	for i, li := range req.LineItems {
		lineSubtotal := li.Quantity.Mul(li.UnitPrice)
		var discountAmt decimal.Decimal
		if li.DiscountPercent != nil && li.DiscountPercent.GreaterThan(decimal.Zero) {
			discountAmt = lineSubtotal.Mul(*li.DiscountPercent).Div(decimal.NewFromInt(100))
		}
		lineAfterDiscount := lineSubtotal.Sub(discountAmt)

		linePph23Rate := pph23Rate
		if li.Pph23Rate != nil {
			linePph23Rate = *li.Pph23Rate
		}
		linePpnRate := ppnRate
		if li.PpnRate != nil {
			linePpnRate = *li.PpnRate
		}
		linePph23Amt := lineAfterDiscount.Mul(linePph23Rate).Div(decimal.NewFromInt(100))
		linePpnAmt := lineAfterDiscount.Mul(linePpnRate).Div(decimal.NewFromInt(100))
		lineTotal := lineAfterDiscount.Add(linePpnAmt).Sub(linePph23Amt)

		subtotal = subtotal.Add(lineSubtotal)
		totalDiscount = totalDiscount.Add(discountAmt)
		totalPph23 = totalPph23.Add(linePph23Amt)
		totalPpn = totalPpn.Add(linePpnAmt)

		lineItemID := uuid.Must(uuid.NewV7())
		lineItems = append(lineItems, paymentmodel.InvoiceLineItems{
			Id:              lineItemID,
			InvoiceId:       invoiceID,
			LineNo:          i + 1,
			Name:            li.Name,
			Description:     null.StringFromPtr(&li.Description),
			Quantity:        li.Quantity,
			UnitPrice:       li.UnitPrice,
			DiscountPercent: orZero(li.DiscountPercent),
			DiscountAmount:  discountAmt,
			PpnRate:         linePpnRate,
			PpnAmount:       linePpnAmt,
			Pph23Rate:       linePph23Rate,
			Pph23Amount:     linePph23Amt,
			Subtotal:        lineSubtotal,
			TotalAmount:     lineTotal,
			Currency:        currency,
			Sku:             null.StringFromPtr(&li.SKU),
			Category:        null.StringFromPtr(&li.Category),
			MetaSignature: shared.MetaSignature{
				MetaCreatedAt: now,
				MetaCreatedBy: req.ActorID,
			},
		})
	}

	totalAmount := subtotal.Sub(totalDiscount).Add(totalPpn).Sub(totalPph23)

	inv := paymentmodel.Invoices{
		Id:              invoiceID,
		InvoiceCode:     operationCode("inv"),
		MerchantId:      req.MerchantID,
		CustomerId:      nuuid.From(req.CustomerID),
		CustomerName:    null.StringFrom(req.CustomerName),
		CustomerEmail:   null.StringFrom(req.CustomerEmail),
		CustomerPhone:   null.StringFrom(req.CustomerPhone),
		Subtotal:        subtotal,
		DiscountAmount:  totalDiscount,
		Pph23Amount:     totalPph23,
		Pph23Rate:       pph23Rate,
		PpnAmount:       totalPpn,
		PpnRate:         ppnRate,
		TotalAmount:     totalAmount,
		PaidAmount:      decimal.Zero,
		RemainingAmount: totalAmount,
		Currency:        currency,
		Status:          paymentmodel.InvoiceStatusDraft,
		Description:     null.StringFrom(req.Description),
		Notes:           null.StringFrom(req.Notes),
		Terms:           null.StringFrom(req.Terms),
		MetaSignature: shared.MetaSignature{
			MetaCreatedAt: now,
			MetaCreatedBy: req.ActorID,
		},
	}
	if req.DueAt != nil {
		inv.DueAt = null.TimeFrom(req.DueAt.UTC())
	}

	if err := s.paymentRepo.CreateInvoice(ctx, &inv, lineItems); err != nil {
		return dto.InvoiceDetailResponse{}, err
	}

	s.recordStatusEvent(ctx, invoiceID, "created", "", string(paymentmodel.InvoiceStatusDraft), "invoice created", req.ActorID)

	resp := dto.NewInvoiceResponse(inv)
	lineItemResp := make([]dto.InvoiceLineItemResponse, len(lineItems))
	for i, li := range lineItems {
		lineItemResp[i] = dto.NewInvoiceLineItemResponse(li)
	}
	return dto.InvoiceDetailResponse{
		InvoiceResponse: resp,
		LineItems:       lineItemResp,
	}, nil
}

// ─── Get Invoice ─────────────────────────────────────────────────────────────

func (s *ServiceImpl) GetInvoice(ctx context.Context, id uuid.UUID) (dto.InvoiceDetailResponse, error) {
	if id == uuid.Nil {
		return dto.InvoiceDetailResponse{}, failure.WithCode(shared.ErrInvalidID, "invoice id is required")
	}

	inv, err := s.paymentRepo.ResolveInvoiceByID(ctx, id)
	if err != nil {
		return dto.InvoiceDetailResponse{}, err
	}

	items, _ := s.paymentRepo.ResolveInvoiceLineItemsByInvoiceID(ctx, id)
	links, _ := s.paymentRepo.ResolveInvoicePaymentLinksByInvoiceID(ctx, id)
	pays, _ := s.paymentRepo.ResolveInvoicePaymentsByInvoiceID(ctx, id)

	lineItemResp := make([]dto.InvoiceLineItemResponse, len(items))
	for i, item := range items {
		lineItemResp[i] = dto.NewInvoiceLineItemResponse(item)
	}
	linkResp := make([]dto.InvoicePaymentLinkResponse, len(links))
	for i, link := range links {
		linkResp[i] = dto.NewInvoicePaymentLinkResponse(link)
	}
	payResp := make([]dto.InvoicePaymentResponse, len(pays))
	for i, pay := range pays {
		payResp[i] = dto.NewInvoicePaymentResponse(pay)
	}

	return dto.InvoiceDetailResponse{
		InvoiceResponse: dto.NewInvoiceResponse(*inv),
		LineItems:       lineItemResp,
		PaymentLinks:    linkResp,
		Payments:        payResp,
	}, nil
}

// ─── List Invoices ───────────────────────────────────────────────────────────

func (s *ServiceImpl) ListInvoices(ctx context.Context, filter paymentmodel.Filter) ([]dto.InvoiceResponse, int, error) {
	results, err := s.paymentRepo.ResolveInvoicesByFilter(ctx, filter)
	if err != nil {
		return nil, 0, err
	}

	resp := make([]dto.InvoiceResponse, len(results))
	total := 0
	for i, r := range results {
		resp[i] = dto.NewInvoiceResponse(r.Invoices)
		if i == 0 {
			total = len(results)
		}
	}
	return resp, total, nil
}

// ─── Issue Invoice ───────────────────────────────────────────────────────────

func (s *ServiceImpl) IssueInvoice(ctx context.Context, id uuid.UUID, actorID uuid.UUID) (dto.InvoiceResponse, error) {
	inv, err := s.paymentRepo.ResolveInvoiceByID(ctx, id)
	if err != nil {
		return dto.InvoiceResponse{}, err
	}
	if inv.Status != paymentmodel.InvoiceStatusDraft {
		return dto.InvoiceResponse{}, failure.Conflict("issue", "invoice", fmt.Sprintf("cannot issue invoice from status %s", inv.Status))
	}
	now := time.Now().UTC()
	inv.Status = paymentmodel.InvoiceStatusIssued
	inv.IssuedAt = null.TimeFrom(now)
	inv.SetSignatureMetaUpdate(actorID)

	if err := s.paymentRepo.UpdateInvoiceByID(ctx, id, inv); err != nil {
		return dto.InvoiceResponse{}, err
	}

	s.recordStatusEvent(ctx, id, "status_change", string(paymentmodel.InvoiceStatusDraft), string(paymentmodel.InvoiceStatusIssued), "invoice issued", actorID)
	return dto.NewInvoiceResponse(*inv), nil
}

// ─── Void Invoice ────────────────────────────────────────────────────────────

func (s *ServiceImpl) VoidInvoice(ctx context.Context, id uuid.UUID, req dto.VoidInvoiceRequest) (dto.InvoiceResponse, error) {
	inv, err := s.paymentRepo.ResolveInvoiceByID(ctx, id)
	if err != nil {
		return dto.InvoiceResponse{}, err
	}
	if inv.Status == paymentmodel.InvoiceStatusVoided {
		return dto.InvoiceResponse{}, failure.Conflict("void", "invoice", "invoice is already voided")
	}
	if inv.Status == paymentmodel.InvoiceStatusPaid || inv.Status == paymentmodel.InvoiceStatusWrittenOff {
		return dto.InvoiceResponse{}, failure.Conflict("void", "invoice", fmt.Sprintf("cannot void invoice from status %s", inv.Status))
	}
	if inv.PaidAmount.GreaterThan(decimal.Zero) {
		return dto.InvoiceResponse{}, failure.Conflict("void", "invoice", "cannot void invoice with existing payments; issue a refund instead")
	}

	now := time.Now().UTC()
	inv.Status = paymentmodel.InvoiceStatusVoided
	inv.VoidedAt = null.TimeFrom(now)
	inv.VoidReason = null.StringFrom(req.Reason)
	inv.SetSignatureMetaUpdate(req.ActorID)

	if err := s.paymentRepo.UpdateInvoiceByID(ctx, id, inv); err != nil {
		return dto.InvoiceResponse{}, err
	}

	s.recordStatusEvent(ctx, id, "status_change", "draft/issued", string(paymentmodel.InvoiceStatusVoided), req.Reason, req.ActorID)
	return dto.NewInvoiceResponse(*inv), nil
}

func (s *ServiceImpl) RecordInvoicePayment(ctx context.Context, req dto.RecordPaymentRequest) (dto.InvoicePaymentResponse, error) {
	inv, err := s.paymentRepo.ResolveInvoiceByID(ctx, req.ActorID)
	if err != nil {
		return dto.InvoicePaymentResponse{}, err
	}
	_ = inv

	invoices, err := s.paymentRepo.ResolveInvoicesByFilter(ctx, paymentmodel.Filter{
		FilterFields: []paymentmodel.FilterField{
			{Field: "merchant_id", Operator: paymentmodel.OperatorEqual, Value: req.ActorID},
		},
		Pagination: paymentmodel.Pagination{Page: 1, PageSize: 100},
	})
	if err != nil {
		return dto.InvoicePaymentResponse{}, err
	}
	_ = invoices

	now := time.Now().UTC()
	pay := paymentmodel.InvoicePayments{
		Id:              uuid.Must(uuid.NewV7()),
		PaymentIntentId: req.PaymentIntentID,
		ProviderCode:    req.ProviderCode,
		Amount:          req.Amount,
		Currency:        strings.ToUpper(strings.TrimSpace(req.Currency)),
		Status:          paymentmodel.InvoicePaymentStatusPending,
		MetaSignature: shared.MetaSignature{
			MetaCreatedAt: now,
			MetaCreatedBy: req.ActorID,
		},
	}

	_ = pay
	return dto.InvoicePaymentResponse{}, nil
}

func (s *ServiceImpl) GenerateInvoicePaymentLinks(ctx context.Context, id uuid.UUID, req dto.GeneratePaymentLinksRequest) ([]dto.InvoicePaymentLinkResponse, error) {
	inv, err := s.paymentRepo.ResolveInvoiceByID(ctx, id)
	if err != nil {
		return nil, err
	}
	if inv.Status != paymentmodel.InvoiceStatusIssued && inv.Status != paymentmodel.InvoiceStatusPartiallyPaid {
		return nil, failure.Conflict("generate_payment_links", "invoice", fmt.Sprintf("cannot generate links from status %s", inv.Status))
	}

	var links []dto.InvoicePaymentLinkResponse
	now := time.Now().UTC()

	for _, providerCode := range req.Provider {
		gw, err := s.gatewayRegistry.Get(pg.ProviderCode(providerCode))
		if err != nil {
			s.log().Warn().Str("provider", providerCode).Err(err).Msg("provider not available for payment link generation")
			continue
		}
		creator, ok := gw.(pg.PaymentLinkCreator)
		if !ok {
			s.log().Warn().Str("provider", providerCode).Msg("provider does not implement PaymentLinkCreator")
			continue
		}

		linkURL, expiresAt, err := creator.CreatePaymentLink(ctx, pg.CreatePaymentLinkRequest{
			InvoiceID:   id.String(),
			Amount:      inv.TotalAmount.StringFixed(2),
			Currency:    inv.Currency,
			CustomerID:  inv.CustomerId.UUID.String(),
			Description: orStr(inv.Description.String, inv.InvoiceCode),
		})
		if err != nil {
			s.log().Error().Err(err).Str("provider", providerCode).Msg("failed to create payment link")
			continue
		}

		linkID := uuid.Must(uuid.NewV7())
		dbLink := paymentmodel.InvoicePaymentLinks{
			Id:           linkID,
			InvoiceId:    id,
			ProviderCode: providerCode,
			LinkUrl:      linkURL,
			LinkType:     "redirect",
			Status:       paymentmodel.PaymentLinkStatusActive,
			MetaSignature: shared.MetaSignature{
				MetaCreatedAt: now,
				MetaCreatedBy: req.ActorID,
			},
		}
		if expiresAt != nil {
			dbLink.ExpiresAt = null.TimeFrom(*expiresAt)
		}

		if err := s.paymentRepo.CreateInvoicePaymentLinks(ctx, []paymentmodel.InvoicePaymentLinks{dbLink}); err != nil {
			continue
		}

		resp := dto.NewInvoicePaymentLinkResponse(dbLink)
		links = append(links, resp)
	}

	return links, nil
}

// ─── Helpers ─────────────────────────────────────────────────────────────────

func (s *ServiceImpl) recordStatusEvent(ctx context.Context, invoiceID uuid.UUID, eventType, oldStatus, newStatus, reason string, actorID uuid.UUID) {
	now := time.Now().UTC()
	event := paymentmodel.InvoiceStatusEvents{
		Id:            uuid.Must(uuid.NewV7()),
		InvoiceId:     invoiceID,
		EventType:     eventType,
		OldStatus:     oldStatus,
		NewStatus:     newStatus,
		Reason:        reason,
		ActorId:       actorID,
		MetaCreatedAt: now,
		MetaCreatedBy: actorID,
	}
	if err := s.paymentRepo.CreateInvoiceStatusEvent(ctx, &event); err != nil {
		s.log().Error().Err(err).Str("invoice_id", invoiceID.String()).Msg("failed to record invoice status event")
	}
}

func orZero(d *decimal.Decimal) decimal.Decimal {
	if d == nil {
		return decimal.Zero
	}
	return *d
}

func orStr(a, b string) string {
	if strings.TrimSpace(a) != "" {
		return a
	}
	return b
}
