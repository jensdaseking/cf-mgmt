// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	"sync"

	uaaa "github.com/cloudfoundry-community/go-uaa"
)

type FakeUaa struct {
	CreateUserStub        func(uaaa.User) (*uaaa.User, error)
	createUserMutex       sync.RWMutex
	createUserArgsForCall []struct {
		arg1 uaaa.User
	}
	createUserReturns struct {
		result1 *uaaa.User
		result2 error
	}
	createUserReturnsOnCall map[int]struct {
		result1 *uaaa.User
		result2 error
	}
	ListAllUsersStub        func(string, string, string, uaaa.SortOrder) ([]uaaa.User, error)
	listAllUsersMutex       sync.RWMutex
	listAllUsersArgsForCall []struct {
		arg1 string
		arg2 string
		arg3 string
		arg4 uaaa.SortOrder
	}
	listAllUsersReturns struct {
		result1 []uaaa.User
		result2 error
	}
	listAllUsersReturnsOnCall map[int]struct {
		result1 []uaaa.User
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeUaa) CreateUser(arg1 uaaa.User) (*uaaa.User, error) {
	fake.createUserMutex.Lock()
	ret, specificReturn := fake.createUserReturnsOnCall[len(fake.createUserArgsForCall)]
	fake.createUserArgsForCall = append(fake.createUserArgsForCall, struct {
		arg1 uaaa.User
	}{arg1})
	fake.recordInvocation("CreateUser", []interface{}{arg1})
	fake.createUserMutex.Unlock()
	if fake.CreateUserStub != nil {
		return fake.CreateUserStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.createUserReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeUaa) CreateUserCallCount() int {
	fake.createUserMutex.RLock()
	defer fake.createUserMutex.RUnlock()
	return len(fake.createUserArgsForCall)
}

func (fake *FakeUaa) CreateUserCalls(stub func(uaaa.User) (*uaaa.User, error)) {
	fake.createUserMutex.Lock()
	defer fake.createUserMutex.Unlock()
	fake.CreateUserStub = stub
}

func (fake *FakeUaa) CreateUserArgsForCall(i int) uaaa.User {
	fake.createUserMutex.RLock()
	defer fake.createUserMutex.RUnlock()
	argsForCall := fake.createUserArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeUaa) CreateUserReturns(result1 *uaaa.User, result2 error) {
	fake.createUserMutex.Lock()
	defer fake.createUserMutex.Unlock()
	fake.CreateUserStub = nil
	fake.createUserReturns = struct {
		result1 *uaaa.User
		result2 error
	}{result1, result2}
}

func (fake *FakeUaa) CreateUserReturnsOnCall(i int, result1 *uaaa.User, result2 error) {
	fake.createUserMutex.Lock()
	defer fake.createUserMutex.Unlock()
	fake.CreateUserStub = nil
	if fake.createUserReturnsOnCall == nil {
		fake.createUserReturnsOnCall = make(map[int]struct {
			result1 *uaaa.User
			result2 error
		})
	}
	fake.createUserReturnsOnCall[i] = struct {
		result1 *uaaa.User
		result2 error
	}{result1, result2}
}

func (fake *FakeUaa) ListAllUsers(arg1 string, arg2 string, arg3 string, arg4 uaaa.SortOrder) ([]uaaa.User, error) {
	fake.listAllUsersMutex.Lock()
	ret, specificReturn := fake.listAllUsersReturnsOnCall[len(fake.listAllUsersArgsForCall)]
	fake.listAllUsersArgsForCall = append(fake.listAllUsersArgsForCall, struct {
		arg1 string
		arg2 string
		arg3 string
		arg4 uaaa.SortOrder
	}{arg1, arg2, arg3, arg4})
	fake.recordInvocation("ListAllUsers", []interface{}{arg1, arg2, arg3, arg4})
	fake.listAllUsersMutex.Unlock()
	if fake.ListAllUsersStub != nil {
		return fake.ListAllUsersStub(arg1, arg2, arg3, arg4)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.listAllUsersReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeUaa) ListAllUsersCallCount() int {
	fake.listAllUsersMutex.RLock()
	defer fake.listAllUsersMutex.RUnlock()
	return len(fake.listAllUsersArgsForCall)
}

func (fake *FakeUaa) ListAllUsersCalls(stub func(string, string, string, uaaa.SortOrder) ([]uaaa.User, error)) {
	fake.listAllUsersMutex.Lock()
	defer fake.listAllUsersMutex.Unlock()
	fake.ListAllUsersStub = stub
}

func (fake *FakeUaa) ListAllUsersArgsForCall(i int) (string, string, string, uaaa.SortOrder) {
	fake.listAllUsersMutex.RLock()
	defer fake.listAllUsersMutex.RUnlock()
	argsForCall := fake.listAllUsersArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3, argsForCall.arg4
}

func (fake *FakeUaa) ListAllUsersReturns(result1 []uaaa.User, result2 error) {
	fake.listAllUsersMutex.Lock()
	defer fake.listAllUsersMutex.Unlock()
	fake.ListAllUsersStub = nil
	fake.listAllUsersReturns = struct {
		result1 []uaaa.User
		result2 error
	}{result1, result2}
}

func (fake *FakeUaa) ListAllUsersReturnsOnCall(i int, result1 []uaaa.User, result2 error) {
	fake.listAllUsersMutex.Lock()
	defer fake.listAllUsersMutex.Unlock()
	fake.ListAllUsersStub = nil
	if fake.listAllUsersReturnsOnCall == nil {
		fake.listAllUsersReturnsOnCall = make(map[int]struct {
			result1 []uaaa.User
			result2 error
		})
	}
	fake.listAllUsersReturnsOnCall[i] = struct {
		result1 []uaaa.User
		result2 error
	}{result1, result2}
}

func (fake *FakeUaa) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.createUserMutex.RLock()
	defer fake.createUserMutex.RUnlock()
	fake.listAllUsersMutex.RLock()
	defer fake.listAllUsersMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeUaa) recordInvocation(key string, args []interface{}) {
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
