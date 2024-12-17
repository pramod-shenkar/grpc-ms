package db

import (
	"grpc-ms/services/order/internal/application/core/domain"

	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	CustomerID int64
	Status     string
	OrderItems OrderItems
}

func (o *Order) FromOrder(order domain.Order) {

	var orderItems OrderItems
	orderItems.FromOrderItmes(order.OrderItems)
	*o = Order{
		CustomerID: order.CustomerID,
		Status:     order.Status,
		OrderItems: orderItems,
	}

}

func (o *Order) ToOrder() domain.Order {

	return domain.Order{

		ID:         int64(o.ID),
		CustomerID: o.CustomerID,
		Status:     o.Status,
		OrderItems: o.OrderItems.ToOrderItems(),
		CreatedAt:  o.CreatedAt.UnixNano(),
	}

}

type OrderItems []OrderItem

func (o *OrderItems) FromOrderItmes(items []domain.OrderItem) {

	for _, v := range items {
		var item OrderItem
		item.FromOrderItem(v)
		*o = append(*o, item)
	}
}

func (o *OrderItems) ToOrderItems() []domain.OrderItem {

	var order []domain.OrderItem
	for _, v := range *o {
		order = append(order, v.ToOrderItem())
	}
	return order

}

type OrderItem struct {
	gorm.Model
	ProductCode string
	UnitPrice   float32
	Quantity    int32
	OrderID     uint
}

func (o *OrderItem) FromOrderItem(item domain.OrderItem) {
	*o = OrderItem{
		ProductCode: item.ProductCode,
		UnitPrice:   item.UnitPrice,
		Quantity:    item.Quantity,
	}
}

func (o *OrderItem) ToOrderItem() domain.OrderItem {
	return domain.OrderItem{
		ProductCode: o.ProductCode,
		UnitPrice:   o.UnitPrice,
		Quantity:    o.Quantity,
	}
}
