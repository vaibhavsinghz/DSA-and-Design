package models

type IVehicle interface {
	GetVehicleType() VehicleType
	GetVehicleNo() string
}
