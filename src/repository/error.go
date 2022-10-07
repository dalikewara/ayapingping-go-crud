package repository

import (
	"github.com/dalikewara/ayapingping-go-crud/src/entity"
	"net/http"
)

// Database: User

var ErrDatabaseUserInsert = entity.NewStdError("DB-USER-01", "error inserting user data into database", http.StatusInternalServerError)
var ErrDatabaseUserInsertNoAffected = entity.NewStdError("DB-USER-02", "no user data was inserted into database", http.StatusInternalServerError)
var ErrDatabaseUserFind = entity.NewStdError("DB-USER-03", "error finding user data from database", http.StatusInternalServerError)
var ErrDatabaseUserUpdate = entity.NewStdError("DB-USER-04", "error updating user data into database", http.StatusInternalServerError)
var ErrDatabaseUserUpdateNoAffected = entity.NewStdError("DB-USER-05", "no user data was updated into database", http.StatusInternalServerError)
var ErrDatabaseUserDelete = entity.NewStdError("DB-USER-06", "error deleting user data from database", http.StatusInternalServerError)
var ErrDatabaseUserDeleteNoAffected = entity.NewStdError("DB-USER-07", "no user data was deleted from database", http.StatusInternalServerError)

// Database: Profile

var ErrDatabaseProfileInsert = entity.NewStdError("DB-PROFILE-01", "error inserting profile data into database", http.StatusInternalServerError)
var ErrDatabaseProfileInsertNoAffected = entity.NewStdError("DB-PROFILE-02", "no profile data was inserted into database", http.StatusInternalServerError)
var ErrDatabaseProfileUpdate = entity.NewStdError("DB-PROFILE-03", "error updating profile data into database", http.StatusInternalServerError)
var ErrDatabaseProfileUpdateNoAffected = entity.NewStdError("DB-PROFILE-04", "no profile data was updated into database", http.StatusInternalServerError)
