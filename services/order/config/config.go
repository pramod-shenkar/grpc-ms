package config

import (
	"fmt"
	"os"
)

func GetAppPort() string {
	appPort := os.Getenv("AppPort")
	if appPort == "" {
		return fmt.Sprintf(":%v", 3000)
	}
	return appPort
}

func GetDataSourceUrl() string {
	dataSourceUrl := os.Getenv("DaraSourceUrl")
	if dataSourceUrl == "" {
		return "user:password@tcp(172.21.0.2:3306)/order"
	}
	return dataSourceUrl
}

type env string

const (
	Dev  env = "dev"
	Test env = "stage"
	Prod env = "prod"
)

func GetEnv() env {
	envStr := os.Getenv("Env")
	if envStr == "" {
		return Dev
	}

	return env(envStr)
}

func GetPaymentServiceUrl() string {
	envStr := os.Getenv("PaymentServiceUrl")
	if envStr == "" {
		return "localhost:3001"
	}

	return envStr
}

func IsDevEnv() bool {
	return GetEnv() == Dev
}
