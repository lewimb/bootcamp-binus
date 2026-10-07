package usecase

import (
	"fmt"

	adapter "github.com/bootcamp-assignment/adapter/jwt"
	"github.com/bootcamp-assignment/config"
	"github.com/bootcamp-assignment/domain"
	port "github.com/bootcamp-assignment/port/user"
	"github.com/bootcamp-assignment/response"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	userRepository port.UserRepository
}

func NewAuthService(userRepository port.UserRepository) *AuthService {
	return &AuthService{userRepository: userRepository}
}

func (auth *AuthService) Login(username, password string) (*response.LoginResponse, error) {
	user, err := auth.userRepository.FindByUsername(username)
	if err != nil {
		return nil, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return nil, fmt.Errorf("Unauthorized")
	}

	jwtToken, err := adapter.NewJWTAdapter(config.SetupConfig().JwtSecret).Generate(adapter.Payload{Username: user.Username, ID: user.ID, Name: user.Name})

	if err != nil {
		return nil, err
	}
	res := response.LoginResponse{
		AccessToken: jwtToken,
	}

	return &res, nil
}

func (auth *AuthService) Me(username string) (*domain.User, error) {
	user, err := auth.userRepository.FindByUsername(username)

	if err != nil {
		return nil, err
	}

	return &domain.User{ID: user.ID, Username: user.Username, Name: user.Name}, nil
}
