//
// Automatically generated. DO NOT EDIT.
//

package types

import (
        "encoding/json"

        "github.com/michaelhenkel/contrail-go-api"
)

const (
	floating_ip_floating_ip_address uint64 = 1 << iota
	floating_ip_floating_ip_is_virtual_ip
	floating_ip_floating_ip_fixed_ip_address
	floating_ip_floating_ip_address_family
	floating_ip_id_perms
	floating_ip_perms2
	floating_ip_display_name
	floating_ip_project_refs
	floating_ip_virtual_machine_interface_refs
	floating_ip_customer_attachment_back_refs
)

type FloatingIp struct {
        contrail.ObjectBase
	floating_ip_address string
	floating_ip_is_virtual_ip bool
	floating_ip_fixed_ip_address string
	floating_ip_address_family string
	id_perms IdPermsType
	perms2 PermType2
	display_name string
	project_refs contrail.ReferenceList
	virtual_machine_interface_refs contrail.ReferenceList
	customer_attachment_back_refs contrail.ReferenceList
        valid uint64
        modified uint64
        baseMap map[string]contrail.ReferenceList
}

func (obj *FloatingIp) GetType() string {
        return "floating-ip"
}

func (obj *FloatingIp) GetDefaultParent() []string {
        name := []string{"default-domain", "default-project", "default-virtual-network", "default-floating-ip-pool"}
        return name
}

func (obj *FloatingIp) GetDefaultParentType() string {
        return "floating-ip-pool"
}

func (obj *FloatingIp) SetName(name string) {
        obj.VSetName(obj, name)
}

func (obj *FloatingIp) SetParent(parent contrail.IObject) {
        obj.VSetParent(obj, parent)
}

func (obj *FloatingIp) storeReferenceBase(
        name string, refList contrail.ReferenceList) {
        if obj.baseMap == nil {
                obj.baseMap = make(map[string]contrail.ReferenceList)
        }
        refCopy := make(contrail.ReferenceList, len(refList))
        copy(refCopy, refList)
        obj.baseMap[name] = refCopy
}

func (obj *FloatingIp) hasReferenceBase(name string) bool {
        if obj.baseMap == nil {
                return false
        }
        _, exists := obj.baseMap[name]
        return exists
}

func (obj *FloatingIp) UpdateDone() {
        obj.modified = 0
        obj.baseMap = nil
}


func (obj *FloatingIp) GetFloatingIpAddress() string {
        return obj.floating_ip_address
}

func (obj *FloatingIp) SetFloatingIpAddress(value string) {
        obj.floating_ip_address = value
        obj.modified |= floating_ip_floating_ip_address
}

func (obj *FloatingIp) GetFloatingIpIsVirtualIp() bool {
        return obj.floating_ip_is_virtual_ip
}

func (obj *FloatingIp) SetFloatingIpIsVirtualIp(value bool) {
        obj.floating_ip_is_virtual_ip = value
        obj.modified |= floating_ip_floating_ip_is_virtual_ip
}

func (obj *FloatingIp) GetFloatingIpFixedIpAddress() string {
        return obj.floating_ip_fixed_ip_address
}

func (obj *FloatingIp) SetFloatingIpFixedIpAddress(value string) {
        obj.floating_ip_fixed_ip_address = value
        obj.modified |= floating_ip_floating_ip_fixed_ip_address
}

func (obj *FloatingIp) GetFloatingIpAddressFamily() string {
        return obj.floating_ip_address_family
}

func (obj *FloatingIp) SetFloatingIpAddressFamily(value string) {
        obj.floating_ip_address_family = value
        obj.modified |= floating_ip_floating_ip_address_family
}

func (obj *FloatingIp) GetIdPerms() IdPermsType {
        return obj.id_perms
}

func (obj *FloatingIp) SetIdPerms(value *IdPermsType) {
        obj.id_perms = *value
        obj.modified |= floating_ip_id_perms
}

func (obj *FloatingIp) GetPerms2() PermType2 {
        return obj.perms2
}

func (obj *FloatingIp) SetPerms2(value *PermType2) {
        obj.perms2 = *value
        obj.modified |= floating_ip_perms2
}

func (obj *FloatingIp) GetDisplayName() string {
        return obj.display_name
}

func (obj *FloatingIp) SetDisplayName(value string) {
        obj.display_name = value
        obj.modified |= floating_ip_display_name
}

func (obj *FloatingIp) readProjectRefs() error {
        if !obj.IsTransient() &&
                (obj.valid & floating_ip_project_refs == 0) {
                err := obj.GetField(obj, "project_refs")
                if err != nil {
                        return err
                }
        }
        return nil
}

