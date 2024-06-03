package queries

import "github.com/google/uuid"

type GetMessagesQuery struct {
	EventId uuid.UUID
}

func (q *GetMessagesQuery) Type() string {
	return "GetMessagesQuery"
}
