package service

import (
	"Self/parkingSlot/models"
	"errors"
	"fmt"
	"sync"
)

type ParkingService struct {
	ID                    string
	ParkingFloors         []models.IParkingFloor
	ActiveTickets         map[string]models.ITicket
	SlotSelectionStrategy ISlotSelectionStrategy
	mu                    sync.Mutex
}

func NewParkingService(ID string, strategy ISlotSelectionStrategy) IParkingService {
	return &ParkingService{
		ID:                    ID,
		ParkingFloors:         []models.IParkingFloor{},
		ActiveTickets:         make(map[string]models.ITicket),
		SlotSelectionStrategy: strategy,
	}
}

func (ps *ParkingService) AddFloor() {
	ps.mu.Lock()
	defer ps.mu.Unlock()

	newFloorID := len(ps.ParkingFloors) + 1
	parkingFloor := models.NewParkingFloor(newFloorID)
	ps.ParkingFloors = append(ps.ParkingFloors, parkingFloor)
}

func (ps *ParkingService) AddSlotToFloor(floorID int, vehicleType models.VehicleType) error {
	ps.mu.Lock()
	defer ps.mu.Unlock()

	if floorID > len(ps.ParkingFloors) {
		return fmt.Errorf("floor %d does not exist", floorID)
	}

	floor := ps.ParkingFloors[floorID-1]
	floor.AddSlot(vehicleType)

	return nil
}

func (ps *ParkingService) Park(vehicleNo, colour string, vehicleType models.VehicleType) (models.ITicket, error) {
	ps.mu.Lock()
	defer ps.mu.Unlock()

	vehicle := models.NewVehicle(vehicleType, vehicleNo, colour)
	slot := ps.SlotSelectionStrategy.FindAndAssignSlot(vehicle, ps.ParkingFloors)
	if slot == nil {
		return nil, errors.New("slot not found")
	}
	ticket := models.NewTicket(ps.ID, vehicle.GetVehicleNo(), slot.GetSlotID(), slot.GetSlotFloorID())
	ps.ActiveTickets[ticket.GetID()] = ticket
	return ticket, nil
}

func (ps *ParkingService) findAndAssignSlot(vehicle models.IVehicle) models.IParkingSlot {
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
