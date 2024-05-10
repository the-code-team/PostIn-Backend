package proposes_mod

import (
	usecases "epsa.upv.es/postin_backend/src/modules/proposes_mod/domain/use-cases"
)

var Handlers = []func(){
	usecases.AcceptProposeUseCase,
	usecases.GenerateProposeUseCase,
	usecases.ListProposesUseCase,
	usecases.RejectProposeUseCase,
}

func ProposesModule() {
	for _, handler := range Handlers {
		handler()
	}
}
