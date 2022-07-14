package rateutils

import "time"

var TimeLayout = "15:04"

var Car_fare = map[string]float32{"Hyundai": 10.75, "Sedan": 11.51, "Hatchback": 19.5, "Hyundai Xcent": 18,
	"Maruti Suzuki": 15, "Ford": 14}

type RateTime struct {
	startTime   string
	endTime     string
	ratePerTime float32
}

var TimeFare = []RateTime{
	{"00:00", "06:00", 1.0}, {"06:00", "09:00", 1.5}, {"09:00", "12:00", 1.8},
	{"12:00", "18:00", 1.1}, {"18:00", "22:00", 1.8}, {"22:00", "24:00", 1.2}}

func GetRatebyTimeAndDistance(carType string, t1 string, distance int32) float32 {
	var timeValue float32
	check, _ := time.Parse(TimeLayout, t1)
	for _, id := range TimeFare {
		start, _ := time.Parse(TimeLayout, id.startTime)
		end, _ := time.Parse(TimeLayout, id.endTime)
		b1 := inTimeSpan(start, end, check)
		if b1 {
			timeValue = id.ratePerTime
			break
		}
	}
	price := Car_fare[carType] * float32(distance) * timeValue
	return price
}

func inTimeSpan(start, end, check time.Time) bool {
	if start.Before(end) {
		return !check.Before(start) && !check.After(end)
	}
	if start.Equal(end) {
		return check.Equal(start)
	}
	return !start.After(check) || !end.Before(check)
}
