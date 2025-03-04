package main

type Product struct {
	ProductID int
	Name      string
	Quantity  int
}
type Order struct {
	OrderID  int
	Products []Product
	Status   int
}

const (
	OrderCreated = iota
	OrderCompleted
	OrderCancelled
)
