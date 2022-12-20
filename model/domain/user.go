package domain

import "time"

type User struct {
	Id        int       `json:"id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	Bio       string    `json:"bio"`
	UpdatedAt time.Time `json:"updated_at"`
}

type GoodUser struct {
	Id        int
	FirstName string
	LastName  string
	Username  string
	Email     string
	Bio       string
	UpdatedAt time.Time
}
