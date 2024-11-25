// Code generated by mockery v2.49.0. DO NOT EDIT.

package mocks

import (
	core "github.com/johnfercher/maroto/v2/pkg/core"
	entity "github.com/johnfercher/maroto/v2/pkg/core/entity"

	mock "github.com/stretchr/testify/mock"

	node "github.com/johnfercher/go-tree/node"
)

// Maroto is an autogenerated mock type for the Maroto type
type Maroto struct {
	mock.Mock
}

type Maroto_Expecter struct {
	mock *mock.Mock
}

func (_m *Maroto) EXPECT() *Maroto_Expecter {
	return &Maroto_Expecter{mock: &_m.Mock}
}

// AddAutoRow provides a mock function with given fields: cols
func (_m *Maroto) AddAutoRow(cols ...core.Col) core.Row {
	_va := make([]interface{}, len(cols))
	for _i := range cols {
		_va[_i] = cols[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for AddAutoRow")
	}

	var r0 core.Row
	if rf, ok := ret.Get(0).(func(...core.Col) core.Row); ok {
		r0 = rf(cols...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(core.Row)
		}
	}

	return r0
}

// Maroto_AddAutoRow_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'AddAutoRow'
type Maroto_AddAutoRow_Call struct {
	*mock.Call
}

// AddAutoRow is a helper method to define mock.On call
//   - cols ...core.Col
func (_e *Maroto_Expecter) AddAutoRow(cols ...interface{}) *Maroto_AddAutoRow_Call {
	return &Maroto_AddAutoRow_Call{Call: _e.mock.On("AddAutoRow",
		append([]interface{}{}, cols...)...)}
}

func (_c *Maroto_AddAutoRow_Call) Run(run func(cols ...core.Col)) *Maroto_AddAutoRow_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]core.Col, len(args)-0)
		for i, a := range args[0:] {
			if a != nil {
				variadicArgs[i] = a.(core.Col)
			}
		}
		run(variadicArgs...)
	})
	return _c
}

func (_c *Maroto_AddAutoRow_Call) Return(_a0 core.Row) *Maroto_AddAutoRow_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Maroto_AddAutoRow_Call) RunAndReturn(run func(...core.Col) core.Row) *Maroto_AddAutoRow_Call {
	_c.Call.Return(run)
	return _c
}

// AddPages provides a mock function with given fields: pages
func (_m *Maroto) AddPages(pages ...core.Page) {
	_va := make([]interface{}, len(pages))
	for _i := range pages {
		_va[_i] = pages[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _va...)
	_m.Called(_ca...)
}

// Maroto_AddPages_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'AddPages'
type Maroto_AddPages_Call struct {
	*mock.Call
}

// AddPages is a helper method to define mock.On call
//   - pages ...core.Page
func (_e *Maroto_Expecter) AddPages(pages ...interface{}) *Maroto_AddPages_Call {
	return &Maroto_AddPages_Call{Call: _e.mock.On("AddPages",
		append([]interface{}{}, pages...)...)}
}

func (_c *Maroto_AddPages_Call) Run(run func(pages ...core.Page)) *Maroto_AddPages_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]core.Page, len(args)-0)
		for i, a := range args[0:] {
			if a != nil {
				variadicArgs[i] = a.(core.Page)
			}
		}
		run(variadicArgs...)
	})
	return _c
}

func (_c *Maroto_AddPages_Call) Return() *Maroto_AddPages_Call {
	_c.Call.Return()
	return _c
}

func (_c *Maroto_AddPages_Call) RunAndReturn(run func(...core.Page)) *Maroto_AddPages_Call {
	_c.Call.Return(run)
	return _c
}

// AddRow provides a mock function with given fields: rowHeight, cols
func (_m *Maroto) AddRow(rowHeight float64, cols ...core.Col) core.Row {
	_va := make([]interface{}, len(cols))
	for _i := range cols {
		_va[_i] = cols[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, rowHeight)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for AddRow")
	}

	var r0 core.Row
	if rf, ok := ret.Get(0).(func(float64, ...core.Col) core.Row); ok {
		r0 = rf(rowHeight, cols...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(core.Row)
		}
	}

	return r0
}

// Maroto_AddRow_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'AddRow'
type Maroto_AddRow_Call struct {
	*mock.Call
}

// AddRow is a helper method to define mock.On call
//   - rowHeight float64
//   - cols ...core.Col
func (_e *Maroto_Expecter) AddRow(rowHeight interface{}, cols ...interface{}) *Maroto_AddRow_Call {
	return &Maroto_AddRow_Call{Call: _e.mock.On("AddRow",
		append([]interface{}{rowHeight}, cols...)...)}
}

