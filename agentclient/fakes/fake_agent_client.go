// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	"sync"

	"github.com/cloudfoundry/bosh-agent/agentclient"
	"github.com/cloudfoundry/bosh-agent/agentclient/applyspec"
)

type FakeAgentClient struct {
	AddPersistentDiskStub        func(string, interface{}) error
	addPersistentDiskMutex       sync.RWMutex
	addPersistentDiskArgsForCall []struct {
		arg1 string
		arg2 interface{}
	}
	addPersistentDiskReturns struct {
		result1 error
	}
	addPersistentDiskReturnsOnCall map[int]struct {
		result1 error
	}
	ApplyStub        func(applyspec.ApplySpec) error
	applyMutex       sync.RWMutex
	applyArgsForCall []struct {
		arg1 applyspec.ApplySpec
	}
	applyReturns struct {
		result1 error
	}
	applyReturnsOnCall map[int]struct {
		result1 error
	}
	CompilePackageStub        func(agentclient.BlobRef, []agentclient.BlobRef) (agentclient.BlobRef, error)
	compilePackageMutex       sync.RWMutex
	compilePackageArgsForCall []struct {
		arg1 agentclient.BlobRef
		arg2 []agentclient.BlobRef
	}
	compilePackageReturns struct {
		result1 agentclient.BlobRef
		result2 error
	}
	compilePackageReturnsOnCall map[int]struct {
		result1 agentclient.BlobRef
		result2 error
	}
	DeleteARPEntriesStub        func([]string) error
	deleteARPEntriesMutex       sync.RWMutex
	deleteARPEntriesArgsForCall []struct {
		arg1 []string
	}
	deleteARPEntriesReturns struct {
		result1 error
	}
	deleteARPEntriesReturnsOnCall map[int]struct {
		result1 error
	}
	DrainStub        func(string) (int64, error)
	drainMutex       sync.RWMutex
	drainArgsForCall []struct {
		arg1 string
	}
	drainReturns struct {
		result1 int64
		result2 error
	}
	drainReturnsOnCall map[int]struct {
		result1 int64
		result2 error
	}
	GetStateStub        func() (agentclient.AgentState, error)
	getStateMutex       sync.RWMutex
	getStateArgsForCall []struct {
	}
	getStateReturns struct {
		result1 agentclient.AgentState
		result2 error
	}
	getStateReturnsOnCall map[int]struct {
		result1 agentclient.AgentState
		result2 error
	}
	ListDiskStub        func() ([]string, error)
	listDiskMutex       sync.RWMutex
	listDiskArgsForCall []struct {
	}
	listDiskReturns struct {
		result1 []string
		result2 error
	}
	listDiskReturnsOnCall map[int]struct {
		result1 []string
		result2 error
	}
	MigrateDiskStub        func() error
	migrateDiskMutex       sync.RWMutex
	migrateDiskArgsForCall []struct {
	}
	migrateDiskReturns struct {
		result1 error
	}
	migrateDiskReturnsOnCall map[int]struct {
		result1 error
	}
	MountDiskStub        func(string) error
	mountDiskMutex       sync.RWMutex
	mountDiskArgsForCall []struct {
		arg1 string
	}
	mountDiskReturns struct {
		result1 error
	}
	mountDiskReturnsOnCall map[int]struct {
		result1 error
	}
	PingStub        func() (string, error)
	pingMutex       sync.RWMutex
	pingArgsForCall []struct {
	}
	pingReturns struct {
		result1 string
		result2 error
	}
	pingReturnsOnCall map[int]struct {
		result1 string
		result2 error
	}
	RemovePersistentDiskStub        func(string) error
	removePersistentDiskMutex       sync.RWMutex
	removePersistentDiskArgsForCall []struct {
		arg1 string
	}
	removePersistentDiskReturns struct {
		result1 error
	}
	removePersistentDiskReturnsOnCall map[int]struct {
		result1 error
	}
	RunScriptStub        func(string, map[string]interface{}) error
	runScriptMutex       sync.RWMutex
	runScriptArgsForCall []struct {
		arg1 string
		arg2 map[string]interface{}
	}
	runScriptReturns struct {
		result1 error
	}
	runScriptReturnsOnCall map[int]struct {
		result1 error
	}
	StartStub        func() error
	startMutex       sync.RWMutex
	startArgsForCall []struct {
	}
	startReturns struct {
		result1 error
	}
	startReturnsOnCall map[int]struct {
		result1 error
	}
	StopStub        func() error
	stopMutex       sync.RWMutex
	stopArgsForCall []struct {
	}
	stopReturns struct {
		result1 error
	}
	stopReturnsOnCall map[int]struct {
		result1 error
	}
	SyncDNSStub        func(string, string, uint64) (string, error)
	syncDNSMutex       sync.RWMutex
	syncDNSArgsForCall []struct {
		arg1 string
		arg2 string
		arg3 uint64
	}
	syncDNSReturns struct {
		result1 string
		result2 error
	}
	syncDNSReturnsOnCall map[int]struct {
		result1 string
		result2 error
	}
	UnmountDiskStub        func(string) error
	unmountDiskMutex       sync.RWMutex
	unmountDiskArgsForCall []struct {
		arg1 string
	}
	unmountDiskReturns struct {
		result1 error
	}
	unmountDiskReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeAgentClient) AddPersistentDisk(arg1 string, arg2 interface{}) error {
	fake.addPersistentDiskMutex.Lock()
	ret, specificReturn := fake.addPersistentDiskReturnsOnCall[len(fake.addPersistentDiskArgsForCall)]
	fake.addPersistentDiskArgsForCall = append(fake.addPersistentDiskArgsForCall, struct {
		arg1 string
		arg2 interface{}
	}{arg1, arg2})
	stub := fake.AddPersistentDiskStub
	fakeReturns := fake.addPersistentDiskReturns
	fake.recordInvocation("AddPersistentDisk", []interface{}{arg1, arg2})
	fake.addPersistentDiskMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeAgentClient) AddPersistentDiskCallCount() int {
	fake.addPersistentDiskMutex.RLock()
	defer fake.addPersistentDiskMutex.RUnlock()
	return len(fake.addPersistentDiskArgsForCall)
}

