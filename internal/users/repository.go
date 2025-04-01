package users

import (
	"context"

	pb "github.com/rightDesire/facultativeproject/proto/users"
)

type UserRepositoryServer struct {
	pb.UnimplementedUserRepositoryServer
}

func (r *UserRepositoryServer) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (res *pb.CreateUserResponse, err error) {
	panic("Implemented me")
}

func (r *UserRepositoryServer) LoginUser(ctx context.Context, req *pb.LoginUserRequest) (res *pb.LoginUserResponse, err error) {
	panic("Implemented me")
}

func (r *UserRepositoryServer) GetUserUserUUID(ctx context.Context, req *pb.GetUserUserUUIDRequest) (res *pb.GetUserUserUUIDResponse, err error) {
	panic("Implemented me")
}

func (r *UserRepositoryServer) PutUserUserUUID(cxt context.Context, req *pb.PutUserUserUUIDRequest) (res *pb.PutUserUserUUIDResponse, err error) {
	panic("Implemented me")
}
