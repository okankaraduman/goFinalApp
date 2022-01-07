// Package entity defines main entities for business logic (services), data base mapping and
// HTTP response objects if suitable. Each logic group entities in own file.
package entity

//Models!

type CreateReviewRequest struct {
	UserId    int64  `json:"userId"`
	ContentId string `json:"contentId"`
	Rate      int    `json:"rate"`
	Comment   string `json:"comment"`
	UserName  string `json:"userName"`
}

type Review struct {
	Id               string `json:"id,omitempty"`
	UserId           int64  `json:"userId,omitempty"`
	ContentId        string `json:"contentId,omitempty"`
	ReviewStatus     string `json:"reviewStatus,omitempty"`
	Rate             int    `json:"rate,omitempty"`
	Comment          string `json:"comment,omitempty"`
	UserName         string `json:"userName,omitempty"`
	CreatedDate      int64  `json:"createdDate,omitempty"`
	LastModifiedDate int64  `json:"lastModifiedDate,omitempty"`
}

type ReviewDTO struct {
	Id        string `json:"id,omitempty"`
	UserId    int64  `json:"userId,omitempty"`
	ContentId string `json:"contentId,omitempty"`
	Rate      int    `json:"rate,omitempty"`
	Comment   string `json:"comment,omitempty"`
	UserName  string `json:"userName,omitempty"`
}
