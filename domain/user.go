package domain

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/rightDesire/facultativeproject/internal/helpers"
)

type User struct {
	UUID      uuid.UUID `validate:"required"`
	Email     string    `validate:"required,email"`
	Password  string    `validate:"required"`
	FullName  string    `validate:"required"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

func NewUser(email string, password string, fullName string) *User {
	le := &User{
		UUID:     uuid.New(),
		Email:    email,
		Password: password,
		FullName: fullName,
	}

	errs, ok := helpers.ValidationStruct(le)
	if !ok {
		panic(errors.New(helpers.Join(errs, ", ")))
	}

	return le
}
