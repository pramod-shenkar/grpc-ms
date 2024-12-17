package api

import (
	"grpc-ms/services/order/internal/application/core/domain"
	"grpc-ms/services/order/internal/ports"
	"log/slog"
	"strings"

	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Application struct {
	db      ports.DBPort
	payment ports.PaymentPort
}

func NewApplication(db ports.DBPort, payment ports.PaymentPort) *Application {
	return &Application{
		db:      db,
		payment: payment,
	}
}

func (a Application) PlaceOrder(order domain.Order) (domain.Order, error) {
	err := a.db.Save(&order)
	if err != nil {
		slog.Error(err.Error())
		return domain.Order{}, err
	}

	err = a.payment.Charge(&order)
	if err != nil {
		slog.Error(err.Error())

		/*
			status.Errorf(
				codes.Internal,
				fmt.Sprintf(
					"error while charging user : %v :: %v",
					order.CustomerID, err.Error(),
				),
			)
		*/

		// sts, _ := status.FromError(err)
		// errString := sts.Message()

		var errStrings []string
		sts := status.Convert(err)
		for _, v := range sts.Details() {

			switch t := v.(type) {
			case *errdetails.BadRequest:
				for _, violation := range t.GetFieldViolations() {
					errStrings = append(errStrings, violation.Description)
				}
			}
		}
		errString := strings.Join(errStrings, "\n")

		statusWithDetails, _ :=
			status.New(
				codes.InvalidArgument,
				"order creation failed",
			).WithDetails(
				&errdetails.BadRequest{
					FieldViolations: []*errdetails.BadRequest_FieldViolation{
						{
							Field:       "payment",
							Description: errString,
						},
					},
				},
			)

		return domain.Order{}, statusWithDetails.Err()

	}
	return order, nil
}
