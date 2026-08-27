package booking

import "sync"

type ConcurrentStore struct {
	// seats -> booking, use map, evetually we will use distributed storage
	// for now we will keep it in-memory(map here)
	bookings map[string]Booking // {SeatId, Booking} eg.{"A2",Booking} O(1)
	sync.RWMutex
}

// lets create a constructor so we can easily replace it on teh test
func NewConcurrentStore() *ConcurrentStore {
	return &ConcurrentStore{
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
func (s *ConcurrentStore) Book(b Booking) error {
	//we use pessimistic approach
	s.Lock()
	defer s.Unlock()

	// seat is taken? error: populate the seat;
	if _, exists := s.bookings[b.SeatId]; exists { //we can use ID simply
		// but the source of truth is SeatId and ID can be used for Redis stuff
		return ErrSeatAlreadyBooked
	}
	s.bookings[b.SeatId] = b
	return nil

}
func (s *ConcurrentStore) ListBooings(movieId string) []Booking {
	s.RLock()
	defer s.RUnlock()
	
	var result []Booking

	for _, b := range s.bookings {
		if b.MovieId == movieId {
			result = append(result, b)
		}
	}
	return result
}
