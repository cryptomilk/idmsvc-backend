// Code generated by mockery v2.33.2. DO NOT EDIT.

package repository

import (
	jwk "github.com/lestrrat-go/jwx/v2/jwk"
	interactor "github.com/podengo-project/idmsvc-backend/internal/interface/interactor"
	gorm "gorm.io/gorm"

	mock "github.com/stretchr/testify/mock"

	model "github.com/podengo-project/idmsvc-backend/internal/domain/model"
)

// HostRepository is an autogenerated mock type for the HostRepository type
type HostRepository struct {
	mock.Mock
}

// MatchDomain provides a mock function with given fields: db, options
func (_m *HostRepository) MatchDomain(db *gorm.DB, options *interactor.HostConfOptions) (*model.Domain, error) {
	ret := _m.Called(db, options)

	var r0 *model.Domain
	var r1 error
	if rf, ok := ret.Get(0).(func(*gorm.DB, *interactor.HostConfOptions) (*model.Domain, error)); ok {
		return rf(db, options)
	}
	if rf, ok := ret.Get(0).(func(*gorm.DB, *interactor.HostConfOptions) *model.Domain); ok {
		r0 = rf(db, options)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Domain)
		}
	}

	if rf, ok := ret.Get(1).(func(*gorm.DB, *interactor.HostConfOptions) error); ok {
		r1 = rf(db, options)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SignHostConfToken provides a mock function with given fields: privs, options, domain
func (_m *HostRepository) SignHostConfToken(privs []jwk.Key, options *interactor.HostConfOptions, domain *model.Domain) (string, error) {
	ret := _m.Called(privs, options, domain)

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func([]jwk.Key, *interactor.HostConfOptions, *model.Domain) (string, error)); ok {
		return rf(privs, options, domain)
	}
	if rf, ok := ret.Get(0).(func([]jwk.Key, *interactor.HostConfOptions, *model.Domain) string); ok {
		r0 = rf(privs, options, domain)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func([]jwk.Key, *interactor.HostConfOptions, *model.Domain) error); ok {
		r1 = rf(privs, options, domain)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewHostRepository creates a new instance of HostRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewHostRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *HostRepository {
	mock := &HostRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
