package main

import "fmt"

/*
1. addProduct(productid, quantity) -> add product in stock
2. createOrder(orderid, list of product ids and their corresponding quantities) -> placeOrder
3. confirmOrder(orderid) -> complete order
4. CancelOrder
5. getStock(productid) -> checkQuantity
*/

func main() {
	inventoryUC := NewInventoryService()
	orderUC := NewOrderService(inventoryUC)

	if err := inventoryUC.AddProduct(nil, 15, "MC Aloo"); err != nil {
		fmt.Println(err)
		return
	}

	if err := inventoryUC.AddProduct(nil, 8, "Aloo Tikki Burger"); err != nil {
		fmt.Println(err)
		return
	}

	order, err := orderUC.CreateOrder(
		[]Product{
			{
				ProductID: 1,
				Quantity:  2,
			},
			{
				ProductID: 2,
				Quantity:  3,
			},
		})

	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Order Created Successfully with order id : ", order.OrderID)

	if err := orderUC.ConfirmOrder(order.OrderID); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Order Confirmed Successfully, order status is : ", order.Status)
}
