// Code generated by MockGen. DO NOT EDIT.
// Source: storj.io/storj/pkg/pointerdb/pdbclient (interfaces: Client)

// Package mock_pointerdb is a generated GoMock package.
package mock_pointerdb

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
	pb "storj.io/storj/pkg/pb"
	pdbclient "storj.io/storj/pkg/pointerdb/pdbclient"
)

// MockClient is a mock of Client interface
type MockClient struct {
	ctrl     *gomock.Controller
	recorder *MockClientMockRecorder
}

// MockClientMockRecorder is the mock recorder for MockClient
type MockClientMockRecorder struct {
	mock *MockClient
}

// NewMockClient creates a new mock instance
func NewMockClient(ctrl *gomock.Controller) *MockClient {
	mock := &MockClient{ctrl: ctrl}
	mock.recorder = &MockClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockClient) EXPECT() *MockClientMockRecorder {
	return m.recorder
}

// Delete mocks base method
func (m *MockClient) Delete(arg0 context.Context, arg1 string) error {
	ret := m.ctrl.Call(m, "Delete", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete
func (mr *MockClientMockRecorder) Delete(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockClient)(nil).Delete), arg0, arg1)
}

// Get mocks base method
func (m *MockClient) Get(arg0 context.Context, arg1 string) (*pb.Pointer, error) {
	ret := m.ctrl.Call(m, "Get", arg0, arg1)
	ret0, _ := ret[0].(*pb.Pointer)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get
func (mr *MockClientMockRecorder) Get(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockClient)(nil).Get), arg0, arg1)
}

// List mocks base method
func (m *MockClient) List(arg0 context.Context, arg1, arg2, arg3 string, arg4 bool, arg5 int, arg6 uint32) ([]pdbclient.ListItem, bool, error) {
	ret := m.ctrl.Call(m, "List", arg0, arg1, arg2, arg3, arg4, arg5, arg6)
	ret0, _ := ret[0].([]pdbclient.ListItem)
	ret1, _ := ret[1].(bool)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// List indicates an expected call of List
func (mr *MockClientMockRecorder) List(arg0, arg1, arg2, arg3, arg4, arg5, arg6 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockClient)(nil).List), arg0, arg1, arg2, arg3, arg4, arg5, arg6)
}

// PayerBandwidthAllocation mocks base method
func (m *MockClient) PayerBandwidthAllocation() *pb.PayerBandwidthAllocation {
	ret := m.ctrl.Call(m, "PayerBandwidthAllocation")
	ret0, _ := ret[0].(*pb.PayerBandwidthAllocation)
	return ret0
}

// PayerBandwidthAllocation indicates an expected call of PayerBandwidthAllocation
func (mr *MockClientMockRecorder) PayerBandwidthAllocation() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PayerBandwidthAllocation", reflect.TypeOf((*MockClient)(nil).PayerBandwidthAllocation))
}

// Put mocks base method
func (m *MockClient) Put(arg0 context.Context, arg1 string, arg2 *pb.Pointer) error {
	ret := m.ctrl.Call(m, "Put", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// Put indicates an expected call of Put
func (mr *MockClientMockRecorder) Put(arg0, arg1, arg2 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Put", reflect.TypeOf((*MockClient)(nil).Put), arg0, arg1, arg2)
}

// SignedMessage mocks base method
func (m *MockClient) SignedMessage() (*pb.SignedMessage, error) {
	ret := m.ctrl.Call(m, "SignedMessage")
	ret0, _ := ret[0].(*pb.SignedMessage)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SignedMessage indicates an expected call of SignedMessage
func (mr *MockClientMockRecorder) SignedMessage() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SignedMessage", reflect.TypeOf((*MockClient)(nil).SignedMessage))
}
