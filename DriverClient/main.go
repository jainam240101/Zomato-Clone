package main

import (
	"context"

	"github.com/hashicorp/go-hclog"
	protos "github.com/jainam240101/zomato-clone/Protos/DriverProtos"
	"google.golang.org/grpc"
)

func main() {
	opts := grpc.WithInsecure()
	log := hclog.Default()
	cc, err := grpc.Dial("localhost:8083", opts)
	if err != nil {
		log.Error("Some error occured %v", err)
	}
	defer cc.Close()

	client := protos.NewDriverServiceClient(cc)

	// client.AddDriverLocation(context.Background(), &protos.DriverDetails{
	// 	DriverId:  "1",
	// 	Latitude:  -33.44091,
	// 	Longitude: -70.6301,
	// })
	// client.AddDriverLocation(context.Background(), &protos.DriverDetails{
	// 	DriverId:  "2",
	// 	Latitude:  -33.44005,
	// 	Longitude: -70.63279,
	// })
	// client.AddDriverLocation(context.Background(), &protos.DriverDetails{
	// 	DriverId:  "3",
	// 	Latitude:  -33.44338,
	// 	Longitude: -70.63335,
	// })
	// client.AddDriverLocation(context.Background(), &protos.DriverDetails{
	// 	DriverId:  "4",
	// 	Latitude:  -33.44186,
	// 	Longitude: -70.62653,
	// })

	// data, _ := client.SearchForDrivers(context.Background(), &protos.DriverSearch{
	// 	Latitude:  -33.44262,
	// 	Longitude: -70.63054,
	// 	Limit:     5,
	// 	OrderId:   "62f7732c66188fd414fd8403",
	// })

	data,_:=client.UpdateOrder(context.Background(),&protos.OrderDetails{
		DriverId: "1",
		OrderId: "62f7732c66188fd414fd8403",
	})
	log.Info("Data ", data)
}
