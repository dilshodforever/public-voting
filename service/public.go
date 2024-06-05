package service

import (
	"context"
	"log"
	pb "root/genprotos"
	s "root/storage"
)

type PublicService struct {
	stg s.InitRoot
	pb.UnimplementedPublicServiceServer
}

func NewPublicService(stg s.InitRoot) *PublicService {
	return &PublicService{stg: stg}
}

func (c *PublicService) CreatePublic(ctx context.Context, public *pb.Public) (*pb.Void, error) {
	pb, err := c.stg.Public().CreatePublic(public)
	if err != nil {
		log.Print(err)
	}
	return pb, err
}

func (c *PublicService) GetAllPublics(ctx context.Context, pb *pb.Void) (*pb.GetAllPublic, error) {
	publics, err := c.stg.Public().GetAllPublic(pb)
	if err != nil {
		log.Print(err)
	}

	return publics, err
}

func (c *PublicService) GetByIdPublic(ctx context.Context, id *pb.ById) (*pb.Public, error) {
	prod, err := c.stg.Public().GetByIdPublic(id)
	if err != nil {
		log.Print(err)
	}

	return prod, err
}

func (c *PublicService) UpdatePublic(ctx context.Context, public *pb.Public) (*pb.Void, error) {
	pb, err := c.stg.Public().UpdatePublic(public)
	if err != nil {
		log.Print(err)
	}

	return pb, err
}

func (c *PublicService) DeletePublic(ctx context.Context, id *pb.ById) (*pb.Void, error) {
	pb, err := c.stg.Public().DeletePublic(id)
	if err != nil {
		log.Print(err)
	}

	return pb, err
}
