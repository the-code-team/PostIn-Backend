package queries

type ListPhotosQuery struct {
	Email string
}

func (q *ListPhotosQuery) Type() string {
	return "ListPhotosQuery"
}
