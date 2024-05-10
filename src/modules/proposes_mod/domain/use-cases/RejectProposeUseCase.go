package use_cases

import (
	"context"
	"epsa.upv.es/postin_backend/src/modules/proposes_mod/domain/commands"
	"epsa.upv.es/postin_backend/src/providers"
	"github.com/gogolfing/cbus"
)

func RejectProposeUseCase() {
	bus := providers.GetCommandBus()
	bus.Handle(&commands.RejectProposeCommand{}, cbus.HandlerFunc(RejectProposeHandler))
}

func RejectProposeHandler(ctx context.Context, command cbus.Command) (interface{}, error) {
	// Get the providers
	db := providers.GetDatabase()

	// Reject the propose
	query := db.Model(&commands.RejectProposeCommand{}).
		Where("UserId = ? AND EventId = ?", command.(*commands.RejectProposeCommand).UserId, command.(*commands.RejectProposeCommand).EventId).
		Update("Status", "Rejected")

	if query.Error != nil {
		return nil, query.Error
	}

	// The proposal has been rejected successfully
	return nil, nil
}
