package main

import (
	"context"
	"log"
	"net"

	"github.com/Fermekoo/game-store/api"
	"github.com/Fermekoo/game-store/gapi"
	"github.com/Fermekoo/game-store/payment"
	"github.com/Fermekoo/game-store/pb"
	"github.com/Fermekoo/game-store/pkg"
	"github.com/Fermekoo/game-store/utils"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	config, err := utils.LoadConfig("./")
	if err != nil {
		log.Fatal("cannot load config", err)
	}

	service := pkg.NewVIPPayment(config)
	payment := payment.NewMidtrans(config)
	go RunGRPCServer(service, payment, config)
	RunHttpServer(service, payment, config)
}

func RunHttpServer(service pkg.ApiGameInterface, payment payment.Payment, config utils.Config) {
	server := api.NewServer(service, payment, config)
	log.Printf("http server run on port %s", config.HTTPServerAddress)
	server.Start(config.HTTPServerAddress, context.Background())
}

func RunGRPCServer(service pkg.ApiGameInterface, payment payment.Payment, config utils.Config) {
	server := gapi.NewServer(service, payment, config)

	gRPCServer := grpc.NewServer()
	pb.RegisterGameStoreServer(gRPCServer, server)
	reflection.Register(gRPCServer)

	listener, err := net.Listen("tcp", config.GRPCServerAddress)
	if err != nil {
		log.Fatal("cannot create listener")
	}

	log.Printf("start gRPC Server on %s", listener.Addr().String())

	err = gRPCServer.Serve(listener)
	if err != nil {
		log.Fatal("cannot start gRPC server")
	}
}
