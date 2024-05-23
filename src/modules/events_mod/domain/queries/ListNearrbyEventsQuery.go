package queries

type ListNearbyEventsQuery struct {
	Longitude   float64
	Latitude    float64
	MaxDistance float64
}

func (q *ListNearbyEventsQuery) Type() string {
	return "ListNearbyEventsQuery"
}
