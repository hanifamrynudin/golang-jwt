package main

import (
	"golang-jwt/config"
	"golang-jwt/controllers"
	"golang-jwt/database"
	"golang-jwt/repositories"
	"golang-jwt/routes"
	"golang-jwt/services"

	"github.com/gin-gonic/gin"
)

func main() {

	cfg := config.LoadConfig()
	db := database.InitializeDatabase(cfg.Database.Username, cfg.Database.Password, cfg.Database.Port, cfg.Database.Name, cfg.Database.Host)

	orderRepository := repositories.ProvideOrderRepository(db)
	itemRepository := repositories.ProvideItemRepository(db)
	orderService := services.ProvideOrderService(orderRepository, itemRepository)
	orderController := controllers.ProvideOrdersController(orderService)

	userRepository := repositories.ProvideUserRepository(db)
	userService := services.ProvideUserService(userRepository)
	authService := services.ProvideAuthService()
	userController := controllers.ProvideUserController(userService, authService)

	router := gin.Default()
	routes.OrderRoutes(router, authService, orderController)
	routes.AuthRoutes(router, authService, userController)

	router.Run(cfg.ServerPort)
}
