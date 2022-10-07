package service

import (
	"context"
	"github.com/dalikewara/ayapingping-go-crud/src/entity"
)

// User

type UserRegisterParam struct {
	Ctx                  context.Context
	Username             entity.Username
	Email                entity.Email
	Password             entity.Password
	PasswordConfirmation entity.Password
	FirstName            entity.FirstName
	LastName             entity.LastName
	Gender               entity.Gender
}

type UserRegisterResult struct {
	UserID entity.ID
	Error  entity.StdError
}

type UserLoginParam struct {
	Ctx             context.Context
	UsernameOrEmail entity.UsernameOrEmail
	Password        entity.Password
}

type UserLoginResult struct {
	User  *entity.User
	Error entity.StdError
}
