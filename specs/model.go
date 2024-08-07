package specs

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type PaymentModel struct {
	gorm.Model
	PaymentGateaway  string        `gorm:"type:text;not null"`
	UniqueKey        uuid.UUID     `gorm:"type:uuid;not null"`
	OrderName        string        `gorm:"type:text"`
	OrderDescription string        `gorm:"type:text"`
	Amount           int64         `gorm:"type:decimal;not null"`
	Currency         string        `gorm:"type:text;not null"`
	Status           PaymentStatus `gorm:"type:integer;not null;default:0"`
	ErrorMsg         string        `gorm:"type:text"`
	Metadata         string        `gorm:"type:json"`
	RawRequest       string        `gorm:"type:json;not null"`
}
