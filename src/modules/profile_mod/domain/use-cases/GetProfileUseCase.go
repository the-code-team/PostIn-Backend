package usecases

import (
	"github.com/gogolfing/cbus"

	"../../../../providers"
	"../queries"
)

func GetProfileUseCase() {
	bus := providers.GetCommandBus()

	bus.Handle(&GetProfileQuery{}, HandlerFunc(func(ctx context.Context, command Command) (interface{}, error) {
		user := &User{
			Name: command.(*CreateUserCommand).Name,
		}
		return user, nil
	}))
	
}