func (obj *FloatingIp) GetProjectRefs() (
        contrail.ReferenceList, error) {
        err := obj.readProjectRefs()
        if err != nil {
                return nil, err
        }
        return obj.project_refs, nil
}

func (obj *FloatingIp) AddProject(
        rhs *Project) error {
        err := obj.readProjectRefs()
        if err != nil {
                return err
        }

        if obj.modified & floating_ip_project_refs == 0 {
                obj.storeReferenceBase("project", obj.project_refs)
        }

        ref := contrail.Reference {
                rhs.GetFQName(), rhs.GetUuid(), rhs.GetHref(), nil}
        obj.project_refs = append(obj.project_refs, ref)
        obj.modified |= floating_ip_project_refs
        return nil
}

func (obj *FloatingIp) DeleteProject(uuid string) error {
        err := obj.readProjectRefs()
        if err != nil {
                return err
        }

        if obj.modified & floating_ip_project_refs == 0 {
                obj.storeReferenceBase("project", obj.project_refs)
        }

        for i, ref := range obj.project_refs {
                if ref.Uuid == uuid {
                        obj.project_refs = append(
                                obj.project_refs[:i],
                                obj.project_refs[i+1:]...)
                        break
                }
        }
        obj.modified |= floating_ip_project_refs
        return nil
}

func (obj *FloatingIp) ClearProject() {
        if (obj.valid & floating_ip_project_refs != 0) &&
           (obj.modified & floating_ip_project_refs == 0) {
                obj.storeReferenceBase("project", obj.project_refs)
        }
        obj.project_refs = make([]contrail.Reference, 0)
        obj.valid |= floating_ip_project_refs
        obj.modified |= floating_ip_project_refs
}

func (obj *FloatingIp) SetProjectList(
        refList []contrail.ReferencePair) {
        obj.ClearProject()
        obj.project_refs = make([]contrail.Reference, len(refList))
        for i, pair := range refList {
                obj.project_refs[i] = contrail.Reference {
                        pair.Object.GetFQName(),
                        pair.Object.GetUuid(),
                        pair.Object.GetHref(),
                        pair.Attribute,
                }
        }
}


func (obj *FloatingIp) readVirtualMachineInterfaceRefs() error {
        if !obj.IsTransient() &&
                (obj.valid & floating_ip_virtual_machine_interface_refs == 0) {
                err := obj.GetField(obj, "virtual_machine_interface_refs")
                if err != nil {
                        return err
                }
        }
        return nil
}

func (obj *FloatingIp) GetVirtualMachineInterfaceRefs() (
        contrail.ReferenceList, error) {
        err := obj.readVirtualMachineInterfaceRefs()
        if err != nil {
                return nil, err
        }
        return obj.virtual_machine_interface_refs, nil
}

func (obj *FloatingIp) AddVirtualMachineInterface(
        rhs *VirtualMachineInterface) error {
        err := obj.readVirtualMachineInterfaceRefs()
        if err != nil {
                return err
        }

        if obj.modified & floating_ip_virtual_machine_interface_refs == 0 {
                obj.storeReferenceBase("virtual-machine-interface", obj.virtual_machine_interface_refs)
        }

        ref := contrail.Reference {
                rhs.GetFQName(), rhs.GetUuid(), rhs.GetHref(), nil}
        obj.virtual_machine_interface_refs = append(obj.virtual_machine_interface_refs, ref)
        obj.modified |= floating_ip_virtual_machine_interface_refs
        return nil
}

func (obj *FloatingIp) DeleteVirtualMachineInterface(uuid string) error {
        err := obj.readVirtualMachineInterfaceRefs()
        if err != nil {
                return err
        }

        if obj.modified & floating_ip_virtual_machine_interface_refs == 0 {
                obj.storeReferenceBase("virtual-machine-interface", obj.virtual_machine_interface_refs)
        }

        for i, ref := range obj.virtual_machine_interface_refs {
                if ref.Uuid == uuid {
                        obj.virtual_machine_interface_refs = append(
                                obj.virtual_machine_interface_refs[:i],
                                obj.virtual_machine_interface_refs[i+1:]...)
                        break
                }
        }
        obj.modified |= floating_ip_virtual_machine_interface_refs
        return nil
}

func (obj *FloatingIp) ClearVirtualMachineInterface() {
        if (obj.valid & floating_ip_virtual_machine_interface_refs != 0) &&
           (obj.modified & floating_ip_virtual_machine_interface_refs == 0) {
                obj.storeReferenceBase("virtual-machine-interface", obj.virtual_machine_interface_refs)
        }
        obj.virtual_machine_interface_refs = make([]contrail.Reference, 0)
        obj.valid |= floating_ip_virtual_machine_interface_refs
        obj.modified |= floating_ip_virtual_machine_interface_refs
}

