package port

import (
	"github.com/bootcamp-assignment/domain"
	"github.com/bootcamp-assignment/response"
)

type AuthService interface {
	Login(username, password string) (*response.LoginResponse, error)
	Me(username string) (*domain.User, error)
}
