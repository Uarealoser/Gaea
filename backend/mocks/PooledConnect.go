// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import (
	context "context"

	mysql "github.com/XiaoMi/Gaea/mysql"
	mock "github.com/stretchr/testify/mock"
)

// PooledConnect is an autogenerated mock type for the PooledConnect type
type PooledConnect struct {
	mock.Mock
}

// Begin provides a mock function with given fields:
func (_m *PooledConnect) Begin() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Close provides a mock function with given fields:
func (_m *PooledConnect) Close() {
	_m.Called()
}

// Commit provides a mock function with given fields:
func (_m *PooledConnect) Commit() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Execute provides a mock function with given fields: sql
func (_m *PooledConnect) Execute(sql string) (*mysql.Result, error) {
	ret := _m.Called(sql)

	var r0 *mysql.Result
	if rf, ok := ret.Get(0).(func(string) *mysql.Result); ok {
		r0 = rf(sql)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*mysql.Result)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(sql)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ExecuteWithCtx provides a mock function with given fields: ctx, sql
func (_m *PooledConnect) ExecuteWithCtx(ctx context.Context, sql string) (*mysql.Result, error) {
	ret := _m.Called(ctx, sql)

	var r0 *mysql.Result
	if rf, ok := ret.Get(0).(func(context.Context, string) *mysql.Result); ok {
		r0 = rf(ctx, sql)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*mysql.Result)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, sql)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FieldList provides a mock function with given fields: table, wildcard
func (_m *PooledConnect) FieldList(table string, wildcard string) ([]*mysql.Field, error) {
	ret := _m.Called(table, wildcard)

	var r0 []*mysql.Field
	if rf, ok := ret.Get(0).(func(string, string) []*mysql.Field); ok {
		r0 = rf(table, wildcard)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*mysql.Field)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(table, wildcard)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAddr provides a mock function with given fields:
func (_m *PooledConnect) GetAddr() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// IsClosed provides a mock function with given fields:
func (_m *PooledConnect) IsClosed() bool {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// Reconnect provides a mock function with given fields:
func (_m *PooledConnect) Reconnect() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Recycle provides a mock function with given fields:
func (_m *PooledConnect) Recycle() {
	_m.Called()
}

// Rollback provides a mock function with given fields:
func (_m *PooledConnect) Rollback() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SetAutoCommit provides a mock function with given fields: v
func (_m *PooledConnect) SetAutoCommit(v uint8) error {
	ret := _m.Called(v)

	var r0 error
	if rf, ok := ret.Get(0).(func(uint8) error); ok {
		r0 = rf(v)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SetCharset provides a mock function with given fields: charset, collation
func (_m *PooledConnect) SetCharset(charset string, collation mysql.CollationID) (bool, error) {
	ret := _m.Called(charset, collation)

	var r0 bool
	if rf, ok := ret.Get(0).(func(string, mysql.CollationID) bool); ok {
		r0 = rf(charset, collation)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, mysql.CollationID) error); ok {
		r1 = rf(charset, collation)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SetSessionVariables provides a mock function with given fields: frontend
func (_m *PooledConnect) SetSessionVariables(frontend *mysql.SessionVariables) (bool, error) {
	ret := _m.Called(frontend)

	var r0 bool
	if rf, ok := ret.Get(0).(func(*mysql.SessionVariables) bool); ok {
		r0 = rf(frontend)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*mysql.SessionVariables) error); ok {
		r1 = rf(frontend)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UseDB provides a mock function with given fields: db
func (_m *PooledConnect) UseDB(db string) error {
	ret := _m.Called(db)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(db)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// WriteSetStatement provides a mock function with given fields:
func (_m *PooledConnect) WriteSetStatement() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
