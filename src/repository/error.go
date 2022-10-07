package repository

import (
	"github.com/dalikewara/ayapingping-go-crud/src/entity"
	"net/http"
)

// Database: User

var ErrDatabaseUserInsert = entity.NewStdError("DB-USER-01", "error inserting user data into database", http.StatusInternalServerError)
var ErrDatabaseUserInsertNoAffected = entity.NewStdError("DB-USER-02", "no user data was inserted into database", http.StatusInternalServerError)
var ErrDatabaseUserFind = entity.NewStdError("DB-USER-03", "error finding user data from database", http.StatusInternalServerError)

// Database: Profile

var ErrDatabaseProfileInsert = entity.NewStdError("DB-PROFILE-01", "error inserting profile data into database", http.StatusInternalServerError)
var ErrDatabaseProfileInsertNoAffected = entity.NewStdError("DB-PROFILE-02", "no profile data was inserted into database", http.StatusInternalServerError)
