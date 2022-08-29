package handlers

import (
	"context"

	"github.com/hashicorp/go-hclog"
	"github.com/jainam240101/zomato-clone/Orders/schemas"
	"github.com/jainam240101/zomato-clone/Orders/utils"
	protos "github.com/jainam240101/zomato-clone/Protos/OrderProtos"
	restaurantProtos "github.com/jainam240101/zomato-clone/Protos/RestaurantProtos"
	restroStruct "github.com/jainam240101/zomato-clone/scripts/schemas"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Server struct {
	Log      hclog.Logger
	RestroDB *mongo.Collection
	DB       *mongo.Collection
	Restro   restaurantProtos.RestaurantServiceClient
	protos.UnimplementedOrderServiceServer
}

func NewServer(log hclog.Logger, DB *mongo.Collection, restro restaurantProtos.RestaurantServiceClient) *Server {
	return &Server{Log: log, DB: DB, Restro: restro}
}

func (s *Server) CreateOrder(ctx context.Context, request *protos.OrderDetails) (*protos.OrderResponse, error) {
	s.Log.Info("Creating Order ")

	var restro restroStruct.RestaurantModel
	restroobjectId, _ := primitive.ObjectIDFromHex(request.RestaurantId)
	s.RestroDB.FindOne(context.TODO(), bson.M{"_id": restroobjectId}).Decode(&restro)

	distance := utils.Distance(float64(request.UserLatitude), float64(request.UserLongitude), restro.Latitude, restro.Longitude, "Kilometer")

	s.Log.Info("Distance ", distance)

	OrderDetails := schemas.OrderStruct{
		ID:             primitive.NewObjectID(),
		UserId:         request.UserId,
		RestaurantId:   request.RestaurantId,
		BillAmount:     float64(request.PayableAmount),
		DeliveryCharge: float64(distance*1.5) + 15,
		PaymentMethod:  request.PaymentMethod.String(),
		PaymentStatus:  "Pending",
		OrderStatus:    "Pending",
		Order:          request.Order,
	}
	val, err := s.DB.InsertOne(context.TODO(), OrderDetails)
	if err != nil {
		s.Log.Error("Some Error occured ", err)
		return nil, err
	}

	orderResponse := &protos.OrderResponse{
		OrderId:             val.InsertedID.(primitive.ObjectID).String(),
		PaymentMethod:       request.PaymentMethod.String(),
		UserId:              request.UserId,
		PaymentStatus:       "Pending",
		OrderStatus:         "Pending",
		RestaurantId:        request.RestaurantId,
		BillAmount:          request.PayableAmount,
		Order:               request.Order,
		RestaurantLatitude:  float32(restro.Latitude),
		RestaurantLongitude: float32(restro.Longitude),
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
	orderResponse := &protos.OrderResponse{
		OrderId:        request.OrderId,
		PaymentMethod:  order.PaymentMethod,
		UserId:         order.UserId,
		PaymentStatus:  order.PaymentStatus,
		OrderStatus:    order.OrderStatus,
		RestaurantId:   order.RestaurantId,
		BillAmount:     float32(order.BillAmount),
		DeliveryCharge: float32(order.DeliveryCharge),
		Order:          order.Order,
	}
	return orderResponse, nil
}
