package schemas

import (
	"github.com/jainam240101/zomato-clone/Protos/OrderProtos"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type OrderStruct struct {
	ID            primitive.ObjectID   `bson:"_id"`
	UserId        string               `bson:"userId"`
	RestaurantId  string               `bson:"restaurantId"`
	OrderStatus   string               `bson:"orderStatus"`
	PayableAmount float64              `bson:"payableAmount"`
	PaymentMethod string               `bson:"paymentMethod"`
	Order         []*OrderProtos.Order `bson:"order"`
}
