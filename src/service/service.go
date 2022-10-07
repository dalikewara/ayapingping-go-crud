package service

import "github.com/dalikewara/ayapingping-go-crud/src/repository"

type User interface {
	// Register registers new user data.
	Register(param UserRegisterParam) UserRegisterResult
	
	// Login logins user based on their account credentials.
	Login(param UserLoginParam) UserLoginResult
}

type Service struct {
	User User
}

type NewParam struct {
	Repo *repository.Repository
}

// New generates new service.
func New(param NewParam) *Service {
	userService := NewUser(NewUserParam{
		UserRepo: param.Repo.UserPostgreSQL,
	})

	return &Service{
		User: userService,
	}
}
