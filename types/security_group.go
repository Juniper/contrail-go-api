//
// Automatically generated. DO NOT EDIT.
//

package types

import (
        "encoding/json"

        "github.com/Juniper/contrail-go-api"
)

const (
	security_group_security_group_id uint64 = 1 << iota
	security_group_configured_security_group_id
	security_group_security_group_entries
	security_group_id_perms
	security_group_perms2
	security_group_display_name
	security_group_access_control_lists
	security_group_virtual_machine_interface_back_refs
)

type SecurityGroup struct {
        contrail.ObjectBase
	security_group_id string
	configured_security_group_id int
	security_group_entries PolicyEntriesType
	id_perms IdPermsType
	perms2 PermType2
	display_name string
	access_control_lists contrail.ReferenceList
	virtual_machine_interface_back_refs contrail.ReferenceList
        valid uint64
        modified uint64
        baseMap map[string]contrail.ReferenceList
}

func (obj *SecurityGroup) GetType() string {
        return "security-group"
}

func (obj *SecurityGroup) GetDefaultParent() []string {
        name := []string{"default-domain", "default-project"}
        return name
}

func (obj *SecurityGroup) GetDefaultParentType() string {
        return "project"
}

func (obj *SecurityGroup) SetName(name string) {
        obj.VSetName(obj, name)
}

func (obj *SecurityGroup) SetParent(parent contrail.IObject) {
        obj.VSetParent(obj, parent)
}

func (obj *SecurityGroup) storeReferenceBase(
        name string, refList contrail.ReferenceList) {
        if obj.baseMap == nil {
                obj.baseMap = make(map[string]contrail.ReferenceList)
        }
        refCopy := make(contrail.ReferenceList, len(refList))
        copy(refCopy, refList)
        obj.baseMap[name] = refCopy
}

func (obj *SecurityGroup) hasReferenceBase(name string) bool {
        if obj.baseMap == nil {
                return false
        }
        _, exists := obj.baseMap[name]
        return exists
}

func (obj *SecurityGroup) UpdateDone() {
        obj.modified = 0
        obj.baseMap = nil
}


func (obj *SecurityGroup) GetSecurityGroupId() string {
        return obj.security_group_id
}

func (obj *SecurityGroup) SetSecurityGroupId(value string) {
        obj.security_group_id = value
        obj.modified |= security_group_security_group_id
}

func (obj *SecurityGroup) GetConfiguredSecurityGroupId() int {
        return obj.configured_security_group_id
}

func (obj *SecurityGroup) SetConfiguredSecurityGroupId(value int) {
        obj.configured_security_group_id = value
        obj.modified |= security_group_configured_security_group_id
}

func (obj *SecurityGroup) GetSecurityGroupEntries() PolicyEntriesType {
        return obj.security_group_entries
}

func (obj *SecurityGroup) SetSecurityGroupEntries(value *PolicyEntriesType) {
        obj.security_group_entries = *value
        obj.modified |= security_group_security_group_entries
}

func (obj *SecurityGroup) GetIdPerms() IdPermsType {
        return obj.id_perms
}

func (obj *SecurityGroup) SetIdPerms(value *IdPermsType) {
        obj.id_perms = *value
        obj.modified |= security_group_id_perms
}

func (obj *SecurityGroup) GetPerms2() PermType2 {
        return obj.perms2
}

func (obj *SecurityGroup) SetPerms2(value *PermType2) {
        obj.perms2 = *value
        obj.modified |= security_group_perms2
}

func (obj *SecurityGroup) GetDisplayName() string {
        return obj.display_name
}

func (obj *SecurityGroup) SetDisplayName(value string) {
        obj.display_name = value
        obj.modified |= security_group_display_name
}

func (obj *SecurityGroup) readAccessControlLists() error {
        if !obj.IsTransient() &&
                (obj.valid & security_group_access_control_lists == 0) {
                err := obj.GetField(obj, "access_control_lists")
                if err != nil {
                        return err
                }
        }
        return nil
}

func (obj *SecurityGroup) GetAccessControlLists() (
        contrail.ReferenceList, error) {
        err := obj.readAccessControlLists()
        if err != nil {
                return nil, err
        }
        return obj.access_control_lists, nil
}

func (obj *SecurityGroup) readVirtualMachineInterfaceBackRefs() error {
        if !obj.IsTransient() &&
                (obj.valid & security_group_virtual_machine_interface_back_refs == 0) {
                err := obj.GetField(obj, "virtual_machine_interface_back_refs")
                if err != nil {
                        return err
                }
        }
        return nil
}

func (obj *SecurityGroup) GetVirtualMachineInterfaceBackRefs() (
        contrail.ReferenceList, error) {
        err := obj.readVirtualMachineInterfaceBackRefs()
        if err != nil {
                return nil, err
        }
        return obj.virtual_machine_interface_back_refs, nil
}

