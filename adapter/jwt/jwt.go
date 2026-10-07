package adapter

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JWTAdapter struct {
	JWTSecret []byte
}

type Claims struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Name     string `json:"name"`
	jwt.RegisteredClaims
}

type Payload struct {
	ID       int
	Username string
	Name     string
}

func NewJWTAdapter(secret []byte) *JWTAdapter {
	return &JWTAdapter{JWTSecret: secret}
}

func (j *JWTAdapter) Generate(obj Payload) (string, error) {
	c := Claims{
		ID:        obj.ID,
		Username:  obj.Username,
		Name:      obj.Name,
		IssuedAt:  jwt.NewNumericDate(time.Now()),
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour)),
	}

	return jwt.NewWithClaims(jwt.SigningMethodHS256, c).SignedString(j.JWTSecret)
}

func (j *JWTAdapter) ParseToken(jwtToken string) (*Claims, error) {
	claims := Claims{}
	_, err := jwt.ParseWithClaims(jwtToken, &claims, func(token *jwt.Token) (any, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return j.JWTSecret, nil
	})

	if err != nil {
		return nil, err
	}

	return &claims, nil
}
