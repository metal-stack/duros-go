// Code generated by mockery v2.12.1. DO NOT EDIT.

package mocks

import (
	testing "testing"

	mock "github.com/stretchr/testify/mock"
)

// isQoSRateLimitPolicy_QoSLimit is an autogenerated mock type for the isQoSRateLimitPolicy_QoSLimit type
type isQoSRateLimitPolicy_QoSLimit struct {
	mock.Mock
}

// isQoSRateLimitPolicy_QoSLimit provides a mock function with given fields:
func (_m *isQoSRateLimitPolicy_QoSLimit) isQoSRateLimitPolicy_QoSLimit() {
	_m.Called()
}

// newIsQoSRateLimitPolicy_QoSLimit creates a new instance of isQoSRateLimitPolicy_QoSLimit. It also registers the testing.TB interface on the mock and a cleanup function to assert the mocks expectations.
func newIsQoSRateLimitPolicy_QoSLimit(t testing.TB) *isQoSRateLimitPolicy_QoSLimit {
	mock := &isQoSRateLimitPolicy_QoSLimit{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
