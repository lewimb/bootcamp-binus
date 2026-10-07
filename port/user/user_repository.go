package port

import "github.com/bootcamp-assignment/domain"

type UserRepository interface {
	FindByUsername(username string) (*domain.User, error)
}
