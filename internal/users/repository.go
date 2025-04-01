package users

import (
	"github.com/google/uuid"
	dm "github.com/rightDesire/facultativeproject/domain"
	"github.com/rightDesire/facultativeproject/pkg/postgres"
)

type UserRepository struct {
	gorm *postgres.GDB
}

func NewRepository(db *postgres.GDB) IUserRepository {
	return &UserRepository{
		gorm: db,
	}
}

func (r *UserRepository) CreateUser(dmUser dm.User) (uid uuid.UUID, err error) {
	var orm User
	orm = User{
		Email:    dmUser.Email,
		Password: dmUser.Password,
		FullName: dmUser.FullName,
	}
	res := r.gorm.DB.Create(&orm)
	if err := res.Error; err != nil {
		return uid, err
	}
	uid = orm.UUID
	return uid, nil
}

func (r *UserRepository) LoginUser() {
	panic("Implemented me")
}

func (r *UserRepository) GetUserUserUUID() {
	panic("Implemented me")
}

func (r *UserRepository) PutUserUserUUID() {
	panic("Implemented me")
}
