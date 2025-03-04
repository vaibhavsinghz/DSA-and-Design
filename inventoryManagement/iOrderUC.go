package main

type OrderUC interface {
	CreateOrder(products []Product) (*Order, error)
	ConfirmOrder(orderID int) error
	CancelOrder(orderID int) error
}
