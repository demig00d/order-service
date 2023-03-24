package repo

import (
	"context"
	"fmt"
	"github.com/demig00d/order-service/pkg/postgres"
	lru "github.com/hashicorp/golang-lru/v2"
	"github.com/jackc/pgx/v4"
)

// OrderRepo -.
type OrderRepo struct {
	*postgres.Postgres
	cache *lru.TwoQueueCache[string, OrderDto]
}

// New -.
func New(pg *postgres.Postgres, cache *lru.TwoQueueCache[string, OrderDto]) *OrderRepo {
	return &OrderRepo{pg, cache}
}

func (r *OrderRepo) GetById(ctx context.Context, id string) (OrderDto, error) {
	order, ok := r.cache.Get(id)

	if ok {
		return order, nil
	}

	order, err := r.SelectById(ctx, id)

	if err != nil {
		return order, err
	}

	r.cache.Add(id, order)

	return order, nil

}

// SelectById -.
func (r *OrderRepo) SelectById(ctx context.Context, id string) (OrderDto, error) {
	sql, _, err := r.Builder.
		Select("*").
		From(`orders`).
		Where("order_uid=$1").
		ToSql()

	if err != nil {
		return OrderDto{}, fmt.Errorf("OrderRepo - GetHistory - r.Builder: %w", err)
	}

	var (
		orderUid  string
		orderInfo []byte
	)

	err = r.Pool.QueryRow(ctx, sql, id).Scan(&orderUid, &orderInfo)

	if err != nil {
		if err == pgx.ErrNoRows {
			return OrderDto{}, fmt.Errorf("OrderRepo - SelectById - r.Pool.Query: %w", err)
		}
		return OrderDto{}, fmt.Errorf("OrderRepo - SelectById - r.Pool.Query: %w", err)
	}

	return OrderDto{orderUid, orderInfo}, nil
}

// Store -.
func (r *OrderRepo) Store(ctx context.Context, o OrderDto) error {
	r.cache.Add(o.OrderUid, o)
	err := r.Insert(ctx, o)

	return err
}

// Insert -.
func (r *OrderRepo) Insert(ctx context.Context, o OrderDto) error {
	sql, args, err := r.Builder.
		Insert("orders").
		Columns("order_uid", `"order"`).
		Values(o.OrderUid, o.OrderInfo).
		ToSql()
	if err != nil {
		return fmt.Errorf("OrderRepo - Store - r.Builder: %w", err)
	}

	_, err = r.Pool.Exec(ctx, sql, args...)
	if err != nil {
		return fmt.Errorf("OrderRepo - Store - r.Pool.Exec: %w", err)
	}

	return nil
}
