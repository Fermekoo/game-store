package api

import (
	"net/http"

	"github.com/Fermekoo/game-store/pkg"
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

	order, err := server.service.Order(pkg.OrderCall(request))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, order)
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

func (server *Server) games(ctx *gin.Context) {
	list, err := server.service.Game()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, list)
}
