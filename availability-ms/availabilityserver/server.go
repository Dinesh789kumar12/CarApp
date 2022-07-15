package availabilityserver

import (
	"log"
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
	for {
		data.BookrandomCar()
		for _, c := range data.Carpool {
			var a, b int
			if c.Available {
				a, b = utils.RandomNumberGenerator()
				CarLocation := data.LocationName[a][b]
				m := utils.Distance(a, b, int(lat), int(lon))
				res := &availabilitypb.AvailabilityResponse{
					CarId:    c.CarId,
					CarType:  c.Model,
					Location: CarLocation,
					Distance: m,
				}
				strm.Send(res)
				log.Printf("Sent:%v", res)
				time.Sleep(2 * time.Nanosecond)
			}
		}
		time.Sleep(20 * time.Second)
	}
}
