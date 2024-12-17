package ports

import "grpc-ms/services/order/internal/application/core/domain"

type PaymentPort interface {
	Charge(order *domain.Order) error
}
