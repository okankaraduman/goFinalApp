// Package app configures and runs application.
package app

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/go-chi/chi/v5"
	"github.com/okankaraduman/goFinalApp/config"
	amqprpc "github.com/okankaraduman/goFinalApp/internal/controller/amqp_rpc"
	v1 "github.com/okankaraduman/goFinalApp/internal/controller/http/v1"
	"github.com/okankaraduman/goFinalApp/internal/usecase"
	"github.com/okankaraduman/goFinalApp/internal/usecase/repo"
	"github.com/okankaraduman/goFinalApp/internal/usecase/webapi"
	"github.com/okankaraduman/goFinalApp/pkg/httpserver"
	"github.com/okankaraduman/goFinalApp/pkg/logger"
	"github.com/okankaraduman/goFinalApp/pkg/postgres"
	"github.com/okankaraduman/goFinalApp/pkg/rabbitmq/rmq_rpc/server"
)

// Run creates objects via constructors.
func Run(cfg *config.Config) {
	l := logger.New(cfg.Log.Level)

	// Repository
	pg, err := postgres.New(cfg.PG.URL, postgres.MaxPoolSize(cfg.PG.PoolMax))
	if err != nil {
		l.Fatal(fmt.Errorf("app - Run - postgres.New: %w", err))
	}
	defer pg.Close()

	// Use case
	commentUseCase := usecase.New(
		repo.New(pg),
		webapi.New(),
	)

	//RabbitMQ RPC Server
	rmqRouter := amqprpc.NewRouter(commentUseCase)

	rmqServer, err := server.New(cfg.RMQ.URL, cfg.RMQ.ServerExchange, rmqRouter, l)
	if err != nil {
		l.Fatal(fmt.Errorf("app - Run - rmqServer - server.New: %w", err))
	}

	// HTTP Server
	r := chi.NewRouter()
	v1.NewRouter(r, l, commentUseCase)
	httpServer := httpserver.New(r, httpserver.Port(cfg.HTTP.Port))
	//Router

	// Waiting signal
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		l.Info("app - Run - signal: " + s.String())
	case err = <-httpServer.Notify():
		l.Error(fmt.Errorf("app - Run - httpServer.Notify: %w", err))
	case err = <-rmqServer.Notify():
		l.Error(fmt.Errorf("app - Run - rmqServer.Notify: %w", err))
	}

	// Shutdown
	err = httpServer.Shutdown()
	if err != nil {
		l.Error(fmt.Errorf("app - Run - httpServer.Shutdown: %w", err))
	}

	err = rmqServer.Shutdown()
	if err != nil {
		l.Error(fmt.Errorf("app - Run - rmqServer.Shutdown: %w", err))
	}
}
