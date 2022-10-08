package repository

// User

var PostgreSQLUserFindAllActiveQuery = `
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
	    active_status = 1
`

var PostgreSQLUserFindDetailByIDQuery = `
	SELECT
		u.id,
		u.username,
		u.email,
		u.active_status,
		u.created_at,
		u.updated_at,
		u.deleted_at,
		p.id AS profile_id,
		p.first_name AS profile_first_name,
		p.last_name AS profile_last_name,
		p.image AS profile_image,
		p.gender AS profile_gender,
		p.created_at AS profile_created_at,
		p.updated_at AS profile_updated_at
	FROM
	    users AS u
	LEFT JOIN 
	    profiles AS p ON p.user_id = u.id
	WHERE
	    u.id = $1
	LIMIT 1
`

var PostgreSQLUserFindByIDAndPasswordQuery = `
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
	    id = $1 AND
	    password = $2
	LIMIT 1
`

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

var PostgreSQLUserUpdateByIDQuery = `
	UPDATE
		users
	SET
	    username = $1
	WHERE
	    id = $2 AND
	    active_status = 1 AND
	    deleted_at IS NOT NULL
`

var PostgreSQLUserDeleteByIDQuery = `
	UPDATE
		users
	SET
	    active_status = 0,
	    deleted_at = CURRENT_TIMESTAMP
	WHERE
	    id = $1 AND
	    active_status = 1 AND
	    deleted_at IS NOT NULL
`

// Profile

var PostgreSQLProfileInsertQuery = `
	INSERT INTO
		profiles
	(user_id, first_name, last_name, gender, created_at, updated_at)
	VALUES
		($1, $2, $3, $4, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP)
	RETURNING id
`

var PostgreSQLProfileUpdateByUserIDQuery = `
	UPDATE
		profiles
	SET
	    first_name = $1,
	    last_name = $2,
	    gender = $3
	WHERE
	    user_id = $4
`

var PostgreSQLProfileUpdateImageByUserIDQuery = `
	UPDATE
		profiles
	SET
	    image = $1
	WHERE
	    user_id = $2
`
