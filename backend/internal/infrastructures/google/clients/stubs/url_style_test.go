// Code generated by mockery. DO NOT EDIT.

package stubs

import mock "github.com/stretchr/testify/mock"

// URLStyle is an autogenerated mock type for the URLStyle type
type URLStyle struct {
	mock.Mock
}

// host provides a mock function with given fields: hostname, bucket
func (_m *URLStyle) host(hostname string, bucket string) string {
	ret := _m.Called(hostname, bucket)

	var r0 string
	if rf, ok := ret.Get(0).(func(string, string) string); ok {
		r0 = rf(hostname, bucket)
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// path provides a mock function with given fields: bucket, object
func (_m *URLStyle) path(bucket string, object string) string {
	ret := _m.Called(bucket, object)

	var r0 string
	if rf, ok := ret.Get(0).(func(string, string) string); ok {
		r0 = rf(bucket, object)
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// NewURLStyle creates a new instance of URLStyle. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewURLStyle(t interface {
	mock.TestingT
	Cleanup(func())
}) *URLStyle {
	mock := &URLStyle{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
