// Code generated by mockery v2.42.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// isCreatePolicyRequest_Policy is an autogenerated mock type for the isCreatePolicyRequest_Policy type
type isCreatePolicyRequest_Policy struct {
	mock.Mock
}

// isCreatePolicyRequest_Policy provides a mock function with given fields:
func (_m *isCreatePolicyRequest_Policy) isCreatePolicyRequest_Policy() {
	_m.Called()
}

// newIsCreatePolicyRequest_Policy creates a new instance of isCreatePolicyRequest_Policy. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func newIsCreatePolicyRequest_Policy(t interface {
	mock.TestingT
	Cleanup(func())
}) *isCreatePolicyRequest_Policy {
	mock := &isCreatePolicyRequest_Policy{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
