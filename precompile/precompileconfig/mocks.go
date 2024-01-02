// Code generated by MockGen. DO NOT EDIT.
// Source: config.go

// Package precompileconfig is a generated GoMock package.
package precompileconfig

import (
	reflect "reflect"

	atomic "github.com/ava-labs/avalanchego/chains/atomic"
	ids "github.com/ava-labs/avalanchego/ids"
	warp "github.com/ava-labs/avalanchego/vms/platformvm/warp"
	common "github.com/ethereum/go-ethereum/common"
	gomock "go.uber.org/mock/gomock"
)

// MockConfig is a mock of Config interface.
type MockConfig struct {
	ctrl     *gomock.Controller
	recorder *MockConfigMockRecorder
}

// MockConfigMockRecorder is the mock recorder for MockConfig.
type MockConfigMockRecorder struct {
	mock *MockConfig
}

// NewMockConfig creates a new mock instance.
func NewMockConfig(ctrl *gomock.Controller) *MockConfig {
	mock := &MockConfig{ctrl: ctrl}
	mock.recorder = &MockConfigMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockConfig) EXPECT() *MockConfigMockRecorder {
	return m.recorder
}

// Equal mocks base method.
func (m *MockConfig) Equal(arg0 Config) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Equal", arg0)
	ret0, _ := ret[0].(bool)
	return ret0
}

// Equal indicates an expected call of Equal.
func (mr *MockConfigMockRecorder) Equal(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Equal", reflect.TypeOf((*MockConfig)(nil).Equal), arg0)
}

// IsDisabled mocks base method.
func (m *MockConfig) IsDisabled() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsDisabled")
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsDisabled indicates an expected call of IsDisabled.
func (mr *MockConfigMockRecorder) IsDisabled() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsDisabled", reflect.TypeOf((*MockConfig)(nil).IsDisabled))
}

// Key mocks base method.
func (m *MockConfig) Key() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Key")
	ret0, _ := ret[0].(string)
	return ret0
}

// Key indicates an expected call of Key.
func (mr *MockConfigMockRecorder) Key() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Key", reflect.TypeOf((*MockConfig)(nil).Key))
}

// Timestamp mocks base method.
func (m *MockConfig) Timestamp() *uint64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Timestamp")
	ret0, _ := ret[0].(*uint64)
	return ret0
}

// Timestamp indicates an expected call of Timestamp.
func (mr *MockConfigMockRecorder) Timestamp() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Timestamp", reflect.TypeOf((*MockConfig)(nil).Timestamp))
}

// Verify mocks base method.
func (m *MockConfig) Verify(arg0 ChainConfig) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Verify", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Verify indicates an expected call of Verify.
func (mr *MockConfigMockRecorder) Verify(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Verify", reflect.TypeOf((*MockConfig)(nil).Verify), arg0)
}

// MockPredicater is a mock of Predicater interface.
type MockPredicater struct {
	ctrl     *gomock.Controller
	recorder *MockPredicaterMockRecorder
}

// MockPredicaterMockRecorder is the mock recorder for MockPredicater.
type MockPredicaterMockRecorder struct {
	mock *MockPredicater
}

// NewMockPredicater creates a new mock instance.
func NewMockPredicater(ctrl *gomock.Controller) *MockPredicater {
	mock := &MockPredicater{ctrl: ctrl}
	mock.recorder = &MockPredicaterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPredicater) EXPECT() *MockPredicaterMockRecorder {
	return m.recorder
}

// PredicateGas mocks base method.
func (m *MockPredicater) PredicateGas(storageSlots []byte) (uint64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PredicateGas", storageSlots)
	ret0, _ := ret[0].(uint64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PredicateGas indicates an expected call of PredicateGas.
func (mr *MockPredicaterMockRecorder) PredicateGas(storageSlots interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PredicateGas", reflect.TypeOf((*MockPredicater)(nil).PredicateGas), storageSlots)
}

// VerifyPredicate mocks base method.
func (m *MockPredicater) VerifyPredicate(arg0 *PredicateContext, arg1 []byte) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "VerifyPredicate", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// VerifyPredicate indicates an expected call of VerifyPredicate.
func (mr *MockPredicaterMockRecorder) VerifyPredicate(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VerifyPredicate", reflect.TypeOf((*MockPredicater)(nil).VerifyPredicate), arg0, arg1)
}

// MockSharedMemoryWriter is a mock of SharedMemoryWriter interface.
type MockSharedMemoryWriter struct {
	ctrl     *gomock.Controller
	recorder *MockSharedMemoryWriterMockRecorder
}

// MockSharedMemoryWriterMockRecorder is the mock recorder for MockSharedMemoryWriter.
type MockSharedMemoryWriterMockRecorder struct {
	mock *MockSharedMemoryWriter
}

// NewMockSharedMemoryWriter creates a new mock instance.
func NewMockSharedMemoryWriter(ctrl *gomock.Controller) *MockSharedMemoryWriter {
	mock := &MockSharedMemoryWriter{ctrl: ctrl}
	mock.recorder = &MockSharedMemoryWriterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSharedMemoryWriter) EXPECT() *MockSharedMemoryWriterMockRecorder {
	return m.recorder
}

// AddSharedMemoryRequests mocks base method.
func (m *MockSharedMemoryWriter) AddSharedMemoryRequests(chainID ids.ID, requests *atomic.Requests) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "AddSharedMemoryRequests", chainID, requests)
}

