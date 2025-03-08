package models

type IParkingService interface {
	AddFloor()
	AddSlotToFloor(floorID int, vehicleType VehicleType) error
	Park(vehicleNo, colour string, vehicleType VehicleType) (ITicket, error)
	UnPark(ticketID string) error
}
