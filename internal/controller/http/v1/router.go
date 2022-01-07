// Package v1 implements routing paths. Each services in own file.
package v1

import (
	"context"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"github.com/prometheus/client_golang/prometheus/promhttp"
	httpSwagger "github.com/swaggo/http-swagger"

	// Swagger docs.

	_ "github.com/okankaraduman/goFinalApp/docs"
	"github.com/okankaraduman/goFinalApp/internal/usecase"
	_ "github.com/okankaraduman/goFinalApp/internal/usecase"
	"github.com/okankaraduman/goFinalApp/pkg/logger"
)

// NewRouter -.
// Swagger spec:
// @title       Go Order App
// @description Creates a new router
// @version     0.1
// @host        localhost:8080
// @BasePath    /v1
func NewRouter(muxChi *chi.Mux, l logger.Interface, c usecase.Comment) {

	fs := http.FileServer(http.Dir("./././static"))

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

	// API version 1.
	muxChi.Route("/v1", func(r chi.Router) {
		r.Use(apiVersionCtx("v1"))
		r.Mount("/comments", newCommentRoutes(c, l))
	})
}

func apiVersionCtx(version string) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			r = r.WithContext(context.WithValue(r.Context(), "api.version", version))
			next.ServeHTTP(w, r)
		})
	}
}

func addHeaders(fs http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("X-Frame-Options", "DENY")
		fs.ServeHTTP(w, r)
	}
}
