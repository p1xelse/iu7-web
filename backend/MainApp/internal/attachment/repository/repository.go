package repository

import "writesend/MainApp/models"

type RepositoryI interface {
	GetPostAttachments(postID uint64) ([]*models.Attachment, error)
	GetMessageAttachments(messageId uint64) ([]models.Attachment, error)
	GetAttachment(attachmentID uint64) (*models.Attachment, error)
	CreateAttachment(Attachment *models.Attachment) error
}
