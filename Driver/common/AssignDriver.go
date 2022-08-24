package common

import (
	"fmt"
	"math/rand"
	"time"
)

func PingDrivers(drivers []AvailableDrivers) AvailableDrivers {
	fmt.Println("Length ", len(drivers))
	limit := len(drivers) - 1
	rand.Seed(time.Now().UnixNano())
	n := 1 + rand.Intn(limit-1+1)
	fmt.Println("random number ", n)
	return drivers[n]
}
