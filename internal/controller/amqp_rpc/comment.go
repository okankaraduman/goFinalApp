package amqprpc

import (
	"github.com/okankaraduman/goFinalApp/internal/entity"
	"github.com/okankaraduman/goFinalApp/internal/usecase"
	"github.com/okankaraduman/goFinalApp/pkg/rabbitmq/rmq_rpc/server"
)

type commentRoutes struct {
	commentUseCase usecase.Comment
}

func newCommentRoutes(routes map[string]server.CallHandler, t usecase.Comment) {
	//
}

type commentResponse struct {
	Reviews []entity.Review `json:"review"`
}

/*
func (a *authRoutes) getHistory() server.CallHandler {
	return func(d *amqp.Delivery) (interface{}, error) {
		translations, err := a.translationUseCase.History(context.Background())
		if err != nil {
			return nil, fmt.Errorf("amqp_rpc - translationRoutes - getHistory - r.translationUseCase.History: %w", err)
		}

		response := historyResponse{translations}

		return response, nil
	}
}
*/
