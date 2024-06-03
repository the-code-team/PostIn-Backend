package commands

import "github.com/google/uuid"

type PublishMessageCommand struct {
	UserId  uuid.UUID
	EventId uuid.UUID
	Content string
}

func (q *PublishMessageCommand) Type() string {
	return "PublishMessageCommand"
}
