// Package usecase implements application business logic. Each logic group in own file.
package usecase

import (
	"context"
	"github.com/demig00d/order-service/internal/usecase/repo"
)

//go:generate mockgen -source=interfaces.go -destination=./mocks_test.go -package=usecase_test

type (
	// Order -.
	Order interface {
		Store(context.Context, repo.OrderDto) error
		GetById(context.Context, string) (repo.OrderDto, error)
	}

	// OrderRepo -.
	OrderRepo interface {
		Store(context.Context, repo.OrderDto) error
		GetById(context.Context, string) (repo.OrderDto, error)
	}
)
