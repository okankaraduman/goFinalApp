// Package entity defines main entities for business logic (services), data base mapping and
// HTTP response objects if suitable. Each logic group entities in own file.
package entity

//Models!

// User holds a user's data.
type User struct {
	Username string
	Email    string
	FullName string
	Password string
}
