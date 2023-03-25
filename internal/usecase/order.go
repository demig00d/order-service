package usecase

import (
	"context"
	"fmt"
	"github.com/demig00d/order-service/internal/entity"
	"github.com/demig00d/order-service/internal/repo"
	lru "github.com/hashicorp/golang-lru/v2"
)

// OrderUseCase -.
type OrderUseCase struct {
	repo   OrderDb
	cache  *lru.TwoQueueCache[string, repo.OrderDto]
	broker OrderBroker
}

// New -.
func New(r OrderDb, b OrderBroker) *OrderUseCase {
	return &OrderUseCase{
		repo:   r,
		broker: b,
	}
}

// GetById - getting single order from store.
func (uc *OrderUseCase) GetById(ctx context.Context, id string) (entity.Order, error) {
	orderDto, err := uc.repo.GetById(ctx, id)
	if err != nil {
		return entity.Order{}, fmt.Errorf("OrderUseCase - order - s.repo.GetById: %w", err)
	}
	return orderDto.ToEntity()
}

// ReceiveOrder -.
func (uc *OrderUseCase) ReceiveOrder(ctx context.Context) error {
	return uc.broker.ConsumeOrder(ctx, uc.repo)
}
