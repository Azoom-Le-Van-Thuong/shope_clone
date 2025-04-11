package domain

type UserRepository interface {
	Create(user *User) error
	GetAll() ([]User, error)
}
