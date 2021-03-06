// Code generated by MockGen. DO NOT EDIT.
// Source: interfaces.go

// Package mock_usecase is a generated GoMock package.
package usecase_test

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	entity "github.com/okankaraduman/goFinalApp/internal/entity"
)

// MockComment is a mock of Comment interface.
type MockComment struct {
	ctrl     *gomock.Controller
	recorder *MockCommentMockRecorder
}

// MockCommentMockRecorder is the mock recorder for MockComment.
type MockCommentMockRecorder struct {
	mock *MockComment
}

// NewMockComment creates a new mock instance.
func NewMockComment(ctrl *gomock.Controller) *MockComment {
	mock := &MockComment{ctrl: ctrl}
	mock.recorder = &MockCommentMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockComment) EXPECT() *MockCommentMockRecorder {
	return m.recorder
}

// CreateReview mocks base method.
func (m *MockComment) CreateReview(request entity.CreateReviewRequest) (*entity.ReviewDTO, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateReview", request)
	ret0, _ := ret[0].(*entity.ReviewDTO)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateReview indicates an expected call of CreateReview.
func (mr *MockCommentMockRecorder) CreateReview(request interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateReview", reflect.TypeOf((*MockComment)(nil).CreateReview), request)
}

// TakeReviews mocks base method.
func (m *MockComment) TakeReviews() (*[]entity.ReviewDTO, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TakeReviews")
	ret0, _ := ret[0].(*[]entity.ReviewDTO)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// TakeReviews indicates an expected call of TakeReviews.
func (mr *MockCommentMockRecorder) TakeReviews() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TakeReviews", reflect.TypeOf((*MockComment)(nil).TakeReviews))
}

// MockCommentRepo is a mock of CommentRepo interface.
type MockCommentRepo struct {
	ctrl     *gomock.Controller
	recorder *MockCommentRepoMockRecorder
}

// MockCommentRepoMockRecorder is the mock recorder for MockCommentRepo.
type MockCommentRepoMockRecorder struct {
	mock *MockCommentRepo
}

// NewMockCommentRepo creates a new mock instance.
func NewMockCommentRepo(ctrl *gomock.Controller) *MockCommentRepo {
	mock := &MockCommentRepo{ctrl: ctrl}
	mock.recorder = &MockCommentRepoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCommentRepo) EXPECT() *MockCommentRepoMockRecorder {
	return m.recorder
}

// GetReviews mocks base method.
func (m *MockCommentRepo) GetReviews(arg0 context.Context) ([]entity.Review, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetReviews", arg0)
	ret0, _ := ret[0].([]entity.Review)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetReviews indicates an expected call of GetReviews.
func (mr *MockCommentRepoMockRecorder) GetReviews(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetReviews", reflect.TypeOf((*MockCommentRepo)(nil).GetReviews), arg0)
}

// InsertReview mocks base method.
func (m *MockCommentRepo) InsertReview(arg0 context.Context, arg1 entity.Review) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InsertReview", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// InsertReview indicates an expected call of InsertReview.
func (mr *MockCommentRepoMockRecorder) InsertReview(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsertReview", reflect.TypeOf((*MockCommentRepo)(nil).InsertReview), arg0, arg1)
}

// MockCommentWebAPI is a mock of CommentWebAPI interface.
type MockCommentWebAPI struct {
	ctrl     *gomock.Controller
	recorder *MockCommentWebAPIMockRecorder
}

// MockCommentWebAPIMockRecorder is the mock recorder for MockCommentWebAPI.
type MockCommentWebAPIMockRecorder struct {
	mock *MockCommentWebAPI
}

// NewMockCommentWebAPI creates a new mock instance.
func NewMockCommentWebAPI(ctrl *gomock.Controller) *MockCommentWebAPI {
	mock := &MockCommentWebAPI{ctrl: ctrl}
	mock.recorder = &MockCommentWebAPIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCommentWebAPI) EXPECT() *MockCommentWebAPIMockRecorder {
	return m.recorder
}
