package handlers

import (
	"context"

	"github.com/hashicorp/go-hclog"
	orderProtos "github.com/jainam240101/zomato-clone/Protos/OrderProtos"
	protos "github.com/jainam240101/zomato-clone/Protos/RestaurantProtos"
)

type Server struct {
	Log hclog.Logger
	protos.UnimplementedRestaurantServiceServer
}

func NewServer(log hclog.Logger) *Server {
	return &Server{Log: log}
}

func (s *Server) AcceptOrder(ctx context.Context, request *orderProtos.OrderResponse) (*orderProtos.OrderResponse, error) {
	s.Log.Info("Accepting Orders from Restaurant")
	orderDetails := request
	orderDetails.OrderStatus = "Accepted"
	return orderDetails, nil
}
