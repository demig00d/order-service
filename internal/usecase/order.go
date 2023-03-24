package usecase

import (
	"context"
	"fmt"
	lru "github.com/hashicorp/golang-lru/v2"
)

// OrderUseCase -.
type OrderUseCase struct {
	repo   OrderRepo
	cache  *lru.TwoQueueCache[string, OrderDto]
	broker OrderBroker
}

// New -.
func New(r OrderRepo, b OrderBroker) *OrderUseCase {
	return &OrderUseCase{
		repo:   r,
		broker: b,
	}
}

// GetById - getting single order from store.
func (uc *OrderUseCase) GetById(ctx context.Context, id string) (OrderDto, error) {
	order, err := uc.repo.GetById(ctx, id)
	if err != nil {
		return order, fmt.Errorf("OrderUseCase - order - s.repo.GetById: %w", err)
	}

	return order, nil
}

// ReceiveOrder -.
func (uc *OrderUseCase) ReceiveOrder(ctx context.Context) error {
	order, err := uc.broker.ConsumeOrder()

	if err != nil {
		return err
	}

	err = uc.repo.Store(ctx, order)
	if err != nil {
		return fmt.Errorf("OrderUseCase - Store - s.repo.Store: %w", err)
	}

	return nil
}
