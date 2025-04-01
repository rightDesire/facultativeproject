package users

import (
	"context"

	"github.com/google/uuid"
	dm "github.com/rightDesire/facultativeproject/domain"
	pb "github.com/rightDesire/facultativeproject/proto/users"
)

type GRPCUserClient struct {
	client pb.UserServiceClient
}

func NewGRPCUserClient(client pb.UserServiceClient) *GRPCUserClient {
	return &GRPCUserClient{client: client}
}

func (r *GRPCUserClient) CreateUser(ctx context.Context, user dm.User) (uuid.UUID, error) {
	pbUser := &pb.User{
		Email:    user.Email,
		Password: user.Password,
		FullName: user.FullName,
	}
	req := &pb.CreateUserRequest{
		User: pbUser,
	}
	res, err := r.client.CreateUser(ctx, req)
	if err != nil {
		return uuid.Nil, err
	}
	uid, err := uuid.Parse(res.UserUUID)
	if err != nil {
		return uuid.Nil, err
	}
	return uid, nil
}
