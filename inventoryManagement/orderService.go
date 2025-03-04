package main

import (
	"errors"
	"github.com/emirpasic/gods/sets/hashset"
	"sync"
)

type OrderService struct {
	inventoryUC InventoryUC
	Orders      map[int]*Order //map[OrderID]Order
	mu          sync.RWMutex
	nextOrderID int
}

func NewOrderService(inventory InventoryUC) OrderUC {
	return &OrderService{
		inventoryUC: inventory,
		Orders:      make(map[int]*Order),
		nextOrderID: 1,
	}
}

var validOrderStatus = hashset.New(OrderCreated, OrderCompleted, OrderCancelled)

func (orderService *OrderService) CreateOrder(products []Product) (*Order, error) {
	orderService.mu.Lock()
	defer orderService.mu.Unlock()

	//verify stock
	for _, product := range products {
		productStockQty, err := orderService.inventoryUC.GetProductStock(product.ProductID)
		if err != nil {
			return nil, err
		}

		if product.Quantity > productStockQty {
			return nil, errors.New("can not place order, insufficient quantity")
		}
	}

	newOrder := &Order{
		Products: products,
		OrderID:  orderService.nextOrderID,
		Status:   OrderCreated,
	}

	orderService.nextOrderID++
	orderService.Orders[newOrder.OrderID] = newOrder

	//update Stock
	for _, product := range products {
		err := orderService.inventoryUC.ConsumeProduct(product.ProductID, product.Quantity)
		if err != nil {
			return nil, err
		}
	}

	return newOrder, nil
}
func (orderService *OrderService) ConfirmOrder(orderID int) error {
	return orderService.updateOrderStatus(orderID, OrderCompleted)
}
func (orderService *OrderService) CancelOrder(orderID int) error {
	//return inventory -- according to requirement
	return orderService.updateOrderStatus(orderID, OrderCancelled)

}
func (orderService *OrderService) updateOrderStatus(orderID, status int) error {
	orderService.mu.Lock()
	defer orderService.mu.Unlock()

	order, exist := orderService.Orders[orderID]
	if !exist {
		return errors.New("order does not exist")
	}

	if !validOrderStatus.Contains(status) {
		return errors.New("invalid order status transition request")
	}
	order.Status = status
	return nil
}
