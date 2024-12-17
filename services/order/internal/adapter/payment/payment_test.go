package payment

import (
	"grpc-ms/services/order/internal/application/core/domain"
	"grpc-ms/services/order/internal/ports"

	"github.com/stretchr/testify/mock"
)

/* mocked grpc adopter */

type mockedPaymentAdopter struct {
	mock.Mock
}

var _ ports.PaymentPort = new(mockedPaymentAdopter)

func (p *mockedPaymentAdopter) Charge(order *domain.Order) error {
	return p.Called(order).Error(0)
}

/* mocked db adopter */
type mockedDb struct {
	mock.Mock
}

var _ ports.DBPort = new(mockedDb)

func (d *mockedDb) Get(id string) (domain.Order, error) {
	return d.Called(id).Get(0).(domain.Order), d.Called(id).Error(1)
}
func (d *mockedDb) Save(order *domain.Order) error {
	return d.Called(order).Error(0)
}

