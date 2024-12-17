package db

import (
	"grpc-ms/services/order/internal/application/core/domain"
	"grpc-ms/services/order/internal/ports"
	"log/slog"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Adapter struct {
	db *gorm.DB
}

var _ ports.DBPort = new(Adapter)

func NewAdapter(dataSourceUrl string) (*Adapter, error) {
	db, err := gorm.Open(mysql.Open(dataSourceUrl), &gorm.Config{})
	if err != nil {
		slog.Error(err.Error())
		return nil, err
	}

	err = db.AutoMigrate(&Order{}, OrderItem{})
	if err != nil {
		slog.Error(err.Error())
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
