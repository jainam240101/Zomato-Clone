package schemas

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserModel struct {
	ID          primitive.ObjectID `bson:"_id"`
	Name        string             `bson:"name"`
	Email       string             `bson:"email"`
	Password    string             `bson:"password"`
	PhoneNumber string             `bson:"phoneNumber"`
}
