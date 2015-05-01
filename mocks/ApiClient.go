package mocks

import (
	"fmt"
	"strings"

	"code.google.com/p/go-uuid/uuid"
	"github.com/Juniper/contrail-go-api"
	"github.com/stretchr/testify/mock"
)

type TypeInterceptor interface {
	Get(contrail.IObject)
}

type ApiClient struct {
	mock.Mock
	IdAssignMap    map[string]string
	InterceptorMap map[string]TypeInterceptor
	objByNameMap   map[string]contrail.IObject
	objByIdMap     map[string]contrail.IObject
}

func (m *ApiClient) Init() {
	m.objByNameMap = make(map[string]contrail.IObject)
	m.objByIdMap = make(map[string]contrail.IObject)
}

func objName(ptr contrail.IObject) string {
	return ptr.GetType() + ":" + strings.Join(ptr.GetFQName(), ":")
}

func (m *ApiClient) AddInterceptor(typename string, interceptor TypeInterceptor) {
	if m.InterceptorMap == nil {
		m.InterceptorMap = make(map[string]TypeInterceptor)
	}
	m.InterceptorMap[typename] = interceptor
}
func (m *ApiClient) interceptGet(ptr contrail.IObject) {
	if m.InterceptorMap == nil {
		return
	}
	if interceptor, ok := m.InterceptorMap[ptr.GetType()]; ok {
		interceptor.Get(ptr)
	}
}

func (m *ApiClient) Create(ptr contrail.IObject) error {
	if ptr.GetUuid() == "" {
		isSet := false
		if m.IdAssignMap != nil {
			if id, ok := m.IdAssignMap[objName(ptr)]; ok {
				ptr.SetUuid(id)
				isSet = true
			}
		}
		if !isSet {
			ptr.SetUuid(uuid.New())
		}
	}
	m.objByNameMap[objName(ptr)] = ptr
	m.objByIdMap[ptr.GetUuid()] = ptr
	return nil
}
func (m *ApiClient) Update(ptr contrail.IObject) error {
	return nil
}
func (m *ApiClient) DeleteByUuid(typename string, uuid string) error {
	obj, ok := m.objByIdMap[uuid]
	if !ok {
		return fmt.Errorf("%s %s: Not found", typename, uuid)
	}
	delete(m.objByIdMap, uuid)
	delete(m.objByNameMap, objName(obj))
	return nil
}
func (m *ApiClient) Delete(ptr contrail.IObject) error {
	delete(m.objByIdMap, ptr.GetUuid())
	delete(m.objByNameMap, objName(ptr))
	return nil
}
func (m *ApiClient) FindByUuid(typename string, uuid string) (contrail.IObject, error) {
	obj, ok := m.objByIdMap[uuid]
	if !ok {
		return nil, fmt.Errorf("%s %s: Not found", typename, uuid)
	}
	m.interceptGet(obj)
	return obj, nil
}
func (m *ApiClient) UuidByName(typename string, fqn string) (string, error) {
	obj, ok := m.objByNameMap[typename+":"+fqn]
	if !ok {
		return "", fmt.Errorf("%s %s: Not found", typename, fqn)
	}
	return obj.GetUuid(), nil
}
func (m *ApiClient) FQNameByUuid(uuid string) ([]string, error) {
	obj, ok := m.objByIdMap[uuid]
	if !ok {
		return []string{}, fmt.Errorf("%s: Not found", uuid)
	}
	return obj.GetFQName(), nil
}
func (m *ApiClient) FindByName(typename string, fqn string) (contrail.IObject, error) {
	obj, ok := m.objByNameMap[typename+":"+fqn]
	if !ok {
		return nil, fmt.Errorf("%s %s: Not found", typename, fqn)
	}
	m.interceptGet(obj)
	return obj, nil
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
