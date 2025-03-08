package models

type IParkingFloor interface {
	AddSlot(vehicleType VehicleType)
	IsSlotAvailable(vehicleType VehicleType) bool
	AssignSlot(vehicle IVehicle) (IParkingSlot, error)
	VacateSpot(slotID int, vehicleNo string) error
	GetFloorID() int
}
