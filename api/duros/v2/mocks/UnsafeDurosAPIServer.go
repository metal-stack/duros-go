// Code generated by mockery v2.41.0. DO NOT EDIT.

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

// NewUnsafeDurosAPIServer creates a new instance of UnsafeDurosAPIServer. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewUnsafeDurosAPIServer(t interface {
	mock.TestingT
	Cleanup(func())
}) *UnsafeDurosAPIServer {
	mock := &UnsafeDurosAPIServer{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
