package repository

import (
	"writesend/AttachmentMicroservice/models"
)

type RepositoryI interface {
	GetPostAttachments(postID uint64) ([]*models.Attachment, error)
	GetMessageAttachments(postID uint64) ([]*models.Attachment, error)
	GetAttachment(attachmentID uint64) (*models.Attachment, error)
	CreateAttachment(attachment *models.Attachment) error
}
