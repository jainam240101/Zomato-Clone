package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	redisModule "github.com/jainam240101/zomato-clone/Driver/Redis"
)

var redis *redisModule.RedisClient

func NewHandlers() {
	redis = redisModule.ConnectRedis()
}

func Search(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Search")
	body := struct {
		Lat   float64 `json:"lat"`
		Lng   float64 `json:"lng"`
		Limit int     `json:"limit"`
	}{}
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		log.Printf("could not decode request: %v", err)
		http.Error(w, "could not decode request", http.StatusInternalServerError)
		return
	}
	drivers := redis.SearchDrivers(body.Limit, body.Lat, body.Lng, 100)
	data, err := json.Marshal(drivers)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

func Tracking(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Tracking")
	// crate an anonymous struct for driver data.
	var driver = struct {
		ID  string  `json:"id"`
		Lat float64 `json:"lat"`
		Lng float64 `json:"lng"`
	}{}
	if err := json.NewDecoder(r.Body).Decode(&driver); err != nil {
		log.Printf("could not decode request: %v", err)
		http.Error(w, "could not decode request", http.StatusInternalServerError)
		return
	}
	// Add new location
	// You can save locations in another db
	redis.AddDriverLocation(driver.Lng, driver.Lat, driver.ID)
	w.WriteHeader(http.StatusOK)
}
