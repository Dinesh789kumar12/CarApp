package utils

import (
	"math"
	"math/rand"
	"time"

	data "github.com/Dinesh789kumar12/CarApp/availability-ms/data"
)

func RandomNumberGenerator() (x int, y int) {
	rand.Seed(time.Now().UnixNano())

	a := rand.Intn(3)
	b := rand.Intn(3)

	return a, b
}

//Distance b/w longitude and latitude
func Distance(x1 int, y1 int, x2 int, y2 int) int32 {
	return int32(math.Sqrt(math.Pow(float64(x2-x1), 2) + math.Pow(float64(y2-y1), 2)))
}

func GetCoordinatebyName(s string) (int, int) {
	var x2, y2 int
	for i, v := range data.LocationName {
		for i2, _ := range v {
			if data.LocationName[i][i2] == s {
				x2 = i
				y2 = i2
			}
		}
	}
	return x2, y2
}
