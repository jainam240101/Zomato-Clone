package main

import (
	"net"
	"os"

	"github.com/hashicorp/go-hclog"
	protos "github.com/jainam240101/zomato-clone/Protos/RestaurantProtos"
	"github.com/jainam240101/zomato-clone/Restaurant/handlers"
	"google.golang.org/grpc"
)

func main() {
	log := hclog.Default()
	address := "0.0.0.0:8081"
	lis, err := net.Listen("tcp", address)
	if err != nil {
		log.Error("Error in Listening ", err)
		os.Exit(0)
	}
	log.Info("Server Listening on ", address)
	s := grpc.NewServer()
	server := &handlers.Server{
		Log: log,
	}
	protos.RegisterRestaurantServiceServer(s, server)
	s.Serve(lis)
}
