package parkingSlot

type VehicleType int

const (
	Car VehicleType = iota
	Motorcycle
	Bus
	Van
)

// Vehicle represents a vehicle that can be parked
type Vehicle struct {
	LicenseNumber string
	Type          VehicleType
	Ticket        *Ticket
}

// NewVehicle creates a new vehicle instance
func NewVehicle(licenseNumber string, vehicleType VehicleType) *Vehicle {
	return &Vehicle{
		LicenseNumber: licenseNumber,
		Type:          vehicleType,
	}
}
