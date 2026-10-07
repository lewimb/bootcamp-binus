package handler

import (
	"github.com/bootcamp-assignment/adapter/http/auth"
	mockdb "github.com/bootcamp-assignment/adapter/mock_db"
	"github.com/bootcamp-assignment/usecase"
	"github.com/gin-gonic/gin"
)

func Handler(router *gin.Engine) {

	authService := usecase.NewAuthService(mockdb.NewUserRepository())
	auth.NewAuthHandler(authService).RegisterRoutes(router)
}
