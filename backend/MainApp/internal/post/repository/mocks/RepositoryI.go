// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	models "writesend/MainApp/models"
	mock "github.com/stretchr/testify/mock"
)

// RepositoryI is an autogenerated mock type for the RepositoryI type
type RepositoryI struct {
	mock.Mock
}

// AddComment provides a mock function with given fields: comment
func (_m *RepositoryI) AddComment(comment *models.Comment) error {
	ret := _m.Called(comment)

	var r0 error
	if rf, ok := ret.Get(0).(func(*models.Comment) error); ok {
		r0 = rf(comment)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CheckLikePost provides a mock function with given fields: id, userID
func (_m *RepositoryI) CheckLikePost(id uint64, userID uint64) (bool, error) {
	ret := _m.Called(id, userID)

	var r0 bool
	if rf, ok := ret.Get(0).(func(uint64, uint64) bool); ok {
		r0 = rf(id, userID)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(uint64, uint64) error); ok {
		r1 = rf(id, userID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreatePost provides a mock function with given fields: p
func (_m *RepositoryI) CreatePost(p *models.Post) error {
	ret := _m.Called(p)

	var r0 error
	if rf, ok := ret.Get(0).(func(*models.Post) error); ok {
		r0 = rf(p)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteComment provides a mock function with given fields: id
func (_m *RepositoryI) DeleteComment(id uint64) error {
	ret := _m.Called(id)

	var r0 error
	if rf, ok := ret.Get(0).(func(uint64) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeletePostById provides a mock function with given fields: id
func (_m *RepositoryI) DeletePostById(id uint64) error {
	ret := _m.Called(id)

	var r0 error
	if rf, ok := ret.Get(0).(func(uint64) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetAllPosts provides a mock function with given fields:
func (_m *RepositoryI) GetAllPosts() ([]*models.Post, error) {
	ret := _m.Called()

	var r0 []*models.Post
	if rf, ok := ret.Get(0).(func() []*models.Post); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*models.Post)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetCommentById provides a mock function with given fields: id
func (_m *RepositoryI) GetCommentById(id uint64) (*models.Comment, error) {
	ret := _m.Called(id)

	var r0 *models.Comment
	if rf, ok := ret.Get(0).(func(uint64) *models.Comment); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Comment)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(uint64) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetComments provides a mock function with given fields: postId
func (_m *RepositoryI) GetComments(postId uint64) ([]*models.Comment, error) {
	ret := _m.Called(postId)

	var r0 []*models.Comment
	if rf, ok := ret.Get(0).(func(uint64) []*models.Comment); ok {
		r0 = rf(postId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*models.Comment)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(uint64) error); ok {
		r1 = rf(postId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetCommunityPosts provides a mock function with given fields: userId
func (_m *RepositoryI) GetCommunityPosts(userId uint64) ([]*models.Post, error) {
	ret := _m.Called(userId)

	var r0 []*models.Post
	if rf, ok := ret.Get(0).(func(uint64) []*models.Post); ok {
		r0 = rf(userId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*models.Post)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(uint64) error); ok {
		r1 = rf(userId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetCountLikesPost provides a mock function with given fields: id
func (_m *RepositoryI) GetCountLikesPost(id uint64) (uint64, error) {
	ret := _m.Called(id)

	var r0 uint64
	if rf, ok := ret.Get(0).(func(uint64) uint64); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(uint64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(uint64) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetPostById provides a mock function with given fields: id
func (_m *RepositoryI) GetPostById(id uint64) (*models.Post, error) {
	ret := _m.Called(id)

	var r0 *models.Post
	if rf, ok := ret.Get(0).(func(uint64) *models.Post); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Post)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(uint64) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetUserPosts provides a mock function with given fields: userId
func (_m *RepositoryI) GetUserPosts(userId uint64) ([]*models.Post, error) {
	ret := _m.Called(userId)

	var r0 []*models.Post
	if rf, ok := ret.Get(0).(func(uint64) []*models.Post); ok {
		r0 = rf(userId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*models.Post)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(uint64) error); ok {
		r1 = rf(userId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// LikePost provides a mock function with given fields: id, userId
func (_m *RepositoryI) LikePost(id uint64, userId uint64) error {
	ret := _m.Called(id, userId)

	var r0 error
	if rf, ok := ret.Get(0).(func(uint64, uint64) error); ok {
		r0 = rf(id, userId)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UnLikePost provides a mock function with given fields: id, userId
func (_m *RepositoryI) UnLikePost(id uint64, userId uint64) error {
	ret := _m.Called(id, userId)

	var r0 error
	if rf, ok := ret.Get(0).(func(uint64, uint64) error); ok {
		r0 = rf(id, userId)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateComment provides a mock function with given fields: comment
func (_m *RepositoryI) UpdateComment(comment *models.Comment) error {
	ret := _m.Called(comment)

	var r0 error
	if rf, ok := ret.Get(0).(func(*models.Comment) error); ok {
		r0 = rf(comment)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdatePost provides a mock function with given fields: post
func (_m *RepositoryI) UpdatePost(post *models.Post) error {
	ret := _m.Called(post)

	var r0 error
	if rf, ok := ret.Get(0).(func(*models.Post) error); ok {
		r0 = rf(post)
	} else {
		r0 = ret.Error(0)
	}

	return r0
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
