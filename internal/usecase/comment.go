package usecase

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/okankaraduman/goFinalApp/internal/entity"
)

// CommentUseCase -.
type CommentUseCase struct {
	repo   CommentRepo
	webAPI CommentWebAPI
}

// New -.
func New(r CommentRepo, w CommentWebAPI) *CommentUseCase {
	return &CommentUseCase{
		repo:   r,
		webAPI: w,
	}
}

func (c *CommentUseCase) CreateReview(request entity.CreateReviewRequest) (*entity.ReviewDTO, error) {

	r := entity.Review{
		Id:               uuid.New().String(),
		UserId:           request.UserId,
		ContentId:        request.ContentId,
		Rate:             request.Rate,
		Comment:          request.Comment,
		UserName:         request.UserName,
		CreatedDate:      time.Now().UnixMilli(),
		LastModifiedDate: time.Now().UnixMilli(),
	}

	ctx := context.Background()
	err := c.repo.InsertReview(ctx, r)

	if err != nil {
		return nil, err
	}
	reviewDTO := convertReview(r)

	return reviewDTO, nil
}

/*
// History - getting translate history from store.
func (uc *TranslationUseCase) History(ctx context.Context) ([]entity.Translation, error) {
	translations, err := uc.repo.GetHistory(ctx)
	if err != nil {
		return nil, fmt.Errorf("TranslationUseCase - History - s.repo.GetHistory: %w", err)
	}

	return translations, nil
}

// Translate -.
func (uc *TranslationUseCase) Translate(ctx context.Context, t entity.Translation) (entity.Translation, error) {
	translation, err := uc.webAPI.Translate(t)
	if err != nil {
		return entity.Translation{}, fmt.Errorf("TranslationUseCase - Translate - s.webAPI.Translate: %w", err)
	}

	err = uc.repo.Store(context.Background(), translation)
	if err != nil {
		return entity.Translation{}, fmt.Errorf("TranslationUseCase - Translate - s.repo.Store: %w", err)
	}

	return translation, nil
}

*/

func convertReview(review entity.Review) *entity.ReviewDTO {
	return &entity.ReviewDTO{
		Id:        review.Id,
		UserId:    review.UserId,
		ContentId: review.ContentId,
		Rate:      review.Rate,
		Comment:   review.Comment,
		UserName:  review.UserName,
	}
}
