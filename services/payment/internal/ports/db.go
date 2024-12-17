package ports

import "grpc-ms/services/payment/internal/application/core/domain"

type DBPort interface {
	Get(id string) (domain.Payment, error)
	Save(*domain.Payment) error
}
