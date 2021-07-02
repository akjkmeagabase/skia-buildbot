// Code generated by mockery v2.4.0. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"
	config "go.skia.org/infra/skcq/go/config"

	time "time"
)

// ThrottlerManager is an autogenerated mock type for the ThrottlerManager type
type ThrottlerManager struct {
	mock.Mock
}

// Throttle provides a mock function with given fields: repoBranch, commitTime
func (_m *ThrottlerManager) Throttle(repoBranch string, commitTime time.Time) bool {
	ret := _m.Called(repoBranch, commitTime)

	var r0 bool
	if rf, ok := ret.Get(0).(func(string, time.Time) bool); ok {
		r0 = rf(repoBranch, commitTime)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// UpdateThrottler provides a mock function with given fields: repoBranch, commitTime, throttlerCfg
func (_m *ThrottlerManager) UpdateThrottler(repoBranch string, commitTime time.Time, throttlerCfg *config.ThrottlerCfg) {
	_m.Called(repoBranch, commitTime, throttlerCfg)
}
