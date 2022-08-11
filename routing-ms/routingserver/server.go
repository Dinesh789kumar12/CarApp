package routingserver

import (
	"context"
	"io"
	"log"
	"sync"

	"github.com/Dinesh789kumar12/CarApp/definitions-store/availabilitypb"
	"github.com/Dinesh789kumar12/CarApp/definitions-store/ratepb"
	"github.com/Dinesh789kumar12/CarApp/definitions-store/routingpb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Server struct {
	routingpb.UnimplementedRoutingServiceServer
}

func (*Server) GetRateBasedonAvailability(stream routingpb.RoutingService_GetRateBasedonAvailabilityServer) error {
	cc, err := grpc.Dial("0.0.0.0:50052", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Error while Dial: %v", err)
	}
	log.Println("rate server started listening on port 50052")
	c := ratepb.NewRoutingServiceClient(cc)
	for {
		req, err := stream.Recv()
		log.Printf("received request from client:%v:", req)
		if err == io.EOF {
			log.Print("no more data")
			break
		}
		if err != nil {
			log.Fatalf("Error when reading client request stream: %v", err)
		}
		reqRate := ratepb.RateRequest{
			Car: &ratepb.Car{
				CarId:   req.GetCar().GetCarId(),
				CarType: req.GetCar().GetCarType(),
			},
		}
		resRate, err := c.GetRate(context.Background(), &reqRate)
		if err != nil {
			log.Fatalf("Error while dailing rate ms:%v", err.Error())
		}

		response := routingpb.RoutingAvailabilityResponse{
			Car: &routingpb.Car{
				CarId:   req.GetCar().GetCarId(),
				CarType: req.GetCar().GetCarType(),
			},
			Location: req.GetLocation(),
			Price:    resRate.GetPrice(),
		}
		res := stream.Send(&response)
		log.Printf("sent to client:%v:", &response)
		if res != nil {
			log.Fatalf("Error when response was sent to the client: %v", res)
		}
	}
	return nil
}

func (*Server) GetCarAvailability(req *routingpb.RoutingAvailabilityCarRequest, stream routingpb.RoutingService_GetCarAvailabilityServer) error {

	cc, err := grpc.Dial("0.0.0.0:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Error while Dial: %v", err)
	}
	log.Println("connection established with Availability MS....")
	defer cc.Close()
	c := availabilitypb.NewAvailabilityServiceClient(cc)
	reqAvailability := availabilitypb.AvailabilityRequest{
		Source: &availabilitypb.Location{
			Latitude:  req.GetSource().GetLatitude(),
			Longitude: req.GetSource().GetLongitude(),
		},
	}
	streamAvailability, err := c.GetAvailability(context.Background(), &reqAvailability)
	if err != nil {
		log.Fatalf("Error while dailing rate ms:%v", err.Error())
	}
	var wg = new(sync.WaitGroup)
	wg.Add(1)
	go func() {
		for {
			resp, err := streamAvailability.Recv()
			if err == io.EOF {
				log.Print("no more data")
				break
			}
			if err != nil {
				log.Fatalf("Error when reading client request stream: %v", err)
			}
			response := routingpb.RoutingAvailabilityCarResponse{
				Car: &routingpb.Car{
					CarId:   resp.GetCar().GetCarId(),
					CarType: resp.GetCar().GetCarType(),
				},
				Location: resp.GetLocation(),
				Distance: resp.GetDistance(),
			}

			stream.Send(&response)
		}
	}()
	wg.Wait()
	return nil
}
