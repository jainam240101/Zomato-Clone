package handlers

import (
	"context"

	"github.com/hashicorp/go-hclog"
	redisModule "github.com/jainam240101/zomato-clone/Driver/Redis"
	protos "github.com/jainam240101/zomato-clone/Protos/DriverProtos"
)

type Server struct {
	Log   hclog.Logger
	redis *redisModule.RedisClient
	protos.UnimplementedDriverServiceServer
}

func NewServer(log hclog.Logger) *Server {
	redis := redisModule.ConnectRedis()
	return &Server{
		redis: redis,
		Log:   log,
	}
}

func (s *Server) SearchForDrivers(ctx context.Context, request *protos.DriverSearch) (*protos.SearchResponse, error) {
	s.Log.Info("Search Function")
	data, _ := s.redis.SearchDrivers(int(request.Limit), float64(request.Latitude), float64(request.Longitude), 10, request.OrderId)
	return &protos.SearchResponse{
		DriverLocations: data,
	}, nil
}

func (s *Server) AddDriverLocation(ctx context.Context, request *protos.DriverDetails) (*protos.DriverResponse, error) {
	s.Log.Info("Tracking")
	s.redis.AddDriverLocation(float64(request.Longitude), float64(request.Latitude), request.DriverId)
	return &protos.DriverResponse{}, nil
}

func (s *Server) UpdateDriverLocation(ctx context.Context, request *protos.DriverDetails) (*protos.DriverResponse, error) {
	s.Log.Info("Tracking")
	s.redis.AddDriverLocation(float64(request.Longitude), float64(request.Latitude), request.DriverId)
	return &protos.DriverResponse{}, nil
}

func (s *Server) UpdateOrder(ctx context.Context, request *protos.OrderDetails) (*protos.OrderResponse, error) {
	s.Log.Info("Updating an order")
	_, err := s.redis.UpdateOrder(request.DriverId, request.OrderId)
	if err != nil {
		return nil, err
	}
	return &protos.OrderResponse{Message: "Done"}, nil
}
