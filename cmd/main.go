package main

import (
	"context"
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
	server := api.NewServer(service, config)
	if err != nil {
		log.Fatal("cannot create server", err)
	}

	server.Start(config.ServerAddress, context.Background())
}
