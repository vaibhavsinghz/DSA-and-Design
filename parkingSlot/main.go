package main

import (
	"Self/parkingSlot/models"
	"Self/parkingSlot/service"
	"fmt"
)

func main() {
	fmt.Println("Welcome to parking slot")
	defaultSlotSelection := service.NewDefaultSlotSelection()
	vPark := service.NewParkingService("vPark", defaultSlotSelection)

	for i := 0; i < 4; i++ {
		vPark.AddFloor()
	}

	for i := 0; i < 5; i++ {
		for i := 1; i <= 4; i++ {
			err := vPark.AddSlotToFloor(i, models.Car)
			if err != nil {
				fmt.Println(err)
				return
			}
		}
	}

	for i := 0; i < 3; i++ {
		for i := 1; i <= 4; i++ {
			err := vPark.AddSlotToFloor(i, models.Bike)
			if err != nil {
				fmt.Println(err)
				return
			}
		}
	}

	for i := 0; i < 3; i++ {
		for i := 1; i <= 4; i++ {
			err := vPark.AddSlotToFloor(i, models.Truck)
			if err != nil {
				fmt.Println(err)
				return
			}
		}
	}

	ticket, err := vPark.Park("UP60V0529", "Black", models.Bike)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Ticket %s booked for vehicle no %s on %d floor and %d slot", ticket.GetID(), ticket.GetVehicleNo(), ticket.GetTicketFloorID(), ticket.GetTicketSlotID())

	err = vPark.UnPark(ticket.GetID())
	if err != nil {
		fmt.Println(err)
		return
	}

}
