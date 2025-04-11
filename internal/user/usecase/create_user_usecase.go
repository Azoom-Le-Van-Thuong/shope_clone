package usecase

import (
	"context"
	"shope_clone/internal/user/domain"
)

type CreateUserUseCase struct {
	repo domain.UserRepository
}

func NewCreateUserUseCase(repo domain.UserRepository) *CreateUserUseCase {
	return &CreateUserUseCase{repo}
}

func (uc *CreateUserUseCase) Execute(context context.Context, user *domain.User) error {
	return uc.repo.Insert(context, user)
}
