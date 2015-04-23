//
// Copyright (c) 2014 Juniper Networks, Inc. All rights reserved.
//

package contrail

import (
	"encoding/json"
	"fmt"
	"os"
	"reflect"
	"sort"
	"strings"
)

// Interface implemented by auto-generated types.
type IObject interface {
	GetDefaultParent() []string
	GetDefaultParentType() string
	GetFQName() []string
	GetName() string
	GetType() string
	GetUuid() string
	GetHref() string
	SetName(string)
	SetUuid(string)
	SetFQName(string, []string)

	SetClient(ClientInterface)
	UpdateObject() ([]byte, error)
	UpdateReferences() error
	UpdateDone()
}

// ObjectBase class
type ObjectBase struct {
	fq_name []string
	href string
	name string
	uuid string
	parent_href string
	parent_type string
	parent_uuid string

	// clientPtr is set once the object is persisted in the API server
	// or for objects that are retrieved via Read/GET
	clientPtr ClientInterface
	parent IObject
}

// Used by IObject.SetName methods.
//
// The implementation must be able to access both the ObjectBase fields
// as well as the IObject interface in order to retrieve data specific to
// a given type.
func (obj *ObjectBase) VSetName(vPtr IObject, name string) {
	obj.name = name
	if obj.parent != nil {
		obj.fq_name = make([]string, len(obj.parent.GetFQName()) + 1)
		copy(obj.fq_name, obj.parent.GetFQName())
		obj.fq_name = append(obj.fq_name, name)
		obj.parent_type = obj.parent.GetType()
	} else {
		obj.fq_name = make([]string, len(vPtr.GetDefaultParent()) + 1)
		copy(obj.fq_name, vPtr.GetDefaultParent())
		obj.fq_name = append(obj.fq_name, name)
		obj.parent_type = vPtr.GetDefaultParentType()
	}
}

func (obj *ObjectBase) VSetParent(vPtr IObject, parent IObject) {
	obj.parent = parent
	if len(obj.name) > 0 {
		obj.VSetName(vPtr, obj.name)
	}
}

// Accessor for object (unqualified) name
func (obj *ObjectBase) GetName() string {
	return obj.name
}

// Accessor for object uuid.
func (obj *ObjectBase) GetUuid() string {
	return obj.uuid
}

// Setter for uuid.
func (obj *ObjectBase) SetUuid(uuid string) {
	if obj.clientPtr != nil {
		panic(fmt.Sprintf("Attempt to override uuid for %s", obj.uuid))
	}
	obj.uuid = uuid
}


// Accessor for href.
func (obj *ObjectBase) GetHref() string {
	return obj.href
}

// Set the fully qualified domain name.
func (obj *ObjectBase) SetFQName(parentType string, fqn []string) {
	obj.fq_name = fqn
	obj.name = fqn[len(fqn) - 1]
	obj.parent_type = parentType
}

func (obj *ObjectBase) GetFQName() []string {
	return obj.fq_name
}

func (obj *ObjectBase) SetClient(c ClientInterface) {
	obj.clientPtr = c
}

func (obj *ObjectBase) IsTransient() bool {
	return obj.clientPtr == nil
}

func (obj *ObjectBase) GetField(ptr IObject, field string) error {
	return obj.clientPtr.GetField(ptr, field)
}

func (obj *ObjectBase) UnmarshalCommon(m map[string]json.RawMessage) error {
	var err error
	err = json.Unmarshal(m["fq_name"], &obj.fq_name)
	if err != nil {
		return err
	}
	err = json.Unmarshal(m["uuid"], &obj.uuid)
	if err != nil {
		return err
	}
	err = json.Unmarshal(m["name"], &obj.name)
	if err != nil {
		return err
	}
	err = json.Unmarshal(m["href"], &obj.href)
	if err != nil {
		return err
	}
	// Older versions of the API server have a bug generating the href
	// on list commands
	helements := strings.Split(obj.href, "/")
	if (helements[len(helements)-1] != obj.uuid) {
		fmt.Fprintf(os.Stderr, "WARN invalid href: %s\n", obj.href)
		helements[len(helements)-1] = obj.uuid
		obj.href = strings.Join(helements, "/")
	}
	return nil
}

