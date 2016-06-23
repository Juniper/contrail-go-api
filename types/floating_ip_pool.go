//
// Automatically generated. DO NOT EDIT.
//

package types

import (
        "encoding/json"

        "github.com/michaelhenkel/contrail-go-api"
)

const (
	floating_ip_pool_floating_ip_pool_prefixes uint64 = 1 << iota
	floating_ip_pool_id_perms
	floating_ip_pool_perms2
	floating_ip_pool_display_name
	floating_ip_pool_floating_ips
	floating_ip_pool_project_back_refs
)

type FloatingIpPool struct {
        contrail.ObjectBase
	floating_ip_pool_prefixes FloatingIpPoolType
	id_perms IdPermsType
	perms2 PermType2
	display_name string
	floating_ips contrail.ReferenceList
	project_back_refs contrail.ReferenceList
        valid uint64
        modified uint64
        baseMap map[string]contrail.ReferenceList
}

func (obj *FloatingIpPool) GetType() string {
        return "floating-ip-pool"
}

func (obj *FloatingIpPool) GetDefaultParent() []string {
        name := []string{"default-domain", "default-project", "default-virtual-network"}
        return name
}

func (obj *FloatingIpPool) GetDefaultParentType() string {
        return "virtual-network"
}

func (obj *FloatingIpPool) SetName(name string) {
        obj.VSetName(obj, name)
}

func (obj *FloatingIpPool) SetParent(parent contrail.IObject) {
        obj.VSetParent(obj, parent)
}

func (obj *FloatingIpPool) storeReferenceBase(
        name string, refList contrail.ReferenceList) {
        if obj.baseMap == nil {
                obj.baseMap = make(map[string]contrail.ReferenceList)
        }
        refCopy := make(contrail.ReferenceList, len(refList))
        copy(refCopy, refList)
        obj.baseMap[name] = refCopy
}

func (obj *FloatingIpPool) hasReferenceBase(name string) bool {
        if obj.baseMap == nil {
                return false
        }
        _, exists := obj.baseMap[name]
        return exists
}

func (obj *FloatingIpPool) UpdateDone() {
        obj.modified = 0
        obj.baseMap = nil
}


func (obj *FloatingIpPool) GetFloatingIpPoolPrefixes() FloatingIpPoolType {
        return obj.floating_ip_pool_prefixes
}

func (obj *FloatingIpPool) SetFloatingIpPoolPrefixes(value *FloatingIpPoolType) {
        obj.floating_ip_pool_prefixes = *value
        obj.modified |= floating_ip_pool_floating_ip_pool_prefixes
}

func (obj *FloatingIpPool) GetIdPerms() IdPermsType {
        return obj.id_perms
}

func (obj *FloatingIpPool) SetIdPerms(value *IdPermsType) {
        obj.id_perms = *value
        obj.modified |= floating_ip_pool_id_perms
}

func (obj *FloatingIpPool) GetPerms2() PermType2 {
        return obj.perms2
}

func (obj *FloatingIpPool) SetPerms2(value *PermType2) {
        obj.perms2 = *value
        obj.modified |= floating_ip_pool_perms2
}

func (obj *FloatingIpPool) GetDisplayName() string {
        return obj.display_name
}

func (obj *FloatingIpPool) SetDisplayName(value string) {
        obj.display_name = value
        obj.modified |= floating_ip_pool_display_name
}

func (obj *FloatingIpPool) readFloatingIps() error {
        if !obj.IsTransient() &&
                (obj.valid & floating_ip_pool_floating_ips == 0) {
                err := obj.GetField(obj, "floating_ips")
                if err != nil {
                        return err
                }
        }
        return nil
}

func (obj *FloatingIpPool) GetFloatingIps() (
        contrail.ReferenceList, error) {
        err := obj.readFloatingIps()
        if err != nil {
                return nil, err
        }
        return obj.floating_ips, nil
}

func (obj *FloatingIpPool) readProjectBackRefs() error {
        if !obj.IsTransient() &&
                (obj.valid & floating_ip_pool_project_back_refs == 0) {
                err := obj.GetField(obj, "project_back_refs")
                if err != nil {
                        return err
                }
        }
        return nil
}

