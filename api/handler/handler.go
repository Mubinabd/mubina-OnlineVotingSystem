package handler

import (
	pb "project/genproto/public"

	"google.golang.org/grpc"
)

type HandlerStruct struct {
	Candidate  pb.CandidateServiceClient
	Election   pb.ElectionServiceClient
	PublicVote pb.PublicVoteServiceClient
}

func NewHandlerStruct(connClient *grpc.ClientConn) *HandlerStruct {
	return &HandlerStruct{
		Election:   pb.NewElectionServiceClient(connClient),
		Candidate:  pb.NewCandidateServiceClient(connClient),
		PublicVote: pb.NewPublicVoteServiceClient(connClient),
	}
}