func (_c *Maroto_AddRow_Call) Run(run func(rowHeight float64, cols ...core.Col)) *Maroto_AddRow_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]core.Col, len(args)-1)
		for i, a := range args[1:] {
			if a != nil {
				variadicArgs[i] = a.(core.Col)
			}
		}
		run(args[0].(float64), variadicArgs...)
	})
	return _c
}

func (_c *Maroto_AddRow_Call) Return(_a0 core.Row) *Maroto_AddRow_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Maroto_AddRow_Call) RunAndReturn(run func(float64, ...core.Col) core.Row) *Maroto_AddRow_Call {
	_c.Call.Return(run)
	return _c
}

// AddRows provides a mock function with given fields: rows
func (_m *Maroto) AddRows(rows ...core.Row) {
	_va := make([]interface{}, len(rows))
	for _i := range rows {
		_va[_i] = rows[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _va...)
	_m.Called(_ca...)
}

// Maroto_AddRows_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'AddRows'
type Maroto_AddRows_Call struct {
	*mock.Call
}

// AddRows is a helper method to define mock.On call
//   - rows ...core.Row
func (_e *Maroto_Expecter) AddRows(rows ...interface{}) *Maroto_AddRows_Call {
	return &Maroto_AddRows_Call{Call: _e.mock.On("AddRows",
		append([]interface{}{}, rows...)...)}
}

func (_c *Maroto_AddRows_Call) Run(run func(rows ...core.Row)) *Maroto_AddRows_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]core.Row, len(args)-0)
		for i, a := range args[0:] {
			if a != nil {
				variadicArgs[i] = a.(core.Row)
			}
		}
		run(variadicArgs...)
	})
	return _c
}

func (_c *Maroto_AddRows_Call) Return() *Maroto_AddRows_Call {
	_c.Call.Return()
	return _c
}

func (_c *Maroto_AddRows_Call) RunAndReturn(run func(...core.Row)) *Maroto_AddRows_Call {
	_c.Call.Return(run)
	return _c
}

// FitlnCurrentPage provides a mock function with given fields: heightNewLine
func (_m *Maroto) FitlnCurrentPage(heightNewLine float64) bool {
	ret := _m.Called(heightNewLine)

	if len(ret) == 0 {
		panic("no return value specified for FitlnCurrentPage")
	}

	var r0 bool
	if rf, ok := ret.Get(0).(func(float64) bool); ok {
		r0 = rf(heightNewLine)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// Maroto_FitlnCurrentPage_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'FitlnCurrentPage'
type Maroto_FitlnCurrentPage_Call struct {
	*mock.Call
}

// FitlnCurrentPage is a helper method to define mock.On call
//   - heightNewLine float64
func (_e *Maroto_Expecter) FitlnCurrentPage(heightNewLine interface{}) *Maroto_FitlnCurrentPage_Call {
	return &Maroto_FitlnCurrentPage_Call{Call: _e.mock.On("FitlnCurrentPage", heightNewLine)}
}

func (_c *Maroto_FitlnCurrentPage_Call) Run(run func(heightNewLine float64)) *Maroto_FitlnCurrentPage_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(float64))
	})
	return _c
}

func (_c *Maroto_FitlnCurrentPage_Call) Return(_a0 bool) *Maroto_FitlnCurrentPage_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Maroto_FitlnCurrentPage_Call) RunAndReturn(run func(float64) bool) *Maroto_FitlnCurrentPage_Call {
	_c.Call.Return(run)
	return _c
}

// Generate provides a mock function with given fields:
func (_m *Maroto) Generate() (core.Document, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Generate")
	}

	var r0 core.Document
	var r1 error
	if rf, ok := ret.Get(0).(func() (core.Document, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() core.Document); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(core.Document)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Maroto_Generate_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Generate'
type Maroto_Generate_Call struct {
	*mock.Call
}

// Generate is a helper method to define mock.On call
func (_e *Maroto_Expecter) Generate() *Maroto_Generate_Call {
	return &Maroto_Generate_Call{Call: _e.mock.On("Generate")}
}

func (_c *Maroto_Generate_Call) Run(run func()) *Maroto_Generate_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Maroto_Generate_Call) Return(_a0 core.Document, _a1 error) *Maroto_Generate_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Maroto_Generate_Call) RunAndReturn(run func() (core.Document, error)) *Maroto_Generate_Call {
	_c.Call.Return(run)
	return _c
}

// GetCurrentConfig provides a mock function with given fields:
func (_m *Maroto) GetCurrentConfig() *entity.Config {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetCurrentConfig")
	}

	var r0 *entity.Config
	if rf, ok := ret.Get(0).(func() *entity.Config); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.Config)
		}
	}

	return r0
}

// Maroto_GetCurrentConfig_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetCurrentConfig'
type Maroto_GetCurrentConfig_Call struct {
	*mock.Call
}

