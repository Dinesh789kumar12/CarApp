package main

import (
	"context"
	"io"
	"log"
	"sync"
	"time"

	"github.com/Dinesh789kumar12/CarApp/definitions-store/routingpb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	log.Println("**********************************************************************************")
	log.Println("")
	log.Println("\t\tMobile App Client\t\t")
	log.Println("")
	log.Println("**********************************************************************************")

	//Routing MS
	cc1, err := grpc.Dial("0.0.0.0:3000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Error while Dial: %v", err)
	}
	log.Println("routing server started listening on port 3000")

	client := routingpb.NewRoutingServiceClient(cc1)
	routingReq := routingpb.RoutingAvailabilityCarRequest{
		Source: &routingpb.Location{
			Latitude:  1,
			Longitude: 2,
		},
	}
	log.Println("Sending the present coordinate :")
	log.Println("Latitute: 1")
	log.Println("Longitute: 2")

	strm, err := client.GetCarAvailability(context.Background(), &routingReq)
	if err != nil {
		log.Fatalf("error while GetCarAvailability: %v", err.Error())
	}
	log.Printf("Sent the mobile app co-ordinates and waiting for availble car for my location")

	stream, err := client.GetRateBasedonAvailability(context.Background())
	if err != nil {
		log.Fatalf("error while client.GetAvailability: %v", err.Error())
	}

	log.Println("Waiting for the available car....")
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
				Car: &routingpb.Car{
					CarId:   resp.GetCar().GetCarId(),
					CarType: resp.GetCar().GetCarType(),
				},
				Location: resp.GetLocation(),
			}
			if err := stream.Send(&routingRequest); err != nil {
				log.Fatalf("Error while send to Routing Server: %v", err)
			}

			//log.Printf("sent request to client: %s", &routingRequest)
		}
		if err := stream.CloseSend(); err != nil {
			log.Fatalf("Error while close send to RS: %v", err)
		}
	}()

	// Use a go routine to receive response messages from the server
	wg.Add(1)
	go func() {
		var count = 1
		for {
			if count == 4 {
				count = 1
			}
			//main code
			res, err := stream.Recv()
			if err != nil {
				break
			}
			if count == 1 {
				log.Println("Received top cars that near to my location :", time.Now().Format(time.Stamp))
			}
			log.Printf("%v:received response from server :%v", count, res)
			count++
		}
	}()
	wg.Wait()
}
