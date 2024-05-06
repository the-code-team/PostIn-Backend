package models

type ProposeStatus string

const (
	Accepted ProposeStatus = "accepted"
	Rejected ProposeStatus = "rejected"
	Pending  ProposeStatus = "pending"
)

type Propose struct {
	UserId  string        `gorm:"primaryKey"`
	EventId string        `gorm:"primaryKey"`
	Status  ProposeStatus `gorm:"default:pending"`

	Profile Profile `gorm:"foreignKey:UserId"`
	Event   Event   `gorm:"foreignKey:EventId"`
}
