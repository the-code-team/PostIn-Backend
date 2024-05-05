package queries

type GetProfileQuery struct {
	userId string
}

func (q *GetProfileQuery) Type() string {
	return "GetProfileQuery"
}
