package service

import (
	"context"
	pb "project/genproto/public"
)

type ElectionService struct {
	stg pb.ElectionServiceClient
	pb.UnimplementedElectionServiceServer
}

func NewElectionService(stg pb.ElectionServiceClient) *ElectionService {
	return &ElectionService{stg: stg}
}

func (es *ElectionService) Create(ctx context.Context, req *pb.CreateElectionReq)(*pb.Void, error ){
	_,err := es.stg.Create(ctx,req)
	if err != nil {
		return nil, err
	}
	return &pb.Void{},nil
}
func (es *ElectionService) GetById(ctx context.Context, req *pb.ElectionReq) (*pb.ElectionReq, error) {
	res,err := es.stg.Get(ctx,req)
	if err!= nil {
		return nil, err
	}
	return res,nil
}
func (es *ElectionService) GetAll(ctx context.Context, req *pb.Filter) (*pb.ElectionsGetAllResp, error) {
	res,err := es.stg.GetAll(ctx,req)
	if err!= nil {
		return nil, err
	}
	return res,nil
}

func (es *ElectionService) UpdateItem(ctx context.Context, elec *pb.ElectionUpdate) (*pb.Void, error) {
	_,err := es.stg.Update(ctx,elec)
	if err != nil {
		return nil, err
	}
	return &pb.Void{}, nil
}

func (es *ElectionService) Delete(ctx context.Context, req *pb.GetByIdReq) (*pb.Void, error ) {
	_,err := es.stg.Delete(ctx,req)
    if err != nil {
        return nil, err
    }
    return &pb.Void{},nil
}