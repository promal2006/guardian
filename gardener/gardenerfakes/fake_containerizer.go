// Code generated by counterfeiter. DO NOT EDIT.
package gardenerfakes

import (
	"io"
	"sync"

	"code.cloudfoundry.org/garden"
	"code.cloudfoundry.org/guardian/gardener"
	"code.cloudfoundry.org/guardian/gardener/container-spec"
	"code.cloudfoundry.org/lager"
)

type FakeContainerizer struct {
	CreateStub        func(log lager.Logger, desiredContainerSpec spec.DesiredContainerSpec) error
	createMutex       sync.RWMutex
	createArgsForCall []struct {
		log                  lager.Logger
		desiredContainerSpec spec.DesiredContainerSpec
	}
	createReturns struct {
		result1 error
	}
	createReturnsOnCall map[int]struct {
		result1 error
	}
	HandlesStub        func() ([]string, error)
	handlesMutex       sync.RWMutex
	handlesArgsForCall []struct{}
	handlesReturns     struct {
		result1 []string
		result2 error
	}
	handlesReturnsOnCall map[int]struct {
		result1 []string
		result2 error
	}
	StreamInStub        func(log lager.Logger, handle string, streamInSpec garden.StreamInSpec) error
	streamInMutex       sync.RWMutex
	streamInArgsForCall []struct {
		log          lager.Logger
		handle       string
		streamInSpec garden.StreamInSpec
	}
	streamInReturns struct {
		result1 error
	}
	streamInReturnsOnCall map[int]struct {
		result1 error
	}
	StreamOutStub        func(log lager.Logger, handle string, streamOutSpec garden.StreamOutSpec) (io.ReadCloser, error)
	streamOutMutex       sync.RWMutex
	streamOutArgsForCall []struct {
		log           lager.Logger
		handle        string
		streamOutSpec garden.StreamOutSpec
	}
	streamOutReturns struct {
		result1 io.ReadCloser
		result2 error
	}
	streamOutReturnsOnCall map[int]struct {
		result1 io.ReadCloser
		result2 error
	}
	RunStub        func(log lager.Logger, handle string, processSpec garden.ProcessSpec, io garden.ProcessIO) (garden.Process, error)
	runMutex       sync.RWMutex
	runArgsForCall []struct {
		log         lager.Logger
		handle      string
		processSpec garden.ProcessSpec
		io          garden.ProcessIO
	}
	runReturns struct {
		result1 garden.Process
		result2 error
	}
	runReturnsOnCall map[int]struct {
		result1 garden.Process
		result2 error
	}
	AttachStub        func(log lager.Logger, handle string, processGUID string, io garden.ProcessIO) (garden.Process, error)
	attachMutex       sync.RWMutex
	attachArgsForCall []struct {
		log         lager.Logger
		handle      string
		processGUID string
		io          garden.ProcessIO
	}
	attachReturns struct {
		result1 garden.Process
		result2 error
	}
	attachReturnsOnCall map[int]struct {
		result1 garden.Process
		result2 error
	}
	StopStub        func(log lager.Logger, handle string, kill bool) error
	stopMutex       sync.RWMutex
	stopArgsForCall []struct {
		log    lager.Logger
		handle string
		kill   bool
	}
	stopReturns struct {
		result1 error
	}
	stopReturnsOnCall map[int]struct {
		result1 error
	}
	DestroyStub        func(log lager.Logger, handle string) error
	destroyMutex       sync.RWMutex
	destroyArgsForCall []struct {
		log    lager.Logger
		handle string
	}
	destroyReturns struct {
		result1 error
	}
	destroyReturnsOnCall map[int]struct {
		result1 error
	}
	RemoveBundleStub        func(log lager.Logger, handle string) error
	removeBundleMutex       sync.RWMutex
	removeBundleArgsForCall []struct {
		log    lager.Logger
		handle string
	}
	removeBundleReturns struct {
		result1 error
	}
	removeBundleReturnsOnCall map[int]struct {
		result1 error
	}
	InfoStub        func(log lager.Logger, handle string) (spec.ActualContainerSpec, error)
	infoMutex       sync.RWMutex
	infoArgsForCall []struct {
		log    lager.Logger
		handle string
	}
	infoReturns struct {
		result1 spec.ActualContainerSpec
		result2 error
	}
	infoReturnsOnCall map[int]struct {
		result1 spec.ActualContainerSpec
		result2 error
	}
	MetricsStub        func(log lager.Logger, handle string) (gardener.ActualContainerMetrics, error)
	metricsMutex       sync.RWMutex
	metricsArgsForCall []struct {
		log    lager.Logger
		handle string
	}
	metricsReturns struct {
		result1 gardener.ActualContainerMetrics
		result2 error
	}
	metricsReturnsOnCall map[int]struct {
		result1 gardener.ActualContainerMetrics
		result2 error
	}
	BulkMetricsStub        func(log lager.Logger, handles []string) (map[string]gardener.ActualContainerMetrics, error)
	bulkMetricsMutex       sync.RWMutex
	bulkMetricsArgsForCall []struct {
		log     lager.Logger
		handles []string
	}
	bulkMetricsReturns struct {
		result1 map[string]gardener.ActualContainerMetrics
		result2 error
	}
	bulkMetricsReturnsOnCall map[int]struct {
		result1 map[string]gardener.ActualContainerMetrics
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeContainerizer) Create(log lager.Logger, desiredContainerSpec spec.DesiredContainerSpec) error {
	fake.createMutex.Lock()
	ret, specificReturn := fake.createReturnsOnCall[len(fake.createArgsForCall)]
	fake.createArgsForCall = append(fake.createArgsForCall, struct {
		log                  lager.Logger
		desiredContainerSpec spec.DesiredContainerSpec
	}{log, desiredContainerSpec})
	fake.recordInvocation("Create", []interface{}{log, desiredContainerSpec})
	fake.createMutex.Unlock()
	if fake.CreateStub != nil {
		return fake.CreateStub(log, desiredContainerSpec)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.createReturns.result1
}

func (fake *FakeContainerizer) CreateCallCount() int {
	fake.createMutex.RLock()
	defer fake.createMutex.RUnlock()
	return len(fake.createArgsForCall)
}

func (fake *FakeContainerizer) CreateArgsForCall(i int) (lager.Logger, spec.DesiredContainerSpec) {
	fake.createMutex.RLock()
	defer fake.createMutex.RUnlock()
	return fake.createArgsForCall[i].log, fake.createArgsForCall[i].desiredContainerSpec
}

func (fake *FakeContainerizer) CreateReturns(result1 error) {
	fake.CreateStub = nil
	fake.createReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeContainerizer) CreateReturnsOnCall(i int, result1 error) {
	fake.CreateStub = nil
	if fake.createReturnsOnCall == nil {
		fake.createReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.createReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeContainerizer) Handles() ([]string, error) {
	fake.handlesMutex.Lock()
	ret, specificReturn := fake.handlesReturnsOnCall[len(fake.handlesArgsForCall)]
	fake.handlesArgsForCall = append(fake.handlesArgsForCall, struct{}{})
	fake.recordInvocation("Handles", []interface{}{})
	fake.handlesMutex.Unlock()
	if fake.HandlesStub != nil {
		return fake.HandlesStub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.handlesReturns.result1, fake.handlesReturns.result2
}

func (fake *FakeContainerizer) HandlesCallCount() int {
	fake.handlesMutex.RLock()
	defer fake.handlesMutex.RUnlock()
	return len(fake.handlesArgsForCall)
}

func (fake *FakeContainerizer) HandlesReturns(result1 []string, result2 error) {
	fake.HandlesStub = nil
	fake.handlesReturns = struct {
		result1 []string
		result2 error
	}{result1, result2}
}

func (fake *FakeContainerizer) HandlesReturnsOnCall(i int, result1 []string, result2 error) {
	fake.HandlesStub = nil
	if fake.handlesReturnsOnCall == nil {
		fake.handlesReturnsOnCall = make(map[int]struct {
			result1 []string
			result2 error
		})
	}
	fake.handlesReturnsOnCall[i] = struct {
		result1 []string
		result2 error
	}{result1, result2}
}

func (fake *FakeContainerizer) StreamIn(log lager.Logger, handle string, streamInSpec garden.StreamInSpec) error {
	fake.streamInMutex.Lock()
	ret, specificReturn := fake.streamInReturnsOnCall[len(fake.streamInArgsForCall)]
	fake.streamInArgsForCall = append(fake.streamInArgsForCall, struct {
		log          lager.Logger
		handle       string
		streamInSpec garden.StreamInSpec
	}{log, handle, streamInSpec})
	fake.recordInvocation("StreamIn", []interface{}{log, handle, streamInSpec})
	fake.streamInMutex.Unlock()
	if fake.StreamInStub != nil {
		return fake.StreamInStub(log, handle, streamInSpec)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.streamInReturns.result1
}

func (fake *FakeContainerizer) StreamInCallCount() int {
	fake.streamInMutex.RLock()
	defer fake.streamInMutex.RUnlock()
	return len(fake.streamInArgsForCall)
}

func (fake *FakeContainerizer) StreamInArgsForCall(i int) (lager.Logger, string, garden.StreamInSpec) {
	fake.streamInMutex.RLock()
	defer fake.streamInMutex.RUnlock()
	return fake.streamInArgsForCall[i].log, fake.streamInArgsForCall[i].handle, fake.streamInArgsForCall[i].streamInSpec
}

func (fake *FakeContainerizer) StreamInReturns(result1 error) {
	fake.StreamInStub = nil
	fake.streamInReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeContainerizer) StreamInReturnsOnCall(i int, result1 error) {
	fake.StreamInStub = nil
	if fake.streamInReturnsOnCall == nil {
		fake.streamInReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.streamInReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeContainerizer) StreamOut(log lager.Logger, handle string, streamOutSpec garden.StreamOutSpec) (io.ReadCloser, error) {
	fake.streamOutMutex.Lock()
	ret, specificReturn := fake.streamOutReturnsOnCall[len(fake.streamOutArgsForCall)]
	fake.streamOutArgsForCall = append(fake.streamOutArgsForCall, struct {
		log           lager.Logger
		handle        string
		streamOutSpec garden.StreamOutSpec
	}{log, handle, streamOutSpec})
	fake.recordInvocation("StreamOut", []interface{}{log, handle, streamOutSpec})
	fake.streamOutMutex.Unlock()
	if fake.StreamOutStub != nil {
		return fake.StreamOutStub(log, handle, streamOutSpec)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.streamOutReturns.result1, fake.streamOutReturns.result2
}

func (fake *FakeContainerizer) StreamOutCallCount() int {
	fake.streamOutMutex.RLock()
	defer fake.streamOutMutex.RUnlock()
	return len(fake.streamOutArgsForCall)
}

func (fake *FakeContainerizer) StreamOutArgsForCall(i int) (lager.Logger, string, garden.StreamOutSpec) {
	fake.streamOutMutex.RLock()
	defer fake.streamOutMutex.RUnlock()
	return fake.streamOutArgsForCall[i].log, fake.streamOutArgsForCall[i].handle, fake.streamOutArgsForCall[i].streamOutSpec
}

func (fake *FakeContainerizer) StreamOutReturns(result1 io.ReadCloser, result2 error) {
	fake.StreamOutStub = nil
	fake.streamOutReturns = struct {
		result1 io.ReadCloser
		result2 error
	}{result1, result2}
}

func (fake *FakeContainerizer) StreamOutReturnsOnCall(i int, result1 io.ReadCloser, result2 error) {
	fake.StreamOutStub = nil
	if fake.streamOutReturnsOnCall == nil {
		fake.streamOutReturnsOnCall = make(map[int]struct {
			result1 io.ReadCloser
			result2 error
		})
	}
	fake.streamOutReturnsOnCall[i] = struct {
		result1 io.ReadCloser
		result2 error
	}{result1, result2}
}

func (fake *FakeContainerizer) Run(log lager.Logger, handle string, processSpec garden.ProcessSpec, io garden.ProcessIO) (garden.Process, error) {
	fake.runMutex.Lock()
	ret, specificReturn := fake.runReturnsOnCall[len(fake.runArgsForCall)]
	fake.runArgsForCall = append(fake.runArgsForCall, struct {
		log         lager.Logger
		handle      string
		processSpec garden.ProcessSpec
		io          garden.ProcessIO
	}{log, handle, processSpec, io})
	fake.recordInvocation("Run", []interface{}{log, handle, processSpec, io})
	fake.runMutex.Unlock()
	if fake.RunStub != nil {
		return fake.RunStub(log, handle, processSpec, io)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.runReturns.result1, fake.runReturns.result2
}

func (fake *FakeContainerizer) RunCallCount() int {
	fake.runMutex.RLock()
	defer fake.runMutex.RUnlock()
	return len(fake.runArgsForCall)
}

func (fake *FakeContainerizer) RunArgsForCall(i int) (lager.Logger, string, garden.ProcessSpec, garden.ProcessIO) {
	fake.runMutex.RLock()
	defer fake.runMutex.RUnlock()
	return fake.runArgsForCall[i].log, fake.runArgsForCall[i].handle, fake.runArgsForCall[i].processSpec, fake.runArgsForCall[i].io
}

func (fake *FakeContainerizer) RunReturns(result1 garden.Process, result2 error) {
	fake.RunStub = nil
	fake.runReturns = struct {
		result1 garden.Process
		result2 error
	}{result1, result2}
}

func (fake *FakeContainerizer) RunReturnsOnCall(i int, result1 garden.Process, result2 error) {
	fake.RunStub = nil
	if fake.runReturnsOnCall == nil {
		fake.runReturnsOnCall = make(map[int]struct {
			result1 garden.Process
			result2 error
		})
	}
	fake.runReturnsOnCall[i] = struct {
		result1 garden.Process
		result2 error
	}{result1, result2}
}

func (fake *FakeContainerizer) Attach(log lager.Logger, handle string, processGUID string, io garden.ProcessIO) (garden.Process, error) {
	fake.attachMutex.Lock()
	ret, specificReturn := fake.attachReturnsOnCall[len(fake.attachArgsForCall)]
	fake.attachArgsForCall = append(fake.attachArgsForCall, struct {
		log         lager.Logger
		handle      string
		processGUID string
		io          garden.ProcessIO
	}{log, handle, processGUID, io})
	fake.recordInvocation("Attach", []interface{}{log, handle, processGUID, io})
	fake.attachMutex.Unlock()
	if fake.AttachStub != nil {
		return fake.AttachStub(log, handle, processGUID, io)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.attachReturns.result1, fake.attachReturns.result2
}

func (fake *FakeContainerizer) AttachCallCount() int {
	fake.attachMutex.RLock()
	defer fake.attachMutex.RUnlock()
	return len(fake.attachArgsForCall)
}

func (fake *FakeContainerizer) AttachArgsForCall(i int) (lager.Logger, string, string, garden.ProcessIO) {
	fake.attachMutex.RLock()
	defer fake.attachMutex.RUnlock()
	return fake.attachArgsForCall[i].log, fake.attachArgsForCall[i].handle, fake.attachArgsForCall[i].processGUID, fake.attachArgsForCall[i].io
}

func (fake *FakeContainerizer) AttachReturns(result1 garden.Process, result2 error) {
	fake.AttachStub = nil
	fake.attachReturns = struct {
		result1 garden.Process
		result2 error
	}{result1, result2}
}

func (fake *FakeContainerizer) AttachReturnsOnCall(i int, result1 garden.Process, result2 error) {
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

func (fake *FakeContainerizer) Stop(log lager.Logger, handle string, kill bool) error {
	fake.stopMutex.Lock()
	ret, specificReturn := fake.stopReturnsOnCall[len(fake.stopArgsForCall)]
	fake.stopArgsForCall = append(fake.stopArgsForCall, struct {
		log    lager.Logger
		handle string
		kill   bool
	}{log, handle, kill})
	fake.recordInvocation("Stop", []interface{}{log, handle, kill})
	fake.stopMutex.Unlock()
	if fake.StopStub != nil {
		return fake.StopStub(log, handle, kill)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.stopReturns.result1
}

func (fake *FakeContainerizer) StopCallCount() int {
	fake.stopMutex.RLock()
	defer fake.stopMutex.RUnlock()
	return len(fake.stopArgsForCall)
}

func (fake *FakeContainerizer) StopArgsForCall(i int) (lager.Logger, string, bool) {
	fake.stopMutex.RLock()
	defer fake.stopMutex.RUnlock()
	return fake.stopArgsForCall[i].log, fake.stopArgsForCall[i].handle, fake.stopArgsForCall[i].kill
}

func (fake *FakeContainerizer) StopReturns(result1 error) {
	fake.StopStub = nil
	fake.stopReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeContainerizer) StopReturnsOnCall(i int, result1 error) {
	fake.StopStub = nil
	if fake.stopReturnsOnCall == nil {
		fake.stopReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.stopReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeContainerizer) Destroy(log lager.Logger, handle string) error {
	fake.destroyMutex.Lock()
	ret, specificReturn := fake.destroyReturnsOnCall[len(fake.destroyArgsForCall)]
	fake.destroyArgsForCall = append(fake.destroyArgsForCall, struct {
		log    lager.Logger
		handle string
	}{log, handle})
	fake.recordInvocation("Destroy", []interface{}{log, handle})
	fake.destroyMutex.Unlock()
	if fake.DestroyStub != nil {
		return fake.DestroyStub(log, handle)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.destroyReturns.result1
}

func (fake *FakeContainerizer) DestroyCallCount() int {
	fake.destroyMutex.RLock()
	defer fake.destroyMutex.RUnlock()
	return len(fake.destroyArgsForCall)
}

func (fake *FakeContainerizer) DestroyArgsForCall(i int) (lager.Logger, string) {
	fake.destroyMutex.RLock()
	defer fake.destroyMutex.RUnlock()
	return fake.destroyArgsForCall[i].log, fake.destroyArgsForCall[i].handle
}

func (fake *FakeContainerizer) DestroyReturns(result1 error) {
	fake.DestroyStub = nil
	fake.destroyReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeContainerizer) DestroyReturnsOnCall(i int, result1 error) {
	fake.DestroyStub = nil
	if fake.destroyReturnsOnCall == nil {
		fake.destroyReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.destroyReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeContainerizer) RemoveBundle(log lager.Logger, handle string) error {
	fake.removeBundleMutex.Lock()
	ret, specificReturn := fake.removeBundleReturnsOnCall[len(fake.removeBundleArgsForCall)]
	fake.removeBundleArgsForCall = append(fake.removeBundleArgsForCall, struct {
		log    lager.Logger
		handle string
	}{log, handle})
	fake.recordInvocation("RemoveBundle", []interface{}{log, handle})
	fake.removeBundleMutex.Unlock()
	if fake.RemoveBundleStub != nil {
		return fake.RemoveBundleStub(log, handle)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.removeBundleReturns.result1
}

func (fake *FakeContainerizer) RemoveBundleCallCount() int {
	fake.removeBundleMutex.RLock()
	defer fake.removeBundleMutex.RUnlock()
	return len(fake.removeBundleArgsForCall)
}

func (fake *FakeContainerizer) RemoveBundleArgsForCall(i int) (lager.Logger, string) {
	fake.removeBundleMutex.RLock()
	defer fake.removeBundleMutex.RUnlock()
	return fake.removeBundleArgsForCall[i].log, fake.removeBundleArgsForCall[i].handle
}

func (fake *FakeContainerizer) RemoveBundleReturns(result1 error) {
	fake.RemoveBundleStub = nil
	fake.removeBundleReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeContainerizer) RemoveBundleReturnsOnCall(i int, result1 error) {
	fake.RemoveBundleStub = nil
	if fake.removeBundleReturnsOnCall == nil {
		fake.removeBundleReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.removeBundleReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeContainerizer) Info(log lager.Logger, handle string) (spec.ActualContainerSpec, error) {
	fake.infoMutex.Lock()
	ret, specificReturn := fake.infoReturnsOnCall[len(fake.infoArgsForCall)]
	fake.infoArgsForCall = append(fake.infoArgsForCall, struct {
		log    lager.Logger
		handle string
	}{log, handle})
	fake.recordInvocation("Info", []interface{}{log, handle})
	fake.infoMutex.Unlock()
	if fake.InfoStub != nil {
		return fake.InfoStub(log, handle)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.infoReturns.result1, fake.infoReturns.result2
}

func (fake *FakeContainerizer) InfoCallCount() int {
	fake.infoMutex.RLock()
	defer fake.infoMutex.RUnlock()
	return len(fake.infoArgsForCall)
}

func (fake *FakeContainerizer) InfoArgsForCall(i int) (lager.Logger, string) {
	fake.infoMutex.RLock()
	defer fake.infoMutex.RUnlock()
	return fake.infoArgsForCall[i].log, fake.infoArgsForCall[i].handle
}

func (fake *FakeContainerizer) InfoReturns(result1 spec.ActualContainerSpec, result2 error) {
	fake.InfoStub = nil
	fake.infoReturns = struct {
		result1 spec.ActualContainerSpec
		result2 error
	}{result1, result2}
}

func (fake *FakeContainerizer) InfoReturnsOnCall(i int, result1 spec.ActualContainerSpec, result2 error) {
	fake.InfoStub = nil
	if fake.infoReturnsOnCall == nil {
		fake.infoReturnsOnCall = make(map[int]struct {
			result1 spec.ActualContainerSpec
			result2 error
		})
	}
	fake.infoReturnsOnCall[i] = struct {
		result1 spec.ActualContainerSpec
		result2 error
	}{result1, result2}
}

func (fake *FakeContainerizer) Metrics(log lager.Logger, handle string) (gardener.ActualContainerMetrics, error) {
	fake.metricsMutex.Lock()
	ret, specificReturn := fake.metricsReturnsOnCall[len(fake.metricsArgsForCall)]
	fake.metricsArgsForCall = append(fake.metricsArgsForCall, struct {
		log    lager.Logger
		handle string
	}{log, handle})
	fake.recordInvocation("Metrics", []interface{}{log, handle})
	fake.metricsMutex.Unlock()
	if fake.MetricsStub != nil {
		return fake.MetricsStub(log, handle)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.metricsReturns.result1, fake.metricsReturns.result2
}

func (fake *FakeContainerizer) MetricsCallCount() int {
	fake.metricsMutex.RLock()
	defer fake.metricsMutex.RUnlock()
	return len(fake.metricsArgsForCall)
}

func (fake *FakeContainerizer) MetricsArgsForCall(i int) (lager.Logger, string) {
	fake.metricsMutex.RLock()
	defer fake.metricsMutex.RUnlock()
	return fake.metricsArgsForCall[i].log, fake.metricsArgsForCall[i].handle
}

func (fake *FakeContainerizer) MetricsReturns(result1 gardener.ActualContainerMetrics, result2 error) {
	fake.MetricsStub = nil
	fake.metricsReturns = struct {
		result1 gardener.ActualContainerMetrics
		result2 error
	}{result1, result2}
}

func (fake *FakeContainerizer) MetricsReturnsOnCall(i int, result1 gardener.ActualContainerMetrics, result2 error) {
	fake.MetricsStub = nil
	if fake.metricsReturnsOnCall == nil {
		fake.metricsReturnsOnCall = make(map[int]struct {
			result1 gardener.ActualContainerMetrics
			result2 error
		})
	}
	fake.metricsReturnsOnCall[i] = struct {
		result1 gardener.ActualContainerMetrics
		result2 error
	}{result1, result2}
}

func (fake *FakeContainerizer) BulkMetrics(log lager.Logger, handles []string) (map[string]gardener.ActualContainerMetrics, error) {
	var handlesCopy []string
	if handles != nil {
		handlesCopy = make([]string, len(handles))
		copy(handlesCopy, handles)
	}
	fake.bulkMetricsMutex.Lock()
	ret, specificReturn := fake.bulkMetricsReturnsOnCall[len(fake.bulkMetricsArgsForCall)]
	fake.bulkMetricsArgsForCall = append(fake.bulkMetricsArgsForCall, struct {
		log     lager.Logger
		handles []string
	}{log, handlesCopy})
	fake.recordInvocation("BulkMetrics", []interface{}{log, handlesCopy})
	fake.bulkMetricsMutex.Unlock()
	if fake.BulkMetricsStub != nil {
		return fake.BulkMetricsStub(log, handles)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.bulkMetricsReturns.result1, fake.bulkMetricsReturns.result2
}

func (fake *FakeContainerizer) BulkMetricsCallCount() int {
	fake.bulkMetricsMutex.RLock()
	defer fake.bulkMetricsMutex.RUnlock()
	return len(fake.bulkMetricsArgsForCall)
}

func (fake *FakeContainerizer) BulkMetricsArgsForCall(i int) (lager.Logger, []string) {
	fake.bulkMetricsMutex.RLock()
	defer fake.bulkMetricsMutex.RUnlock()
	return fake.bulkMetricsArgsForCall[i].log, fake.bulkMetricsArgsForCall[i].handles
}

func (fake *FakeContainerizer) BulkMetricsReturns(result1 map[string]gardener.ActualContainerMetrics, result2 error) {
	fake.BulkMetricsStub = nil
	fake.bulkMetricsReturns = struct {
		result1 map[string]gardener.ActualContainerMetrics
		result2 error
	}{result1, result2}
}

func (fake *FakeContainerizer) BulkMetricsReturnsOnCall(i int, result1 map[string]gardener.ActualContainerMetrics, result2 error) {
	fake.BulkMetricsStub = nil
	if fake.bulkMetricsReturnsOnCall == nil {
		fake.bulkMetricsReturnsOnCall = make(map[int]struct {
			result1 map[string]gardener.ActualContainerMetrics
			result2 error
		})
	}
	fake.bulkMetricsReturnsOnCall[i] = struct {
		result1 map[string]gardener.ActualContainerMetrics
		result2 error
	}{result1, result2}
}

func (fake *FakeContainerizer) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.createMutex.RLock()
	defer fake.createMutex.RUnlock()
	fake.handlesMutex.RLock()
	defer fake.handlesMutex.RUnlock()
	fake.streamInMutex.RLock()
	defer fake.streamInMutex.RUnlock()
	fake.streamOutMutex.RLock()
	defer fake.streamOutMutex.RUnlock()
	fake.runMutex.RLock()
	defer fake.runMutex.RUnlock()
	fake.attachMutex.RLock()
	defer fake.attachMutex.RUnlock()
	fake.stopMutex.RLock()
	defer fake.stopMutex.RUnlock()
	fake.destroyMutex.RLock()
	defer fake.destroyMutex.RUnlock()
	fake.removeBundleMutex.RLock()
	defer fake.removeBundleMutex.RUnlock()
	fake.infoMutex.RLock()
	defer fake.infoMutex.RUnlock()
	fake.metricsMutex.RLock()
	defer fake.metricsMutex.RUnlock()
	fake.bulkMetricsMutex.RLock()
	defer fake.bulkMetricsMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeContainerizer) recordInvocation(key string, args []interface{}) {
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

var _ gardener.Containerizer = new(FakeContainerizer)
