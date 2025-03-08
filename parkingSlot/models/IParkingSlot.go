package models

type IParkingSlot interface {
	GetSlotVehicleType() VehicleType
	IsSlotAvailable() bool
	ParkVehicle(vehicle IVehicle) error
	RemoveVehicle(vehicleNo string) error
	GetSlotFloorID() int
	GetSlotID() int
}
