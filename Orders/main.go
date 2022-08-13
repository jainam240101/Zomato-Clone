package main

import (
	"fmt"
	"net"

	"github.com/hashicorp/go-hclog"
	"github.com/jainam240101/zomato-clone/Orders/db"
	"github.com/jainam240101/zomato-clone/Orders/handlers"
	orderProtos "github.com/jainam240101/zomato-clone/Protos/OrderProtos"
	restaurantProtos "github.com/jainam240101/zomato-clone/Protos/RestaurantProtos"
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

	restoConn, err := CreateRestaurantConnection(log)
	if err != nil {
		log.Error("Error %v", err)
		return
	}

	fmt.Println("Server is listening on ", address)
	s := grpc.NewServer()
	server := &handlers.Server{
		Log:    log,
		Restro: *restoConn,
		DB:     db.OrderDB,
	}
	orderProtos.RegisterOrderServiceServer(s, server)
	s.Serve(lis)
}

func CreateRestaurantConnection(log hclog.Logger) (*restaurantProtos.RestaurantServiceClient, error) {
	opts := grpc.WithInsecure()
	cc, err := grpc.Dial("localhost:8081", opts)
	if err != nil {
		log.Error("Error ", err)
		return nil, err
	}
	// defer cc.Close()
	client := restaurantProtos.NewRestaurantServiceClient(cc)
	return &client, nil
}
