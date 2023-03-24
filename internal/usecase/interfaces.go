// Package usecase implements application business logic. Each logic group in own file.
package usecase

import (
	"context"
)

//go:generate mockgen -source=interfaces.go -destination=./mocks_test.go -package=usecase_test

type (
	// Order -.
	Order interface {
		ReceiveOrder(context.Context) error
		GetById(context.Context, string) (OrderDto, error)
	}

	// OrderRepo -.
	OrderRepo interface {
		Store(context.Context, OrderDto) error
		GetById(context.Context, string) (OrderDto, error)
		GetAll(ctx context.Context) ([]OrderDto, error)
	}

	// OrderBroker -.
	OrderBroker interface {
		ConsumeOrder() (OrderDto, error)
	}
)
