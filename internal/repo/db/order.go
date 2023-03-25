package db

import (
	"context"
	"fmt"
	"github.com/demig00d/order-service/internal/repo"
	"github.com/demig00d/order-service/pkg/postgres"
	"github.com/georgysavva/scany/pgxscan"
	lru "github.com/hashicorp/golang-lru/v2"
	"github.com/jackc/pgx/v4"
)

// OrderDb -.
type OrderDb struct {
	*postgres.Postgres
	cache *lru.TwoQueueCache[string, repo.OrderDto]
}

// New -.
func New(pg *postgres.Postgres, cache *lru.TwoQueueCache[string, repo.OrderDto]) *OrderDb {
	return &OrderDb{pg, cache}
}

// GetById -.
func (r *OrderDb) GetById(ctx context.Context, id string) (repo.OrderDto, error) {
	var err error

	orderDto, isOrderInCache := r.cache.Get(id)

	if !isOrderInCache {
		orderDto, err = r.SelectById(ctx, id)

		if err != nil {
			return repo.OrderDto{}, nil
		}
	}

	if err == nil && !isOrderInCache {
		r.cache.Add(id, orderDto)
	}

	return orderDto, nil

}

// SelectById -.
func (r *OrderDb) SelectById(ctx context.Context, id string) (repo.OrderDto, error) {
	sql, _, err := r.Builder.
		Select("*").
		From(`orders`).
		Where("order_uid=$1").
		ToSql()

	if err != nil {
		return repo.OrderDto{}, fmt.Errorf("OrderDb - SelectById - r.Builder: %w", err)
	}

	var (
		orderUid  string
		orderInfo []byte
	)

	err = r.Pool.QueryRow(ctx, sql, id).Scan(&orderUid, &orderInfo)

	if err != nil {
		return repo.OrderDto{}, err
	}

	return repo.OrderDto{
		OrderUid:  orderUid,
		OrderInfo: orderInfo,
	}, nil
}

func (r *OrderDb) RecoverCache(ctx context.Context, cacheCapacity uint64) error {
	ordersDto, err := r.Select(ctx, cacheCapacity)

	if err != nil {
		return err
	}

	for _, orderDto := range ordersDto {
		r.cache.Add(orderDto.OrderUid, orderDto)
	}

	return err
}

func (r *OrderDb) Select(ctx context.Context, limit uint64) ([]repo.OrderDto, error) {
	sql, _, err := r.Builder.
		Select("*").
		From("orders").
		Limit(limit).
		ToSql()

	if err != nil {
		return nil, fmt.Errorf("OrderDb - SelectById - r.Pool.Query: %w", err)
	}

	var ordersDto []repo.OrderDto

	err = pgxscan.Select(ctx, r.Pool, &ordersDto, sql)

	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, fmt.Errorf("OrderDb - SelectById - r.Pool.Query: %w", err)
		}
		return nil, fmt.Errorf("OrderDb - SelectById - r.Pool.Query: %w", err)
	}

	return ordersDto, nil
}

// Store -.
func (r *OrderDb) Store(ctx context.Context, o repo.OrderDto) error {
	r.cache.Add(o.OrderUid, o)
	err := r.Insert(ctx, o)

	return err
}

// Insert -.
func (r *OrderDb) Insert(ctx context.Context, o repo.OrderDto) error {
	sql, args, err := r.Builder.
		Insert("orders").
		Columns("order_uid", `"order"`).
		Values(o.OrderUid, o.OrderInfo).
		ToSql()
	if err != nil {
		return fmt.Errorf("OrderDb - Store - r.Builder: %w", err)
	}

	_, err = r.Pool.Exec(ctx, sql, args...)
	if err != nil {
		return fmt.Errorf("OrderDb - Store - r.Pool.Exec: %w", err)
	}

	return nil
}
