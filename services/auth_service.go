package services

import (
	"errors"
	"golang-jwt/utils"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

type AuthService interface {
	VerifyToken(c *gin.Context) (interface{}, error)
	GenerateToken(userID uint) string
}

type authService struct {
	SecretKey string
}

func (s *authService) VerifyToken(c *gin.Context) (interface{}, error) {
	errResponse := errors.New("unauthorized")
	headerToken := c.Request.Header.Get("Authorization")
	bearer := strings.HasPrefix(headerToken, "Bearer")

	if !bearer {
		return nil, errResponse
	}

	if len(strings.Split(headerToken, " ")) < 2 {
		return nil, errResponse
	}

	stringToken := strings.Split(headerToken, " ")[1]

	token, err := jwt.Parse(stringToken, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errResponse
		}

		return []byte(s.SecretKey), nil
	})

	if token == nil {
		return nil, errResponse
	}

	if err != nil {
		v, _ := err.(*jwt.ValidationError)
		if v.Errors == jwt.ValidationErrorExpired {
			return nil, errResponse
		}
	}

	if _, ok := token.Claims.(jwt.MapClaims); !ok && !token.Valid {
		return nil, errResponse
	}

	return token.Claims.(jwt.MapClaims), nil
}

func (s *authService) GenerateToken(userID uint) string {
	ttl := 24 * time.Hour
	claims := jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().UTC().Add(ttl).Unix(),
	}

	parseToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, _ := parseToken.SignedString([]byte(s.SecretKey))

	return signedToken
}

func ProvideAuthService() *authService {
	return &authService{
		SecretKey: utils.GetSecretKey(),
	}
}