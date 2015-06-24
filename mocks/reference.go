package mocks

import (
	"reflect"
	"strings"
	"unsafe"

	"github.com/Juniper/contrail-go-api"
)

func GetReferenceList(obj contrail.IObject) UIDList {
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

func ClearReferenceMask(obj contrail.IObject) {
	value := reflect.ValueOf(obj).Elem()
	field := value.FieldByName("valid")
	maskptr := (*uint64)(unsafe.Pointer(field.UnsafeAddr()))
	*maskptr = 0
}
