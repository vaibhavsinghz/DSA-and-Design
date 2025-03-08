package models

type ParkingFloor struct {
	ID            int
	ParkingSlots  []IParkingSlot
	AvailableSlot map[VehicleType]int
}

func NewParkingFloor(id int) IParkingFloor {
	return &ParkingFloor{
		ID:            id,
		ParkingSlots:  []IParkingSlot{},
		AvailableSlot: make(map[VehicleType]int),
	}
}

func (pf *ParkingFloor) AddSlot(vehicleType VehicleType) {
	slotID := len(pf.ParkingSlots) + 1
	slot := NewParkingSlot(slotID, pf.ID, vehicleType)
	pf.ParkingSlots = append(pf.ParkingSlots, slot)
	pf.AvailableSlot[vehicleType]++
}

func (pf *ParkingFloor) IsSlotAvailable(vehicleType VehicleType) bool {
	return pf.AvailableSlot[vehicleType] > 0
}

func (pf *ParkingFloor) AssignSlot(vehicle IVehicle) (IParkingSlot, error) {
	for _, slot := range pf.ParkingSlots {
		if slot.GetSlotVehicleType() == vehicle.GetVehicleType() && slot.IsSlotAvailable() {
			if err := slot.ParkVehicle(vehicle); err != nil {
				return nil, err
			}
			pf.AvailableSlot[slot.GetSlotVehicleType()]--
			return slot, nil
		}
	}
	return nil, nil
}

func (pf *ParkingFloor) VacateSpot(slotID int, vehicleNo string) error {
	for _, slot := range pf.ParkingSlots {
		if slot.GetSlotID() == slotID {
			if err := slot.RemoveVehicle(vehicleNo); err != nil {
				return err
			}
			pf.AvailableSlot[slot.GetSlotVehicleType()]++
			break
		}
	}
	return nil
}

func (pf *ParkingFloor) GetFloorID() int {
	return pf.ID
}
