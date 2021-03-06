// Package usecase implements application business logic. Each logic group in own file.
package usecase

import (
	"context"

	"github.com/okankaraduman/goFinalApp/internal/entity"
)

//go:generate mockgen -source=interfaces.go -destination=./mocks_test.go -package=usecase_test

type (
	// Comment -.
	Comment interface {
		CreateReview(request entity.CreateReviewRequest) (*entity.ReviewDTO, error)
		TakeReviews() (*[]entity.ReviewDTO, error)
	}

	// CommentRepo -.
	CommentRepo interface {
		InsertReview(context.Context, entity.Review) error
		GetReviews(context.Context) ([]entity.Review, error)
		//reponun içindelileri
	}

	// CommentWebAPI -.
	CommentWebAPI interface {
		//Web API'nin içindeki fonksiyonlar
	}
)
