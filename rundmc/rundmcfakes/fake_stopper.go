// This file was generated by counterfeiter
package rundmcfakes

import (
	"sync"

	"github.com/cloudfoundry-incubator/guardian/rundmc"
	"code.cloudfoundry.org/lager"
)

type FakeStopper struct {
	StopAllStub        func(log lager.Logger, cgroupName string, save []int, kill bool) error
	stopAllMutex       sync.RWMutex
	stopAllArgsForCall []struct {
		log        lager.Logger
		cgroupName string
		save       []int
		kill       bool
	}
	stopAllReturns struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeStopper) StopAll(log lager.Logger, cgroupName string, save []int, kill bool) error {
	var saveCopy []int
	if save != nil {
		saveCopy = make([]int, len(save))
		copy(saveCopy, save)
	}
	fake.stopAllMutex.Lock()
	fake.stopAllArgsForCall = append(fake.stopAllArgsForCall, struct {
		log        lager.Logger
		cgroupName string
		save       []int
		kill       bool
	}{log, cgroupName, saveCopy, kill})
	fake.recordInvocation("StopAll", []interface{}{log, cgroupName, saveCopy, kill})
	fake.stopAllMutex.Unlock()
	if fake.StopAllStub != nil {
		return fake.StopAllStub(log, cgroupName, save, kill)
	} else {
		return fake.stopAllReturns.result1
	}
}

func (fake *FakeStopper) StopAllCallCount() int {
	fake.stopAllMutex.RLock()
	defer fake.stopAllMutex.RUnlock()
	return len(fake.stopAllArgsForCall)
}

func (fake *FakeStopper) StopAllArgsForCall(i int) (lager.Logger, string, []int, bool) {
	fake.stopAllMutex.RLock()
	defer fake.stopAllMutex.RUnlock()
	return fake.stopAllArgsForCall[i].log, fake.stopAllArgsForCall[i].cgroupName, fake.stopAllArgsForCall[i].save, fake.stopAllArgsForCall[i].kill
}

func (fake *FakeStopper) StopAllReturns(result1 error) {
	fake.StopAllStub = nil
	fake.stopAllReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeStopper) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.stopAllMutex.RLock()
	defer fake.stopAllMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeStopper) recordInvocation(key string, args []interface{}) {
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

var _ rundmc.Stopper = new(FakeStopper)
