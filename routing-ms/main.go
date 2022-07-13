package main

import (
	"log"
	"net"

	"github.com/Dinesh789kumar12/CarApp/definitions-store/routingpb"
	"github.com/Dinesh789kumar12/CarApp/routing-ms/routingserver"
	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", "0.0.0.0:3000")
	if err != nil {
		log.Fatalf("error while listening on port 3000: %v", err.Error())
	}
	log.Println("routing server started listening on port 3000")
	routingServiceServer := grpc.NewServer()
	routingpb.RegisterRoutingServiceServer(routingServiceServer, &routingserver.Server{})
	if err := routingServiceServer.Serve(lis); err != nil {
		log.Fatalf("error while availabilityServer.Serve: %v", err.Error())
	}
}
