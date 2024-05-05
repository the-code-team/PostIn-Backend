package models

type Profile struct {
	UserId         string `gorm:"primaryKey"`
	Email          string
	FirstName      string
	LastName       string
	ProfilePicture string

	Tag []Tag
}
