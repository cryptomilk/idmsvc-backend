// Code generated by mockery v2.33.2. DO NOT EDIT.

package repository

import (
	model "github.com/podengo-project/idmsvc-backend/internal/domain/model"
	mock "github.com/stretchr/testify/mock"
	gorm "gorm.io/gorm"

	public "github.com/podengo-project/idmsvc-backend/internal/api/public"

	repository "github.com/podengo-project/idmsvc-backend/internal/interface/repository"

	time "time"

	uuid "github.com/google/uuid"
)

// DomainRepository is an autogenerated mock type for the DomainRepository type
type DomainRepository struct {
	mock.Mock
}

// CreateDomainToken provides a mock function with given fields: key, validity, orgID, domainType
func (_m *DomainRepository) CreateDomainToken(key []byte, validity time.Duration, orgID string, domainType public.DomainType) (*repository.DomainRegToken, error) {
	ret := _m.Called(key, validity, orgID, domainType)

	var r0 *repository.DomainRegToken
	var r1 error
	if rf, ok := ret.Get(0).(func([]byte, time.Duration, string, public.DomainType) (*repository.DomainRegToken, error)); ok {
		return rf(key, validity, orgID, domainType)
	}
	if rf, ok := ret.Get(0).(func([]byte, time.Duration, string, public.DomainType) *repository.DomainRegToken); ok {
		r0 = rf(key, validity, orgID, domainType)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*repository.DomainRegToken)
		}
	}

	if rf, ok := ret.Get(1).(func([]byte, time.Duration, string, public.DomainType) error); ok {
		r1 = rf(key, validity, orgID, domainType)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteById provides a mock function with given fields: db, orgID, UUID
func (_m *DomainRepository) DeleteById(db *gorm.DB, orgID string, UUID uuid.UUID) error {
	ret := _m.Called(db, orgID, UUID)

	var r0 error
	if rf, ok := ret.Get(0).(func(*gorm.DB, string, uuid.UUID) error); ok {
		r0 = rf(db, orgID, UUID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// FindByID provides a mock function with given fields: db, orgID, UUID
func (_m *DomainRepository) FindByID(db *gorm.DB, orgID string, UUID uuid.UUID) (*model.Domain, error) {
	ret := _m.Called(db, orgID, UUID)

	var r0 *model.Domain
	var r1 error
	if rf, ok := ret.Get(0).(func(*gorm.DB, string, uuid.UUID) (*model.Domain, error)); ok {
		return rf(db, orgID, UUID)
	}
	if rf, ok := ret.Get(0).(func(*gorm.DB, string, uuid.UUID) *model.Domain); ok {
		r0 = rf(db, orgID, UUID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Domain)
		}
	}

	if rf, ok := ret.Get(1).(func(*gorm.DB, string, uuid.UUID) error); ok {
		r1 = rf(db, orgID, UUID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// List provides a mock function with given fields: db, orgID, offset, limit
func (_m *DomainRepository) List(db *gorm.DB, orgID string, offset int, limit int) ([]model.Domain, int64, error) {
	ret := _m.Called(db, orgID, offset, limit)

	var r0 []model.Domain
	var r1 int64
	var r2 error
	if rf, ok := ret.Get(0).(func(*gorm.DB, string, int, int) ([]model.Domain, int64, error)); ok {
		return rf(db, orgID, offset, limit)
	}
	if rf, ok := ret.Get(0).(func(*gorm.DB, string, int, int) []model.Domain); ok {
		r0 = rf(db, orgID, offset, limit)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]model.Domain)
		}
	}

	if rf, ok := ret.Get(1).(func(*gorm.DB, string, int, int) int64); ok {
		r1 = rf(db, orgID, offset, limit)
	} else {
		r1 = ret.Get(1).(int64)
	}

	if rf, ok := ret.Get(2).(func(*gorm.DB, string, int, int) error); ok {
		r2 = rf(db, orgID, offset, limit)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// Register provides a mock function with given fields: db, orgID, data
func (_m *DomainRepository) Register(db *gorm.DB, orgID string, data *model.Domain) error {
	ret := _m.Called(db, orgID, data)

	var r0 error
	if rf, ok := ret.Get(0).(func(*gorm.DB, string, *model.Domain) error); ok {
		r0 = rf(db, orgID, data)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateAgent provides a mock function with given fields: db, orgID, data
func (_m *DomainRepository) UpdateAgent(db *gorm.DB, orgID string, data *model.Domain) error {
	ret := _m.Called(db, orgID, data)

	var r0 error
	if rf, ok := ret.Get(0).(func(*gorm.DB, string, *model.Domain) error); ok {
		r0 = rf(db, orgID, data)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateUser provides a mock function with given fields: db, orgID, data
func (_m *DomainRepository) UpdateUser(db *gorm.DB, orgID string, data *model.Domain) error {
	ret := _m.Called(db, orgID, data)

	var r0 error
	if rf, ok := ret.Get(0).(func(*gorm.DB, string, *model.Domain) error); ok {
		r0 = rf(db, orgID, data)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewDomainRepository creates a new instance of DomainRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewDomainRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *DomainRepository {
	mock := &DomainRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
