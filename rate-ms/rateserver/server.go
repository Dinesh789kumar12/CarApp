package rateserver

import (
	"log"
	"time"

	"github.com/Dinesh789kumar12/CarApp/availability-ms/data"
	"github.com/Dinesh789kumar12/CarApp/availability-ms/utils"
	"github.com/Dinesh789kumar12/CarApp/definitions-store/ratepb"
	"github.com/Dinesh789kumar12/CarApp/rate-ms/rateutils"
)

type Server struct {
	ratepb.UnimplementedRoutingServiceServer
}

func (*Server) GetRate(req *ratepb.RateRequest, stream ratepb.RoutingService_GetRateServer) error {
	//Getting request from client
	rateDateReq := req.GetDate()
	rateSourceReq := req.GetSource()
	rateDestinationReq := req.GetDestination()

	//Calculating Distance
	x1, y1 := utils.GetCoordinatebyName(rateSourceReq)
	x2, y2 := utils.GetCoordinatebyName(rateDestinationReq)
	rideDistance := utils.Distance(x1, x2, y1, y2)

	for _, c := range data.Carpool {
		if c.Available {
			res := &ratepb.RateResponse{
				CarType:  c.Model,
				Location: "Delhi",
				Distance: 0,
				Price:    rateutils.GetRatebyTimeAndDistance(c.Model, rateDateReq, rideDistance),
			}
			stream.Send(res)
			log.Printf("Sent:%v", res)
			time.Sleep(2 * time.Nanosecond)
		}
	}
	return nil
}
