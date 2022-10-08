package service

import (
	"github.com/dalikewara/ayapingping-go-crud/src/entity"
	"net/http"
)

// Param: Username

var ErrParamUsername = entity.NewStdError("PRM-USERNAME-01", "invalid username format", http.StatusBadRequest)

// Param: Email

var ErrParamEmail = entity.NewStdError("PRM-EMAIL-01", "invalid email format", http.StatusBadRequest)

// Param: Password

var ErrParamPassword = entity.NewStdError("PRM-PASSWORD-01", "invalid password format", http.StatusBadRequest)
var ErrParamPasswordConfirmationDoesntMatch = entity.NewStdError("PRM-PASSWORD-02", "password confirmation doesn't match", http.StatusBadRequest)

// Param: FirstName

var ErrParamFirstName = entity.NewStdError("PRM-FIRSTNAME-01", "invalid first name format", http.StatusBadRequest)

// Param: LastName

var ErrParamLastName = entity.NewStdError("PRM-LASTNAME-01", "invalid last name format", http.StatusBadRequest)

// Param: Username or email

var ErrParamUsernameOrEmail = entity.NewStdError("PRM-USERNAMEOREMAIL-01", "invalid username or email format", http.StatusBadRequest)

// Param: User id

var ErrParamUserID = entity.NewStdError("PRM-USERID-01", "invalid user id format", http.StatusBadRequest)

// Param: Gender

var ErrParamGender = entity.NewStdError("PRM-GENDER-01", "invalid gender format", http.StatusBadRequest)

// Service: User

var ErrServiceUserUsernameOrEmailAlreadyExists = entity.NewStdError("SVC-USER-01", "username or email already exists", http.StatusConflict)
var ErrServiceUserUsernameOrPasswordWrong = entity.NewStdError("SVC-USER-02", "username/email or password wrong", http.StatusNotFound)
var ErrServiceUserNotFound = entity.NewStdError("SVC-USER-03", "user not found", http.StatusNotFound)
var ErrServiceUserUsernameAlreadyExists = entity.NewStdError("SVC-USER-04", "username already exists", http.StatusConflict)

// Service: Profile

var ErrServiceProfileAlreadyExists = entity.NewStdError("SVC-PROFILE-01", "profile already exists", http.StatusConflict)
