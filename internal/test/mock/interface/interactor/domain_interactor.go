// Code generated by mockery v2.16.0. DO NOT EDIT.

package interactor

import (
	header "github.com/podengo-project/idmsvc-backend/internal/api/header"
	identity "github.com/redhatinsights/platform-go-middlewares/identity"

	mock "github.com/stretchr/testify/mock"

	model "github.com/podengo-project/idmsvc-backend/internal/domain/model"

	public "github.com/podengo-project/idmsvc-backend/internal/api/public"
)

// DomainInteractor is an autogenerated mock type for the DomainInteractor type
type DomainInteractor struct {
	mock.Mock
}

// Create provides a mock function with given fields: xrhid, params, body
func (_m *DomainInteractor) Create(xrhid *identity.XRHID, params *public.CreateDomainParams, body *public.CreateDomain) (string, *model.Domain, error) {
	ret := _m.Called(xrhid, params, body)

	var r0 string
	if rf, ok := ret.Get(0).(func(*identity.XRHID, *public.CreateDomainParams, *public.CreateDomain) string); ok {
		r0 = rf(xrhid, params, body)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 *model.Domain
	if rf, ok := ret.Get(1).(func(*identity.XRHID, *public.CreateDomainParams, *public.CreateDomain) *model.Domain); ok {
		r1 = rf(xrhid, params, body)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*model.Domain)
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(*identity.XRHID, *public.CreateDomainParams, *public.CreateDomain) error); ok {
		r2 = rf(xrhid, params, body)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// Delete provides a mock function with given fields: xrhid, uuid, params
func (_m *DomainInteractor) Delete(xrhid *identity.XRHID, uuid string, params *public.DeleteDomainParams) (string, string, error) {
	ret := _m.Called(xrhid, uuid, params)

	var r0 string
	if rf, ok := ret.Get(0).(func(*identity.XRHID, string, *public.DeleteDomainParams) string); ok {
		r0 = rf(xrhid, uuid, params)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 string
	if rf, ok := ret.Get(1).(func(*identity.XRHID, string, *public.DeleteDomainParams) string); ok {
		r1 = rf(xrhid, uuid, params)
	} else {
		r1 = ret.Get(1).(string)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(*identity.XRHID, string, *public.DeleteDomainParams) error); ok {
		r2 = rf(xrhid, uuid, params)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// GetByID provides a mock function with given fields: xrhid, params
func (_m *DomainInteractor) GetByID(xrhid *identity.XRHID, params *public.ReadDomainParams) (string, error) {
	ret := _m.Called(xrhid, params)

	var r0 string
	if rf, ok := ret.Get(0).(func(*identity.XRHID, *public.ReadDomainParams) string); ok {
		r0 = rf(xrhid, params)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*identity.XRHID, *public.ReadDomainParams) error); ok {
		r1 = rf(xrhid, params)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// List provides a mock function with given fields: xrhid, params
func (_m *DomainInteractor) List(xrhid *identity.XRHID, params *public.ListDomainsParams) (string, int, int, error) {
	ret := _m.Called(xrhid, params)

	var r0 string
	if rf, ok := ret.Get(0).(func(*identity.XRHID, *public.ListDomainsParams) string); ok {
		r0 = rf(xrhid, params)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 int
	if rf, ok := ret.Get(1).(func(*identity.XRHID, *public.ListDomainsParams) int); ok {
		r1 = rf(xrhid, params)
	} else {
		r1 = ret.Get(1).(int)
	}

	var r2 int
	if rf, ok := ret.Get(2).(func(*identity.XRHID, *public.ListDomainsParams) int); ok {
		r2 = rf(xrhid, params)
	} else {
		r2 = ret.Get(2).(int)
	}

	var r3 error
	if rf, ok := ret.Get(3).(func(*identity.XRHID, *public.ListDomainsParams) error); ok {
		r3 = rf(xrhid, params)
	} else {
		r3 = ret.Error(3)
	}

	return r0, r1, r2, r3
}

// Register provides a mock function with given fields: xrhid, UUID, params, body
func (_m *DomainInteractor) Register(xrhid *identity.XRHID, UUID string, params *public.RegisterDomainParams, body *public.Domain) (string, *header.XRHIDMVersion, *model.Domain, error) {
	ret := _m.Called(xrhid, UUID, params, body)

	var r0 string
	if rf, ok := ret.Get(0).(func(*identity.XRHID, string, *public.RegisterDomainParams, *public.Domain) string); ok {
		r0 = rf(xrhid, UUID, params, body)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 *header.XRHIDMVersion
	if rf, ok := ret.Get(1).(func(*identity.XRHID, string, *public.RegisterDomainParams, *public.Domain) *header.XRHIDMVersion); ok {
		r1 = rf(xrhid, UUID, params, body)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*header.XRHIDMVersion)
		}
	}

	var r2 *model.Domain
	if rf, ok := ret.Get(2).(func(*identity.XRHID, string, *public.RegisterDomainParams, *public.Domain) *model.Domain); ok {
		r2 = rf(xrhid, UUID, params, body)
	} else {
		if ret.Get(2) != nil {
			r2 = ret.Get(2).(*model.Domain)
		}
	}

	var r3 error
	if rf, ok := ret.Get(3).(func(*identity.XRHID, string, *public.RegisterDomainParams, *public.Domain) error); ok {
		r3 = rf(xrhid, UUID, params, body)
	} else {
		r3 = ret.Error(3)
	}

	return r0, r1, r2, r3
}

// Update provides a mock function with given fields: xrhid, UUID, params, body
func (_m *DomainInteractor) Update(xrhid *identity.XRHID, UUID string, params *public.UpdateDomainParams, body *public.Domain) (string, *header.XRHIDMVersion, *model.Domain, error) {
	ret := _m.Called(xrhid, UUID, params, body)

	var r0 string
	if rf, ok := ret.Get(0).(func(*identity.XRHID, string, *public.UpdateDomainParams, *public.Domain) string); ok {
		r0 = rf(xrhid, UUID, params, body)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 *header.XRHIDMVersion
	if rf, ok := ret.Get(1).(func(*identity.XRHID, string, *public.UpdateDomainParams, *public.Domain) *header.XRHIDMVersion); ok {
		r1 = rf(xrhid, UUID, params, body)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*header.XRHIDMVersion)
		}
	}

	var r2 *model.Domain
	if rf, ok := ret.Get(2).(func(*identity.XRHID, string, *public.UpdateDomainParams, *public.Domain) *model.Domain); ok {
		r2 = rf(xrhid, UUID, params, body)
	} else {
		if ret.Get(2) != nil {
			r2 = ret.Get(2).(*model.Domain)
		}
	}

	var r3 error
	if rf, ok := ret.Get(3).(func(*identity.XRHID, string, *public.UpdateDomainParams, *public.Domain) error); ok {
		r3 = rf(xrhid, UUID, params, body)
	} else {
		r3 = ret.Error(3)
	}

	return r0, r1, r2, r3
}

type mockConstructorTestingTNewDomainInteractor interface {
	mock.TestingT
	Cleanup(func())
}

// NewDomainInteractor creates a new instance of DomainInteractor. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewDomainInteractor(t mockConstructorTestingTNewDomainInteractor) *DomainInteractor {
	mock := &DomainInteractor{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
