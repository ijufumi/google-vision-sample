// Code generated by mockery. DO NOT EDIT.

package stubs

import mock "github.com/stretchr/testify/mock"

// HMACKeyOption is an autogenerated mock type for the HMACKeyOption type
type HMACKeyOption struct {
	mock.Mock
}

// withHMACKeyDesc provides a mock function with given fields: _a0
func (_m *HMACKeyOption) withHMACKeyDesc(_a0 *storage.hmacKeyDesc) {
	_m.Called(_a0)
}

// NewHMACKeyOption creates a new instance of HMACKeyOption. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewHMACKeyOption(t interface {
	mock.TestingT
	Cleanup(func())
}) *HMACKeyOption {
	mock := &HMACKeyOption{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}