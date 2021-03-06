// Code generated by counterfeiter. DO NOT EDIT.
package depotfakes

import (
	"sync"

	spec "code.cloudfoundry.org/guardian/gardener/container-spec"
	"code.cloudfoundry.org/guardian/rundmc/depot"
	"code.cloudfoundry.org/guardian/rundmc/goci"
)

type FakeBundleGenerator struct {
	GenerateStub        func(spec.DesiredContainerSpec, string) (goci.Bndl, error)
	generateMutex       sync.RWMutex
	generateArgsForCall []struct {
		arg1 spec.DesiredContainerSpec
		arg2 string
	}
	generateReturns struct {
		result1 goci.Bndl
		result2 error
	}
	generateReturnsOnCall map[int]struct {
		result1 goci.Bndl
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeBundleGenerator) Generate(arg1 spec.DesiredContainerSpec, arg2 string) (goci.Bndl, error) {
	fake.generateMutex.Lock()
	ret, specificReturn := fake.generateReturnsOnCall[len(fake.generateArgsForCall)]
	fake.generateArgsForCall = append(fake.generateArgsForCall, struct {
		arg1 spec.DesiredContainerSpec
		arg2 string
	}{arg1, arg2})
	fake.recordInvocation("Generate", []interface{}{arg1, arg2})
	fake.generateMutex.Unlock()
	if fake.GenerateStub != nil {
		return fake.GenerateStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.generateReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeBundleGenerator) GenerateCallCount() int {
	fake.generateMutex.RLock()
	defer fake.generateMutex.RUnlock()
	return len(fake.generateArgsForCall)
}

func (fake *FakeBundleGenerator) GenerateCalls(stub func(spec.DesiredContainerSpec, string) (goci.Bndl, error)) {
	fake.generateMutex.Lock()
	defer fake.generateMutex.Unlock()
	fake.GenerateStub = stub
}

func (fake *FakeBundleGenerator) GenerateArgsForCall(i int) (spec.DesiredContainerSpec, string) {
	fake.generateMutex.RLock()
	defer fake.generateMutex.RUnlock()
	argsForCall := fake.generateArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeBundleGenerator) GenerateReturns(result1 goci.Bndl, result2 error) {
	fake.generateMutex.Lock()
	defer fake.generateMutex.Unlock()
	fake.GenerateStub = nil
	fake.generateReturns = struct {
		result1 goci.Bndl
		result2 error
	}{result1, result2}
}

func (fake *FakeBundleGenerator) GenerateReturnsOnCall(i int, result1 goci.Bndl, result2 error) {
	fake.generateMutex.Lock()
	defer fake.generateMutex.Unlock()
	fake.GenerateStub = nil
	if fake.generateReturnsOnCall == nil {
		fake.generateReturnsOnCall = make(map[int]struct {
			result1 goci.Bndl
			result2 error
		})
	}
	fake.generateReturnsOnCall[i] = struct {
		result1 goci.Bndl
		result2 error
	}{result1, result2}
}

func (fake *FakeBundleGenerator) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.generateMutex.RLock()
	defer fake.generateMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeBundleGenerator) recordInvocation(key string, args []interface{}) {
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

var _ depot.BundleGenerator = new(FakeBundleGenerator)
