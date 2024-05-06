package profile_mod

import (
	usecases "epsa.upv.es/postin_backend/src/modules/profile_mod/domain/use-cases"
)

var Handlers = []func(){
	usecases.UpsertProfileUseCase,
	usecases.DeleteProfileUseCase,
	usecases.GetProfileUseCase,
	usecases.UpdatePhotosUseCase,
}

func ProfileModule() {
	for _, handler := range Handlers {
		handler()
	}
}
