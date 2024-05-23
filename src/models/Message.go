package models

import (
	"github.com/google/uuid"
	"time"
)

type Message struct {
	UserId    string    `gorm:"primaryKey" json:"user_id,omitempty"`
	EventId   string    `gorm:"primaryKey" json:"event_id,omitempty"`
	MessageId uuid.UUID `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()" json:"message_id,omitempty"`

	CreatedAt time.Time `json:"created_at,omitempty"`
	Content   string    `json:"content,omitempty"`

	Profile Profile `gorm:"foreignKey:UserId"`
	Event   Event   `gorm:"foreignKey:EventId"`
}
