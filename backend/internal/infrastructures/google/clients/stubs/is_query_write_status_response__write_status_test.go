// Code generated by mockery. DO NOT EDIT.

package stubs

import mock "github.com/stretchr/testify/mock"

// isQueryWriteStatusResponse_WriteStatus is an autogenerated mock type for the isQueryWriteStatusResponse_WriteStatus type
type isQueryWriteStatusResponse_WriteStatus struct {
	mock.Mock
}

// isQueryWriteStatusResponse_WriteStatus provides a mock function with given fields:
func (_m *isQueryWriteStatusResponse_WriteStatus) isQueryWriteStatusResponse_WriteStatus() {
	_m.Called()
}

// newIsQueryWriteStatusResponse_WriteStatus creates a new instance of isQueryWriteStatusResponse_WriteStatus. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func newIsQueryWriteStatusResponse_WriteStatus(t interface {
	mock.TestingT
	Cleanup(func())
}) *isQueryWriteStatusResponse_WriteStatus {
	mock := &isQueryWriteStatusResponse_WriteStatus{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
