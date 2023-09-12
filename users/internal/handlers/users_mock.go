// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package handlers

import (
	"context"
	"github.com/mikenai/gowork/internal/models"
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
//			CreateFunc: func(ctx context.Context, name string) (models.User, error) {
//				panic("mock out the Create method")
//			},
//			GetOneFunc: func(ctx context.Context, id string) (models.User, error) {
//				panic("mock out the GetOne method")
//			},
//			UpdateUserFunc: func(ctx context.Context, user models.User) error {
//				panic("mock out the UpdateUser method")
//			},
//		}
//
//		// use mockedUsersService in code that requires UsersService
//		// and then make assertions.
//
//	}
type UsersServiceMock struct {
	// CreateFunc mocks the Create method.
	CreateFunc func(ctx context.Context, name string) (models.User, error)

	// GetOneFunc mocks the GetOne method.
	GetOneFunc func(ctx context.Context, id string) (models.User, error)

	// UpdateUserFunc mocks the UpdateUser method.
	UpdateUserFunc func(ctx context.Context, user models.User) error

	// calls tracks calls to the methods.
	calls struct {
		// Create holds details about calls to the Create method.
		Create []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Name is the name argument value.
			Name string
		}
		// GetOne holds details about calls to the GetOne method.
		GetOne []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// ID is the id argument value.
			ID string
		}
		// UpdateUser holds details about calls to the UpdateUser method.
		UpdateUser []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// User is the user argument value.
			User models.User
		}
	}
	lockCreate     sync.RWMutex
	lockGetOne     sync.RWMutex
	lockUpdateUser sync.RWMutex
}

// Create calls CreateFunc.
func (mock *UsersServiceMock) Create(ctx context.Context, name string) (models.User, error) {
	if mock.CreateFunc == nil {
		panic("UsersServiceMock.CreateFunc: method is nil but UsersService.Create was just called")
	}
	callInfo := struct {
		Ctx  context.Context
		Name string
	}{
		Ctx:  ctx,
		Name: name,
	}
	mock.lockCreate.Lock()
	mock.calls.Create = append(mock.calls.Create, callInfo)
	mock.lockCreate.Unlock()
	return mock.CreateFunc(ctx, name)
}

// CreateCalls gets all the calls that were made to Create.
// Check the length with:
//
//	len(mockedUsersService.CreateCalls())
func (mock *UsersServiceMock) CreateCalls() []struct {
	Ctx  context.Context
	Name string
} {
	var calls []struct {
		Ctx  context.Context
		Name string
	}
	mock.lockCreate.RLock()
	calls = mock.calls.Create
	mock.lockCreate.RUnlock()
	return calls
}

// GetOne calls GetOneFunc.
func (mock *UsersServiceMock) GetOne(ctx context.Context, id string) (models.User, error) {
	if mock.GetOneFunc == nil {
		panic("UsersServiceMock.GetOneFunc: method is nil but UsersService.GetOne was just called")
	}
	callInfo := struct {
		Ctx context.Context
		ID  string
	}{
		Ctx: ctx,
		ID:  id,
	}
	mock.lockGetOne.Lock()
	mock.calls.GetOne = append(mock.calls.GetOne, callInfo)
	mock.lockGetOne.Unlock()
	return mock.GetOneFunc(ctx, id)
}

// GetOneCalls gets all the calls that were made to GetOne.
// Check the length with:
//
//	len(mockedUsersService.GetOneCalls())
func (mock *UsersServiceMock) GetOneCalls() []struct {
	Ctx context.Context
	ID  string
} {
	var calls []struct {
		Ctx context.Context
		ID  string
	}
	mock.lockGetOne.RLock()
	calls = mock.calls.GetOne
	mock.lockGetOne.RUnlock()
	return calls
}

// UpdateUser calls UpdateUserFunc.
func (mock *UsersServiceMock) UpdateUser(ctx context.Context, user models.User) error {
	if mock.UpdateUserFunc == nil {
		panic("UsersServiceMock.UpdateUserFunc: method is nil but UsersService.UpdateUser was just called")
	}
	callInfo := struct {
		Ctx  context.Context
		User models.User
	}{
		Ctx:  ctx,
		User: user,
	}
	mock.lockUpdateUser.Lock()
	mock.calls.UpdateUser = append(mock.calls.UpdateUser, callInfo)
	mock.lockUpdateUser.Unlock()
	return mock.UpdateUserFunc(ctx, user)
}

// UpdateUserCalls gets all the calls that were made to UpdateUser.
// Check the length with:
//
//	len(mockedUsersService.UpdateUserCalls())
func (mock *UsersServiceMock) UpdateUserCalls() []struct {
	Ctx  context.Context
	User models.User
} {
	var calls []struct {
		Ctx  context.Context
		User models.User
	}
	mock.lockUpdateUser.RLock()
	calls = mock.calls.UpdateUser
	mock.lockUpdateUser.RUnlock()
	return calls
}
