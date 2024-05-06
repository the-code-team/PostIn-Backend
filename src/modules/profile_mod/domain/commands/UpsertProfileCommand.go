package commands

import "epsa.upv.es/postin_backend/src/models"

type UpsertProfileCommand struct {
	models.Profile
}

func (q *UpsertProfileCommand) Type() string {
	return "UpsertProfileCommand"
}
