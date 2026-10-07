package auth

import (
	"net/http"

	"github.com/bootcamp-assignment/request"
	"github.com/gin-gonic/gin"
)

func (h *AuthHandler) LoginHandler(c *gin.Context) {
	var user request.LoginRequest
	if err := c.ShouldBindBodyWithJSON(&user); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if user.Username == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Please input your username!"})
		return
	}

	if user.Password == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Please input your password!"})
		return
	}

	res, err := h.AuthService.Login(user.Username, user.Password)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"accessToken": res.AccessToken})
	c.Abort()
}
