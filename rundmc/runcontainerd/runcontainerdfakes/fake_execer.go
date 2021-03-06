// Code generated by counterfeiter. DO NOT EDIT.
package runcontainerdfakes

import (
	"sync"

	"code.cloudfoundry.org/garden"
	"code.cloudfoundry.org/guardian/rundmc/runcontainerd"
	"code.cloudfoundry.org/lager"
)

type FakeExecer struct {
	AttachStub        func(lager.Logger, string, string, string, garden.ProcessIO) (garden.Process, error)
	attachMutex       sync.RWMutex
	attachArgsForCall []struct {
		arg1 lager.Logger
		arg2 string
		arg3 string
		arg4 string
		arg5 garden.ProcessIO
	}
	attachReturns struct {
		result1 garden.Process
		result2 error
	}
	attachReturnsOnCall map[int]struct {
		result1 garden.Process
		result2 error
	}
	ExecStub        func(lager.Logger, string, string, garden.ProcessSpec, garden.ProcessIO) (garden.Process, error)
	execMutex       sync.RWMutex
	execArgsForCall []struct {
		arg1 lager.Logger
		arg2 string
		arg3 string
		arg4 garden.ProcessSpec
		arg5 garden.ProcessIO
	}
	execReturns struct {
		result1 garden.Process
		result2 error
	}
	execReturnsOnCall map[int]struct {
		result1 garden.Process
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeExecer) Attach(arg1 lager.Logger, arg2 string, arg3 string, arg4 string, arg5 garden.ProcessIO) (garden.Process, error) {
	fake.attachMutex.Lock()
	ret, specificReturn := fake.attachReturnsOnCall[len(fake.attachArgsForCall)]
	fake.attachArgsForCall = append(fake.attachArgsForCall, struct {
		arg1 lager.Logger
		arg2 string
		arg3 string
		arg4 string
		arg5 garden.ProcessIO
	}{arg1, arg2, arg3, arg4, arg5})
	fake.recordInvocation("Attach", []interface{}{arg1, arg2, arg3, arg4, arg5})
	fake.attachMutex.Unlock()
	if fake.AttachStub != nil {
		return fake.AttachStub(arg1, arg2, arg3, arg4, arg5)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.attachReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeExecer) AttachCallCount() int {
	fake.attachMutex.RLock()
	defer fake.attachMutex.RUnlock()
	return len(fake.attachArgsForCall)
}

func (fake *FakeExecer) AttachCalls(stub func(lager.Logger, string, string, string, garden.ProcessIO) (garden.Process, error)) {
	fake.attachMutex.Lock()
	defer fake.attachMutex.Unlock()
	fake.AttachStub = stub
}

func (fake *FakeExecer) AttachArgsForCall(i int) (lager.Logger, string, string, string, garden.ProcessIO) {
	fake.attachMutex.RLock()
	defer fake.attachMutex.RUnlock()
	argsForCall := fake.attachArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3, argsForCall.arg4, argsForCall.arg5
}

func (fake *FakeExecer) AttachReturns(result1 garden.Process, result2 error) {
	fake.attachMutex.Lock()
	defer fake.attachMutex.Unlock()
	fake.AttachStub = nil
	fake.attachReturns = struct {
		result1 garden.Process
		result2 error
	}{result1, result2}
}

func (fake *FakeExecer) AttachReturnsOnCall(i int, result1 garden.Process, result2 error) {
	fake.attachMutex.Lock()
	defer fake.attachMutex.Unlock()
	fake.AttachStub = nil
	if fake.attachReturnsOnCall == nil {
		fake.attachReturnsOnCall = make(map[int]struct {
			result1 garden.Process
			result2 error
		})
	}
	fake.attachReturnsOnCall[i] = struct {
		result1 garden.Process
		result2 error
	}{result1, result2}
}

func (fake *FakeExecer) Exec(arg1 lager.Logger, arg2 string, arg3 string, arg4 garden.ProcessSpec, arg5 garden.ProcessIO) (garden.Process, error) {
	fake.execMutex.Lock()
	ret, specificReturn := fake.execReturnsOnCall[len(fake.execArgsForCall)]
	fake.execArgsForCall = append(fake.execArgsForCall, struct {
		arg1 lager.Logger
		arg2 string
		arg3 string
		arg4 garden.ProcessSpec
		arg5 garden.ProcessIO
	}{arg1, arg2, arg3, arg4, arg5})
	fake.recordInvocation("Exec", []interface{}{arg1, arg2, arg3, arg4, arg5})
	fake.execMutex.Unlock()
	if fake.ExecStub != nil {
		return fake.ExecStub(arg1, arg2, arg3, arg4, arg5)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.execReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeExecer) ExecCallCount() int {
	fake.execMutex.RLock()
	defer fake.execMutex.RUnlock()
	return len(fake.execArgsForCall)
}

func (fake *FakeExecer) ExecCalls(stub func(lager.Logger, string, string, garden.ProcessSpec, garden.ProcessIO) (garden.Process, error)) {
	fake.execMutex.Lock()
	defer fake.execMutex.Unlock()
	fake.ExecStub = stub
}

func (fake *FakeExecer) ExecArgsForCall(i int) (lager.Logger, string, string, garden.ProcessSpec, garden.ProcessIO) {
	fake.execMutex.RLock()
	defer fake.execMutex.RUnlock()
	argsForCall := fake.execArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3, argsForCall.arg4, argsForCall.arg5
}

func (fake *FakeExecer) ExecReturns(result1 garden.Process, result2 error) {
	fake.execMutex.Lock()
	defer fake.execMutex.Unlock()
	fake.ExecStub = nil
	fake.execReturns = struct {
		result1 garden.Process
		result2 error
	}{result1, result2}
}

func (fake *FakeExecer) ExecReturnsOnCall(i int, result1 garden.Process, result2 error) {
	fake.execMutex.Lock()
	defer fake.execMutex.Unlock()
	fake.ExecStub = nil
	if fake.execReturnsOnCall == nil {
		fake.execReturnsOnCall = make(map[int]struct {
			result1 garden.Process
			result2 error
		})
	}
	fake.execReturnsOnCall[i] = struct {
		result1 garden.Process
		result2 error
	}{result1, result2}
}

func (fake *FakeExecer) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.attachMutex.RLock()
	defer fake.attachMutex.RUnlock()
	fake.execMutex.RLock()
	defer fake.execMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeExecer) recordInvocation(key string, args []interface{}) {
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

var _ runcontainerd.Execer = new(FakeExecer)
