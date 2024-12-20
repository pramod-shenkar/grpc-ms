package payment

import (
	"context"
	"grpc-ms/proto/payment"
	"grpc-ms/services/order/internal/application/core/domain"
	"grpc-ms/services/order/internal/ports"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Adapter struct {
	payment payment.PaymentClient
}

var _ ports.PaymentPort = new(Adapter)

func NewAdapter(paymentServiceUrl string) (*Adapter, error) {
	var opts = []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}

	conn, err := grpc.NewClient(paymentServiceUrl, opts...)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	// defer conn.Close()

	client := payment.NewPaymentClient(conn)
	return &Adapter{client}, nil
}

func (a *Adapter) Charge(order *domain.Order) error {

	_, err := a.payment.Create(
		context.Background(),
		&payment.CreatePaymentRequest{
			UserId:     order.CustomerID,
			OrderId:    order.ID,
			TotalPrice: float32(order.TotalPrice()),
		},
	)

	if err != nil {
		log.Println(err.Error())
		return err
	}

	return nil
}
