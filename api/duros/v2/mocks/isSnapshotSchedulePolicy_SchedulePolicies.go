// Code generated by mockery v2.50.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// isSnapshotSchedulePolicy_SchedulePolicies is an autogenerated mock type for the isSnapshotSchedulePolicy_SchedulePolicies type
type isSnapshotSchedulePolicy_SchedulePolicies struct {
	mock.Mock
}

// isSnapshotSchedulePolicy_SchedulePolicies provides a mock function with no fields
func (_m *isSnapshotSchedulePolicy_SchedulePolicies) isSnapshotSchedulePolicy_SchedulePolicies() {
	_m.Called()
}

// newIsSnapshotSchedulePolicy_SchedulePolicies creates a new instance of isSnapshotSchedulePolicy_SchedulePolicies. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func newIsSnapshotSchedulePolicy_SchedulePolicies(t interface {
	mock.TestingT
	Cleanup(func())
}) *isSnapshotSchedulePolicy_SchedulePolicies {
	mock := &isSnapshotSchedulePolicy_SchedulePolicies{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
