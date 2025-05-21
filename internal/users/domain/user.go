package domain

import "github.com/google/uuid"

type User struct {
	ID       uuid.UUID `json:"id"`
	Name     string    `json:"name"`
	Email    string    `json:"email"`
	Password string    `json:"-"`
}

type Repository interface {
	Create(user *User) error
	GetByID(id uuid.UUID) (*User, error)
	GetAll() ([]*User, error)
	GetByEmail(email string) (*User, error)
} 