package main

import (
	"context"
	"io"
	"log"
	"sync"

	"github.com/Dinesh789kumar12/CarApp/definitions-store/availabilitypb"
	"github.com/Dinesh789kumar12/CarApp/definitions-store/routingpb"
	"google.golang.org/grpc"
)

func main() {

	cc, err := grpc.Dial("0.0.0.0:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Error while Dial: %v", err)
	}
	log.Println("availability server started listening on port 50051")
	cc1, err := grpc.Dial("0.0.0.0:3000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Error while Dial: %v", err)
	}
	log.Println("routing server started listening on port 3000")
	availabilityReq := availabilitypb.AvailabilityRequest{
		Source: &availabilitypb.Location{
			Latitude:  1,
			Longitude: 2,
		},
	}
	c := availabilitypb.NewAvailabilityServiceClient(cc)
	client := routingpb.NewRoutingServiceClient(cc1)
	strm, err := c.GetAvailability(context.Background(), &availabilityReq)
	if err != nil {
		log.Fatalf("error while c.GetAvailability: %v", err.Error())
	}
	stream, err := client.GetAvailability(context.Background())
	if err != nil {
		log.Fatalf("error while client.GetAvailability: %v", err.Error())
	}
	var wg = new(sync.WaitGroup)
	wg.Add(1)
	go func() {
		for {
			resp, err := strm.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatalf("Error when receiving response: %v", err)
			}
			routingRequest := routingpb.RoutingAvailabilityRequest{
				CarId:   resp.GetCarId(),
				CarType: resp.GetCarType(),
			}
			if err := stream.Send(&routingRequest); err != nil {
				log.Fatalf("Error while send to Routing Server: %v", err)
			}

			//time.Sleep(100 * time.Millisecond)
			log.Printf("sent request to client: %s", &routingRequest)
		}
		if err := stream.CloseSend(); err != nil {
			log.Fatalf("Error while close send to RS: %v", err)
		}
	}()

	// Use a go routine to receive response messages from the server
	wg.Add(1)
	go func() {
		for {
			res, err := stream.Recv()
			if err != nil {
				break
			}
			log.Printf("resceived response from server :%v", res)
		}
	}()
	wg.Wait()
}
