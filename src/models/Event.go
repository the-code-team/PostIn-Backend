package models

import (
	"github.com/google/uuid"
	"time"
)

type Event struct {
	EventId    uuid.UUID `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	HostUserId string    `json:"host_user_id","omitempty"`

	Title       string `json:"title","omitempty"`
	Description string `json:"description","omitempty"`

	Price 	    float32

	LocationName string
	Latitude  float64 `json:"latitude","omitempty"`
	Longitude float64 `json:"longitude","omitempty"`

	DateEvent             time.Time `json:"date_event","omitempty"`
	OpenInscriptionsUntil time.Time `json:"open_inscriptions_until","omitempty"`

	Tags      []Tag    `json:"tags","omitempty"`
	PhotosUri []string `json:"photos_uri","omitempty"`

	Profile Profile `gorm:"foreignKey:HostUserId" json:"profile","omitempty"`
}
