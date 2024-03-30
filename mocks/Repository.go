// Code generated by mockery v2.42.1. DO NOT EDIT.

package mocks

import (
	context "context"

	config "github.com/VanLavr/auth/internal/pkg/config"

	mock "github.com/stretchr/testify/mock"

	models "github.com/VanLavr/auth/internal/models"
)

// Repository is an autogenerated mock type for the Repository type
type Repository struct {
	mock.Mock
}

// CloseConnetion provides a mock function with given fields: _a0
func (_m *Repository) CloseConnetion(_a0 context.Context) error {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for CloseConnetion")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Connect provides a mock function with given fields: _a0, _a1
func (_m *Repository) Connect(_a0 context.Context, _a1 *config.Config) error {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for Connect")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *config.Config) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetToken provides a mock function with given fields: _a0, _a1
func (_m *Repository) GetToken(_a0 context.Context, _a1 models.RefreshToken) (*models.RefreshToken, error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for GetToken")
	}

	var r0 *models.RefreshToken
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, models.RefreshToken) (*models.RefreshToken, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, models.RefreshToken) *models.RefreshToken); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.RefreshToken)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, models.RefreshToken) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// StoreToken provides a mock function with given fields: _a0, _a1
func (_m *Repository) StoreToken(_a0 context.Context, _a1 models.RefreshToken) error {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for StoreToken")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, models.RefreshToken) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateToken provides a mock function with given fields: _a0, _a1
func (_m *Repository) UpdateToken(_a0 context.Context, _a1 models.RefreshToken) error {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for UpdateToken")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, models.RefreshToken) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewRepository creates a new instance of Repository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *Repository {
	mock := &Repository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
