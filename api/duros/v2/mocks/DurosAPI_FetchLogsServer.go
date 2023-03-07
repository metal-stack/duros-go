// Code generated by mockery v2.21.1. DO NOT EDIT.

package mocks

import (
	context "context"

	httpbody "google.golang.org/genproto/googleapis/api/httpbody"
	metadata "google.golang.org/grpc/metadata"

	mock "github.com/stretchr/testify/mock"
)

// DurosAPI_FetchLogsServer is an autogenerated mock type for the DurosAPI_FetchLogsServer type
type DurosAPI_FetchLogsServer struct {
	mock.Mock
}

// Context provides a mock function with given fields:
func (_m *DurosAPI_FetchLogsServer) Context() context.Context {
	ret := _m.Called()

	var r0 context.Context
	if rf, ok := ret.Get(0).(func() context.Context); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(context.Context)
		}
	}

	return r0
}

// RecvMsg provides a mock function with given fields: m
func (_m *DurosAPI_FetchLogsServer) RecvMsg(m interface{}) error {
	ret := _m.Called(m)

	var r0 error
	if rf, ok := ret.Get(0).(func(interface{}) error); ok {
		r0 = rf(m)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Send provides a mock function with given fields: _a0
func (_m *DurosAPI_FetchLogsServer) Send(_a0 *httpbody.HttpBody) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(*httpbody.HttpBody) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SendHeader provides a mock function with given fields: _a0
func (_m *DurosAPI_FetchLogsServer) SendHeader(_a0 metadata.MD) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(metadata.MD) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SendMsg provides a mock function with given fields: m
func (_m *DurosAPI_FetchLogsServer) SendMsg(m interface{}) error {
	ret := _m.Called(m)

	var r0 error
	if rf, ok := ret.Get(0).(func(interface{}) error); ok {
		r0 = rf(m)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SetHeader provides a mock function with given fields: _a0
func (_m *DurosAPI_FetchLogsServer) SetHeader(_a0 metadata.MD) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(metadata.MD) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SetTrailer provides a mock function with given fields: _a0
func (_m *DurosAPI_FetchLogsServer) SetTrailer(_a0 metadata.MD) {
	_m.Called(_a0)
}

type mockConstructorTestingTNewDurosAPI_FetchLogsServer interface {
	mock.TestingT
	Cleanup(func())
}

// NewDurosAPI_FetchLogsServer creates a new instance of DurosAPI_FetchLogsServer. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewDurosAPI_FetchLogsServer(t mockConstructorTestingTNewDurosAPI_FetchLogsServer) *DurosAPI_FetchLogsServer {
	mock := &DurosAPI_FetchLogsServer{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
