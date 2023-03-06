package presenter

import (
	"fmt"
	"testing"

	"github.com/google/uuid"
	"github.com/hmsidm/internal/api/public"
	"github.com/hmsidm/internal/domain/model"
	"github.com/hmsidm/internal/interface/presenter"
	"github.com/openlyinc/pointy"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gorm.io/gorm"
)

func TestNewTodoPresenter(t *testing.T) {
	assert.NotPanics(t, func() {
		NewDomainPresenter()
	})
}

type mynewerror struct{}

func (e *mynewerror) Error() string {
	return "mynewerror"
}

func TestGet(t *testing.T) {
	testUuid := uuid.New()
	type TestCaseGiven struct {
		Input  *model.Domain
		Output *public.ReadDomainResponse
	}
	type TestCaseExpected struct {
		Err    error
		Output *public.ReadDomainResponse
	}
	type TestCase struct {
		Name     string
		Given    TestCaseGiven
		Expected TestCaseExpected
	}
	testCases := []TestCase{
		{
			Name: "error when 'in' is nil",
			Given: TestCaseGiven{
				Input: nil,
			},
			Expected: TestCaseExpected{
				Err:    fmt.Errorf("'domain' cannot be nil"),
				Output: nil,
			},
		},
		{
			Name: "Success case",
			Given: TestCaseGiven{
				Input: &model.Domain{
					Model:                 gorm.Model{ID: 1},
					OrgId:                 "12345",
					DomainUuid:            testUuid,
					DomainName:            pointy.String("domain.example"),
					DomainType:            pointy.Uint(model.DomainTypeIpa),
					AutoEnrollmentEnabled: pointy.Bool(true),
					IpaDomain: &model.Ipa{
						RealmName:  pointy.String("DOMAIN.EXAMPLE"),
						CaList:     pointy.String(""),
						ServerList: pointy.String("server1.domain.example,server2.domain.example"),
					},
				},
			},
			Expected: TestCaseExpected{
				Err: nil,
				Output: &public.ReadDomainResponse{
					AutoEnrollmentEnabled: pointy.Bool(true),
					DomainUuid:            pointy.String(testUuid.String()),
					DomainName:            pointy.String("domain.example"),
					DomainType:            pointy.String(model.DomainTypeString(model.DomainTypeIpa)),
					Ipa: &public.ReadDomainIpa{
						RealmName:  pointy.String("DOMAIN.EXAMPLE"),
						CaList:     "",
						ServerList: &[]string{"server1.domain.example", "server2.domain.example"},
					},
				},
			},
		},
	}
	for _, testCase := range testCases {
		t.Log(testCase.Name)
		obj := NewDomainPresenter()
		output, err := obj.Get(testCase.Given.Input)
		if testCase.Expected.Err != nil {
			require.Error(t, err)
			assert.Equal(t, testCase.Expected.Err.Error(), err.Error())
			assert.Nil(t, output)
		} else {
			assert.NoError(t, err)
			assert.Equal(t,
				*testCase.Expected.Output.DomainUuid,
				*output.DomainUuid)
			assert.Equal(t,
				*testCase.Expected.Output.DomainName,
				*output.DomainName)
			assert.Equal(t,
				*testCase.Expected.Output.DomainType,
				*output.DomainType)
			assert.Equal(t,
				*testCase.Expected.Output.AutoEnrollmentEnabled,
				*output.AutoEnrollmentEnabled)
			assert.Equal(t,
				*testCase.Expected.Output.Ipa.RealmName,
				*output.Ipa.RealmName)
			assert.Equal(t,
				testCase.Expected.Output.Ipa.CaList,
				output.Ipa.CaList)
			assert.Equal(t,
				testCase.Expected.Output.Ipa.ServerList,
				output.Ipa.ServerList)
		}
	}
}

