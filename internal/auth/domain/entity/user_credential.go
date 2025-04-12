package entity

import "shope_clone/internal/auth/domain/valueobject"

type UserCredential struct {
	ID       int
	Email    valueobject.Email
	Password *valueobject.Password
	Name     string
}

func ToUserAuthCreation(u UserCredential) *UserAuthCreation {
	return &UserAuthCreation{
		Name:     u.Name,
		Password: u.Password.Hash(),
		Salt:     u.Password.Salt(),
		Email:    u.Email.Value(),
	}
}

type UserAuthCreation struct {
	Name     string `json:"name" gorm:"name"`
	Password string `json:"password" gorm:"password"`
	Salt     string `json:"salt" gorm:"salt"`
	Email    string `json:"email" gorm:"email uniqueIndex"`
}

func (receiver UserAuthCreation) TableName() string {
	return "users"
}
