package v1

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

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

// @Summary     Get Reviews
// @Description Returns all comments
// @ID          getReviews
// @Tags  	    comment
// @Accept      json
// @Produce     json
// @Success     200 {array} entity.ReviewDTO
// @Router      /v1/comments [get]
func (c *commentRoutes) getReviews(w http.ResponseWriter, r *http.Request) {
	//Get All comments
	arr, err := c.c.TakeReviews()
	if err != nil {
		c.l.Error(err, "http - v1 - getReviews")
		resp := Response{Resp: w}
		resp.Text(http.StatusInternalServerError, "Get Reviews service problems", "text/plain")

		return
	}
	body, err := json.Marshal(arr)
	c.l.Debug(arr)
	if err != nil {
		c.l.Error(err, "http - v1 - getReviews")
	}

	resp := Response{Resp: w}
	resp.Text(http.StatusOK, string(body), "text/json")

}

// @Summary     Delete Review
// @Description Deletes comments by its id
// @ID          deleteReview
// @Tags  	    comment
// @Accept      json
// @Produce     text
// @Success     200 {text}
// @Router      /v1/comments/{commentID} [delete]
func (c *commentRoutes) deleteReview(w http.ResponseWriter, r *http.Request) {
	url := r.URL.Path
	ourid := strings.SplitN(url, "/", 6)[3]
	c.l.Debug(ourid)

	err := c.c.DeleteReview(ourid)
	if err != nil {
		c.l.Error(err, "http - v1 - deleteReview")
		resp := Response{Resp: w}
		resp.Text(http.StatusInternalServerError, "Delete Review service problems", "text/plain")

		return
	}
	resp := Response{Resp: w}
	resp.Text(http.StatusOK, "Resource successfully deleted!", "text/plaing")
}

// @Summary     Insert Review
// @Description creates review
// @ID          insertReview
// @Tags  	    comment
// @Accept      json
// @Produce     json
// @Success     200 {object} entity.ReviewDTO
// @Router      /v1/comments [post]
func (c *commentRoutes) insertReview(w http.ResponseWriter, r *http.Request) {
	var requestBody entity.CreateReviewRequest
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

	resp := Response{Resp: w}
	resp.Text(http.StatusOK, string(body), "text/json")
}
func (c *commentRoutes) updateReview(w http.ResponseWriter, r *http.Request) {
	fmt.Println("sadas")
}
func (c *commentRoutes) getReview(w http.ResponseWriter, r *http.Request) {
	//Get All comments
}
