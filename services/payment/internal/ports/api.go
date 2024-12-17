package ports

import (
	"context"
	"grpc-ms/services/payment/internal/application/core/domain"
)

type APIPort interface {
	Charge(ctx context.Context, payment *domain.Payment) (int64, error)
}
