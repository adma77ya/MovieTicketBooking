package booking

type Service struct{
	// Dependency Injection
	store BookingStore
}

func (s *Service) Book(b Booking) error{
	return s.store.Book(b)

}


func NewService(store BookingStore) *Service{
	return &Service{store}
}





