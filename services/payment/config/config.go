package config

import (
	"fmt"
	"os"
)

func AppPort() string {
	appPort := os.Getenv("AppPort")
	if appPort == "" {
		return fmt.Sprintf(":%v", 3001)
	}
	return appPort
}

func DataSourceUrl() string {
	dataSourceUrl := os.Getenv("DaraSourceUrl")
	if dataSourceUrl == "" {
		// return fmt.Sprintf(
		// 	"%s:%s@tcp(%s)/%s?multiStatements=true&parseTime=true&tls=false&multiStatements=true&interpolateParams=true",
		// 	"root", "verysecretpass", "localhost:3306", "payment",
		// )
		return "root:verysecretpass@tcp(127.0.0.1:3306)/payment"
	}
	return dataSourceUrl
}

type env string

const (
	Dev  env = "dev"
	Test env = "stage"
	Prod env = "prod"
)

func Env() env {
	envStr := os.Getenv("env")
	if envStr == "" {
		return Dev
	}

	return env(envStr)
}

func IsDevEnv() bool {
	return Env() == Dev
}
