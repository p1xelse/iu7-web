package repository

import "writesend/MainApp/models"

type RepositoryI interface {
	GetCommunity(id uint64) (*models.Community, error)
	UpdateCommunity(comm *models.Community) error
	CreateCommunity(comm *models.Community) error
	SearchCommunities(searchString string) ([]*models.Community, error)
	DeleteCommunity(id uint64) error
	GetAllCommunities() ([]*models.Community, error)
}
