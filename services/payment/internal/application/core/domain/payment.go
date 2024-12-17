package domain

import (
	"time"
)

type Payment struct {
	ID         int64
	CustomerID int64
	Status     string
	Price      float64
	CreatedAt  int64
}

func NewPayment(customerId int64, price float64) Payment {
	return Payment{
		CustomerID: customerId,
		Status:     "pending",
		CreatedAt:  time.Now().Unix(),
	}
}
