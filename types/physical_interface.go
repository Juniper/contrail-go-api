//
// Automatically generated. DO NOT EDIT.
//

package types

import (
        "encoding/json"

        "github.com/Juniper/contrail-go-api"
)

const (
	physical_interface_id_perms uint64 = 1 << iota
	physical_interface_perms2
	physical_interface_display_name
	physical_interface_logical_interfaces
	physical_interface_physical_interface_refs
	physical_interface_service_appliance_back_refs
	physical_interface_virtual_machine_interface_back_refs
	physical_interface_physical_interface_back_refs
)

type PhysicalInterface struct {
        contrail.ObjectBase
	id_perms IdPermsType
	perms2 PermType2
	display_name string
	logical_interfaces contrail.ReferenceList
	physical_interface_refs contrail.ReferenceList
	service_appliance_back_refs contrail.ReferenceList
	virtual_machine_interface_back_refs contrail.ReferenceList
	physical_interface_back_refs contrail.ReferenceList
        valid uint64
        modified uint64
        baseMap map[string]contrail.ReferenceList
}

func (obj *PhysicalInterface) GetType() string {
        return "physical-interface"
}

func (obj *PhysicalInterface) GetDefaultParent() []string {
        name := []string{"default-global-system-config", "default-physical-router"}
        return name
}

func (obj *PhysicalInterface) GetDefaultParentType() string {
        return "physical-router"
}

func (obj *PhysicalInterface) SetName(name string) {
        obj.VSetName(obj, name)
}

func (obj *PhysicalInterface) SetParent(parent contrail.IObject) {
        obj.VSetParent(obj, parent)
}

func (obj *PhysicalInterface) storeReferenceBase(
        name string, refList contrail.ReferenceList) {
        if obj.baseMap == nil {
                obj.baseMap = make(map[string]contrail.ReferenceList)
        }
        refCopy := make(contrail.ReferenceList, len(refList))
        copy(refCopy, refList)
        obj.baseMap[name] = refCopy
}

func (obj *PhysicalInterface) hasReferenceBase(name string) bool {
        if obj.baseMap == nil {
                return false
        }
        _, exists := obj.baseMap[name]
        return exists
}

func (obj *PhysicalInterface) UpdateDone() {
        obj.modified = 0
        obj.baseMap = nil
}


func (obj *PhysicalInterface) GetIdPerms() IdPermsType {
        return obj.id_perms
}

func (obj *PhysicalInterface) SetIdPerms(value *IdPermsType) {
        obj.id_perms = *value
        obj.modified |= physical_interface_id_perms
}

func (obj *PhysicalInterface) GetPerms2() PermType2 {
        return obj.perms2
}

func (obj *PhysicalInterface) SetPerms2(value *PermType2) {
        obj.perms2 = *value
        obj.modified |= physical_interface_perms2
}

func (obj *PhysicalInterface) GetDisplayName() string {
        return obj.display_name
}

func (obj *PhysicalInterface) SetDisplayName(value string) {
        obj.display_name = value
        obj.modified |= physical_interface_display_name
}

func (obj *PhysicalInterface) readLogicalInterfaces() error {
        if !obj.IsTransient() &&
                (obj.valid & physical_interface_logical_interfaces == 0) {
                err := obj.GetField(obj, "logical_interfaces")
                if err != nil {
                        return err
                }
        }
        return nil
}

func (obj *PhysicalInterface) GetLogicalInterfaces() (
        contrail.ReferenceList, error) {
        err := obj.readLogicalInterfaces()
        if err != nil {
                return nil, err
        }
        return obj.logical_interfaces, nil
}

func (obj *PhysicalInterface) readPhysicalInterfaceRefs() error {
        if !obj.IsTransient() &&
                (obj.valid & physical_interface_physical_interface_refs == 0) {
                err := obj.GetField(obj, "physical_interface_refs")
                if err != nil {
                        return err
                }
        }
        return nil
}

func (obj *PhysicalInterface) GetPhysicalInterfaceRefs() (
        contrail.ReferenceList, error) {
        err := obj.readPhysicalInterfaceRefs()
        if err != nil {
                return nil, err
        }
        return obj.physical_interface_refs, nil
}

func (obj *PhysicalInterface) AddPhysicalInterface(
        rhs *PhysicalInterface) error {
        err := obj.readPhysicalInterfaceRefs()
        if err != nil {
                return err
        }

        if obj.modified & physical_interface_physical_interface_refs == 0 {
                obj.storeReferenceBase("physical-interface", obj.physical_interface_refs)
        }

        ref := contrail.Reference {
                rhs.GetFQName(), rhs.GetUuid(), rhs.GetHref(), nil}
        obj.physical_interface_refs = append(obj.physical_interface_refs, ref)
        obj.modified |= physical_interface_physical_interface_refs
        return nil
}

