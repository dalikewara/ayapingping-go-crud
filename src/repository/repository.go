package repository

import (
	"github.com/dalikewara/pgxpoolgo"
)

type User interface {
	// FindAllActive finds all active user from database.
	FindAllActive(param UserFindAllActiveParam) UserFindAllActiveResult

	// FindDetailByID finds user detail by id from database.
	FindDetailByID(param UserFindDetailByIDParam) UserFindDetailByIDResult

	// FindByUsernameOrEmailAndPassword finds user data by username or email, and password from database.
	FindByUsernameOrEmailAndPassword(param UserFindByUsernameOrEmailAndPasswordParam) UserFindByUsernameOrEmailAndPasswordResult

	// InsertTx inserts new user data into database.
	// It will insert user and profile data using transaction.
	InsertTx(param UserInsertTxParam) UserInsertTxResult

	// UpdateByIDTx updates user data by id into database.
	// It will update user and profile data using transaction.
	UpdateByIDTx(param UserUpdateByIDTxParam) UserUpdateByIDTxResult

	// DeleteByIDTx deletes user data by id from database.
	// It will delete user and profile data using transaction.
	DeleteByIDTx(param UserDeleteByIDTxParam) UserDeleteByIDTxResult
}

type Profile interface {
	// UpdateImageByUserID updates profile image by user id into database.
	UpdateImageByUserID(param ProfileUpdateImageByUserIDParam) ProfileUpdateImageByUserIDResult
}

type Repository struct {
	UserPostgreSQL    User
	ProfilePostgreSQL Profile
}

type NewParam struct {
	TimezoneName   string
	TimezoneOffset int
	PostgreSQLPool pgxpoolgo.Pool
}

// New generates new repository.
func New(param NewParam) *Repository {
	return &Repository{
		UserPostgreSQL: NewUserPostgreSQL(NewUserPostgreSQLParam{
			TimezoneName:   param.TimezoneName,
			TimezoneOffset: param.TimezoneOffset,
			Pool:           param.PostgreSQLPool,
		}),
		ProfilePostgreSQL: NewProfilePostgreSQL(NewProfilePostgreSQLParam{
			Pool: param.PostgreSQLPool,
		}),
	}
}
