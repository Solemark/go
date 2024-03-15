package cli

import (
	booking "bookings/data"
	"fmt"
)

func NewBooking() booking.Luxury {
	return booking.Luxury{
		Booking: booking.Booking{
			BookingID:         getBookingID(),
			GardenArea:        getGardenArea(),
			NumberOfWeeks:     getNumberOfWeeks(),
			Rooms:             getNumberOfRooms(),
			Address:           getAddress(),
			BookingDate:       getBookingDate(),
			ContactNumber:     getContactNumber(),
			PropertyOwnerName: getPropertyOwnerName(),
		},
		SecurityAlarmCheck: getSecurityAlarmCheck(),
		PoolMaintenance:    getPoolMaintenance(),
	}
}

func getGardenArea() int {
	var output int = 0
	for {
		if output != 0 {
			return output
		}
		fmt.Print("Enter garden area: ")
		fmt.Scanln(&output)
	}
}

func getNumberOfWeeks() int {
	var output int = 0
	for {
		if output != 0 {
			return output
		}
		fmt.Print("Enter number of weeks: ")
		fmt.Scanln(&output)
	}
}

func getNumberOfRooms() int {
	var output int = 0
	for {
		if output != 0 {
			return output
		}
		fmt.Print("Enter number of rooms: ")
		fmt.Scanln(&output)
	}
}

func getAddress() string {
	var output string = ""
	for {
		if output != "" {
			return output
		}
		fmt.Print("Enter address: ")
		fmt.Scanln(&output)
	}
}

func getBookingDate() string {
	var output string = ""
	for {
		if output != "" {
			return output
		}
		fmt.Print("Enter booking date: ")
		fmt.Scanln(&output)
	}
}

func getBookingID() string {
	var output string = ""
	for {
		if output != "" {
			return output
		}
		fmt.Print("Enter booking id: ")
		fmt.Scanln(&output)
	}
}

func getContactNumber() string {
	var output string = ""
	for {
		if output != "" {
			return output
		}
		fmt.Print("Enter number: ")
		fmt.Scanln(&output)
	}
}

func getPropertyOwnerName() string {
	var output string = ""
	for {
		if output != "" {
			return output
		}
		fmt.Print("Enter property owner name: ")
		fmt.Scanln(&output)
	}
}

func getSecurityAlarmCheck() bool {
	var input string = ""
	for {
		if input == "0" || input == "1" {
			break
		}
		fmt.Print("Do a security alarm check(0/1): ")
		fmt.Scanln(&input)
	}
	var output bool = false
	if input == "1" {
		output = true
	}
	return output
}

func getPoolMaintenance() bool {
	var input string = ""
	for {
		if input == "0" || input == "1" {
			break
		}
		fmt.Print("Do pool maintenance(0/1): ")
		fmt.Scanln(&input)
	}
	var output bool = false
	if input == "1" {
		output = true
	}
	return output
}
