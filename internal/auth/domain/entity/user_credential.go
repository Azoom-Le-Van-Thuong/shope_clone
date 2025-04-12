package entity

import "shope_clone/internal/auth/domain/valueobject"

type UserCredential struct {
	ID       int
	Email    valueobject.Email
	Password valueobject.Password
	Name     string
}
