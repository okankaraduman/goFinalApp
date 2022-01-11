package usecase

import (
	"context"
	"fmt"
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

	//TO DO: Translate the comment before inserting to database
	r := entity.Review{
		Id:               uuid.New().String(),
		UserId:           request.UserId,
		ContentId:        request.ContentId,
		ReviewStatus:     "true",
		Rate:             request.Rate,
		Comment:          request.Comment,
		UserName:         request.UserName,
		CreatedDate:      time.Now().UnixMilli(),
		LastModifiedDate: time.Now().UnixMilli(),
	}

	ctx := context.Background()
	err := c.repo.InsertReview(ctx, r)

	if err != nil {
		return nil, fmt.Errorf("CommentUseCase - CreateReview - s.repo.InsertReview: %w", err)
	}
	reviewDTO := convertReview(r)
	if err != nil {
		return nil, fmt.Errorf("CommentUseCase - TakeReviews - convertReview: %w", err)
	}

	return reviewDTO, nil
}
func (c *CommentUseCase) TakeReviews() (*[]entity.ReviewDTO, error) {
	ctx := context.Background()
	arr, err := c.repo.GetReviews(ctx)
	if err != nil {
		return nil, fmt.Errorf("CommentUseCase - TakeReviews - s.repo.GetReviews: %w", err)
	}
	var arr_dto []entity.ReviewDTO
	for i := 0; i < len(arr); i++ {
		temp := convertReview(arr[i])
		if err != nil {
			return nil, fmt.Errorf("CommentUseCase - TakeReviews - convertReview: %w", err)
		}
		arr_dto = append(arr_dto, *temp)
	}
	return &arr_dto, nil
}
func (c *CommentUseCase) DeleteReview(uid string) error {
	err := c.repo.DeleteReview(context.Background(), uid)
	if err != nil {
		return fmt.Errorf("CommentUseCase - DeleteReview - s.repo.DeleteReview: %w", err)
	}
	return nil
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
