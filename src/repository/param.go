package repository

import (
	"context"
	"github.com/dalikewara/ayapingping-go-crud/src/entity"
)

// User

type UserFindAllActiveParam struct {
	Ctx context.Context
}

type UserFindAllActiveResult struct {
	User  []*entity.User
	Error entity.StdError
}

type UserFindDetailByIDParam struct {
	Ctx context.Context
	ID  entity.ID
}

type UserFindDetailByIDResult struct {
	User  *entity.UserWithProfile
	Error entity.StdError
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

type UserUpdateByIDTxParam struct {
	Ctx       context.Context
	ID        entity.ID
	Username  entity.Username
	FirstName entity.FirstName
	LastName  entity.LastName
	Gender    entity.Gender
}

type UserUpdateByIDTxResult struct {
	Error entity.StdError
}

type UserDeleteByIDTxParam struct {
	Ctx      context.Context
	ID       entity.ID
	Password entity.Password
}

type UserDeleteByIDTxResult struct {
	Error entity.StdError
}

// Profile

type ProfileUpdateImageByUserIDParam struct {
	Ctx    context.Context
	UserID entity.ID
	Image  entity.Image
}

type ProfileUpdateImageByUserIDResult struct {
	Error entity.StdError
}
