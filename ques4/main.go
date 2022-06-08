// Q4. Make a struct with following fields: Seats , RWMutex, Waitgroup.
// Create two methods over struct. First method returns number of
// seats available. Second method books a seat. If seat is equal to zero
// no one can book it.
// You can allow simultaneous reading of seats data but make sure only
// one go routine can book a seat at a time.
// In main function set available seats field of struct to 4 and run 10 go
// routines to book a seat and 15 go routines which will try to read
// available seats.
// Use wait groups as necessary. (20 Marks)

package main

import (
	"fmt"
	"sync"
)

type booking struct {
	Seats int
	sync.RWMutex
	sync.WaitGroup
}

func seatsAvailable(a booking, wg *sync.WaitGroup) (s int) {
	wg.Done()
	if a.Seats == 0 {
		fmt.Println("no seats")
	} else {
		fmt.Println("seats available ")
	}
}
func seatBooking(wg *sync.WaitGroup) {
	wg.Done()

}

func main() {
	var wg sync.WaitGroup
	seat := booking{Seats: 4}

	wg.Add(10)
	go seatsAvailable(seat, &wg)
	go seatBooking(&wg)

	//
}
