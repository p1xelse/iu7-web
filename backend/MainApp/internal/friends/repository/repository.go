package repository

import (
	"writesend/MainApp/models"
)

type RepositoryI interface {
	AddFriend(friends models.Friends) error
	DeleteFriend(friends models.Friends) error
	CheckFriends(friends models.Friends) (bool, error)
	SelectFriends(id uint64) ([]models.User, error)
}
