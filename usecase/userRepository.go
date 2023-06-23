package usecase

import (
    "github.com/observation-system/observation-auth/domain"
)

type UserRepository interface {
	FindById(id int) (domain.User, error)
	FindByEmail(email string) (domain.User, error)
	Store(domain.User) (domain.User, error)
	DeleteByUserKey(domain.User) error
}
