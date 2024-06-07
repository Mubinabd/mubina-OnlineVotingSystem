package main

import (
	"log"
	"project/api-gateway"
	"project/api-gateway/handler"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	connVote, err := grpc.NewClient("localhost:50050", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to connect:%v", err)
	}
	defer connVote.Close()
	
	engine := api.NewGin(handler.NewHandlerStruct(connVote))
	engine.Run(":8070")
}
