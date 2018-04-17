// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	"sync"

	"github.com/pivotal-cf/om/api"
)

type SetErrandStateService struct {
	GetStagedProductByNameStub        func(productName string) (api.StagedProductsFindOutput, error)
	getStagedProductByNameMutex       sync.RWMutex
	getStagedProductByNameArgsForCall []struct {
		productName string
	}
	getStagedProductByNameReturns struct {
		result1 api.StagedProductsFindOutput
		result2 error
	}
	getStagedProductByNameReturnsOnCall map[int]struct {
		result1 api.StagedProductsFindOutput
		result2 error
	}
	UpdateStagedProductErrandsStub        func(productID, errandName string, postDeployState, preDeleteState interface{}) error
	updateStagedProductErrandsMutex       sync.RWMutex
	updateStagedProductErrandsArgsForCall []struct {
		productID       string
		errandName      string
		postDeployState interface{}
		preDeleteState  interface{}
	}
	updateStagedProductErrandsReturns struct {
		result1 error
	}
	updateStagedProductErrandsReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *SetErrandStateService) GetStagedProductByName(productName string) (api.StagedProductsFindOutput, error) {
	fake.getStagedProductByNameMutex.Lock()
	ret, specificReturn := fake.getStagedProductByNameReturnsOnCall[len(fake.getStagedProductByNameArgsForCall)]
	fake.getStagedProductByNameArgsForCall = append(fake.getStagedProductByNameArgsForCall, struct {
		productName string
	}{productName})
	fake.recordInvocation("GetStagedProductByName", []interface{}{productName})
	fake.getStagedProductByNameMutex.Unlock()
	if fake.GetStagedProductByNameStub != nil {
		return fake.GetStagedProductByNameStub(productName)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.getStagedProductByNameReturns.result1, fake.getStagedProductByNameReturns.result2
}

func (fake *SetErrandStateService) GetStagedProductByNameCallCount() int {
	fake.getStagedProductByNameMutex.RLock()
	defer fake.getStagedProductByNameMutex.RUnlock()
	return len(fake.getStagedProductByNameArgsForCall)
}

func (fake *SetErrandStateService) GetStagedProductByNameArgsForCall(i int) string {
	fake.getStagedProductByNameMutex.RLock()
	defer fake.getStagedProductByNameMutex.RUnlock()
	return fake.getStagedProductByNameArgsForCall[i].productName
}

func (fake *SetErrandStateService) GetStagedProductByNameReturns(result1 api.StagedProductsFindOutput, result2 error) {
	fake.GetStagedProductByNameStub = nil
	fake.getStagedProductByNameReturns = struct {
		result1 api.StagedProductsFindOutput
		result2 error
	}{result1, result2}
}

func (fake *SetErrandStateService) GetStagedProductByNameReturnsOnCall(i int, result1 api.StagedProductsFindOutput, result2 error) {
	fake.GetStagedProductByNameStub = nil
	if fake.getStagedProductByNameReturnsOnCall == nil {
		fake.getStagedProductByNameReturnsOnCall = make(map[int]struct {
			result1 api.StagedProductsFindOutput
			result2 error
		})
	}
	fake.getStagedProductByNameReturnsOnCall[i] = struct {
		result1 api.StagedProductsFindOutput
		result2 error
	}{result1, result2}
}

func (fake *SetErrandStateService) UpdateStagedProductErrands(productID string, errandName string, postDeployState interface{}, preDeleteState interface{}) error {
	fake.updateStagedProductErrandsMutex.Lock()
	ret, specificReturn := fake.updateStagedProductErrandsReturnsOnCall[len(fake.updateStagedProductErrandsArgsForCall)]
	fake.updateStagedProductErrandsArgsForCall = append(fake.updateStagedProductErrandsArgsForCall, struct {
		productID       string
		errandName      string
		postDeployState interface{}
		preDeleteState  interface{}
	}{productID, errandName, postDeployState, preDeleteState})
	fake.recordInvocation("UpdateStagedProductErrands", []interface{}{productID, errandName, postDeployState, preDeleteState})
	fake.updateStagedProductErrandsMutex.Unlock()
	if fake.UpdateStagedProductErrandsStub != nil {
		return fake.UpdateStagedProductErrandsStub(productID, errandName, postDeployState, preDeleteState)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.updateStagedProductErrandsReturns.result1
}

func (fake *SetErrandStateService) UpdateStagedProductErrandsCallCount() int {
	fake.updateStagedProductErrandsMutex.RLock()
	defer fake.updateStagedProductErrandsMutex.RUnlock()
	return len(fake.updateStagedProductErrandsArgsForCall)
}

func (fake *SetErrandStateService) UpdateStagedProductErrandsArgsForCall(i int) (string, string, interface{}, interface{}) {
	fake.updateStagedProductErrandsMutex.RLock()
	defer fake.updateStagedProductErrandsMutex.RUnlock()
	return fake.updateStagedProductErrandsArgsForCall[i].productID, fake.updateStagedProductErrandsArgsForCall[i].errandName, fake.updateStagedProductErrandsArgsForCall[i].postDeployState, fake.updateStagedProductErrandsArgsForCall[i].preDeleteState
}

func (fake *SetErrandStateService) UpdateStagedProductErrandsReturns(result1 error) {
	fake.UpdateStagedProductErrandsStub = nil
	fake.updateStagedProductErrandsReturns = struct {
		result1 error
	}{result1}
}

func (fake *SetErrandStateService) UpdateStagedProductErrandsReturnsOnCall(i int, result1 error) {
	fake.UpdateStagedProductErrandsStub = nil
	if fake.updateStagedProductErrandsReturnsOnCall == nil {
		fake.updateStagedProductErrandsReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.updateStagedProductErrandsReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *SetErrandStateService) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.getStagedProductByNameMutex.RLock()
	defer fake.getStagedProductByNameMutex.RUnlock()
	fake.updateStagedProductErrandsMutex.RLock()
	defer fake.updateStagedProductErrandsMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *SetErrandStateService) recordInvocation(key string, args []interface{}) {
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
