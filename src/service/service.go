package service

import "github.com/dalikewara/ayapingping-go-crud/src/repository"

type User interface {
	// GetAllActive gets all active users.
	GetAllActive(param UserGetAllActiveParam) UserGetAllActiveResult

	// GetDetail gets user detail.
	GetDetail(param UserGetDetailParam) UserGetDetailResult

	// Register registers new user data.
	Register(param UserRegisterParam) UserRegisterResult

	// Login logins user based on their account credentials.
	Login(param UserLoginParam) UserLoginResult

	// Update updates user data.
	Update(param UserUpdateParam) UserUpdateResult

	// Delete deletes user data.
	Delete(param UserDeleteParam) UserDeleteResult
}

type Profile interface {
	// UpdateImage updates profile image.
	UpdateImage(param ProfileUpdateImageParam) ProfileUpdateImageResult
}

type Service struct {
	User    User
	Profile Profile
}

type NewParam struct {
	Repo *repository.Repository
}

// New generates new service.
func New(param NewParam) *Service {
	return &Service{
		User: NewUser(NewUserParam{
			UserRepo: param.Repo.UserPostgreSQL,
		}),
		Profile: NewProfile(NewProfileParam{
			ProfileRepo: param.Repo.ProfilePostgreSQL,
		}),
	}
}
