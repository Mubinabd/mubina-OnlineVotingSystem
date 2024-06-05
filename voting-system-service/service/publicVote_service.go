package service

import (
	"context"
	pb "project/genproto/public"
)

type PublicVoteService struct {
	stg pb.PublicVoteServiceClient
	pb.UnimplementedPublicVoteServiceServer
}

func NewPublicVoteService(stg pb.PublicVoteServiceClient) *PublicVoteService {
	return &PublicVoteService{stg: stg}
}

func (pvs *PublicVoteService)Create(ctx context.Context, req *pb.CreatePublicVoteReq)(*pb.Void, error ){
	_,err := pvs.stg.Create(ctx,req)
	if err != nil {
		return nil, err
	}
	return &pb.Void{},nil
}
func (pvs *PublicVoteService) GetById(ctx context.Context, req *pb.PublicVoteReq) (*pb.PublicVote, error) {
	res,err := pvs.stg.Get(ctx,req)
	if err!= nil {
		return nil, err
	}
	return res,nil
}
func (pvs *PublicVoteService) GetAll(ctx context.Context, req *pb.Filter) (*pb.PublicVotesGetAllResp, error) {
	res,err := pvs.stg.GetAll(ctx,req)
	if err!= nil {
		return nil, err
	}
	return res,nil
}

func (pvs *PublicVoteService) Delete(ctx context.Context, req *pb.GetByIdReq) (*pb.Void, error ) {
	_,err := pvs.stg.Delete(ctx,req)
    if err != nil {
        return nil, err
    }
    return &pb.Void{},nil
}