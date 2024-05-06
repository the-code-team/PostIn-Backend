package usecases

import (
	"context"
	"epsa.upv.es/postin_backend/src/modules/profile_mod/domain/commands"
	"epsa.upv.es/postin_backend/src/providers"
	"github.com/gogolfing/cbus"
)

func UpdatePhotosUseCase() {
	bus := providers.GetCommandBus()

	bus.Handle(&commands.UpdatePhotosCommand{}, cbus.HandlerFunc(
		func(ctx context.Context, command cbus.Command) (interface{}, error) {
			// TODO: Implement UpdatePhotosUseCase
			return nil, nil
		}),
	)
}
