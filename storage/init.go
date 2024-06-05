package postgres

import (
	pb "root/genprotos"
)

type InitRoot interface {
	Party() Party
	Public() Public
}
type Party interface {
	CreateParty(party *pb.Party) (*pb.Void, error)
	GetByIdParty(id *pb.ById) (*pb.Party, error)
	GetAllParty(_ *pb.Void) (*pb.GetAllParty, error)
	UpdateParty(party *pb.Party) (*pb.Void, error)
	DeleteParty(id *pb.ById) (*pb.Void, error)
}

type Public interface {
	CreatePublic(pub *pb.Public) (*pb.Void, error)
	GetByIdPublic(id *pb.ById) (*pb.Public, error)
	GetAllPublic(_ *pb.Void) (*pb.GetAllPublic, error)
	UpdatePublic(pub *pb.Public) (*pb.Void, error)
	DeletePublic(id *pb.ById) (*pb.Void, error)
}
