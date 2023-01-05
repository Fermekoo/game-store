package api

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Fermekoo/game-store/pkg"
	"github.com/Fermekoo/game-store/utils"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Server struct {
	router  *gin.Engine
	service pkg.ApiGameInterface
	config  utils.Config
}

func NewServer(service pkg.ApiGameInterface, config utils.Config) *Server {
	server := &Server{
		service: service,
		config:  config,
	}
	server.SetupRouter()
	return server
}

func (server *Server) SetupRouter() {
	router := gin.Default()
	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "game store api",
		})
	})

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "PUT", "PATCH", "DELETE", "POST", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Authorization", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	router.GET("/profile", server.profile)
	router.POST("/order", server.order)
	router.GET("/game-service", server.services)
	router.GET("/game-service/:code", server.detailService)
	router.GET("/game", server.games)
	server.router = router
}

func (server *Server) Start(address string, ctx context.Context) {
	srv := &http.Server{
		Addr:    address,
		Handler: server.router,
	}

	server_err := make(chan error, 1)
	go func() {
		server_err <- srv.ListenAndServe()
	}()

	shutdown_channel := make(chan os.Signal, 1)
	signal.Notify(shutdown_channel, syscall.SIGINT, syscall.SIGTERM)

	select {
	case sig := <-shutdown_channel:
		log.Println("signal", sig)
		const timeout = 10 * time.Second
		ctx, cancel := context.WithTimeout(ctx, timeout)
		defer cancel()

		if err := srv.Shutdown(ctx); err != nil {
			srv.Close()
		}
	case err := <-server_err:
		if err != nil {
			log.Fatalf("server: %v", err)
		}

	}
}

func errorResponse(err error) gin.H {
	return gin.H{
		"error": err.Error(),
	}
}

func succesResponse(status bool, message string, data interface{}) interface{} {
	return &pkg.GeneralResponse{
		Result:  status,
		Message: message,
		Data:    data,
	}
}
