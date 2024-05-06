package usecases

import (
	"context"
	"epsa.upv.es/postin_backend/src/models"
	"epsa.upv.es/postin_backend/src/modules/profile_mod/domain/commands"
	"epsa.upv.es/postin_backend/src/providers"
	"github.com/gogolfing/cbus"
)

func UpsertProfileUseCase() {
	bus := providers.GetCommandBus()
	bus.Handle(&commands.UpsertProfileCommand{}, cbus.HandlerFunc(UpsertProfileHandler))
}

func UpsertProfileHandler(ctx context.Context, command cbus.Command) (interface{}, error) {
	// Get the providers
	db := providers.GetDatabase()

	// Save the profile
	query := db.Model(&models.Profile{}).Save(&command)

	if query.Error != nil {
		return nil, query.Error
	}

	// The profile has been saved successfully
	return nil, nil
}
