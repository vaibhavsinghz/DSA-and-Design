package parkingSlot

import (
	"time"
)

// TicketStatus represents the status of a parking ticket
type TicketStatus int

const (
	Active TicketStatus = iota
	Paid
	Expired
	Lost
)

// Ticket represents a parking ticket
type Ticket struct {
	ID          string
	IssuedAt    time.Time
	PaidAt      time.Time
	PaidAmount  float64
	Status      TicketStatus
	Vehicle     *Vehicle
	ParkingSpot *ParkingSpot
}

// NewTicket creates a new parking ticket
func NewTicket(id string, vehicle *Vehicle, parkingSpot *ParkingSpot) *Ticket {
	return &Ticket{
		ID:          id,
		IssuedAt:    time.Now(),
		Status:      Active,
		Vehicle:     vehicle,
		ParkingSpot: parkingSpot,
	}
}

// CalculateFee calculates the parking fee
func (t *Ticket) CalculateFee() float64 {
	endTime := time.Now()
	if !t.PaidAt.IsZero() {
		endTime = t.PaidAt
	}

	duration := endTime.Sub(t.IssuedAt)
	hours := int(duration.Hours())

	// Get appropriate pricing strategy
	pricingStrategy := GetPricingStrategy(t.Type)
	return pricingStrategy.CalculatePrice(hours)
}

// MarkPaid marks the ticket as paid
func (t *Ticket) MarkPaid() {
	t.PaidAt = time.Now()
	t.PaidAmount = t.CalculateFee()
	t.Status = Paid
}