func TestCreate(t *testing.T) {
	type TestCaseExpected struct {
		Response *public.CreateDomainResponse
		Err      error
	}
	type TestCase struct {
		Name     string
		Given    *model.Domain
		Expected TestCaseExpected
	}
	var (
		obj presenter.DomainPresenter
	)
	obj = NewDomainPresenter()

	testCases := []TestCase{
		{
			Name:  "domain is nil",
			Given: nil,
			Expected: TestCaseExpected{
				Response: nil,
				Err:      fmt.Errorf("domain cannot be nil"),
			},
		},
		{
			Name: "AutoenrollmentEnabled is nil",
			Given: &model.Domain{
				AutoEnrollmentEnabled: nil,
			},
			Expected: TestCaseExpected{
				Response: nil,
				Err:      fmt.Errorf("AutoenrollmentEnabled cannot be nil"),
			},
		},
		{
			Name: "DomainName is nil",
			Given: &model.Domain{
				AutoEnrollmentEnabled: pointy.Bool(true),
				DomainName:            nil,
			},
			Expected: TestCaseExpected{
				Response: nil,
				Err:      fmt.Errorf("DomainName cannot be nil"),
			},
		},
		{
			Name: "DomainType is nil",
			Given: &model.Domain{
				AutoEnrollmentEnabled: pointy.Bool(true),
				DomainName:            pointy.String("domain.example"),
				DomainType:            nil,
			},
			Expected: TestCaseExpected{
				Response: nil,
				Err:      fmt.Errorf("DomainType cannot be nil"),
			},
		},
		{
			Name: "DomainType is nil",
			Given: &model.Domain{
				AutoEnrollmentEnabled: pointy.Bool(true),
				DomainName:            pointy.String("domain.example"),
				DomainType:            nil,
			},
			Expected: TestCaseExpected{
				Response: nil,
				Err:      fmt.Errorf("DomainType cannot be nil"),
			},
		},
		{
			Name: "IpaDomain is nil",
			Given: &model.Domain{
				AutoEnrollmentEnabled: pointy.Bool(true),
				DomainName:            pointy.String("domain.example"),
				DomainType:            pointy.Uint(model.DomainTypeIpa),
				IpaDomain:             nil,
			},
			Expected: TestCaseExpected{
				Response: nil,
				Err:      fmt.Errorf("IpaDomain cannot be nil"),
			},
		},
		{
			Name: "CaList is nil",
			Given: &model.Domain{
				AutoEnrollmentEnabled: pointy.Bool(true),
				DomainName:            pointy.String("domain.example"),
				DomainType:            pointy.Uint(model.DomainTypeIpa),
				IpaDomain: &model.Ipa{
					CaList: nil,
				},
			},
			Expected: TestCaseExpected{
				Response: nil,
				Err:      fmt.Errorf("CaList cannot be nil"),
			},
		},
		{
			Name: "RealmName is nil",
			Given: &model.Domain{
				AutoEnrollmentEnabled: pointy.Bool(true),
				DomainName:            pointy.String("domain.example"),
				DomainType:            pointy.Uint(model.DomainTypeIpa),
				IpaDomain: &model.Ipa{
					CaList:    pointy.String(""),
					RealmName: nil,
				},
			},
			Expected: TestCaseExpected{
				Response: nil,
				Err:      fmt.Errorf("RealmName cannot be nil"),
			},
		},
		{
			Name: "ServerList is nil",
			Given: &model.Domain{
				AutoEnrollmentEnabled: pointy.Bool(true),
				DomainName:            pointy.String("domain.example"),
				DomainType:            pointy.Uint(model.DomainTypeIpa),
				IpaDomain: &model.Ipa{
					CaList:     pointy.String(""),
					RealmName:  pointy.String("DOMAIN.EXAMPLE"),
					ServerList: nil,
				},
			},
			Expected: TestCaseExpected{
				Response: nil,
				Err:      fmt.Errorf("ServerList cannot be nil"),
			},
		},
		{
			Name: "Success scenario",
			Given: &model.Domain{
				AutoEnrollmentEnabled: pointy.Bool(true),
				DomainName:            pointy.String("domain.example"),
				DomainType:            pointy.Uint(model.DomainTypeIpa),
				IpaDomain: &model.Ipa{
					CaList:     pointy.String(""),
					RealmName:  pointy.String("DOMAIN.EXAMPLE"),
					ServerList: pointy.String("server1.domain.example,server2.domain.example"),
				},
			},
			Expected: TestCaseExpected{
				Response: &public.CreateDomainResponseSchema{
					DomainName:            "domain.example",
					AutoEnrollmentEnabled: true,
					DomainType:            "ipa",
					DomainUuid:            "00000000-0000-0000-0000-000000000000",
					Ipa: public.CreateDomainIpa{
						CaList:     "",
						ServerList: &[]string{"server1.domain.example", "server2.domain.example"},
						RealmName:  pointy.String("DOMAIN.EXAMPLE"),
					},
				},
				Err: nil,
			},
		},
	}

	for _, testCase := range testCases {
		response, err := obj.Create(testCase.Given)
		if testCase.Expected.Err != nil {
			require.Error(t, err)
			assert.EqualError(t, err, testCase.Expected.Err.Error())
			assert.Nil(t, response)
		} else {
			assert.NoError(t, err)
			assert.Equal(t, testCase.Expected.Response.DomainName, response.DomainName)
			assert.Equal(t, testCase.Expected.Response.AutoEnrollmentEnabled, response.AutoEnrollmentEnabled)
			assert.Equal(t, testCase.Expected.Response.DomainType, response.DomainType)
			assert.Equal(t, testCase.Expected.Response.DomainUuid, response.DomainUuid)
			assert.Equal(t, testCase.Expected.Response.Ipa.CaList, response.Ipa.CaList)
			assert.Equal(t, testCase.Expected.Response.Ipa.ServerList, response.Ipa.ServerList)
			assert.Equal(t, *testCase.Expected.Response.Ipa.RealmName, *response.Ipa.RealmName)
		}
	}

}
