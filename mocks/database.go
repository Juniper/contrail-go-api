package mocks

import (
	"fmt"
	"sort"
	"strings"

	"github.com/Juniper/contrail-go-api"
	"github.com/pborman/uuid"
)

// Database defines the interface used to mock the behavior of the contrail API server.
type Database interface {
	Put(obj contrail.IObject, parent contrail.IObject, refs UIDList) error
	Update(obj contrail.IObject, refs UIDList) error
	Delete(obj contrail.IObject) error
	GetByUUID(id uuid.UUID) (contrail.IObject, error)
	GetByName(typename string, fqn string) (contrail.IObject, error)
	List(typename string) []contrail.IObject
	GetChildren(uid UID, typename string) (UIDList, error)
	GetBackReferences(uid UID, typename string) (UIDList, error)
}

// TypeMap stores per typename map of name to object.
type TypeMap map[string]contrail.IObject

// ReferenceMap stores references by typename
type ReferenceMap map[string]UIDList

// ObjectData stores reference data
type ObjectData struct {
	parent UID

	// Forward references stored on a database update.
	refs UIDList

	// Children
	children ReferenceMap

	// Reverse references
	backRefs ReferenceMap
}

// InMemDatabase implements the Database interface.
type InMemDatabase struct {
	objByIDMap map[UID]contrail.IObject
	typeDB     map[string]TypeMap
	objectData map[UID]*ObjectData
}

// NewInMemDatabase creates and initializes a InMemDatabase.
func NewInMemDatabase() Database {
	db := new(InMemDatabase)
	db.objByIDMap = make(map[UID]contrail.IObject)
	db.typeDB = make(map[string]TypeMap)
	db.objectData = make(map[UID]*ObjectData)
	return db
}

func (db *InMemDatabase) addChild(parent, obj contrail.IObject) error {
	data, ok := db.objectData[parseUID(parent.GetUuid())]
	if !ok {
		return fmt.Errorf("Parent %s: not in database", parent.GetUuid())
	}
	typename := strings.Replace(obj.GetType(), "-", "_", -1)
	rList, ok := data.children[typename]
	if !ok {
		rList = make(UIDList, 0, 1)
	}
	uid := parseUID(obj.GetUuid())
	data.children[typename] = append(rList, uid)
	return nil
}

func (db *InMemDatabase) deleteChild(parent UID, obj contrail.IObject) error {
	data, ok := db.objectData[parent]
	if !ok {
		return fmt.Errorf("Parent %s: not in database", parent.Interface().String())
	}
	typename := strings.Replace(obj.GetType(), "-", "_", -1)
	rList, ok := data.children[typename]
	if !ok {
		return fmt.Errorf("Parent %s: no children", parent.Interface().String())
	}
	uid := parseUID(obj.GetUuid())
	var index = -1
	for i, r := range rList {
		if r == uid {
			index = i
			break
		}
	}
	if index < 0 {
		return fmt.Errorf("Parent %s: has no child %s", parent.Interface().String(), obj.GetUuid())
	}
	if len(rList) == 1 {
		delete(data.children, typename)
	} else {
		rList[index], rList = rList[len(rList)-1], rList[:len(rList)-1]
		data.children[typename] = rList
	}
	return nil
}

func (db *InMemDatabase) addBackReference(uid UID, typename string, ref UID) error {
	data, ok := db.objectData[uid]
	if !ok {
		return fmt.Errorf("Object %s: not in database", uid.Interface().String())
	}
	rList, ok := data.backRefs[typename]
	if !ok {
		rList = make(UIDList, 0, 1)
	}
	data.backRefs[typename] = append(rList, ref)

	if obj, ok := db.objByIDMap[uid]; ok {
		clearReferenceMask(obj)
	}

	return nil
}

