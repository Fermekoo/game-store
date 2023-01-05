package gapi

import (
	"context"

	"github.com/Fermekoo/game-store/db"
	"github.com/Fermekoo/game-store/pb"
	"github.com/Fermekoo/game-store/repositories/order"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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

	response := &pb.OrderResponse{Order: pb_order}
	return response, nil
}
