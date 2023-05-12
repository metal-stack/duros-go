// Code generated by mockery v2.26.1. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// UnsafeDurosAPIServer is an autogenerated mock type for the UnsafeDurosAPIServer type
type UnsafeDurosAPIServer struct {
	mock.Mock
}

// mustEmbedUnimplementedDurosAPIServer provides a mock function with given fields:
func (_m *UnsafeDurosAPIServer) mustEmbedUnimplementedDurosAPIServer() {
	_m.Called()
}

type mockConstructorTestingTNewUnsafeDurosAPIServer interface {
	mock.TestingT
	Cleanup(func())
}

// NewUnsafeDurosAPIServer creates a new instance of UnsafeDurosAPIServer. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewUnsafeDurosAPIServer(t mockConstructorTestingTNewUnsafeDurosAPIServer) *UnsafeDurosAPIServer {
	mock := &UnsafeDurosAPIServer{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
