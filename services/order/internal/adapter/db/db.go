package db

import (
	"grpc-ms/services/order/config"
	"grpc-ms/services/order/internal/application/core/domain"
	"grpc-ms/services/order/internal/ports"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	_ "github.com/go-sql-driver/mysql"
)

type Adapter struct {
	db *gorm.DB
}

var _ ports.DBPort = new(Adapter)

func NewAdapter() (*Adapter, error) {

	db, err := gorm.Open(mysql.Open(config.Env.DataSourceUrl), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	result := db.Exec("CREATE DATABASE IF NOT EXISTS " + config.Env.ServiceName + ";")
	if result.Error != nil {
		return nil, result.Error
	}

	db, err = gorm.Open(mysql.Open(config.Env.DataSourceUrl+config.Env.ServiceName), &gorm.Config{})
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	err = db.AutoMigrate(&Order{}, OrderItem{})
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	return &Adapter{db: db}, nil
}

func (a *Adapter) Save(order *domain.Order) error {
	var orderEntity Order
	orderEntity.FromOrder(*order)
	result := a.db.Create(&orderEntity)
	if result.Error != nil {
		return result.Error
	}
	order.ID = int64(orderEntity.ID)
	return nil
}

func (a *Adapter) Get(id string) (domain.Order, error) {
	var order Order
	err := a.db.First(&order).Error
	if err != nil {
		return domain.Order{}, err
	}

	return order.ToOrder(), nil
}
