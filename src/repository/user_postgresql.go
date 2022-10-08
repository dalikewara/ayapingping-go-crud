package repository

import (
	"github.com/dalikewara/ayapingping-go-crud/src/entity"
	"github.com/dalikewara/pgxpoolgo"
	"github.com/jackc/pgx/v4"
)

type userPostgreSQL struct {
	timezoneName   string
	timezoneOffset int
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
		pool:           param.Pool,
	}
}

// FindAllActive finds all active user from PostgreSQL database.
func (r *userPostgreSQL) FindAllActive(param UserFindAllActiveParam) UserFindAllActiveResult {
	result := UserFindAllActiveResult{}
	users := entity.Users{}

	rows, err := r.pool.Query(param.Ctx, PostgreSQLUserFindAllActiveQuery)
	if err != nil {
		if pgxpoolgo.ErrDB(err).IsNoRows() {
			return result
		}
		result.Error = ErrDatabaseUserFind
		return result
	}
	defer rows.Close()

	var exists bool

	for rows.Next() {
		user := &entity.User{}
		createdAt := user.CreatedAt.GetPostgreSQLType()
		updatedAt := user.UpdatedAt.GetPostgreSQLType()
		deletedAt := user.DeletedAt.GetPostgreSQLType()

		if err = rows.Scan(
			&user.ID,
			&user.Username,
			&user.Email,
			&user.ActiveStatus,
			&createdAt,
			&updatedAt,
			&deletedAt,
		); err != nil {
			result.Error = ErrDatabaseUserFind
			return result
		}

		user.CreatedAt.SetFromTime(user.CreatedAt.PrimitiveFromPostgreSQLType(createdAt), entity.Timezone{
			Name:   r.timezoneName,
			Offset: r.timezoneOffset,
		})
		user.UpdatedAt.SetFromTime(user.UpdatedAt.PrimitiveFromPostgreSQLType(updatedAt), entity.Timezone{
			Name:   r.timezoneName,
			Offset: r.timezoneOffset,
		})
		user.DeletedAt.SetFromTime(user.DeletedAt.PrimitiveFromPostgreSQLType(deletedAt), entity.Timezone{
			Name:   r.timezoneName,
			Offset: r.timezoneOffset,
		})

		users = append(users, user)
		exists = true
	}

	if exists {
		result.Users = &users
	}

	return result
}

// FindDetailByID finds user detail by id from PostgreSQL database.
func (r *userPostgreSQL) FindDetailByID(param UserFindDetailByIDParam) UserFindDetailByIDResult {
	result := UserFindDetailByIDResult{}
	user := &entity.UserWithProfile{}
	profile := &entity.Profile{}
	userCreatedAt := user.CreatedAt.GetPostgreSQLType()
	userUpdatedAt := user.UpdatedAt.GetPostgreSQLType()
	userDeletedAt := user.DeletedAt.GetPostgreSQLType()
	profileImage := profile.Image.GetPostgreSQLType()
	profileCreatedAt := profile.CreatedAt.GetPostgreSQLType()
	profileUpdatedAt := profile.UpdatedAt.GetPostgreSQLType()

	if err := r.pool.QueryRow(
		param.Ctx,
		PostgreSQLUserFindDetailByIDQuery,
		param.ID,
	).Scan(
		&user.ID,
		&user.Username,
		&user.Email,
		&user.ActiveStatus,
		&userCreatedAt,
		&userUpdatedAt,
		&userDeletedAt,
		&profile.ID,
		&profile.FirstName,
		&profile.LastName,
		&profileImage,
		&profile.Gender,
		&profileCreatedAt,
		&profileUpdatedAt,
	); err != nil {
		if pgxpoolgo.ErrDB(err).IsNoRows() {
			return result
		}
		result.Error = ErrDatabaseUserFind
		return result
	}

	user.CreatedAt.SetFromTime(user.CreatedAt.PrimitiveFromPostgreSQLType(userCreatedAt), entity.Timezone{
		Name:   r.timezoneName,
		Offset: r.timezoneOffset,
	})
	user.UpdatedAt.SetFromTime(user.UpdatedAt.PrimitiveFromPostgreSQLType(userUpdatedAt), entity.Timezone{
		Name:   r.timezoneName,
		Offset: r.timezoneOffset,
	})
	user.DeletedAt.SetFromTime(user.DeletedAt.PrimitiveFromPostgreSQLType(userDeletedAt), entity.Timezone{
		Name:   r.timezoneName,
		Offset: r.timezoneOffset,
	})
	profile.Image.Set(profile.Image.PrimitiveFromPostgreSQLType(profileImage))
	profile.CreatedAt.SetFromTime(profile.CreatedAt.PrimitiveFromPostgreSQLType(profileCreatedAt), entity.Timezone{
		Name:   r.timezoneName,
		Offset: r.timezoneOffset,
	})
	profile.UpdatedAt.SetFromTime(profile.UpdatedAt.PrimitiveFromPostgreSQLType(profileUpdatedAt), entity.Timezone{
		Name:   r.timezoneName,
		Offset: r.timezoneOffset,
	})

	profile.UserID = user.ID
	user.Profile = profile

	return result
}

