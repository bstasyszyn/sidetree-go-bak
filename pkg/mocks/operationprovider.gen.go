// Code generated by counterfeiter. DO NOT EDIT.
package mocks

import (
	"sync"

	"github.com/trustbloc/sidetree-go/pkg/api/operation"
	"github.com/trustbloc/sidetree-go/pkg/api/protocol"
	"github.com/trustbloc/sidetree-go/pkg/api/txn"
)

type OperationProvider struct {
	GetTxnOperationsStub        func(*txn.SidetreeTxn) ([]*operation.AnchoredOperation, error)
	getTxnOperationsMutex       sync.RWMutex
	getTxnOperationsArgsForCall []struct {
		arg1 *txn.SidetreeTxn
	}
	getTxnOperationsReturns struct {
		result1 []*operation.AnchoredOperation
		result2 error
	}
	getTxnOperationsReturnsOnCall map[int]struct {
		result1 []*operation.AnchoredOperation
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *OperationProvider) GetTxnOperations(arg1 *txn.SidetreeTxn) ([]*operation.AnchoredOperation, error) {
	fake.getTxnOperationsMutex.Lock()
	ret, specificReturn := fake.getTxnOperationsReturnsOnCall[len(fake.getTxnOperationsArgsForCall)]
	fake.getTxnOperationsArgsForCall = append(fake.getTxnOperationsArgsForCall, struct {
		arg1 *txn.SidetreeTxn
	}{arg1})
	fake.recordInvocation("GetTxnOperations", []interface{}{arg1})
	fake.getTxnOperationsMutex.Unlock()
	if fake.GetTxnOperationsStub != nil {
		return fake.GetTxnOperationsStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.getTxnOperationsReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *OperationProvider) GetTxnOperationsCallCount() int {
	fake.getTxnOperationsMutex.RLock()
	defer fake.getTxnOperationsMutex.RUnlock()
	return len(fake.getTxnOperationsArgsForCall)
}

func (fake *OperationProvider) GetTxnOperationsCalls(stub func(*txn.SidetreeTxn) ([]*operation.AnchoredOperation, error)) {
	fake.getTxnOperationsMutex.Lock()
	defer fake.getTxnOperationsMutex.Unlock()
	fake.GetTxnOperationsStub = stub
}

func (fake *OperationProvider) GetTxnOperationsArgsForCall(i int) *txn.SidetreeTxn {
	fake.getTxnOperationsMutex.RLock()
	defer fake.getTxnOperationsMutex.RUnlock()
	argsForCall := fake.getTxnOperationsArgsForCall[i]
	return argsForCall.arg1
}

func (fake *OperationProvider) GetTxnOperationsReturns(result1 []*operation.AnchoredOperation, result2 error) {
	fake.getTxnOperationsMutex.Lock()
	defer fake.getTxnOperationsMutex.Unlock()
	fake.GetTxnOperationsStub = nil
	fake.getTxnOperationsReturns = struct {
		result1 []*operation.AnchoredOperation
		result2 error
	}{result1, result2}
}

func (fake *OperationProvider) GetTxnOperationsReturnsOnCall(i int, result1 []*operation.AnchoredOperation, result2 error) {
	fake.getTxnOperationsMutex.Lock()
	defer fake.getTxnOperationsMutex.Unlock()
	fake.GetTxnOperationsStub = nil
	if fake.getTxnOperationsReturnsOnCall == nil {
		fake.getTxnOperationsReturnsOnCall = make(map[int]struct {
			result1 []*operation.AnchoredOperation
			result2 error
		})
	}
	fake.getTxnOperationsReturnsOnCall[i] = struct {
		result1 []*operation.AnchoredOperation
		result2 error
	}{result1, result2}
}

func (fake *OperationProvider) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.getTxnOperationsMutex.RLock()
	defer fake.getTxnOperationsMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *OperationProvider) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}

var _ protocol.OperationProvider = new(OperationProvider)
