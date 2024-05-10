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
	UserId  string        `gorm:"primaryKey"`
	EventId uuid.UUID     `gorm:"primaryKey"`
	Status  ProposeStatus `gorm:"default:pending"`

	Profile Profile `gorm:"foreignKey:UserId"`
	Event   Event   `gorm:"foreignKey:EventId"`
}
