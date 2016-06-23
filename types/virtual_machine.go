//
// Automatically generated. DO NOT EDIT.
//

package types

import (
        "encoding/json"

        "github.com/michaelhenkel/contrail-go-api"
)

const (
	virtual_machine_id_perms uint64 = 1 << iota
	virtual_machine_perms2
	virtual_machine_display_name
	virtual_machine_virtual_machine_interfaces
	virtual_machine_service_instance_refs
	virtual_machine_virtual_machine_interface_back_refs
	virtual_machine_virtual_router_back_refs
)

type VirtualMachine struct {
        contrail.ObjectBase
	id_perms IdPermsType
	perms2 PermType2
	display_name string
	virtual_machine_interfaces contrail.ReferenceList
	service_instance_refs contrail.ReferenceList
	virtual_machine_interface_back_refs contrail.ReferenceList
	virtual_router_back_refs contrail.ReferenceList
        valid uint64
        modified uint64
        baseMap map[string]contrail.ReferenceList
}

func (obj *VirtualMachine) GetType() string {
        return "virtual-machine"
}

func (obj *VirtualMachine) GetDefaultParent() []string {
        name := []string{}
        return name
}

func (obj *VirtualMachine) GetDefaultParentType() string {
        return ""
}

func (obj *VirtualMachine) SetName(name string) {
        obj.VSetName(obj, name)
}

func (obj *VirtualMachine) SetParent(parent contrail.IObject) {
        obj.VSetParent(obj, parent)
}

func (obj *VirtualMachine) storeReferenceBase(
        name string, refList contrail.ReferenceList) {
        if obj.baseMap == nil {
                obj.baseMap = make(map[string]contrail.ReferenceList)
        }
        refCopy := make(contrail.ReferenceList, len(refList))
        copy(refCopy, refList)
        obj.baseMap[name] = refCopy
}

func (obj *VirtualMachine) hasReferenceBase(name string) bool {
        if obj.baseMap == nil {
                return false
        }
        _, exists := obj.baseMap[name]
        return exists
}

func (obj *VirtualMachine) UpdateDone() {
        obj.modified = 0
        obj.baseMap = nil
}


func (obj *VirtualMachine) GetIdPerms() IdPermsType {
        return obj.id_perms
}

func (obj *VirtualMachine) SetIdPerms(value *IdPermsType) {
        obj.id_perms = *value
        obj.modified |= virtual_machine_id_perms
}

func (obj *VirtualMachine) GetPerms2() PermType2 {
        return obj.perms2
}

func (obj *VirtualMachine) SetPerms2(value *PermType2) {
        obj.perms2 = *value
        obj.modified |= virtual_machine_perms2
}

func (obj *VirtualMachine) GetDisplayName() string {
        return obj.display_name
}

func (obj *VirtualMachine) SetDisplayName(value string) {
        obj.display_name = value
        obj.modified |= virtual_machine_display_name
}

func (obj *VirtualMachine) readVirtualMachineInterfaces() error {
        if !obj.IsTransient() &&
                (obj.valid & virtual_machine_virtual_machine_interfaces == 0) {
                err := obj.GetField(obj, "virtual_machine_interfaces")
                if err != nil {
                        return err
                }
        }
        return nil
}

func (obj *VirtualMachine) GetVirtualMachineInterfaces() (
        contrail.ReferenceList, error) {
        err := obj.readVirtualMachineInterfaces()
        if err != nil {
                return nil, err
        }
        return obj.virtual_machine_interfaces, nil
}

func (obj *VirtualMachine) readServiceInstanceRefs() error {
        if !obj.IsTransient() &&
                (obj.valid & virtual_machine_service_instance_refs == 0) {
                err := obj.GetField(obj, "service_instance_refs")
                if err != nil {
                        return err
                }
        }
        return nil
}

func (obj *VirtualMachine) GetServiceInstanceRefs() (
        contrail.ReferenceList, error) {
        err := obj.readServiceInstanceRefs()
        if err != nil {
                return nil, err
        }
        return obj.service_instance_refs, nil
}

