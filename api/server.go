package api

import (
	"net/http"

	"github.com/Fermekoo/game-store/pkg"
	"github.com/gin-gonic/gin"
)

type Server struct {
	router  *gin.Engine
	service pkg.ApiGameInterface
}

func NewServer(service pkg.ApiGameInterface) (*Server, error) {
	server := &Server{
		service: service,
	}
	server.SetupRouter()
	return server, nil
}

func (server *Server) SetupRouter() {
	router := gin.Default()
	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "game store api",
		})
	})

	router.GET("/profile", server.profile)
	router.POST("/order", server.order)
	router.GET("/game-service", server.services)
	router.GET("/game", server.games)
	server.router = router
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{
		"error": err.Error(),
	}
}
