package controllers

import (
	"golang-jwt/models"
	"golang-jwt/services"
	"golang-jwt/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type OrdersController interface {
	GetAllOrders(ctx *gin.Context)
	GetOrderByID(ctx *gin.Context)
	CreateOrder(ctx *gin.Context)
	UpdateOrder(ctx *gin.Context)
	DeleteOrder(ctx *gin.Context)
}

type ordersController struct {
	orderService services.OrdersService
}

func (c *ordersController) GetAllOrders(ctx *gin.Context) {
	var orders []models.Order

	orders, err := c.orderService.GetOrders()
	if err != nil {
		ctx.JSON(http.StatusNotFound, utils.NewResponse(
			http.StatusNotFound,
			orders,
			err.Error(),
		))
		return
	}

	ctx.JSON(http.StatusOK, utils.NewResponse(
		http.StatusOK,
		orders,
		"Success get orders",
	))
}

func (c *ordersController) GetOrderByID(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.NewResponse(
			http.StatusBadRequest,
			nil,
			err.Error(),
		))
		return
	}

	order, err := c.orderService.GetOrder(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, utils.NewResponse(
			http.StatusNotFound,
			nil,
			err.Error(),
		))
		return
	}

	ctx.JSON(http.StatusOK, utils.NewResponse(
		http.StatusOK,
		order,
		"Success get order",
	))
}

func (c *ordersController) CreateOrder(ctx *gin.Context) {
	var input models.OrderInput

	err := ctx.ShouldBindJSON(&input)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.NewResponse(
			http.StatusBadRequest,
			nil,
			err.Error(),
		))
		return
	}

	order, err := c.orderService.CreateOrder(input)
	if err != nil {
		ctx.JSON(http.StatusNotFound, utils.NewResponse(
			http.StatusNotFound,
			nil,
			err.Error(),
		))
		return
	}

	ctx.JSON(http.StatusOK, utils.NewResponse(
		http.StatusOK,
		order,
		"Success create order",
	))
}

func (c *ordersController) UpdateOrder(ctx *gin.Context) {
	var input models.OrderInput

	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.NewResponse(
			http.StatusBadRequest,
			nil,
			err.Error(),
		))
		return
	}

	err = ctx.ShouldBindJSON(&input)
	if err != nil {
		ctx.JSON(http.StatusNotFound, utils.NewResponse(
			http.StatusNotFound,
			nil,
			err.Error(),
		))
		return
	}

	order, err := c.orderService.UpdateOrder(id, input)
	if err != nil {
		ctx.JSON(http.StatusNotFound, utils.NewResponse(
			http.StatusNotFound,
			nil,
			err.Error(),
		))
		return
	}

	ctx.JSON(http.StatusOK, utils.NewResponse(
		http.StatusOK,
		order,
		"Success update order",
	))
}

func (c *ordersController) DeleteOrder(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.NewResponse(
			http.StatusBadRequest,
			nil,
			err.Error(),
		))
		return
	}

	err = c.orderService.DeleteOrder(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, utils.NewResponse(
			http.StatusNotFound,
			false,
			err.Error(),
		))
		return
	}

	ctx.JSON(http.StatusOK, utils.NewResponse(
		http.StatusOK,
		true,
		"Success delete order",
	))
}

func ProvideOrdersController(orderService services.OrdersService) *ordersController {
	return &ordersController{orderService}
}