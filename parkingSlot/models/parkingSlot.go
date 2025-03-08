package models

import "errors"

type ParkingSlot struct {
	ID         int
	Type       VehicleType
	Floor      int
	IsOccupied bool
	Vehicle    IVehicle
}

func NewParkingSlot(slotID, floorID int, vehicleType VehicleType) IParkingSlot {
	return &ParkingSlot{
		ID:    slotID,
		Type:  vehicleType,
		Floor: floorID,
	}
}

func (p *ParkingSlot) GetSlotID() int {
	return p.ID
}

func (p *ParkingSlot) GetSlotFloorID() int {
	return p.Floor
}

func (p *ParkingSlot) IsSlotAvailable() bool {
	return !p.IsOccupied
}

func (p *ParkingSlot) GetSlotVehicleType() VehicleType {
	return p.Type
}

func (p *ParkingSlot) ParkVehicle(vehicle IVehicle) error {
	if p.IsOccupied {
		return errors.New("slot is already occupied")
	}
	p.Vehicle = vehicle
	p.IsOccupied = true
	return nil
}

func (p *ParkingSlot) RemoveVehicle(vehicleNo string) error {
	if !p.IsOccupied {
		return errors.New("slot is not occupied")
	}
	if p.Vehicle.GetVehicleNo() != vehicleNo {
		return errors.New("vehicle does not match")
	}
	p.IsOccupied = false
	p.Vehicle = nil
	return nil
}
