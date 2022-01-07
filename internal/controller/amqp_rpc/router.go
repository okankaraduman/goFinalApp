package amqprpc

import (
	"github.com/okankaraduman/goFinalApp/internal/usecase"
	"github.com/okankaraduman/goFinalApp/pkg/rabbitmq/rmq_rpc/server"
)

// NewRouter -.
func NewRouter(a usecase.Comment) map[string]server.CallHandler {
	routes := make(map[string]server.CallHandler)
	{
		newCommentRoutes(routes, a)
	}

	return routes
}
