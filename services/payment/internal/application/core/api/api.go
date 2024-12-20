package api

import (
	"context"
	"grpc-ms/services/payment/internal/application/core/domain"
	"grpc-ms/services/payment/internal/ports"
	"log"
)

type Application struct {
	db ports.DBPort
}

func NewApplication(db ports.DBPort) *Application {
	return &Application{
		db: db,
	}
}

var _ ports.APIPort = new(Application)

func (a Application) Charge(ctx context.Context, payment *domain.Payment) (int64, error) {
	err := a.db.Save(payment)
	if err != nil {
		log.Println(err.Error())
		return -1, err
	}
	return payment.ID, nil
}
