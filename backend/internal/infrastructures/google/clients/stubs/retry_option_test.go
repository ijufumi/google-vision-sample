// Code generated by mockery. DO NOT EDIT.

package stubs

import mock "github.com/stretchr/testify/mock"

// RetryOption is an autogenerated mock type for the RetryOption type
type RetryOption struct {
	mock.Mock
}

// apply provides a mock function with given fields: config
func (_m *RetryOption) apply(config *storage.retryConfig) {
	_m.Called(config)
}

// NewRetryOption creates a new instance of RetryOption. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewRetryOption(t interface {
	mock.TestingT
	Cleanup(func())
}) *RetryOption {
	mock := &RetryOption{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}