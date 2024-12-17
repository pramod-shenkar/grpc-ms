package main

import (
	"grpc-ms/services/order/config"
	"grpc-ms/services/order/internal/adapter/db"
	"grpc-ms/services/order/internal/adapter/payment"
	"grpc-ms/services/order/internal/application/core/api"
	"log/slog"

	"grpc-ms/services/order/internal/adapter/grpc"
)

func init() {
	slog.Info(config.GetDataSourceUrl())
	slog.Info(config.GetAppPort())
	slog.Info(string(config.GetEnv()))
	slog.Info(config.GetPaymentServiceUrl())
}

func main() {

	defer func() {
		for {

		}
	}()

	dbAdapter, err := db.NewAdapter(config.GetDataSourceUrl())
	if err != nil {
		slog.Error(err.Error())
		return
	}

	paymentAdapter, err := payment.NewAdapter(config.GetPaymentServiceUrl())
	if err != nil {
		slog.Error(err.Error())
		return
	}

	application := api.NewApplication(dbAdapter, paymentAdapter)

	grpcAdapter := grpc.NewAdapter(application)
	grpcAdapter.Run()

}
