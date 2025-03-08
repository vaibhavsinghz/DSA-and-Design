package service

import (
	"Self/parkingSlot/models"
	"fmt"
)

type DefaultSlotSelection struct{}

func NewDefaultSlotSelection() ISlotSelectionStrategy {
	return &DefaultSlotSelection{}
}

func (d *DefaultSlotSelection) FindAndAssignSlot(vehicle models.IVehicle, floors []models.IParkingFloor) models.IParkingSlot {
	for _, floor := range floors {
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
