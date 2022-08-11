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
	//getting value from User
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

			for i, c := range data.AvailableCarPool {
				//show top 3 car to user
				if i == 3 {
					break
				}
				res := &availabilitypb.AvailabilityResponse{
					Car: &availabilitypb.Car{
						CarId:   c.CarId,
						CarType: c.Model,
					},
					Location: c.Location,
					Distance: c.CarDistance,
				}
				strm.Send(res)             //sending response back to client
				log.Printf("Sent:%v", res) //logging messsage in Availability Server after sending
			}
			data.AvailableCarPool = nil
			time.Sleep(20 * time.Second) //wait for 20 second for show the next available car
		}
	}()
	wg.Wait() // wait for all the goroutines launched here to finish
	return nil
}

func sendAvailableCarToClient(carPool []data.Car, x int, y int) {
	for _, c := range carPool {
		if c.Available {
			a, b := utils.RandomNumberGenerator()
			location := data.LocationName[a][b]    //mocking location on basis of getting value from Random Generator
			distance := utils.Distance(a, b, x, y) //calculating distance b/w User location and Car location
			data.AvailableCarPool = append(data.AvailableCarPool, data.AvailableCar{CarId: c.CarId, Available: c.Available, Model: c.Model, Location: location, CarDistance: distance})
			time.Sleep(2 * time.Nanosecond)
		}
	}
}
