// Code generated by counterfeiter. DO NOT EDIT.
package mocks

import (
	"sync"

	"github.com/trustbloc/sidetree-go/pkg/api/protocol"
	"github.com/trustbloc/sidetree-go/pkg/document"
	"github.com/trustbloc/sidetree-go/pkg/patch"
)

type DocumentComposer struct {
	ApplyPatchesStub        func(document.Document, []patch.Patch) (document.Document, error)
	applyPatchesMutex       sync.RWMutex
	applyPatchesArgsForCall []struct {
		arg1 document.Document
		arg2 []patch.Patch
	}
	applyPatchesReturns struct {
		result1 document.Document
		result2 error
	}
	applyPatchesReturnsOnCall map[int]struct {
		result1 document.Document
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *DocumentComposer) ApplyPatches(arg1 document.Document, arg2 []patch.Patch) (document.Document, error) {
	var arg2Copy []patch.Patch
	if arg2 != nil {
		arg2Copy = make([]patch.Patch, len(arg2))
		copy(arg2Copy, arg2)
	}
	fake.applyPatchesMutex.Lock()
	ret, specificReturn := fake.applyPatchesReturnsOnCall[len(fake.applyPatchesArgsForCall)]
	fake.applyPatchesArgsForCall = append(fake.applyPatchesArgsForCall, struct {
		arg1 document.Document
		arg2 []patch.Patch
	}{arg1, arg2Copy})
	fake.recordInvocation("ApplyPatches", []interface{}{arg1, arg2Copy})
	fake.applyPatchesMutex.Unlock()
	if fake.ApplyPatchesStub != nil {
		return fake.ApplyPatchesStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.applyPatchesReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *DocumentComposer) ApplyPatchesCallCount() int {
	fake.applyPatchesMutex.RLock()
	defer fake.applyPatchesMutex.RUnlock()
	return len(fake.applyPatchesArgsForCall)
}

func (fake *DocumentComposer) ApplyPatchesCalls(stub func(document.Document, []patch.Patch) (document.Document, error)) {
	fake.applyPatchesMutex.Lock()
	defer fake.applyPatchesMutex.Unlock()
	fake.ApplyPatchesStub = stub
}

func (fake *DocumentComposer) ApplyPatchesArgsForCall(i int) (document.Document, []patch.Patch) {
	fake.applyPatchesMutex.RLock()
	defer fake.applyPatchesMutex.RUnlock()
	argsForCall := fake.applyPatchesArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *DocumentComposer) ApplyPatchesReturns(result1 document.Document, result2 error) {
	fake.applyPatchesMutex.Lock()
	defer fake.applyPatchesMutex.Unlock()
	fake.ApplyPatchesStub = nil
	fake.applyPatchesReturns = struct {
		result1 document.Document
		result2 error
	}{result1, result2}
}

func (fake *DocumentComposer) ApplyPatchesReturnsOnCall(i int, result1 document.Document, result2 error) {
	fake.applyPatchesMutex.Lock()
	defer fake.applyPatchesMutex.Unlock()
	fake.ApplyPatchesStub = nil
	if fake.applyPatchesReturnsOnCall == nil {
		fake.applyPatchesReturnsOnCall = make(map[int]struct {
			result1 document.Document
			result2 error
		})
	}
	fake.applyPatchesReturnsOnCall[i] = struct {
		result1 document.Document
		result2 error
	}{result1, result2}
}

func (fake *DocumentComposer) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.applyPatchesMutex.RLock()
	defer fake.applyPatchesMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *DocumentComposer) recordInvocation(key string, args []interface{}) {
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

var _ protocol.DocumentComposer = new(DocumentComposer)
