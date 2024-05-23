package use_cases

import (
	"context"
	"epsa.upv.es/postin_backend/src/models"
	"epsa.upv.es/postin_backend/src/modules/events_mod/domain/queries"
	"epsa.upv.es/postin_backend/src/providers"
	"github.com/gogolfing/cbus"
)

func ListNearbyEventsUseCase() {
	bus := providers.GetCommandBus()
	bus.Handle(&queries.ListNearbyEventsQuery{}, cbus.HandlerFunc(ListNearbyEventsHandler))
}

func ListNearbyEventsHandler(ctx context.Context, command cbus.Command) (interface{}, error) {
	// Get the providers
	db := providers.GetDatabase()
	cache := providers.GetQueryCacheProvider()

	// Get the events
	events := &models.Event{}

	// Get the events from the database
	query := cache.Wrap(
		db.Exec(""+
			"SELECT * FROM Events WHERE earth_distance(ll_to_earth(Events.Latitude, Events.Longitude), ll_to_earth(:Latitude, :Longitude)) < :MaxDistance;",
			map[string]interface{}{
				"Latitude":    command.(*queries.ListNearbyEventsQuery).Latitude,
				"Longitude":   command.(*queries.ListNearbyEventsQuery).Longitude,
				"MaxDistance": command.(*queries.ListNearbyEventsQuery).MaxDistance,
			}),
	)

	// Return the events into the result variable
	query = query.First(&events)

	if query.Error != nil {
		return nil, query.Error
	}

	// The events have been retrieved successfully
	return events, nil
}
