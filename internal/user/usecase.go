package user

import "github.com/Patrikesm/basic-solid-api/internal/models"

type UseCase interface {
	Create(user *models.User) error
	ListAll() ([]*models.User, error)
}
