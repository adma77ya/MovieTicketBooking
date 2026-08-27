package booking

import (
	"errors"
)

var (
	ErrSeatAlreadyBooked = errors.New("Seat is already taken")
)

type Booking struct{
	ID string
	MovieId string
	SeatId string
	UserId string
	Status string
}

// here we do teh DI(dependency injection), we are going to have
// multiple instances of the different Storeages, and inject it here
type BookingStore interface{
	Book(b Booking) error
	ListBooings(movieId string) []Booking
}