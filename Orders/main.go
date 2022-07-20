package main

import (
	"fmt"
	"net"

	"github.com/hashicorp/go-hclog"
	"github.com/jainam240101/zomato-clone/Orders/db"
	"github.com/jainam240101/zomato-clone/Orders/handlers"
	protos "github.com/jainam240101/zomato-clone/Protos/OrderProtos"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
)

func main() {
	log := hclog.Default()
	godotenv.Load("../.env")
	db.ConnectDb()
	log.Info("Database Connected")
	address := "0.0.0.0:8080"
	lis, err := net.Listen("tcp", address)
	if err != nil {
		log.Error("Error %v", err)
		return
	}
	fmt.Println("Server is listening on ", address)
	s := grpc.NewServer()
	server := &handlers.Server{
		Log: log,
		DB:  db.OrderDB,
	}
	protos.RegisterOrderServiceServer(s, server)
	s.Serve(lis)
}
