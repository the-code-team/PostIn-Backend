package use_cases

import (
	"context"
	"epsa.upv.es/postin_backend/src/models"
	"epsa.upv.es/postin_backend/src/modules/messages_mod/domain/commands"
	"epsa.upv.es/postin_backend/src/providers"
	"fmt"
	"github.com/gogolfing/cbus"
)

func PublishMessageUseCase() {
	bus := providers.GetCommandBus()
	bus.Handle(&commands.PublishMessageCommand{}, cbus.HandlerFunc(PublishMessageHandle))
}

func PublishMessageHandle(ctx context.Context, command cbus.Command) (interface{}, error) {
	// Get the providers
	db := providers.GetDatabase()
	eventBus := providers.GetEventBusProvider()

	// Add content with ids to database
	query := db.Model(&models.Message{}).Save(&command)

	if query.Error != nil {
		return nil, query.Error
	}

	err := eventBus.Publish(
		fmt.Sprintf("message|%s", command.(*commands.PublishMessageCommand).EventId),
		command.(*commands.PublishMessageCommand).Content,
	)

	if err != nil {
		return nil, err
	}

	// The message has been added successfully
	return nil, nil
}