// FindByIDAndPassword finds user data by id and password from database.
func (r *userPostgreSQL) FindByIDAndPassword(param UserFindByIDAndPasswordParam) UserFindByIDAndPasswordResult {
	result := UserFindByIDAndPasswordResult{}
	user := &entity.User{}
	createdAt := user.CreatedAt.GetPostgreSQLType()
	updatedAt := user.UpdatedAt.GetPostgreSQLType()
	deletedAt := user.DeletedAt.GetPostgreSQLType()

	if err := r.pool.QueryRow(
		param.Ctx,
		PostgreSQLUserFindByIDAndPasswordQuery,
		param.ID,
		param.Password,
	).Scan(
		&user.ID,
		&user.Username,
		&user.Email,
		&user.ActiveStatus,
		&createdAt,
		&updatedAt,
		&deletedAt,
	); err != nil {
		if pgxpoolgo.ErrDB(err).IsNoRows() {
			return result
		}
		result.Error = ErrDatabaseUserFind
		return result
	}

	user.CreatedAt.SetFromTime(user.CreatedAt.PrimitiveFromPostgreSQLType(createdAt), entity.Timezone{
		Name:   r.timezoneName,
		Offset: r.timezoneOffset,
	})
	user.UpdatedAt.SetFromTime(user.UpdatedAt.PrimitiveFromPostgreSQLType(updatedAt), entity.Timezone{
		Name:   r.timezoneName,
		Offset: r.timezoneOffset,
	})
	user.DeletedAt.SetFromTime(user.DeletedAt.PrimitiveFromPostgreSQLType(deletedAt), entity.Timezone{
		Name:   r.timezoneName,
		Offset: r.timezoneOffset,
	})

	result.User = user

	return result
}

