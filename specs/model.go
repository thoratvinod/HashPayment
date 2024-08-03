package specs

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type PaymentStatus uint32

const (
	PaymentStatusCreated PaymentStatus = iota
	PaymentStatusSuccess
	PaymentStatusCanceled
	PaymentStatusFailed
)

type PaymentModel struct {
	gorm.Model
	UniqueKey   uuid.UUID     `gorm:"type:uuid;not null"`
	Amount      int64         `gorm:"type:decimal;not null"`
	TenantInfo  string        `gorm:"type:text;not null"`
	WebhookURL  string        `gorm:"type:text;not null"`
	Status      PaymentStatus `gorm:"type:integer;not null;default:0"`
	RedirectURL string        `gorm:"type:text"`
	ErrorMsg    string        `gorm:"type:text"`
}
