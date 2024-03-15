package cli

import (
	booking "bookings/data"
	"fmt"
)

func RemoveBookings(bookings []booking.Luxury) []booking.Luxury {
	s := ""
	b := []booking.Luxury{}
	fmt.Print("Enter booking id: ")
	fmt.Scanln(&s)
	for _, booking := range bookings {
		if booking.GetBookingID() != s {
			b = append(b, booking)
		}
	}
	return b
}
