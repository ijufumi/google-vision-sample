// Code generated by mockery. DO NOT EDIT.

package mocks

import (
	entities "github.com/ijufumi/google-vision-sample/internal/models/entities"
	gorm "gorm.io/gorm"

	mock "github.com/stretchr/testify/mock"
)

// ExtractedTextRepository is an autogenerated mock type for the ExtractedTextRepository type
type ExtractedTextRepository struct {
	mock.Mock
}

// Create provides a mock function with given fields: db, entity
func (_m *ExtractedTextRepository) Create(db *gorm.DB, entity ...*entities.ExtractedText) error {
	_va := make([]interface{}, len(entity))
	for _i := range entity {
		_va[_i] = entity[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, db)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 error
	if rf, ok := ret.Get(0).(func(*gorm.DB, ...*entities.ExtractedText) error); ok {
		r0 = rf(db, entity...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteByOutputFileID provides a mock function with given fields: db, outputFileID
func (_m *ExtractedTextRepository) DeleteByOutputFileID(db *gorm.DB, outputFileID string) error {
	ret := _m.Called(db, outputFileID)

	var r0 error
	if rf, ok := ret.Get(0).(func(*gorm.DB, string) error); ok {
		r0 = rf(db, outputFileID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetByID provides a mock function with given fields: db, id
func (_m *ExtractedTextRepository) GetByID(db *gorm.DB, id string) (*entities.ExtractedText, error) {
	ret := _m.Called(db, id)

	var r0 *entities.ExtractedText
	var r1 error
	if rf, ok := ret.Get(0).(func(*gorm.DB, string) (*entities.ExtractedText, error)); ok {
		return rf(db, id)
	}
	if rf, ok := ret.Get(0).(func(*gorm.DB, string) *entities.ExtractedText); ok {
		r0 = rf(db, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entities.ExtractedText)
		}
	}

	if rf, ok := ret.Get(1).(func(*gorm.DB, string) error); ok {
		r1 = rf(db, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewExtractedTextRepository creates a new instance of ExtractedTextRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewExtractedTextRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *ExtractedTextRepository {
	mock := &ExtractedTextRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
