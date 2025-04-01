package users

import (
	pb "github.com/rightDesire/facultativeproject/proto/users"
)

type UserRepositoryServer struct {
	pb.UnimplementedUserRepositoryServer
}
