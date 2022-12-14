package entity

type User struct {
	ID           ID           `json:"id"`
	Username     Username     `json:"username"`
	Email        Email        `json:"email"`
	Password     Password     `json:"password"`
	ActiveStatus ActiveStatus `json:"active_status"`
	CreatedAt    Time         `json:"created_at"`
	UpdatedAt    Time         `json:"updated_at"`
	DeletedAt    Time         `json:"deleted_at"`
}

type UserWithProfile struct {
	User
	Profile *Profile `json:"profile"`
}

type Users []*User

type UsersWithProfile []*UserWithProfile

// For client

type UserForClient struct {
	ID           ID           `json:"id"`
	Username     Username     `json:"username"`
	Email        Email        `json:"email"`
	ActiveStatus ActiveStatus `json:"active_status"`
	CreatedAt    Time         `json:"created_at"`
	UpdatedAt    Time         `json:"updated_at"`
}

type UserWithProfileForClient struct {
	UserForClient
	Profile *ProfileForClient `json:"profile"`
}

type UsersForClient []*UserForClient

type UsersWithProfileForClient []*UserWithProfileForClient
