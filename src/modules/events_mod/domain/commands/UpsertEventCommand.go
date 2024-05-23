package commands

import "epsa.upv.es/postin_backend/src/models"

type UpsertEventCommand struct {
	*models.Event
}

func (q *UpsertEventCommand) Type() string {
	return "UpsertEventCommand"
}
