package main

import (
	"log"
	"project/api/handler"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"project/api"
	_ "project/api/docs"
)

// @title Online Voting System API
// @version 1.0
// @description API for managing online voting system resources
// @host localhost:8080
// @BasePath /api/v1
// @in header
// @name Authorization
func main() {
	connVote, err := grpc.NewClient("localhost:8081", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect to Vote service: %v", err)
	}
	defer connVote.Close()

	connPublic, err := grpc.NewClient("localhost:8082", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect to Vote service: %v", err)
	}
	defer connPublic.Close()

	h := handler.HandlerStruct{}

	engine := api.NewGin(&h)
	if err := engine.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
