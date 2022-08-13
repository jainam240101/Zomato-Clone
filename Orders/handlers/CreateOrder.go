package handlers

import (
	"context"

	"github.com/hashicorp/go-hclog"
	"github.com/jainam240101/zomato-clone/Orders/schemas"
	protos "github.com/jainam240101/zomato-clone/Protos/OrderProtos"
	restaurantProtos "github.com/jainam240101/zomato-clone/Protos/RestaurantProtos"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Server struct {
	Log    hclog.Logger
	DB     *mongo.Collection
	Restro restaurantProtos.RestaurantServiceClient
	protos.UnimplementedOrderServiceServer
}

func NewServer(log hclog.Logger, DB *mongo.Collection, restro restaurantProtos.RestaurantServiceClient) *Server {
	return &Server{Log: log, DB: DB, Restro: restro}
}

func (s *Server) CreateOrder(ctx context.Context, request *protos.OrderDetails) (*protos.OrderResponse, error) {
	s.Log.Info("Creating Order ")

	OrderDetails := schemas.OrderStruct{
		ID:            primitive.NewObjectID(),
		UserId:        request.UserId,
		RestaurantId:  request.RestaurantId,
		PayableAmount: float64(request.PayableAmount),
		PaymentMethod: request.PaymentMethod.String(),
		OrderStatus:   "Pending",
		Order:         request.Order,
	}

	val, err := s.DB.InsertOne(context.TODO(), OrderDetails)
	if err != nil {
		s.Log.Error("Some Error occured ", err)
		return nil, err
	}

	orderResponse := &protos.OrderResponse{
		OrderId:       val.InsertedID.(primitive.ObjectID).String(),
		PaymentMethod: request.PaymentMethod.String(),
		UserId:        request.UserId,
		OrderStatus:   "Pending",
		RestaurantId:  request.RestaurantId,
		PayableAmount: request.PayableAmount,
		Order:         request.Order,
	}

	s.Log.Info("Sending Orders to restro")

	data, err := s.Restro.AcceptOrder(context.Background(), orderResponse)
	if err != nil {
		s.Log.Error("Some Error Occured", err)
		return nil, err
	}

	filter := bson.D{{"_id", OrderDetails.ID}}
	update := bson.D{{"$set", bson.D{{"orderStatus", "Accepted"}}}}
	mongoData, err := s.DB.UpdateOne(context.TODO(), filter, update)
	s.Log.Info("Data ", mongoData)
	if err != nil {
		s.Log.Error("Error ", err)
	}

	return data, nil
}

func (s *Server) FindOrder(ctx context.Context, request *protos.OrderID) (*protos.OrderResponse, error) {
	s.Log.Info("Finding Order ")
	var order schemas.OrderStruct

	objectId, _ := primitive.ObjectIDFromHex(request.OrderId)

	s.DB.FindOne(context.TODO(), bson.M{"_id": objectId}).Decode(&order)
	// if err != nil {
	// 	s.Log.Error("Some Error occured ", err)
	// 	return nil, err
	// }
	orderResponse := &protos.OrderResponse{
		OrderId:       request.OrderId,
		PaymentMethod: order.PaymentMethod,
		UserId:        order.UserId,
		OrderStatus:   order.OrderStatus,
		RestaurantId:  order.RestaurantId,
		PayableAmount: float32(order.PayableAmount),
		Order:         order.Order,
	}
	return orderResponse, nil
}
