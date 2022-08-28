package db

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var OrderDB *mongo.Collection
var RestaurantDB *mongo.Collection

func ConnectDb() {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		panic(err)
	}
	if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
		log.Fatal("Some error occured ", err)
	}
	OrderDB = client.Database("Zomato-Clone").Collection("Orders")
	RestaurantDB = client.Database("Zomato-Clone").Collection("restaurants")
}
