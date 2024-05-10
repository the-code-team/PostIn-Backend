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
	bus.Handle(&queries.GetProfileQuery{}, cbus.HandlerFunc(GetProfileHandler))
}

func GetProfileHandler(ctx context.Context, command cbus.Command) (interface{}, error) {
	// Get the providers
	db := providers.GetDatabase()
	cache := providers.GetQueryCacheProvider()

	// Get the profile
	result := models.Profile{}

	// Get the profile from the database
	query := cache.Wrap(
		db.Model(&models.Profile{}).
			Where("Email = ?", command.(*queries.GetProfileQuery).Email),
	)

	// Return the profile into the result variable
	query = query.First(&result)

	if query.Error != nil {
		return nil, query.Error
	}

	// Return the profile
	return result, nil
}
