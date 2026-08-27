package booking

type MemoryStore struct {
	// seats -> booking, use map, evetually we will use distributed storage
	// for now we will keep it in-memory(map here)
	bookings map[string]Booking // {SeatId, Booking} eg.{"A2",Booking} O(1)

}

// lets create a constructor so we can easily replace it on teh test
func NewMemoryStore() *MemoryStore {
	return &MemoryStore{
		bookings: map[string]Booking{},
	}
}

//now inject the MemoryStore, so we can say MS.Book adn MS.ListBookings
/*
type Booking struct{
	ID string
	MovieId string
	SeatId string
	UserId string
	Status string
}

*/
func (s *MemoryStore) Book(b Booking) error {
	// seat is taken? error: populate the seat;
	if _, exists := s.bookings[b.SeatId]; exists { //we can use ID simply
		// but the source of truth is SeatId and ID can be used for Redis stuff
		return ErrSeatAlreadyBooked
	}
	s.bookings[b.SeatId] = b
	return nil

}
func (s *MemoryStore) ListBooings(movieId string) []Booking {
	var result []Booking

	for _, b := range s.bookings {
		if b.MovieId == movieId {
			result = append(result, b)
		}
	}
	return result
}
