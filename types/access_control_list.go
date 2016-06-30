//
// Automatically generated. DO NOT EDIT.
//

package types

import (
        "encoding/json"

        "github.com/Juniper/contrail-go-api"
)

const (
	access_control_list_access_control_list_entries uint64 = 1 << iota
	access_control_list_id_perms
	access_control_list_perms2
	access_control_list_display_name
)

type AccessControlList struct {
        contrail.ObjectBase
	access_control_list_entries AclEntriesType
	id_perms IdPermsType
	perms2 PermType2
	display_name string
        valid uint64
        modified uint64
        baseMap map[string]contrail.ReferenceList
}

func (obj *AccessControlList) GetType() string {
        return "access-control-list"
}

func (obj *AccessControlList) GetDefaultParent() []string {
        name := []string{"default-domain", "default-project", "default-virtual-network"}
        return name
}

func (obj *AccessControlList) GetDefaultParentType() string {
        return "virtual-network"
}

func (obj *AccessControlList) SetName(name string) {
        obj.VSetName(obj, name)
}

func (obj *AccessControlList) SetParent(parent contrail.IObject) {
        obj.VSetParent(obj, parent)
}

func (obj *AccessControlList) storeReferenceBase(
        name string, refList contrail.ReferenceList) {
        if obj.baseMap == nil {
                obj.baseMap = make(map[string]contrail.ReferenceList)
        }
        refCopy := make(contrail.ReferenceList, len(refList))
        copy(refCopy, refList)
        obj.baseMap[name] = refCopy
}

func (obj *AccessControlList) hasReferenceBase(name string) bool {
        if obj.baseMap == nil {
                return false
        }
        _, exists := obj.baseMap[name]
        return exists
}

func (obj *AccessControlList) UpdateDone() {
        obj.modified = 0
        obj.baseMap = nil
}


func (obj *AccessControlList) GetAccessControlListEntries() AclEntriesType {
        return obj.access_control_list_entries
}

func (obj *AccessControlList) SetAccessControlListEntries(value *AclEntriesType) {
        obj.access_control_list_entries = *value
        obj.modified |= access_control_list_access_control_list_entries
}

func (obj *AccessControlList) GetIdPerms() IdPermsType {
        return obj.id_perms
}

func (obj *AccessControlList) SetIdPerms(value *IdPermsType) {
        obj.id_perms = *value
        obj.modified |= access_control_list_id_perms
}

func (obj *AccessControlList) GetPerms2() PermType2 {
        return obj.perms2
}

func (obj *AccessControlList) SetPerms2(value *PermType2) {
        obj.perms2 = *value
        obj.modified |= access_control_list_perms2
}

func (obj *AccessControlList) GetDisplayName() string {
        return obj.display_name
}

func (obj *AccessControlList) SetDisplayName(value string) {
        obj.display_name = value
        obj.modified |= access_control_list_display_name
}

func (obj *AccessControlList) MarshalJSON() ([]byte, error) {
        msg := map[string]*json.RawMessage {
        }
        err := obj.MarshalCommon(msg)
        if err != nil {
                return nil, err
        }

        if obj.modified & access_control_list_access_control_list_entries != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.access_control_list_entries)
                if err != nil {
                        return nil, err
                }
                msg["access_control_list_entries"] = &value
        }

        if obj.modified & access_control_list_id_perms != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.id_perms)
                if err != nil {
                        return nil, err
                }
                msg["id_perms"] = &value
        }

        if obj.modified & access_control_list_perms2 != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.perms2)
                if err != nil {
                        return nil, err
                }
                msg["perms2"] = &value
        }

        if obj.modified & access_control_list_display_name != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.display_name)
                if err != nil {
                        return nil, err
                }
                msg["display_name"] = &value
        }

        return json.Marshal(msg)
}

func (obj *AccessControlList) UnmarshalJSON(body []byte) error {
        var m map[string]json.RawMessage
        err := json.Unmarshal(body, &m)
        if err != nil {
                return err
        }
        err = obj.UnmarshalCommon(m)
        if err != nil {
                return err
        }
        for key, value := range m {
                switch key {
                case "access_control_list_entries":
                        err = json.Unmarshal(value, &obj.access_control_list_entries)
                        if err == nil {
                                obj.valid |= access_control_list_access_control_list_entries
                        }
                        break
                case "id_perms":
                        err = json.Unmarshal(value, &obj.id_perms)
                        if err == nil {
                                obj.valid |= access_control_list_id_perms
                        }
                        break
                case "perms2":
                        err = json.Unmarshal(value, &obj.perms2)
                        if err == nil {
                                obj.valid |= access_control_list_perms2
                        }
                        break
                case "display_name":
                        err = json.Unmarshal(value, &obj.display_name)
                        if err == nil {
                                obj.valid |= access_control_list_display_name
                        }
                        break
                }
                if err != nil {
                        return err
                }
        }
        return nil
}

func (obj *AccessControlList) UpdateObject() ([]byte, error) {
        msg := map[string]*json.RawMessage {
        }
        err := obj.MarshalId(msg)
        if err != nil {
                return nil, err
        }

        if obj.modified & access_control_list_access_control_list_entries != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.access_control_list_entries)
                if err != nil {
                        return nil, err
                }
                msg["access_control_list_entries"] = &value
        }

        if obj.modified & access_control_list_id_perms != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.id_perms)
                if err != nil {
                        return nil, err
                }
                msg["id_perms"] = &value
        }

        if obj.modified & access_control_list_perms2 != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.perms2)
                if err != nil {
                        return nil, err
                }
                msg["perms2"] = &value
        }

        if obj.modified & access_control_list_display_name != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.display_name)
                if err != nil {
                        return nil, err
                }
                msg["display_name"] = &value
        }

        return json.Marshal(msg)
}

func (obj *AccessControlList) UpdateReferences() error {

        return nil
}

func AccessControlListByName(c contrail.ApiClient, fqn string) (*AccessControlList, error) {
    obj, err := c.FindByName("access-control-list", fqn)
    if err != nil {
        return nil, err
    }
    return obj.(*AccessControlList), nil
}

func AccessControlListByUuid(c contrail.ApiClient, uuid string) (*AccessControlList, error) {
    obj, err := c.FindByUuid("access-control-list", uuid)
    if err != nil {
        return nil, err
    }
    return obj.(*AccessControlList), nil
}
