package main

import (
	"bookings/cli"
	booking "bookings/data"
	"fmt"
	"os"
	"strings"
)

func linebreak() string {
	return "--------------------"
}

func newHandler() (booking.Luxury, string) {
	booking := cli.NewBooking()
	return booking, fmt.Sprintf("%s\n%v", linebreak(), booking)
}

func searchHandler(bookings []booking.Luxury) string {
	return cli.SearchBookings(bookings)
}

func updateHandler(bookings []booking.Luxury) ([]booking.Luxury, string) {
	return cli.UpdateBookings(bookings)
}

func removeHandler(bookings []booking.Luxury) []booking.Luxury {
	return cli.RemoveBookings(bookings)
}

func main() {
	fmt.Println("Booking Management System")
	var instructions []string = []string{
		"1 to add new booking",
		"2 to search existing booking",
		"3 to list all bookings",
		"4 to update a booking",
		"5 to remove a booking",
		"Enter any other key to quit",
	}
	var bookings []booking.Luxury = []booking.Luxury{}
	for {
		var input int = 0
		fmt.Printf("%s\n%s\n", linebreak(), strings.Join(instructions, ", "))
		fmt.Scanln(&input)

		switch input {
		case 1:
			b, out := newHandler()
			bookings = append(bookings, b)
			fmt.Println(out)
		case 2:
			fmt.Println(searchHandler(bookings))
		case 3:
			fmt.Printf("%s\n%v\n", linebreak(), bookings)
		case 4:
			b, s := updateHandler(bookings)
			bookings = b
			fmt.Println(s)
		case 5:
			bookings = removeHandler(bookings)
		default:
			os.Exit(0)
		}
	}
}
