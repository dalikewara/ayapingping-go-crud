package repository

import (
	"context"
	"github.com/dalikewara/ayapingping-go-crud/src/entity"
)

// User

type UserInsertTxParam struct {
	Ctx       context.Context
	Username  entity.Username
	Email     entity.Email
	Password  entity.Password
	FirstName entity.FirstName
	LastName  entity.LastName
	Gender    entity.Gender
}

type UserInsertTxResult struct {
	ID                    entity.ID
	ProfileID             entity.ID
	Error                 entity.StdError
	IsUserDuplicateKey    bool
	IsProfileDuplicateKey bool
}

type UserFindByUsernameOrEmailAndPasswordParam struct {
	Ctx             context.Context
	UsernameOrEmail entity.UsernameOrEmail
	Password        entity.Password
}

type UserFindByUsernameOrEmailAndPasswordResult struct {
	User  *entity.User
	Error entity.StdError
}
