package queries

type ListNearbyEventsQuery struct {
	Longitude   float64
	Latitude    float64
	MaxDistance float64 // INFO: In meters
}

func (q *ListNearbyEventsQuery) Type() string {
	return "ListNearbyEventsQuery"
}
