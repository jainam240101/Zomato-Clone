package redis

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/go-redis/redis/v8"
	"github.com/jainam240101/zomato-clone/Driver/common"
	"github.com/jainam240101/zomato-clone/Driver/db"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const key = "drivers"

func (c *RedisClient) AddDriverLocation(lng, lat float64, id string) {
	c.GeoAdd(
		context.TODO(),
		key,
		&redis.GeoLocation{Longitude: lng, Latitude: lat, Name: id},
	)
}

func (c *RedisClient) RemoveDriverLocation(id string) {
	c.ZRem(context.TODO(), key, id)
}

func (c *RedisClient) SearchDrivers(limit int, lat, lng, r float64, orderId string) ([]byte, error) {
	/*
		WITHDIST: Also return the distance of the returned items from    the specified center. The distance is returned in the same unit as the unit specified as the radius argument of the command.

		WITHCOORD: Also return the longitude,latitude coordinates of the  matching items.

		WITHHASH: Also return the raw geohash-encoded sorted set score of the item, in the form of a 52 bit unsigned integer. This is only useful for low level hacks or debugging and is otherwise of little interest for the general user.
	*/
	res, _ := c.GeoRadius(context.TODO(), key, lng, lat, &redis.GeoRadiusQuery{
		Radius:      r,
		Unit:        "km",
		WithGeoHash: false,
		WithCoord:   false,
		WithDist:    true,
		Count:       limit,
		Sort:        "ASC",
	}).Result()

	var drivers []common.AvailableDrivers

	for i := 0; i < len(res); i++ {
		val := c.Get(context.TODO(), res[i].Name)
		final, _ := val.Result()
		if final == "" {
			drivers = append(drivers, common.AvailableDrivers{
				DriverId: res[i].Name,
				Distance: res[i].Dist,
			})
		}
	}
	selectedDriver := common.PingDrivers(drivers)
	data, err := json.Marshal(selectedDriver)
	if err != nil {
		return nil, err
	}

	id, _ := primitive.ObjectIDFromHex(orderId)
	result, err := db.OrderDB.UpdateOne(
		context.TODO(),
		bson.M{"_id": id},
		bson.D{
			{"$set", bson.D{{"driverId", selectedDriver.DriverId}}},
		},
	)
	if err != nil {
		return nil, err
	}
	fmt.Println("Result is ", result)
	err = c.Set(context.TODO(), selectedDriver.DriverId, "busy", 0).Err()
	if err != nil {
		panic(err)
	}
	return data, nil
}

func (c *RedisClient) UpdateOrder(driverId, orderId string) (bool, error) {
	id, _ := primitive.ObjectIDFromHex(orderId)
	_, err := db.OrderDB.UpdateOne(
		context.TODO(),
		bson.M{"_id": id},
		bson.D{
			{"$set", bson.D{{"orderStatus", "Delivered"}}},
		},
	)
	if err != nil {
		return false, err
	}
	c.Del(context.TODO(), driverId)
	return true, nil
}
