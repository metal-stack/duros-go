// Code generated by mockery v2.26.1. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// isEvent_ComponentInfo is an autogenerated mock type for the isEvent_ComponentInfo type
type isEvent_ComponentInfo struct {
	mock.Mock
}

// isEvent_ComponentInfo provides a mock function with given fields:
func (_m *isEvent_ComponentInfo) isEvent_ComponentInfo() {
	_m.Called()
}

type mockConstructorTestingTnewIsEvent_ComponentInfo interface {
	mock.TestingT
	Cleanup(func())
}

// newIsEvent_ComponentInfo creates a new instance of isEvent_ComponentInfo. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func newIsEvent_ComponentInfo(t mockConstructorTestingTnewIsEvent_ComponentInfo) *isEvent_ComponentInfo {
	mock := &isEvent_ComponentInfo{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
