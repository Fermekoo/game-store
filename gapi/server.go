package gapi

import (
	"github.com/Fermekoo/game-store/pb"
	"github.com/Fermekoo/game-store/pkg"
	"github.com/Fermekoo/game-store/utils"
)

type Server struct {
	pb.UnimplementedGameStoreServer
	service pkg.ApiGameInterface
	config  utils.Config
}

func NewServer(service pkg.ApiGameInterface, config utils.Config) *Server {
	server := &Server{
		service: service,
		config:  config,
	}

	return server
}
