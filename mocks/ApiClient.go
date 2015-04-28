package mocks

import "github.com/Juniper/contrail-go-api"
import "github.com/stretchr/testify/mock"

type ApiClient struct {
	mock.Mock
}

func (m *ApiClient) Create(ptr contrail.IObject) error {
	ret := m.Called(ptr)

	r0 := ret.Error(0)

	return r0
}
func (m *ApiClient) Update(ptr contrail.IObject) error {
	ret := m.Called(ptr)

	r0 := ret.Error(0)

	return r0
}
func (m *ApiClient) DeleteByUuid(typename string, uuid string) error {
	ret := m.Called(typename, uuid)

	r0 := ret.Error(0)

	return r0
}
func (m *ApiClient) Delete(ptr contrail.IObject) error {
	ret := m.Called(ptr)

	r0 := ret.Error(0)

	return r0
}
func (m *ApiClient) FindByUuid(typename string, uuid string) (contrail.IObject, error) {
	ret := m.Called(typename, uuid)

	r0 := ret.Get(0).(contrail.IObject)
	r1 := ret.Error(1)

	return r0, r1
}
func (m *ApiClient) UuidByName(typename string, fqn string) (string, error) {
	ret := m.Called(typename, fqn)

	r0 := ret.Get(0).(string)
	r1 := ret.Error(1)

	return r0, r1
}
func (m *ApiClient) FQNameByUuid(uuid string) ([]string, error) {
	ret := m.Called(uuid)

	var r0 []string
	if ret.Get(0) != nil {
		r0 = ret.Get(0).([]string)
	}
	r1 := ret.Error(1)

	return r0, r1
}
func (m *ApiClient) FindByName(typename string, fqn string) (contrail.IObject, error) {
	ret := m.Called(typename, fqn)

	r0 := ret.Get(0).(contrail.IObject)
	r1 := ret.Error(1)

	return r0, r1
}
func (m *ApiClient) List(typename string, count int) ([]contrail.ListResult, error) {
	ret := m.Called(typename, count)

	var r0 []contrail.ListResult
	if ret.Get(0) != nil {
		r0 = ret.Get(0).([]contrail.ListResult)
	}
	r1 := ret.Error(1)

	return r0, r1
}
func (m *ApiClient) ListByParent(typename string, parent_id string, count int) ([]contrail.ListResult, error) {
	ret := m.Called(typename, parent_id, count)

	var r0 []contrail.ListResult
	if ret.Get(0) != nil {
		r0 = ret.Get(0).([]contrail.ListResult)
	}
	r1 := ret.Error(1)

	return r0, r1
}
func (m *ApiClient) ListDetail(typename string, fields []string, count int) ([]contrail.IObject, error) {
	ret := m.Called(typename, fields, count)

	var r0 []contrail.IObject
	if ret.Get(0) != nil {
		r0 = ret.Get(0).([]contrail.IObject)
	}
	r1 := ret.Error(1)

	return r0, r1
}
func (m *ApiClient) ListDetailByParent(typename string, parent_id string, fields []string, count int) ([]contrail.IObject, error) {
	ret := m.Called(typename, parent_id, fields, count)

	var r0 []contrail.IObject
	if ret.Get(0) != nil {
		r0 = ret.Get(0).([]contrail.IObject)
	}
	r1 := ret.Error(1)

	return r0, r1
}
