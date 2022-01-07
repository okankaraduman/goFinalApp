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
		return fmt.Errorf("CommentRepo - InsertReview - c.Pool.Exec: %w", err)
	}
	return nil
}
