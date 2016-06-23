//
// Automatically generated. DO NOT EDIT.
//

package types

import (
        "encoding/json"

        "github.com/michaelhenkel/contrail-go-api"
)

const (
	dsa_rule_dsa_rule_entry uint64 = 1 << iota
	dsa_rule_id_perms
	dsa_rule_perms2
	dsa_rule_display_name
)

type DsaRule struct {
        contrail.ObjectBase
	dsa_rule_entry DiscoveryServiceAssignmentType
	id_perms IdPermsType
	perms2 PermType2
	display_name string
        valid uint64
        modified uint64
        baseMap map[string]contrail.ReferenceList
}

func (obj *DsaRule) GetType() string {
        return "dsa-rule"
}

func (obj *DsaRule) GetDefaultParent() []string {
        name := []string{"default-discovery-service-assignment"}
        return name
}

func (obj *DsaRule) GetDefaultParentType() string {
        return "discovery-service-assignment"
}

func (obj *DsaRule) SetName(name string) {
        obj.VSetName(obj, name)
}

func (obj *DsaRule) SetParent(parent contrail.IObject) {
        obj.VSetParent(obj, parent)
}

func (obj *DsaRule) storeReferenceBase(
        name string, refList contrail.ReferenceList) {
        if obj.baseMap == nil {
                obj.baseMap = make(map[string]contrail.ReferenceList)
        }
        refCopy := make(contrail.ReferenceList, len(refList))
        copy(refCopy, refList)
        obj.baseMap[name] = refCopy
}

func (obj *DsaRule) hasReferenceBase(name string) bool {
        if obj.baseMap == nil {
                return false
        }
        _, exists := obj.baseMap[name]
        return exists
}

func (obj *DsaRule) UpdateDone() {
        obj.modified = 0
        obj.baseMap = nil
}


func (obj *DsaRule) GetDsaRuleEntry() DiscoveryServiceAssignmentType {
        return obj.dsa_rule_entry
}

func (obj *DsaRule) SetDsaRuleEntry(value *DiscoveryServiceAssignmentType) {
        obj.dsa_rule_entry = *value
        obj.modified |= dsa_rule_dsa_rule_entry
}

func (obj *DsaRule) GetIdPerms() IdPermsType {
        return obj.id_perms
}

func (obj *DsaRule) SetIdPerms(value *IdPermsType) {
        obj.id_perms = *value
        obj.modified |= dsa_rule_id_perms
}

func (obj *DsaRule) GetPerms2() PermType2 {
        return obj.perms2
}

func (obj *DsaRule) SetPerms2(value *PermType2) {
        obj.perms2 = *value
        obj.modified |= dsa_rule_perms2
}

func (obj *DsaRule) GetDisplayName() string {
        return obj.display_name
}

func (obj *DsaRule) SetDisplayName(value string) {
        obj.display_name = value
        obj.modified |= dsa_rule_display_name
}

func (obj *DsaRule) MarshalJSON() ([]byte, error) {
        msg := map[string]*json.RawMessage {
        }
        err := obj.MarshalCommon(msg)
        if err != nil {
                return nil, err
        }

        if obj.modified & dsa_rule_dsa_rule_entry != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.dsa_rule_entry)
                if err != nil {
                        return nil, err
                }
                msg["dsa_rule_entry"] = &value
        }

        if obj.modified & dsa_rule_id_perms != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.id_perms)
                if err != nil {
                        return nil, err
                }
                msg["id_perms"] = &value
        }

        if obj.modified & dsa_rule_perms2 != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.perms2)
                if err != nil {
                        return nil, err
                }
                msg["perms2"] = &value
        }

        if obj.modified & dsa_rule_display_name != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.display_name)
                if err != nil {
                        return nil, err
                }
                msg["display_name"] = &value
        }

        return json.Marshal(msg)
}

func (obj *DsaRule) UnmarshalJSON(body []byte) error {
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
                case "dsa_rule_entry":
                        err = json.Unmarshal(value, &obj.dsa_rule_entry)
                        if err == nil {
                                obj.valid |= dsa_rule_dsa_rule_entry
                        }
                        break
                case "id_perms":
                        err = json.Unmarshal(value, &obj.id_perms)
                        if err == nil {
                                obj.valid |= dsa_rule_id_perms
                        }
                        break
                case "perms2":
                        err = json.Unmarshal(value, &obj.perms2)
                        if err == nil {
                                obj.valid |= dsa_rule_perms2
                        }
                        break
                case "display_name":
                        err = json.Unmarshal(value, &obj.display_name)
                        if err == nil {
                                obj.valid |= dsa_rule_display_name
                        }
                        break
                }
                if err != nil {
                        return err
                }
        }
        return nil
}

func (obj *DsaRule) UpdateObject() ([]byte, error) {
        msg := map[string]*json.RawMessage {
        }
        err := obj.MarshalId(msg)
        if err != nil {
                return nil, err
        }

        if obj.modified & dsa_rule_dsa_rule_entry != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.dsa_rule_entry)
                if err != nil {
                        return nil, err
                }
                msg["dsa_rule_entry"] = &value
        }

        if obj.modified & dsa_rule_id_perms != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.id_perms)
                if err != nil {
                        return nil, err
                }
                msg["id_perms"] = &value
        }

        if obj.modified & dsa_rule_perms2 != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.perms2)
                if err != nil {
                        return nil, err
                }
                msg["perms2"] = &value
        }

        if obj.modified & dsa_rule_display_name != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.display_name)
                if err != nil {
                        return nil, err
                }
                msg["display_name"] = &value
        }

        return json.Marshal(msg)
}

func (obj *DsaRule) UpdateReferences() error {

        return nil
}

func DsaRuleByName(c contrail.ApiClient, fqn string) (*DsaRule, error) {
    obj, err := c.FindByName("dsa-rule", fqn)
    if err != nil {
        return nil, err
    }
    return obj.(*DsaRule), nil
}

func DsaRuleByUuid(c contrail.ApiClient, uuid string) (*DsaRule, error) {
    obj, err := c.FindByUuid("dsa-rule", uuid)
    if err != nil {
        return nil, err
    }
    return obj.(*DsaRule), nil
}
