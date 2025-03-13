package main

import (
	"fmt"
	"math/rand"
	"sync"
)

/*
location
cab
drives/ trip
user -> driver, passeneger

bookCab(passenegerID, location)
endRide()

*/

//models

type Location int

func NewLocation(val int) Location {
	return Location(val)
}

type Cab struct {
	CabID    int
	DriverID int
	IsBooked bool
	Location Location
}

func NewCab(id int, driverID int, location Location) *Cab {
	return &Cab{CabID: id, DriverID: driverID, Location: location}
}

type User struct {
	ID         int
	Name       string
	TotalRides int
	Location   Location
}

func NewUser(id int, name string, loc Location) *User {
	return &User{
		ID:         id,
		Name:       name,
		Location:   loc,
		TotalRides: 0,
	}
}

type Passenger struct {
	User
}

func NewPassenger(id int, name string, loc Location) *Passenger {
	user := NewUser(id, name, loc)
	return &Passenger{User: *user}
}

type Driver struct {
	User
	TotalEarnings int
}

func NewDriver(id int, name string, loc Location) *Driver {
	user := NewUser(id, name, loc)
	return &Driver{User: *user}
}

type Trip struct {
	TripID   int
	Fare     int
	From, To Location
}

func NewTrip(id int, fare int, from, to Location) *Trip {
	return &Trip{
		TripID: id,
		Fare:   fare,
		From:   from,
		To:     to,
	}
}

type Lane struct {
	Cab *Cab
}

//[, 0  , , , ]

func NewLane(cab *Cab) *Lane {
	return &Lane{Cab: cab}
}

// service
type CabService struct {
	City                      []*Lane
	Cabs                      map[int]*Cab
	Passengers                map[int]*Passenger //passenger id to passeneger
	Drivers                   map[int]*Driver    //driver id to driver
	Trips                     map[int]*Trip      //trip id to trip
	CabBookingRadius          int
	cabIDCount, driverIDCount int
	passengerIDCount          int
	tripIDCount               int
	mu                        sync.Mutex
}

func InitCabService(noOfCabs int, cityLaneSize int, cabBookingRadius int) *CabService {
	cabService := &CabService{
		City:             make([]*Lane, cityLaneSize),
		Cabs:             make(map[int]*Cab),
		Passengers:       make(map[int]*Passenger),
		Drivers:          make(map[int]*Driver),
		Trips:            make(map[int]*Trip),
		CabBookingRadius: cabBookingRadius,
		cabIDCount:       1,
		driverIDCount:    1,
		passengerIDCount: 1,
	}

	RandomCabAssign(noOfCabs, cabService)
	return cabService
}

func RandomCabAssign(noOfCabs int, cabService *CabService) {
	for i := 0; i < noOfCabs; i++ {
		loc := randInt(len(cabService.City) - 1)
		driver := NewDriver(cabService.driverIDCount, "Uber", NewLocation(loc))
		cab := NewCab(cabService.cabIDCount, driver.ID, NewLocation(loc))

		cabService.Cabs[cab.CabID] = cab
		cabService.Drivers[driver.ID] = driver
		cabService.City[loc] = NewLane(cab)
		cabService.driverIDCount++
		cabService.cabIDCount++
	}
}

func (cabService *CabService) AddPassenger(name string, location int) int {
	passengerID := cabService.passengerIDCount
	passenger := NewPassenger(passengerID, name, NewLocation(location))
	cabService.Passengers[passengerID] = passenger
	cabService.passengerIDCount++
	return passengerID
}

func (cabService *CabService) BookCab(passengerID, destination int) (int, error) {
	cabService.mu.Lock()
	defer cabService.mu.Unlock()

	passenger, exist := cabService.Passengers[passengerID]
	if !exist {
		return -1, fmt.Errorf("passenger %d does not exist", passengerID)
	}
	curLoc := passenger.Location
	destinationLoc := NewLocation(destination)

	cabID := cabService.checkCabAvailabilityInRadius(curLoc, cabService.CabBookingRadius)
	if cabID == -1 {
		return -1, fmt.Errorf("cab %d does not exist", cabID)
	}
	cabService.AssignCab(passengerID, cabID, curLoc, destinationLoc)
	return cabID, nil
}

func (cabService *CabService) AssignCab(passengerID, cabID int, cur, destination Location) {
	distance := int(destination) - int(cur)
	distance = abs(distance)
	fare := distance * 10

	cab, _ := cabService.Cabs[cabID]
	cab.IsBooked = true

	driver, _ := cabService.Drivers[cab.DriverID]
	driver.TotalEarnings += fare

	passenger, _ := cabService.Passengers[passengerID]
	passenger.TotalRides++

	trip := NewTrip(cabService.tripIDCount, fare, cur, destination)
	cabService.Trips[trip.TripID] = trip

	fmt.Printf("%d has been assigned cab %d with fare %d \n", passengerID, cabID, fare)
}

func (cabService *CabService) checkCabAvailabilityInRadius(current Location, radius int) int {

	if current < 0 && int(current) >= len(cabService.City) {
		fmt.Println("Not serving")
		return -1
	}

	cur := int(current)
	left, right := int(current), int(current)
	//[, , ]
	for left > cur-radius && right < cur+radius {
		lane1 := cabService.City[left]
		if lane1 != nil && !lane1.Cab.IsBooked {
			return lane1.Cab.CabID
		}

		lane2 := cabService.City[right]
		if lane2 != nil && !lane2.Cab.IsBooked {
			return lane2.Cab.CabID
		}

		left--
		right++
	}

	return -1
}

// util
func randInt(max int) int {
	return rand.Intn(max)
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func main() {
	//fmt.Println("Hello World")
	cabService := InitCabService(4, 10, 3)

	passID1 := cabService.AddPassenger("A", 3)
	passID2 := cabService.AddPassenger("B", 4)

	cabID1, err := cabService.BookCab(passID1, 6)
	if err != nil {
		fmt.Println(err)
	}

	cabID2, err := cabService.BookCab(passID2, 1)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("first Cab", cabID1)
	fmt.Println("second Cab", cabID2)
}
