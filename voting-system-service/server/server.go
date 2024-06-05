package main

import (
	"log"
	pb "project/genproto/public"
	"project/service"
	"project/storage/postgres"
	"net"

	"google.golang.org/grpc"
)

func main() {
	_, err := postgres.NewPostgresManager()
	if err != nil {
		log.Fatal(err)
	}
	liss, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatal(err)
	}
	s := grpc.NewServer()

	pb.RegisterCandidateServiceServer(s, service.NewCandidateService(pb.NewCandidateServiceClient(&grpc.ClientConn{})))
	pb.RegisterElectionServiceServer(s, service.NewElectionService(pb.NewElectionServiceClient(&grpc.ClientConn{})))
	pb.RegisterPublicVoteServiceServer(s, service.NewPublicVoteService(pb.NewPublicVoteServiceClient(&grpc.ClientConn{})))
	

	log.Printf("server listening at %v", liss.Addr())

	if err := s.Serve(liss); err != nil {
		log.Fatal(err)
	}
}
