package amqprpc

import (
	"github.com/demig00d/order-service/internal/usecase"
	"github.com/demig00d/order-service/pkg/rabbitmq/rmq_rpc/server"
)

// NewRouter -.
func NewRouter(t usecase.Translation) map[string]server.CallHandler {
	routes := make(map[string]server.CallHandler)
	{
		newTranslationRoutes(routes, t)
	}

	return routes
}
