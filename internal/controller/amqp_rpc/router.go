package amqprpc

import (
	"github.com/Valdym/goFinalApp/internal/usecase"
	"github.com/Valdym/goFinalApp/pkg/rabbitmq/rmq_rpc/server"
)

// NewRouter -.
func NewRouter(t usecase.Translation) map[string]server.CallHandler {
	routes := make(map[string]server.CallHandler)
	{
		newTranslationRoutes(routes, t)
	}

	return routes
}
