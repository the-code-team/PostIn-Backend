package queries

type GetProfileQuery struct {
	Email string
}

func (q *GetProfileQuery) Type() string {
	return "GetProfileQuery"
}
