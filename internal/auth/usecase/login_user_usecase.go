package authUseCase

import (
	"context"
	authdomain "shope_clone/internal/auth/domain"
	domain "shope_clone/internal/auth/domain/repository"
	"shope_clone/internal/auth/domain/valueobject"
	"shope_clone/internal/auth/dto"
	"shope_clone/internal/auth/token"
)

type LoginUseCase struct {
	JWT           token.Token
	LoginUserRepo domain.LoginUserRepository
}

func NewLoginUseCase(repo domain.LoginUserRepository, jwt token.Token) *LoginUseCase {
	return &LoginUseCase{LoginUserRepo: repo, JWT: jwt}
}

func (uc *LoginUseCase) Execute(ctx context.Context, req dto.LoginRequest) (*dto.LoginResponse, error) {
	user, err := uc.LoginUserRepo.FindUserByEmail(ctx, req.Email)
	if err != nil {
		return nil, authdomain.ErrInvalidEmailOrPassword
	}

	password := valueobject.NewPasswordFromHash(user.Salt, user.Password)
	if x := password.Compare(ctx, req.Password); x != nil {
		return nil, authdomain.ErrInvalidEmailOrPassword
	}

	accessToken, err := uc.JWT.GenerateToken(user.ID, user.Email)
	if err != nil {
		return nil, err
	}

	return &dto.LoginResponse{
		AccessToken: accessToken,
		UserID:      user.ID,
		Email:       user.Email,
	}, nil
}
