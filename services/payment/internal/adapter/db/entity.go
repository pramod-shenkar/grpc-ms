package db

import (
	"grpc-ms/services/payment/internal/application/core/domain"

	"gorm.io/gorm"
)

type Payment struct {
	gorm.Model
	CustomerID int64
	Status     string
	Price      float64
}

func (o *Payment) FromPayment(payment domain.Payment) {

	*o = Payment{
		CustomerID: payment.CustomerID,
		Status:     payment.Status,
		Price:      payment.Price,
	}

}

func (o *Payment) ToPayment() domain.Payment {

	return domain.Payment{
		ID:         int64(o.ID),
		CustomerID: o.CustomerID,
		Status:     o.Status,
		Price:      o.Price,
		CreatedAt:  o.CreatedAt.UnixNano(),
	}

}
