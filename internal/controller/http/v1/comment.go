package v1

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/okankaraduman/goFinalApp/internal/entity"
	"github.com/okankaraduman/goFinalApp/internal/usecase"

	"github.com/go-chi/chi/v5"
	_ "github.com/okankaraduman/goFinalApp/docs"
	"github.com/okankaraduman/goFinalApp/pkg/logger"
)

type commentRoutes struct {
	c usecase.Comment
	l logger.Interface
}

func newCommentRoutes(a usecase.Comment, l logger.Interface) http.Handler {
	comment := &commentRoutes{a, l}
	r := chi.NewRouter()

	r.Get("/", comment.getReviews)
	r.Post("/", comment.insertReview)
	r.Route("/{commentID}", func(r chi.Router) {
		r.Get("/", comment.getReview)
		r.Put("/", comment.updateReview)
		r.Delete("/", comment.deleteReview)
	})
	return r
}

// @Summary     Returns all comments
// @Description Says hello
// @ID          hello
// @Tags  	    comment
// @Accept      json
// @Produce     json
// @Success     200 {array} entity.ReviewDTO
// @Router      /comment/hello [get]
func (c *commentRoutes) getReviews(w http.ResponseWriter, r *http.Request) {
	//Get All comments
}
func (c *commentRoutes) deleteReview(w http.ResponseWriter, r *http.Request) {
	fmt.Println("sadas")
}

// @Summary     creates review
// @Description Says hello
// @ID          hello
// @Tags  	    comment
// @Accept      json
// @Produce     json
// @Success     200 {object} entity.ReviewDTO
// @Router      /comment/hello [get]

func (c *commentRoutes) insertReview(w http.ResponseWriter, r *http.Request) {
	var requestBody entity.CreateReviewRequest
	c.l.Debug(r.Body)
	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		c.l.Error(err, "http - v1 - insertReview")
		resp := Response{Resp: w}
		resp.Text(http.StatusBadRequest, "Invalid JSON Format", "text/plain")

		return
	}
	var responseObject *entity.ReviewDTO
	responseObject, err = c.c.CreateReview(requestBody)
	if err != nil {
		c.l.Error(err, "http - v1 - insertReview")
		resp := Response{Resp: w}
		resp.Text(http.StatusInternalServerError, "Insert Review service problems", "text/plain")

		return
	}
	body, err := json.Marshal(responseObject)

	if err != nil {
		c.l.Error(err, "http - v1 - insertReview")
	}
	fmt.Println(string(body))

	resp := Response{Resp: w}
	resp.Text(http.StatusOK, string(body), "text/json")
}
func (c *commentRoutes) updateReview(w http.ResponseWriter, r *http.Request) {
	fmt.Println("sadas")
}
func (c *commentRoutes) getReview(w http.ResponseWriter, r *http.Request) {
	//Get All comments
}
