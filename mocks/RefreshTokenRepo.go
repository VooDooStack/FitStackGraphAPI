// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
	fitstackapi "github.com/voodoostack/fitstackapi"
)

// RefreshTokenRepo is an autogenerated mock type for the RefreshTokenRepo type
type RefreshTokenRepo struct {
	mock.Mock
}

// Create provides a mock function with given fields: ctx, params
func (_m *RefreshTokenRepo) Create(ctx context.Context, params fitstackapi.CreateRefreshTokenParams) (fitstackapi.RefreshToken, error) {
	ret := _m.Called(ctx, params)

	var r0 fitstackapi.RefreshToken
	if rf, ok := ret.Get(0).(func(context.Context, fitstackapi.CreateRefreshTokenParams) fitstackapi.RefreshToken); ok {
		r0 = rf(ctx, params)
	} else {
		r0 = ret.Get(0).(fitstackapi.RefreshToken)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, fitstackapi.CreateRefreshTokenParams) error); ok {
		r1 = rf(ctx, params)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByID provides a mock function with given fields: ctx, id
func (_m *RefreshTokenRepo) GetByID(ctx context.Context, id string) (fitstackapi.RefreshToken, error) {
	ret := _m.Called(ctx, id)

	var r0 fitstackapi.RefreshToken
	if rf, ok := ret.Get(0).(func(context.Context, string) fitstackapi.RefreshToken); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(fitstackapi.RefreshToken)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
