package api

import (
	"net/http"

	"github.com/Fermekoo/game-store/db"
	"github.com/Fermekoo/game-store/pkg"
	"github.com/Fermekoo/game-store/repositories/order"
	"github.com/gin-gonic/gin"
)

func (server *Server) profile(ctx *gin.Context) {
	profile, err := server.service.Profile()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, profile)
}

func (server *Server) order(ctx *gin.Context) {
	var request pkg.OrderCallRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	game_service, err := server.service.DetailService(request.ServiceCode)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	price := game_service.Data.Price.Basic
	fee := server.config.ServiceFee
	total_price := price + fee
	order_repo := order.NewOrder(db.Connect())
	order_payload := &order.Order{
		ServiceCode: request.ServiceCode,
		AccountId:   request.AccountID,
		AccountZone: request.ServiceCode,
		TotalPrice:  total_price,
		Fee:         fee,
		Price:       price,
		Status:      order.Pending,
	}

	order, err := order_repo.Create(order_payload)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, succesResponse(true, "success create order", order))
}

func (server *Server) services(ctx *gin.Context) {
	var request pkg.FilterRequestListService
	if err := ctx.ShouldBindQuery(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	filter := pkg.FilterListService{
		FilterType:  "game",
		FilterValue: request.Game,
	}
	list, err := server.service.ListService(filter)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, list)
}

func (server *Server) detailService(ctx *gin.Context) {
	service_code := ctx.Param("code")

	detail, err := server.service.DetailService(service_code)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
	}

	ctx.JSON(http.StatusOK, detail)
}

func (server *Server) games(ctx *gin.Context) {
	list, err := server.service.Game()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, list)
}
