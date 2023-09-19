package repository

import (
	"writesend/AuthMicroservice/models"
)

type RepositoryI interface {
	CreateCookie(cookie *models.Cookie) error
	GetCookie(value string) (string, error)
	DeleteCookie(value string) error
}
