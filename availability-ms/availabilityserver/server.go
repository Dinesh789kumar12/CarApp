package availabilityserver

import (
	"log"
	"sort"
	"sync"
	"time"

	"github.com/Dinesh789kumar12/CarApp/availability-ms/data"
	"github.com/Dinesh789kumar12/CarApp/availability-ms/utils"
	"github.com/Dinesh789kumar12/CarApp/definitions-store/availabilitypb"
)

type Server struct {
	availabilitypb.UnimplementedAvailabilityServiceServer
}

func (*Server) GetAvailability(req *availabilitypb.AvailabilityRequest, strm availabilitypb.AvailabilityService_GetAvailabilityServer) error {
	usrLoc := req.GetSource()
	lat := usrLoc.Latitude
	lon := usrLoc.Longitude
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			data.BookrandomCar() //Book Random Car from slice
			sendAvailableCarToClient(data.Carpool, int(lat), int(lon))

			//Sorting on carDistance
			sort.Slice(data.AvailableCarPool, func(p, q int) bool {
				return data.AvailableCarPool[p].CarDistance < data.AvailableCarPool[q].CarDistance
			})

			for _, c := range data.AvailableCarPool {
				res := &availabilitypb.AvailabilityResponse{
					Car: &availabilitypb.Car{
						CarId:   c.CarId,
						CarType: c.Model,
					},
					Location: c.Location,
					Distance: c.CarDistance,
				}
				strm.Send(res)
				log.Printf("Sent:%v", res)
			}
			data.AvailableCarPool = nil
			time.Sleep(20 * time.Second)
		}
	}()
	wg.Wait()
	return nil
}

func sendAvailableCarToClient(carPool []data.Car, x int, y int) {
	for _, c := range carPool {
		if c.Available {
			a, b := utils.RandomNumberGenerator()
			location := data.LocationName[a][b]
			distance := utils.Distance(a, b, x, y)
			data.AvailableCarPool = append(data.AvailableCarPool, data.AvailableCar{CarId: c.CarId, Available: c.Available, Model: c.Model, Location: location, CarDistance: distance})
			time.Sleep(2 * time.Nanosecond)
		}
	}
}