func (obj *FloatingIp) SetVirtualMachineInterfaceList(
        refList []contrail.ReferencePair) {
        obj.ClearVirtualMachineInterface()
        obj.virtual_machine_interface_refs = make([]contrail.Reference, len(refList))
        for i, pair := range refList {
                obj.virtual_machine_interface_refs[i] = contrail.Reference {
                        pair.Object.GetFQName(),
                        pair.Object.GetUuid(),
                        pair.Object.GetHref(),
                        pair.Attribute,
                }
        }
}


func (obj *FloatingIp) readCustomerAttachmentBackRefs() error {
        if !obj.IsTransient() &&
                (obj.valid & floating_ip_customer_attachment_back_refs == 0) {
                err := obj.GetField(obj, "customer_attachment_back_refs")
                if err != nil {
                        return err
                }
        }
        return nil
}

func (obj *FloatingIp) GetCustomerAttachmentBackRefs() (
        contrail.ReferenceList, error) {
        err := obj.readCustomerAttachmentBackRefs()
        if err != nil {
                return nil, err
        }
        return obj.customer_attachment_back_refs, nil
}

func (obj *FloatingIp) MarshalJSON() ([]byte, error) {
        msg := map[string]*json.RawMessage {
        }
        err := obj.MarshalCommon(msg)
        if err != nil {
                return nil, err
        }

        if obj.modified & floating_ip_floating_ip_address != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.floating_ip_address)
                if err != nil {
                        return nil, err
                }
                msg["floating_ip_address"] = &value
        }

        if obj.modified & floating_ip_floating_ip_is_virtual_ip != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.floating_ip_is_virtual_ip)
                if err != nil {
                        return nil, err
                }
                msg["floating_ip_is_virtual_ip"] = &value
        }

        if obj.modified & floating_ip_floating_ip_fixed_ip_address != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.floating_ip_fixed_ip_address)
                if err != nil {
                        return nil, err
                }
                msg["floating_ip_fixed_ip_address"] = &value
        }

        if obj.modified & floating_ip_floating_ip_address_family != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.floating_ip_address_family)
                if err != nil {
                        return nil, err
                }
                msg["floating_ip_address_family"] = &value
        }

        if obj.modified & floating_ip_id_perms != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.id_perms)
                if err != nil {
                        return nil, err
                }
                msg["id_perms"] = &value
        }

        if obj.modified & floating_ip_perms2 != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.perms2)
                if err != nil {
                        return nil, err
                }
                msg["perms2"] = &value
        }

        if obj.modified & floating_ip_display_name != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.display_name)
                if err != nil {
                        return nil, err
                }
                msg["display_name"] = &value
        }

        if len(obj.project_refs) > 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.project_refs)
                if err != nil {
                        return nil, err
                }
                msg["project_refs"] = &value
        }

        if len(obj.virtual_machine_interface_refs) > 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.virtual_machine_interface_refs)
                if err != nil {
                        return nil, err
                }
                msg["virtual_machine_interface_refs"] = &value
        }

        return json.Marshal(msg)
}

func (obj *FloatingIp) UnmarshalJSON(body []byte) error {
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
                case "floating_ip_address":
                        err = json.Unmarshal(value, &obj.floating_ip_address)
                        if err == nil {
                                obj.valid |= floating_ip_floating_ip_address
                        }
                        break
                case "floating_ip_is_virtual_ip":
                        err = json.Unmarshal(value, &obj.floating_ip_is_virtual_ip)
                        if err == nil {
                                obj.valid |= floating_ip_floating_ip_is_virtual_ip
                        }
                        break
                case "floating_ip_fixed_ip_address":
                        err = json.Unmarshal(value, &obj.floating_ip_fixed_ip_address)
                        if err == nil {
                                obj.valid |= floating_ip_floating_ip_fixed_ip_address
                        }
                        break
                case "floating_ip_address_family":
                        err = json.Unmarshal(value, &obj.floating_ip_address_family)
                        if err == nil {
                                obj.valid |= floating_ip_floating_ip_address_family
                        }
                        break
                case "id_perms":
                        err = json.Unmarshal(value, &obj.id_perms)
                        if err == nil {
                                obj.valid |= floating_ip_id_perms
                        }
                        break
                case "perms2":
                        err = json.Unmarshal(value, &obj.perms2)
                        if err == nil {
                                obj.valid |= floating_ip_perms2
                        }
                        break
                case "display_name":
                        err = json.Unmarshal(value, &obj.display_name)
                        if err == nil {
                                obj.valid |= floating_ip_display_name
                        }
                        break
                case "project_refs":
                        err = json.Unmarshal(value, &obj.project_refs)
                        if err == nil {
                                obj.valid |= floating_ip_project_refs
                        }
                        break
                case "virtual_machine_interface_refs":
                        err = json.Unmarshal(value, &obj.virtual_machine_interface_refs)
                        if err == nil {
                                obj.valid |= floating_ip_virtual_machine_interface_refs
                        }
                        break
                case "customer_attachment_back_refs":
                        err = json.Unmarshal(value, &obj.customer_attachment_back_refs)
                        if err == nil {
                                obj.valid |= floating_ip_customer_attachment_back_refs
                        }
                        break
                }
                if err != nil {
                        return err
                }
        }
        return nil
}

