// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	context "context"

	__ "writesend/MainApp/proto/auth"

	mock "github.com/stretchr/testify/mock"
)

// AuthServer is an autogenerated mock type for the AuthServer type
type AuthServer struct {
	mock.Mock
}

// CreateCookie provides a mock function with given fields: _a0, _a1
func (_m *AuthServer) CreateCookie(_a0 context.Context, _a1 *__.Cookie) (*__.Nothing, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *__.Nothing
	if rf, ok := ret.Get(0).(func(context.Context, *__.Cookie) *__.Nothing); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*__.Nothing)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *__.Cookie) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteCookie provides a mock function with given fields: _a0, _a1
func (_m *AuthServer) DeleteCookie(_a0 context.Context, _a1 *__.ValueCookieRequest) (*__.Nothing, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *__.Nothing
	if rf, ok := ret.Get(0).(func(context.Context, *__.ValueCookieRequest) *__.Nothing); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*__.Nothing)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *__.ValueCookieRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetCookie provides a mock function with given fields: _a0, _a1
func (_m *AuthServer) GetCookie(_a0 context.Context, _a1 *__.ValueCookieRequest) (*__.GetCookieResponse, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *__.GetCookieResponse
	if rf, ok := ret.Get(0).(func(context.Context, *__.ValueCookieRequest) *__.GetCookieResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*__.GetCookieResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *__.ValueCookieRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// mustEmbedUnimplementedAuthServer provides a mock function with given fields:
func (_m *AuthServer) mustEmbedUnimplementedAuthServer() {
	_m.Called()
}

type mockConstructorTestingTNewAuthServer interface {
	mock.TestingT
	Cleanup(func())
}

// NewAuthServer creates a new instance of AuthServer. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewAuthServer(t mockConstructorTestingTNewAuthServer) *AuthServer {
	mock := &AuthServer{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
