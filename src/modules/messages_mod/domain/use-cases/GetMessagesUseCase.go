package use_cases

import (
	"context"
	"epsa.upv.es/postin_backend/src/models"
	"epsa.upv.es/postin_backend/src/modules/messages_mod/domain/queries"
	"epsa.upv.es/postin_backend/src/providers"
	"github.com/gogolfing/cbus"
)

func GetMessagesUseCase() {
	bus := providers.GetCommandBus()
	bus.Handle(&queries.GetMessagesQuery{}, cbus.HandlerFunc(GetMessagesHandler))
}

func GetMessagesHandler(ctx context.Context, command cbus.Command) (interface{}, error) {
	// Get the providers
	db := providers.GetDatabase()
	cache := providers.GetQueryCacheProvider()

	// Get the event
	message := &models.Message{}

	// Get the event from the database
	query := cache.Wrap(
		db.Model(&models.Message{}).
			Where("EventId = ?", command.(*queries.GetMessagesQuery).EventId),
	)

	// Return the event into the result variable
	query = query.First(&message)

	if query.Error != nil {
		return nil, query.Error
	}

	// The event has been retrieved successfully
	return message, nil
}
