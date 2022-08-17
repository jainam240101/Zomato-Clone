package schemas

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type DriverModel struct {
	ID             primitive.ObjectID
	Name           string    `bson:"name"`
	Email          string    `bson:"email"`
	Password       string    `bson:"password"`
	PhoneNumber    string    `bson:"phoneNumber"`
	Current_Rating float64   `bson:"currentRating"`
	Joining_Date   time.Time `bson:"start_date"`
}
