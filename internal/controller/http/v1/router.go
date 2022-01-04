// Package v1 implements routing paths. Each services in own file.
package v1

import (
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	httpSwagger "github.com/swaggo/http-swagger"

	// Swagger docs.

	_ "github.com/okankaraduman/goFinalApp/docs"
	"github.com/okankaraduman/goFinalApp/internal/usecase"
	"github.com/okankaraduman/goFinalApp/pkg/logger"
)

// NewRouter -.
// Swagger spec:
// @title       Go Clean Template
// @description Using a translation service as an example
// @version     1.0
// @host        localhost:8080
// @BasePath    /v1
func NewRouter(handler *gin.Engine, l logger.Interface, t usecase.Translation) {

	r := chi.NewRouter()

	// A good base middleware stack
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// Set a timeout value on the request context (ctx), that will signal
	// through ctx.Done() that the request has timed out and further
	// processing should be stopped.
	r.Use(middleware.Timeout(60 * time.Second))

	// Swagger
	swaggerHandler := httpSwagger.Handler()

	r.Get("/swagger/*any", swaggerHandler)

	// K8s probe
	r.Get("/healthz", func(w http.ResponseWriter, r *http.Request) {
		resp := Response{Resp: w}
		resp.Text(http.StatusOK, "404 Not Found", "text/plain")
	})

	// Prometheus metrics
	r.Get("/metrics", func(w http.ResponseWriter, r *http.Request) {
		return promhttp.Handler()
	})

}
