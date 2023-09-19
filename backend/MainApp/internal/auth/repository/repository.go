package repository

import (
	"writesend/MainApp/models"
)

type RepositoryI interface {
	CreateCookie(cookie *models.Cookie) error
	GetCookie(value string) (string, error)
	DeleteCookie(value string) error
}
