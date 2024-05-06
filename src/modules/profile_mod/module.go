package profile_mod

import (
	usecases "epsa.upv.es/postin_backend/src/modules/profile_mod/domain/use-cases"
)

var Handlers = []func(){
	usecases.DeleteProfileUseCase,
	usecases.GetProfileUseCase,
	usecases.ListPhotosUseCase,
	usecases.UpdatePhotosUseCase,
	usecases.UpsertProfileUseCase,
}

func ProfileModule() {
	for _, handler := range Handlers {
		handler()
	}
}
