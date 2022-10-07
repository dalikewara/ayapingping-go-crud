package repository

import (
	"github.com/dalikewara/pgxpoolgo"
)

type profilePostgreSQL struct {
	pool pgxpoolgo.Pool
}

type NewProfilePostgreSQLParam struct {
	Pool pgxpoolgo.Pool
}

// NewProfilePostgreSQL generates new profile PostgreSQL repository that implements Profile.
func NewProfilePostgreSQL(param NewProfilePostgreSQLParam) Profile {
	return &profilePostgreSQL{
		pool: param.Pool,
	}
}

// UpdateImageByUserID updates profile image by user id into database.
func (r *profilePostgreSQL) UpdateImageByUserID(param ProfileUpdateImageByUserIDParam) ProfileUpdateImageByUserIDResult {
	result := ProfileUpdateImageByUserIDResult{}

	reply, err := r.pool.Exec(
		param.Ctx,
		PostgreSQLProfileUpdateImageByUserIDQuery,
		param.Image,
		param.UserID,
	)
	if err != nil {
		result.Error = ErrDatabaseProfileUpdate
		return result
	}
	if reply.RowsAffected() < 1 {
		result.Error = ErrDatabaseProfileUpdateNoAffected
		return result
	}

	return result
}
