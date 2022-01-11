package repo

import (
	"context"
	"fmt"

	"github.com/okankaraduman/goFinalApp/internal/entity"
	"github.com/okankaraduman/goFinalApp/pkg/postgres"
)

const _defaultEntityCap = 64

// CommentRepo -.
type CommentRepo struct {
	*postgres.Postgres
}

// New -.
func New(pg *postgres.Postgres) *CommentRepo {
	return &CommentRepo{pg}
}

// InsertReview -.
func (c *CommentRepo) InsertReview(ctx context.Context, r entity.Review) error {
	sql, args, err := c.Builder.
		Insert("Reviews").
		Columns("Id, UserId, ContentId, ReviewStatus, Rate, Comment, UserName,CreatedDate,LastModifiedDate").
		Values(r.Id, r.UserId, r.ContentId, r.ReviewStatus, r.Rate, r.Comment, r.UserName, r.CreatedDate, r.LastModifiedDate).
		ToSql()
	if err != nil {
		return fmt.Errorf("CommentRepo - InsertReview - c.Builder: %w", err)
	}

	_, err = c.Pool.Exec(ctx, sql, args...)
	if err != nil {
		return fmt.Errorf("CommentRepo - InsertReview - c.Pool.Exec: %w sql:"+sql, err)
	}
	return nil
}

// GetReviews -.
func (c *CommentRepo) GetReviews(ctx context.Context) ([]entity.Review, error) {
	sql, args, err := c.Builder.
		Select("*").
		From("Reviews").
		ToSql()
	if err != nil {
		return nil, fmt.Errorf("CommentRepo - GetReviews - c.Builder: %w sql:"+sql, err)
	}
	rows, err := c.Pool.Query(ctx, sql, args...)
	if err != nil {
		return nil, fmt.Errorf("CommentRepo - GetReviews - c.Pool.Exec: %w sql:"+sql, err)
	}
	defer rows.Close()
	var array []entity.Review
	for rows.Next() {
		var temp entity.Review
		rows.Scan(&temp.Id, &temp.UserId, &temp.ContentId, &temp.ReviewStatus, &temp.Rate, &temp.Comment, &temp.UserName, &temp.CreatedDate, &temp.LastModifiedDate)
		if err != nil {
			return nil, fmt.Errorf("CommentRepo - GetReviews - rows.Scan: %w ", err)
		}
		array = append(array, temp)
	}

	return array, nil
}

// DeleteReview -.
func (c *CommentRepo) DeleteReview(ctx context.Context, uid string) error {
	sql, args, err := c.Builder.
		Delete("*").
		Where("Id=$1", uid).
		From("Reviews").
		ToSql()
	if err != nil {
		return fmt.Errorf("CommentRepo - DeleteReview - c.Builder: %w sql:"+sql, err)
	}
	_, err = c.Pool.Exec(ctx, sql, args...)
	if err != nil {
		return fmt.Errorf("CommentRepo - DeleteReview - c.Pool.Exec: %w sql:"+sql, err)
	}

	return nil
}
