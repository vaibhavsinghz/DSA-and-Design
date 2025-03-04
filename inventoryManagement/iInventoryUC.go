package main

type InventoryUC interface {
	AddProduct(productID *int, quantity int, name string) error
	GetProduct(productID int) (*Product, error)
	ConsumeProduct(productID, quantity int) error
	GetProductStock(productID int) (int, error)
}
