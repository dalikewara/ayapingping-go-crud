package rest

import (
	"github.com/dalikewara/ayapingping-go-crud/src/entity"
)

type Response struct {
	Code    string      `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// User

type UserGetAllActiveResponseData struct {
	Rows *entity.UsersForClient `json:"rows"`
}

type UserGetDetailResponseData entity.UserWithProfileForClient

type UserRegisterRequest struct {
	Username             entity.Username  `json:"username" binding:"required"`
	Email                entity.Email     `json:"email" binding:"required"`
	Password             entity.Password  `json:"password" binding:"required"`
	PasswordConfirmation entity.Password  `json:"password_confirmation" binding:"required"`
	FirstName            entity.FirstName `json:"first_name"`
	LastName             entity.LastName  `json:"last_name"`
	Gender               entity.Gender    `json:"gender"`
}

type UserRegisterResponseData struct {
	ID entity.ID `json:"id"`
}

type UserLoginRequest struct {
	UsernameOrEmail entity.UsernameOrEmail `json:"username_or_email" binding:"required"`
	Password        entity.Password        `json:"password" binding:"required"`
}

type UserLoginResponseData entity.UserForClient

type UserUpdateRequest struct {
	ID        entity.ID        `json:"id" binding:"required"`
	Username  entity.Username  `json:"username"`
	FirstName entity.FirstName `json:"first_name"`
	LastName  entity.LastName  `json:"last_name"`
	Gender    entity.Gender    `json:"gender"`
}

type UserDeleteRequest struct {
	ID       entity.ID       `json:"id" binding:"required"`
	Password entity.Password `json:"password" binding:"required"`
}

// Profile

type ProfileUpdateImageRequest struct {
	UserID entity.ID    `json:"user_id" binding:"required"`
	Image  entity.Image `json:"image"`
}
