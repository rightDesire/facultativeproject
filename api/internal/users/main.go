package users

import (
	"context"

	"github.com/google/uuid"
	dm "github.com/rightDesire/facultativeproject/common/domain"
)

type UserService struct {
	client GRPCUserClient
}

func (s *UserService) CreateUser(ctx context.Context, domain dm.User) (uuid.UUID, error) {
	return s.client.CreateUser(ctx, domain)
}
