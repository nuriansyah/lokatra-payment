package shared

import (
	"time"

	"github.com/guregu/null"
)

type Signature struct {
	CreatedAt time.Time `db:"created_at"`
	CreatedBy int64     `db:"created_by"`
	UpdatedAt null.Time `db:"updated_at"`
	UpdatedBy null.Int  `db:"updated_by"`
	DeletedAt null.Time `db:"deleted_at"`
	DeletedBy null.Int  `db:"deleted_by"`
}

func (s *Signature) SetSignatureUpdate(userId int64) {
	s.UpdatedAt = null.TimeFrom(time.Now())
	s.UpdatedBy = null.IntFrom(userId)
}

func (s *Signature) SetSignatureCreate(userId int64) {
	s.CreatedAt = time.Now()
	s.CreatedBy = userId
}

func (s *Signature) SetSignatureDelete(userId int64) {
	s.DeletedAt = null.TimeFrom(time.Now())
	s.DeletedBy = null.IntFrom(userId)
}
