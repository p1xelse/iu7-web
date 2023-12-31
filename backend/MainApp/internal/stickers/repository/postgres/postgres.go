package postgres

import (
	"github.com/pkg/errors"
	"gorm.io/gorm"

	stickerRep "writesend/MainApp/internal/stickers/repository"
	"writesend/MainApp/models"
)

type dataBase struct {
	db *gorm.DB
}

func New(db *gorm.DB) stickerRep.RepositoryI {
	return &dataBase{
		db: db,
	}
}

func (dbStickers *dataBase) GetStickerByID(id uint64) (*models.Sticker, error) {
	sticker := models.Sticker{}

	tx := dbStickers.db.Where("id = ?", id).Take(&sticker)
	if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
		return nil, models.ErrNotFound
	} else if tx.Error != nil {
		return nil, errors.Wrap(tx.Error, "database error (table stickers)")
	}

	return &sticker, nil
}

func (dbStickers *dataBase) GetAllStickers() ([]*models.Sticker, error) {
	stickers := make([]*models.Sticker, 0, 10)
	tx := dbStickers.db.Find(&stickers)
	if tx.Error != nil {
		return nil, errors.Wrap(tx.Error, "database error (table stickers)")
	}

	return stickers, nil
}
