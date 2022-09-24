package routes

import (
	"golang-jwt/controllers"
	"golang-jwt/middleware"
	"golang-jwt/services"

	"github.com/gin-gonic/gin"
)

func AuthRoutes(router *gin.Engine, authService services.AuthService, authController controllers.UserController) {
	authRoutes := router.Group("/auth")
	{
		authRoutes.POST("/login", authController.LogIn)
		authRoutes.POST("/register", authController.Register)
		authRoutes.GET("/me", middleware.Authentication(authService), authController.Profile)
	}
}
