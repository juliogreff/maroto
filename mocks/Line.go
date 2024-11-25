// Code generated by mockery v2.49.0. DO NOT EDIT.

package mocks

import (
	entity "github.com/johnfercher/maroto/v2/pkg/core/entity"
	mock "github.com/stretchr/testify/mock"

	props "github.com/johnfercher/maroto/v2/pkg/props"
)

// Line is an autogenerated mock type for the Line type
type Line struct {
	mock.Mock
}

type Line_Expecter struct {
	mock *mock.Mock
}

func (_m *Line) EXPECT() *Line_Expecter {
	return &Line_Expecter{mock: &_m.Mock}
}

// Add provides a mock function with given fields: cell, prop
func (_m *Line) Add(cell *entity.Cell, prop *props.Line) {
	_m.Called(cell, prop)
}

// Line_Add_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Add'
type Line_Add_Call struct {
	*mock.Call
}

// Add is a helper method to define mock.On call
//   - cell *entity.Cell
//   - prop *props.Line
func (_e *Line_Expecter) Add(cell interface{}, prop interface{}) *Line_Add_Call {
	return &Line_Add_Call{Call: _e.mock.On("Add", cell, prop)}
}

func (_c *Line_Add_Call) Run(run func(cell *entity.Cell, prop *props.Line)) *Line_Add_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*entity.Cell), args[1].(*props.Line))
	})
	return _c
}

func (_c *Line_Add_Call) Return() *Line_Add_Call {
	_c.Call.Return()
	return _c
}

func (_c *Line_Add_Call) RunAndReturn(run func(*entity.Cell, *props.Line)) *Line_Add_Call {
	_c.Call.Return(run)
	return _c
}

// NewLine creates a new instance of Line. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewLine(t interface {
	mock.TestingT
	Cleanup(func())
},
) *Line {
	mock := &Line{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
