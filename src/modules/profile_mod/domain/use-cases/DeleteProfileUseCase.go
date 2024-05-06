package usecases

import (
	"context"
	"epsa.upv.es/postin_backend/src/models"
	"epsa.upv.es/postin_backend/src/modules/profile_mod/domain/commands"
	"epsa.upv.es/postin_backend/src/modules/profile_mod/domain/queries"
	"epsa.upv.es/postin_backend/src/providers"
	"github.com/gogolfing/cbus"
)

func DeleteProfileUseCase() {
	bus := providers.GetCommandBus()

	bus.Handle(&commands.DeleteProfileCommand{}, cbus.HandlerFunc(
		func(ctx context.Context, command cbus.Command) (interface{}, error) {
			db := providers.GetDatabase()
			db.Model(&models.Profile{}).Where("Email = ?", command.(*queries.GetProfileQuery).Email).Delete(&command)

			return nil, nil
		}),
	)
}
