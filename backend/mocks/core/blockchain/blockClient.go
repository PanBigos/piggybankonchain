// Code generated by mockery v2.40.1. DO NOT EDIT.

package blockchain

import (
	context "context"
	big "math/big"

	common "github.com/ethereum/go-ethereum/common"

	ethereum "github.com/ethereum/go-ethereum"

	mock "github.com/stretchr/testify/mock"

	types "github.com/ethereum/go-ethereum/core/types"
)

// blockClient is an autogenerated mock type for the blockClient type
type blockClient struct {
	mock.Mock
}

type blockClient_Expecter struct {
	mock *mock.Mock
}

func (_m *blockClient) EXPECT() *blockClient_Expecter {
	return &blockClient_Expecter{mock: &_m.Mock}
}

// BlockByNumber provides a mock function with given fields: ctx, number
func (_m *blockClient) BlockByNumber(ctx context.Context, number *big.Int) (*types.Block, error) {
	ret := _m.Called(ctx, number)

	if len(ret) == 0 {
		panic("no return value specified for BlockByNumber")
	}

	var r0 *types.Block
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *big.Int) (*types.Block, error)); ok {
		return rf(ctx, number)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *big.Int) *types.Block); ok {
		r0 = rf(ctx, number)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Block)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *big.Int) error); ok {
		r1 = rf(ctx, number)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// blockClient_BlockByNumber_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'BlockByNumber'
type blockClient_BlockByNumber_Call struct {
	*mock.Call
}

// BlockByNumber is a helper method to define mock.On call
//   - ctx context.Context
//   - number *big.Int
func (_e *blockClient_Expecter) BlockByNumber(ctx interface{}, number interface{}) *blockClient_BlockByNumber_Call {
	return &blockClient_BlockByNumber_Call{Call: _e.mock.On("BlockByNumber", ctx, number)}
}

func (_c *blockClient_BlockByNumber_Call) Run(run func(ctx context.Context, number *big.Int)) *blockClient_BlockByNumber_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*big.Int))
	})
	return _c
}

func (_c *blockClient_BlockByNumber_Call) Return(_a0 *types.Block, _a1 error) *blockClient_BlockByNumber_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *blockClient_BlockByNumber_Call) RunAndReturn(run func(context.Context, *big.Int) (*types.Block, error)) *blockClient_BlockByNumber_Call {
	_c.Call.Return(run)
	return _c
}

// ChainID provides a mock function with given fields: ctx
func (_m *blockClient) ChainID(ctx context.Context) (*big.Int, error) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for ChainID")
	}

	var r0 *big.Int
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) (*big.Int, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) *big.Int); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*big.Int)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// blockClient_ChainID_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ChainID'
type blockClient_ChainID_Call struct {
	*mock.Call
}

// ChainID is a helper method to define mock.On call
//   - ctx context.Context
func (_e *blockClient_Expecter) ChainID(ctx interface{}) *blockClient_ChainID_Call {
	return &blockClient_ChainID_Call{Call: _e.mock.On("ChainID", ctx)}
}

func (_c *blockClient_ChainID_Call) Run(run func(ctx context.Context)) *blockClient_ChainID_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *blockClient_ChainID_Call) Return(_a0 *big.Int, _a1 error) *blockClient_ChainID_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *blockClient_ChainID_Call) RunAndReturn(run func(context.Context) (*big.Int, error)) *blockClient_ChainID_Call {
	_c.Call.Return(run)
	return _c
}

// HeaderByHash provides a mock function with given fields: ctx, hash
func (_m *blockClient) HeaderByHash(ctx context.Context, hash common.Hash) (*types.Header, error) {
	ret := _m.Called(ctx, hash)

	if len(ret) == 0 {
		panic("no return value specified for HeaderByHash")
	}

	var r0 *types.Header
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, common.Hash) (*types.Header, error)); ok {
		return rf(ctx, hash)
	}
	if rf, ok := ret.Get(0).(func(context.Context, common.Hash) *types.Header); ok {
		r0 = rf(ctx, hash)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Header)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, common.Hash) error); ok {
		r1 = rf(ctx, hash)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// blockClient_HeaderByHash_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'HeaderByHash'
type blockClient_HeaderByHash_Call struct {
	*mock.Call
}

// HeaderByHash is a helper method to define mock.On call
//   - ctx context.Context
//   - hash common.Hash
func (_e *blockClient_Expecter) HeaderByHash(ctx interface{}, hash interface{}) *blockClient_HeaderByHash_Call {
	return &blockClient_HeaderByHash_Call{Call: _e.mock.On("HeaderByHash", ctx, hash)}
}

func (_c *blockClient_HeaderByHash_Call) Run(run func(ctx context.Context, hash common.Hash)) *blockClient_HeaderByHash_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(common.Hash))
	})
	return _c
}

