// Code generated by mockery v2.42.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// isSchedulePolicy_SchedulePolicies is an autogenerated mock type for the isSchedulePolicy_SchedulePolicies type
type isSchedulePolicy_SchedulePolicies struct {
	mock.Mock
}

// isSchedulePolicy_SchedulePolicies provides a mock function with given fields:
func (_m *isSchedulePolicy_SchedulePolicies) isSchedulePolicy_SchedulePolicies() {
	_m.Called()
}

// newIsSchedulePolicy_SchedulePolicies creates a new instance of isSchedulePolicy_SchedulePolicies. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func newIsSchedulePolicy_SchedulePolicies(t interface {
	mock.TestingT
	Cleanup(func())
}) *isSchedulePolicy_SchedulePolicies {
	mock := &isSchedulePolicy_SchedulePolicies{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
