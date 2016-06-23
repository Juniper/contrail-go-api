//
// Automatically generated. DO NOT EDIT.
//

package types

import (
        "encoding/json"

        "github.com/michaelhenkel/contrail-go-api"
)

const (
	analytics_node_analytics_node_ip_address uint64 = 1 << iota
	analytics_node_id_perms
	analytics_node_perms2
	analytics_node_display_name
)

type AnalyticsNode struct {
        contrail.ObjectBase
	analytics_node_ip_address string
	id_perms IdPermsType
	perms2 PermType2
	display_name string
        valid uint64
        modified uint64
        baseMap map[string]contrail.ReferenceList
}

func (obj *AnalyticsNode) GetType() string {
        return "analytics-node"
}

func (obj *AnalyticsNode) GetDefaultParent() []string {
        name := []string{"default-global-system-config"}
        return name
}

func (obj *AnalyticsNode) GetDefaultParentType() string {
        return "global-system-config"
}

func (obj *AnalyticsNode) SetName(name string) {
        obj.VSetName(obj, name)
}

func (obj *AnalyticsNode) SetParent(parent contrail.IObject) {
        obj.VSetParent(obj, parent)
}

func (obj *AnalyticsNode) storeReferenceBase(
        name string, refList contrail.ReferenceList) {
        if obj.baseMap == nil {
                obj.baseMap = make(map[string]contrail.ReferenceList)
        }
        refCopy := make(contrail.ReferenceList, len(refList))
        copy(refCopy, refList)
        obj.baseMap[name] = refCopy
}

func (obj *AnalyticsNode) hasReferenceBase(name string) bool {
        if obj.baseMap == nil {
                return false
        }
        _, exists := obj.baseMap[name]
        return exists
}

func (obj *AnalyticsNode) UpdateDone() {
        obj.modified = 0
        obj.baseMap = nil
}


func (obj *AnalyticsNode) GetAnalyticsNodeIpAddress() string {
        return obj.analytics_node_ip_address
}

func (obj *AnalyticsNode) SetAnalyticsNodeIpAddress(value string) {
        obj.analytics_node_ip_address = value
        obj.modified |= analytics_node_analytics_node_ip_address
}

func (obj *AnalyticsNode) GetIdPerms() IdPermsType {
        return obj.id_perms
}

func (obj *AnalyticsNode) SetIdPerms(value *IdPermsType) {
        obj.id_perms = *value
        obj.modified |= analytics_node_id_perms
}

func (obj *AnalyticsNode) GetPerms2() PermType2 {
        return obj.perms2
}

func (obj *AnalyticsNode) SetPerms2(value *PermType2) {
        obj.perms2 = *value
        obj.modified |= analytics_node_perms2
}

func (obj *AnalyticsNode) GetDisplayName() string {
        return obj.display_name
}

func (obj *AnalyticsNode) SetDisplayName(value string) {
        obj.display_name = value
        obj.modified |= analytics_node_display_name
}

func (obj *AnalyticsNode) MarshalJSON() ([]byte, error) {
        msg := map[string]*json.RawMessage {
        }
        err := obj.MarshalCommon(msg)
        if err != nil {
                return nil, err
        }

        if obj.modified & analytics_node_analytics_node_ip_address != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.analytics_node_ip_address)
                if err != nil {
                        return nil, err
                }
                msg["analytics_node_ip_address"] = &value
        }

        if obj.modified & analytics_node_id_perms != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.id_perms)
                if err != nil {
                        return nil, err
                }
                msg["id_perms"] = &value
        }

        if obj.modified & analytics_node_perms2 != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.perms2)
                if err != nil {
                        return nil, err
                }
                msg["perms2"] = &value
        }

        if obj.modified & analytics_node_display_name != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.display_name)
                if err != nil {
                        return nil, err
                }
                msg["display_name"] = &value
        }

        return json.Marshal(msg)
}

func (obj *AnalyticsNode) UnmarshalJSON(body []byte) error {
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
                case "analytics_node_ip_address":
                        err = json.Unmarshal(value, &obj.analytics_node_ip_address)
                        if err == nil {
                                obj.valid |= analytics_node_analytics_node_ip_address
                        }
                        break
                case "id_perms":
                        err = json.Unmarshal(value, &obj.id_perms)
                        if err == nil {
                                obj.valid |= analytics_node_id_perms
                        }
                        break
                case "perms2":
                        err = json.Unmarshal(value, &obj.perms2)
                        if err == nil {
                                obj.valid |= analytics_node_perms2
                        }
                        break
                case "display_name":
                        err = json.Unmarshal(value, &obj.display_name)
                        if err == nil {
                                obj.valid |= analytics_node_display_name
                        }
                        break
                }
                if err != nil {
                        return err
                }
        }
        return nil
}

func (obj *AnalyticsNode) UpdateObject() ([]byte, error) {
        msg := map[string]*json.RawMessage {
        }
        err := obj.MarshalId(msg)
        if err != nil {
                return nil, err
        }

        if obj.modified & analytics_node_analytics_node_ip_address != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.analytics_node_ip_address)
                if err != nil {
                        return nil, err
                }
                msg["analytics_node_ip_address"] = &value
        }

        if obj.modified & analytics_node_id_perms != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.id_perms)
                if err != nil {
                        return nil, err
                }
                msg["id_perms"] = &value
        }

        if obj.modified & analytics_node_perms2 != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.perms2)
                if err != nil {
                        return nil, err
                }
                msg["perms2"] = &value
        }

        if obj.modified & analytics_node_display_name != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.display_name)
                if err != nil {
                        return nil, err
                }
                msg["display_name"] = &value
        }

        return json.Marshal(msg)
}

func (obj *AnalyticsNode) UpdateReferences() error {

        return nil
}

func AnalyticsNodeByName(c contrail.ApiClient, fqn string) (*AnalyticsNode, error) {
    obj, err := c.FindByName("analytics-node", fqn)
    if err != nil {
        return nil, err
    }
    return obj.(*AnalyticsNode), nil
}

func AnalyticsNodeByUuid(c contrail.ApiClient, uuid string) (*AnalyticsNode, error) {
    obj, err := c.FindByUuid("analytics-node", uuid)
    if err != nil {
        return nil, err
    }
    return obj.(*AnalyticsNode), nil
}
