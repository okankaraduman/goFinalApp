package repo

import (
	"github.com/okankaraduman/goFinalApp/pkg/postgres"
)

const _defaultEntityCap = 64

// TranslationRepo -.
type UserRepo struct {
	*postgres.Postgres
}

// New -.
func New(pg *postgres.Postgres) *UserRepo {
	return &UserRepo{pg}
}
