package middleware

import (
	"golang-jwt/services"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func Authentication(authService services.AuthService) gin.HandlerFunc {
	return func(c *gin.Context) {
		verifyToken, err := authService.VerifyToken(c)
		_ = verifyToken

		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"status_code": http.StatusUnauthorized,
				"message":     err.Error(),
			})
			return
		}

		c.Set("user_id", verifyToken.(jwt.MapClaims)["user_id"])
		c.Next()
	}
}