func (db *InMemDatabase) deleteBackReference(uid UID, typename string, ref UID) error {
	data, ok := db.objectData[uid]
	if !ok {
		return fmt.Errorf("Object %s: not in database", uid.Interface().String())
	}
	rList, ok := data.backRefs[typename]
	if !ok {
		return fmt.Errorf("Object %s: no %s_back_refs", uid.Interface().String(), typename)
	}
	var index = -1
	for i, r := range rList {
		if r == ref {
			index = i
			break
		}
	}
	if index < 0 {
		return fmt.Errorf("Object %s: has no back_ref %s", uid.Interface().String(), ref.Interface().String())
	}
	if len(rList) == 1 {
		delete(data.backRefs, typename)
	} else {
		rList[index], rList = rList[len(rList)-1], rList[:len(rList)-1]
		data.backRefs[typename] = rList
	}

	if obj, ok := db.objByIDMap[uid]; ok {
		clearReferenceMask(obj)
	}

	return nil
}

func (db *InMemDatabase) addBackReferences(obj contrail.IObject, refs UIDList) error {
	typename := strings.Replace(obj.GetType(), "-", "_", -1)
	uid := parseUID(obj.GetUuid())
	for _, r := range refs {
		db.addBackReference(r, typename, uid)
	}
	return nil
}

type uidListSorter []UID

func (x uidListSorter) Len() int {
	return len(x)
}
func (x uidListSorter) Swap(i, j int) {
	y := x[i]
	x[i], x[j] = x[j], y
}
func (x uidListSorter) Less(i, j int) bool {
	return Compare(x[i], x[j]) < 0
}
func (db *InMemDatabase) updateBackReferences(obj contrail.IObject, newRefs, oldRefs UIDList) error {
	typename := strings.Replace(obj.GetType(), "-", "_", -1)
	uid := parseUID(obj.GetUuid())
	sort.Sort(uidListSorter(newRefs))
	sort.Sort(uidListSorter(oldRefs))
	var i, j int = 0, 0
	for i < len(oldRefs) && j < len(newRefs) {
		lhs := oldRefs[i]
		rhs := newRefs[j]
		cmp := Compare(lhs, rhs)
		if cmp < 0 {
			db.deleteBackReference(lhs, typename, uid)
			i++
		} else if cmp > 0 {
			db.addBackReference(rhs, typename, uid)
			j++
		} else {
			i++
			j++
		}
	}
	for ; i < len(oldRefs); i++ {
		db.deleteBackReference(oldRefs[i], typename, uid)
	}
	for ; j < len(newRefs); j++ {
		db.addBackReference(newRefs[j], typename, uid)
	}
	return nil
}

func (db *InMemDatabase) deleteBackReferences(obj contrail.IObject, oldRefs UIDList) error {
	typename := strings.Replace(obj.GetType(), "-", "_", -1)
	uid := parseUID(obj.GetUuid())

	for _, r := range oldRefs {
		db.deleteBackReference(r, typename, uid)
	}
	return nil
}

// Put adds an object to the database and updates the children list on the parent and the
// back_references lists of objects that obj refers to.
func (db *InMemDatabase) Put(obj contrail.IObject, parent contrail.IObject, refs UIDList) error {
	id := parseUID(obj.GetUuid())
	if _, ok := db.objByIDMap[id]; ok {
		return fmt.Errorf("Duplicate id %s", id.Interface().String())
	}

	typeMap, ok := db.typeDB[obj.GetType()]
	if !ok {
		typeMap = make(map[string]contrail.IObject)
		db.typeDB[obj.GetType()] = typeMap
	}
	fqn := strings.Join(obj.GetFQName(), ":")
	if _, ok := typeMap[fqn]; ok {
		return fmt.Errorf("Duplicate name %s", fqn)
	}
	db.objByIDMap[id] = obj
	typeMap[fqn] = obj

	var parentID UID
	if parent != nil {
		parentID = parseUID(parent.GetUuid())
	}
	db.objectData[id] = &ObjectData{
		parent:   parentID,
		refs:     refs,
		children: make(ReferenceMap),
		backRefs: make(ReferenceMap),
	}
	if parent != nil {
		db.addChild(parent, obj)
		clearReferenceMask(parent)
	}
	db.addBackReferences(obj, refs)
	return nil
}

