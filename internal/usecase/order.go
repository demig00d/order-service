package usecase

import (
	"context"
	"fmt"
	"github.com/demig00d/order-service/internal/entity"
	lru "github.com/hashicorp/golang-lru/v2"
)

// OrderUseCase -.
type OrderUseCase struct {
	repo  OrderRepo
	cache *lru.TwoQueueCache[string, entity.Order]
}

// New -.
func New(r OrderRepo) *OrderUseCase {
	return &OrderUseCase{
		repo: r,
	}
}

// GetById - getting single order from store.
func (uc *OrderUseCase) GetById(ctx context.Context, id string) (entity.Order, error) {
	order, err := uc.repo.GetById(ctx, id)
	if err != nil {
		return order, fmt.Errorf("OrderUseCase - order - s.repo.GetById: %w", err)
	}

	return order, nil
}

// Store -.
func (uc *OrderUseCase) Store(ctx context.Context, order entity.Order) error {
	err := uc.repo.Store(ctx, order)
	if err != nil {
		return fmt.Errorf("OrderUseCase - Store - s.repo.Store: %w", err)
	}

	return nil
}
