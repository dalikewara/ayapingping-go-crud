package repository

// User

var PostgreSQLUserFindByUsernameOrEmailAndPasswordQuery = `
	SELECT
		id,
		username,
		email,
		active_status,
		created_at,
		updated_at,
		deleted_at
	FROM
	    users
	WHERE
	    (username = $1 OR email = $2) AND
	    password = $3
	LIMIT 1
`

var PostgreSQLUserInsertQuery = `
	INSERT INTO
		users
	(username, email, password, active_status, created_at, updated_at, deleted_at)
	VALUES
		($1, $2, $3, 0, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP, NULL)
	RETURNING id
`

// Profile

var PostgreSQLProfileInsertQuery = `
	INSERT INTO
		profiles
	(user_id, first_name, last_name, gender, created_at, updated_at, deleted_at)
	VALUES
		($1, $2, $3, $4, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP, NULL)
	RETURNING id
`
