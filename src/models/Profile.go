package models

type Profile struct {
	UserId         string `gorm:"primaryKey" json:"user_id,omitempty"`
	Email          string `json:"email,omitempty"`
	FirstName      string `json:"first_name,omitempty"`
	LastName       string `json:"last_name,omitempty"`
	ProfilePicture string `json:"profile_picture,omitempty"`

	Tag []Tag `json:"tag,omitempty"`
}
