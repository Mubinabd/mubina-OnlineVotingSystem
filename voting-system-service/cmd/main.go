package main

import (
	"fmt"
	"log"
	"project/api"
	handlers "project/api/handler"
	cf "project/config"
	pb "project/genproto/public"
	"project/service"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	config := cf.Load()

	conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	candS := pb.NewCandidateServiceClient(conn)
	elecS := pb.NewElectionServiceClient(conn)
	publicVS := pb.NewPublicVoteServiceClient(conn)

	cs := service.NewCandidateService(candS)
	es := service.NewElectionService(elecS)
	ps := service.NewPublicVoteService(publicVS)

	h := handlers.NewHTTPHandler(service.Service{CM: cs, EM: es, PVM: ps})

	r := api.NewGin(h)

	fmt.Printf("Server started on port %s\n", config.HTTPPort)
	err = r.Run(config.HTTPPort)
	if err != nil {
		log.Fatal(err)
	}

}
