// Code generated by mockery 2.7.5. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"

	indexer "github.com/tendermint/tendermint/internal/state/indexer"

	query "github.com/tendermint/tendermint/internal/pubsub/query"

	tenderminttypes "github.com/tendermint/tendermint/types"

	types "github.com/tendermint/tendermint/abci/types"
)

// EventSink is an autogenerated mock type for the EventSink type
type EventSink struct {
	mock.Mock
}

// GetTxByHash provides a mock function with given fields: _a0
func (_m *EventSink) GetTxByHash(_a0 []byte) (*types.TxResult, error) {
	ret := _m.Called(_a0)

	var r0 *types.TxResult
	if rf, ok := ret.Get(0).(func([]byte) *types.TxResult); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.TxResult)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func([]byte) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// HasBlock provides a mock function with given fields: _a0
func (_m *EventSink) HasBlock(_a0 int64) (bool, error) {
	ret := _m.Called(_a0)

	var r0 bool
	if rf, ok := ret.Get(0).(func(int64) bool); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int64) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// IndexBlockEvents provides a mock function with given fields: _a0
func (_m *EventSink) IndexBlockEvents(_a0 tenderminttypes.EventDataNewBlockHeader) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(tenderminttypes.EventDataNewBlockHeader) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// IndexTxEvents provides a mock function with given fields: _a0
func (_m *EventSink) IndexTxEvents(_a0 []*types.TxResult) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func([]*types.TxResult) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SearchBlockEvents provides a mock function with given fields: _a0, _a1
func (_m *EventSink) SearchBlockEvents(_a0 context.Context, _a1 *query.Query) ([]int64, error) {
	ret := _m.Called(_a0, _a1)

	var r0 []int64
	if rf, ok := ret.Get(0).(func(context.Context, *query.Query) []int64); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]int64)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *query.Query) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SearchTxEvents provides a mock function with given fields: _a0, _a1
func (_m *EventSink) SearchTxEvents(_a0 context.Context, _a1 *query.Query) ([]*types.TxResult, error) {
	ret := _m.Called(_a0, _a1)

	var r0 []*types.TxResult
	if rf, ok := ret.Get(0).(func(context.Context, *query.Query) []*types.TxResult); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*types.TxResult)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *query.Query) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Stop provides a mock function with given fields:
func (_m *EventSink) Stop() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Type provides a mock function with given fields:
func (_m *EventSink) Type() indexer.EventSinkType {
	ret := _m.Called()

	var r0 indexer.EventSinkType
	if rf, ok := ret.Get(0).(func() indexer.EventSinkType); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(indexer.EventSinkType)
	}

	return r0
}
