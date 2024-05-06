package models

import "time"

type ProfilePhoto struct {
	PhotoUri  string
	UpdatedAt *time.Time
}
