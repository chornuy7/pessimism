// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/base-org/pessimism/internal/client (interfaces: EthClientInterface)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	big "math/big"
	reflect "reflect"

	ethereum "github.com/ethereum/go-ethereum"
	common "github.com/ethereum/go-ethereum/common"
	types "github.com/ethereum/go-ethereum/core/types"
	gomock "github.com/golang/mock/gomock"
)

// MockEthClientInterface is a mock of EthClientInterface interface.
type MockEthClientInterface struct {
	ctrl     *gomock.Controller
	recorder *MockEthClientInterfaceMockRecorder
}

// MockEthClientInterfaceMockRecorder is the mock recorder for MockEthClientInterface.
type MockEthClientInterfaceMockRecorder struct {
	mock *MockEthClientInterface
}

// NewMockEthClientInterface creates a new mock instance.
func NewMockEthClientInterface(ctrl *gomock.Controller) *MockEthClientInterface {
	mock := &MockEthClientInterface{ctrl: ctrl}
	mock.recorder = &MockEthClientInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockEthClientInterface) EXPECT() *MockEthClientInterfaceMockRecorder {
	return m.recorder
}

// BalanceAt mocks base method.
func (m *MockEthClientInterface) BalanceAt(arg0 context.Context, arg1 common.Address, arg2 *big.Int) (*big.Int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BalanceAt", arg0, arg1, arg2)
	ret0, _ := ret[0].(*big.Int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BalanceAt indicates an expected call of BalanceAt.
func (mr *MockEthClientInterfaceMockRecorder) BalanceAt(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BalanceAt", reflect.TypeOf((*MockEthClientInterface)(nil).BalanceAt), arg0, arg1, arg2)
}

// BlockByNumber mocks base method.
func (m *MockEthClientInterface) BlockByNumber(arg0 context.Context, arg1 *big.Int) (*types.Block, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BlockByNumber", arg0, arg1)
	ret0, _ := ret[0].(*types.Block)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BlockByNumber indicates an expected call of BlockByNumber.
func (mr *MockEthClientInterfaceMockRecorder) BlockByNumber(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BlockByNumber", reflect.TypeOf((*MockEthClientInterface)(nil).BlockByNumber), arg0, arg1)
}

// FilterLogs mocks base method.
func (m *MockEthClientInterface) FilterLogs(arg0 context.Context, arg1 ethereum.FilterQuery) ([]types.Log, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FilterLogs", arg0, arg1)
	ret0, _ := ret[0].([]types.Log)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FilterLogs indicates an expected call of FilterLogs.
func (mr *MockEthClientInterfaceMockRecorder) FilterLogs(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FilterLogs", reflect.TypeOf((*MockEthClientInterface)(nil).FilterLogs), arg0, arg1)
}

// HeaderByNumber mocks base method.
func (m *MockEthClientInterface) HeaderByNumber(arg0 context.Context, arg1 *big.Int) (*types.Header, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HeaderByNumber", arg0, arg1)
	ret0, _ := ret[0].(*types.Header)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// HeaderByNumber indicates an expected call of HeaderByNumber.
func (mr *MockEthClientInterfaceMockRecorder) HeaderByNumber(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HeaderByNumber", reflect.TypeOf((*MockEthClientInterface)(nil).HeaderByNumber), arg0, arg1)
}
