package usecases

import (
	"context"
	"github.com/gogolfing/cbus"

	"epsa.upv.es/postin_backend/src/modules/profile_mod/domain/queries"
	"epsa.upv.es/postin_backend/src/providers"
)

func GetProfileUseCase() {
	bus := providers.GetCommandBus()

	bus.Handle(&queries.GetProfileQuery{}, cbus.HandlerFunc(
		func(ctx context.Context, command cbus.Command) (interface{}, error) {
			db := providers.GetDatabase()

			/*
				user := &User{
					Name: command.(*CreateUserCommand).Name,
				}
				return user, nil
			*/
		}),
	)

}
