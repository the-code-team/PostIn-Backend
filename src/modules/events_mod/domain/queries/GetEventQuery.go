package queries

import "github.com/google/uuid"

type GetEventQuery struct {
	EventId uuid.UUID
}

func (q *GetEventQuery) Type() string {
	return "GetEventQuery"
}
