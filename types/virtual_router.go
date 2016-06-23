//
// Automatically generated. DO NOT EDIT.
//

package types

import (
        "encoding/json"

        "github.com/michaelhenkel/contrail-go-api"
)

const (
	virtual_router_virtual_router_type uint64 = 1 << iota
	virtual_router_virtual_router_dpdk_enabled
	virtual_router_virtual_router_ip_address
	virtual_router_id_perms
	virtual_router_perms2
	virtual_router_display_name
	virtual_router_virtual_machine_refs
	virtual_router_physical_router_back_refs
	virtual_router_provider_attachment_back_refs
)

type VirtualRouter struct {
        contrail.ObjectBase
	virtual_router_type string
	virtual_router_dpdk_enabled bool
	virtual_router_ip_address string
	id_perms IdPermsType
	perms2 PermType2
	display_name string
	virtual_machine_refs contrail.ReferenceList
	physical_router_back_refs contrail.ReferenceList
	provider_attachment_back_refs contrail.ReferenceList
        valid uint64
        modified uint64
        baseMap map[string]contrail.ReferenceList
}

func (obj *VirtualRouter) GetType() string {
        return "virtual-router"
}

func (obj *VirtualRouter) GetDefaultParent() []string {
        name := []string{"default-global-system-config"}
        return name
}

func (obj *VirtualRouter) GetDefaultParentType() string {
        return "global-system-config"
}

func (obj *VirtualRouter) SetName(name string) {
        obj.VSetName(obj, name)
}

func (obj *VirtualRouter) SetParent(parent contrail.IObject) {
        obj.VSetParent(obj, parent)
}

func (obj *VirtualRouter) storeReferenceBase(
        name string, refList contrail.ReferenceList) {
        if obj.baseMap == nil {
                obj.baseMap = make(map[string]contrail.ReferenceList)
        }
        refCopy := make(contrail.ReferenceList, len(refList))
        copy(refCopy, refList)
        obj.baseMap[name] = refCopy
}

func (obj *VirtualRouter) hasReferenceBase(name string) bool {
        if obj.baseMap == nil {
                return false
        }
        _, exists := obj.baseMap[name]
        return exists
}

func (obj *VirtualRouter) UpdateDone() {
        obj.modified = 0
        obj.baseMap = nil
}


func (obj *VirtualRouter) GetVirtualRouterType() string {
        return obj.virtual_router_type
}

func (obj *VirtualRouter) SetVirtualRouterType(value string) {
        obj.virtual_router_type = value
        obj.modified |= virtual_router_virtual_router_type
}

func (obj *VirtualRouter) GetVirtualRouterDpdkEnabled() bool {
        return obj.virtual_router_dpdk_enabled
}

func (obj *VirtualRouter) SetVirtualRouterDpdkEnabled(value bool) {
        obj.virtual_router_dpdk_enabled = value
        obj.modified |= virtual_router_virtual_router_dpdk_enabled
}

func (obj *VirtualRouter) GetVirtualRouterIpAddress() string {
        return obj.virtual_router_ip_address
}

func (obj *VirtualRouter) SetVirtualRouterIpAddress(value string) {
        obj.virtual_router_ip_address = value
        obj.modified |= virtual_router_virtual_router_ip_address
}

func (obj *VirtualRouter) GetIdPerms() IdPermsType {
        return obj.id_perms
}

func (obj *VirtualRouter) SetIdPerms(value *IdPermsType) {
        obj.id_perms = *value
        obj.modified |= virtual_router_id_perms
}

func (obj *VirtualRouter) GetPerms2() PermType2 {
        return obj.perms2
}

func (obj *VirtualRouter) SetPerms2(value *PermType2) {
        obj.perms2 = *value
        obj.modified |= virtual_router_perms2
}

func (obj *VirtualRouter) GetDisplayName() string {
        return obj.display_name
}

func (obj *VirtualRouter) SetDisplayName(value string) {
        obj.display_name = value
        obj.modified |= virtual_router_display_name
}

func (obj *VirtualRouter) readVirtualMachineRefs() error {
        if !obj.IsTransient() &&
                (obj.valid & virtual_router_virtual_machine_refs == 0) {
                err := obj.GetField(obj, "virtual_machine_refs")
                if err != nil {
                        return err
                }
        }
        return nil
}

func (obj *VirtualRouter) GetVirtualMachineRefs() (
        contrail.ReferenceList, error) {
        err := obj.readVirtualMachineRefs()
        if err != nil {
                return nil, err
        }
        return obj.virtual_machine_refs, nil
}

