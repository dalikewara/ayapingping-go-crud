package repository

import (
	"github.com/dalikewara/pgxpoolgo"
)

type User interface {
	// InsertTx inserts new user data into database.
	// InsertTx will insert user and profile data using transaction.
	InsertTx(param UserInsertTxParam) UserInsertTxResult

	// FindByUsernameOrEmailAndPassword finds user data by username or email, and password from database.
	FindByUsernameOrEmailAndPassword(param UserFindByUsernameOrEmailAndPasswordParam) UserFindByUsernameOrEmailAndPasswordResult
}

type Repository struct {
	UserPostgreSQL User
}

type NewParam struct {
	TimezoneName   string
	TimezoneOffset int
	PostgreSQLPool pgxpoolgo.Pool
}

// New generates new repository.
func New(param NewParam) *Repository {
	userPostgreSQLRepo := NewUserPostgreSQL(NewUserPostgreSQLParam{
		TimezoneName:   param.TimezoneName,
		TimezoneOffset: param.TimezoneOffset,
		Pool:           param.PostgreSQLPool,
	})

	return &Repository{
		UserPostgreSQL: userPostgreSQLRepo,
	}
}
