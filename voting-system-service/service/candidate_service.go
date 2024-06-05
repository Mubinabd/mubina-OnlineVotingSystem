package service

import (
	"context"
	pb "project/genproto/public"
)

type CandidateService struct {
	stg pb.CandidateServiceClient
	pb.UnimplementedCandidateServiceServer
}

func NewCandidateService(stg pb.CandidateServiceClient) *CandidateService {
	return &CandidateService{stg: stg}
}

func (cs *CandidateService) Create(ctx context.Context, req *pb.CreateCandidateReq) (*pb.Void, error) {
	_,err := cs.stg.Create(ctx,req)
	if err != nil {
		return nil, err
	}
	return &pb.Void{}, nil
}

func (cs *CandidateService) GetById(ctx context.Context, req *pb.GetCandidate) (*pb.CandidateRes, error) {
	res,err := cs.stg.Get(ctx,req)
	if err!= nil {
		return nil, err
	}
	return res,nil
}
func (cs *CandidateService) GetAll(ctx context.Context, req *pb.Filter) (*pb.CandidatiesGetAllResp, error) {
	res,err := cs.stg.GetAll(ctx,req)
	if err!= nil {
		return nil, err
	}
	return res,nil
}

func (cs *CandidateService) Delete(ctx context.Context, req *pb.GetByIdReq) (*pb.Void, error ) {
	_,err := cs.stg.Delete(ctx,req)
    if err != nil {
        return nil, err
    }
    return &pb.Void{},nil
}