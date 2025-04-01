package users

import (
	"context"

	"github.com/google/uuid"
	dm "github.com/rightDesire/facultativeproject/domain"
	pb "github.com/rightDesire/facultativeproject/proto/users"
)

type UserService struct {
	rpcService pb.UserServiceServer
}

func (s *UserService) CreateUser(ctx context.Context, domain dm.User) (uuid.UUID, error) {
	var uid uuid.UUID
	pbUser := &pb.User{
		Email:    domain.Email,
		Password: domain.Password,
		FullName: domain.FullName,
	}
	proto := &pb.CreateUserRequest{
		User: pbUser,
	}
	res, err := s.rpcService.CreateUser(ctx, proto)
	if err != nil {
		return uid, err
	}
	uid, err = uuid.Parse(res.UserUUID)
	if err != nil {
		return uid, err
	}
	return uid, nil
}
