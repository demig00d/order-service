// Package usecase implements application business logic. Each logic group in own file.
package usecase

import (
	"context"
	"github.com/demig00d/order-service/internal/entity"
	"github.com/demig00d/order-service/internal/repo"
)

//go:generate mockgen -source=interfaces.go -destination=./mocks_test.go -package=usecase_test

type (
	// Order -.
	Order interface {
		ReceiveOrder(context.Context) error
		GetById(context.Context, string) (entity.Order, error)
	}

	// OrderDb -.
	OrderDb interface {
		Store(context.Context, repo.OrderDto) error
		GetById(context.Context, string) (repo.OrderDto, error)
		RecoverCache(cacheCapacity uint64) error
	}

	// OrderBroker -.
	OrderBroker interface {
		ConsumeOrder(ctx context.Context, db OrderDb) error
	}
)
