// Package v1 implements routing paths. Each services in own file.
package v1

import (
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"github.com/prometheus/client_golang/prometheus/promhttp"
	httpSwagger "github.com/swaggo/http-swagger"

	// Swagger docs.

	_ "github.com/okankaraduman/goFinalApp/docs"
	_ "github.com/okankaraduman/goFinalApp/internal/usecase"
	"github.com/okankaraduman/goFinalApp/pkg/logger"
)

// NewRouter -.
// Swagger spec:
// @title       Go Order App
// @description Creates an
// @version     1.0
// @host        localhost:8080
// @BasePath    /v1
func NewRouter(muxChi *chi.Mux, l logger.Interface) {

	fs := http.FileServer(http.Dir("./././static"))
	fmt.Println(http.Dir("../../static"))
	// A good base middleware stack
	muxChi.Use(middleware.RequestID)
	muxChi.Use(middleware.RealIP)
	muxChi.Use(middleware.Logger)
	muxChi.Use(middleware.Recoverer)

	// Set a timeout value on the request context (ctx), that will signal
	// through ctx.Done() that the request has timed out and further
	// processing should be stopped.
	muxChi.Use(middleware.Timeout(60 * time.Second))

	//Fileserver

	muxChi.Handle("/", addHeaders(fs))

	// Swagger
	swaggerHandler := httpSwagger.Handler()

	muxChi.Get("/swagger/*", swaggerHandler)

	// K8s probe
	muxChi.Get("/healthz", func(w http.ResponseWriter, r *http.Request) {
		resp := Response{Resp: w}
		resp.Text(http.StatusOK, "200 OK", "text/plain")
	})
	// Prometheus metrics
	muxChi.Handle("/metrics", promhttp.Handler())
}

func addHeaders(fs http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("X-Frame-Options", "DENY")
		fs.ServeHTTP(w, r)
	}
}
