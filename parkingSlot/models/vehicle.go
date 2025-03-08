package models

type Vehicle struct {
	Type           VehicleType
	RegistrationNo string
	Colour         string
}

func NewVehicle(vehicleType VehicleType, registrationNo string, colour string) IVehicle {
	return Vehicle{
		Type:           vehicleType,
		RegistrationNo: registrationNo,
		Colour:         colour,
	}
}

func (v Vehicle) GetVehicleType() VehicleType {
	return v.Type
}

func (v Vehicle) GetVehicleNo() string {
	return v.RegistrationNo
}
