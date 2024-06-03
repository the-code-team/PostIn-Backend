package profile_mod

import (
	usecases "epsa.upv.es/postin_backend/src/modules/messages_mod/domain/use-cases"
)

var Handlers = []func(){
	usecases.GetMessagesUseCase,
	usecases.PublishMessageUseCase,
}

func MessagesModule() {
	for _, handler := range Handlers {
		handler()
	}
}
