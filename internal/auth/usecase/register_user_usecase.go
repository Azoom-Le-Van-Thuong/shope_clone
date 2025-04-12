package authUseCase

import (
	"context"
	"shope_clone/internal/auth/domain/entity"
	"shope_clone/internal/auth/domain/repository"
	"shope_clone/internal/auth/domain/valueobject"
	"shope_clone/internal/auth/dto"
)

type RegisterUserUseCase struct {
	registerUserRepository domain.RegisterUserRepository
}
 
func NewRegisterUserUseCase(userCredentialRepository domain.RegisterUserRepository) *RegisterUserUseCase {
	return &RegisterUserUseCase{
		registerUserRepository: userCredentialRepository,
	}
}

func (uc *RegisterUserUseCase) Execute(ctx context.Context, userBody dto.RegisterRequest) error {
	emailRaw := userBody.Email
	passwordRaw := userBody.Password
	email, err := valueobject.NewEmail(emailRaw)
	if err != nil {
		return err
	}
	password, err := valueobject.NewPassword(passwordRaw)
	if err != nil {
		return err
	}

	userInDb, err := uc.registerUserRepository.FindUserByEmail(ctx, email.Value())
	if err != nil {
		return err
	}
	if userInDb != nil {

	}

	user := entity.UserCredential{
		ID:       0,
		Email:    email,
		Password: password,
		Name:     userBody.Name,
	}
	return uc.registerUserRepository.Insert(ctx, user)
}
