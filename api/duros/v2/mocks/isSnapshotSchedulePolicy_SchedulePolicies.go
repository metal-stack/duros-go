// Code generated by mockery v2.26.1. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// isSnapshotSchedulePolicy_SchedulePolicies is an autogenerated mock type for the isSnapshotSchedulePolicy_SchedulePolicies type
type isSnapshotSchedulePolicy_SchedulePolicies struct {
	mock.Mock
}

// isSnapshotSchedulePolicy_SchedulePolicies provides a mock function with given fields:
func (_m *isSnapshotSchedulePolicy_SchedulePolicies) isSnapshotSchedulePolicy_SchedulePolicies() {
	_m.Called()
}

type mockConstructorTestingTnewIsSnapshotSchedulePolicy_SchedulePolicies interface {
	mock.TestingT
	Cleanup(func())
}

// newIsSnapshotSchedulePolicy_SchedulePolicies creates a new instance of isSnapshotSchedulePolicy_SchedulePolicies. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func newIsSnapshotSchedulePolicy_SchedulePolicies(t mockConstructorTestingTnewIsSnapshotSchedulePolicy_SchedulePolicies) *isSnapshotSchedulePolicy_SchedulePolicies {
	mock := &isSnapshotSchedulePolicy_SchedulePolicies{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
