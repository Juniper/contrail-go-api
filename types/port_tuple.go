//
// Automatically generated. DO NOT EDIT.
//

package types

import (
        "encoding/json"

        "github.com/michaelhenkel/contrail-go-api"
)

const (
	port_tuple_id_perms uint64 = 1 << iota
	port_tuple_perms2
	port_tuple_display_name
	port_tuple_virtual_machine_interface_back_refs
)

type PortTuple struct {
        contrail.ObjectBase
	id_perms IdPermsType
	perms2 PermType2
	display_name string
	virtual_machine_interface_back_refs contrail.ReferenceList
        valid uint64
        modified uint64
        baseMap map[string]contrail.ReferenceList
}

func (obj *PortTuple) GetType() string {
        return "port-tuple"
}

func (obj *PortTuple) GetDefaultParent() []string {
        name := []string{"default-domain", "default-project", "default-service-instance"}
        return name
}

func (obj *PortTuple) GetDefaultParentType() string {
        return "service-instance"
}

func (obj *PortTuple) SetName(name string) {
        obj.VSetName(obj, name)
}

func (obj *PortTuple) SetParent(parent contrail.IObject) {
        obj.VSetParent(obj, parent)
}

func (obj *PortTuple) storeReferenceBase(
        name string, refList contrail.ReferenceList) {
        if obj.baseMap == nil {
                obj.baseMap = make(map[string]contrail.ReferenceList)
        }
        refCopy := make(contrail.ReferenceList, len(refList))
        copy(refCopy, refList)
        obj.baseMap[name] = refCopy
}

func (obj *PortTuple) hasReferenceBase(name string) bool {
        if obj.baseMap == nil {
                return false
        }
        _, exists := obj.baseMap[name]
        return exists
}

func (obj *PortTuple) UpdateDone() {
        obj.modified = 0
        obj.baseMap = nil
}


func (obj *PortTuple) GetIdPerms() IdPermsType {
        return obj.id_perms
}

func (obj *PortTuple) SetIdPerms(value *IdPermsType) {
        obj.id_perms = *value
        obj.modified |= port_tuple_id_perms
}

func (obj *PortTuple) GetPerms2() PermType2 {
        return obj.perms2
}

func (obj *PortTuple) SetPerms2(value *PermType2) {
        obj.perms2 = *value
        obj.modified |= port_tuple_perms2
}

func (obj *PortTuple) GetDisplayName() string {
        return obj.display_name
}

func (obj *PortTuple) SetDisplayName(value string) {
        obj.display_name = value
        obj.modified |= port_tuple_display_name
}

func (obj *PortTuple) readVirtualMachineInterfaceBackRefs() error {
        if !obj.IsTransient() &&
                (obj.valid & port_tuple_virtual_machine_interface_back_refs == 0) {
                err := obj.GetField(obj, "virtual_machine_interface_back_refs")
                if err != nil {
                        return err
                }
        }
        return nil
}

func (obj *PortTuple) GetVirtualMachineInterfaceBackRefs() (
        contrail.ReferenceList, error) {
        err := obj.readVirtualMachineInterfaceBackRefs()
        if err != nil {
                return nil, err
        }
        return obj.virtual_machine_interface_back_refs, nil
}

func (obj *PortTuple) MarshalJSON() ([]byte, error) {
        msg := map[string]*json.RawMessage {
        }
        err := obj.MarshalCommon(msg)
        if err != nil {
                return nil, err
        }

        if obj.modified & port_tuple_id_perms != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.id_perms)
                if err != nil {
                        return nil, err
                }
                msg["id_perms"] = &value
        }

        if obj.modified & port_tuple_perms2 != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.perms2)
                if err != nil {
                        return nil, err
                }
                msg["perms2"] = &value
        }

        if obj.modified & port_tuple_display_name != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.display_name)
                if err != nil {
                        return nil, err
                }
                msg["display_name"] = &value
        }

        return json.Marshal(msg)
}

func (obj *PortTuple) UnmarshalJSON(body []byte) error {
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
                                obj.valid |= port_tuple_id_perms
                        }
                        break
                case "perms2":
                        err = json.Unmarshal(value, &obj.perms2)
                        if err == nil {
                                obj.valid |= port_tuple_perms2
                        }
                        break
                case "display_name":
                        err = json.Unmarshal(value, &obj.display_name)
                        if err == nil {
                                obj.valid |= port_tuple_display_name
                        }
                        break
                case "virtual_machine_interface_back_refs":
                        err = json.Unmarshal(value, &obj.virtual_machine_interface_back_refs)
                        if err == nil {
                                obj.valid |= port_tuple_virtual_machine_interface_back_refs
                        }
                        break
                }
                if err != nil {
                        return err
                }
        }
        return nil
}

func (obj *PortTuple) UpdateObject() ([]byte, error) {
        msg := map[string]*json.RawMessage {
        }
        err := obj.MarshalId(msg)
        if err != nil {
                return nil, err
        }

        if obj.modified & port_tuple_id_perms != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.id_perms)
                if err != nil {
                        return nil, err
                }
                msg["id_perms"] = &value
        }

        if obj.modified & port_tuple_perms2 != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.perms2)
                if err != nil {
                        return nil, err
                }
                msg["perms2"] = &value
        }

        if obj.modified & port_tuple_display_name != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.display_name)
                if err != nil {
                        return nil, err
                }
                msg["display_name"] = &value
        }

        return json.Marshal(msg)
}

func (obj *PortTuple) UpdateReferences() error {

        return nil
}

func PortTupleByName(c contrail.ApiClient, fqn string) (*PortTuple, error) {
    obj, err := c.FindByName("port-tuple", fqn)
    if err != nil {
        return nil, err
    }
    return obj.(*PortTuple), nil
}

func PortTupleByUuid(c contrail.ApiClient, uuid string) (*PortTuple, error) {
    obj, err := c.FindByUuid("port-tuple", uuid)
    if err != nil {
        return nil, err
    }
    return obj.(*PortTuple), nil
}
