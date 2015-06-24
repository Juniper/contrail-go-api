package mocks

import (
	"encoding/json"
	"strings"

	"github.com/Juniper/contrail-go-api"
)

type objectUpdater struct {
	db Database
}

func NewObjectUpdater(db Database) *objectUpdater {
	updater := new(objectUpdater)
	updater.db = db
	return updater
}

// On forward refs do nothing. For children and back_refs fetch content from DB.
func (u *objectUpdater) GetField(obj contrail.IObject, field string) error {
	var idList UIDList
	if strings.HasSuffix(field, "_back_refs") {
		typename := field[:len(field)-len("_back_refs")]
		var err error
		idList, err = u.db.GetBackReferences(parseUID(obj.GetUuid()), typename)
		if err != nil {
			return err
		}
	} else if strings.HasSuffix(field, "_refs") {
		return nil
	} else {
		typename := field[:len(field)-1]
		var err error
		idList, err = u.db.GetChildren(parseUID(obj.GetUuid()), typename)
		if err != nil {
			return err
		}
	}
	if len(idList) == 0 {
		return nil
	}

	refList := make(contrail.ReferenceList, len(idList))
	for i, id := range idList {
		refList[i].Uuid = id.Interface().String()
		obj, err := u.db.GetByUuid(id.Interface())
		if err != nil {
			continue
		}
		fqn := obj.GetFQName()
		refList[i].To = make([]string, len(fqn))
		copy(refList[i].To, fqn)
	}
	listJson, err := json.Marshal(refList)
	if err != nil {
		return err
	}
	var listRaw json.RawMessage = listJson

	fqnJson, err := json.Marshal(obj.GetFQName())
	if err != nil {
		return err
	}
	fqnRaw := json.RawMessage(fqnJson)

	uuidJson, err := json.Marshal(obj.GetUuid())
	if err != nil {
		return err
	}
	uuidRaw := json.RawMessage(uuidJson)

	nameJson, err := json.Marshal(obj.GetName())
	if err != nil {
		return err
	}
	nameRaw := json.RawMessage(nameJson)

	msg := map[string]*json.RawMessage{
		"fq_name": &fqnRaw,
		"uuid":    &uuidRaw,
		"name":    &nameRaw,
		field:     &listRaw,
	}
	data, err := json.Marshal(msg)
	if err != nil {
		return err
	}
	return json.Unmarshal(data, obj)
}

func (u *objectUpdater) UpdateReference(msg *contrail.ReferenceUpdateMsg) error {
	return nil
}
