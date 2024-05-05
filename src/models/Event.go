package models

import "github.com/google/uuid"

type Event struct {
	EventId    uuid.UUID `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	HostUserId string

	Title       string
	Description string

	LocationName string
	Latitude     float64
	Longitude    float64

	Tags      []Tag
	PhotosUri []string

	Profile Profile `gorm:"foreignKey:HostUserId"`
}
