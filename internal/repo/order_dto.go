package repo

import (
	"encoding/json"
	"github.com/demig00d/order-service/internal/entity"
)

type OrderDto struct {
	OrderUid  string
	OrderInfo []byte
}

func (orderDto *OrderDto) ToEntity() (entity.Order, error) {
	var order entity.Order
	err := json.Unmarshal(orderDto.OrderInfo, &order)

	return order, err
}