func (obj *VirtualMachine) AddServiceInstance(
        rhs *ServiceInstance) error {
        err := obj.readServiceInstanceRefs()
        if err != nil {
                return err
        }

        if obj.modified & virtual_machine_service_instance_refs == 0 {
                obj.storeReferenceBase("service-instance", obj.service_instance_refs)
        }

        ref := contrail.Reference {
                rhs.GetFQName(), rhs.GetUuid(), rhs.GetHref(), nil}
        obj.service_instance_refs = append(obj.service_instance_refs, ref)
        obj.modified |= virtual_machine_service_instance_refs
        return nil
}

func (obj *VirtualMachine) DeleteServiceInstance(uuid string) error {
        err := obj.readServiceInstanceRefs()
        if err != nil {
                return err
        }

        if obj.modified & virtual_machine_service_instance_refs == 0 {
                obj.storeReferenceBase("service-instance", obj.service_instance_refs)
        }

        for i, ref := range obj.service_instance_refs {
                if ref.Uuid == uuid {
                        obj.service_instance_refs = append(
                                obj.service_instance_refs[:i],
                                obj.service_instance_refs[i+1:]...)
                        break
                }
        }
        obj.modified |= virtual_machine_service_instance_refs
        return nil
}

func (obj *VirtualMachine) ClearServiceInstance() {
        if (obj.valid & virtual_machine_service_instance_refs != 0) &&
           (obj.modified & virtual_machine_service_instance_refs == 0) {
                obj.storeReferenceBase("service-instance", obj.service_instance_refs)
        }
        obj.service_instance_refs = make([]contrail.Reference, 0)
        obj.valid |= virtual_machine_service_instance_refs
        obj.modified |= virtual_machine_service_instance_refs
}

func (obj *VirtualMachine) SetServiceInstanceList(
        refList []contrail.ReferencePair) {
        obj.ClearServiceInstance()
        obj.service_instance_refs = make([]contrail.Reference, len(refList))
        for i, pair := range refList {
                obj.service_instance_refs[i] = contrail.Reference {
                        pair.Object.GetFQName(),
                        pair.Object.GetUuid(),
                        pair.Object.GetHref(),
                        pair.Attribute,
                }
        }
}


func (obj *VirtualMachine) readVirtualMachineInterfaceBackRefs() error {
        if !obj.IsTransient() &&
                (obj.valid & virtual_machine_virtual_machine_interface_back_refs == 0) {
                err := obj.GetField(obj, "virtual_machine_interface_back_refs")
                if err != nil {
                        return err
                }
        }
        return nil
}

func (obj *VirtualMachine) GetVirtualMachineInterfaceBackRefs() (
        contrail.ReferenceList, error) {
        err := obj.readVirtualMachineInterfaceBackRefs()
        if err != nil {
                return nil, err
        }
        return obj.virtual_machine_interface_back_refs, nil
}

func (obj *VirtualMachine) readVirtualRouterBackRefs() error {
        if !obj.IsTransient() &&
                (obj.valid & virtual_machine_virtual_router_back_refs == 0) {
                err := obj.GetField(obj, "virtual_router_back_refs")
                if err != nil {
                        return err
                }
        }
        return nil
}

func (obj *VirtualMachine) GetVirtualRouterBackRefs() (
        contrail.ReferenceList, error) {
        err := obj.readVirtualRouterBackRefs()
        if err != nil {
                return nil, err
        }
        return obj.virtual_router_back_refs, nil
}

func (obj *VirtualMachine) MarshalJSON() ([]byte, error) {
        msg := map[string]*json.RawMessage {
        }
        err := obj.MarshalCommon(msg)
        if err != nil {
                return nil, err
        }

        if obj.modified & virtual_machine_id_perms != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.id_perms)
                if err != nil {
                        return nil, err
                }
                msg["id_perms"] = &value
        }

        if obj.modified & virtual_machine_perms2 != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.perms2)
                if err != nil {
                        return nil, err
                }
                msg["perms2"] = &value
        }

        if obj.modified & virtual_machine_display_name != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.display_name)
                if err != nil {
                        return nil, err
                }
                msg["display_name"] = &value
        }

        if len(obj.service_instance_refs) > 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.service_instance_refs)
                if err != nil {
                        return nil, err
                }
                msg["service_instance_refs"] = &value
        }

        return json.Marshal(msg)
}