func (obj *VirtualRouter) AddVirtualMachine(
        rhs *VirtualMachine) error {
        err := obj.readVirtualMachineRefs()
        if err != nil {
                return err
        }

        if obj.modified & virtual_router_virtual_machine_refs == 0 {
                obj.storeReferenceBase("virtual-machine", obj.virtual_machine_refs)
        }

        ref := contrail.Reference {
                rhs.GetFQName(), rhs.GetUuid(), rhs.GetHref(), nil}
        obj.virtual_machine_refs = append(obj.virtual_machine_refs, ref)
        obj.modified |= virtual_router_virtual_machine_refs
        return nil
}

func (obj *VirtualRouter) DeleteVirtualMachine(uuid string) error {
        err := obj.readVirtualMachineRefs()
        if err != nil {
                return err
        }

        if obj.modified & virtual_router_virtual_machine_refs == 0 {
                obj.storeReferenceBase("virtual-machine", obj.virtual_machine_refs)
        }

        for i, ref := range obj.virtual_machine_refs {
                if ref.Uuid == uuid {
                        obj.virtual_machine_refs = append(
                                obj.virtual_machine_refs[:i],
                                obj.virtual_machine_refs[i+1:]...)
                        break
                }
        }
        obj.modified |= virtual_router_virtual_machine_refs
        return nil
}

func (obj *VirtualRouter) ClearVirtualMachine() {
        if (obj.valid & virtual_router_virtual_machine_refs != 0) &&
           (obj.modified & virtual_router_virtual_machine_refs == 0) {
                obj.storeReferenceBase("virtual-machine", obj.virtual_machine_refs)
        }
        obj.virtual_machine_refs = make([]contrail.Reference, 0)
        obj.valid |= virtual_router_virtual_machine_refs
        obj.modified |= virtual_router_virtual_machine_refs
}

func (obj *VirtualRouter) SetVirtualMachineList(
        refList []contrail.ReferencePair) {
        obj.ClearVirtualMachine()
        obj.virtual_machine_refs = make([]contrail.Reference, len(refList))
        for i, pair := range refList {
                obj.virtual_machine_refs[i] = contrail.Reference {
                        pair.Object.GetFQName(),
                        pair.Object.GetUuid(),
                        pair.Object.GetHref(),
                        pair.Attribute,
                }
        }
}


func (obj *VirtualRouter) readPhysicalRouterBackRefs() error {
        if !obj.IsTransient() &&
                (obj.valid & virtual_router_physical_router_back_refs == 0) {
                err := obj.GetField(obj, "physical_router_back_refs")
                if err != nil {
                        return err
                }
        }
        return nil
}

func (obj *VirtualRouter) GetPhysicalRouterBackRefs() (
        contrail.ReferenceList, error) {
        err := obj.readPhysicalRouterBackRefs()
        if err != nil {
                return nil, err
        }
        return obj.physical_router_back_refs, nil
}

func (obj *VirtualRouter) readProviderAttachmentBackRefs() error {
        if !obj.IsTransient() &&
                (obj.valid & virtual_router_provider_attachment_back_refs == 0) {
                err := obj.GetField(obj, "provider_attachment_back_refs")
                if err != nil {
                        return err
                }
        }
        return nil
}

func (obj *VirtualRouter) GetProviderAttachmentBackRefs() (
        contrail.ReferenceList, error) {
        err := obj.readProviderAttachmentBackRefs()
        if err != nil {
                return nil, err
        }
        return obj.provider_attachment_back_refs, nil
}

func (obj *VirtualRouter) MarshalJSON() ([]byte, error) {
        msg := map[string]*json.RawMessage {
        }
        err := obj.MarshalCommon(msg)
        if err != nil {
                return nil, err
        }

        if obj.modified & virtual_router_virtual_router_type != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.virtual_router_type)
                if err != nil {
                        return nil, err
                }
                msg["virtual_router_type"] = &value
        }

        if obj.modified & virtual_router_virtual_router_dpdk_enabled != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.virtual_router_dpdk_enabled)
                if err != nil {
                        return nil, err
                }
                msg["virtual_router_dpdk_enabled"] = &value
        }

        if obj.modified & virtual_router_virtual_router_ip_address != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.virtual_router_ip_address)
                if err != nil {
                        return nil, err
                }
                msg["virtual_router_ip_address"] = &value
        }

        if obj.modified & virtual_router_id_perms != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.id_perms)
                if err != nil {
                        return nil, err
                }
                msg["id_perms"] = &value
        }

        if obj.modified & virtual_router_perms2 != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.perms2)
                if err != nil {
                        return nil, err
                }
                msg["perms2"] = &value
        }

        if obj.modified & virtual_router_display_name != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.display_name)
                if err != nil {
                        return nil, err
                }
                msg["display_name"] = &value
        }

        if len(obj.virtual_machine_refs) > 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.virtual_machine_refs)
                if err != nil {
                        return nil, err
                }
                msg["virtual_machine_refs"] = &value
        }

        return json.Marshal(msg)
}

