package models

import (
	"errors"
	"fmt"
	"sync"
)

type ParkingService struct {
	ID            string
	ParkingFloors []IParkingFloor
	ActiveTickets map[string]ITicket
	mu            sync.Mutex
}

func NewParkingService(ID string) IParkingService {
	return &ParkingService{
		ID:            ID,
		ParkingFloors: []IParkingFloor{},
		ActiveTickets: make(map[string]ITicket),
	}
}

func (ps *ParkingService) AddFloor() {
	ps.mu.Lock()
	defer ps.mu.Unlock()

	newFloorID := len(ps.ParkingFloors) + 1
	parkingFloor := NewParkingFloor(newFloorID)
	ps.ParkingFloors = append(ps.ParkingFloors, parkingFloor)
}

func (ps *ParkingService) AddSlotToFloor(floorID int, vehicleType VehicleType) error {
	ps.mu.Lock()
	defer ps.mu.Unlock()

	if floorID > len(ps.ParkingFloors) {
		return fmt.Errorf("floor %d does not exist", floorID)
	}

	floor := ps.ParkingFloors[floorID-1]
	floor.AddSlot(vehicleType)

	return nil
}

func (ps *ParkingService) Park(vehicleNo, colour string, vehicleType VehicleType) (ITicket, error) {
	ps.mu.Lock()
	defer ps.mu.Unlock()

	vehicle := NewVehicle(vehicleType, vehicleNo, colour)
	slot := ps.findAndAssignSlot(vehicle)
	if slot == nil {
		return nil, errors.New("slot not found")
	}
	ticket := NewTicket(ps.ID, vehicle.GetVehicleNo(), slot.GetSlotID(), slot.GetSlotFloorID())
	ps.ActiveTickets[ticket.GetID()] = ticket
	return ticket, nil
}

func (ps *ParkingService) findAndAssignSlot(vehicle IVehicle) IParkingSlot {
	for _, floor := range ps.ParkingFloors {
		if floor.IsSlotAvailable(vehicle.GetVehicleType()) {
			slot, err := floor.AssignSlot(vehicle)
			if err != nil {
				fmt.Printf("error assigning %s to floor %d, retrying...\n", vehicle.GetVehicleNo(), floor.GetFloorID())
				continue
			}
			return slot
		}
	}
	return nil
}

func (ps *ParkingService) UnPark(ticketID string) error {
	ps.mu.Lock()
	defer ps.mu.Unlock()

	ticket, exist := ps.ActiveTickets[ticketID]
	if !exist {
		return errors.New("ticket not found")
	}

	if ticket.GetTicketFloorID() > len(ps.ParkingFloors) {
		return errors.New("ticket floor does not exist")
	}

	floor := ps.ParkingFloors[ticket.GetTicketFloorID()-1]
	if err := floor.VacateSpot(ticket.GetTicketSlotID(), ticket.GetVehicleNo()); err != nil {
		return err
	}
	delete(ps.ActiveTickets, ticket.GetID())
	return nil
}
