package main

import (
	"context"
	"fmt"

	"github.com/hashicorp/go-hclog"
	protos "github.com/jainam240101/zomato-clone/Protos/OrderProtos"
	// driverProtos "github.com/jainam240101/zomato-clone/Protos/DriverProtos"
	"google.golang.org/grpc"
)

func main() {
	opts := grpc.WithInsecure()
	log := hclog.Default()
	cc, err := grpc.Dial("localhost:8080", opts)
	if err != nil {
		log.Error("Some error occured ", err)
	}
	defer cc.Close()

	client := protos.NewOrderServiceClient(cc)

	items := GetItems()

	orderRequest := &protos.OrderDetails{
		UserId:        "0517cc14-9918-4230-a4f1-3670683e3431",
		RestaurantId:  "630b248769fbbb44c75f271d",
		PayableAmount: 2578.34,
		PaymentMethod: protos.PaymentMethod_CARD,
		Order:         items,
		UserLatitude: -3.496589660644531,
		UserLongitude: -62.961456298828125,
	}

	orderResp, _ := client.CreateOrder(context.Background(), orderRequest)
	fmt.Println(orderResp)

	// orderResp, _ := client.FindOrder(context.Background(), &protos.OrderID{OrderId: "62d7c981ccefcbb3d1b858bd"})
	// fmt.Println(orderResp)
	// fmt.Println(orderResp.OrderId)
	// fmt.Println(orderResp.OrderStatus)
	// fmt.Println(orderResp.RestaurantId)

	

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
