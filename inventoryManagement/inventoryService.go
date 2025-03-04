package main

import (
	"errors"
	"sync"
)

type InventoryService struct {
	Stock         map[int]*Product // map[productID]Product
	mu            sync.RWMutex
	nextProductID int
}

func NewInventoryService() InventoryUC {
	return &InventoryService{
		Stock:         make(map[int]*Product),
		nextProductID: 1,
	}
}

func (inventoryService *InventoryService) AddProduct(productID *int, quantity int, name string) error {
	inventoryService.mu.Lock()
	defer inventoryService.mu.Unlock()

	if productID == nil {
		newProduct := &Product{
			ProductID: inventoryService.nextProductID,
			Quantity:  quantity,
			Name:      name,
		}
		inventoryService.nextProductID++
		inventoryService.Stock[newProduct.ProductID] = newProduct
	} else {
		product, exist := inventoryService.Stock[*productID]
		if !exist {
			return errors.New("product not found")
		}
		product.Quantity += quantity
		product.Name = name
	}
	return nil
}
func (inventoryService *InventoryService) ConsumeProduct(productID, quantity int) error {
	inventoryService.mu.Lock()
	defer inventoryService.mu.Unlock()

	product, exist := inventoryService.Stock[productID]
	if !exist {
		return errors.New("product not found")
	}

	if product.Quantity < quantity {
		return errors.New("not enough quantity")
	}

	product.Quantity -= quantity
	return nil
}
func (inventoryService *InventoryService) GetProduct(productID int) (*Product, error) {
	inventoryService.mu.RLock()
	defer inventoryService.mu.RUnlock()

	product, exist := inventoryService.Stock[productID]
	if !exist {
		return nil, errors.New("product not found")
	}

	return product, nil
}
func (inventoryService *InventoryService) GetProductStock(productID int) (int, error) {
	product, err := inventoryService.GetProduct(productID)
	if err != nil {
		return -1, err
	}
	return product.Quantity, nil
}
