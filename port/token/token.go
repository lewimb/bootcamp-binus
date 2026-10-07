package port

import adapter "github.com/bootcamp-assignment/adapter/jwt"

type JWTAdapter interface {
	Generate(p adapter.Payload) (string, error)
	ParseToken(jwtToken string) (*adapter.Claims, error)
}
