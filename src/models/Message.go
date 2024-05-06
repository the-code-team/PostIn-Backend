package models

import (
	"github.com/google/uuid"
	"time"
)

type Message struct {
	UserId    string    `gorm:"primaryKey"`
	EventId   string    `gorm:"primaryKey"`
	MessageId uuid.UUID `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	CreatedAt time.Time
	Content   string

	Profile Profile `gorm:"foreignKey:UserId"`
	Event   Event   `gorm:"foreignKey:EventId"`
}
