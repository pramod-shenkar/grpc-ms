package db

import (
	"grpc-ms/services/payment/internal/application/core/domain"
	"grpc-ms/services/payment/internal/ports"
	"log"

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
		log.Println(err.Error())
		return nil, err
	}

	err = db.AutoMigrate(&Payment{})
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	return &Adapter{db: db}, nil
}

func (a *Adapter) Save(payment *domain.Payment) error {
	var paymentEntity Payment
	paymentEntity.FromPayment(*payment)
	err := a.db.Create(&paymentEntity).Error
	if err != nil {
		return err
	}
	payment.ID = int64(paymentEntity.ID)
	return err
}

func (a *Adapter) Get(id string) (domain.Payment, error) {
	var payment Payment
	err := a.db.First(&payment).Error
	if err != nil {
		return domain.Payment{}, err
	}

	return payment.ToPayment(), nil
}
