package mocks

import (
	"fmt"
	"strings"

	"github.com/Juniper/contrail-go-api"
	"github.com/Juniper/contrail-go-api/types"
	"github.com/pborman/uuid"
)

// TypeInterceptor defines an interface that intercepts Put and Get requests to the API.
// Typically used to simulate API server behavior that is type specific. An example is
// the API server determining the Gateway address on a Subnet.
type TypeInterceptor interface {
	Get(contrail.IObject)
	Put(contrail.IObject)
}

// ApiClient mocks the contrail.ApiClient interface.
type ApiClient struct {
	IDAssignMap    map[string]string
	InterceptorMap map[string]TypeInterceptor
	db             Database
	updater        *objectUpdater
}

// Init initializes the database with default values such as the default-project and ipam.
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

// AddInterceptor is used to apply a type-specific hook for GET and PUT requests.
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
	parentName := fqn[:len(fqn)-1]
	return m.db.GetByName(typename, strings.Join(parentName, ":"))
}

// Create adds an object to the database.
func (m *ApiClient) Create(ptr contrail.IObject) error {
	if ptr.GetUuid() == "" {
		isSet := false
		if m.IDAssignMap != nil {
			if id, ok := m.IDAssignMap[objName(ptr)]; ok {
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

	refList := getReferenceList(ptr)

	m.db.Put(ptr, parent, refList)
	ptr.SetClient(m.updater)

	return nil
}

// Update modifies the object in the database.
func (m *ApiClient) Update(ptr contrail.IObject) error {
	refList := getReferenceList(ptr)
	m.db.Update(ptr, refList)
	return nil
}

// DeleteByUuid removes the object identified by the specified uuid.
func (m *ApiClient) DeleteByUuid(typename string, id string) error {
	obj, err := m.db.GetByUUID(uuid.Parse(id))
	if err != nil {
		return err
	}
	// Ensure the object has no children and/or back_refs
	return m.db.Delete(obj)
}

// Delete an object.
func (m *ApiClient) Delete(ptr contrail.IObject) error {
	return m.db.Delete(ptr)
}

// FindByUuid retrieves the object specified by uuid.
func (m *ApiClient) FindByUuid(typename string, id string) (contrail.IObject, error) {
	obj, err := m.db.GetByUUID(uuid.Parse(id))
	if err != nil {
		return nil, err
	}
	m.interceptGet(obj)
	return obj, nil
}

// UuidByName retrieves the uuid given a fully qualified name as a colon (:) delimited string.
func (m *ApiClient) UuidByName(typename string, fqn string) (string, error) {
	obj, err := m.db.GetByName(typename, fqn)
	if err != nil {
		return "", err
	}
	return obj.GetUuid(), nil
}

// FQNameByUuid retrieves the fully qualified name corresponding to a UUID.
func (m *ApiClient) FQNameByUuid(id string) ([]string, error) {
	obj, err := m.db.GetByUUID(uuid.Parse(id))
	if err != nil {
		return []string{}, err
	}
	return obj.GetFQName(), nil
}

// FindByName reads the database object identified by a colon (:) delimited string.
func (m *ApiClient) FindByName(typename string, fqn string) (contrail.IObject, error) {
	obj, err := m.db.GetByName(typename, fqn)
	if err != nil {
		return nil, err
	}
	m.interceptGet(obj)
	return obj, nil
}

// List retrives the identifiers of all objects of a given type.
func (m *ApiClient) List(typename string) ([]contrail.ListResult, error) {
	return m.listByParentImpl(typename, nil)
}

// Return true if the object should be excluded from the results.
func (m *ApiClient) filterByParent(obj contrail.IObject, parentID uuid.UUID) bool {
	if parentID == nil {
		return false
	}
	parent, err := m.getParent(obj)
	if err != nil {
		return true
	}

	return !uuid.Equal(uuid.Parse(parent.GetUuid()), parentID)
}

// ListByParent retrieves the identifiers of objects of the specified type that are descendents
// of parent (as identified by UUID)
func (m *ApiClient) ListByParent(typename string, parentID string) ([]contrail.ListResult, error) {
	parentUUID := uuid.Parse(parentID)
	if parentUUID == nil {
		return nil, fmt.Errorf("Invalid uuid: %s", parentID)
	}
	return m.listByParentImpl(typename, parentUUID)
}

func (m *ApiClient) listByParentImpl(typename string, parentID uuid.UUID) ([]contrail.ListResult, error) {
	if _, ok := types.TypeMap[typename]; !ok {
		return nil, fmt.Errorf("404 Not Found")
	}
	objList := m.db.List(typename)
	cap := 0
	if parentID == nil {
		cap = len(objList)
	}
	result := make([]contrail.ListResult, 0, cap)
	for _, obj := range objList {
		if m.filterByParent(obj, parentID) {
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

// ListDetail reads all the objects of a given type.
func (m *ApiClient) ListDetail(typename string, fields []string) ([]contrail.IObject, error) {
	nilList := []contrail.IObject{}
	if _, ok := types.TypeMap[typename]; !ok {
		return nilList, fmt.Errorf("404 Not Found")
	}
	objList := m.db.List(typename)
	return objList, nil
}

// ListDetailByParent reads all the objects of the specified type that are descendants of the parent
// object (as specified by UUID).
func (m *ApiClient) ListDetailByParent(typename string, parentID string, fields []string) ([]contrail.IObject, error) {
	elements := make([]contrail.IObject, 0)
	if _, ok := types.TypeMap[typename]; !ok {
		return nil, fmt.Errorf("404 Not Found")
	}
	parentUUID := uuid.Parse(parentID)
	if parentUUID == nil {
		return nil, fmt.Errorf("Invalid uuid: %s", parentID)
	}

	objList := m.db.List(typename)
	for _, obj := range objList {
		if m.filterByParent(obj, parentUUID) {
			continue
		}
		elements = append(elements, obj)
	}
	return elements, nil
}
