package view

import "shope_clone/internal/user/domain"

type SimpleUser struct {
	ID       uint   `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"-"`
	Salt     string `json:"-"`
}

// add Function Table Namme
func (SimpleUser) TableName() string {
	return "users"
}

func FromDomainToSimple(user *domain.User) *SimpleUser {
	return &SimpleUser{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}
}
