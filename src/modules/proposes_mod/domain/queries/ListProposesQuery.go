package queries

import (
	"github.com/google/uuid"
)

type ListProposesQuery struct {
	EventIds []uuid.UUID
}

func (q *ListProposesQuery) Type() string {
	return "ListProposesQuery"
}
