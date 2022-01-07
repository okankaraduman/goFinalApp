// Package entity defines main entities for business logic (services), data base mapping and
// HTTP response objects if suitable. Each logic group entities in own file.
package entity

import "time"

//Models!

// User holds a user's data.
type User struct {
	UserID    int       `json:"user_id"`
	Username  string    `json:"user_name"`
	Email     string    `json:"email"`
	FullName  string    `json:"full_name"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
