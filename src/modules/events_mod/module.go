package events_mod

import (
	usecases "epsa.upv.es/postin_backend/src/modules/events_mod/domain/use-cases"
)

var Handlers = []func(){
	usecases.GetEventUseCase,
	usecases.ListNearbyEventsUseCase,
	usecases.UpsertEventUseCase,
	usecases.PublishPhotoHeaderUseCase,
}

func EventsModule() {
	for _, handler := range Handlers {
		handler()
	}
}
