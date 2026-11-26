package user

import "time"

type User struct {
	ID           string    `json:"id"`
	Email        string    `json:"email"`
	PasswordHash string    `json:"-"` //the value will not be returned to the user in the json response
	FullName     string    `json:"full_name"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}
