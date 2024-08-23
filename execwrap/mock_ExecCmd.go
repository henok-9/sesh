// Code generated by mockery v2.45.0. DO NOT EDIT.

package execwrap

import mock "github.com/stretchr/testify/mock"

// MockExecCmd is an autogenerated mock type for the ExecCmd type
type MockExecCmd struct {
	mock.Mock
}

type MockExecCmd_Expecter struct {
	mock *mock.Mock
}

func (_m *MockExecCmd) EXPECT() *MockExecCmd_Expecter {
	return &MockExecCmd_Expecter{mock: &_m.Mock}
}

// CombinedOutput provides a mock function with given fields:
func (_m *MockExecCmd) CombinedOutput() ([]byte, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for CombinedOutput")
	}

	var r0 []byte
	var r1 error
	if rf, ok := ret.Get(0).(func() ([]byte, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() []byte); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockExecCmd_CombinedOutput_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CombinedOutput'
type MockExecCmd_CombinedOutput_Call struct {
	*mock.Call
}

// CombinedOutput is a helper method to define mock.On call
func (_e *MockExecCmd_Expecter) CombinedOutput() *MockExecCmd_CombinedOutput_Call {
	return &MockExecCmd_CombinedOutput_Call{Call: _e.mock.On("CombinedOutput")}
}

func (_c *MockExecCmd_CombinedOutput_Call) Run(run func()) *MockExecCmd_CombinedOutput_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockExecCmd_CombinedOutput_Call) Return(_a0 []byte, _a1 error) *MockExecCmd_CombinedOutput_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockExecCmd_CombinedOutput_Call) RunAndReturn(run func() ([]byte, error)) *MockExecCmd_CombinedOutput_Call {
	_c.Call.Return(run)
	return _c
}

// Output provides a mock function with given fields:
func (_m *MockExecCmd) Output() ([]byte, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Output")
	}

	var r0 []byte
	var r1 error
	if rf, ok := ret.Get(0).(func() ([]byte, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() []byte); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockExecCmd_Output_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Output'
type MockExecCmd_Output_Call struct {
	*mock.Call
}

// Output is a helper method to define mock.On call
func (_e *MockExecCmd_Expecter) Output() *MockExecCmd_Output_Call {
	return &MockExecCmd_Output_Call{Call: _e.mock.On("Output")}
}

func (_c *MockExecCmd_Output_Call) Run(run func()) *MockExecCmd_Output_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockExecCmd_Output_Call) Return(_a0 []byte, _a1 error) *MockExecCmd_Output_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockExecCmd_Output_Call) RunAndReturn(run func() ([]byte, error)) *MockExecCmd_Output_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockExecCmd creates a new instance of MockExecCmd. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockExecCmd(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockExecCmd {
	mock := &MockExecCmd{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