func (obj *SecurityGroup) MarshalJSON() ([]byte, error) {
        msg := map[string]*json.RawMessage {
        }
        err := obj.MarshalCommon(msg)
        if err != nil {
                return nil, err
        }

        if obj.modified & security_group_security_group_id != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.security_group_id)
                if err != nil {
                        return nil, err
                }
                msg["security_group_id"] = &value
        }

        if obj.modified & security_group_configured_security_group_id != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.configured_security_group_id)
                if err != nil {
                        return nil, err
                }
                msg["configured_security_group_id"] = &value
        }

        if obj.modified & security_group_security_group_entries != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.security_group_entries)
                if err != nil {
                        return nil, err
                }
                msg["security_group_entries"] = &value
        }

        if obj.modified & security_group_id_perms != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.id_perms)
                if err != nil {
                        return nil, err
                }
                msg["id_perms"] = &value
        }

        if obj.modified & security_group_perms2 != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.perms2)
                if err != nil {
                        return nil, err
                }
                msg["perms2"] = &value
        }

        if obj.modified & security_group_display_name != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.display_name)
                if err != nil {
                        return nil, err
                }
                msg["display_name"] = &value
        }

        return json.Marshal(msg)
}

func (obj *SecurityGroup) UnmarshalJSON(body []byte) error {
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
                case "security_group_id":
                        err = json.Unmarshal(value, &obj.security_group_id)
                        if err == nil {
                                obj.valid |= security_group_security_group_id
                        }
                        break
                case "configured_security_group_id":
                        err = json.Unmarshal(value, &obj.configured_security_group_id)
                        if err == nil {
                                obj.valid |= security_group_configured_security_group_id
                        }
                        break
                case "security_group_entries":
                        err = json.Unmarshal(value, &obj.security_group_entries)
                        if err == nil {
                                obj.valid |= security_group_security_group_entries
                        }
                        break
                case "id_perms":
                        err = json.Unmarshal(value, &obj.id_perms)
                        if err == nil {
                                obj.valid |= security_group_id_perms
                        }
                        break
                case "perms2":
                        err = json.Unmarshal(value, &obj.perms2)
                        if err == nil {
                                obj.valid |= security_group_perms2
                        }
                        break
                case "display_name":
                        err = json.Unmarshal(value, &obj.display_name)
                        if err == nil {
                                obj.valid |= security_group_display_name
                        }
                        break
                case "access_control_lists":
                        err = json.Unmarshal(value, &obj.access_control_lists)
                        if err == nil {
                                obj.valid |= security_group_access_control_lists
                        }
                        break
                case "virtual_machine_interface_back_refs":
                        err = json.Unmarshal(value, &obj.virtual_machine_interface_back_refs)
                        if err == nil {
                                obj.valid |= security_group_virtual_machine_interface_back_refs
                        }
                        break
                }
                if err != nil {
                        return err
                }
        }
        return nil
}

func (obj *SecurityGroup) UpdateObject() ([]byte, error) {
        msg := map[string]*json.RawMessage {
        }
        err := obj.MarshalId(msg)
        if err != nil {
                return nil, err
        }

        if obj.modified & security_group_security_group_id != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.security_group_id)
                if err != nil {
                        return nil, err
                }
                msg["security_group_id"] = &value
        }

        if obj.modified & security_group_configured_security_group_id != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.configured_security_group_id)
                if err != nil {
                        return nil, err
                }
                msg["configured_security_group_id"] = &value
        }

        if obj.modified & security_group_security_group_entries != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.security_group_entries)
                if err != nil {
                        return nil, err
                }
                msg["security_group_entries"] = &value
        }

        if obj.modified & security_group_id_perms != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.id_perms)
                if err != nil {
                        return nil, err
                }
                msg["id_perms"] = &value
        }

        if obj.modified & security_group_perms2 != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.perms2)
                if err != nil {
                        return nil, err
                }
                msg["perms2"] = &value
        }

        if obj.modified & security_group_display_name != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.display_name)
                if err != nil {
                        return nil, err
                }
                msg["display_name"] = &value
        }

        return json.Marshal(msg)
}

func (obj *SecurityGroup) UpdateReferences() error {

        return nil
}

func SecurityGroupByName(c contrail.ApiClient, fqn string) (*SecurityGroup, error) {
    obj, err := c.FindByName("security-group", fqn)
    if err != nil {
        return nil, err
    }
    return obj.(*SecurityGroup), nil
}

func SecurityGroupByUuid(c contrail.ApiClient, uuid string) (*SecurityGroup, error) {
    obj, err := c.FindByUuid("security-group", uuid)
    if err != nil {
        return nil, err
    }
    return obj.(*SecurityGroup), nil
}
