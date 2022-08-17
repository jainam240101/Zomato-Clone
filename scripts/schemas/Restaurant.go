package schemas

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type RestaurantModel struct {
	ID            primitive.ObjectID `bson:"_id"`
	Name          string             `json:"name"`
	Email         string             `json:"email"`
	Password      string             `json:"password"`
	PhoneNumber   string             `json:"phoneNumber"`
	Manager_Name  string             `json:"managerName"`
	Manager_Phone string             `json:"managerPhone"`
	Listing_Date  string             `json:"listingDate"`
}
