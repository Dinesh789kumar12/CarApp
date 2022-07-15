package rateserver

import (
	"context"
	"log"
	"time"

	"github.com/Dinesh789kumar12/CarApp/definitions-store/ratepb"
	"github.com/Dinesh789kumar12/CarApp/rate-ms/rateutils"
)

type Server struct {
	ratepb.UnimplementedRoutingServiceServer
}

func (*Server) GetRate(ctx context.Context, req *ratepb.RateRequest) (*ratepb.RateResponse, error) {
	reqCar := req.GetCarType()
	price := rateutils.GetRatebyTimeAndDistance(reqCar, time.Now().String(), 2)
	resRate := ratepb.RateResponse{
		CarId:   req.GetCarId(),
		CarType: reqCar,
		Price:   price,
	}
	log.Printf("sent to client:%v", &resRate)
	return &resRate, nil
}
