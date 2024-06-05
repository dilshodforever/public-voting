package service

import (
	"context"
	"log"
	pb "root/genprotos"
	s "root/storage"

)

type PartyService struct {
	stg s.InitRoot
	pb.UnimplementedPartyServiceServer
}

func NewPartyService(stg s.InitRoot) *PartyService {
	return &PartyService{stg: stg}
}

func (c *PartyService) CreateParty(ctx context.Context, Party *pb.Party) (*pb.Void, error) {
	pb, err := c.stg.Party().CreateParty(Party)
	if err != nil {
		log.Print(err)
	}
	return pb, err
}

func (c *PartyService) GetAllPartys(ctx context.Context, pb *pb.Void) (*pb.GetAllParty, error) {
	Partys, err := c.stg.Party().GetAllParty(pb)
	if err != nil {
		log.Print(err)
	}

	return Partys, err
}

func (c *PartyService) GetByIdParty(ctx context.Context, id *pb.ById) (*pb.Party, error) {
	prod, err := c.stg.Party().GetByIdParty(id)
	if err != nil {
		log.Print(err)
	}

	return prod, err
}

func (c *PartyService) UpdateParty(ctx context.Context, Party *pb.Party) (*pb.Void, error) {
	pb, err := c.stg.Party().UpdateParty(Party)
	if err != nil {
		log.Print(err)
	}

	return pb, err
}

func (c *PartyService) DeleteParty(ctx context.Context, id *pb.ById) (*pb.Void, error) {
	pb, err := c.stg.Party().DeleteParty(id)
	if err != nil {
		log.Print(err)
	}

	return pb, err
}
