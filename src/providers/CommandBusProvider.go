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