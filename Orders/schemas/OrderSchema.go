package schemas

import (
	"github.com/jainam240101/zomato-clone/Protos/OrderProtos"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type OrderStruct struct {
	ID             primitive.ObjectID   `bson:"_id"`
	UserId         string               `bson:"userId"`
	RestaurantId   string               `bson:"restaurantId"`
	PaymentStatus  string               `bson:"paymentStatus"`
	OrderStatus    string               `bson:"orderStatus"`
	BillAmount     float64              `bson:"billAmount"`
	DeliveryCharge float64              `bson:"deliveryCharge"`
	PaymentMethod  string               `bson:"paymentMethod"`
	Order          []*OrderProtos.Order `bson:"order"`
}
