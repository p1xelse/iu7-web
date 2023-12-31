package repository

import (
	"writesend/MainApp/models"
)

type RepositoryI interface {
	SelectUserByNickName(name string) (*models.User, error)
	SelectUserByEmail(email string) (*models.User, error)
	CreateUser(u *models.User) error
	SelectUserById(id uint64) (*models.User, error)
	UpdateUser(user models.User) error
	SelectAllUsers() ([]models.User, error)
	SearchUsers(name string) ([]models.User, error)
}
