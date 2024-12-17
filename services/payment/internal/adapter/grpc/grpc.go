package grpc

import (
	"context"
	"grpc-ms/proto/payment"
	"grpc-ms/services/payment/config"
	"grpc-ms/services/payment/internal/application/core/domain"
	"grpc-ms/services/payment/internal/ports"
	"log/slog"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type Adapter struct {
	api ports.APIPort
	payment.UnimplementedPaymentServer
}

var _ payment.PaymentServer = new(Adapter)

func NewAdapter(api ports.APIPort) *Adapter {
	return &Adapter{
		api: api,
	}
}

func (a Adapter) Run() error {

	listener, err := net.Listen("tcp", config.AppPort())
	if err != nil {
		slog.Error(err.Error())
		return err
	}

	server := grpc.NewServer()

	payment.RegisterPaymentServer(server, a)
	if config.IsDevEnv() {
		reflection.Register(server)
	}

	slog.Info("starting service at ", config.AppPort())

	if err := server.Serve(listener); err != nil {
		slog.Error(err.Error())
		return err
	}

	return nil

}

func (a Adapter) Create(ctx context.Context, request *payment.CreatePaymentRequest) (response *payment.CreatePaymentResponse, err error) {

	var paymentRequest domain.Payment = domain.NewPayment(request.UserId, float64(request.TotalPrice))

	billId, err := a.api.Charge(ctx, &paymentRequest)
	if err != nil {
		return nil, err
	}

	return &payment.CreatePaymentResponse{BillId: int64(billId)}, nil
}

func (a Adapter) Get(ctx context.Context, request *payment.CreatePaymentRequest) (response *payment.CreatePaymentResponse, err error) {
	return nil, nil
}
