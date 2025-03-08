package parkingSlot

import (
	"sync"
)

// ParkingFloor represents a floor in the parking lot
type ParkingFloor struct {
	ID             string
	ParkingSpots   []*ParkingSpot
	AvailableSpots map[SpotType]int
	mu             sync.RWMutex // For concurrent access
}

// NewParkingFloor creates a new parking floor
func NewParkingFloor(id string) *ParkingFloor {
	floor := &ParkingFloor{
		ID:             id,
		ParkingSpots:   make([]*ParkingSpot, 0),
		AvailableSpots: make(map[SpotType]int),
	}

	// Initialize counts for each spot type
	for i := Compact; i <= MotorcycleSpot; i++ {
		floor.AvailableSpots[i] = 0
	}

	return floor
}

// AddParkingSpot adds a parking spot to the floor
func (pf *ParkingFloor) AddParkingSpot(spot *ParkingSpot) {
	pf.mu.Lock()
	defer pf.mu.Unlock()

	pf.ParkingSpots = append(pf.ParkingSpots, spot)
	pf.AvailableSpots[spot.Type]++
}

// IsSpotAvailable checks if a spot is available for the vehicle type
func (pf *ParkingFloor) IsSpotAvailable(vehicleType VehicleType) bool {
	pf.mu.RLock()
	defer pf.mu.RUnlock()

	compatibleSpots := getCompatibleSpotTypes(vehicleType)

	for _, spotType := range compatibleSpots {
		if pf.AvailableSpots[spotType] > 0 {
			return true
		}
	}

	return false
}

// FindAndAssignSpot finds and assigns a spot for a vehicle
func (pf *ParkingFloor) FindAndAssignSpot(v *Vehicle) *ParkingSpot {
	pf.mu.Lock()
	defer pf.mu.Unlock()

	compatibleSpots := getCompatibleSpotTypes(v.Type)

	for _, spotType := range compatibleSpots {
		for _, spot := range pf.ParkingSpots {
			if spot.Type == spotType && spot.IsAvailable() {
				spot.ParkVehicle(v)
				pf.AvailableSpots[spotType]--
				return spot
			}
		}
	}

	return nil
}

// VacateSpot vacates a parking spot
func (pf *ParkingFloor) VacateSpot(spot *ParkingSpot) bool {
	pf.mu.Lock()
	defer pf.mu.Unlock()

	if spot.RemoveVehicle() {
		pf.AvailableSpots[spot.Type]++
		return true
	}

	return false
}

// getCompatibleSpotTypes returns compatible spot types for a vehicle type
func getCompatibleSpotTypes(vehicleType VehicleType) []SpotType {
	var compatibleSpots []SpotType

	switch vehicleType {
	case Car:
		compatibleSpots = append(compatibleSpots, Regular, Compact)
	case Motorcycle:
		compatibleSpots = append(compatibleSpots, MotorcycleSpot)
	case Bus:
		compatibleSpots = append(compatibleSpots, Large)
	case Van:
		compatibleSpots = append(compatibleSpots, Large, Regular)
	}

	return compatibleSpots
}
