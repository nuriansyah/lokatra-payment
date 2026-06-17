package shared

import (
	"time"

	"github.com/gofrs/uuid"
	"github.com/guregu/null"
)

type MetaSignature struct {
	MetaCreatedAt time.Time  `db:"meta_created_at"`
	MetaCreatedBy uuid.UUID  `db:"meta_created_by"`
	MetaUpdatedAt null.Time  `db:"meta_updated_at"`
	MetaUpdatedBy *uuid.UUID `db:"meta_updated_by"`
	MetaDeletedAt null.Time  `db:"meta_deleted_at"`
	MetaDeletedBy *uuid.UUID `db:"meta_deleted_by"`
}

func (s *MetaSignature) SetSignatureMetaUpdate(userId uuid.UUID) *MetaSignature {
	s.MetaUpdatedAt = null.TimeFrom(time.Now())
	s.MetaUpdatedBy = &userId
	return s
}

func (s *MetaSignature) SetSignatureMetaCreate(userId uuid.UUID) *MetaSignature {
	s.MetaCreatedAt = time.Now()
	s.MetaCreatedBy = userId
	return s
}

func (s *MetaSignature) SetSignatureMetaDelete(userId uuid.UUID) {
	s.MetaDeletedAt = null.TimeFrom(time.Now())
	s.MetaDeletedBy = &userId
}
