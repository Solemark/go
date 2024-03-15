package cli

import (
	booking "bookings/data"
	"fmt"
)

func SearchBookings(bookings []booking.Luxury) string {
	s := ""
	fmt.Print("Enter booking id: ")
	fmt.Scanln(&s)
	for _, booking := range bookings {
		if booking.GetBookingID() == s {
			return fmt.Sprintf("found booking: %v", booking)
		}
	}
	return "booking not found!"
}
