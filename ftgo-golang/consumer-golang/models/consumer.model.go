package models

import (
	"github.com/google/uuid"
)

type Consumer struct {
	ID      uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key" json:"id,omitempty"`
	Name    string    `gorm:"uniqueIndex;not null" json:"name,omitempty"`
	Address string    `gorm:"not null" json:"address,omitempty"`
	Email   string    `gorm:"not null" json:"email,omitempty"`
}
