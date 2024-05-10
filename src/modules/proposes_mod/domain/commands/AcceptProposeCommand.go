package commands

import "github.com/google/uuid"

type AcceptProposeCommand struct {
	UserId  string
	EventId uuid.UUID
}

func (c *AcceptProposeCommand) Type() string {
	return "AcceptProposeCommand"
}
