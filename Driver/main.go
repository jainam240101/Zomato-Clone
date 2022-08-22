package main

import (
	"net"

	"github.com/hashicorp/go-hclog"
	handlers "github.com/jainam240101/zomato-clone/Driver/Handlers"
	protos "github.com/jainam240101/zomato-clone/Protos/DriverProtos"
	"google.golang.org/grpc"
)

func main() {
	log := hclog.Default()
	address := "0.0.0.0:8083"
	lis, err := net.Listen("tcp", address)
	if err != nil {
		log.Error("Error %v", err)
		return
	}

	log.Info("Server is listening on ", address)
	s := grpc.NewServer()
	server := handlers.NewServer(log)
	protos.RegisterDriverServiceServer(s, server)
	s.Serve(lis)
}
