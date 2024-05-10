package use_cases

import (
	"context"
	"encoding/json"
	"epsa.upv.es/postin_backend/src/modules/proposes_mod/domain/commands"
	"epsa.upv.es/postin_backend/src/modules/proposes_mod/domain/models"
	"epsa.upv.es/postin_backend/src/providers"
	"fmt"
	"github.com/gogolfing/cbus"
)

func AcceptProposeUseCase() {
	bus := providers.GetCommandBus()
	bus.Handle(&commands.AcceptProposeCommand{}, cbus.HandlerFunc(AcceptProposeHandler))
}

func AcceptProposeHandler(ctx context.Context, command cbus.Command) (interface{}, error) {
	// Get the providers
	db := providers.GetDatabase()
	eventBus := providers.GetEventBusProvider()

	// Save the UserId locally
	userId := command.(*commands.AcceptProposeCommand).UserId

	// Accept the proposal
	query := db.Model(&commands.AcceptProposeCommand{}).
		Where("UserId = ? AND EventId = ?", userId, command.(*commands.AcceptProposeCommand).EventId).
		Update("Status", "Accepted")

	if query.Error != nil {
		return nil, query.Error
	}

	// Prepare JSON message
	message, err := json.Marshal(models.StatusChangeNotification{
		UserId:  userId,
		EventId: command.(*commands.AcceptProposeCommand).EventId,
		Status:  "Accepted",
		Message: "Your propose has been accepted!",
	})

	if err != nil {
		panic("BUG: failed to marshal message from AcceptProposeUseCase")
	}

	// Publish the message
	err = eventBus.Publish(fmt.Sprintf("notification|%s", userId), string(message))

	if err != nil {
		return nil, err
	}

	// The proposal has been accepted successfully
	return nil, nil
}
