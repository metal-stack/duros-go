// Code generated by mockery v2.41.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// isCreateVolumeRequest_QosPolicyID is an autogenerated mock type for the isCreateVolumeRequest_QosPolicyID type
type isCreateVolumeRequest_QosPolicyID struct {
	mock.Mock
}

// isCreateVolumeRequest_QosPolicyID provides a mock function with given fields:
func (_m *isCreateVolumeRequest_QosPolicyID) isCreateVolumeRequest_QosPolicyID() {
	_m.Called()
}

// newIsCreateVolumeRequest_QosPolicyID creates a new instance of isCreateVolumeRequest_QosPolicyID. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func newIsCreateVolumeRequest_QosPolicyID(t interface {
	mock.TestingT
	Cleanup(func())
}) *isCreateVolumeRequest_QosPolicyID {
	mock := &isCreateVolumeRequest_QosPolicyID{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
