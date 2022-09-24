package models

import (
	"time"

	"github.com/gin-gonic/gin"
)

type User struct {
	ID        uint       `gorm:"primaryKey" json:"id"`
	FullName  string     `gorm:"not null" json:"full_name"`
	Email     string     `gorm:"not null;uniqueIndex" json:"email"`
	Password  string     `gorm:"not null" json:"password"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `sql:"index" json:"deleted_at"`
}

type UserInput struct {
	Fullname string `json:"full_name" form:"full_name" binding:"required"`
	Email    string `json:"email" form:"email" binding:"required"`
	Password string `json:"password" form:"password" binding:"required"`
}

type LoginInput struct {
	Email    string `json:"email" form:"email" binding:"required"`
	Password string `json:"password" form:"password" binding:"required"`
}

func GetContentType(c *gin.Context) string {
	return c.Request.Header.Get("Content-Type")
}