func (obj *VirtualRouter) UnmarshalJSON(body []byte) error {
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
                case "virtual_router_type":
                        err = json.Unmarshal(value, &obj.virtual_router_type)
                        if err == nil {
                                obj.valid |= virtual_router_virtual_router_type
                        }
                        break
                case "virtual_router_dpdk_enabled":
                        err = json.Unmarshal(value, &obj.virtual_router_dpdk_enabled)
                        if err == nil {
                                obj.valid |= virtual_router_virtual_router_dpdk_enabled
                        }
                        break
                case "virtual_router_ip_address":
                        err = json.Unmarshal(value, &obj.virtual_router_ip_address)
                        if err == nil {
                                obj.valid |= virtual_router_virtual_router_ip_address
                        }
                        break
                case "id_perms":
                        err = json.Unmarshal(value, &obj.id_perms)
                        if err == nil {
                                obj.valid |= virtual_router_id_perms
                        }
                        break
                case "perms2":
                        err = json.Unmarshal(value, &obj.perms2)
                        if err == nil {
                                obj.valid |= virtual_router_perms2
                        }
                        break
                case "display_name":
                        err = json.Unmarshal(value, &obj.display_name)
                        if err == nil {
                                obj.valid |= virtual_router_display_name
                        }
                        break
                case "virtual_machine_refs":
                        err = json.Unmarshal(value, &obj.virtual_machine_refs)
                        if err == nil {
                                obj.valid |= virtual_router_virtual_machine_refs
                        }
                        break
                case "physical_router_back_refs":
                        err = json.Unmarshal(value, &obj.physical_router_back_refs)
                        if err == nil {
                                obj.valid |= virtual_router_physical_router_back_refs
                        }
                        break
                case "provider_attachment_back_refs":
                        err = json.Unmarshal(value, &obj.provider_attachment_back_refs)
                        if err == nil {
                                obj.valid |= virtual_router_provider_attachment_back_refs
                        }
                        break
                }
                if err != nil {
                        return err
                }
        }
        return nil
}

func (obj *VirtualRouter) UpdateObject() ([]byte, error) {
        msg := map[string]*json.RawMessage {
        }
        err := obj.MarshalId(msg)
        if err != nil {
                return nil, err
        }

        if obj.modified & virtual_router_virtual_router_type != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.virtual_router_type)
                if err != nil {
                        return nil, err
                }
                msg["virtual_router_type"] = &value
        }

        if obj.modified & virtual_router_virtual_router_dpdk_enabled != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.virtual_router_dpdk_enabled)
                if err != nil {
                        return nil, err
                }
                msg["virtual_router_dpdk_enabled"] = &value
        }

        if obj.modified & virtual_router_virtual_router_ip_address != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.virtual_router_ip_address)
                if err != nil {
                        return nil, err
                }
                msg["virtual_router_ip_address"] = &value
        }

        if obj.modified & virtual_router_id_perms != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.id_perms)
                if err != nil {
                        return nil, err
                }
                msg["id_perms"] = &value
        }

        if obj.modified & virtual_router_perms2 != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.perms2)
                if err != nil {
                        return nil, err
                }
                msg["perms2"] = &value
        }

        if obj.modified & virtual_router_display_name != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.display_name)
                if err != nil {
                        return nil, err
                }
                msg["display_name"] = &value
        }

        if obj.modified & virtual_router_virtual_machine_refs != 0 {
                if len(obj.virtual_machine_refs) == 0 {
                        var value json.RawMessage
                        value, err := json.Marshal(
                                          make([]contrail.Reference, 0))
                        if err != nil {
                                return nil, err
                        }
                        msg["virtual_machine_refs"] = &value
                } else if !obj.hasReferenceBase("virtual-machine") {
                        var value json.RawMessage
                        value, err := json.Marshal(&obj.virtual_machine_refs)
                        if err != nil {
                                return nil, err
                        }
                        msg["virtual_machine_refs"] = &value
                }
        }


        return json.Marshal(msg)
}

func (obj *VirtualRouter) UpdateReferences() error {

        if (obj.modified & virtual_router_virtual_machine_refs != 0) &&
           len(obj.virtual_machine_refs) > 0 &&
           obj.hasReferenceBase("virtual-machine") {
                err := obj.UpdateReference(
                        obj, "virtual-machine",
                        obj.virtual_machine_refs,
                        obj.baseMap["virtual-machine"])
                if err != nil {
                        return err
                }
        }

        return nil
}

func VirtualRouterByName(c contrail.ApiClient, fqn string) (*VirtualRouter, error) {
    obj, err := c.FindByName("virtual-router", fqn)
    if err != nil {
        return nil, err
    }
    return obj.(*VirtualRouter), nil
}

func VirtualRouterByUuid(c contrail.ApiClient, uuid string) (*VirtualRouter, error) {
    obj, err := c.FindByUuid("virtual-router", uuid)
    if err != nil {
        return nil, err
    }
    return obj.(*VirtualRouter), nil
}
