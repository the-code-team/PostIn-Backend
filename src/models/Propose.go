package models

import (
	"github.com/google/uuid"
)

type ProposeStatus string

const (
	Accepted ProposeStatus = "accepted"
	Rejected ProposeStatus = "rejected"
	Pending  ProposeStatus = "pending"
)

type Propose struct {
	UserId  string        `gorm:"primaryKey" json:"user_id,omitempty"`
	EventId uuid.UUID     `gorm:"primaryKey" json:"event_id,omitempty"`
	Status  ProposeStatus `gorm:"default:pending" json:"status,omitempty"`

	Profile Profile `gorm:"foreignKey:UserId"`
	Event   Event   `gorm:"foreignKey:EventId"`
}
