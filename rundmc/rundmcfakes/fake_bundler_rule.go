// This file was generated by counterfeiter
package rundmcfakes

import (
	"sync"

	"code.cloudfoundry.org/guardian/gardener"
	"code.cloudfoundry.org/guardian/rundmc"
	"code.cloudfoundry.org/guardian/rundmc/goci"
)

type FakeBundlerRule struct {
	ApplyStub        func(bndle goci.Bndl, spec gardener.DesiredContainerSpec) goci.Bndl
	applyMutex       sync.RWMutex
	applyArgsForCall []struct {
		bndle goci.Bndl
		spec  gardener.DesiredContainerSpec
	}
	applyReturns struct {
		result1 goci.Bndl
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeBundlerRule) Apply(bndle goci.Bndl, spec gardener.DesiredContainerSpec) goci.Bndl {
	fake.applyMutex.Lock()
	fake.applyArgsForCall = append(fake.applyArgsForCall, struct {
		bndle goci.Bndl
		spec  gardener.DesiredContainerSpec
	}{bndle, spec})
	fake.recordInvocation("Apply", []interface{}{bndle, spec})
	fake.applyMutex.Unlock()
	if fake.ApplyStub != nil {
		return fake.ApplyStub(bndle, spec)
	}
	return fake.applyReturns.result1
}

func (fake *FakeBundlerRule) ApplyCallCount() int {
	fake.applyMutex.RLock()
	defer fake.applyMutex.RUnlock()
	return len(fake.applyArgsForCall)
}

func (fake *FakeBundlerRule) ApplyArgsForCall(i int) (goci.Bndl, gardener.DesiredContainerSpec) {
	fake.applyMutex.RLock()
	defer fake.applyMutex.RUnlock()
	return fake.applyArgsForCall[i].bndle, fake.applyArgsForCall[i].spec
}

func (fake *FakeBundlerRule) ApplyReturns(result1 goci.Bndl) {
	fake.ApplyStub = nil
	fake.applyReturns = struct {
		result1 goci.Bndl
	}{result1}
}

func (fake *FakeBundlerRule) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.applyMutex.RLock()
	defer fake.applyMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeBundlerRule) recordInvocation(key string, args []interface{}) {
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

var _ rundmc.BundlerRule = new(FakeBundlerRule)
