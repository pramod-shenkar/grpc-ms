package config

import (
	conf "github.com/caarlos0/env"
)

type env struct {
	AppPort           string `env:"AppPort" envDefault:"3000"`
	DataSourceUrl     string `env:"DataSourceUrl" envDefault:"root:password@tcp(localhost:3306)/"`
	ServiceName       string `env:"ServiceName" envDefault:"common"`
	Mode              mode   `env:"Mode" envDefault:"dev"`
	PaymentServiceUrl string `env:"PaymentServiceUrl" envDefault:"localhost:3001"`
}

var Env = New()

func New() env {

	var envVars env
	conf.Parse(&envVars)

	return envVars
}

type mode string

const (
	Dev  mode = "dev"
	Test mode = "stage"
	Prod mode = "prod"
)
