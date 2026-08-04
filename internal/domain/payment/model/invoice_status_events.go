package model

import (
	"time"

	"github.com/gofrs/uuid"
)

type InvoiceStatusEventsDBFieldNameType string

type invoiceStatusEventsDBFieldName struct {
	Id            InvoiceStatusEventsDBFieldNameType
	InvoiceId     InvoiceStatusEventsDBFieldNameType
	EventType     InvoiceStatusEventsDBFieldNameType
	OldStatus     InvoiceStatusEventsDBFieldNameType
	NewStatus     InvoiceStatusEventsDBFieldNameType
	Reason        InvoiceStatusEventsDBFieldNameType
	ActorId       InvoiceStatusEventsDBFieldNameType
	MetaCreatedAt InvoiceStatusEventsDBFieldNameType
	MetaCreatedBy InvoiceStatusEventsDBFieldNameType
}

var InvoiceStatusEventsDBFieldName = invoiceStatusEventsDBFieldName{
	Id:            "id",
	InvoiceId:     "invoice_id",
	EventType:     "event_type",
	OldStatus:     "old_status",
	NewStatus:     "new_status",
	Reason:        "reason",
	ActorId:       "actor_id",
	MetaCreatedAt: "meta_created_at",
	MetaCreatedBy: "meta_created_by",
}

// InvoiceStatusEvent is append-only: no MetaSignature (no update/delete columns)
type InvoiceStatusEvents struct {
	Id        uuid.UUID `db:"id"`
	InvoiceId uuid.UUID `db:"invoice_id"`
	EventType string    `db:"event_type"`
	OldStatus string    `db:"old_status"`
	NewStatus string    `db:"new_status"`
	Reason    string    `db:"reason"`
	ActorId   uuid.UUID `db:"actor_id"`

	MetaCreatedAt time.Time `db:"meta_created_at"`
	MetaCreatedBy uuid.UUID `db:"meta_created_by"`
}

type InvoiceStatusEventsPrimaryID struct {
	Id uuid.UUID `db:"id"`
}

func (d InvoiceStatusEvents) ToInvoiceStatusEventsPrimaryID() InvoiceStatusEventsPrimaryID {
	return InvoiceStatusEventsPrimaryID{Id: d.Id}
}

type InvoiceStatusEventsList []*InvoiceStatusEvents

type InvoiceStatusEventsFilterResult struct {
	InvoiceStatusEvents
}

type InvoiceStatusEventsFilterResultList []InvoiceStatusEventsFilterResult
