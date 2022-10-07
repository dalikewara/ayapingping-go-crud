package service

import (
	"github.com/dalikewara/ayapingping-go-crud/src/repository"
)

type user struct {
	userRepo repository.User
}

type NewUserParam struct {
	UserRepo repository.User
}

// NewUser generates new user service that implements User.
func NewUser(param NewUserParam) User {
	return &user{
		userRepo: param.UserRepo,
	}
}

// Register registers new user data.
func (s *user) Register(param UserRegisterParam) UserRegisterResult {
	result := UserRegisterResult{}

	if param.Username.Validate() != nil {
		result.Error = ErrParamUsername
		return result
	}
	if param.Email.Validate() != nil {
		result.Error = ErrParamEmail
		return result
	}
	if param.Password.Validate() != nil {
		result.Error = ErrParamPassword
		return result
	}
	if param.Password != param.PasswordConfirmation {
		result.Error = ErrParamPasswordConfirmationDoesntMatch
		return result
	}
	if param.FirstName.Validate() != nil {
		result.Error = ErrParamFirstName
		return result
	}
	if !param.LastName.IsEmpty() && param.LastName.Validate() != nil {
		result.Error = ErrParamLastName
		return result
	}

	userInsert := s.userRepo.InsertTx(repository.UserInsertTxParam{
		Ctx:       param.Ctx,
		Username:  param.Username,
		Email:     param.Email,
		Password:  param.Password,
		FirstName: param.FirstName,
		LastName:  param.LastName,
		Gender:    param.Gender,
	})
	if userInsert.Error != nil {
		if userInsert.IsUserDuplicateKey {
			result.Error = ErrServiceUsernameOrEmailAlreadyExists
		} else if userInsert.IsProfileDuplicateKey {
			result.Error = ErrServiceProfileAlreadyExists
		} else {
			result.Error = userInsert.Error
		}
		return result
	}

	result.UserID = userInsert.ID

	return result
}

// Login logins user based on their account credentials.
func (s *user) Login(param UserLoginParam) UserLoginResult {
	result := UserLoginResult{}

	if param.UsernameOrEmail.Validate() != nil {
		result.Error = ErrParamUsernameOrEmail
		return result
	}
	if param.Password.Validate() != nil {
		result.Error = ErrParamPassword
		return result
	}

	userLogin := s.userRepo.FindByUsernameOrEmailAndPassword(repository.UserFindByUsernameOrEmailAndPasswordParam{
		Ctx:             param.Ctx,
		UsernameOrEmail: param.UsernameOrEmail,
		Password:        param.Password,
	})
	if userLogin.Error != nil {
		result.Error = userLogin.Error
		return result
	}
	if userLogin.User == nil {
		result.Error = ErrServiceUsernameOrPasswordWrong
		return result
	}

	result.User = userLogin.User

	return result
}
