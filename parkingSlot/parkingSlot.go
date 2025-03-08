package parkingSlot

type SpotType int

const (
	Compact SpotType = iota
	Regular
	Large
	Handicapped
	MotorcycleSpot
)

// ParkingSpot represents a single parking spot
type ParkingSpot struct {
	ID         string
	Type       SpotType
	IsOccupied bool
	Vehicle    *Vehicle
	Floor      int
}

// IsAvailable checks if the spot is available
func (ps *ParkingSpot) IsAvailable() bool {
	return !ps.IsOccupied
}

// ParkVehicle parks a vehicle in the spot
func (ps *ParkingSpot) ParkVehicle(vehicle *Vehicle) bool {
	if !ps.IsAvailable() {
		return false
	}

	ps.Vehicle = vehicle
	ps.IsOccupied = true
	return true
}

// RemoveVehicle removes a vehicle from the spot
func (ps *ParkingSpot) RemoveVehicle() bool {
	if ps.IsAvailable() {
		return false
	}

	ps.Vehicle = nil
	ps.IsOccupied = false
	return true
}

// NewParkingSpot creates a new parking spot
func NewParkingSpot(id string, spotType SpotType, floor int) *ParkingSpot {
	return &ParkingSpot{
		ID:         id,
		Type:       spotType,
		Floor:      floor,
		IsOccupied: false,
	}
}
