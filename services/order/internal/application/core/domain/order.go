package domain

import (
	"grpc-ms/proto/order"
	"time"
)

type Order struct {
	ID         int64
	CustomerID int64
	Status     string
	OrderItems OrderItems
	CreatedAt  int64
}

func NewOrder(customerId int64, orderItems OrderItems) Order {
	return Order{
		CustomerID: customerId,
		Status:     "pending",
		OrderItems: orderItems,
		CreatedAt:  time.Now().Unix(),
	}
}

func (o *Order) TotalPrice() float32 {
	var price float32
	for _, v := range o.OrderItems {
		price += v.UnitPrice * float32(v.Quantity)
	}
	return price
}

type OrderItems []OrderItem

func (o *OrderItems) FromOrderRequest(items []*order.Item) {
	for _, v := range items {
		*o = append(*o, OrderItem{
			ProductCode: v.ProductCode,
			UnitPrice:   float32(v.UnitPrice),
			Quantity:    v.Quantity,
		})
	}
}

type OrderItem struct {
	ProductCode string
	UnitPrice   float32
	Quantity    int32
}