func (obj *VirtualMachine) UnmarshalJSON(body []byte) error {
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
                                obj.valid |= virtual_machine_id_perms
                        }
                        break
                case "perms2":
                        err = json.Unmarshal(value, &obj.perms2)
                        if err == nil {
                                obj.valid |= virtual_machine_perms2
                        }
                        break
                case "display_name":
                        err = json.Unmarshal(value, &obj.display_name)
                        if err == nil {
                                obj.valid |= virtual_machine_display_name
                        }
                        break
                case "virtual_machine_interfaces":
                        err = json.Unmarshal(value, &obj.virtual_machine_interfaces)
                        if err == nil {
                                obj.valid |= virtual_machine_virtual_machine_interfaces
                        }
                        break
                case "service_instance_refs":
                        err = json.Unmarshal(value, &obj.service_instance_refs)
                        if err == nil {
                                obj.valid |= virtual_machine_service_instance_refs
                        }
                        break
                case "virtual_machine_interface_back_refs":
                        err = json.Unmarshal(value, &obj.virtual_machine_interface_back_refs)
                        if err == nil {
                                obj.valid |= virtual_machine_virtual_machine_interface_back_refs
                        }
                        break
                case "virtual_router_back_refs":
                        err = json.Unmarshal(value, &obj.virtual_router_back_refs)
                        if err == nil {
                                obj.valid |= virtual_machine_virtual_router_back_refs
                        }
                        break
                }
                if err != nil {
                        return err
                }
        }
        return nil
}

func (obj *VirtualMachine) UpdateObject() ([]byte, error) {
        msg := map[string]*json.RawMessage {
        }
        err := obj.MarshalId(msg)
        if err != nil {
                return nil, err
        }

        if obj.modified & virtual_machine_id_perms != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.id_perms)
                if err != nil {
                        return nil, err
                }
                msg["id_perms"] = &value
        }

        if obj.modified & virtual_machine_perms2 != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.perms2)
                if err != nil {
                        return nil, err
                }
                msg["perms2"] = &value
        }

        if obj.modified & virtual_machine_display_name != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.display_name)
                if err != nil {
                        return nil, err
                }
                msg["display_name"] = &value
        }

        if obj.modified & virtual_machine_service_instance_refs != 0 {
                if len(obj.service_instance_refs) == 0 {
                        var value json.RawMessage
                        value, err := json.Marshal(
                                          make([]contrail.Reference, 0))
                        if err != nil {
                                return nil, err
                        }
                        msg["service_instance_refs"] = &value
                } else if !obj.hasReferenceBase("service-instance") {
                        var value json.RawMessage
                        value, err := json.Marshal(&obj.service_instance_refs)
                        if err != nil {
                                return nil, err
                        }
                        msg["service_instance_refs"] = &value
                }
        }


        return json.Marshal(msg)
}

func (obj *VirtualMachine) UpdateReferences() error {

        if (obj.modified & virtual_machine_service_instance_refs != 0) &&
           len(obj.service_instance_refs) > 0 &&
           obj.hasReferenceBase("service-instance") {
                err := obj.UpdateReference(
                        obj, "service-instance",
                        obj.service_instance_refs,
                        obj.baseMap["service-instance"])
                if err != nil {
                        return err
                }
        }

        return nil
}

func VirtualMachineByName(c contrail.ApiClient, fqn string) (*VirtualMachine, error) {
    obj, err := c.FindByName("virtual-machine", fqn)
    if err != nil {
        return nil, err
    }
    return obj.(*VirtualMachine), nil
}

func VirtualMachineByUuid(c contrail.ApiClient, uuid string) (*VirtualMachine, error) {
    obj, err := c.FindByUuid("virtual-machine", uuid)
    if err != nil {
        return nil, err
    }
    return obj.(*VirtualMachine), nil
}
