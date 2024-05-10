package use_cases

import (
	"context"
	"epsa.upv.es/postin_backend/src/models"
	"epsa.upv.es/postin_backend/src/modules/proposes_mod/domain/queries"
	"epsa.upv.es/postin_backend/src/providers"
	"github.com/gogolfing/cbus"
)

func ListProposesUseCase() {
	bus := providers.GetCommandBus()
	bus.Handle(&queries.ListProposesQuery{}, cbus.HandlerFunc(ListProposesHandler))
}

func ListProposesHandler(ctx context.Context, command cbus.Command) (interface{}, error) {
	// Get the providers
	db := providers.GetDatabase()
	cache := providers.GetQueryCacheProvider()

	// The ListProposesQuery command has multiple EventIds to match
	query := cache.Wrap(
		db.Model(&models.Propose{}).
			Where("EventId IN (?) AND Status = ?", command.(*queries.ListProposesQuery).EventIds, "Pending"),
	)

	// Return the proposes into the result variable
	var proposes []models.Propose

	query = query.Find(&proposes)

	if query.Error != nil {
		return nil, query.Error
	}

	// The proposes have been found successfully
	return proposes, nil
}
