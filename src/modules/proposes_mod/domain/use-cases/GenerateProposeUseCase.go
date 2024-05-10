package use_cases

import (
	"context"
	"epsa.upv.es/postin_backend/src/modules/proposes_mod/domain/commands"
	"epsa.upv.es/postin_backend/src/providers"
	"github.com/gogolfing/cbus"
)

func GenerateProposeUseCase() {
	bus := providers.GetCommandBus()
	bus.Handle(&commands.GenerateProposeCommand{}, cbus.HandlerFunc(GenerateProposeHandler))
}

func GenerateProposeHandler(ctx context.Context, command cbus.Command) (interface{}, error) {
	// Get the providers
	db := providers.GetDatabase()

	// Add the proposal status
	command.(*commands.GenerateProposeCommand).Status = "Pending"

	// Generate the proposal
	query := db.Create(&commands.GenerateProposeCommand{})

	if query.Error != nil {
		return nil, query.Error
	}

	// The proposal has been generated successfully
	return nil, nil
}
