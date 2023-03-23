// Package usecase implements application business logic. Each logic group in own file.
package usecase

import (
	"context"

	"github.com/demig00d/order-service/internal/entity"
)

//go:generate mockgen -source=interfaces.go -destination=./mocks_test.go -package=usecase_test

type (
	// Order -.
	Order interface {
		Store(context.Context, entity.Order) error
		GetById(context.Context, string) (entity.Order, error)
	}

	// OrderRepo -.
	OrderRepo interface {
		Store(context.Context, entity.Order) error
		GetById(context.Context, string) (entity.Order, error)
	}
)
