package handler

import (
	"atro/internal/helper"
	"atro/internal/model"
	"atro/internal/model/request"
	"atro/internal/model/response"
	"atro/internal/repository"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

//OrderHandler --> Handler for Order Entity
type OrderHandler interface {
	OrderProduct(*gin.Context)
	GetOrderProduct(*gin.Context)
	GetAllOrderProduct(*gin.Context)
	UpdateOrderProduct(*gin.Context)
}

type orderHandler struct {
	repo        repository.OrderRepository
	repoProduct repository.ProductRepository
}

//NewOrderHandler --> return new Order Handler
func NewOrderHandler() OrderHandler {
	return &orderHandler{
		repo:        repository.NewOrderRepository(),
		repoProduct: repository.NewProductRepository(),
	}
}

func (h *orderHandler) OrderProduct(ctx *gin.Context) { // TODO: transaction?
	var orderForm request.OrderRequest
	if err := ctx.ShouldBindJSON(&orderForm); err != nil {
		ctx.JSON(http.StatusBadRequest, helper.BuildResponse(-1, "invalid input", err.Error()))
		return
	}

	id, _ := ctx.Get("userID")
	if id == nil {
		ctx.JSON(http.StatusInternalServerError, helper.BuildResponse(-1, "get the lost user id", ""))
		return
	}

	order, err := h.OrderRequestToOrder(&orderForm, fmt.Sprint(id)) // parse asset id to int syntax
	if err != nil {
		ctx.JSON(http.StatusBadRequest, helper.BuildResponse(-1, "Cannot create order", err.Error()))
		return
	}

	order.OrderId = uuid.NewString()
	rsOrder, err := h.repo.OrderProduct(order)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, helper.BuildResponse(-1, "error when add product (id is wrong)", err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, helper.BuildResponse(1, "create order successfully!", rsOrder))

}

func (h *orderHandler) GetOrderProduct(ctx *gin.Context) {

	id := ctx.Param("id")
	order, err := h.repo.GetOrder(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, helper.BuildResponse(-1, "error when find product", err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, helper.BuildResponse(1, "get product successfully!", order))

}

func (h *orderHandler) GetAllOrderProduct(ctx *gin.Context) {

	sortBy := ctx.Query("sort-by")
	if sortBy == "" {
		sortBy = "order_id.asc" // sortBy is expected to look like field.orderdirection i. e. id.asc
	}
	sortQuery, err := helper.ValidateAndReturnSortQuery(model.Order{}, sortBy)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, helper.BuildResponse(-1, "invalid param sort", err.Error()))
		return
	}

	// tao query limit
	strLimit := ctx.Query("limit")
	fmt.Println("param limit", strLimit)
	limit := -1 // with a value as -1 for gorms Limit method, we'll get a request without limit as default
	if strLimit != "" {
		limit, err = strconv.Atoi(strLimit)
		if err != nil || limit < -1 {
			ctx.JSON(http.StatusBadRequest, helper.BuildResponse(-1, "limit query parameter is no valid number", err.Error()))
			return
		}
	}

	strOffset := ctx.Query("offset")
	offset := -1
	if strOffset != "" {
		offset, err = strconv.Atoi(strOffset)
		if err != nil || offset < -1 {
			ctx.JSON(http.StatusBadRequest, helper.BuildResponse(-1, "offset query parameter is no valid number", err.Error()))
			return
		}
	}

	filter := ctx.Query("filter")
	filterMap := map[string]interface{}{}
	if filter != "" {
		filterMap, err = helper.ValidateAndReturnFilterMap(model.Order{}, filter)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, helper.BuildResponse(-1, "invalid filter param ", err.Error()))
			return
		}
	}

	rsOrders, err := h.repo.GetAllOrderOptions(filterMap, limit, offset, sortQuery)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, helper.BuildResponse(-1, "not found !", err.Error()))
		return
	}

	res := response.OrderResponse{
		Orders:       rsOrders,
		OrdersLength: len(rsOrders),
	}
	ctx.JSON(http.StatusOK, helper.BuildResponse(1, "get list products successfully!", res))
}

func (h *orderHandler) UpdateOrderProduct(ctx *gin.Context) {

}

func (h *orderHandler) OrderRequestToOrder(orderForm *request.OrderRequest, userId string) (model.Order, error) {

	total := 0.0
	for _, element := range orderForm.ProductOrders {
		if productDetail, err := h.repoProduct.GetProduct(element.ProductId); err == nil {
			fmt.Println("product detail ", productDetail)
			total = total + productDetail.ProductPrice*float64(element.Quantity)

		} else {
			return model.Order{}, err
		}

	}

	productOrdersInfo, err := json.Marshal(orderForm)
	if err != nil {
		return model.Order{}, err
	}

	order := model.Order{
		UserId:         userId,
		OrderDetail:    string(productOrdersInfo),
		OrderPrice:     float32(total),
		OrderCreatedAt: time.Now(),
		OrderStatus:    orderForm.TypeOrder,
	}
	return order, nil
}
