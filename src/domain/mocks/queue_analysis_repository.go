// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// QueueAnalysisRepository is an autogenerated mock type for the QueueAnalysisRepository type
type QueueAnalysisRepository struct {
	mock.Mock
}

// PublishMessage provides a mock function with given fields: message
func (_m *QueueAnalysisRepository) PublishMessage(message string) error {
	ret := _m.Called(message)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(message)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
