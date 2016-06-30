package mocks

import (
	"reflect"
	"strings"
	"unsafe"

	"github.com/Juniper/contrail-go-api"
)

// getReferenceList retrieves the list of UUIDs this object has a forward reference to
// using reflection.
func getReferenceList(obj contrail.IObject) UIDList {
	refList := make([]UID, 0, 8)
	value := reflect.ValueOf(obj).Elem()
	typeval := value.Type()

	for i := 0; i < value.NumField(); i++ {
		name := typeval.Field(i).Name
		if strings.HasSuffix(name, "_back_refs") {
			continue
		}
		if strings.HasSuffix(name, "_refs") {
			field := value.Field(i)
			for i := 0; i < field.Len(); i++ {
				v := field.Index(i)
				ref := (*contrail.Reference)(unsafe.Pointer(v.UnsafeAddr()))
				refList = append(refList, parseUID(ref.Uuid))
			}
		}
	}
	return refList
}

// clearReferenceMask resets the valid bitmask on an object, after an update operation.
// The valid mask is used by the generated types object to keep track of what reference
// list it has a valid cached entry for.
func clearReferenceMask(obj contrail.IObject) {
	value := reflect.ValueOf(obj).Elem()
	field := value.FieldByName("valid")
	maskptr := (*uint64)(unsafe.Pointer(field.UnsafeAddr()))
	*maskptr = 0
}
