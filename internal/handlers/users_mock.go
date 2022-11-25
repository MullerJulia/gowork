// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package handlers

import (
	"sync"
)

// Ensure, that UsersServiceMock does implement UsersService.
// If this is not the case, regenerate this file with moq.
var _ UsersService = &UsersServiceMock{}

// UsersServiceMock is a mock implementation of UsersService.
//
//	func TestSomethingThatUsesUsersService(t *testing.T) {
//
//		// make and configure a mocked UsersService
//		mockedUsersService := &UsersServiceMock{
//			CreateFunc: func(name string) (User, error) {
//				panic("mock out the Create method")
//			},
//		}
//
//		// use mockedUsersService in code that requires UsersService
//		// and then make assertions.
//
//	}
type UsersServiceMock struct {
	// CreateFunc mocks the Create method.
	CreateFunc func(name string) (User, error)

	// calls tracks calls to the methods.
	calls struct {
		// Create holds details about calls to the Create method.
		Create []struct {
			// Name is the name argument value.
			Name string
		}
	}
	lockCreate sync.RWMutex
}

// Create calls CreateFunc.
func (mock *UsersServiceMock) Create(name string) (User, error) {
	if mock.CreateFunc == nil {
		panic("UsersServiceMock.CreateFunc: method is nil but UsersService.Create was just called")
	}
	callInfo := struct {
		Name string
	}{
		Name: name,
	}
	mock.lockCreate.Lock()
	mock.calls.Create = append(mock.calls.Create, callInfo)
	mock.lockCreate.Unlock()
	return mock.CreateFunc(name)
}

// CreateCalls gets all the calls that were made to Create.
// Check the length with:
//
//	len(mockedUsersService.CreateCalls())
func (mock *UsersServiceMock) CreateCalls() []struct {
	Name string
} {
	var calls []struct {
		Name string
	}
	mock.lockCreate.RLock()
	calls = mock.calls.Create
	mock.lockCreate.RUnlock()
	return calls
}
