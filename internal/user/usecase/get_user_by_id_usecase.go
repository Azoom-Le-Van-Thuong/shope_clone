package usecase

// Implementation of the GetUser use case

import (
	"context"
	"shope_clone/internal/user/domain"
)

type GetUserByIDUseCase struct {
	repo domain.UserRepository
}

func NewGetUserByIDUseCase(repo domain.UserRepository) *GetUserByIDUseCase {
	return &GetUserByIDUseCase{repo: repo}
}

func (uc *GetUserByIDUseCase) Execute(ctx context.Context, id uint) (*domain.User, error) {
	return uc.repo.GetByID(ctx, id)
}
