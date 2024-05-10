package models

import "github.com/google/uuid"

type StatusChangeNotification struct {
	UserId  string
	EventId uuid.UUID
	Message string
	Status  string
}
