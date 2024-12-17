package main

import (
	"grpc-ms/services/payment/config"
	"grpc-ms/services/payment/internal/adapter/db"
	"grpc-ms/services/payment/internal/application/core/api"
	"log/slog"

	"grpc-ms/services/payment/internal/adapter/grpc"
)

func main() {

	dbAdapter, err := db.NewAdapter(config.DataSourceUrl())
	if err != nil {
		slog.Error(err.Error())
		return
	}

	application := api.NewApplication(dbAdapter)

	grpcAdapter := grpc.NewAdapter(application)
	grpcAdapter.Run()

}
