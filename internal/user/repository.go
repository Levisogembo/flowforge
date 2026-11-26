package user

import "context"

type Repository interface {
	Create(ctx context.Context, u *User) error
	GetByEmail(ctx context.Context, email string) (*User, error)
	GetById(ctx context.Context, id string) (*User, error)
}
