package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"math"
	"net/http"

	"github.com/hashicorp/go-hclog"
	"github.com/jainam240101/zomato-clone/Driver/db"
	protos "github.com/jainam240101/zomato-clone/Protos/OrderProtos"

	driverProtos "github.com/jainam240101/zomato-clone/Protos/DriverProtos"
	"google.golang.org/grpc"
)

type Body struct {
	Amount      float64 `json:"amount"`
	UserId      string  `json:"userId"`
	Description string  `json:"description"`
}
type PaymentResponse struct {
	Amount float64 `json:"amount"`
	Id     string  `json:"id"`
	Object string  `json:"object"`
}

func round(num float64) int {
	return int(num + math.Copysign(0.5, num))
}

func toFixed(num float64, precision int) float64 {
	output := math.Pow(10, float64(precision))
	return float64(round(num*output)) / output
}

func main() {
	// Default Code
	db.ConnectDb()
	opts := grpc.WithInsecure()
	log := hclog.Default()
	cc, err := grpc.Dial("localhost:8080", opts)
	if err != nil {
		log.Error("Some error occured ", err)
	}

	defer cc.Close()

	client := protos.NewOrderServiceClient(cc)

	driverCC, err := grpc.Dial("localhost:8083", opts)
	if err != nil {
		log.Error("Some error occured %v", err)
	}
	defer driverCC.Close()

	driverClient := driverProtos.NewDriverServiceClient(cc)

	items := GetItems()

	// Ending Default Code

	orderRequest := &protos.OrderDetails{
		UserId:        "0517cc14-9918-4230-a4f1-3670683e3431",
		RestaurantId:  "630b248769fbbb44c75f271d",
		PayableAmount: 2578.34,
		PaymentMethod: protos.PaymentMethod_CARD,
		Order:         items,
		UserLatitude:  23.0301728, //Users Location (Ours)
		UserLongitude: 72.4859199, //Users Location (Ours)
	}

	// Calling Create Order (GRPC)
	orderResp, err := client.CreateOrder(context.Background(), orderRequest)
	if err != nil {
		fmt.Println("error ", err)
	}
	total := toFixed(float64(orderResp.DeliveryCharge), 2) + toFixed(float64(orderRequest.PayableAmount), 2)
	fmt.Println("Order id ", orderResp.OrderId)
	// Calling Create Order (GRPC)

	// Payment
	posturl := "http://localhost:4000/payment/create"
	body := Body{
		Amount:      total,
		UserId:      orderResp.UserId,
		Description: "Placing Order",
	}

	marshalledData, err := json.Marshal(body)
	if err != nil {
		fmt.Println("error ", err)
	}

	r, err := http.NewRequest("POST", posturl, bytes.NewBuffer(marshalledData))
	if err != nil {
		panic(err)
	}

	r.Header.Add("Content-Type", "application/json")

	httpClient := &http.Client{}
	res, err := httpClient.Do(r)
	if err != nil {
		panic(err)
	}

	post := &PaymentResponse{}
	derr := json.NewDecoder(res.Body).Decode(post)
	if derr != nil {
		panic(derr)
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		panic(res.Status)
	}

	fmt.Println("Stripe Payment ID ", post.Id)

	// Ending Payment

	driverData, _ := driverClient.SearchForDrivers(context.Background(), &driverProtos.DriverSearch{
		Latitude:  -33.44262, //Restaurant Location
		Longitude: -70.63054, //Restaurant Location
		Limit:     5,
		OrderId:   orderResp.OrderId,
	})
	fmt.Println("Driver Details ", driverData)

	// data, _ := driverClient.UpdateOrder(context.Background(), &protos.OrderDetails{
	// 	DriverId: "4",
	// 	OrderId:   orderResp.OrderId,
	// })
	// log.Info("Data ", data)

	// fmt.Println("Driver Details ", driverData)
}

func GetItems() []*protos.Order {
	items := []*protos.Order{}
	items = append(items, CreateStruct("Paneer Butter Masala", 2, 250.13, "abcv"))
	items = append(items, CreateStruct("Dal Makhni", 3, 350.13, "1234xabcv"))
	items = append(items, CreateStruct("Lasagna", 1, 550, "9871xuqhud"))
	items = append(items, CreateStruct("Mohito", 2, 150, "144211xuqhud"))
	return items
}

func CreateStruct(DishName string, Quantity int32, Price float32, DishId string) *protos.Order {
	return &protos.Order{
		DishName: DishName,
		Quantity: Quantity,
		Price:    Price,
		DishId:   DishId,
	}
}
