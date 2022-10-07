package repository

import (
	"github.com/dalikewara/ayapingping-go-crud/src/entity"
	"github.com/dalikewara/pgxpoolgo"
	"github.com/jackc/pgtype"
	"github.com/jackc/pgx/v4"
)

type userPostgreSQL struct {
	timezoneName   string
	timezoneOffset int
	typeTimestamp  pgtype.Timestamp
	pool           pgxpoolgo.Pool
}

type NewUserPostgreSQLParam struct {
	TimezoneName   string
	TimezoneOffset int
	Pool           pgxpoolgo.Pool
}

// NewUserPostgreSQL generates new user PostgreSQL repository that implements User.
func NewUserPostgreSQL(param NewUserPostgreSQLParam) User {
	return &userPostgreSQL{
		timezoneName:   param.TimezoneName,
		timezoneOffset: param.TimezoneOffset,
		typeTimestamp:  pgtype.Timestamp{},
		pool:           param.Pool,
	}
}

// FindAllActive finds all active user from database.
func (r *userPostgreSQL) FindAllActive(param UserFindAllActiveParam) UserFindAllActiveResult {
	result := UserFindAllActiveResult{}
	return result
}

// FindDetailByID finds user detail by id from database.
func (r *userPostgreSQL) FindDetailByID(param UserFindDetailByIDParam) UserFindDetailByIDResult {
	result := UserFindDetailByIDResult{}
	return result
}

// FindByUsernameOrEmailAndPassword finds user data by username or email, and password from PostgreSQL database.
func (r *userPostgreSQL) FindByUsernameOrEmailAndPassword(param UserFindByUsernameOrEmailAndPasswordParam) UserFindByUsernameOrEmailAndPasswordResult {
	result := UserFindByUsernameOrEmailAndPasswordResult{}
	user := &entity.User{}

	createdAt := r.typeTimestamp
	updatedAt := r.typeTimestamp
	deletedAt := r.typeTimestamp

	err := r.pool.QueryRow(param.Ctx, PostgreSQLUserFindByUsernameOrEmailAndPasswordQuery, param.UsernameOrEmail, param.UsernameOrEmail, param.Password).
		Scan(
			&user.ID,
			&user.Username,
			&user.Email,
			&user.ActiveStatus,
			&createdAt,
			&updatedAt,
			&deletedAt,
		)
	if err != nil {
		if pgxpoolgo.ErrDB(err).IsNoRows() {
			return result
		}
		result.Error = ErrDatabaseUserFind
		return result
	}

	user.CreatedAt.SetFromTime(createdAt.Time, entity.Timezone{
		Name:   r.timezoneName,
		Offset: r.timezoneOffset,
	})
	user.UpdatedAt.SetFromTime(updatedAt.Time, entity.Timezone{
		Name:   r.timezoneName,
		Offset: r.timezoneOffset,
	})
	user.DeletedAt.SetFromTime(deletedAt.Time, entity.Timezone{
		Name:   r.timezoneName,
		Offset: r.timezoneOffset,
	})

	result.User = user

	return result
}

// InsertTx inserts new user data into PostgreSQL database.
// It will insert user and profile data using transaction.
func (r *userPostgreSQL) InsertTx(param UserInsertTxParam) UserInsertTxResult {
	result := UserInsertTxResult{}

	tx, err := r.pool.BeginTx(param.Ctx, pgx.TxOptions{})
	if err != nil {
		result.Error = ErrDatabaseUserInsert
		return result
	}

	user := &entity.User{}
	err = tx.QueryRow(param.Ctx, PostgreSQLUserInsertQuery, param.Username, param.Email, param.Password).Scan(&user.ID)
	if err != nil {
		if errR := tx.Rollback(param.Ctx); errR != nil {
			result.Error = ErrDatabaseUserInsert
			return result
		}
		e := pgxpoolgo.ErrDB(err)
		if e.IsNoRows() {
			result.Error = ErrDatabaseUserInsertNoAffected
		} else {
			result.Error = ErrDatabaseUserInsert
		}
		result.IsUserDuplicateKey = e.IsDuplicateKey()
		return result
	}

	profile := &entity.Profile{}
	err = tx.QueryRow(param.Ctx, PostgreSQLProfileInsertQuery, user.ID, param.FirstName, param.LastName, param.Gender).Scan(&profile.ID)
	if err != nil {
		if errR := tx.Rollback(param.Ctx); errR != nil {
			result.Error = ErrDatabaseProfileInsert
			return result
		}
		e := pgxpoolgo.ErrDB(err)
		if e.IsNoRows() {
			result.Error = ErrDatabaseProfileInsertNoAffected
		} else {
			result.Error = ErrDatabaseProfileInsert
		}
		result.IsProfileDuplicateKey = e.IsDuplicateKey()
		return result
	}

	if err = tx.Commit(param.Ctx); err != nil {
		result.Error = ErrDatabaseUserInsert
		return result
	}

	result.ID = user.ID
	result.ProfileID = profile.ID

	return result
}

// UpdateByIDTx updates user data by id into database.
// It will update user and profile data using transaction.
func (r *userPostgreSQL) UpdateByIDTx(param UserUpdateByIDTxParam) UserUpdateByIDTxResult {
	result := UserUpdateByIDTxResult{}
	return result
}

// DeleteByIDTx deletes user data by id from database.
// It will delete user and profile data using transaction.
func (r *userPostgreSQL) DeleteByIDTx(param UserDeleteByIDTxParam) UserDeleteByIDTxResult {
	result := UserDeleteByIDTxResult{}
	return result
}
