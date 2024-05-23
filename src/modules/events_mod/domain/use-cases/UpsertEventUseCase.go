package use_cases

import (
	"context"
	"epsa.upv.es/postin_backend/src/models"
	"epsa.upv.es/postin_backend/src/modules/events_mod/domain/commands"
	"epsa.upv.es/postin_backend/src/providers"
	"github.com/gogolfing/cbus"
)

func UpsertEventUseCase() {
	bus := providers.GetCommandBus()
	bus.Handle(&commands.UpsertEventCommand{}, cbus.HandlerFunc(PublishEventHandler))
}

func PublishEventHandler(ctx context.Context, command cbus.Command) (interface{}, error) {
	// Get the providers
	db := providers.GetDatabase()

	// Publish the event
	query := db.Model(&models.Event{}).Save(&command)

	if query.Error != nil {
		return nil, query.Error
	}

	// The event has been published successfully
	return nil, nil
}
