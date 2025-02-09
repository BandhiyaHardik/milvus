// Code generated by mockery v2.32.4. DO NOT EDIT.

package metacache

import mock "github.com/stretchr/testify/mock"

// MockMetaCache is an autogenerated mock type for the MetaCache type
type MockMetaCache struct {
	mock.Mock
}

type MockMetaCache_Expecter struct {
	mock *mock.Mock
}

func (_m *MockMetaCache) EXPECT() *MockMetaCache_Expecter {
	return &MockMetaCache_Expecter{mock: &_m.Mock}
}

// CompactSegments provides a mock function with given fields: newSegmentID, partitionID, oldSegmentIDs
func (_m *MockMetaCache) CompactSegments(newSegmentID int64, partitionID int64, oldSegmentIDs ...int64) {
	_va := make([]interface{}, len(oldSegmentIDs))
	for _i := range oldSegmentIDs {
		_va[_i] = oldSegmentIDs[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, newSegmentID, partitionID)
	_ca = append(_ca, _va...)
	_m.Called(_ca...)
}

// MockMetaCache_CompactSegments_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CompactSegments'
type MockMetaCache_CompactSegments_Call struct {
	*mock.Call
}

// CompactSegments is a helper method to define mock.On call
//   - newSegmentID int64
//   - partitionID int64
//   - oldSegmentIDs ...int64
func (_e *MockMetaCache_Expecter) CompactSegments(newSegmentID interface{}, partitionID interface{}, oldSegmentIDs ...interface{}) *MockMetaCache_CompactSegments_Call {
	return &MockMetaCache_CompactSegments_Call{Call: _e.mock.On("CompactSegments",
		append([]interface{}{newSegmentID, partitionID}, oldSegmentIDs...)...)}
}

func (_c *MockMetaCache_CompactSegments_Call) Run(run func(newSegmentID int64, partitionID int64, oldSegmentIDs ...int64)) *MockMetaCache_CompactSegments_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]int64, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(int64)
			}
		}
		run(args[0].(int64), args[1].(int64), variadicArgs...)
	})
	return _c
}

func (_c *MockMetaCache_CompactSegments_Call) Return() *MockMetaCache_CompactSegments_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockMetaCache_CompactSegments_Call) RunAndReturn(run func(int64, int64, ...int64)) *MockMetaCache_CompactSegments_Call {
	_c.Call.Return(run)
	return _c
}

// GetSegmentIDsBy provides a mock function with given fields: filters
func (_m *MockMetaCache) GetSegmentIDsBy(filters ...SegmentFilter) []int64 {
	_va := make([]interface{}, len(filters))
	for _i := range filters {
		_va[_i] = filters[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 []int64
	if rf, ok := ret.Get(0).(func(...SegmentFilter) []int64); ok {
		r0 = rf(filters...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]int64)
		}
	}

	return r0
}

// MockMetaCache_GetSegmentIDsBy_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetSegmentIDsBy'
type MockMetaCache_GetSegmentIDsBy_Call struct {
	*mock.Call
}

// GetSegmentIDsBy is a helper method to define mock.On call
//   - filters ...SegmentFilter
func (_e *MockMetaCache_Expecter) GetSegmentIDsBy(filters ...interface{}) *MockMetaCache_GetSegmentIDsBy_Call {
	return &MockMetaCache_GetSegmentIDsBy_Call{Call: _e.mock.On("GetSegmentIDsBy",
		append([]interface{}{}, filters...)...)}
}

func (_c *MockMetaCache_GetSegmentIDsBy_Call) Run(run func(filters ...SegmentFilter)) *MockMetaCache_GetSegmentIDsBy_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]SegmentFilter, len(args)-0)
		for i, a := range args[0:] {
			if a != nil {
				variadicArgs[i] = a.(SegmentFilter)
			}
		}
		run(variadicArgs...)
	})
	return _c
}

func (_c *MockMetaCache_GetSegmentIDsBy_Call) Return(_a0 []int64) *MockMetaCache_GetSegmentIDsBy_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockMetaCache_GetSegmentIDsBy_Call) RunAndReturn(run func(...SegmentFilter) []int64) *MockMetaCache_GetSegmentIDsBy_Call {
	_c.Call.Return(run)
	return _c
}

