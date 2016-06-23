//
// Automatically generated. DO NOT EDIT.
//

package types

import (
        "encoding/json"

        "github.com/michaelhenkel/contrail-go-api"
)

const (
	loadbalancer_member_loadbalancer_member_properties uint64 = 1 << iota
	loadbalancer_member_id_perms
	loadbalancer_member_perms2
	loadbalancer_member_display_name
)

type LoadbalancerMember struct {
        contrail.ObjectBase
	loadbalancer_member_properties LoadbalancerMemberType
	id_perms IdPermsType
	perms2 PermType2
	display_name string
        valid uint64
        modified uint64
        baseMap map[string]contrail.ReferenceList
}

func (obj *LoadbalancerMember) GetType() string {
        return "loadbalancer-member"
}

func (obj *LoadbalancerMember) GetDefaultParent() []string {
        name := []string{"default-domain", "default-project", "default-loadbalancer-pool"}
        return name
}

func (obj *LoadbalancerMember) GetDefaultParentType() string {
        return "loadbalancer-pool"
}

func (obj *LoadbalancerMember) SetName(name string) {
        obj.VSetName(obj, name)
}

func (obj *LoadbalancerMember) SetParent(parent contrail.IObject) {
        obj.VSetParent(obj, parent)
}

func (obj *LoadbalancerMember) storeReferenceBase(
        name string, refList contrail.ReferenceList) {
        if obj.baseMap == nil {
                obj.baseMap = make(map[string]contrail.ReferenceList)
        }
        refCopy := make(contrail.ReferenceList, len(refList))
        copy(refCopy, refList)
        obj.baseMap[name] = refCopy
}

func (obj *LoadbalancerMember) hasReferenceBase(name string) bool {
        if obj.baseMap == nil {
                return false
        }
        _, exists := obj.baseMap[name]
        return exists
}

func (obj *LoadbalancerMember) UpdateDone() {
        obj.modified = 0
        obj.baseMap = nil
}


func (obj *LoadbalancerMember) GetLoadbalancerMemberProperties() LoadbalancerMemberType {
        return obj.loadbalancer_member_properties
}

func (obj *LoadbalancerMember) SetLoadbalancerMemberProperties(value *LoadbalancerMemberType) {
        obj.loadbalancer_member_properties = *value
        obj.modified |= loadbalancer_member_loadbalancer_member_properties
}

func (obj *LoadbalancerMember) GetIdPerms() IdPermsType {
        return obj.id_perms
}

func (obj *LoadbalancerMember) SetIdPerms(value *IdPermsType) {
        obj.id_perms = *value
        obj.modified |= loadbalancer_member_id_perms
}

func (obj *LoadbalancerMember) GetPerms2() PermType2 {
        return obj.perms2
}

func (obj *LoadbalancerMember) SetPerms2(value *PermType2) {
        obj.perms2 = *value
        obj.modified |= loadbalancer_member_perms2
}

func (obj *LoadbalancerMember) GetDisplayName() string {
        return obj.display_name
}

func (obj *LoadbalancerMember) SetDisplayName(value string) {
        obj.display_name = value
        obj.modified |= loadbalancer_member_display_name
}

func (obj *LoadbalancerMember) MarshalJSON() ([]byte, error) {
        msg := map[string]*json.RawMessage {
        }
        err := obj.MarshalCommon(msg)
        if err != nil {
                return nil, err
        }

        if obj.modified & loadbalancer_member_loadbalancer_member_properties != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.loadbalancer_member_properties)
                if err != nil {
                        return nil, err
                }
                msg["loadbalancer_member_properties"] = &value
        }

        if obj.modified & loadbalancer_member_id_perms != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.id_perms)
                if err != nil {
                        return nil, err
                }
                msg["id_perms"] = &value
        }

        if obj.modified & loadbalancer_member_perms2 != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.perms2)
                if err != nil {
                        return nil, err
                }
                msg["perms2"] = &value
        }

        if obj.modified & loadbalancer_member_display_name != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.display_name)
                if err != nil {
                        return nil, err
                }
                msg["display_name"] = &value
        }

        return json.Marshal(msg)
}

func (obj *LoadbalancerMember) UnmarshalJSON(body []byte) error {
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
                case "loadbalancer_member_properties":
                        err = json.Unmarshal(value, &obj.loadbalancer_member_properties)
                        if err == nil {
                                obj.valid |= loadbalancer_member_loadbalancer_member_properties
                        }
                        break
                case "id_perms":
                        err = json.Unmarshal(value, &obj.id_perms)
                        if err == nil {
                                obj.valid |= loadbalancer_member_id_perms
                        }
                        break
                case "perms2":
                        err = json.Unmarshal(value, &obj.perms2)
                        if err == nil {
                                obj.valid |= loadbalancer_member_perms2
                        }
                        break
                case "display_name":
                        err = json.Unmarshal(value, &obj.display_name)
                        if err == nil {
                                obj.valid |= loadbalancer_member_display_name
                        }
                        break
                }
                if err != nil {
                        return err
                }
        }
        return nil
}

func (obj *LoadbalancerMember) UpdateObject() ([]byte, error) {
        msg := map[string]*json.RawMessage {
        }
        err := obj.MarshalId(msg)
        if err != nil {
                return nil, err
        }

        if obj.modified & loadbalancer_member_loadbalancer_member_properties != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.loadbalancer_member_properties)
                if err != nil {
                        return nil, err
                }
                msg["loadbalancer_member_properties"] = &value
        }

        if obj.modified & loadbalancer_member_id_perms != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.id_perms)
                if err != nil {
                        return nil, err
                }
                msg["id_perms"] = &value
        }

        if obj.modified & loadbalancer_member_perms2 != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.perms2)
                if err != nil {
                        return nil, err
                }
                msg["perms2"] = &value
        }

        if obj.modified & loadbalancer_member_display_name != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.display_name)
                if err != nil {
                        return nil, err
                }
                msg["display_name"] = &value
        }

        return json.Marshal(msg)
}

func (obj *LoadbalancerMember) UpdateReferences() error {

        return nil
}

func LoadbalancerMemberByName(c contrail.ApiClient, fqn string) (*LoadbalancerMember, error) {
    obj, err := c.FindByName("loadbalancer-member", fqn)
    if err != nil {
        return nil, err
    }
    return obj.(*LoadbalancerMember), nil
}

func LoadbalancerMemberByUuid(c contrail.ApiClient, uuid string) (*LoadbalancerMember, error) {
    obj, err := c.FindByUuid("loadbalancer-member", uuid)
    if err != nil {
        return nil, err
    }
    return obj.(*LoadbalancerMember), nil
}
