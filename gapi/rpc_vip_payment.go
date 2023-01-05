package gapi

import (
	"context"

	"github.com/Fermekoo/game-store/db"
	"github.com/Fermekoo/game-store/pb"
	"github.com/Fermekoo/game-store/pkg"
	"github.com/Fermekoo/game-store/repositories/order"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (server *Server) Order(ctx context.Context, request *pb.OrderCallRequest) (*pb.OrderResponse, error) {

	game_service, err := server.service.DetailService(request.GetService())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "error %s", err)
	}

	price := game_service.Data.Price.Basic
	fee := server.config.ServiceFee
	total_price := price + fee
	order_repo := order.NewOrder(db.Connect())
	order_payload := &order.Order{
		ServiceCode: request.GetService(),
		AccountId:   request.GetAccountId(),
		AccountZone: request.GetAccountZone(),
		TotalPrice:  total_price,
		Fee:         fee,
		Price:       price,
		Status:      order.Pending,
	}

	order, err := order_repo.Create(order_payload)

	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed create order %s", err)
	}
	pb_order := &pb.Order{
		OrderId:     uint32(order.ID),
		AccountId:   order.AccountId,
		AccountZone: order.AccountZone,
		ServiceCode: order.ServiceCode,
		TotalPrice:  uint32(order.TotalPrice),
		Price:       uint32(order.Price),
		Fee:         uint32(order.Fee),
		Status:      order.Status,
		CreatedAt:   timestamppb.New(order.CreatedAt),
		UpdatedAt:   timestamppb.New(order.UpdatedAt),
	}

	response := &pb.OrderResponse{
		Result:  true,
		Message: "success create order",
		Data:    pb_order,
	}
	return response, nil
}

func (server *Server) Profile(ctx context.Context, _ *emptypb.Empty) (*pb.ProfileResponse, error) {
	profile, err := server.service.Profile()
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get profile %s", err)
	}

	pb_profile := &pb.Profile{
		Fullname:   profile.Data.Fullname,
		Username:   profile.Data.Username,
		Balance:    profile.Data.Balance,
		Point:      profile.Data.Balance,
		Level:      profile.Data.Level,
		Registered: profile.Data.Registered,
	}

	response := &pb.ProfileResponse{
		Result:  profile.Result,
		Message: profile.Message,
		Data:    pb_profile,
	}

	return response, nil
}

func (server *Server) Service(ctx context.Context, request *pb.ServiceRequest) (*pb.ServiceResponse, error) {

	var pb_list []*pb.Service
	filter := pkg.FilterListService{
		FilterType:  "game",
		FilterValue: request.GetGame(),
	}

	list, err := server.service.ListService(filter)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get services %s", err)
	}

	for _, service := range list.Data {

		pb_service_price := &pb.Price{
			Basic:   uint32(service.Price.Basic),
			Premium: uint32(service.Price.Premium),
			Special: uint32(service.Price.Special),
		}

		pb_service := &pb.Service{
			Code:         service.Code,
			Game:         service.Game,
			Name:         service.Name,
			Server:       service.Server,
			Status:       service.Status,
			ServicePrice: pb_service_price,
		}

		pb_list = append(pb_list, pb_service)
	}

	response := &pb.ServiceResponse{
		Result:  list.Result,
		Message: list.Message,
		Data:    pb_list,
	}

	return response, nil
}

func (server *Server) Game(ctx context.Context, _ *emptypb.Empty) (*pb.GameResponse, error) {
	games, err := server.service.Game()
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get game %s", err)
	}

	var pb_games []*pb.Game
	for _, game := range games {
		game := &pb.Game{
			Name: game.Name,
		}

		pb_games = append(pb_games, game)
	}

	response := &pb.GameResponse{
		Result:  true,
		Message: "game list",
		Data:    pb_games,
	}

	return response, nil
}
