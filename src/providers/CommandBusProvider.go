package providers

import (
	"github.com/gogolfing/cbus"
	"sync"
)

var (
	busOnce sync.Once
	bus     *cbus.Bus
)

func GetCommandBus() *cbus.Bus {
	busOnce.Do(func() {
		bus = &cbus.Bus{}
	})

	return bus
}

func InitCommandBus() {
	GetCommandBus()
}
