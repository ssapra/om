// This file was generated by counterfeiter
package fakes

import "sync"

type Extractor struct {
	ExtractMetadataStub        func(string) (string, string, error)
	extractMetadataMutex       sync.RWMutex
	extractMetadataArgsForCall []struct {
		arg1 string
	}
	extractMetadataReturns struct {
		result1 string
		result2 string
		result3 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *Extractor) ExtractMetadata(arg1 string) (string, string, error) {
	fake.extractMetadataMutex.Lock()
	fake.extractMetadataArgsForCall = append(fake.extractMetadataArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("ExtractMetadata", []interface{}{arg1})
	fake.extractMetadataMutex.Unlock()
	if fake.ExtractMetadataStub != nil {
		return fake.ExtractMetadataStub(arg1)
	} else {
		return fake.extractMetadataReturns.result1, fake.extractMetadataReturns.result2, fake.extractMetadataReturns.result3
	}
}

func (fake *Extractor) ExtractMetadataCallCount() int {
	fake.extractMetadataMutex.RLock()
	defer fake.extractMetadataMutex.RUnlock()
	return len(fake.extractMetadataArgsForCall)
}

func (fake *Extractor) ExtractMetadataArgsForCall(i int) string {
	fake.extractMetadataMutex.RLock()
	defer fake.extractMetadataMutex.RUnlock()
	return fake.extractMetadataArgsForCall[i].arg1
}

func (fake *Extractor) ExtractMetadataReturns(result1 string, result2 string, result3 error) {
	fake.ExtractMetadataStub = nil
	fake.extractMetadataReturns = struct {
		result1 string
		result2 string
		result3 error
	}{result1, result2, result3}
}

func (fake *Extractor) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.extractMetadataMutex.RLock()
	defer fake.extractMetadataMutex.RUnlock()
	return fake.invocations
}

func (fake *Extractor) recordInvocation(key string, args []interface{}) {
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
