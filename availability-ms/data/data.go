package data

import (
	"log"
	"math/rand"
	"time"
)

var LocationName = [][]string{
	{"New Delhi", "Mandoli", "Tughlaqabad", "Yammuna Vihar", "Bhajanpura"},
	{"Ambala", "Gurugram ", "Faridabad", "Bhiwani", "Fatehabad"},
	{"Marathahalli", "BTM Layout", "Banashankari", "Domlur", " Rajajinagar"},
	{"Bhopura", "Mohan Nagar", "Ghaziabad", "Anand Vihar", "Noida"},
}

type Car struct {
	CarId     int32
	Available bool
	Model     string
}

var Carpool = []Car{
	{111, false, "Hyundai"}, {112, false, "Sedan"}, {113, true, "Maruti Suzuki"}, {114, true, "Hatchback"}, {115, true, "Hyundai Xcent"}, {116, true, "Maruti Suzuki"}, {117, true, "Maruti Suzuki"}, {118, true, "Maruti Suzuki"}, {119, true, "Hatchback"}, {120, true, "Hyundai Xcent"},
	{121, true, "Sedan"}, {122, true, "Ford"},
}

func BookrandomCar() {

	rand.Seed(time.Now().UnixNano())
	randomIndex := rand.Intn(len(Carpool))
	selectedCar := Carpool[randomIndex]
	if selectedCar.Available {
		selectedCar.Available = false
	} else {
		selectedCar.Available = true
	}
	Carpool[randomIndex] = selectedCar
	log.Printf("Randomly selected car : "+"%+v\n", selectedCar)
}

type AvailableCar struct {
	CarId       int32
	Available   bool
	Model       string
	Location    string
	CarDistance int32
}

var AvailableCarPool = []AvailableCar{}
