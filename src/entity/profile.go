package entity

type Profile struct {
	ID        ID        `json:"id"`
	UserID    ID        `json:"user_id"`
	FirstName FirstName `json:"first_name"`
	LastName  LastName  `json:"last_name"`
	Image     Image     `json:"image"`
	Gender    Gender    `json:"gender"`
	CreatedAt Time      `json:"created_at"`
	UpdatedAt Time      `json:"updated_at"`
}

type Profiles []*Profile

// For client

type ProfileForClient Profile

type ProfilesForClient []*ProfileForClient