func (obj *PhysicalInterface) DeletePhysicalInterface(uuid string) error {
        err := obj.readPhysicalInterfaceRefs()
        if err != nil {
                return err
        }

        if obj.modified & physical_interface_physical_interface_refs == 0 {
                obj.storeReferenceBase("physical-interface", obj.physical_interface_refs)
        }

        for i, ref := range obj.physical_interface_refs {
                if ref.Uuid == uuid {
                        obj.physical_interface_refs = append(
                                obj.physical_interface_refs[:i],
                                obj.physical_interface_refs[i+1:]...)
                        break
                }
        }
        obj.modified |= physical_interface_physical_interface_refs
        return nil
}

func (obj *PhysicalInterface) ClearPhysicalInterface() {
        if (obj.valid & physical_interface_physical_interface_refs != 0) &&
           (obj.modified & physical_interface_physical_interface_refs == 0) {
                obj.storeReferenceBase("physical-interface", obj.physical_interface_refs)
        }
        obj.physical_interface_refs = make([]contrail.Reference, 0)
        obj.valid |= physical_interface_physical_interface_refs
        obj.modified |= physical_interface_physical_interface_refs
}

func (obj *PhysicalInterface) SetPhysicalInterfaceList(
        refList []contrail.ReferencePair) {
        obj.ClearPhysicalInterface()
        obj.physical_interface_refs = make([]contrail.Reference, len(refList))
        for i, pair := range refList {
                obj.physical_interface_refs[i] = contrail.Reference {
                        pair.Object.GetFQName(),
                        pair.Object.GetUuid(),
                        pair.Object.GetHref(),
                        pair.Attribute,
                }
        }
}


func (obj *PhysicalInterface) readServiceApplianceBackRefs() error {
        if !obj.IsTransient() &&
                (obj.valid & physical_interface_service_appliance_back_refs == 0) {
                err := obj.GetField(obj, "service_appliance_back_refs")
                if err != nil {
                        return err
                }
        }
        return nil
}

func (obj *PhysicalInterface) GetServiceApplianceBackRefs() (
        contrail.ReferenceList, error) {
        err := obj.readServiceApplianceBackRefs()
        if err != nil {
                return nil, err
        }
        return obj.service_appliance_back_refs, nil
}

func (obj *PhysicalInterface) readVirtualMachineInterfaceBackRefs() error {
        if !obj.IsTransient() &&
                (obj.valid & physical_interface_virtual_machine_interface_back_refs == 0) {
                err := obj.GetField(obj, "virtual_machine_interface_back_refs")
                if err != nil {
                        return err
                }
        }
        return nil
}

func (obj *PhysicalInterface) GetVirtualMachineInterfaceBackRefs() (
        contrail.ReferenceList, error) {
        err := obj.readVirtualMachineInterfaceBackRefs()
        if err != nil {
                return nil, err
        }
        return obj.virtual_machine_interface_back_refs, nil
}

func (obj *PhysicalInterface) readPhysicalInterfaceBackRefs() error {
        if !obj.IsTransient() &&
                (obj.valid & physical_interface_physical_interface_back_refs == 0) {
                err := obj.GetField(obj, "physical_interface_back_refs")
                if err != nil {
                        return err
                }
        }
        return nil
}

func (obj *PhysicalInterface) GetPhysicalInterfaceBackRefs() (
        contrail.ReferenceList, error) {
        err := obj.readPhysicalInterfaceBackRefs()
        if err != nil {
                return nil, err
        }
        return obj.physical_interface_back_refs, nil
}

func (obj *PhysicalInterface) MarshalJSON() ([]byte, error) {
        msg := map[string]*json.RawMessage {
        }
        err := obj.MarshalCommon(msg)
        if err != nil {
                return nil, err
        }

        if obj.modified & physical_interface_id_perms != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.id_perms)
                if err != nil {
                        return nil, err
                }
                msg["id_perms"] = &value
        }

        if obj.modified & physical_interface_perms2 != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.perms2)
                if err != nil {
                        return nil, err
                }
                msg["perms2"] = &value
        }

        if obj.modified & physical_interface_display_name != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.display_name)
                if err != nil {
                        return nil, err
                }
                msg["display_name"] = &value
        }

        if len(obj.physical_interface_refs) > 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.physical_interface_refs)
                if err != nil {
                        return nil, err
                }
                msg["physical_interface_refs"] = &value
        }

        return json.Marshal(msg)
}

