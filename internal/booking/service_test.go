package booking

import (
	"sync"
	"sync/atomic"
	"testing"

	"github.com/google/uuid"
)

func TestConcurrentBooking_ExactlyOneWins(t *testing.T){
	store := NewConcurrentStore()
	svc := NewService(store)

	const numGoroutines = 100_000 //100k users trying to book a seat at a time
	
	var (
		successes atomic.Int64
		failures atomic.Int64
		wg sync.WaitGroup
	)	 
	wg.Add(numGoroutines)

	for i:= range numGoroutines{
		go func(userNum int){
			defer wg.Done()
			err := svc.Book(Booking{
				MovieId: "screen-1",
				SeatId: "A1",
				UserId: uuid.New().String(),
			})
			if err == nil{
				successes.Add(1)
			}else{
				failures.Add(1)
			}
		}(i)
	}
wg.Wait()
}
