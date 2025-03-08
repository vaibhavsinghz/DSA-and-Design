package models

import (
	"fmt"
	"time"
)

type Ticket struct {
	ID              string
	IssuedAt        time.Time
	PaidAt          time.Time
	VehicleNo       string
	SlotID, FloorID int
}

const (
	ticketIDFormat = "%s_%d_%d" //<parking_lot_id>_<floor_no>_<slot_no>
)

func NewTicket(parkingLotID, vehicleNo string, slotID, floorID int) ITicket {
	return &Ticket{
		ID:        fmt.Sprintf(ticketIDFormat, parkingLotID, floorID, slotID),
		IssuedAt:  time.Now(),
		VehicleNo: vehicleNo,
		SlotID:    slotID,
		FloorID:   floorID,
	}
}

func (t *Ticket) GetID() string {
	return t.ID
}

func (t *Ticket) GetTicketSlotID() int {
	return t.SlotID
}

func (t *Ticket) GetTicketFloorID() int {
	return t.FloorID
}

func (t *Ticket) GetVehicleNo() string {
	return t.VehicleNo
}
