package utils

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) string {
	salt := 8
	passwordByte := []byte(password)
	hash, _ := bcrypt.GenerateFromPassword(passwordByte, salt)

	return string(hash)
}

func ComparePassword(password []byte, hash []byte) bool {
	err := bcrypt.CompareHashAndPassword(hash, password)
	fmt.Println(err, string(password), string(hash))
	return err == nil
}

func GetContentType(c *gin.Context) string {
	return c.Request.Header.Get("Content-Type")
}

func GetSecretKey() string {
	secret := os.Getenv("SECRET_KEY")
	if secret == "" {
		secret = "This is secret"
	}

	return secret
}
