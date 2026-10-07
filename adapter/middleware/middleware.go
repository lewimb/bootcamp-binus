package middleware

import (
	"net/http"
	"strings"

	adapter "github.com/bootcamp-assignment/adapter/jwt"
	"github.com/bootcamp-assignment/config"
	"github.com/gin-gonic/gin"
)

func TokenAuth(c *gin.Context) {
	headers := c.Request.Header
	token := headers.Get("Authorization")
	splittedToken := strings.Split(token, " ")

	if len(splittedToken) != 2 {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	if splittedToken[0] != "Bearer" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	jwtPayload, err := adapter.NewJWTAdapter(config.SetupConfig().JwtSecret).ParseToken(splittedToken[1])

	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	c.Set("username", jwtPayload.Username)
	c.Next()
}
