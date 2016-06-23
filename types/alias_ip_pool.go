//
// Automatically generated. DO NOT EDIT.
//

package types

import (
        "encoding/json"

        "github.com/Juniper/contrail-go-api"
)

const (
	alias_ip_pool_id_perms uint64 = 1 << iota
	alias_ip_pool_perms2
	alias_ip_pool_display_name
	alias_ip_pool_alias_ips
	alias_ip_pool_project_back_refs
)

type AliasIpPool struct {
        contrail.ObjectBase
	id_perms IdPermsType
	perms2 PermType2
	display_name string
	alias_ips contrail.ReferenceList
	project_back_refs contrail.ReferenceList
        valid uint64
        modified uint64
        baseMap map[string]contrail.ReferenceList
}

func (obj *AliasIpPool) GetType() string {
        return "alias-ip-pool"
}

func (obj *AliasIpPool) GetDefaultParent() []string {
        name := []string{"default-domain", "default-project", "default-virtual-network"}
        return name
}

func (obj *AliasIpPool) GetDefaultParentType() string {
        return "virtual-network"
}

func (obj *AliasIpPool) SetName(name string) {
        obj.VSetName(obj, name)
}

func (obj *AliasIpPool) SetParent(parent contrail.IObject) {
        obj.VSetParent(obj, parent)
}

func (obj *AliasIpPool) storeReferenceBase(
        name string, refList contrail.ReferenceList) {
        if obj.baseMap == nil {
                obj.baseMap = make(map[string]contrail.ReferenceList)
        }
        refCopy := make(contrail.ReferenceList, len(refList))
        copy(refCopy, refList)
        obj.baseMap[name] = refCopy
}

func (obj *AliasIpPool) hasReferenceBase(name string) bool {
        if obj.baseMap == nil {
                return false
        }
        _, exists := obj.baseMap[name]
        return exists
}

func (obj *AliasIpPool) UpdateDone() {
        obj.modified = 0
        obj.baseMap = nil
}


func (obj *AliasIpPool) GetIdPerms() IdPermsType {
        return obj.id_perms
}

func (obj *AliasIpPool) SetIdPerms(value *IdPermsType) {
        obj.id_perms = *value
        obj.modified |= alias_ip_pool_id_perms
}

func (obj *AliasIpPool) GetPerms2() PermType2 {
        return obj.perms2
}

func (obj *AliasIpPool) SetPerms2(value *PermType2) {
        obj.perms2 = *value
        obj.modified |= alias_ip_pool_perms2
}

func (obj *AliasIpPool) GetDisplayName() string {
        return obj.display_name
}

func (obj *AliasIpPool) SetDisplayName(value string) {
        obj.display_name = value
        obj.modified |= alias_ip_pool_display_name
}

func (obj *AliasIpPool) readAliasIps() error {
        if !obj.IsTransient() &&
                (obj.valid & alias_ip_pool_alias_ips == 0) {
                err := obj.GetField(obj, "alias_ips")
                if err != nil {
                        return err
                }
        }
        return nil
}

func (obj *AliasIpPool) GetAliasIps() (
        contrail.ReferenceList, error) {
        err := obj.readAliasIps()
        if err != nil {
                return nil, err
        }
        return obj.alias_ips, nil
}

func (obj *AliasIpPool) readProjectBackRefs() error {
        if !obj.IsTransient() &&
                (obj.valid & alias_ip_pool_project_back_refs == 0) {
                err := obj.GetField(obj, "project_back_refs")
                if err != nil {
                        return err
                }
        }
        return nil
}

func (obj *AliasIpPool) GetProjectBackRefs() (
        contrail.ReferenceList, error) {
        err := obj.readProjectBackRefs()
        if err != nil {
                return nil, err
        }
        return obj.project_back_refs, nil
}

func (obj *AliasIpPool) MarshalJSON() ([]byte, error) {
        msg := map[string]*json.RawMessage {
        }
        err := obj.MarshalCommon(msg)
        if err != nil {
                return nil, err
        }

        if obj.modified & alias_ip_pool_id_perms != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.id_perms)
                if err != nil {
                        return nil, err
                }
                msg["id_perms"] = &value
        }

        if obj.modified & alias_ip_pool_perms2 != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.perms2)
                if err != nil {
                        return nil, err
                }
                msg["perms2"] = &value
        }

        if obj.modified & alias_ip_pool_display_name != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.display_name)
                if err != nil {
                        return nil, err
                }
                msg["display_name"] = &value
        }

        return json.Marshal(msg)
}

func (obj *AliasIpPool) UnmarshalJSON(body []byte) error {
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
                case "id_perms":
                        err = json.Unmarshal(value, &obj.id_perms)
                        if err == nil {
                                obj.valid |= alias_ip_pool_id_perms
                        }
                        break
                case "perms2":
                        err = json.Unmarshal(value, &obj.perms2)
                        if err == nil {
                                obj.valid |= alias_ip_pool_perms2
                        }
                        break
                case "display_name":
                        err = json.Unmarshal(value, &obj.display_name)
                        if err == nil {
                                obj.valid |= alias_ip_pool_display_name
                        }
                        break
                case "alias_ips":
                        err = json.Unmarshal(value, &obj.alias_ips)
                        if err == nil {
                                obj.valid |= alias_ip_pool_alias_ips
                        }
                        break
                case "project_back_refs":
                        err = json.Unmarshal(value, &obj.project_back_refs)
                        if err == nil {
                                obj.valid |= alias_ip_pool_project_back_refs
                        }
                        break
                }
                if err != nil {
                        return err
                }
        }
        return nil
}

func (obj *AliasIpPool) UpdateObject() ([]byte, error) {
        msg := map[string]*json.RawMessage {
        }
        err := obj.MarshalId(msg)
        if err != nil {
                return nil, err
        }

        if obj.modified & alias_ip_pool_id_perms != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.id_perms)
                if err != nil {
                        return nil, err
                }
                msg["id_perms"] = &value
        }

        if obj.modified & alias_ip_pool_perms2 != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.perms2)
                if err != nil {
                        return nil, err
                }
                msg["perms2"] = &value
        }

        if obj.modified & alias_ip_pool_display_name != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.display_name)
                if err != nil {
                        return nil, err
                }
                msg["display_name"] = &value
        }

        return json.Marshal(msg)
}

func (obj *AliasIpPool) UpdateReferences() error {

        return nil
}

func AliasIpPoolByName(c contrail.ApiClient, fqn string) (*AliasIpPool, error) {
    obj, err := c.FindByName("alias-ip-pool", fqn)
    if err != nil {
        return nil, err
    }
    return obj.(*AliasIpPool), nil
}

func AliasIpPoolByUuid(c contrail.ApiClient, uuid string) (*AliasIpPool, error) {
    obj, err := c.FindByUuid("alias-ip-pool", uuid)
    if err != nil {
        return nil, err
    }
    return obj.(*AliasIpPool), nil
}
