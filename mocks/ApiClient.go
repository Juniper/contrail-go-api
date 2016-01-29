package mocks

import (
	"fmt"
	"strings"

	"github.com/Juniper/contrail-go-api"
	"github.com/Juniper/contrail-go-api/types"
	"github.com/pborman/uuid"
)

type TypeInterceptor interface {
	Get(contrail.IObject)
	Put(contrail.IObject)
}

type ApiClient struct {
	IdAssignMap    map[string]string
	InterceptorMap map[string]TypeInterceptor
	db             Database
	updater        *objectUpdater
}

func (m *ApiClient) Init() {
	m.db = NewInMemDatabase()

	m.updater = NewObjectUpdater(m.db)

	domain := new(types.Domain)
	domain.SetName("default-domain")
	err := m.Create(domain)
	if err != nil {
		panic(err.Error())
	}

	project := new(types.Project)
	project.SetFQName("domain", []string{"default-domain", "default-project"})
	err = m.Create(project)
	if err != nil {
		panic(err.Error())
	}

	ipam := new(types.NetworkIpam)
	ipam.SetFQName("project", []string{"default-domain", "default-project", "default-network-ipam"})
	err = m.Create(ipam)
	if err != nil {
		panic(err.Error())
	}
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

func (m *ApiClient) interceptPut(ptr contrail.IObject) {
	if m.InterceptorMap == nil {
		return
	}
	if interceptor, ok := m.InterceptorMap[ptr.GetType()]; ok {
		interceptor.Put(ptr)
	}
}

func (m *ApiClient) interceptGet(ptr contrail.IObject) {
	if m.InterceptorMap == nil {
		return
	}
	if interceptor, ok := m.InterceptorMap[ptr.GetType()]; ok {
		interceptor.Get(ptr)
	}
}

func (m *ApiClient) getParent(obj contrail.IObject) (contrail.IObject, error) {
	typename := obj.GetParentType()
	if len(typename) == 0 || typename == "config-root" {
		return nil, nil
	}
	fqn := obj.GetFQName()
	parent_name := fqn[:len(fqn)-1]
	return m.db.GetByName(typename, strings.Join(parent_name, ":"))
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

	parent, err := m.getParent(ptr)
	if err != nil {
		return err
	}
	m.interceptPut(ptr)

	refList := GetReferenceList(ptr)

	m.db.Put(ptr, parent, refList)
	ptr.SetClient(m.updater)

	return nil
}

func (m *ApiClient) Update(ptr contrail.IObject) error {
	refList := GetReferenceList(ptr)
	m.db.Update(ptr, refList)
	return nil
}

func (m *ApiClient) DeleteByUuid(typename string, id string) error {
	obj, err := m.db.GetByUuid(uuid.Parse(id))
	if err != nil {
		return err
	}
	// Ensure the object has no children and/or back_refs
	return m.db.Delete(obj)
}

func (m *ApiClient) Delete(ptr contrail.IObject) error {
	return m.db.Delete(ptr)
}

func (m *ApiClient) FindByUuid(typename string, id string) (contrail.IObject, error) {
	obj, err := m.db.GetByUuid(uuid.Parse(id))
	if err != nil {
		return nil, err
	}
	m.interceptGet(obj)
	return obj, nil
}
func (m *ApiClient) UuidByName(typename string, fqn string) (string, error) {
	obj, err := m.db.GetByName(typename, fqn)
	if err != nil {
		return "", err
	}
	return obj.GetUuid(), nil
}
func (m *ApiClient) FQNameByUuid(id string) ([]string, error) {
	obj, err := m.db.GetByUuid(uuid.Parse(id))
	if err != nil {
		return []string{}, err
	}
	return obj.GetFQName(), nil
}
func (m *ApiClient) FindByName(typename string, fqn string) (contrail.IObject, error) {
	obj, err := m.db.GetByName(typename, fqn)
	if err != nil {
		return nil, err
	}
	m.interceptGet(obj)
	return obj, nil
}

func (m *ApiClient) List(typename string) ([]contrail.ListResult, error) {
	return m.ListByParent(typename, "")
}

func filterByParent(obj contrail.IObject, parent_id string) bool {
	if parent_id == "" {
		return false
	}
	fqn := obj.GetFQName()
	parentName := strings.Join(fqn[0:len(fqn)-2], ":")
	return parentName != parent_id
}

func (m *ApiClient) ListByParent(typename string, parent_id string) ([]contrail.ListResult, error) {
	nilList := []contrail.ListResult{}
	if _, ok := types.TypeMap[typename]; !ok {
		return nilList, fmt.Errorf("404 Not Found")
	}
	objList := m.db.List(typename)
	cap := 0
	if parent_id == "" {
		cap = len(objList)
	}
	result := make([]contrail.ListResult, 0, cap)
	for _, obj := range objList {
		if filterByParent(obj, parent_id) {
			continue
		}
		element := contrail.ListResult{
			Fq_name: obj.GetFQName(),
			Href:    obj.GetHref(),
			Uuid:    obj.GetUuid(),
		}
		result = append(result, element)
	}
	return result, nil
}

func (m *ApiClient) ListDetail(typename string, fields []string) ([]contrail.IObject, error) {
	nilList := []contrail.IObject{}
	if _, ok := types.TypeMap[typename]; !ok {
		return nilList, fmt.Errorf("404 Not Found")
	}
	objList := m.db.List(typename)
	return objList, nil
}

func (m *ApiClient) ListDetailByParent(typename string, parent_id string, fields []string) ([]contrail.IObject, error) {
	elements := make([]contrail.IObject, 0)
	if _, ok := types.TypeMap[typename]; !ok {
		return elements, fmt.Errorf("404 Not Found")
	}

	objList := m.db.List(typename)
	for _, obj := range objList {
		if filterByParent(obj, parent_id) {
			continue
		}
		elements = append(elements, obj)
	}
	return elements, nil
}