// AddSharedMemoryRequests indicates an expected call of AddSharedMemoryRequests.
func (mr *MockSharedMemoryWriterMockRecorder) AddSharedMemoryRequests(chainID, requests interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddSharedMemoryRequests", reflect.TypeOf((*MockSharedMemoryWriter)(nil).AddSharedMemoryRequests), chainID, requests)
}

// MockWarpMessageWriter is a mock of WarpMessageWriter interface.
type MockWarpMessageWriter struct {
	ctrl     *gomock.Controller
	recorder *MockWarpMessageWriterMockRecorder
}

// MockWarpMessageWriterMockRecorder is the mock recorder for MockWarpMessageWriter.
type MockWarpMessageWriterMockRecorder struct {
	mock *MockWarpMessageWriter
}

// NewMockWarpMessageWriter creates a new mock instance.
func NewMockWarpMessageWriter(ctrl *gomock.Controller) *MockWarpMessageWriter {
	mock := &MockWarpMessageWriter{ctrl: ctrl}
	mock.recorder = &MockWarpMessageWriterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockWarpMessageWriter) EXPECT() *MockWarpMessageWriterMockRecorder {
	return m.recorder
}

// AddMessage mocks base method.
func (m *MockWarpMessageWriter) AddMessage(unsignedMessage *warp.UnsignedMessage) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddMessage", unsignedMessage)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddMessage indicates an expected call of AddMessage.
func (mr *MockWarpMessageWriterMockRecorder) AddMessage(unsignedMessage interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddMessage", reflect.TypeOf((*MockWarpMessageWriter)(nil).AddMessage), unsignedMessage)
}

// MockAccepter is a mock of Accepter interface.
type MockAccepter struct {
	ctrl     *gomock.Controller
	recorder *MockAccepterMockRecorder
}

// MockAccepterMockRecorder is the mock recorder for MockAccepter.
type MockAccepterMockRecorder struct {
	mock *MockAccepter
}

// NewMockAccepter creates a new mock instance.
func NewMockAccepter(ctrl *gomock.Controller) *MockAccepter {
	mock := &MockAccepter{ctrl: ctrl}
	mock.recorder = &MockAccepterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAccepter) EXPECT() *MockAccepterMockRecorder {
	return m.recorder
}

// Accept mocks base method.
func (m *MockAccepter) Accept(acceptCtx *AcceptContext, blockHash common.Hash, blockNumber uint64, txHash common.Hash, logIndex int, topics []common.Hash, logData []byte) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Accept", acceptCtx, blockHash, blockNumber, txHash, logIndex, topics, logData)
	ret0, _ := ret[0].(error)
	return ret0
}

// Accept indicates an expected call of Accept.
func (mr *MockAccepterMockRecorder) Accept(acceptCtx, blockHash, blockNumber, txHash, logIndex, topics, logData interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Accept", reflect.TypeOf((*MockAccepter)(nil).Accept), acceptCtx, blockHash, blockNumber, txHash, logIndex, topics, logData)
}

// MockChainConfig is a mock of ChainConfig interface.
type MockChainConfig struct {
	ctrl     *gomock.Controller
	recorder *MockChainConfigMockRecorder
}

// MockChainConfigMockRecorder is the mock recorder for MockChainConfig.
type MockChainConfigMockRecorder struct {
	mock *MockChainConfig
}

// NewMockChainConfig creates a new mock instance.
func NewMockChainConfig(ctrl *gomock.Controller) *MockChainConfig {
	mock := &MockChainConfig{ctrl: ctrl}
	mock.recorder = &MockChainConfigMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockChainConfig) EXPECT() *MockChainConfigMockRecorder {
	return m.recorder
}

// IsDUpgrade mocks base method.
func (m *MockChainConfig) IsDUpgrade(time uint64) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsDUpgrade", time)
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsDUpgrade indicates an expected call of IsDUpgrade.
func (mr *MockChainConfigMockRecorder) IsDUpgrade(time interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsDUpgrade", reflect.TypeOf((*MockChainConfig)(nil).IsDUpgrade), time)
}
