package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"

	"github.com/bxcodec/faker/v3"
	"github.com/jainam240101/zomato-clone/scripts/schemas"
	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 4)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

func GenerateUsers(DB *mongo.Database) {
	for i := 0; i < 20; i++ {
		hashedPassword, _ := HashPassword(faker.Password())
		user := schemas.UserModel{
			ID:          primitive.NewObjectID(),
			Name:        faker.FirstName() + " " + faker.LastName(),
			Email:       faker.Email(),
			Password:    hashedPassword,
			PhoneNumber: faker.Phonenumber(),
		}
		val, err := DB.Collection("users").InsertOne(context.TODO(), user)
		if err != nil {
			fmt.Println("error ", err)
			return
		}
		fmt.Println("user ---> ", val)
	}

}

func GenerateRestaurants(DB *mongo.Database) {
	for i := 0; i < 20; i++ {
		hashedPassword, _ := HashPassword(faker.Password())

		restro := schemas.RestaurantModel{
			ID:            primitive.NewObjectID(),
			Name:          faker.FirstName() + " " + faker.LastName(),
			Email:         faker.Email(),
			Latitude:      faker.Latitude(),
			Longitude:     faker.Latitude(),
			Password:      hashedPassword,
			PhoneNumber:   faker.Phonenumber(),
			Manager_Name:  faker.FirstName() + " " + faker.LastName(),
			Manager_Phone: faker.Phonenumber(),
			Listing_Date:  faker.Date(),
		}
		val, err := DB.Collection("restaurants").InsertOne(context.TODO(), restro)
		if err != nil {
			fmt.Println("error ", err)
			return
		}
		fmt.Println("restro ---> ", val)
	}
}

func GenerateDrivers(DB *mongo.Database) {
	for i := 0; i < 20; i++ {
		hashedPassword, _ := HashPassword(faker.Password())

		driver := schemas.DriverModel{
			ID:             primitive.NewObjectID(),
			Name:           faker.FirstName() + " " + faker.LastName(),
			Email:          faker.Email(),
			Password:       hashedPassword,
			PhoneNumber:    faker.Phonenumber(),
			Current_Rating: 0,
			Joining_Date:   time.Now(),
		}
		val, err := DB.Collection("drivers").InsertOne(context.TODO(), driver)
		if err != nil {
			fmt.Println("error ", err)
			return
		}
		fmt.Println("drivers ---> ", val)
	}
}

func ConnectDb() *mongo.Database {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		panic(err)
	}
	if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
		log.Fatal("Some error occured ", err)
		return nil
	}
	return client.Database("Zomato-Clone")
}

func main() {
	godotenv.Load("../.env")

	DB := ConnectDb()

	// GenerateUsers(DB)
	GenerateRestaurants(DB)
	// GenerateDrivers(DB)
}