// Encode fq_name and uuid.
func (obj *ObjectBase) MarshalId(m map[string]*json.RawMessage) error {
	{
		var value json.RawMessage
		value, err := json.Marshal(obj.fq_name)
		if err != nil {
			return err
		}
		m["fq_name"] = &value
	}
	if len(obj.uuid) > 0 {
		var value json.RawMessage
		value, err := json.Marshal(&obj.uuid)
		if err != nil {
			return err
		}
		m["uuid"] = &value
	}
	return nil
}

// Encode the information stored in the ObjectBase struct.
func (obj *ObjectBase) MarshalCommon(m map[string]*json.RawMessage) error {
	err := obj.MarshalId(m)
	if err != nil {
		return err
	}
 	if len(obj.parent_type) > 0 {
		var value json.RawMessage
		value, err := json.Marshal(&obj.parent_type)
		if err != nil {
			return err
		}
		m["parent_type"] = &value
	}
	if len(obj.parent_uuid) > 0 {
		var value json.RawMessage
		value, err := json.Marshal(&obj.parent_uuid)
		if err != nil {
			return err
		}
		m["parent_uuid"] = &value
	}

	return nil
}

type referenceUuidSorter []Reference
func (s referenceUuidSorter) Len() int {
	return len(s)
}
func (s referenceUuidSorter) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s referenceUuidSorter) Less(i, j int) bool {
	lhs, rhs := s[i], s[j]
	return lhs.Uuid < rhs.Uuid
}

func sliceEqual(lhs, rhs reflect.Value) bool {
	len := lhs.Len()
	if len != rhs.Len() {
		return false
	}
	for i := 0; i < len; i++ {
		a1 := lhs.Index(i)
		a2 := rhs.Index(i)
		if a1.Interface() != a2.Interface() {
			return false
		}
	}
	return true
}

func attributeEqual(lhs, rhs LinkAttribute) bool {
	if lhs == nil && rhs == nil {
		return true
	}
	if lhs == nil || rhs == nil {
		return false
	}
	t1 := reflect.TypeOf(lhs).Elem()
	t2 := reflect.TypeOf(rhs).Elem()
	if t1 != t2 {
		return false
	}
	a1 := reflect.ValueOf(lhs).Elem()
	a2 := reflect.ValueOf(rhs).Elem()
	if t1.Kind() == reflect.Slice {
		return sliceEqual(a1, a2)
	}
	return a1.Interface() == a2.Interface()
}

// Helper function that compares two reference lists and generates the
// appropriate list of changes to be resented as POST requests to the
// ref-update URL on the API server.
func (obj *ObjectBase) UpdateReference(
	ptr IObject, field string, current, prev ReferenceList) error {

	sort.Sort(referenceUuidSorter(current))
	sort.Sort(referenceUuidSorter(prev))

	var adds ReferenceList
	var deletes ReferenceList

	i, j := 0, 0
	for i < len(current) && j < len(prev) {
		lhs := current[i]
		rhs := prev[j]
		if lhs.Uuid < rhs.Uuid {
			adds = append(adds, lhs)
			i++
			continue
		} else if lhs.Uuid > rhs.Uuid {
			deletes = append(deletes, rhs)
			j++
			continue
		} else if !attributeEqual(lhs.Attr, rhs.Attr) {
			adds = append(adds, lhs)
			deletes = append(deletes, rhs)
		}
		i++
		j++

	}
	for ;i < len(current); i++ {
		adds = append(adds, current[i])
	}
	for ;j < len(prev); j++ {
		deletes = append(deletes, prev[j])
	}

	for _, ref := range deletes {
		err := obj.clientPtr.UpdateReference(
			&ReferenceUpdateMsg {
				ptr.GetType(),
				obj.uuid, field, ref.Uuid, ref.To,
				"DELETE",
				nil,
			})
		if err != nil {
			return err
		}			
	}

	for _, ref := range adds {
		err := obj.clientPtr.UpdateReference(
			&ReferenceUpdateMsg {
				ptr.GetType(),
				obj.uuid, field, ref.Uuid, ref.To,
				"ADD",
				ref.Attr,
			})
		if err != nil {
			return err
		}
	}

	return nil
}