// GetSegmentsBy provides a mock function with given fields: filters
func (_m *MockMetaCache) GetSegmentsBy(filters ...SegmentFilter) []*SegmentInfo {
	_va := make([]interface{}, len(filters))
	for _i := range filters {
		_va[_i] = filters[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 []*SegmentInfo
	if rf, ok := ret.Get(0).(func(...SegmentFilter) []*SegmentInfo); ok {
		r0 = rf(filters...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*SegmentInfo)
		}
	}

	return r0
}

// MockMetaCache_GetSegmentsBy_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetSegmentsBy'
type MockMetaCache_GetSegmentsBy_Call struct {
	*mock.Call
}

// GetSegmentsBy is a helper method to define mock.On call
//   - filters ...SegmentFilter
func (_e *MockMetaCache_Expecter) GetSegmentsBy(filters ...interface{}) *MockMetaCache_GetSegmentsBy_Call {
	return &MockMetaCache_GetSegmentsBy_Call{Call: _e.mock.On("GetSegmentsBy",
		append([]interface{}{}, filters...)...)}
}

func (_c *MockMetaCache_GetSegmentsBy_Call) Run(run func(filters ...SegmentFilter)) *MockMetaCache_GetSegmentsBy_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]SegmentFilter, len(args)-0)
		for i, a := range args[0:] {
			if a != nil {
				variadicArgs[i] = a.(SegmentFilter)
			}
		}
		run(variadicArgs...)
	})
	return _c
}

func (_c *MockMetaCache_GetSegmentsBy_Call) Return(_a0 []*SegmentInfo) *MockMetaCache_GetSegmentsBy_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockMetaCache_GetSegmentsBy_Call) RunAndReturn(run func(...SegmentFilter) []*SegmentInfo) *MockMetaCache_GetSegmentsBy_Call {
	_c.Call.Return(run)
	return _c
}

// NewSegment provides a mock function with given fields: segmentID, partitionID
func (_m *MockMetaCache) NewSegment(segmentID int64, partitionID int64) {
	_m.Called(segmentID, partitionID)
}

// MockMetaCache_NewSegment_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'NewSegment'
type MockMetaCache_NewSegment_Call struct {
	*mock.Call
}

// NewSegment is a helper method to define mock.On call
//   - segmentID int64
//   - partitionID int64
func (_e *MockMetaCache_Expecter) NewSegment(segmentID interface{}, partitionID interface{}) *MockMetaCache_NewSegment_Call {
	return &MockMetaCache_NewSegment_Call{Call: _e.mock.On("NewSegment", segmentID, partitionID)}
}

func (_c *MockMetaCache_NewSegment_Call) Run(run func(segmentID int64, partitionID int64)) *MockMetaCache_NewSegment_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(int64), args[1].(int64))
	})
	return _c
}

func (_c *MockMetaCache_NewSegment_Call) Return() *MockMetaCache_NewSegment_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockMetaCache_NewSegment_Call) RunAndReturn(run func(int64, int64)) *MockMetaCache_NewSegment_Call {
	_c.Call.Return(run)
	return _c
}

// UpdateSegments provides a mock function with given fields: action, filters
func (_m *MockMetaCache) UpdateSegments(action SegmentAction, filters ...SegmentFilter) {
	_va := make([]interface{}, len(filters))
	for _i := range filters {
		_va[_i] = filters[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, action)
	_ca = append(_ca, _va...)
	_m.Called(_ca...)
}

// MockMetaCache_UpdateSegments_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UpdateSegments'
type MockMetaCache_UpdateSegments_Call struct {
	*mock.Call
}

// UpdateSegments is a helper method to define mock.On call
//   - action SegmentAction
//   - filters ...SegmentFilter
func (_e *MockMetaCache_Expecter) UpdateSegments(action interface{}, filters ...interface{}) *MockMetaCache_UpdateSegments_Call {
	return &MockMetaCache_UpdateSegments_Call{Call: _e.mock.On("UpdateSegments",
		append([]interface{}{action}, filters...)...)}
}

func (_c *MockMetaCache_UpdateSegments_Call) Run(run func(action SegmentAction, filters ...SegmentFilter)) *MockMetaCache_UpdateSegments_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]SegmentFilter, len(args)-1)
		for i, a := range args[1:] {
			if a != nil {
				variadicArgs[i] = a.(SegmentFilter)
			}
		}
		run(args[0].(SegmentAction), variadicArgs...)
	})
	return _c
}

func (_c *MockMetaCache_UpdateSegments_Call) Return() *MockMetaCache_UpdateSegments_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockMetaCache_UpdateSegments_Call) RunAndReturn(run func(SegmentAction, ...SegmentFilter)) *MockMetaCache_UpdateSegments_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockMetaCache creates a new instance of MockMetaCache. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockMetaCache(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockMetaCache {
	mock := &MockMetaCache{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