func (obj *FloatingIp) UpdateObject() ([]byte, error) {
        msg := map[string]*json.RawMessage {
        }
        err := obj.MarshalId(msg)
        if err != nil {
                return nil, err
        }

        if obj.modified & floating_ip_floating_ip_address != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.floating_ip_address)
                if err != nil {
                        return nil, err
                }
                msg["floating_ip_address"] = &value
        }

        if obj.modified & floating_ip_floating_ip_is_virtual_ip != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.floating_ip_is_virtual_ip)
                if err != nil {
                        return nil, err
                }
                msg["floating_ip_is_virtual_ip"] = &value
        }

        if obj.modified & floating_ip_floating_ip_fixed_ip_address != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.floating_ip_fixed_ip_address)
                if err != nil {
                        return nil, err
                }
                msg["floating_ip_fixed_ip_address"] = &value
        }

        if obj.modified & floating_ip_floating_ip_address_family != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.floating_ip_address_family)
                if err != nil {
                        return nil, err
                }
                msg["floating_ip_address_family"] = &value
        }

        if obj.modified & floating_ip_id_perms != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.id_perms)
                if err != nil {
                        return nil, err
                }
                msg["id_perms"] = &value
        }

        if obj.modified & floating_ip_perms2 != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.perms2)
                if err != nil {
                        return nil, err
                }
                msg["perms2"] = &value
        }

        if obj.modified & floating_ip_display_name != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.display_name)
                if err != nil {
                        return nil, err
                }
                msg["display_name"] = &value
        }

        if obj.modified & floating_ip_project_refs != 0 {
                if len(obj.project_refs) == 0 {
                        var value json.RawMessage
                        value, err := json.Marshal(
                                          make([]contrail.Reference, 0))
                        if err != nil {
                                return nil, err
                        }
                        msg["project_refs"] = &value
                } else if !obj.hasReferenceBase("project") {
                        var value json.RawMessage
                        value, err := json.Marshal(&obj.project_refs)
                        if err != nil {
                                return nil, err
                        }
                        msg["project_refs"] = &value
                }
        }


        if obj.modified & floating_ip_virtual_machine_interface_refs != 0 {
                if len(obj.virtual_machine_interface_refs) == 0 {
                        var value json.RawMessage
                        value, err := json.Marshal(
                                          make([]contrail.Reference, 0))
                        if err != nil {
                                return nil, err
                        }
                        msg["virtual_machine_interface_refs"] = &value
                } else if !obj.hasReferenceBase("virtual-machine-interface") {
                        var value json.RawMessage
                        value, err := json.Marshal(&obj.virtual_machine_interface_refs)
                        if err != nil {
                                return nil, err
                        }
                        msg["virtual_machine_interface_refs"] = &value
                }
        }


        return json.Marshal(msg)
}

func (obj *FloatingIp) UpdateReferences() error {

        if (obj.modified & floating_ip_project_refs != 0) &&
           len(obj.project_refs) > 0 &&
           obj.hasReferenceBase("project") {
                err := obj.UpdateReference(
                        obj, "project",
                        obj.project_refs,
                        obj.baseMap["project"])
                if err != nil {
                        return err
                }
        }

        if (obj.modified & floating_ip_virtual_machine_interface_refs != 0) &&
           len(obj.virtual_machine_interface_refs) > 0 &&
           obj.hasReferenceBase("virtual-machine-interface") {
                err := obj.UpdateReference(
                        obj, "virtual-machine-interface",
                        obj.virtual_machine_interface_refs,
                        obj.baseMap["virtual-machine-interface"])
                if err != nil {
                        return err
                }
        }

        return nil
}

func FloatingIpByName(c contrail.ApiClient, fqn string) (*FloatingIp, error) {
    obj, err := c.FindByName("floating-ip", fqn)
    if err != nil {
        return nil, err
    }
    return obj.(*FloatingIp), nil
}

func FloatingIpByUuid(c contrail.ApiClient, uuid string) (*FloatingIp, error) {
    obj, err := c.FindByUuid("floating-ip", uuid)
    if err != nil {
        return nil, err
    }
    return obj.(*FloatingIp), nil
}