func (fake *FakeAgentClient) AddPersistentDiskCalls(stub func(string, interface{}) error) {
	fake.addPersistentDiskMutex.Lock()
	defer fake.addPersistentDiskMutex.Unlock()
	fake.AddPersistentDiskStub = stub
}

func (fake *FakeAgentClient) AddPersistentDiskArgsForCall(i int) (string, interface{}) {
	fake.addPersistentDiskMutex.RLock()
	defer fake.addPersistentDiskMutex.RUnlock()
	argsForCall := fake.addPersistentDiskArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeAgentClient) AddPersistentDiskReturns(result1 error) {
	fake.addPersistentDiskMutex.Lock()
	defer fake.addPersistentDiskMutex.Unlock()
	fake.AddPersistentDiskStub = nil
	fake.addPersistentDiskReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeAgentClient) AddPersistentDiskReturnsOnCall(i int, result1 error) {
	fake.addPersistentDiskMutex.Lock()
	defer fake.addPersistentDiskMutex.Unlock()
	fake.AddPersistentDiskStub = nil
	if fake.addPersistentDiskReturnsOnCall == nil {
		fake.addPersistentDiskReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.addPersistentDiskReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeAgentClient) Apply(arg1 applyspec.ApplySpec) error {
	fake.applyMutex.Lock()
	ret, specificReturn := fake.applyReturnsOnCall[len(fake.applyArgsForCall)]
	fake.applyArgsForCall = append(fake.applyArgsForCall, struct {
		arg1 applyspec.ApplySpec
	}{arg1})
	stub := fake.ApplyStub
	fakeReturns := fake.applyReturns
	fake.recordInvocation("Apply", []interface{}{arg1})
	fake.applyMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeAgentClient) ApplyCallCount() int {
	fake.applyMutex.RLock()
	defer fake.applyMutex.RUnlock()
	return len(fake.applyArgsForCall)
}

func (fake *FakeAgentClient) ApplyCalls(stub func(applyspec.ApplySpec) error) {
	fake.applyMutex.Lock()
	defer fake.applyMutex.Unlock()
	fake.ApplyStub = stub
}

func (fake *FakeAgentClient) ApplyArgsForCall(i int) applyspec.ApplySpec {
	fake.applyMutex.RLock()
	defer fake.applyMutex.RUnlock()
	argsForCall := fake.applyArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeAgentClient) ApplyReturns(result1 error) {
	fake.applyMutex.Lock()
	defer fake.applyMutex.Unlock()
	fake.ApplyStub = nil
	fake.applyReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeAgentClient) ApplyReturnsOnCall(i int, result1 error) {
	fake.applyMutex.Lock()
	defer fake.applyMutex.Unlock()
	fake.ApplyStub = nil
	if fake.applyReturnsOnCall == nil {
		fake.applyReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.applyReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeAgentClient) CompilePackage(arg1 agentclient.BlobRef, arg2 []agentclient.BlobRef) (agentclient.BlobRef, error) {
	var arg2Copy []agentclient.BlobRef
	if arg2 != nil {
		arg2Copy = make([]agentclient.BlobRef, len(arg2))
		copy(arg2Copy, arg2)
	}
	fake.compilePackageMutex.Lock()
	ret, specificReturn := fake.compilePackageReturnsOnCall[len(fake.compilePackageArgsForCall)]
	fake.compilePackageArgsForCall = append(fake.compilePackageArgsForCall, struct {
		arg1 agentclient.BlobRef
		arg2 []agentclient.BlobRef
	}{arg1, arg2Copy})
	stub := fake.CompilePackageStub
	fakeReturns := fake.compilePackageReturns
	fake.recordInvocation("CompilePackage", []interface{}{arg1, arg2Copy})
	fake.compilePackageMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeAgentClient) CompilePackageCallCount() int {
	fake.compilePackageMutex.RLock()
	defer fake.compilePackageMutex.RUnlock()
	return len(fake.compilePackageArgsForCall)
}

func (fake *FakeAgentClient) CompilePackageCalls(stub func(agentclient.BlobRef, []agentclient.BlobRef) (agentclient.BlobRef, error)) {
	fake.compilePackageMutex.Lock()
	defer fake.compilePackageMutex.Unlock()
	fake.CompilePackageStub = stub
}

func (fake *FakeAgentClient) CompilePackageArgsForCall(i int) (agentclient.BlobRef, []agentclient.BlobRef) {
	fake.compilePackageMutex.RLock()
	defer fake.compilePackageMutex.RUnlock()
	argsForCall := fake.compilePackageArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeAgentClient) CompilePackageReturns(result1 agentclient.BlobRef, result2 error) {
	fake.compilePackageMutex.Lock()
	defer fake.compilePackageMutex.Unlock()
	fake.CompilePackageStub = nil
	fake.compilePackageReturns = struct {
		result1 agentclient.BlobRef
		result2 error
	}{result1, result2}
}

func (fake *FakeAgentClient) CompilePackageReturnsOnCall(i int, result1 agentclient.BlobRef, result2 error) {
	fake.compilePackageMutex.Lock()
	defer fake.compilePackageMutex.Unlock()
	fake.CompilePackageStub = nil
	if fake.compilePackageReturnsOnCall == nil {
		fake.compilePackageReturnsOnCall = make(map[int]struct {
			result1 agentclient.BlobRef
			result2 error
		})
	}
	fake.compilePackageReturnsOnCall[i] = struct {
		result1 agentclient.BlobRef
		result2 error
	}{result1, result2}
}

func (fake *FakeAgentClient) DeleteARPEntries(arg1 []string) error {
	var arg1Copy []string
	if arg1 != nil {
		arg1Copy = make([]string, len(arg1))
		copy(arg1Copy, arg1)
	}
	fake.deleteARPEntriesMutex.Lock()
	ret, specificReturn := fake.deleteARPEntriesReturnsOnCall[len(fake.deleteARPEntriesArgsForCall)]
	fake.deleteARPEntriesArgsForCall = append(fake.deleteARPEntriesArgsForCall, struct {
		arg1 []string
	}{arg1Copy})
	stub := fake.DeleteARPEntriesStub
	fakeReturns := fake.deleteARPEntriesReturns
	fake.recordInvocation("DeleteARPEntries", []interface{}{arg1Copy})
	fake.deleteARPEntriesMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeAgentClient) DeleteARPEntriesCallCount() int {
	fake.deleteARPEntriesMutex.RLock()
	defer fake.deleteARPEntriesMutex.RUnlock()
	return len(fake.deleteARPEntriesArgsForCall)
}

func (fake *FakeAgentClient) DeleteARPEntriesCalls(stub func([]string) error) {
	fake.deleteARPEntriesMutex.Lock()
	defer fake.deleteARPEntriesMutex.Unlock()
	fake.DeleteARPEntriesStub = stub
}

func (fake *FakeAgentClient) DeleteARPEntriesArgsForCall(i int) []string {
	fake.deleteARPEntriesMutex.RLock()
	defer fake.deleteARPEntriesMutex.RUnlock()
	argsForCall := fake.deleteARPEntriesArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeAgentClient) DeleteARPEntriesReturns(result1 error) {
	fake.deleteARPEntriesMutex.Lock()
	defer fake.deleteARPEntriesMutex.Unlock()
	fake.DeleteARPEntriesStub = nil
	fake.deleteARPEntriesReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeAgentClient) DeleteARPEntriesReturnsOnCall(i int, result1 error) {
	fake.deleteARPEntriesMutex.Lock()
	defer fake.deleteARPEntriesMutex.Unlock()
	fake.DeleteARPEntriesStub = nil
	if fake.deleteARPEntriesReturnsOnCall == nil {
		fake.deleteARPEntriesReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.deleteARPEntriesReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeAgentClient) Drain(arg1 string) (int64, error) {
	fake.drainMutex.Lock()
	ret, specificReturn := fake.drainReturnsOnCall[len(fake.drainArgsForCall)]
	fake.drainArgsForCall = append(fake.drainArgsForCall, struct {
		arg1 string
	}{arg1})
	stub := fake.DrainStub
	fakeReturns := fake.drainReturns
	fake.recordInvocation("Drain", []interface{}{arg1})
	fake.drainMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeAgentClient) DrainCallCount() int {
	fake.drainMutex.RLock()
	defer fake.drainMutex.RUnlock()
	return len(fake.drainArgsForCall)
}

func (fake *FakeAgentClient) DrainCalls(stub func(string) (int64, error)) {
	fake.drainMutex.Lock()
	defer fake.drainMutex.Unlock()
	fake.DrainStub = stub
}

func (fake *FakeAgentClient) DrainArgsForCall(i int) string {
	fake.drainMutex.RLock()
	defer fake.drainMutex.RUnlock()
	argsForCall := fake.drainArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeAgentClient) DrainReturns(result1 int64, result2 error) {
	fake.drainMutex.Lock()
	defer fake.drainMutex.Unlock()
	fake.DrainStub = nil
	fake.drainReturns = struct {
		result1 int64
		result2 error
	}{result1, result2}
}

func (fake *FakeAgentClient) DrainReturnsOnCall(i int, result1 int64, result2 error) {
	fake.drainMutex.Lock()
	defer fake.drainMutex.Unlock()
	fake.DrainStub = nil
	if fake.drainReturnsOnCall == nil {
		fake.drainReturnsOnCall = make(map[int]struct {
			result1 int64
			result2 error
		})
	}
	fake.drainReturnsOnCall[i] = struct {
		result1 int64
		result2 error
	}{result1, result2}
}

func (fake *FakeAgentClient) GetState() (agentclient.AgentState, error) {
	fake.getStateMutex.Lock()
	ret, specificReturn := fake.getStateReturnsOnCall[len(fake.getStateArgsForCall)]
	fake.getStateArgsForCall = append(fake.getStateArgsForCall, struct {
	}{})
	stub := fake.GetStateStub
	fakeReturns := fake.getStateReturns
	fake.recordInvocation("GetState", []interface{}{})
	fake.getStateMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeAgentClient) GetStateCallCount() int {
	fake.getStateMutex.RLock()
	defer fake.getStateMutex.RUnlock()
	return len(fake.getStateArgsForCall)
}

func (fake *FakeAgentClient) GetStateCalls(stub func() (agentclient.AgentState, error)) {
	fake.getStateMutex.Lock()
	defer fake.getStateMutex.Unlock()
	fake.GetStateStub = stub
}

func (fake *FakeAgentClient) GetStateReturns(result1 agentclient.AgentState, result2 error) {
	fake.getStateMutex.Lock()
	defer fake.getStateMutex.Unlock()
	fake.GetStateStub = nil
	fake.getStateReturns = struct {
		result1 agentclient.AgentState
		result2 error
	}{result1, result2}
}

func (fake *FakeAgentClient) GetStateReturnsOnCall(i int, result1 agentclient.AgentState, result2 error) {
	fake.getStateMutex.Lock()
	defer fake.getStateMutex.Unlock()
	fake.GetStateStub = nil
	if fake.getStateReturnsOnCall == nil {
		fake.getStateReturnsOnCall = make(map[int]struct {
			result1 agentclient.AgentState
			result2 error
		})
	}
	fake.getStateReturnsOnCall[i] = struct {
		result1 agentclient.AgentState
		result2 error
	}{result1, result2}
}

func (fake *FakeAgentClient) ListDisk() ([]string, error) {
	fake.listDiskMutex.Lock()
	ret, specificReturn := fake.listDiskReturnsOnCall[len(fake.listDiskArgsForCall)]
	fake.listDiskArgsForCall = append(fake.listDiskArgsForCall, struct {
	}{})
	stub := fake.ListDiskStub
	fakeReturns := fake.listDiskReturns
	fake.recordInvocation("ListDisk", []interface{}{})
	fake.listDiskMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeAgentClient) ListDiskCallCount() int {
	fake.listDiskMutex.RLock()
	defer fake.listDiskMutex.RUnlock()
	return len(fake.listDiskArgsForCall)
}

func (fake *FakeAgentClient) ListDiskCalls(stub func() ([]string, error)) {
	fake.listDiskMutex.Lock()
	defer fake.listDiskMutex.Unlock()
	fake.ListDiskStub = stub
}

func (fake *FakeAgentClient) ListDiskReturns(result1 []string, result2 error) {
	fake.listDiskMutex.Lock()
	defer fake.listDiskMutex.Unlock()
	fake.ListDiskStub = nil
	fake.listDiskReturns = struct {
		result1 []string
		result2 error
	}{result1, result2}
}

func (fake *FakeAgentClient) ListDiskReturnsOnCall(i int, result1 []string, result2 error) {
	fake.listDiskMutex.Lock()
	defer fake.listDiskMutex.Unlock()
	fake.ListDiskStub = nil
	if fake.listDiskReturnsOnCall == nil {
		fake.listDiskReturnsOnCall = make(map[int]struct {
			result1 []string
			result2 error
		})
	}
	fake.listDiskReturnsOnCall[i] = struct {
		result1 []string
		result2 error
	}{result1, result2}
}

func (fake *FakeAgentClient) MigrateDisk() error {
	fake.migrateDiskMutex.Lock()
	ret, specificReturn := fake.migrateDiskReturnsOnCall[len(fake.migrateDiskArgsForCall)]
	fake.migrateDiskArgsForCall = append(fake.migrateDiskArgsForCall, struct {
	}{})
	stub := fake.MigrateDiskStub
	fakeReturns := fake.migrateDiskReturns
	fake.recordInvocation("MigrateDisk", []interface{}{})
	fake.migrateDiskMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeAgentClient) MigrateDiskCallCount() int {
	fake.migrateDiskMutex.RLock()
	defer fake.migrateDiskMutex.RUnlock()
	return len(fake.migrateDiskArgsForCall)
}

func (fake *FakeAgentClient) MigrateDiskCalls(stub func() error) {
	fake.migrateDiskMutex.Lock()
	defer fake.migrateDiskMutex.Unlock()
	fake.MigrateDiskStub = stub
}

func (fake *FakeAgentClient) MigrateDiskReturns(result1 error) {
	fake.migrateDiskMutex.Lock()
	defer fake.migrateDiskMutex.Unlock()
	fake.MigrateDiskStub = nil
	fake.migrateDiskReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeAgentClient) MigrateDiskReturnsOnCall(i int, result1 error) {
	fake.migrateDiskMutex.Lock()
	defer fake.migrateDiskMutex.Unlock()
	fake.MigrateDiskStub = nil
	if fake.migrateDiskReturnsOnCall == nil {
		fake.migrateDiskReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.migrateDiskReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeAgentClient) MountDisk(arg1 string) error {
	fake.mountDiskMutex.Lock()
	ret, specificReturn := fake.mountDiskReturnsOnCall[len(fake.mountDiskArgsForCall)]
	fake.mountDiskArgsForCall = append(fake.mountDiskArgsForCall, struct {
		arg1 string
	}{arg1})
	stub := fake.MountDiskStub
	fakeReturns := fake.mountDiskReturns
	fake.recordInvocation("MountDisk", []interface{}{arg1})
	fake.mountDiskMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeAgentClient) MountDiskCallCount() int {
	fake.mountDiskMutex.RLock()
	defer fake.mountDiskMutex.RUnlock()
	return len(fake.mountDiskArgsForCall)
}

func (fake *FakeAgentClient) MountDiskCalls(stub func(string) error) {
	fake.mountDiskMutex.Lock()
	defer fake.mountDiskMutex.Unlock()
	fake.MountDiskStub = stub
}

func (fake *FakeAgentClient) MountDiskArgsForCall(i int) string {
	fake.mountDiskMutex.RLock()
	defer fake.mountDiskMutex.RUnlock()
	argsForCall := fake.mountDiskArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeAgentClient) MountDiskReturns(result1 error) {
	fake.mountDiskMutex.Lock()
	defer fake.mountDiskMutex.Unlock()
	fake.MountDiskStub = nil
	fake.mountDiskReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeAgentClient) MountDiskReturnsOnCall(i int, result1 error) {
	fake.mountDiskMutex.Lock()
	defer fake.mountDiskMutex.Unlock()
	fake.MountDiskStub = nil
	if fake.mountDiskReturnsOnCall == nil {
		fake.mountDiskReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.mountDiskReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeAgentClient) Ping() (string, error) {
	fake.pingMutex.Lock()
	ret, specificReturn := fake.pingReturnsOnCall[len(fake.pingArgsForCall)]
	fake.pingArgsForCall = append(fake.pingArgsForCall, struct {
	}{})
	stub := fake.PingStub
	fakeReturns := fake.pingReturns
	fake.recordInvocation("Ping", []interface{}{})
	fake.pingMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeAgentClient) PingCallCount() int {
	fake.pingMutex.RLock()
	defer fake.pingMutex.RUnlock()
	return len(fake.pingArgsForCall)
}

func (fake *FakeAgentClient) PingCalls(stub func() (string, error)) {
	fake.pingMutex.Lock()
	defer fake.pingMutex.Unlock()
	fake.PingStub = stub
}

func (fake *FakeAgentClient) PingReturns(result1 string, result2 error) {
	fake.pingMutex.Lock()
	defer fake.pingMutex.Unlock()
	fake.PingStub = nil
	fake.pingReturns = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeAgentClient) PingReturnsOnCall(i int, result1 string, result2 error) {
	fake.pingMutex.Lock()
	defer fake.pingMutex.Unlock()
	fake.PingStub = nil
	if fake.pingReturnsOnCall == nil {
		fake.pingReturnsOnCall = make(map[int]struct {
			result1 string
			result2 error
		})
	}
	fake.pingReturnsOnCall[i] = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeAgentClient) RemovePersistentDisk(arg1 string) error {
	fake.removePersistentDiskMutex.Lock()
	ret, specificReturn := fake.removePersistentDiskReturnsOnCall[len(fake.removePersistentDiskArgsForCall)]
	fake.removePersistentDiskArgsForCall = append(fake.removePersistentDiskArgsForCall, struct {
		arg1 string
	}{arg1})
	stub := fake.RemovePersistentDiskStub
	fakeReturns := fake.removePersistentDiskReturns
	fake.recordInvocation("RemovePersistentDisk", []interface{}{arg1})
	fake.removePersistentDiskMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeAgentClient) RemovePersistentDiskCallCount() int {
	fake.removePersistentDiskMutex.RLock()
	defer fake.removePersistentDiskMutex.RUnlock()
	return len(fake.removePersistentDiskArgsForCall)
}

func (fake *FakeAgentClient) RemovePersistentDiskCalls(stub func(string) error) {
	fake.removePersistentDiskMutex.Lock()
	defer fake.removePersistentDiskMutex.Unlock()
	fake.RemovePersistentDiskStub = stub
}

func (fake *FakeAgentClient) RemovePersistentDiskArgsForCall(i int) string {
	fake.removePersistentDiskMutex.RLock()
	defer fake.removePersistentDiskMutex.RUnlock()
	argsForCall := fake.removePersistentDiskArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeAgentClient) RemovePersistentDiskReturns(result1 error) {
	fake.removePersistentDiskMutex.Lock()
	defer fake.removePersistentDiskMutex.Unlock()
	fake.RemovePersistentDiskStub = nil
	fake.removePersistentDiskReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeAgentClient) RemovePersistentDiskReturnsOnCall(i int, result1 error) {
	fake.removePersistentDiskMutex.Lock()
	defer fake.removePersistentDiskMutex.Unlock()
	fake.RemovePersistentDiskStub = nil
	if fake.removePersistentDiskReturnsOnCall == nil {
		fake.removePersistentDiskReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.removePersistentDiskReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeAgentClient) RunScript(arg1 string, arg2 map[string]interface{}) error {
	fake.runScriptMutex.Lock()
	ret, specificReturn := fake.runScriptReturnsOnCall[len(fake.runScriptArgsForCall)]
	fake.runScriptArgsForCall = append(fake.runScriptArgsForCall, struct {
		arg1 string
		arg2 map[string]interface{}
	}{arg1, arg2})
	stub := fake.RunScriptStub
	fakeReturns := fake.runScriptReturns
	fake.recordInvocation("RunScript", []interface{}{arg1, arg2})
	fake.runScriptMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeAgentClient) RunScriptCallCount() int {
	fake.runScriptMutex.RLock()
	defer fake.runScriptMutex.RUnlock()
	return len(fake.runScriptArgsForCall)
}

func (fake *FakeAgentClient) RunScriptCalls(stub func(string, map[string]interface{}) error) {
	fake.runScriptMutex.Lock()
	defer fake.runScriptMutex.Unlock()
	fake.RunScriptStub = stub
}

func (fake *FakeAgentClient) RunScriptArgsForCall(i int) (string, map[string]interface{}) {
	fake.runScriptMutex.RLock()
	defer fake.runScriptMutex.RUnlock()
	argsForCall := fake.runScriptArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeAgentClient) RunScriptReturns(result1 error) {
	fake.runScriptMutex.Lock()
	defer fake.runScriptMutex.Unlock()
	fake.RunScriptStub = nil
	fake.runScriptReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeAgentClient) RunScriptReturnsOnCall(i int, result1 error) {
	fake.runScriptMutex.Lock()
	defer fake.runScriptMutex.Unlock()
	fake.RunScriptStub = nil
	if fake.runScriptReturnsOnCall == nil {
		fake.runScriptReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.runScriptReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeAgentClient) Start() error {
	fake.startMutex.Lock()
	ret, specificReturn := fake.startReturnsOnCall[len(fake.startArgsForCall)]
	fake.startArgsForCall = append(fake.startArgsForCall, struct {
	}{})
	stub := fake.StartStub
	fakeReturns := fake.startReturns
	fake.recordInvocation("Start", []interface{}{})
	fake.startMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeAgentClient) StartCallCount() int {
	fake.startMutex.RLock()
	defer fake.startMutex.RUnlock()
	return len(fake.startArgsForCall)
}

func (fake *FakeAgentClient) StartCalls(stub func() error) {
	fake.startMutex.Lock()
	defer fake.startMutex.Unlock()
	fake.StartStub = stub
}

func (fake *FakeAgentClient) StartReturns(result1 error) {
	fake.startMutex.Lock()
	defer fake.startMutex.Unlock()
	fake.StartStub = nil
	fake.startReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeAgentClient) StartReturnsOnCall(i int, result1 error) {
	fake.startMutex.Lock()
	defer fake.startMutex.Unlock()
	fake.StartStub = nil
	if fake.startReturnsOnCall == nil {
		fake.startReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.startReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeAgentClient) Stop() error {
	fake.stopMutex.Lock()
	ret, specificReturn := fake.stopReturnsOnCall[len(fake.stopArgsForCall)]
	fake.stopArgsForCall = append(fake.stopArgsForCall, struct {
	}{})
	stub := fake.StopStub
	fakeReturns := fake.stopReturns
	fake.recordInvocation("Stop", []interface{}{})
	fake.stopMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeAgentClient) StopCallCount() int {
	fake.stopMutex.RLock()
	defer fake.stopMutex.RUnlock()
	return len(fake.stopArgsForCall)
}

func (fake *FakeAgentClient) StopCalls(stub func() error) {
	fake.stopMutex.Lock()
	defer fake.stopMutex.Unlock()
	fake.StopStub = stub
}

func (fake *FakeAgentClient) StopReturns(result1 error) {
	fake.stopMutex.Lock()
	defer fake.stopMutex.Unlock()
	fake.StopStub = nil
	fake.stopReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeAgentClient) StopReturnsOnCall(i int, result1 error) {
	fake.stopMutex.Lock()
	defer fake.stopMutex.Unlock()
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

func (fake *FakeAgentClient) SyncDNS(arg1 string, arg2 string, arg3 uint64) (string, error) {
	fake.syncDNSMutex.Lock()
	ret, specificReturn := fake.syncDNSReturnsOnCall[len(fake.syncDNSArgsForCall)]
	fake.syncDNSArgsForCall = append(fake.syncDNSArgsForCall, struct {
		arg1 string
		arg2 string
		arg3 uint64
	}{arg1, arg2, arg3})
	stub := fake.SyncDNSStub
	fakeReturns := fake.syncDNSReturns
	fake.recordInvocation("SyncDNS", []interface{}{arg1, arg2, arg3})
	fake.syncDNSMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeAgentClient) SyncDNSCallCount() int {
	fake.syncDNSMutex.RLock()
	defer fake.syncDNSMutex.RUnlock()
	return len(fake.syncDNSArgsForCall)
}

func (fake *FakeAgentClient) SyncDNSCalls(stub func(string, string, uint64) (string, error)) {
	fake.syncDNSMutex.Lock()
	defer fake.syncDNSMutex.Unlock()
	fake.SyncDNSStub = stub
}

func (fake *FakeAgentClient) SyncDNSArgsForCall(i int) (string, string, uint64) {
	fake.syncDNSMutex.RLock()
	defer fake.syncDNSMutex.RUnlock()
	argsForCall := fake.syncDNSArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeAgentClient) SyncDNSReturns(result1 string, result2 error) {
	fake.syncDNSMutex.Lock()
	defer fake.syncDNSMutex.Unlock()
	fake.SyncDNSStub = nil
	fake.syncDNSReturns = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeAgentClient) SyncDNSReturnsOnCall(i int, result1 string, result2 error) {
	fake.syncDNSMutex.Lock()
	defer fake.syncDNSMutex.Unlock()
	fake.SyncDNSStub = nil
	if fake.syncDNSReturnsOnCall == nil {
		fake.syncDNSReturnsOnCall = make(map[int]struct {
			result1 string
			result2 error
		})
	}
	fake.syncDNSReturnsOnCall[i] = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeAgentClient) UnmountDisk(arg1 string) error {
	fake.unmountDiskMutex.Lock()
	ret, specificReturn := fake.unmountDiskReturnsOnCall[len(fake.unmountDiskArgsForCall)]
	fake.unmountDiskArgsForCall = append(fake.unmountDiskArgsForCall, struct {
		arg1 string
	}{arg1})
	stub := fake.UnmountDiskStub
	fakeReturns := fake.unmountDiskReturns
	fake.recordInvocation("UnmountDisk", []interface{}{arg1})
	fake.unmountDiskMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeAgentClient) UnmountDiskCallCount() int {
	fake.unmountDiskMutex.RLock()
	defer fake.unmountDiskMutex.RUnlock()
	return len(fake.unmountDiskArgsForCall)
}

func (fake *FakeAgentClient) UnmountDiskCalls(stub func(string) error) {
	fake.unmountDiskMutex.Lock()
	defer fake.unmountDiskMutex.Unlock()
	fake.UnmountDiskStub = stub
}

func (fake *FakeAgentClient) UnmountDiskArgsForCall(i int) string {
	fake.unmountDiskMutex.RLock()
	defer fake.unmountDiskMutex.RUnlock()
	argsForCall := fake.unmountDiskArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeAgentClient) UnmountDiskReturns(result1 error) {
	fake.unmountDiskMutex.Lock()
	defer fake.unmountDiskMutex.Unlock()
	fake.UnmountDiskStub = nil
	fake.unmountDiskReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeAgentClient) UnmountDiskReturnsOnCall(i int, result1 error) {
	fake.unmountDiskMutex.Lock()
	defer fake.unmountDiskMutex.Unlock()
	fake.UnmountDiskStub = nil
	if fake.unmountDiskReturnsOnCall == nil {
		fake.unmountDiskReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.unmountDiskReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeAgentClient) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.addPersistentDiskMutex.RLock()
	defer fake.addPersistentDiskMutex.RUnlock()
	fake.applyMutex.RLock()
	defer fake.applyMutex.RUnlock()
	fake.compilePackageMutex.RLock()
	defer fake.compilePackageMutex.RUnlock()
	fake.deleteARPEntriesMutex.RLock()
	defer fake.deleteARPEntriesMutex.RUnlock()
	fake.drainMutex.RLock()
	defer fake.drainMutex.RUnlock()
	fake.getStateMutex.RLock()
	defer fake.getStateMutex.RUnlock()
	fake.listDiskMutex.RLock()
	defer fake.listDiskMutex.RUnlock()
	fake.migrateDiskMutex.RLock()
	defer fake.migrateDiskMutex.RUnlock()
	fake.mountDiskMutex.RLock()
	defer fake.mountDiskMutex.RUnlock()
	fake.pingMutex.RLock()
	defer fake.pingMutex.RUnlock()
	fake.removePersistentDiskMutex.RLock()
	defer fake.removePersistentDiskMutex.RUnlock()
	fake.runScriptMutex.RLock()
	defer fake.runScriptMutex.RUnlock()
	fake.startMutex.RLock()
	defer fake.startMutex.RUnlock()
	fake.stopMutex.RLock()
	defer fake.stopMutex.RUnlock()
	fake.syncDNSMutex.RLock()
	defer fake.syncDNSMutex.RUnlock()
	fake.unmountDiskMutex.RLock()
	defer fake.unmountDiskMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeAgentClient) recordInvocation(key string, args []interface{}) {
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

var _ agentclient.AgentClient = new(FakeAgentClient)
