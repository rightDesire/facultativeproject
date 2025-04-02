package users

import (
	"context"

	"github.com/google/wire"
	dm "github.com/rightDesire/facultativeproject/common/domain"
	"github.com/rightDesire/facultativeproject/common/helpers"
	pb "github.com/rightDesire/facultativeproject/common/proto/generated"
)

type UserServiceServer struct {
	pb.UnimplementedUserServiceServer
	repo IUserRepository
}

func NewUserServiceServer(r IUserRepository) *UserServiceServer {
	return &UserServiceServer{
		repo: r,
	}
}

func (s *UserServiceServer) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	var domain dm.User
	domain = dm.User{
		Email:    req.GetUser().GetEmail(),
		Password: helpers.Hash(req.GetUser().GetPassword()),
		FullName: req.GetUser().GetFullName(),
	}
	uid, err := s.repo.CreateUser(domain)
	if err != nil {
		return nil, err
	}
	res := &pb.CreateUserResponse{
		UserUUID: uid.String(),
	}
	return res, nil
}

func (s *UserServiceServer) LoginUser(ctx context.Context, req *pb.LoginUserRequest) (res *pb.LoginUserResponse, err error) {
	panic("Implemented me")
}

func (s *UserServiceServer) GetUserUserUUID(ctx context.Context, req *pb.GetUserUserUUIDRequest) (res *pb.GetUserUserUUIDResponse, err error) {
	panic("Implemented me")
}

func (s *UserServiceServer) PutUserUserUUID(cxt context.Context, req *pb.PutUserUserUUIDRequest) (res *pb.PutUserUserUUIDResponse, err error) {
	panic("Implemented me")
}

var ServiceProviderSet = wire.NewSet(NewUserServiceServer)
