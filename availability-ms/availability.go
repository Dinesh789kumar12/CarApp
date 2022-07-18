package main

import (
	"fmt"
	"log"
	"net"

	"github.com/Dinesh789kumar12/CarApp/availability-ms/availabilityserver"
	"github.com/Dinesh789kumar12/CarApp/definitions-store/availabilitypb"
	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("error while listening on port 50051: %v", err.Error())
	}
	log.Println("availability server started listening on port 50051")
	log.Printf("Waiting for mobile app coordinates...")
	availabilityServer := grpc.NewServer()
	availabilitypb.RegisterAvailabilityServiceServer(availabilityServer, &availabilityserver.Server{})

	if err := availabilityServer.Serve(lis); err != nil {
		log.Fatalf("error while availabilityServer.Serve: %v", err.Error())
	}
	fmt.Print()

}
