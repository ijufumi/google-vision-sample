// Code generated by mockery. DO NOT EDIT.

package stubs

import mock "github.com/stretchr/testify/mock"

// storageOption is an autogenerated mock type for the storageOption type
type storageOption struct {
	mock.Mock
}

// Apply provides a mock function with given fields: s
func (_m *storageOption) Apply(s *storage.settings) {
	_m.Called(s)
}

// newStorageOption creates a new instance of storageOption. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func newStorageOption(t interface {
	mock.TestingT
	Cleanup(func())
}) *storageOption {
	mock := &storageOption{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
