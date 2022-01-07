package amqprpc

import (
	"github.com/okankaraduman/goFinalApp/internal/entity"
	"github.com/okankaraduman/goFinalApp/internal/usecase"
	"github.com/okankaraduman/goFinalApp/pkg/rabbitmq/rmq_rpc/server"
)

type authRoutes struct {
	authUseCase usecase.Auth
}

func newTranslationRoutes(routes map[string]server.CallHandler, t usecase.Auth) {
	r := &authRoutes{t}
	{
		routes["getHistory"] = r.he()
	}
}

type historyResponse struct {
	History []entity.User `json:"history"`
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