// Update an object in the database, modifying corresponding references (forward and back).
func (db *InMemDatabase) Update(obj contrail.IObject, refs UIDList) error {
	uid := parseUID(obj.GetUuid())
	data, ok := db.objectData[uid]
	if !ok {
		return fmt.Errorf("Object %s: not in database", obj.GetUuid())
	}
	db.updateBackReferences(obj, refs, data.refs)
	data.refs = refs
	return nil
}

// Delete an object and any back references on objects that it refers to. Deletes as rejected
// for objects that have children or back references mimicing the behavior of
// the Contrail API server.
func (db *InMemDatabase) Delete(obj contrail.IObject) error {
	uid := parseUID(obj.GetUuid())
	data, ok := db.objectData[uid]
	if !ok {
		return fmt.Errorf("Object %s: not in database", obj.GetUuid())
	}
	if len(data.children) > 0 {
		return fmt.Errorf("Delete %s: has children %+v", obj.GetUuid(), data.children)
	}
	if len(data.backRefs) > 0 {
		return fmt.Errorf("Delete %s: has references %+v", obj.GetUuid(), data.backRefs)
	}

	if !data.parent.IsNIL() {
		db.deleteChild(data.parent, obj)
		if parentObj, ok := db.objByIDMap[data.parent]; ok {
			clearReferenceMask(parentObj)
		}
	}
	db.deleteBackReferences(obj, data.refs)

	delete(db.objByIDMap, uid)
	delete(db.objectData, uid)
	typeMap, ok := db.typeDB[obj.GetType()]
	if !ok {
		return fmt.Errorf("No objects of type %s", obj.GetType())
	}
	fqn := strings.Join(obj.GetFQName(), ":")
	delete(typeMap, fqn)
	return nil
}

// GetByUUID retrieves and object by UUID.
func (db *InMemDatabase) GetByUUID(id uuid.UUID) (contrail.IObject, error) {
	uid := makeUID(id)
	obj, ok := db.objByIDMap[uid]
	if !ok {
		return nil, fmt.Errorf("%s: Not found", id.String())
	}
	return obj, nil
}

// GetByName retrieves an object specified by colon separated fully qualified name.
func (db *InMemDatabase) GetByName(typename string, fqn string) (contrail.IObject, error) {
	typeMap, ok := db.typeDB[typename]
	if !ok {
		return nil, fmt.Errorf("%s %s: Not found", typename, fqn)
	}
	obj, ok := typeMap[fqn]
	if !ok {
		return nil, fmt.Errorf("%s %s: Not Found", typename, fqn)
	}
	return obj, nil
}

// GetChildren retrieves the list of children.
func (db *InMemDatabase) GetChildren(uid UID, typename string) (UIDList, error) {
	nilList := []UID{}
	data, ok := db.objectData[uid]
	if !ok {
		return nilList, fmt.Errorf("Object %s: not in database", uid.Interface().String())
	}
	refList, ok := data.children[typename]
	if !ok {
		return nilList, nil
	}
	return refList, nil
}

// GetBackReferences retrieves the list of back references.
func (db *InMemDatabase) GetBackReferences(uid UID, typename string) (UIDList, error) {
	nilList := []UID{}
	data, ok := db.objectData[uid]
	if !ok {
		return nilList, fmt.Errorf("Object %s: not in database", uid.Interface().String())
	}
	refList, ok := data.backRefs[typename]
	if !ok {
		return nilList, nil
	}
	return refList, nil
}

// List retrieves all objects of a given type.
func (db *InMemDatabase) List(typename string) []contrail.IObject {
	nilList := []contrail.IObject{}
	typeMap, ok := db.typeDB[typename]
	if !ok {
		return nilList
	}
	values := make([]contrail.IObject, len(typeMap))
	i := 0
	for _, v := range typeMap {
		values[i] = v
		i++
	}
	return values
}
