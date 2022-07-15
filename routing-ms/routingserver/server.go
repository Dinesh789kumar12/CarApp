package routingserver

import (
	"context"
	"io"
	"log"

	"github.com/Dinesh789kumar12/CarApp/definitions-store/ratepb"
	"github.com/Dinesh789kumar12/CarApp/definitions-store/routingpb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Server struct {
	routingpb.UnimplementedRoutingServiceServer
}

func (*Server) GetAvailability(stream routingpb.RoutingService_GetAvailabilityServer) error {
	cc, err := grpc.Dial("0.0.0.0:50052", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Error while Dial: %v", err)
	}
	log.Println("rate server started listening on port 50052")
	c := ratepb.NewRoutingServiceClient(cc)
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			log.Fatalf("Error when reading client request stream: %v", err)
		}
		reqRate := ratepb.RateRequest{
			CarId:   req.GetCarId(),
			CarType: req.GetCarType(),
		}
		resRate, err := c.GetRate(context.Background(), &reqRate)
		if err != nil {
			log.Fatalf("Error while dailing rate ms:%v", err.Error())
		}
		pr := resRate.GetPrice()
		response := routingpb.RoutingAvailabilityResponse{
			CarId:    req.GetCarId(),
			CarType:  req.GetCarType(),
			Location: "New Delhi",
			Price:    pr,
		}
		res := stream.Send(&response)
		log.Printf("sent to client:%v:", &response)
		if res != nil {
			log.Fatalf("Error when response was sent to the client: %v", res)
		}
	}
	return nil
}
