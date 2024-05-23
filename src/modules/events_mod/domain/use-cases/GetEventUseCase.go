package use_cases

import (
	"context"
	"epsa.upv.es/postin_backend/src/models"
	"epsa.upv.es/postin_backend/src/modules/events_mod/domain/queries"
	"epsa.upv.es/postin_backend/src/providers"
	"github.com/gogolfing/cbus"
)

func GetEventUseCase() {
	bus := providers.GetCommandBus()
	bus.Handle(&queries.GetEventQuery{}, cbus.HandlerFunc(GetEventHandler))
}

func GetEventHandler(ctx context.Context, command cbus.Command) (interface{}, error) {
	// Get the providers
	db := providers.GetDatabase()
	cache := providers.GetQueryCacheProvider()

	// Get the event
	event := &models.Event{}

	// Get the event from the database
	query := cache.Wrap(
		db.Model(&models.Event{}).
			Where("EventId = ?", command.(*queries.GetEventQuery).EventId),
	)

	// Return the event into the result variable
	query = query.First(&event)

	if query.Error != nil {
		return nil, query.Error
	}

	// The event has been retrieved successfully
	return event, nil
}
