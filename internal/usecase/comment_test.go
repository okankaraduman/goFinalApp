package usecase_test

import (
	"context"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"

	"github.com/okankaraduman/goFinalApp/internal/entity"
	"github.com/okankaraduman/goFinalApp/internal/usecase"
)

var errInternalServErr = errors.New("internal server error")

type test struct {
	name string
	mock func()
	res  interface{}
	err  error
}

func comment(t *testing.T) (*usecase.CommentUseCase, *MockCommentRepo, *MockCommentWebAPI) {
	t.Helper()

	mockCtl := gomock.NewController(t)
	defer mockCtl.Finish()

	repo := NewMockCommentRepo(mockCtl)
	webAPI := NewMockCommentWebAPI(mockCtl)

	comment := usecase.New(repo, webAPI)

	return comment, repo, webAPI
}
func TestCreateReview(t *testing.T) {
	t.Parallel()

	comment, repo, _ := comment(t)

	tests := []test{
		{
			name: "result with error",
			mock: func() {
				repo.EXPECT().InsertReview(context.Background(), entity.Review{}).Return(errInternalServErr)
			},
			res: nil,
			err: errInternalServErr,
		},
	}

	for _, tc := range tests {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			tc.mock()

			_, err := comment.CreateReview(entity.CreateReviewRequest{})

			require.ErrorIs(t, err, tc.err, "The two errors should be the same.")
		})
	}
}