// FindByUsernameOrEmailAndPassword finds user data by username or email, and password from PostgreSQL database.
func (r *userPostgreSQL) FindByUsernameOrEmailAndPassword(param UserFindByUsernameOrEmailAndPasswordParam) UserFindByUsernameOrEmailAndPasswordResult {
	result := UserFindByUsernameOrEmailAndPasswordResult{}
	user := &entity.User{}
	createdAt := user.CreatedAt.GetPostgreSQLType()
	updatedAt := user.UpdatedAt.GetPostgreSQLType()
	deletedAt := user.DeletedAt.GetPostgreSQLType()

	if err := r.pool.QueryRow(
		param.Ctx,
		PostgreSQLUserFindByUsernameOrEmailAndPasswordQuery,
		param.UsernameOrEmail,
		param.UsernameOrEmail,
		param.Password,
	).Scan(
		&user.ID,
		&user.Username,
		&user.Email,
		&user.ActiveStatus,
		&createdAt,
		&updatedAt,
		&deletedAt,
	); err != nil {
		if pgxpoolgo.ErrDB(err).IsNoRows() {
			return result
		}
		result.Error = ErrDatabaseUserFind
		return result
	}

	user.CreatedAt.SetFromTime(user.CreatedAt.PrimitiveFromPostgreSQLType(createdAt), entity.Timezone{
		Name:   r.timezoneName,
		Offset: r.timezoneOffset,
	})
	user.UpdatedAt.SetFromTime(user.UpdatedAt.PrimitiveFromPostgreSQLType(updatedAt), entity.Timezone{
		Name:   r.timezoneName,
		Offset: r.timezoneOffset,
	})
	user.DeletedAt.SetFromTime(user.DeletedAt.PrimitiveFromPostgreSQLType(deletedAt), entity.Timezone{
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
	if err = tx.QueryRow(
		param.Ctx,
		PostgreSQLUserInsertQuery,
		param.Username,
		param.Email,
		param.Password,
	).Scan(
		&user.ID,
	); err != nil {
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
	if err = tx.QueryRow(
		param.Ctx,
		PostgreSQLProfileInsertQuery,
		user.ID,
		param.FirstName,
		param.LastName,
		param.Gender,
	).Scan(
		&profile.ID,
	); err != nil {
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

// UpdateByIDTx updates user data by id into PostgreSQL database.
// It will update user and profile data using transaction.
func (r *userPostgreSQL) UpdateByIDTx(param UserUpdateByIDTxParam) UserUpdateByIDTxResult {
	result := UserUpdateByIDTxResult{}

	tx, err := r.pool.BeginTx(param.Ctx, pgx.TxOptions{})
	if err != nil {
		result.Error = ErrDatabaseUserUpdate
		return result
	}

	reply, err := tx.Exec(
		param.Ctx,
		PostgreSQLUserUpdateByIDQuery,
		param.Username,
		param.ID,
	)
	if err != nil {
		if errR := tx.Rollback(param.Ctx); errR != nil {
			result.Error = ErrDatabaseUserUpdate
			return result
		}
		result.Error = ErrDatabaseUserUpdate
		result.IsUserDuplicateKey = pgxpoolgo.ErrDB(err).IsDuplicateKey()
		return result
	}
	if reply.RowsAffected() < 1 {
		if errR := tx.Rollback(param.Ctx); errR != nil {
			result.Error = ErrDatabaseUserUpdate
			return result
		}
		result.Error = ErrDatabaseUserUpdateNoAffected
		return result
	}

	reply, err = tx.Exec(
		param.Ctx,
		PostgreSQLProfileUpdateByUserIDQuery,
		param.FirstName,
		param.LastName,
		param.Gender,
		param.ID,
	)
	if err != nil {
		if errR := tx.Rollback(param.Ctx); errR != nil {
			result.Error = ErrDatabaseProfileUpdate
			return result
		}
		result.Error = ErrDatabaseProfileUpdate
		return result
	}
	if reply.RowsAffected() < 1 {
		if errR := tx.Rollback(param.Ctx); errR != nil {
			result.Error = ErrDatabaseUserUpdate
			return result
		}
		result.Error = ErrDatabaseProfileUpdateNoAffected
		return result
	}

	if err = tx.Commit(param.Ctx); err != nil {
		result.Error = ErrDatabaseUserUpdate
		return result
	}

	return result
}

// DeleteByID deletes user data by id from PostgreSQL database.
func (r *userPostgreSQL) DeleteByID(param UserDeleteByIDParam) UserDeleteByIDResult {
	result := UserDeleteByIDResult{}

	reply, err := r.pool.Exec(
		param.Ctx,
		PostgreSQLUserDeleteByIDQuery,
		param.ID,
	)
	if err != nil {
		result.Error = ErrDatabaseUserDelete
		return result
	}
	if reply.RowsAffected() < 1 {
		result.Error = ErrDatabaseUserDeleteNoAffected
		return result
	}

	return result
}
