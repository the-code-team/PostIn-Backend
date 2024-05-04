package providers

import (
	"sync"
	"github.com/gogolfing/cbus"
)

var (
	busOnce sync.Once
	bus  *cbus.Bus
)

func GetCommandBus() *cbus.Bus {
	busOnce.Do(func() {
		bus = &cbus.Bus{}
	})

	return bus
}

func ConfigureCommandBus() {
	bus = GetCommandBus()

	/* Commented out for now
	bus.Handle(&CreateUserCommand{}, HandlerFunc(func(ctx context.Context, command Command) (interface{}, error) {
		user := &User{
			Name: command.(*CreateUserCommand).Name,
		}
		return user, nil
	}))
	*/
}