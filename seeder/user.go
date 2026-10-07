package seeder

import (
	"github.com/bootcamp-assignment/domain"
	"golang.org/x/crypto/bcrypt"
)

var hashPassword, _ = bcrypt.GenerateFromPassword([]byte("Password123"), 10)

var User = []domain.User{
	{
		ID:       1,
		Name:     "Lewi",
		Username: "lewimb",
		Password: string(hashPassword),
	},
}
