package usecases

import (
	"context"
	"github.com/gogolfing/cbus"

	"epsa.upv.es/postin_backend/src/models"
	"epsa.upv.es/postin_backend/src/modules/profile_mod/domain/queries"
	"epsa.upv.es/postin_backend/src/providers"
)

func GetProfileUseCase() {
	bus := providers.GetCommandBus()

	bus.Handle(&queries.GetProfileQuery{}, cbus.HandlerFunc(
		func(ctx context.Context, command cbus.Command) (interface{}, error) {
			db := providers.GetDatabase()
			result := models.Profile{}

			db.Model(&models.Profile{}).Where("Email = ?", command.(*queries.GetProfileQuery).Email).First(&result)

			return result, nil
		}),
	)
}
