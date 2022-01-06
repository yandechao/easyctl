// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import (
	fs "io/fs"

	logrus "github.com/sirupsen/logrus"
	mock "github.com/stretchr/testify/mock"

	runner "github.com/weiliang-ms/easyctl/pkg/runner"

	time "time"
)

// SftpInterface is an autogenerated mock type for the SftpInterface type
type SftpInterface struct {
	mock.Mock
}

// IOCopy64 provides a mock function with given fields: size, srcPath, dstPath, hostSign, logger
func (_m *SftpInterface) IOCopy64(size int64, srcPath string, dstPath string, hostSign string, logger *logrus.Logger) error {
	ret := _m.Called(size, srcPath, dstPath, hostSign, logger)

	var r0 error
	if rf, ok := ret.Get(0).(func(int64, string, string, string, *logrus.Logger) error); ok {
		r0 = rf(size, srcPath, dstPath, hostSign, logger)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewSftpClient provides a mock function with given fields: server, timeout
func (_m *SftpInterface) NewSftpClient(server runner.ServerInternal, timeout time.Duration) error {
	ret := _m.Called(server, timeout)

	var r0 error
	if rf, ok := ret.Get(0).(func(runner.ServerInternal, time.Duration) error); ok {
		r0 = rf(server, timeout)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SftpChmod provides a mock function with given fields: dstPath, mode
func (_m *SftpInterface) SftpChmod(dstPath string, mode fs.FileMode) error {
	ret := _m.Called(dstPath, mode)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, fs.FileMode) error); ok {
		r0 = rf(dstPath, mode)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SftpCreate provides a mock function with given fields: path
func (_m *SftpInterface) SftpCreate(path string) error {
	ret := _m.Called(path)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(path)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
