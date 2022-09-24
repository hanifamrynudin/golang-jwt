package controllers

import (
	"golang-jwt/models"
	"golang-jwt/services"
	"golang-jwt/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController interface {
	LogIn(c *gin.Context)
	Register(c *gin.Context)
	Profile(c *gin.Context)
}

type userController struct {
	userService services.UserService
	authService services.AuthService
}

func (con *userController) LogIn(c *gin.Context) {
	var input models.LoginInput

	contentType := utils.GetContentType(c)
	var err error = nil
	if contentType == "application/json" {
		err = c.ShouldBindJSON(&input)
	} else {
		err = c.ShouldBind(&input)
	}

	if err != nil {
		c.JSON(http.StatusBadRequest, utils.NewResponse(
			http.StatusBadRequest,
			nil,
			err.Error(),
		))
		return
	}

	user, err := con.userService.AttemptLogin(input)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.NewResponse(
			http.StatusBadRequest,
			nil,
			err.Error(),
		))
		return
	}

	token := con.authService.GenerateToken(user.ID)

	c.JSON(http.StatusOK, utils.NewResponse(
		http.StatusOK,
		token,
		"Success login",
	))
}

func (con *userController) Register(c *gin.Context) {
	var input models.UserInput

	contentType := utils.GetContentType(c)
	var err error = nil
	if contentType == "application/json" {
		err = c.ShouldBindJSON(&input)
	} else {
		err = c.ShouldBind(&input)
	}

	if err != nil {
		c.JSON(http.StatusBadRequest, utils.NewResponse(
			http.StatusBadRequest,
			nil,
			err.Error(),
		))
		return
	}

	user, err := con.userService.Register(input)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.NewResponse(
			http.StatusBadRequest,
			nil,
			err.Error(),
		))
		return
	}

	c.JSON(http.StatusOK, utils.NewResponse(
		http.StatusOK,
		user,
		"Success register",
	))
}

func (con *userController) Profile(c *gin.Context) {
	userID, ok := c.MustGet("user_id").(float64)
	if !ok {
		c.JSON(http.StatusBadRequest, utils.NewResponse(
			http.StatusBadRequest,
			nil,
			"Cannot parsing user",
		))
		return
	}

	user, err := con.userService.GetUserById(uint(userID))
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.NewResponse(
			http.StatusBadRequest,
			nil,
			"Cannot get user",
		))
		return
	}

	c.JSON(http.StatusOK, utils.NewResponse(
		http.StatusOK,
		user,
		"Success get user",
	))
}

func ProvideUserController(userService services.UserService, authService services.AuthService) *userController {
	return &userController{userService, authService}
}