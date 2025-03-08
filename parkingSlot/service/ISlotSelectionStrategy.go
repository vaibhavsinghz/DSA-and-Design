package service

import "Self/parkingSlot/models"

type ISlotSelectionStrategy interface {
	FindAndAssignSlot(vehicle models.IVehicle, floors []models.IParkingFloor) models.IParkingSlot
}
