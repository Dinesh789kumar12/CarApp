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
	reqCar := &ratepb.Car{
		CarId:   req.GetCar().GetCarId(),
		CarType: req.GetCar().GetCarType(),
	}
	price := rateutils.GetRatebyTimeAndDistance(req.GetCar().GetCarType(), time.Now().String(), 2)
	resRate := ratepb.RateResponse{
		Car:   reqCar,
		Price: price,
	}
	log.Printf("sent to client:%v", &resRate)
	return &resRate, nil
}
