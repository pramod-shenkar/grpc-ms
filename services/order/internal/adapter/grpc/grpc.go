package grpc

import (
	"context"
	"grpc-ms/proto/order"
	"grpc-ms/services/order/config"
	"grpc-ms/services/order/internal/application/core/domain"
	"grpc-ms/services/order/internal/ports"
	"log/slog"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type Adapter struct {
	api ports.APIPort
	order.UnimplementedOrderServer
}

var _ order.OrderServer = new(Adapter)

func NewAdapter(api ports.APIPort) *Adapter {
	return &Adapter{
		api: api,
	}
}

func (a Adapter) Run() error {

	listener, err := net.Listen("tcp", config.GetAppPort())
	if err != nil {
		slog.Error(err.Error())
		return err
	}

	server := grpc.NewServer()

	order.RegisterOrderServer(server, a)
	if config.IsDevEnv() {
		reflection.Register(server)
	}

	slog.Info("starting service at ", config.GetAppPort())

	if err := server.Serve(listener); err != nil {
		slog.Error(err.Error())
		return err
	}

	return nil

}

func (a Adapter) Create(ctx context.Context, request *order.CreateOrderRequest) (response *order.CreateOrderResponse, err error) {

	var items domain.OrderItems
	items.FromOrderRequest(request.Items)
	var orderRequest domain.Order = domain.NewOrder(request.UserId, items)

	result, err := a.api.PlaceOrder(orderRequest)
	if err != nil {
		return nil, err
	}

	return &order.CreateOrderResponse{OrderId: result.ID}, nil
}

func (a Adapter) Get(ctx context.Context, request *order.CreateOrderRequest) (response *order.CreateOrderResponse, err error) {
	return nil, nil
}
