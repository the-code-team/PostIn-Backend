package commands

import "github.com/google/uuid"

type RejectProposeCommand struct {
	UserId  string
	EventId uuid.UUID
}

func (c *RejectProposeCommand) Type() string {
	return "RejectProposeCommand"
}
