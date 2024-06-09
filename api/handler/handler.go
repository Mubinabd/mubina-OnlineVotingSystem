package handler

import (
	"project/config/logger"
	pb "project/genproto"

	"google.golang.org/grpc"
)

type HandlerStruct struct {
	Logger     *logger.Loggerr
	Candidate  pb.CandidateServiceClient
	Election   pb.ElectionServiceClient
	PublicVote pb.PublicVoteServiceClient
	Public     pb.PublicServiceClient
	Party      pb.PartyServiceClient
}

func NewHandlerStruct(connClient *grpc.ClientConn) *HandlerStruct {
	return &HandlerStruct{
		Election:   pb.NewElectionServiceClient(connClient),
		Candidate:  pb.NewCandidateServiceClient(connClient),
		PublicVote: pb.NewPublicVoteServiceClient(connClient),
		Public:     pb.NewPublicServiceClient(connClient),
		Party:      pb.NewPartyServiceClient(connClient),
	}
}
