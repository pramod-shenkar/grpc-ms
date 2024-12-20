package main

import (
	"encoding/json"
	"grpc-ms/services/order/config"
	"grpc-ms/services/order/internal/adapter/db"
	"grpc-ms/services/order/internal/adapter/payment"
	"grpc-ms/services/order/internal/application/core/api"
	"log"

	"grpc-ms/services/order/internal/adapter/grpc"
)

func init() {

	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.SetFlags(log.Flags() &^ (log.Ldate | log.Ltime))

	configBytes, _ := json.MarshalIndent(config.Env, "", " ")
	log.Printf("%v", string(configBytes))
}

func main() {

	defer func() {
		for {

		}
	}()

	dbAdapter, err := db.NewAdapter()
	if err != nil {
		log.Println(err.Error())
		return
	}

	paymentAdapter, err := payment.NewAdapter(config.Env.PaymentServiceUrl)
	if err != nil {
		log.Println(err.Error())
		return
	}

	application := api.NewApplication(dbAdapter, paymentAdapter)

	grpcAdapter := grpc.NewAdapter(application)
	grpcAdapter.Run()

}
