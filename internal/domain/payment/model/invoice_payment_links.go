package model

import (
	"github.com/gofrs/uuid"
	"github.com/guregu/null"
	"github.com/nuriansyah/lokatra-payment/shared"
)

type InvoicePaymentLinksDBFieldNameType string

type invoicePaymentLinksDBFieldName struct {
	Id            InvoicePaymentLinksDBFieldNameType
	InvoiceId     InvoicePaymentLinksDBFieldNameType
	ProviderCode  InvoicePaymentLinksDBFieldNameType
	LinkUrl       InvoicePaymentLinksDBFieldNameType
	LinkType      InvoicePaymentLinksDBFieldNameType
	Status        InvoicePaymentLinksDBFieldNameType
	ClickCount    InvoicePaymentLinksDBFieldNameType
	PaymentCount  InvoicePaymentLinksDBFieldNameType
	ExpiresAt     InvoicePaymentLinksDBFieldNameType
	MetaCreatedAt InvoicePaymentLinksDBFieldNameType
	MetaCreatedBy InvoicePaymentLinksDBFieldNameType
	MetaUpdatedAt InvoicePaymentLinksDBFieldNameType
	MetaUpdatedBy InvoicePaymentLinksDBFieldNameType
	MetaDeletedAt InvoicePaymentLinksDBFieldNameType
	MetaDeletedBy InvoicePaymentLinksDBFieldNameType
}

var InvoicePaymentLinksDBFieldName = invoicePaymentLinksDBFieldName{
	Id:            "id",
	InvoiceId:     "invoice_id",
	ProviderCode:  "provider_code",
	LinkUrl:       "link_url",
	LinkType:      "link_type",
	Status:        "status",
	ClickCount:    "click_count",
	PaymentCount:  "payment_count",
	ExpiresAt:     "expires_at",
	MetaCreatedAt: "meta_created_at",
	MetaCreatedBy: "meta_created_by",
	MetaUpdatedAt: "meta_updated_at",
	MetaUpdatedBy: "meta_updated_by",
	MetaDeletedAt: "meta_deleted_at",
	MetaDeletedBy: "meta_deleted_by",
}

type PaymentLinkStatus string

const (
	PaymentLinkStatusActive   PaymentLinkStatus = "active"
	PaymentLinkStatusDisabled PaymentLinkStatus = "disabled"
	PaymentLinkStatusExpired  PaymentLinkStatus = "expired"
)

type InvoicePaymentLinks struct {
	Id           uuid.UUID        `db:"id"`
	InvoiceId    uuid.UUID        `db:"invoice_id"`
	ProviderCode string           `db:"provider_code"`
	LinkUrl      string           `db:"link_url"`
	LinkType     string           `db:"link_type"`
	Status       PaymentLinkStatus `db:"status"`
	ClickCount   int              `db:"click_count"`
	PaymentCount int              `db:"payment_count"`
	ExpiresAt    null.Time        `db:"expires_at"`

	shared.MetaSignature
}

type InvoicePaymentLinksPrimaryID struct {
	Id uuid.UUID `db:"id"`
}

func (d InvoicePaymentLinks) ToInvoicePaymentLinksPrimaryID() InvoicePaymentLinksPrimaryID {
	return InvoicePaymentLinksPrimaryID{Id: d.Id}
}

type InvoicePaymentLinksList []*InvoicePaymentLinks

type InvoicePaymentLinksFilterResult struct {
	InvoicePaymentLinks
}

type InvoicePaymentLinksFilterResultList []InvoicePaymentLinksFilterResult
