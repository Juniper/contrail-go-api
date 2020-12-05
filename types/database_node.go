//
// Automatically generated. DO NOT EDIT.
//

package types

import (
        "encoding/json"

        "github.com/Juniper/contrail-go-api"
)

const (
	database_node_database_node_ip_address uint64 = 1 << iota
	database_node_id_perms
	database_node_perms2
	database_node_display_name
)

type DatabaseNode struct {
        contrail.ObjectBase
	database_node_ip_address string
	id_perms IdPermsType
	perms2 PermType2
	display_name string
        valid uint64
        modified uint64
        baseMap map[string]contrail.ReferenceList
}

func (obj *DatabaseNode) GetType() string {
        return "database-node"
}

func (obj *DatabaseNode) GetDefaultParent() []string {
        name := []string{"default-global-system-config"}
        return name
}

func (obj *DatabaseNode) GetDefaultParentType() string {
        return "global-system-config"
}

func (obj *DatabaseNode) SetName(name string) {
        obj.VSetName(obj, name)
}

func (obj *DatabaseNode) SetParent(parent contrail.IObject) {
        obj.VSetParent(obj, parent)
}

func (obj *DatabaseNode) storeReferenceBase(
        name string, refList contrail.ReferenceList) {
        if obj.baseMap == nil {
                obj.baseMap = make(map[string]contrail.ReferenceList)
        }
        refCopy := make(contrail.ReferenceList, len(refList))
        copy(refCopy, refList)
        obj.baseMap[name] = refCopy
}

func (obj *DatabaseNode) hasReferenceBase(name string) bool {
        if obj.baseMap == nil {
                return false
        }
        _, exists := obj.baseMap[name]
        return exists
}

func (obj *DatabaseNode) UpdateDone() {
        obj.modified = 0
        obj.baseMap = nil
}


func (obj *DatabaseNode) GetDatabaseNodeIpAddress() string {
        return obj.database_node_ip_address
}

func (obj *DatabaseNode) SetDatabaseNodeIpAddress(value string) {
        obj.database_node_ip_address = value
        obj.modified |= database_node_database_node_ip_address
}

func (obj *DatabaseNode) GetIdPerms() IdPermsType {
        return obj.id_perms
}

func (obj *DatabaseNode) SetIdPerms(value *IdPermsType) {
        obj.id_perms = *value
        obj.modified |= database_node_id_perms
}

func (obj *DatabaseNode) GetPerms2() PermType2 {
        return obj.perms2
}

func (obj *DatabaseNode) SetPerms2(value *PermType2) {
        obj.perms2 = *value
        obj.modified |= database_node_perms2
}

func (obj *DatabaseNode) GetDisplayName() string {
        return obj.display_name
}

func (obj *DatabaseNode) SetDisplayName(value string) {
        obj.display_name = value
        obj.modified |= database_node_display_name
}

func (obj *DatabaseNode) MarshalJSON() ([]byte, error) {
        msg := map[string]*json.RawMessage {
        }
        err := obj.MarshalCommon(msg)
        if err != nil {
                return nil, err
        }

        if obj.modified & database_node_database_node_ip_address != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.database_node_ip_address)
                if err != nil {
                        return nil, err
                }
                msg["database_node_ip_address"] = &value
        }

        if obj.modified & database_node_id_perms != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.id_perms)
                if err != nil {
                        return nil, err
                }
                msg["id_perms"] = &value
        }

        if obj.modified & database_node_perms2 != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.perms2)
                if err != nil {
                        return nil, err
                }
                msg["perms2"] = &value
        }

        if obj.modified & database_node_display_name != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.display_name)
                if err != nil {
                        return nil, err
                }
                msg["display_name"] = &value
        }

        return json.Marshal(msg)
}

func (obj *DatabaseNode) UnmarshalJSON(body []byte) error {
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
                case "database_node_ip_address":
                        err = json.Unmarshal(value, &obj.database_node_ip_address)
                        if err == nil {
                                obj.valid |= database_node_database_node_ip_address
                        }
                        break
                case "id_perms":
                        err = json.Unmarshal(value, &obj.id_perms)
                        if err == nil {
                                obj.valid |= database_node_id_perms
                        }
                        break
                case "perms2":
                        err = json.Unmarshal(value, &obj.perms2)
                        if err == nil {
                                obj.valid |= database_node_perms2
                        }
                        break
                case "display_name":
                        err = json.Unmarshal(value, &obj.display_name)
                        if err == nil {
                                obj.valid |= database_node_display_name
                        }
                        break
                }
                if err != nil {
                        return err
                }
        }
        return nil
}

func (obj *DatabaseNode) UpdateObject() ([]byte, error) {
        msg := map[string]*json.RawMessage {
        }
        err := obj.MarshalId(msg)
        if err != nil {
                return nil, err
        }

        if obj.modified & database_node_database_node_ip_address != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.database_node_ip_address)
                if err != nil {
                        return nil, err
                }
                msg["database_node_ip_address"] = &value
        }

        if obj.modified & database_node_id_perms != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.id_perms)
                if err != nil {
                        return nil, err
                }
                msg["id_perms"] = &value
        }

        if obj.modified & database_node_perms2 != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.perms2)
                if err != nil {
                        return nil, err
                }
                msg["perms2"] = &value
        }

        if obj.modified & database_node_display_name != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.display_name)
                if err != nil {
                        return nil, err
                }
                msg["display_name"] = &value
        }

        return json.Marshal(msg)
}

func (obj *DatabaseNode) UpdateReferences() error {

        return nil
}

func DatabaseNodeByName(c contrail.ApiClient, fqn string) (*DatabaseNode, error) {
    obj, err := c.FindByName("database-node", fqn)
    if err != nil {
        return nil, err
    }
    return obj.(*DatabaseNode), nil
}

func DatabaseNodeByUuid(c contrail.ApiClient, uuid string) (*DatabaseNode, error) {
    obj, err := c.FindByUuid("database-node", uuid)
    if err != nil {
        return nil, err
    }
    return obj.(*DatabaseNode), nil
}