// GetCurrentConfig is a helper method to define mock.On call
func (_e *Maroto_Expecter) GetCurrentConfig() *Maroto_GetCurrentConfig_Call {
	return &Maroto_GetCurrentConfig_Call{Call: _e.mock.On("GetCurrentConfig")}
}

func (_c *Maroto_GetCurrentConfig_Call) Run(run func()) *Maroto_GetCurrentConfig_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Maroto_GetCurrentConfig_Call) Return(_a0 *entity.Config) *Maroto_GetCurrentConfig_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Maroto_GetCurrentConfig_Call) RunAndReturn(run func() *entity.Config) *Maroto_GetCurrentConfig_Call {
	_c.Call.Return(run)
	return _c
}

// GetStructure provides a mock function with given fields:
func (_m *Maroto) GetStructure() *node.Node[core.Structure] {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetStructure")
	}

	var r0 *node.Node[core.Structure]
	if rf, ok := ret.Get(0).(func() *node.Node[core.Structure]); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*node.Node[core.Structure])
		}
	}

	return r0
}

// Maroto_GetStructure_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetStructure'
type Maroto_GetStructure_Call struct {
	*mock.Call
}

// GetStructure is a helper method to define mock.On call
func (_e *Maroto_Expecter) GetStructure() *Maroto_GetStructure_Call {
	return &Maroto_GetStructure_Call{Call: _e.mock.On("GetStructure")}
}

func (_c *Maroto_GetStructure_Call) Run(run func()) *Maroto_GetStructure_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Maroto_GetStructure_Call) Return(_a0 *node.Node[core.Structure]) *Maroto_GetStructure_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Maroto_GetStructure_Call) RunAndReturn(run func() *node.Node[core.Structure]) *Maroto_GetStructure_Call {
	_c.Call.Return(run)
	return _c
}

// RegisterFooter provides a mock function with given fields: rows
func (_m *Maroto) RegisterFooter(rows ...core.Row) error {
	_va := make([]interface{}, len(rows))
	for _i := range rows {
		_va[_i] = rows[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for RegisterFooter")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(...core.Row) error); ok {
		r0 = rf(rows...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Maroto_RegisterFooter_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'RegisterFooter'
type Maroto_RegisterFooter_Call struct {
	*mock.Call
}

// RegisterFooter is a helper method to define mock.On call
//   - rows ...core.Row
func (_e *Maroto_Expecter) RegisterFooter(rows ...interface{}) *Maroto_RegisterFooter_Call {
	return &Maroto_RegisterFooter_Call{Call: _e.mock.On("RegisterFooter",
		append([]interface{}{}, rows...)...)}
}

func (_c *Maroto_RegisterFooter_Call) Run(run func(rows ...core.Row)) *Maroto_RegisterFooter_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]core.Row, len(args)-0)
		for i, a := range args[0:] {
			if a != nil {
				variadicArgs[i] = a.(core.Row)
			}
		}
		run(variadicArgs...)
	})
	return _c
}

func (_c *Maroto_RegisterFooter_Call) Return(_a0 error) *Maroto_RegisterFooter_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Maroto_RegisterFooter_Call) RunAndReturn(run func(...core.Row) error) *Maroto_RegisterFooter_Call {
	_c.Call.Return(run)
	return _c
}

// RegisterHeader provides a mock function with given fields: rows
func (_m *Maroto) RegisterHeader(rows ...core.Row) error {
	_va := make([]interface{}, len(rows))
	for _i := range rows {
		_va[_i] = rows[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for RegisterHeader")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(...core.Row) error); ok {
		r0 = rf(rows...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Maroto_RegisterHeader_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'RegisterHeader'
type Maroto_RegisterHeader_Call struct {
	*mock.Call
}

// RegisterHeader is a helper method to define mock.On call
//   - rows ...core.Row
func (_e *Maroto_Expecter) RegisterHeader(rows ...interface{}) *Maroto_RegisterHeader_Call {
	return &Maroto_RegisterHeader_Call{Call: _e.mock.On("RegisterHeader",
		append([]interface{}{}, rows...)...)}
}

func (_c *Maroto_RegisterHeader_Call) Run(run func(rows ...core.Row)) *Maroto_RegisterHeader_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]core.Row, len(args)-0)
		for i, a := range args[0:] {
			if a != nil {
				variadicArgs[i] = a.(core.Row)
			}
		}
		run(variadicArgs...)
	})
	return _c
}

func (_c *Maroto_RegisterHeader_Call) Return(_a0 error) *Maroto_RegisterHeader_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Maroto_RegisterHeader_Call) RunAndReturn(run func(...core.Row) error) *Maroto_RegisterHeader_Call {
	_c.Call.Return(run)
	return _c
}

// NewMaroto creates a new instance of Maroto. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMaroto(t interface {
	mock.TestingT
	Cleanup(func())
},
) *Maroto {
	mock := &Maroto{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
