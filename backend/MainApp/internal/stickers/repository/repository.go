package repository

import "writesend/MainApp/models"

type RepositoryI interface {
	GetAllStickers() ([]*models.Sticker, error)
	GetStickerByID(id uint64) (*models.Sticker, error)
}
