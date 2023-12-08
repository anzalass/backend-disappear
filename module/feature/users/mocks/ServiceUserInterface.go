// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import (
	entities "github.com/capstone-kelompok-7/backend-disappear/module/entities"
	dto "github.com/capstone-kelompok-7/backend-disappear/module/feature/users/dto"

	mock "github.com/stretchr/testify/mock"
)

// ServiceUserInterface is an autogenerated mock type for the ServiceUserInterface type
type ServiceUserInterface struct {
	mock.Mock
}

// AddUserPreference provides a mock function with given fields: userID, request
func (_m *ServiceUserInterface) AddUserPreference(userID uint64, request *dto.UserPreferenceRequest) (*entities.UserModels, error) {
	ret := _m.Called(userID, request)

	var r0 *entities.UserModels
	var r1 error
	if rf, ok := ret.Get(0).(func(uint64, *dto.UserPreferenceRequest) (*entities.UserModels, error)); ok {
		return rf(userID, request)
	}
	if rf, ok := ret.Get(0).(func(uint64, *dto.UserPreferenceRequest) *entities.UserModels); ok {
		r0 = rf(userID, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entities.UserModels)
		}
	}

	if rf, ok := ret.Get(1).(func(uint64, *dto.UserPreferenceRequest) error); ok {
		r1 = rf(userID, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CalculatePaginationValues provides a mock function with given fields: page, totalItems, perPage
func (_m *ServiceUserInterface) CalculatePaginationValues(page int, totalItems int, perPage int) (int, int) {
	ret := _m.Called(page, totalItems, perPage)

	var r0 int
	var r1 int
	if rf, ok := ret.Get(0).(func(int, int, int) (int, int)); ok {
		return rf(page, totalItems, perPage)
	}
	if rf, ok := ret.Get(0).(func(int, int, int) int); ok {
		r0 = rf(page, totalItems, perPage)
	} else {
		r0 = ret.Get(0).(int)
	}

	if rf, ok := ret.Get(1).(func(int, int, int) int); ok {
		r1 = rf(page, totalItems, perPage)
	} else {
		r1 = ret.Get(1).(int)
	}

	return r0, r1
}

// ChangePassword provides a mock function with given fields: userID, updateRequest
func (_m *ServiceUserInterface) ChangePassword(userID uint64, updateRequest dto.UpdatePasswordRequest) error {
	ret := _m.Called(userID, updateRequest)

	var r0 error
	if rf, ok := ret.Get(0).(func(uint64, dto.UpdatePasswordRequest) error); ok {
		r0 = rf(userID, updateRequest)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteAccount provides a mock function with given fields: userID
func (_m *ServiceUserInterface) DeleteAccount(userID uint64) error {
	ret := _m.Called(userID)

	var r0 error
	if rf, ok := ret.Get(0).(func(uint64) error); ok {
		r0 = rf(userID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// EditProfile provides a mock function with given fields: userID, updatedData
func (_m *ServiceUserInterface) EditProfile(userID uint64, updatedData dto.EditProfileRequest) (*entities.UserModels, error) {
	ret := _m.Called(userID, updatedData)

	var r0 *entities.UserModels
	var r1 error
	if rf, ok := ret.Get(0).(func(uint64, dto.EditProfileRequest) (*entities.UserModels, error)); ok {
		return rf(userID, updatedData)
	}
	if rf, ok := ret.Get(0).(func(uint64, dto.EditProfileRequest) *entities.UserModels); ok {
		r0 = rf(userID, updatedData)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entities.UserModels)
		}
	}

	if rf, ok := ret.Get(1).(func(uint64, dto.EditProfileRequest) error); ok {
		r1 = rf(userID, updatedData)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAllEnvironmentsIssues provides a mock function with given fields:
func (_m *ServiceUserInterface) GetAllEnvironmentsIssues() ([]*entities.EnvironmentIssuesModels, error) {
	ret := _m.Called()

	var r0 []*entities.EnvironmentIssuesModels
	var r1 error
	if rf, ok := ret.Get(0).(func() ([]*entities.EnvironmentIssuesModels, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() []*entities.EnvironmentIssuesModels); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*entities.EnvironmentIssuesModels)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAllUsers provides a mock function with given fields: page, perPage
func (_m *ServiceUserInterface) GetAllUsers(page int, perPage int) ([]*entities.UserModels, int64, error) {
	ret := _m.Called(page, perPage)

	var r0 []*entities.UserModels
	var r1 int64
	var r2 error
	if rf, ok := ret.Get(0).(func(int, int) ([]*entities.UserModels, int64, error)); ok {
		return rf(page, perPage)
	}
	if rf, ok := ret.Get(0).(func(int, int) []*entities.UserModels); ok {
		r0 = rf(page, perPage)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*entities.UserModels)
		}
	}

	if rf, ok := ret.Get(1).(func(int, int) int64); ok {
		r1 = rf(page, perPage)
	} else {
		r1 = ret.Get(1).(int64)
	}

	if rf, ok := ret.Get(2).(func(int, int) error); ok {
		r2 = rf(page, perPage)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// GetLeaderboardByExp provides a mock function with given fields: limit
func (_m *ServiceUserInterface) GetLeaderboardByExp(limit int) ([]*entities.UserModels, error) {
	ret := _m.Called(limit)

	var r0 []*entities.UserModels
	var r1 error
	if rf, ok := ret.Get(0).(func(int) ([]*entities.UserModels, error)); ok {
		return rf(limit)
	}
	if rf, ok := ret.Get(0).(func(int) []*entities.UserModels); ok {
		r0 = rf(limit)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*entities.UserModels)
		}
	}

	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(limit)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetNextPage provides a mock function with given fields: currentPage, totalPages
func (_m *ServiceUserInterface) GetNextPage(currentPage int, totalPages int) int {
	ret := _m.Called(currentPage, totalPages)

	var r0 int
	if rf, ok := ret.Get(0).(func(int, int) int); ok {
		r0 = rf(currentPage, totalPages)
	} else {
		r0 = ret.Get(0).(int)
	}

	return r0
}

// GetPrevPage provides a mock function with given fields: currentPage
func (_m *ServiceUserInterface) GetPrevPage(currentPage int) int {
	ret := _m.Called(currentPage)

	var r0 int
	if rf, ok := ret.Get(0).(func(int) int); ok {
		r0 = rf(currentPage)
	} else {
		r0 = ret.Get(0).(int)
	}

	return r0
}

// GetUserChallengeActivity provides a mock function with given fields: userID
func (_m *ServiceUserInterface) GetUserChallengeActivity(userID uint64) (int, int, int, error) {
	ret := _m.Called(userID)

	var r0 int
	var r1 int
	var r2 int
	var r3 error
	if rf, ok := ret.Get(0).(func(uint64) (int, int, int, error)); ok {
		return rf(userID)
	}
	if rf, ok := ret.Get(0).(func(uint64) int); ok {
		r0 = rf(userID)
	} else {
		r0 = ret.Get(0).(int)
	}

	if rf, ok := ret.Get(1).(func(uint64) int); ok {
		r1 = rf(userID)
	} else {
		r1 = ret.Get(1).(int)
	}

	if rf, ok := ret.Get(2).(func(uint64) int); ok {
		r2 = rf(userID)
	} else {
		r2 = ret.Get(2).(int)
	}

	if rf, ok := ret.Get(3).(func(uint64) error); ok {
		r3 = rf(userID)
	} else {
		r3 = ret.Error(3)
	}

	return r0, r1, r2, r3
}

// GetUserLevel provides a mock function with given fields: userID
func (_m *ServiceUserInterface) GetUserLevel(userID uint64) (string, error) {
	ret := _m.Called(userID)

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(uint64) (string, error)); ok {
		return rf(userID)
	}
	if rf, ok := ret.Get(0).(func(uint64) string); ok {
		r0 = rf(userID)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(uint64) error); ok {
		r1 = rf(userID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetUserProfile provides a mock function with given fields: userID
func (_m *ServiceUserInterface) GetUserProfile(userID uint64) (*entities.UserModels, error) {
	ret := _m.Called(userID)

	var r0 *entities.UserModels
	var r1 error
	if rf, ok := ret.Get(0).(func(uint64) (*entities.UserModels, error)); ok {
		return rf(userID)
	}
	if rf, ok := ret.Get(0).(func(uint64) *entities.UserModels); ok {
		r0 = rf(userID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entities.UserModels)
		}
	}

	if rf, ok := ret.Get(1).(func(uint64) error); ok {
		r1 = rf(userID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetUserTransactionActivity provides a mock function with given fields: userID
func (_m *ServiceUserInterface) GetUserTransactionActivity(userID uint64) (int, int, int, error) {
	ret := _m.Called(userID)

	var r0 int
	var r1 int
	var r2 int
	var r3 error
	if rf, ok := ret.Get(0).(func(uint64) (int, int, int, error)); ok {
		return rf(userID)
	}
	if rf, ok := ret.Get(0).(func(uint64) int); ok {
		r0 = rf(userID)
	} else {
		r0 = ret.Get(0).(int)
	}

	if rf, ok := ret.Get(1).(func(uint64) int); ok {
		r1 = rf(userID)
	} else {
		r1 = ret.Get(1).(int)
	}

	if rf, ok := ret.Get(2).(func(uint64) int); ok {
		r2 = rf(userID)
	} else {
		r2 = ret.Get(2).(int)
	}

	if rf, ok := ret.Get(3).(func(uint64) error); ok {
		r3 = rf(userID)
	} else {
		r3 = ret.Error(3)
	}

	return r0, r1, r2, r3
}

// GetUsersByEmail provides a mock function with given fields: email
func (_m *ServiceUserInterface) GetUsersByEmail(email string) (*entities.UserModels, error) {
	ret := _m.Called(email)

	var r0 *entities.UserModels
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (*entities.UserModels, error)); ok {
		return rf(email)
	}
	if rf, ok := ret.Get(0).(func(string) *entities.UserModels); ok {
		r0 = rf(email)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entities.UserModels)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(email)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetUsersById provides a mock function with given fields: userId
func (_m *ServiceUserInterface) GetUsersById(userId uint64) (*entities.UserModels, error) {
	ret := _m.Called(userId)

	var r0 *entities.UserModels
	var r1 error
	if rf, ok := ret.Get(0).(func(uint64) (*entities.UserModels, error)); ok {
		return rf(userId)
	}
	if rf, ok := ret.Get(0).(func(uint64) *entities.UserModels); ok {
		r0 = rf(userId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entities.UserModels)
		}
	}

	if rf, ok := ret.Get(1).(func(uint64) error); ok {
		r1 = rf(userId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetUsersByLevel provides a mock function with given fields: page, perPage, level
func (_m *ServiceUserInterface) GetUsersByLevel(page int, perPage int, level string) ([]*entities.UserModels, int64, error) {
	ret := _m.Called(page, perPage, level)

	var r0 []*entities.UserModels
	var r1 int64
	var r2 error
	if rf, ok := ret.Get(0).(func(int, int, string) ([]*entities.UserModels, int64, error)); ok {
		return rf(page, perPage, level)
	}
	if rf, ok := ret.Get(0).(func(int, int, string) []*entities.UserModels); ok {
		r0 = rf(page, perPage, level)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*entities.UserModels)
		}
	}

	if rf, ok := ret.Get(1).(func(int, int, string) int64); ok {
		r1 = rf(page, perPage, level)
	} else {
		r1 = ret.Get(1).(int64)
	}

	if rf, ok := ret.Get(2).(func(int, int, string) error); ok {
		r2 = rf(page, perPage, level)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// GetUsersByName provides a mock function with given fields: page, perPage, name
func (_m *ServiceUserInterface) GetUsersByName(page int, perPage int, name string) ([]*entities.UserModels, int64, error) {
	ret := _m.Called(page, perPage, name)

	var r0 []*entities.UserModels
	var r1 int64
	var r2 error
	if rf, ok := ret.Get(0).(func(int, int, string) ([]*entities.UserModels, int64, error)); ok {
		return rf(page, perPage, name)
	}
	if rf, ok := ret.Get(0).(func(int, int, string) []*entities.UserModels); ok {
		r0 = rf(page, perPage, name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*entities.UserModels)
		}
	}

	if rf, ok := ret.Get(1).(func(int, int, string) int64); ok {
		r1 = rf(page, perPage, name)
	} else {
		r1 = ret.Get(1).(int64)
	}

	if rf, ok := ret.Get(2).(func(int, int, string) error); ok {
		r2 = rf(page, perPage, name)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// GetUsersBySearchAndFilter provides a mock function with given fields: page, perPage, search, levelFilter
func (_m *ServiceUserInterface) GetUsersBySearchAndFilter(page int, perPage int, search string, levelFilter string) ([]*entities.UserModels, int64, error) {
	ret := _m.Called(page, perPage, search, levelFilter)

	var r0 []*entities.UserModels
	var r1 int64
	var r2 error
	if rf, ok := ret.Get(0).(func(int, int, string, string) ([]*entities.UserModels, int64, error)); ok {
		return rf(page, perPage, search, levelFilter)
	}
	if rf, ok := ret.Get(0).(func(int, int, string, string) []*entities.UserModels); ok {
		r0 = rf(page, perPage, search, levelFilter)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*entities.UserModels)
		}
	}

	if rf, ok := ret.Get(1).(func(int, int, string, string) int64); ok {
		r1 = rf(page, perPage, search, levelFilter)
	} else {
		r1 = ret.Get(1).(int64)
	}

	if rf, ok := ret.Get(2).(func(int, int, string, string) error); ok {
		r2 = rf(page, perPage, search, levelFilter)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// UpdateUserChallengeFollow provides a mock function with given fields: userID, totalChallenge
func (_m *ServiceUserInterface) UpdateUserChallengeFollow(userID uint64, totalChallenge uint64) (*entities.UserModels, error) {
	ret := _m.Called(userID, totalChallenge)

	var r0 *entities.UserModels
	var r1 error
	if rf, ok := ret.Get(0).(func(uint64, uint64) (*entities.UserModels, error)); ok {
		return rf(userID, totalChallenge)
	}
	if rf, ok := ret.Get(0).(func(uint64, uint64) *entities.UserModels); ok {
		r0 = rf(userID, totalChallenge)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entities.UserModels)
		}
	}

	if rf, ok := ret.Get(1).(func(uint64, uint64) error); ok {
		r1 = rf(userID, totalChallenge)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateUserContribution provides a mock function with given fields: userID, gramPlastic
func (_m *ServiceUserInterface) UpdateUserContribution(userID uint64, gramPlastic uint64) (*entities.UserModels, error) {
	ret := _m.Called(userID, gramPlastic)

	var r0 *entities.UserModels
	var r1 error
	if rf, ok := ret.Get(0).(func(uint64, uint64) (*entities.UserModels, error)); ok {
		return rf(userID, gramPlastic)
	}
	if rf, ok := ret.Get(0).(func(uint64, uint64) *entities.UserModels); ok {
		r0 = rf(userID, gramPlastic)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entities.UserModels)
		}
	}

	if rf, ok := ret.Get(1).(func(uint64, uint64) error); ok {
		r1 = rf(userID, gramPlastic)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateUserExp provides a mock function with given fields: userID, exp
func (_m *ServiceUserInterface) UpdateUserExp(userID uint64, exp uint64) (*entities.UserModels, error) {
	ret := _m.Called(userID, exp)

	var r0 *entities.UserModels
	var r1 error
	if rf, ok := ret.Get(0).(func(uint64, uint64) (*entities.UserModels, error)); ok {
		return rf(userID, exp)
	}
	if rf, ok := ret.Get(0).(func(uint64, uint64) *entities.UserModels); ok {
		r0 = rf(userID, exp)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entities.UserModels)
		}
	}

	if rf, ok := ret.Get(1).(func(uint64, uint64) error); ok {
		r1 = rf(userID, exp)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ValidatePassword provides a mock function with given fields: userID, oldPassword, newPassword, confirmPassword
func (_m *ServiceUserInterface) ValidatePassword(userID uint64, oldPassword string, newPassword string, confirmPassword string) error {
	ret := _m.Called(userID, oldPassword, newPassword, confirmPassword)

	var r0 error
	if rf, ok := ret.Get(0).(func(uint64, string, string, string) error); ok {
		r0 = rf(userID, oldPassword, newPassword, confirmPassword)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewServiceUserInterface creates a new instance of ServiceUserInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewServiceUserInterface(t interface {
	mock.TestingT
	Cleanup(func())
}) *ServiceUserInterface {
	mock := &ServiceUserInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
