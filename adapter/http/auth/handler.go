package auth

import (
	"github.com/bootcamp-assignment/adapter/middleware"
	port "github.com/bootcamp-assignment/port/auth"
	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	AuthService port.AuthService
}

func NewAuthHandler(authService port.AuthService) *AuthHandler {
	return &AuthHandler{AuthService: authService}
}

func (a *AuthHandler) RegisterRoutes(router *gin.Engine) {
	auth := router.Group("/api/auth")
	{
		auth.POST("/login", a.LoginHandler)
		auth.GET("/me", middleware.TokenAuth, a.MeHandler)
	}
}
