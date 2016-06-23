//
// Automatically generated. DO NOT EDIT.
//

package types

import (
        "encoding/json"

        "github.com/michaelhenkel/contrail-go-api"
)

const (
	network_policy_network_policy_entries uint64 = 1 << iota
	network_policy_id_perms
	network_policy_perms2
	network_policy_display_name
	network_policy_virtual_network_back_refs
)

type NetworkPolicy struct {
        contrail.ObjectBase
	network_policy_entries PolicyEntriesType
	id_perms IdPermsType
	perms2 PermType2
	display_name string
	virtual_network_back_refs contrail.ReferenceList
        valid uint64
        modified uint64
        baseMap map[string]contrail.ReferenceList
}

func (obj *NetworkPolicy) GetType() string {
        return "network-policy"
}

func (obj *NetworkPolicy) GetDefaultParent() []string {
        name := []string{"default-domain", "default-project"}
        return name
}

func (obj *NetworkPolicy) GetDefaultParentType() string {
        return "project"
}

func (obj *NetworkPolicy) SetName(name string) {
        obj.VSetName(obj, name)
}

func (obj *NetworkPolicy) SetParent(parent contrail.IObject) {
        obj.VSetParent(obj, parent)
}

func (obj *NetworkPolicy) storeReferenceBase(
        name string, refList contrail.ReferenceList) {
        if obj.baseMap == nil {
                obj.baseMap = make(map[string]contrail.ReferenceList)
        }
        refCopy := make(contrail.ReferenceList, len(refList))
        copy(refCopy, refList)
        obj.baseMap[name] = refCopy
}

func (obj *NetworkPolicy) hasReferenceBase(name string) bool {
        if obj.baseMap == nil {
                return false
        }
        _, exists := obj.baseMap[name]
        return exists
}

func (obj *NetworkPolicy) UpdateDone() {
        obj.modified = 0
        obj.baseMap = nil
}


func (obj *NetworkPolicy) GetNetworkPolicyEntries() PolicyEntriesType {
        return obj.network_policy_entries
}

func (obj *NetworkPolicy) SetNetworkPolicyEntries(value *PolicyEntriesType) {
        obj.network_policy_entries = *value
        obj.modified |= network_policy_network_policy_entries
}

func (obj *NetworkPolicy) GetIdPerms() IdPermsType {
        return obj.id_perms
}

func (obj *NetworkPolicy) SetIdPerms(value *IdPermsType) {
        obj.id_perms = *value
        obj.modified |= network_policy_id_perms
}

func (obj *NetworkPolicy) GetPerms2() PermType2 {
        return obj.perms2
}

func (obj *NetworkPolicy) SetPerms2(value *PermType2) {
        obj.perms2 = *value
        obj.modified |= network_policy_perms2
}

func (obj *NetworkPolicy) GetDisplayName() string {
        return obj.display_name
}

func (obj *NetworkPolicy) SetDisplayName(value string) {
        obj.display_name = value
        obj.modified |= network_policy_display_name
}

func (obj *NetworkPolicy) readVirtualNetworkBackRefs() error {
        if !obj.IsTransient() &&
                (obj.valid & network_policy_virtual_network_back_refs == 0) {
                err := obj.GetField(obj, "virtual_network_back_refs")
                if err != nil {
                        return err
                }
        }
        return nil
}

func (obj *NetworkPolicy) GetVirtualNetworkBackRefs() (
        contrail.ReferenceList, error) {
        err := obj.readVirtualNetworkBackRefs()
        if err != nil {
                return nil, err
        }
        return obj.virtual_network_back_refs, nil
}

func (obj *NetworkPolicy) MarshalJSON() ([]byte, error) {
        msg := map[string]*json.RawMessage {
        }
        err := obj.MarshalCommon(msg)
        if err != nil {
                return nil, err
        }

        if obj.modified & network_policy_network_policy_entries != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.network_policy_entries)
                if err != nil {
                        return nil, err
                }
                msg["network_policy_entries"] = &value
        }

        if obj.modified & network_policy_id_perms != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.id_perms)
                if err != nil {
                        return nil, err
                }
                msg["id_perms"] = &value
        }

        if obj.modified & network_policy_perms2 != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.perms2)
                if err != nil {
                        return nil, err
                }
                msg["perms2"] = &value
        }

        if obj.modified & network_policy_display_name != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.display_name)
                if err != nil {
                        return nil, err
                }
                msg["display_name"] = &value
        }

        return json.Marshal(msg)
}

func (obj *NetworkPolicy) UnmarshalJSON(body []byte) error {
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
                case "network_policy_entries":
                        err = json.Unmarshal(value, &obj.network_policy_entries)
                        if err == nil {
                                obj.valid |= network_policy_network_policy_entries
                        }
                        break
                case "id_perms":
                        err = json.Unmarshal(value, &obj.id_perms)
                        if err == nil {
                                obj.valid |= network_policy_id_perms
                        }
                        break
                case "perms2":
                        err = json.Unmarshal(value, &obj.perms2)
                        if err == nil {
                                obj.valid |= network_policy_perms2
                        }
                        break
                case "display_name":
                        err = json.Unmarshal(value, &obj.display_name)
                        if err == nil {
                                obj.valid |= network_policy_display_name
                        }
                        break
                case "virtual_network_back_refs": {
                        type ReferenceElement struct {
                                To []string
                                Uuid string
                                Href string
                                Attr VirtualNetworkPolicyType
                        }
                        var array []ReferenceElement
                        err = json.Unmarshal(value, &array)
                        if err != nil {
                            break
                        }
                        obj.valid |= network_policy_virtual_network_back_refs
                        obj.virtual_network_back_refs = make(contrail.ReferenceList, 0)
                        for _, element := range array {
                                ref := contrail.Reference {
                                        element.To,
                                        element.Uuid,
                                        element.Href,
                                        element.Attr,
                                }
                                obj.virtual_network_back_refs = append(obj.virtual_network_back_refs, ref)
                        }
                        break
                }
                }
                if err != nil {
                        return err
                }
        }
        return nil
}

func (obj *NetworkPolicy) UpdateObject() ([]byte, error) {
        msg := map[string]*json.RawMessage {
        }
        err := obj.MarshalId(msg)
        if err != nil {
                return nil, err
        }

        if obj.modified & network_policy_network_policy_entries != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.network_policy_entries)
                if err != nil {
                        return nil, err
                }
                msg["network_policy_entries"] = &value
        }

        if obj.modified & network_policy_id_perms != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.id_perms)
                if err != nil {
                        return nil, err
                }
                msg["id_perms"] = &value
        }

        if obj.modified & network_policy_perms2 != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.perms2)
                if err != nil {
                        return nil, err
                }
                msg["perms2"] = &value
        }

        if obj.modified & network_policy_display_name != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.display_name)
                if err != nil {
                        return nil, err
                }
                msg["display_name"] = &value
        }

        return json.Marshal(msg)
}

func (obj *NetworkPolicy) UpdateReferences() error {

        return nil
}

func NetworkPolicyByName(c contrail.ApiClient, fqn string) (*NetworkPolicy, error) {
    obj, err := c.FindByName("network-policy", fqn)
    if err != nil {
        return nil, err
    }
    return obj.(*NetworkPolicy), nil
}

func NetworkPolicyByUuid(c contrail.ApiClient, uuid string) (*NetworkPolicy, error) {
    obj, err := c.FindByUuid("network-policy", uuid)
    if err != nil {
        return nil, err
    }
    return obj.(*NetworkPolicy), nil
}
