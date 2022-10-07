package service

import (
	"context"
	"github.com/dalikewara/ayapingping-go-crud/src/entity"
)

// User

type UserGetAllActiveParam struct {
	Ctx context.Context
}

type UserGetAllActiveResult struct {
	User  []*entity.User
	Error entity.StdError
}

type UserGetDetailParam struct {
	Ctx context.Context
	ID  entity.ID
}

type UserGetDetailResult struct {
	User  *entity.UserWithProfile
	Error entity.StdError
}

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
	ID        entity.ID
	ProfileID entity.ID
	Error     entity.StdError
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

type UserUpdateParam struct {
	Ctx       context.Context
	ID        entity.ID
	Username  entity.Username
	FirstName entity.FirstName
	LastName  entity.LastName
	Gender    entity.Gender
}

type UserUpdateResult struct {
	Error entity.StdError
}

type UserDeleteParam struct {
	Ctx      context.Context
	ID       entity.ID
	Password entity.Password
}

type UserDeleteResult struct {
	Error entity.StdError
}

// Profile

type ProfileUpdateImageParam struct {
	Ctx    context.Context
	UserID entity.ID
	Image  entity.Image
}

type ProfileUpdateImageResult struct {
	Error entity.StdError
}
