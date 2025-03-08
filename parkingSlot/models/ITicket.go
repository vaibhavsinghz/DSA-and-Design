package models

type ITicket interface {
	GetID() string
	GetTicketSlotID() int
	GetTicketFloorID() int
	GetVehicleNo() string
}