func (obj *FloatingIpPool) GetProjectBackRefs() (
        contrail.ReferenceList, error) {
        err := obj.readProjectBackRefs()
        if err != nil {
                return nil, err
        }
        return obj.project_back_refs, nil
}

func (obj *FloatingIpPool) MarshalJSON() ([]byte, error) {
        msg := map[string]*json.RawMessage {
        }
        err := obj.MarshalCommon(msg)
        if err != nil {
                return nil, err
        }

        if obj.modified & floating_ip_pool_floating_ip_pool_prefixes != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.floating_ip_pool_prefixes)
                if err != nil {
                        return nil, err
                }
                msg["floating_ip_pool_prefixes"] = &value
        }

        if obj.modified & floating_ip_pool_id_perms != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.id_perms)
                if err != nil {
                        return nil, err
                }
                msg["id_perms"] = &value
        }

        if obj.modified & floating_ip_pool_perms2 != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.perms2)
                if err != nil {
                        return nil, err
                }
                msg["perms2"] = &value
        }

        if obj.modified & floating_ip_pool_display_name != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.display_name)
                if err != nil {
                        return nil, err
                }
                msg["display_name"] = &value
        }

        return json.Marshal(msg)
}

func (obj *FloatingIpPool) UnmarshalJSON(body []byte) error {
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
                case "floating_ip_pool_prefixes":
                        err = json.Unmarshal(value, &obj.floating_ip_pool_prefixes)
                        if err == nil {
                                obj.valid |= floating_ip_pool_floating_ip_pool_prefixes
                        }
                        break
                case "id_perms":
                        err = json.Unmarshal(value, &obj.id_perms)
                        if err == nil {
                                obj.valid |= floating_ip_pool_id_perms
                        }
                        break
                case "perms2":
                        err = json.Unmarshal(value, &obj.perms2)
                        if err == nil {
                                obj.valid |= floating_ip_pool_perms2
                        }
                        break
                case "display_name":
                        err = json.Unmarshal(value, &obj.display_name)
                        if err == nil {
                                obj.valid |= floating_ip_pool_display_name
                        }
                        break
                case "floating_ips":
                        err = json.Unmarshal(value, &obj.floating_ips)
                        if err == nil {
                                obj.valid |= floating_ip_pool_floating_ips
                        }
                        break
                case "project_back_refs":
                        err = json.Unmarshal(value, &obj.project_back_refs)
                        if err == nil {
                                obj.valid |= floating_ip_pool_project_back_refs
                        }
                        break
                }
                if err != nil {
                        return err
                }
        }
        return nil
}

func (obj *FloatingIpPool) UpdateObject() ([]byte, error) {
        msg := map[string]*json.RawMessage {
        }
        err := obj.MarshalId(msg)
        if err != nil {
                return nil, err
        }

        if obj.modified & floating_ip_pool_floating_ip_pool_prefixes != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.floating_ip_pool_prefixes)
                if err != nil {
                        return nil, err
                }
                msg["floating_ip_pool_prefixes"] = &value
        }

        if obj.modified & floating_ip_pool_id_perms != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.id_perms)
                if err != nil {
                        return nil, err
                }
                msg["id_perms"] = &value
        }

        if obj.modified & floating_ip_pool_perms2 != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.perms2)
                if err != nil {
                        return nil, err
                }
                msg["perms2"] = &value
        }

        if obj.modified & floating_ip_pool_display_name != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.display_name)
                if err != nil {
                        return nil, err
                }
                msg["display_name"] = &value
        }

        return json.Marshal(msg)
}

func (obj *FloatingIpPool) UpdateReferences() error {

        return nil
}

func FloatingIpPoolByName(c contrail.ApiClient, fqn string) (*FloatingIpPool, error) {
    obj, err := c.FindByName("floating-ip-pool", fqn)
    if err != nil {
        return nil, err
    }
    return obj.(*FloatingIpPool), nil
}

func FloatingIpPoolByUuid(c contrail.ApiClient, uuid string) (*FloatingIpPool, error) {
    obj, err := c.FindByUuid("floating-ip-pool", uuid)
    if err != nil {
        return nil, err
    }
    return obj.(*FloatingIpPool), nil
}