func (_c *blockClient_HeaderByHash_Call) Return(_a0 *types.Header, _a1 error) *blockClient_HeaderByHash_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *blockClient_HeaderByHash_Call) RunAndReturn(run func(context.Context, common.Hash) (*types.Header, error)) *blockClient_HeaderByHash_Call {
	_c.Call.Return(run)
	return _c
}

// HeaderByNumber provides a mock function with given fields: ctx, number
func (_m *blockClient) HeaderByNumber(ctx context.Context, number *big.Int) (*types.Header, error) {
	ret := _m.Called(ctx, number)

	if len(ret) == 0 {
		panic("no return value specified for HeaderByNumber")
	}

	var r0 *types.Header
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *big.Int) (*types.Header, error)); ok {
		return rf(ctx, number)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *big.Int) *types.Header); ok {
		r0 = rf(ctx, number)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Header)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *big.Int) error); ok {
		r1 = rf(ctx, number)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// blockClient_HeaderByNumber_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'HeaderByNumber'
type blockClient_HeaderByNumber_Call struct {
	*mock.Call
}

// HeaderByNumber is a helper method to define mock.On call
//   - ctx context.Context
//   - number *big.Int
func (_e *blockClient_Expecter) HeaderByNumber(ctx interface{}, number interface{}) *blockClient_HeaderByNumber_Call {
	return &blockClient_HeaderByNumber_Call{Call: _e.mock.On("HeaderByNumber", ctx, number)}
}

func (_c *blockClient_HeaderByNumber_Call) Run(run func(ctx context.Context, number *big.Int)) *blockClient_HeaderByNumber_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*big.Int))
	})
	return _c
}

func (_c *blockClient_HeaderByNumber_Call) Return(_a0 *types.Header, _a1 error) *blockClient_HeaderByNumber_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *blockClient_HeaderByNumber_Call) RunAndReturn(run func(context.Context, *big.Int) (*types.Header, error)) *blockClient_HeaderByNumber_Call {
	_c.Call.Return(run)
	return _c
}

// SubscribeNewHead provides a mock function with given fields: ctx, ch
func (_m *blockClient) SubscribeNewHead(ctx context.Context, ch chan<- *types.Header) (ethereum.Subscription, error) {
	ret := _m.Called(ctx, ch)

	if len(ret) == 0 {
		panic("no return value specified for SubscribeNewHead")
	}

	var r0 ethereum.Subscription
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, chan<- *types.Header) (ethereum.Subscription, error)); ok {
		return rf(ctx, ch)
	}
	if rf, ok := ret.Get(0).(func(context.Context, chan<- *types.Header) ethereum.Subscription); ok {
		r0 = rf(ctx, ch)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(ethereum.Subscription)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, chan<- *types.Header) error); ok {
		r1 = rf(ctx, ch)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// blockClient_SubscribeNewHead_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SubscribeNewHead'
type blockClient_SubscribeNewHead_Call struct {
	*mock.Call
}

// SubscribeNewHead is a helper method to define mock.On call
//   - ctx context.Context
//   - ch chan<- *types.Header
func (_e *blockClient_Expecter) SubscribeNewHead(ctx interface{}, ch interface{}) *blockClient_SubscribeNewHead_Call {
	return &blockClient_SubscribeNewHead_Call{Call: _e.mock.On("SubscribeNewHead", ctx, ch)}
}

func (_c *blockClient_SubscribeNewHead_Call) Run(run func(ctx context.Context, ch chan<- *types.Header)) *blockClient_SubscribeNewHead_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(chan<- *types.Header))
	})
	return _c
}

func (_c *blockClient_SubscribeNewHead_Call) Return(_a0 ethereum.Subscription, _a1 error) *blockClient_SubscribeNewHead_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *blockClient_SubscribeNewHead_Call) RunAndReturn(run func(context.Context, chan<- *types.Header) (ethereum.Subscription, error)) *blockClient_SubscribeNewHead_Call {
	_c.Call.Return(run)
	return _c
}

// newBlockClient creates a new instance of blockClient. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func newBlockClient(t interface {
	mock.TestingT
	Cleanup(func())
}) *blockClient {
	mock := &blockClient{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}