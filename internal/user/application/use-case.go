package application

import "shope_clone/internal/user/domain"

type UserUseCase struct {
	Repo domain.UserRepository
}

func NewUserUseCase(repo domain.UserRepository) *UserUseCase {
	return &UserUseCase{Repo: repo}
}

func (uc *UserUseCase) CreateUser(user *domain.User) error {
	return uc.Repo.Create(user)
}

func (uc *UserUseCase) GetUsers() ([]domain.User, error) {
	return uc.Repo.GetAll()
}
