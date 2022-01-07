package v1

import (

	// Swagger docs.

	"fmt"
	"net/http"

	"github.com/okankaraduman/goFinalApp/internal/usecase"

	"github.com/go-chi/chi/v5"
	_ "github.com/okankaraduman/goFinalApp/docs"
	"github.com/okankaraduman/goFinalApp/pkg/logger"
)

type authRoutes struct {
	t usecase.Auth
	l logger.Interface
}

func newAuthRoutes(a usecase.Auth, l logger.Interface) http.Handler {
	auth := &authRoutes{a, l}
	r := chi.NewRouter()

	r.Get("/", auth.hello)

	r.Route("/{userID}", func(r chi.Router) {
		//r.Get("/", getArticle)
		// r.Put("/", updateArticle)
		// r.Delete("/", deleteArticle)
	})
	return r
}

// @Summary     Show history
// @Description Show all translation history
// @ID          history
// @Tags  	    translation
// @Accept      json
// @Produce     json
// @Success     200 {object} historyResponse
// @Failure     500 {object} response
// @Router      /translation/history [get]
func (r *authRoutes) hello(w http.ResponseWriter, r *http.Request) {
	fmt.Println("sadas")
}
