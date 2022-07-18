package main

import (
	"log"
	"net"

	"github.com/Dinesh789kumar12/CarApp/definitions-store/ratepb"
	"github.com/Dinesh789kumar12/CarApp/rate-ms/rateserver"
	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", "0.0.0.0:50052")
	if err != nil {
		log.Fatalf("error while listening on port 50051: %v", err.Error())
	}
	log.Println("rate server started listening on port 50052")
	log.Printf("Waiting for availble car....")
	rateServer := grpc.NewServer()
	ratepb.RegisterRoutingServiceServer(rateServer, &rateserver.Server{})
	if err := rateServer.Serve(lis); err != nil {
		log.Fatalf("error while rateServer Serve: %v", err.Error())
	}
}