func (obj *PhysicalInterface) UnmarshalJSON(body []byte) error {
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
                                obj.valid |= physical_interface_id_perms
                        }
                        break
                case "perms2":
                        err = json.Unmarshal(value, &obj.perms2)
                        if err == nil {
                                obj.valid |= physical_interface_perms2
                        }
                        break
                case "display_name":
                        err = json.Unmarshal(value, &obj.display_name)
                        if err == nil {
                                obj.valid |= physical_interface_display_name
                        }
                        break
                case "logical_interfaces":
                        err = json.Unmarshal(value, &obj.logical_interfaces)
                        if err == nil {
                                obj.valid |= physical_interface_logical_interfaces
                        }
                        break
                case "physical_interface_refs":
                        err = json.Unmarshal(value, &obj.physical_interface_refs)
                        if err == nil {
                                obj.valid |= physical_interface_physical_interface_refs
                        }
                        break
                case "virtual_machine_interface_back_refs":
                        err = json.Unmarshal(value, &obj.virtual_machine_interface_back_refs)
                        if err == nil {
                                obj.valid |= physical_interface_virtual_machine_interface_back_refs
                        }
                        break
                case "physical_interface_back_refs":
                        err = json.Unmarshal(value, &obj.physical_interface_back_refs)
                        if err == nil {
                                obj.valid |= physical_interface_physical_interface_back_refs
                        }
                        break
                case "service_appliance_back_refs": {
                        type ReferenceElement struct {
                                To []string
                                Uuid string
                                Href string
                                Attr ServiceApplianceInterfaceType
                        }
                        var array []ReferenceElement
                        err = json.Unmarshal(value, &array)
                        if err != nil {
                            break
                        }
                        obj.valid |= physical_interface_service_appliance_back_refs
                        obj.service_appliance_back_refs = make(contrail.ReferenceList, 0)
                        for _, element := range array {
                                ref := contrail.Reference {
                                        element.To,
                                        element.Uuid,
                                        element.Href,
                                        element.Attr,
                                }
                                obj.service_appliance_back_refs = append(obj.service_appliance_back_refs, ref)
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

func (obj *PhysicalInterface) UpdateObject() ([]byte, error) {
        msg := map[string]*json.RawMessage {
        }
        err := obj.MarshalId(msg)
        if err != nil {
                return nil, err
        }

        if obj.modified & physical_interface_id_perms != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.id_perms)
                if err != nil {
                        return nil, err
                }
                msg["id_perms"] = &value
        }

        if obj.modified & physical_interface_perms2 != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.perms2)
                if err != nil {
                        return nil, err
                }
                msg["perms2"] = &value
        }

        if obj.modified & physical_interface_display_name != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.display_name)
                if err != nil {
                        return nil, err
                }
                msg["display_name"] = &value
        }

        if obj.modified & physical_interface_physical_interface_refs != 0 {
                if len(obj.physical_interface_refs) == 0 {
                        var value json.RawMessage
                        value, err := json.Marshal(
                                          make([]contrail.Reference, 0))
                        if err != nil {
                                return nil, err
                        }
                        msg["physical_interface_refs"] = &value
                } else if !obj.hasReferenceBase("physical-interface") {
                        var value json.RawMessage
                        value, err := json.Marshal(&obj.physical_interface_refs)
                        if err != nil {
                                return nil, err
                        }
                        msg["physical_interface_refs"] = &value
                }
        }


        return json.Marshal(msg)
}

func (obj *PhysicalInterface) UpdateReferences() error {

        if (obj.modified & physical_interface_physical_interface_refs != 0) &&
           len(obj.physical_interface_refs) > 0 &&
           obj.hasReferenceBase("physical-interface") {
                err := obj.UpdateReference(
                        obj, "physical-interface",
                        obj.physical_interface_refs,
                        obj.baseMap["physical-interface"])
                if err != nil {
                        return err
                }
        }

        return nil
}

func PhysicalInterfaceByName(c contrail.ApiClient, fqn string) (*PhysicalInterface, error) {
    obj, err := c.FindByName("physical-interface", fqn)
    if err != nil {
        return nil, err
    }
    return obj.(*PhysicalInterface), nil
}

func PhysicalInterfaceByUuid(c contrail.ApiClient, uuid string) (*PhysicalInterface, error) {
    obj, err := c.FindByUuid("physical-interface", uuid)
    if err != nil {
        return nil, err
    }
    return obj.(*PhysicalInterface), nil
}
