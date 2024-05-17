package models

import "github.com/google/uuid"

/*
INFO: Below is the Prefered way to query the events table:
SELECT *
FROM Events
WHERE earth_distance(

	ll_to_earth(Events.Latitude, Events.Longitude),
	ll_to_earth(:Latitude, :Longitude),

) < :Radius;
*/
type Event struct {
	EventId    uuid.UUID `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	HostUserId string

	Title       string
	Description string

	Price 	    float32

	LocationName string
	Latitude     float64
	Longitude    float64

	Tags      []Tag
	PhotosUri []string

	Profile Profile `gorm:"foreignKey:HostUserId"`
}
