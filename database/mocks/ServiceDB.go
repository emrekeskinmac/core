// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"
import service "github.com/mesg-foundation/core/service"

// ServiceDB is an autogenerated mock type for the ServiceDB type
type ServiceDB struct {
	mock.Mock
}

// All provides a mock function with given fields:
func (_m *ServiceDB) All() ([]*service.Service, error) {
	ret := _m.Called()

	var r0 []*service.Service
	if rf, ok := ret.Get(0).(func() []*service.Service); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*service.Service)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Close provides a mock function with given fields:
func (_m *ServiceDB) Close() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Delete provides a mock function with given fields: hashOrSid
func (_m *ServiceDB) Delete(hashOrSid string) error {
	ret := _m.Called(hashOrSid)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(hashOrSid)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Get provides a mock function with given fields: hashOrSid
func (_m *ServiceDB) Get(hashOrSid string) (*service.Service, error) {
	ret := _m.Called(hashOrSid)

	var r0 *service.Service
	if rf, ok := ret.Get(0).(func(string) *service.Service); ok {
		r0 = rf(hashOrSid)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*service.Service)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(hashOrSid)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Save provides a mock function with given fields: s
func (_m *ServiceDB) Save(s *service.Service) error {
	ret := _m.Called(s)

	var r0 error
	if rf, ok := ret.Get(0).(func(*service.Service) error); ok {
		r0 = rf(s)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
