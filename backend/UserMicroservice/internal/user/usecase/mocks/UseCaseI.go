// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	__ "writesend/UserMicroservice/proto"
	mock "github.com/stretchr/testify/mock"
)

// UseCaseI is an autogenerated mock type for the UseCaseI type
type UseCaseI struct {
	mock.Mock
}

// AddFriend provides a mock function with given fields: _a0
func (_m *UseCaseI) AddFriend(_a0 *__.Friends) (*__.Nothing, error) {
	ret := _m.Called(_a0)

	var r0 *__.Nothing
	if rf, ok := ret.Get(0).(func(*__.Friends) *__.Nothing); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*__.Nothing)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*__.Friends) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CheckFriends provides a mock function with given fields: _a0
func (_m *UseCaseI) CheckFriends(_a0 *__.Friends) (*__.CheckFriendsResponse, error) {
	ret := _m.Called(_a0)

	var r0 *__.CheckFriendsResponse
	if rf, ok := ret.Get(0).(func(*__.Friends) *__.CheckFriendsResponse); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*__.CheckFriendsResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*__.Friends) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateUser provides a mock function with given fields: _a0
func (_m *UseCaseI) CreateUser(_a0 *__.User) (*__.Nothing, error) {
	ret := _m.Called(_a0)

	var r0 *__.Nothing
	if rf, ok := ret.Get(0).(func(*__.User) *__.Nothing); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*__.Nothing)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*__.User) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteFriend provides a mock function with given fields: _a0
func (_m *UseCaseI) DeleteFriend(_a0 *__.Friends) (*__.Nothing, error) {
	ret := _m.Called(_a0)

	var r0 *__.Nothing
	if rf, ok := ret.Get(0).(func(*__.Friends) *__.Nothing); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*__.Nothing)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*__.Friends) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SearchUsers provides a mock function with given fields: _a0
func (_m *UseCaseI) SearchUsers(_a0 *__.SearchUsersRequest) (*__.UsersList, error) {
	ret := _m.Called(_a0)

	var r0 *__.UsersList
	if rf, ok := ret.Get(0).(func(*__.SearchUsersRequest) *__.UsersList); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*__.UsersList)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*__.SearchUsersRequest) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SelectAllUsers provides a mock function with given fields: _a0
func (_m *UseCaseI) SelectAllUsers(_a0 *__.Nothing) (*__.UsersList, error) {
	ret := _m.Called(_a0)

	var r0 *__.UsersList
	if rf, ok := ret.Get(0).(func(*__.Nothing) *__.UsersList); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*__.UsersList)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*__.Nothing) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SelectFriends provides a mock function with given fields: _a0
func (_m *UseCaseI) SelectFriends(_a0 *__.UserId) (*__.UsersList, error) {
	ret := _m.Called(_a0)

	var r0 *__.UsersList
	if rf, ok := ret.Get(0).(func(*__.UserId) *__.UsersList); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*__.UsersList)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*__.UserId) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SelectUserByEmail provides a mock function with given fields: _a0
func (_m *UseCaseI) SelectUserByEmail(_a0 *__.SelectUserByEmailRequest) (*__.User, error) {
	ret := _m.Called(_a0)

	var r0 *__.User
	if rf, ok := ret.Get(0).(func(*__.SelectUserByEmailRequest) *__.User); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*__.User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*__.SelectUserByEmailRequest) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SelectUserById provides a mock function with given fields: _a0
func (_m *UseCaseI) SelectUserById(_a0 *__.UserId) (*__.User, error) {
	ret := _m.Called(_a0)

	var r0 *__.User
	if rf, ok := ret.Get(0).(func(*__.UserId) *__.User); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*__.User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*__.UserId) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SelectUserByNickName provides a mock function with given fields: _a0
func (_m *UseCaseI) SelectUserByNickName(_a0 *__.SelectUserByNickNameRequest) (*__.User, error) {
	ret := _m.Called(_a0)

	var r0 *__.User
	if rf, ok := ret.Get(0).(func(*__.SelectUserByNickNameRequest) *__.User); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*__.User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*__.SelectUserByNickNameRequest) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateUser provides a mock function with given fields: _a0
func (_m *UseCaseI) UpdateUser(_a0 *__.User) (*__.Nothing, error) {
	ret := _m.Called(_a0)

	var r0 *__.Nothing
	if rf, ok := ret.Get(0).(func(*__.User) *__.Nothing); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*__.Nothing)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*__.User) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewUseCaseI interface {
	mock.TestingT
	Cleanup(func())
}

// NewUseCaseI creates a new instance of UseCaseI. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewUseCaseI(t mockConstructorTestingTNewUseCaseI) *UseCaseI {
	mock := &UseCaseI{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
