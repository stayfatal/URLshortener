package middleware

import (
	"net/http"
	"url/internal/auth"

	"github.com/gin-gonic/gin"
)

func Authentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("Authorization")

		_, err := auth.ValidateToken(token)
		if err != nil {
			c.String(http.StatusForbidden, err.Error())
			return
		}
		c.Next()
	}
}
