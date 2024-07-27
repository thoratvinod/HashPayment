package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type PaymentStatus uint32

const (
	PaymentStatusNone PaymentStatus = iota
	PaymentStatusSuccess
	PaymentStatusFailed
)

type Payment struct {
	gorm.Model
	UniqueKey   uuid.UUID     `json:"-" gorm:"type:uuid;not null"`
	Amount      float64       `json:"amount" gorm:"type:decimal;not null"`
	Keys        string        `json:"keys" gorm:"type:text;not null"`
	TenantInfo  string        `json:"tenantInfo" gorm:"type:text;not null"`
	WebhookURL  string        `json:"webhookURL" gorm:"type:text;not null"`
	Status      PaymentStatus `json:"-" gorm:"type:integer;not null;default:0"`
	RedirectURL string        `json:"-" gorm:"type:text"`
}
