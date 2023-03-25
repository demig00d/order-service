package broker

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/demig00d/order-service/internal/entity"
	"github.com/demig00d/order-service/internal/repo"
	"github.com/demig00d/order-service/internal/usecase"
	"github.com/demig00d/order-service/pkg/jetstream"
	"github.com/nats-io/nats.go"
)

type OrderBroker struct {
	js *jetstream.JetStream
}

func New(js *jetstream.JetStream) *OrderBroker {
	return &OrderBroker{js}
}

func (b *OrderBroker) ConsumeOrder(ctx context.Context, db usecase.OrderDb) error {
	var order entity.Order
	var subsError error

	_, err := b.js.Subscribe(b.js.SubjectNameOrderCreated, func(m *nats.Msg) {
		err := m.Ack()

		if err != nil {
			subsError = fmt.Errorf("OrderBroker - ConsumeOrder - unable to ack: %w", err)
			return
		}

		err = json.Unmarshal(m.Data, &order)
		if err != nil {
			subsError = fmt.Errorf("OrderBroker - ConsumeOrder - can't unmarshal message to entity.Order: %w", err)
			return
		}

		orderDto :=
			repo.OrderDto{
				OrderUid:  order.OrderUid,
				OrderInfo: m.Data,
			}

		subsError = db.Store(ctx, orderDto)
	})

	if subsError != nil {
		return subsError
	}
	if err != nil {
		return err
	}
	return nil

}
