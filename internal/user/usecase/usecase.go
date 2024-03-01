package usecase

import (
	"github.com/Patrikesm/basic-solid-api/internal/models"
	"github.com/Patrikesm/basic-solid-api/internal/user"
	"github.com/Patrikesm/basic-solid-api/internal/user/repository"
)

type UserUC struct {
	UserRepo repository.UserRepository
}

func NewUserUseCase(userRepo repository.UserRepository) user.UseCase {
	return &UserUC{
		UserRepo: userRepo,
	}
}

func (u *UserUC) Create(user *models.User) error {
	err := u.UserRepo.CreateUser(user)

	if err != nil {
		return err
	}

	return nil
}

func (u *UserUC) ListAll() ([]*models.User, error) {
	users, err := u.UserRepo.ListAll()

	if err != nil {
		return nil, err
	}

	return users, nil
}
