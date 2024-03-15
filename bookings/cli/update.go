package cli

import (
	booking "bookings/data"
	"fmt"
)

// returns the updated booking and the index
func UpdateBookings(bookings []booking.Luxury) ([]booking.Luxury, string) {
	s := ""
	fmt.Print("Enter booking id: ")
	fmt.Scanln(&s)
	for i, booking := range bookings {
		if booking.GetBookingID() == s {
			bookings[i] = NewBooking()
			return bookings, fmt.Sprintf("updated booking %d: %v", i, booking)
		}
	}
	return bookings, "booking not found!"
}
