package parkingSlot

import (
	"sync"
	"time"
)

// ParkingLot represents the main parking lot system
type ParkingLot struct {
	Name           string
	Floors         []*ParkingFloor
	EntrancePoints []string
	ExitPoints     []string
	ActiveTickets  map[string]*Ticket
	mu             sync.RWMutex
}

// NewParkingLot creates a new parking lot
func NewParkingLot(name string) *ParkingLot {
	return &ParkingLot{
		Name:           name,
		Floors:         make([]*ParkingFloor, 0),
		EntrancePoints: make([]string, 0),
		ExitPoints:     make([]string, 0),
		ActiveTickets:  make(map[string]*Ticket),
	}
}

// AddFloor adds a floor to the parking lot
func (pl *ParkingLot) AddFloor(floor *ParkingFloor) {
	pl.mu.Lock()
	defer pl.mu.Unlock()

	pl.Floors = append(pl.Floors, floor)
}

// AddEntrancePoint adds an entrance point
func (pl *ParkingLot) AddEntrancePoint(entranceID string) {
	pl.mu.Lock()
	defer pl.mu.Unlock()

	pl.EntrancePoints = append(pl.EntrancePoints, entranceID)
}

// AddExitPoint adds an exit point
func (pl *ParkingLot) AddExitPoint(exitID string) {
	pl.mu.Lock()
	defer pl.mu.Unlock()

	pl.ExitPoints = append(pl.ExitPoints, exitID)
}

// IsFull checks if the parking lot is full
func (pl *ParkingLot) IsFull(vehicleType VehicleType) bool {
	pl.mu.RLock()
	defer pl.mu.RUnlock()

	for _, floor := range pl.Floors {
		if floor.IsSpotAvailable(vehicleType) {
			return false
		}
	}

	return true
}

// ParkVehicle parks a vehicle and issues a ticket
func (pl *ParkingLot) ParkVehicle(v *Vehicle) *Ticket {
	if pl.IsFull(v.Type) {
		return nil
	}

	var parkingSpot *ParkingSpot

	// Find a parking spot across all floors
	for _, floor := range pl.Floors {
		parkingSpot = floor.FindAndAssignSpot(v)
		if parkingSpot != nil {
			break
		}
	}

	if parkingSpot == nil {
		return nil
	}

	// Generate ticket ID (in practice, use a proper ID generator)
	ticketID := "TKT-" + time.Now().Format("20060102150405")

	// Create a parking ticket
	newTicket := NewTicket(ticketID, v, parkingSpot)
	v.Ticket = newTicket

	pl.mu.Lock()
	pl.ActiveTickets[ticketID] = newTicket
	pl.mu.Unlock()

	return newTicket
}

// ProcessTicket processes a ticket for payment
func (pl *ParkingLot) ProcessTicket(ticketID string) (float64, bool) {
	pl.mu.Lock()
	defer pl.mu.Unlock()

	ticket, exists := pl.ActiveTickets[ticketID]
	if !exists {
		return 0, false
	}

	fee := CalculateFee()
	return fee, true
}

// PayTicket processes payment for a ticket
func (pl *ParkingLot) PayTicket(ticketID string) bool {
	pl.mu.Lock()
	defer pl.mu.Unlock()

	ticket, exists := pl.ActiveTickets[ticketID]
	if !exists {
		return false
	}

	MarkPaid()
	return true
}

// ExitParkingLot processes a vehicle exit
func (pl *ParkingLot) ExitParkingLot(ticketID string) bool {
	pl.mu.Lock()
	defer pl.mu.Unlock()

	ticket, exists := pl.ActiveTickets[ticketID]
	if !exists {
		return false
	}

	if Status != Paid {
		return false
	}

	// Find the floor for this parking spot
	for _, floor := range pl.Floors {
		for _, spot := range floor.ParkingSpots {
			if spot == ParkingSpot {
				floor.VacateSpot(spot)
				delete(pl.ActiveTickets, ticketID)
				return true
			}
		}
	}

	return false
}
