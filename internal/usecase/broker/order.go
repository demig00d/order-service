package broker

import (
	"encoding/json"
	"github.com/demig00d/order-service/internal/entity"
	"github.com/demig00d/order-service/internal/usecase"
	"github.com/demig00d/order-service/pkg/jetstream"
	"github.com/nats-io/nats.go"
	"log"
)

type OrderBroker struct {
	js *jetstream.JetStream
}

func New(js *jetstream.JetStream) *OrderBroker {
	return &OrderBroker{js}
}

func (b *OrderBroker) ConsumeOrder() (usecase.OrderDto, error) {
	var order entity.Order
	var orderDto usecase.OrderDto

	_, err := b.js.Subscribe(b.js.SubjectNameOrderCreated, func(m *nats.Msg) {
		err := m.Ack()

		if err != nil {
			log.Println("Unable to Ack", err)
			return
		}

		err = json.Unmarshal(m.Data, &order)
		if err != nil {
			return
		}

		orderDto =
			usecase.OrderDto{
				OrderUid:  order.OrderUid,
				OrderInfo: m.Data,
			}

	})

	return orderDto, err
}
