package booking

type Booking struct {
	BookingID         string
	GardenArea        int
	NumberOfWeeks     int
	Rooms             int
	Address           string
	BookingDate       string
	ContactNumber     string
	PropertyOwnerName string
}

func (booking *Booking) GetGardenArea() int {
	return booking.GardenArea
}
func (booking *Booking) SetGardenArea(input int) {
	booking.GardenArea = input
}

func (booking *Booking) GetNumberOfWeeks() int {
	return booking.NumberOfWeeks
}
func (booking *Booking) SetNumberOfWeeks(input int) {
	booking.NumberOfWeeks = input
}

func (booking *Booking) GetRooms() int {
	return booking.Rooms
}
func (booking *Booking) SetRooms(input int) {
	booking.Rooms = input
}

func (booking *Booking) GetAddress() string {
	return booking.Address
}
func (booking *Booking) SetAddress(input string) {
	booking.Address = input
}

func (booking *Booking) GetBookingDate() string {
	return booking.BookingDate
}
func (booking *Booking) SetBookingDate(input string) {
	booking.BookingDate = input
}

func (booking *Booking) GetBookingID() string {
	return booking.BookingID
}
func (booking *Booking) SetBookingID(input string) {
	booking.BookingID = input
}

func (booking *Booking) GetContactNumber() string {
	return booking.ContactNumber
}
func (booking *Booking) SetContactNumber(input string) {
	booking.ContactNumber = input
}

func (booking *Booking) GetPropertyOwnerName() string {
	return booking.PropertyOwnerName
}
func (booking *Booking) SetPropertyOwnerName(input string) {
	booking.PropertyOwnerName = input
}

type Luxury struct {
	Booking
	SecurityAlarmCheck bool
	PoolMaintenance    bool
}

func (luxury *Luxury) GetSecurityAlarmCheck() bool {
	return luxury.SecurityAlarmCheck
}
func (luxury *Luxury) SetSecurityAlarmCheck(input bool) {
	luxury.SecurityAlarmCheck = input
}

func (luxury *Luxury) GetPoolMaintenance() bool {
	return luxury.PoolMaintenance
}
func (luxury *Luxury) SetPoolMaintenance(input bool) {
	luxury.PoolMaintenance = input
}
