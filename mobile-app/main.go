package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/Dinesh789kumar12/CarApp/definitions-store/routingpb"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

type userRequest struct {
	Location location `json:"source"`
}

type location struct {
	Latitude  int `json:"latitude"`
	Longitude int `json:"longitude"`
}

type availabilityResponse struct {
	CarName  string `json:"carName"`
	Location string `json:"location"`
	Distance int    `json:"distance"`
}

func main() {
	router := gin.Default()
	rg := router.Group("api/v1/carapp")
	{
		rg.GET("/car", fetchAvailableCarNearby)
	}

	router.Run()

}

func fetchAvailableCarNearby(c *gin.Context) {
	myCounter := 1
	myCounter++
	log.Println(myCounter)

	var usrLoc userRequest
	if err := json.NewDecoder(c.Request.Body).Decode(&usrLoc); err != nil {
		log.Fatalf("error while json unmarshal: %v", err.Error())
	}
	routingRequest := routingpb.RoutingAvailabilityRequest{
		Source: &routingpb.Location{
			Latitude:  int32(usrLoc.Location.Latitude),
			Longitude: int32(usrLoc.Location.Longitude),
		},
	}
	cc, err := grpc.Dial("0.0.0.0:3000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Error while Dial: %v", err)
	}

	client := routingpb.NewRoutingServiceClient(cc)
	strm, err := client.GetAvailability(context.Background())
	if err != nil {
		log.Fatalf("error while client.GetAvailability: %v", err.Error())
	}
	strm.Send(&routingRequest)
	resp := &routingpb.ListRoutingAvailabilityResponse{}
	for {
		resp, err = strm.Recv()
		if err != nil {
			break
		}
		c.JSON(http.StatusOK, &resp)
		break
	}
}
