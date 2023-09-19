// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	models "writesend/AttachmentMicroservice/models"
	mock "github.com/stretchr/testify/mock"
)

// RepositoryI is an autogenerated mock type for the RepositoryI type
type RepositoryI struct {
	mock.Mock
}

// CreateAttachment provides a mock function with given fields: attachment
func (_m *RepositoryI) CreateAttachment(attachment *models.Attachment) error {
	ret := _m.Called(attachment)

	var r0 error
	if rf, ok := ret.Get(0).(func(*models.Attachment) error); ok {
		r0 = rf(attachment)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetAttachment provides a mock function with given fields: attachmentID
func (_m *RepositoryI) GetAttachment(attachmentID uint64) (*models.Attachment, error) {
	ret := _m.Called(attachmentID)

	var r0 *models.Attachment
	if rf, ok := ret.Get(0).(func(uint64) *models.Attachment); ok {
		r0 = rf(attachmentID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Attachment)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(uint64) error); ok {
		r1 = rf(attachmentID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetMessageAttachments provides a mock function with given fields: postID
func (_m *RepositoryI) GetMessageAttachments(postID uint64) ([]*models.Attachment, error) {
	ret := _m.Called(postID)

	var r0 []*models.Attachment
	if rf, ok := ret.Get(0).(func(uint64) []*models.Attachment); ok {
		r0 = rf(postID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*models.Attachment)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(uint64) error); ok {
		r1 = rf(postID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetPostAttachments provides a mock function with given fields: postID
func (_m *RepositoryI) GetPostAttachments(postID uint64) ([]*models.Attachment, error) {
	ret := _m.Called(postID)

	var r0 []*models.Attachment
	if rf, ok := ret.Get(0).(func(uint64) []*models.Attachment); ok {
		r0 = rf(postID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*models.Attachment)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(uint64) error); ok {
		r1 = rf(postID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewRepositoryI interface {
	mock.TestingT
	Cleanup(func())
}

// NewRepositoryI creates a new instance of RepositoryI. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewRepositoryI(t mockConstructorTestingTNewRepositoryI) *RepositoryI {
	mock := &RepositoryI{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
