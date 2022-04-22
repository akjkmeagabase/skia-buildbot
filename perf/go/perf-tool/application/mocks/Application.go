// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"
	config "go.skia.org/infra/perf/go/config"

	testing "testing"

	tracestore "go.skia.org/infra/perf/go/tracestore"

	types "go.skia.org/infra/perf/go/types"
)

// Application is an autogenerated mock type for the Application type
type Application struct {
	mock.Mock
}

// ConfigCreatePubSubTopics provides a mock function with given fields: instanceConfig
func (_m *Application) ConfigCreatePubSubTopics(instanceConfig *config.InstanceConfig) error {
	ret := _m.Called(instanceConfig)

	var r0 error
	if rf, ok := ret.Get(0).(func(*config.InstanceConfig) error); ok {
		r0 = rf(instanceConfig)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DatabaseBackupAlerts provides a mock function with given fields: local, instanceConfig, outputFile
func (_m *Application) DatabaseBackupAlerts(local bool, instanceConfig *config.InstanceConfig, outputFile string) error {
	ret := _m.Called(local, instanceConfig, outputFile)

	var r0 error
	if rf, ok := ret.Get(0).(func(bool, *config.InstanceConfig, string) error); ok {
		r0 = rf(local, instanceConfig, outputFile)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DatabaseBackupRegressions provides a mock function with given fields: local, instanceConfig, outputFile, backupTo
func (_m *Application) DatabaseBackupRegressions(local bool, instanceConfig *config.InstanceConfig, outputFile string, backupTo string) error {
	ret := _m.Called(local, instanceConfig, outputFile, backupTo)

	var r0 error
	if rf, ok := ret.Get(0).(func(bool, *config.InstanceConfig, string, string) error); ok {
		r0 = rf(local, instanceConfig, outputFile, backupTo)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DatabaseBackupShortcuts provides a mock function with given fields: local, instanceConfig, outputFile
func (_m *Application) DatabaseBackupShortcuts(local bool, instanceConfig *config.InstanceConfig, outputFile string) error {
	ret := _m.Called(local, instanceConfig, outputFile)

	var r0 error
	if rf, ok := ret.Get(0).(func(bool, *config.InstanceConfig, string) error); ok {
		r0 = rf(local, instanceConfig, outputFile)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DatabaseMigrate provides a mock function with given fields: instanceConfig
func (_m *Application) DatabaseMigrate(instanceConfig *config.InstanceConfig) error {
	ret := _m.Called(instanceConfig)

	var r0 error
	if rf, ok := ret.Get(0).(func(*config.InstanceConfig) error); ok {
		r0 = rf(instanceConfig)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DatabaseRestoreAlerts provides a mock function with given fields: local, instanceConfig, inputFile
func (_m *Application) DatabaseRestoreAlerts(local bool, instanceConfig *config.InstanceConfig, inputFile string) error {
	ret := _m.Called(local, instanceConfig, inputFile)

	var r0 error
	if rf, ok := ret.Get(0).(func(bool, *config.InstanceConfig, string) error); ok {
		r0 = rf(local, instanceConfig, inputFile)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DatabaseRestoreRegressions provides a mock function with given fields: local, instanceConfig, inputFile
func (_m *Application) DatabaseRestoreRegressions(local bool, instanceConfig *config.InstanceConfig, inputFile string) error {
	ret := _m.Called(local, instanceConfig, inputFile)

	var r0 error
	if rf, ok := ret.Get(0).(func(bool, *config.InstanceConfig, string) error); ok {
		r0 = rf(local, instanceConfig, inputFile)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DatabaseRestoreShortcuts provides a mock function with given fields: local, instanceConfig, inputFile
func (_m *Application) DatabaseRestoreShortcuts(local bool, instanceConfig *config.InstanceConfig, inputFile string) error {
	ret := _m.Called(local, instanceConfig, inputFile)

	var r0 error
	if rf, ok := ret.Get(0).(func(bool, *config.InstanceConfig, string) error); ok {
		r0 = rf(local, instanceConfig, inputFile)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// IngestForceReingest provides a mock function with given fields: local, instanceConfig, start, stop, dryrun
func (_m *Application) IngestForceReingest(local bool, instanceConfig *config.InstanceConfig, start string, stop string, dryrun bool) error {
	ret := _m.Called(local, instanceConfig, start, stop, dryrun)

	var r0 error
	if rf, ok := ret.Get(0).(func(bool, *config.InstanceConfig, string, string, bool) error); ok {
		r0 = rf(local, instanceConfig, start, stop, dryrun)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// IngestValidate provides a mock function with given fields: inputFile, verbose
func (_m *Application) IngestValidate(inputFile string, verbose bool) error {
	ret := _m.Called(inputFile, verbose)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, bool) error); ok {
		r0 = rf(inputFile, verbose)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// TilesLast provides a mock function with given fields: store
func (_m *Application) TilesLast(store tracestore.TraceStore) error {
	ret := _m.Called(store)

	var r0 error
	if rf, ok := ret.Get(0).(func(tracestore.TraceStore) error); ok {
		r0 = rf(store)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// TilesList provides a mock function with given fields: store, num
func (_m *Application) TilesList(store tracestore.TraceStore, num int) error {
	ret := _m.Called(store, num)

	var r0 error
	if rf, ok := ret.Get(0).(func(tracestore.TraceStore, int) error); ok {
		r0 = rf(store, num)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// TracesExport provides a mock function with given fields: store, queryString, begin, end, outputFile
func (_m *Application) TracesExport(store tracestore.TraceStore, queryString string, begin types.CommitNumber, end types.CommitNumber, outputFile string) error {
	ret := _m.Called(store, queryString, begin, end, outputFile)

	var r0 error
	if rf, ok := ret.Get(0).(func(tracestore.TraceStore, string, types.CommitNumber, types.CommitNumber, string) error); ok {
		r0 = rf(store, queryString, begin, end, outputFile)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// TracesList provides a mock function with given fields: store, queryString, tileNumber
func (_m *Application) TracesList(store tracestore.TraceStore, queryString string, tileNumber types.TileNumber) error {
	ret := _m.Called(store, queryString, tileNumber)

	var r0 error
	if rf, ok := ret.Get(0).(func(tracestore.TraceStore, string, types.TileNumber) error); ok {
		r0 = rf(store, queryString, tileNumber)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// TrybotReference provides a mock function with given fields: local, store, instanceConfig, trybotFilename, outputFilename, numCommits
func (_m *Application) TrybotReference(local bool, store tracestore.TraceStore, instanceConfig *config.InstanceConfig, trybotFilename string, outputFilename string, numCommits int) error {
	ret := _m.Called(local, store, instanceConfig, trybotFilename, outputFilename, numCommits)

	var r0 error
	if rf, ok := ret.Get(0).(func(bool, tracestore.TraceStore, *config.InstanceConfig, string, string, int) error); ok {
		r0 = rf(local, store, instanceConfig, trybotFilename, outputFilename, numCommits)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewApplication creates a new instance of Application. It also registers a cleanup function to assert the mocks expectations.
func NewApplication(t testing.TB) *Application {
	mock := &Application{}

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
