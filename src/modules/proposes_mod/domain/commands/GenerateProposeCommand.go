package commands

import "epsa.upv.es/postin_backend/src/models"

type GenerateProposeCommand struct {
	models.Propose
}

func (c *GenerateProposeCommand) Type() string {
	return "GenerateProposeCommand"
}
