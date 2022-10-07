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

type UserRegisterRequest struct {
	Username             entity.Username  `json:"username" binding:"required"`
	Email                entity.Email     `json:"email" binding:"required"`
	Password             entity.Password  `json:"password" binding:"required"`
	PasswordConfirmation entity.Password  `json:"password_confirmation" binding:"required"`
	FirstName            entity.FirstName `json:"first_name" binding:"required"`
	LastName             entity.LastName  `json:"last_name"`
	Gender               entity.Gender    `json:"gender"`
}

type UserRegisterResponseData struct {
	UserID entity.ID `json:"user_id"`
}

type UserLoginRequest struct {
	UsernameOrEmail entity.UsernameOrEmail `json:"username_or_email" binding:"required"`
	Password        entity.Password        `json:"password" binding:"required"`
}

type UserLoginResponseData struct {
	ID           entity.ID           `json:"id"`
	Username     entity.Username     `json:"username"`
	Email        entity.Email        `json:"email"`
	ActiveStatus entity.ActiveStatus `json:"active_status"`
	CreatedAt    entity.Time         `json:"created_at"`
	UpdatedAt    entity.Time         `json:"updated_at"`
}
