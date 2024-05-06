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

	bus.Handle(&commands.UpsertProfileCommand{}, cbus.HandlerFunc(
		func(ctx context.Context, command cbus.Command) (interface{}, error) {
			db := providers.GetDatabase()
			db.Model(&models.Profile{}).Save(&command)

			return nil, nil
		}),
	)
}
