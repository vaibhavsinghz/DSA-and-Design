package service

import "Self/parkingSlot/models"

type IParkingService interface {
	AddFloor()
	AddSlotToFloor(floorID int, vehicleType models.VehicleType) error
	Park(vehicleNo, colour string, vehicleType models.VehicleType) (models.ITicket, error)
	UnPark(ticketID string) error
}
