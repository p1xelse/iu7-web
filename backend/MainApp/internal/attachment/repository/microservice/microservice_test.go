package microservice_test

import (
	"context"
	"testing"

	"github.com/bxcodec/faker"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	attachmentRep "writesend/MainApp/internal/attachment/repository/microservice"
	"writesend/MainApp/models"
	Attachment "writesend/MainApp/proto/attachment"
	attachmentMocks "writesend/MainApp/proto/attachment/mocks"
)

type TestCaseCreateAttachment struct {
	ArgData *models.Attachment
	Error   error
}

type TestCaseGetPostAttachments struct {
	ArgData     uint64
	ExpectedRes []*models.Attachment
	Error       error
}

type TestCaseGetAttachment struct {
	ArgData     uint64
	ExpectedRes *models.Attachment
	Error       error
}

func TestMicroserviceCreateAttachment(t *testing.T) {
	mockPbAttachment := Attachment.Attachment{
		AttLink: "link1",
	}

	att := models.Attachment{
		AttLink: mockPbAttachment.AttLink,
	}

	pbAttachmentId := Attachment.AttachmentId{
		AttachmentId: 1,
	}

	mockPbAttachmentError := Attachment.Attachment{
		AttLink: "link2",
	}

	attError := models.Attachment{
		AttLink: mockPbAttachmentError.AttLink,
	}

	mockAttachmentClient := attachmentMocks.NewAttachmentsClient(t)

	ctx := context.Background()

	createErr := errors.New("error")

	mockAttachmentClient.On("CreateAttachment", ctx, &mockPbAttachment).Return(&pbAttachmentId, nil)
	mockAttachmentClient.On("CreateAttachment", ctx, &mockPbAttachmentError).Return(nil, createErr)

	repository := attachmentRep.New(mockAttachmentClient)

	cases := map[string]TestCaseCreateAttachment{
		"success": {
			ArgData: &att,
			Error:   nil,
		},
		"error": {
			ArgData: &attError,
			Error:   createErr,
		},
	}

	for name, test := range cases {
		t.Run(name, func(t *testing.T) {
			err := repository.CreateAttachment(test.ArgData)
			require.Equal(t, test.Error, errors.Cause(err))
		})
	}
	mockAttachmentClient.AssertExpectations(t)
}

func TestMicroserviceGetPostAttachments(t *testing.T) {
	pbPostId := Attachment.GetPostAttachmentsRequest{
		PostId: 1,
	}

	var mockPbAttachments Attachment.GetPostAttachmentsResponse
	err := faker.FakeData(&mockPbAttachments)
	assert.NoError(t, err)

	attachments := make([]*models.Attachment, 0)

	for idx := range mockPbAttachments.Attachments {
		att := models.Attachment{
			ID:      mockPbAttachments.Attachments[idx].Id,
			AttLink: mockPbAttachments.Attachments[idx].AttLink,
			Type:    mockPbAttachments.Attachments[idx].Type,
		}

		attachments = append(attachments, &att)
	}

	pbPostIdError := Attachment.GetPostAttachmentsRequest{
		PostId: 2,
	}

	mockAttachmentClient := attachmentMocks.NewAttachmentsClient(t)

	ctx := context.Background()

	getErr := errors.New("error")

	mockAttachmentClient.On("GetPostAttachments", ctx, &pbPostId).Return(&mockPbAttachments, nil)
	mockAttachmentClient.On("GetPostAttachments", ctx, &pbPostIdError).Return(nil, getErr)

	repository := attachmentRep.New(mockAttachmentClient)

	cases := map[string]TestCaseGetPostAttachments{
		"success": {
			ArgData:     pbPostId.PostId,
			ExpectedRes: attachments,
			Error:       nil,
		},
		"error": {
			ArgData:     pbPostIdError.PostId,
			ExpectedRes: nil,
			Error:       getErr,
		},
	}

	for name, test := range cases {
		t.Run(name, func(t *testing.T) {
			selectedAttachments, err := repository.GetPostAttachments(test.ArgData)
			require.Equal(t, test.Error, errors.Cause(err))

			if err == nil {
				assert.Equal(t, test.ExpectedRes, selectedAttachments)
			}
		})
	}
	mockAttachmentClient.AssertExpectations(t)
}

func TestMicroserviceGetAttachment(t *testing.T) {
	pbAttachmentId := Attachment.AttachmentId{
		AttachmentId: 1,
	}

	mockPbAttachment := Attachment.Attachment{
		Id:      pbAttachmentId.AttachmentId,
		AttLink: "link1",
	}

	att := &models.Attachment{
		ID:      mockPbAttachment.Id,
		AttLink: mockPbAttachment.AttLink,
		Type:    mockPbAttachment.Type,
	}

	pbAttachmentIdError := Attachment.AttachmentId{
		AttachmentId: 2,
	}

	mockAttachmentClient := attachmentMocks.NewAttachmentsClient(t)

	ctx := context.Background()

	getErr := errors.New("error")

	mockAttachmentClient.On("GetAttachment", ctx, &pbAttachmentId).Return(&mockPbAttachment, nil)
	mockAttachmentClient.On("GetAttachment", ctx, &pbAttachmentIdError).Return(nil, getErr)

	repository := attachmentRep.New(mockAttachmentClient)

	cases := map[string]TestCaseGetAttachment{
		"success": {
			ArgData:     pbAttachmentId.AttachmentId,
			ExpectedRes: att,
			Error:       nil,
		},
		"error": {
			ArgData:     pbAttachmentIdError.AttachmentId,
			ExpectedRes: nil,
			Error:       getErr,
		},
	}

	for name, test := range cases {
		t.Run(name, func(t *testing.T) {
			selectedAttachment, err := repository.GetAttachment(test.ArgData)
			require.Equal(t, test.Error, errors.Cause(err))

			if err == nil {
				assert.Equal(t, test.ExpectedRes, selectedAttachment)
			}
		})
	}
	mockAttachmentClient.AssertExpectations(t)
}
