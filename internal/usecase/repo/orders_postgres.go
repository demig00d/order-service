package repo

import (
	"context"
	"fmt"
	"github.com/demig00d/order-service/internal/entity"
	"github.com/demig00d/order-service/pkg/postgres"
	"github.com/jackc/pgx/v4"
)

// OrderRepo -.
type OrderRepo struct {
	*postgres.Postgres
}

// New -.
func New(pg *postgres.Postgres) *OrderRepo {
	return &OrderRepo{pg}
}

// SelectById -.
func (r *OrderRepo) SelectById(ctx context.Context, id string) (entity.Order, error) {
	sql, _, err := r.Builder.
		Select("*").
		From(`orders`).
		Where("order_uid=$1").
		ToSql()

	if err != nil {
		return entity.Order{}, fmt.Errorf("OrderRepo - GetHistory - r.Builder: %w", err)
	}

	var (
		orderUid  string
		orderInfo []byte
	)

	err = r.Pool.QueryRow(ctx, sql, id).Scan(&orderUid, &orderInfo)

	if err != nil {
		if err == pgx.ErrNoRows {
			return entity.Order{}, fmt.Errorf("OrderRepo - SelectById - r.Pool.Query: %w", err)
		}
		return entity.Order{}, fmt.Errorf("OrderRepo - SelectById - r.Pool.Query: %w", err)
	}

	return entity.Order{orderUid, orderInfo}, nil
}

// Store -.
func (r *OrderRepo) Store(ctx context.Context, o entity.Order) error {
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
