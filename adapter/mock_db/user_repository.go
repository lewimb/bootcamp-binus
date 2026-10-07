package mockdb

import (
	"fmt"

	"github.com/bootcamp-assignment/domain"
	"github.com/bootcamp-assignment/seeder"
)

type UserRepository struct{}

func NewUserRepository() *UserRepository {
	return &UserRepository{}
}

func (u *UserRepository) FindByUsername(username string) (*domain.User, error) {
	for _, u := range seeder.User {
		if u.Username == username {
			return &u, nil
		}
	}
	return nil, fmt.Errorf("Unauthorized")

}
