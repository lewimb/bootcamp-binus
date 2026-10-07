package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *AuthHandler) MeHandler(c *gin.Context) {
	username, _ := c.Get("username")

	user, err := h.AuthService.Me(username.(string))

	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, user)
	c.Abort()
}
