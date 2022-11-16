package main

import (
	"log"

	"github.com/Fermekoo/game-store/api"
	"github.com/Fermekoo/game-store/pkg"
	"github.com/Fermekoo/game-store/utils"
)

func main() {
	config, err := utils.LoadConfig("./")
	if err != nil {
		log.Fatal("cannot load config", err)
	}

	service := pkg.NewVIPPayment(config)
	server, err := api.NewServer(service)
	if err != nil {
		log.Fatal("cannot create server", err)
	}

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("cannot start server", err)
	}
}
