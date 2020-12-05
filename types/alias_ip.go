//
// Automatically generated. DO NOT EDIT.
//

package types

import (
        "encoding/json"

        "github.com/Juniper/contrail-go-api"
)

const (
	alias_ip_alias_ip_address uint64 = 1 << iota
	alias_ip_alias_ip_address_family
	alias_ip_id_perms
	alias_ip_perms2
	alias_ip_display_name
	alias_ip_project_refs
	alias_ip_virtual_machine_interface_refs
)

type AliasIp struct {
        contrail.ObjectBase
	alias_ip_address string
	alias_ip_address_family string
	id_perms IdPermsType
	perms2 PermType2
	display_name string
	project_refs contrail.ReferenceList
	virtual_machine_interface_refs contrail.ReferenceList
        valid uint64
        modified uint64
        baseMap map[string]contrail.ReferenceList
}

func (obj *AliasIp) GetType() string {
        return "alias-ip"
}

func (obj *AliasIp) GetDefaultParent() []string {
        name := []string{"default-domain", "default-project", "default-virtual-network", "default-alias-ip-pool"}
        return name
}

func (obj *AliasIp) GetDefaultParentType() string {
        return "alias-ip-pool"
}

func (obj *AliasIp) SetName(name string) {
        obj.VSetName(obj, name)
}

func (obj *AliasIp) SetParent(parent contrail.IObject) {
        obj.VSetParent(obj, parent)
}

func (obj *AliasIp) storeReferenceBase(
        name string, refList contrail.ReferenceList) {
        if obj.baseMap == nil {
                obj.baseMap = make(map[string]contrail.ReferenceList)
        }
        refCopy := make(contrail.ReferenceList, len(refList))
        copy(refCopy, refList)
        obj.baseMap[name] = refCopy
}

func (obj *AliasIp) hasReferenceBase(name string) bool {
        if obj.baseMap == nil {
                return false
        }
        _, exists := obj.baseMap[name]
        return exists
}

func (obj *AliasIp) UpdateDone() {
        obj.modified = 0
        obj.baseMap = nil
}


func (obj *AliasIp) GetAliasIpAddress() string {
        return obj.alias_ip_address
}

func (obj *AliasIp) SetAliasIpAddress(value string) {
        obj.alias_ip_address = value
        obj.modified |= alias_ip_alias_ip_address
}

func (obj *AliasIp) GetAliasIpAddressFamily() string {
        return obj.alias_ip_address_family
}

func (obj *AliasIp) SetAliasIpAddressFamily(value string) {
        obj.alias_ip_address_family = value
        obj.modified |= alias_ip_alias_ip_address_family
}

func (obj *AliasIp) GetIdPerms() IdPermsType {
        return obj.id_perms
}

func (obj *AliasIp) SetIdPerms(value *IdPermsType) {
        obj.id_perms = *value
        obj.modified |= alias_ip_id_perms
}

func (obj *AliasIp) GetPerms2() PermType2 {
        return obj.perms2
}

func (obj *AliasIp) SetPerms2(value *PermType2) {
        obj.perms2 = *value
        obj.modified |= alias_ip_perms2
}

func (obj *AliasIp) GetDisplayName() string {
        return obj.display_name
}

func (obj *AliasIp) SetDisplayName(value string) {
        obj.display_name = value
        obj.modified |= alias_ip_display_name
}

func (obj *AliasIp) readProjectRefs() error {
        if !obj.IsTransient() &&
                (obj.valid & alias_ip_project_refs == 0) {
                err := obj.GetField(obj, "project_refs")
                if err != nil {
                        return err
                }
        }
        return nil
}

func (obj *AliasIp) GetProjectRefs() (
        contrail.ReferenceList, error) {
        err := obj.readProjectRefs()
        if err != nil {
                return nil, err
        }
        return obj.project_refs, nil
}

func (obj *AliasIp) AddProject(
        rhs *Project) error {
        err := obj.readProjectRefs()
        if err != nil {
                return err
        }

        if obj.modified & alias_ip_project_refs == 0 {
                obj.storeReferenceBase("project", obj.project_refs)
        }

        ref := contrail.Reference {
                rhs.GetFQName(), rhs.GetUuid(), rhs.GetHref(), nil}
        obj.project_refs = append(obj.project_refs, ref)
        obj.modified |= alias_ip_project_refs
        return nil
}

func (obj *AliasIp) DeleteProject(uuid string) error {
        err := obj.readProjectRefs()
        if err != nil {
                return err
        }

        if obj.modified & alias_ip_project_refs == 0 {
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
        obj.modified |= alias_ip_project_refs
        return nil
}

func (obj *AliasIp) ClearProject() {
        if (obj.valid & alias_ip_project_refs != 0) &&
           (obj.modified & alias_ip_project_refs == 0) {
                obj.storeReferenceBase("project", obj.project_refs)
        }
        obj.project_refs = make([]contrail.Reference, 0)
        obj.valid |= alias_ip_project_refs
        obj.modified |= alias_ip_project_refs
}

func (obj *AliasIp) SetProjectList(
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


func (obj *AliasIp) readVirtualMachineInterfaceRefs() error {
        if !obj.IsTransient() &&
                (obj.valid & alias_ip_virtual_machine_interface_refs == 0) {
                err := obj.GetField(obj, "virtual_machine_interface_refs")
                if err != nil {
                        return err
                }
        }
        return nil
}

func (obj *AliasIp) GetVirtualMachineInterfaceRefs() (
        contrail.ReferenceList, error) {
        err := obj.readVirtualMachineInterfaceRefs()
        if err != nil {
                return nil, err
        }
        return obj.virtual_machine_interface_refs, nil
}

func (obj *AliasIp) AddVirtualMachineInterface(
        rhs *VirtualMachineInterface) error {
        err := obj.readVirtualMachineInterfaceRefs()
        if err != nil {
                return err
        }

        if obj.modified & alias_ip_virtual_machine_interface_refs == 0 {
                obj.storeReferenceBase("virtual-machine-interface", obj.virtual_machine_interface_refs)
        }

        ref := contrail.Reference {
                rhs.GetFQName(), rhs.GetUuid(), rhs.GetHref(), nil}
        obj.virtual_machine_interface_refs = append(obj.virtual_machine_interface_refs, ref)
        obj.modified |= alias_ip_virtual_machine_interface_refs
        return nil
}

func (obj *AliasIp) DeleteVirtualMachineInterface(uuid string) error {
        err := obj.readVirtualMachineInterfaceRefs()
        if err != nil {
                return err
        }

        if obj.modified & alias_ip_virtual_machine_interface_refs == 0 {
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
        obj.modified |= alias_ip_virtual_machine_interface_refs
        return nil
}

func (obj *AliasIp) ClearVirtualMachineInterface() {
        if (obj.valid & alias_ip_virtual_machine_interface_refs != 0) &&
           (obj.modified & alias_ip_virtual_machine_interface_refs == 0) {
                obj.storeReferenceBase("virtual-machine-interface", obj.virtual_machine_interface_refs)
        }
        obj.virtual_machine_interface_refs = make([]contrail.Reference, 0)
        obj.valid |= alias_ip_virtual_machine_interface_refs
        obj.modified |= alias_ip_virtual_machine_interface_refs
}

func (obj *AliasIp) SetVirtualMachineInterfaceList(
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


func (obj *AliasIp) MarshalJSON() ([]byte, error) {
        msg := map[string]*json.RawMessage {
        }
        err := obj.MarshalCommon(msg)
        if err != nil {
                return nil, err
        }

        if obj.modified & alias_ip_alias_ip_address != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.alias_ip_address)
                if err != nil {
                        return nil, err
                }
                msg["alias_ip_address"] = &value
        }

        if obj.modified & alias_ip_alias_ip_address_family != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.alias_ip_address_family)
                if err != nil {
                        return nil, err
                }
                msg["alias_ip_address_family"] = &value
        }

        if obj.modified & alias_ip_id_perms != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.id_perms)
                if err != nil {
                        return nil, err
                }
                msg["id_perms"] = &value
        }

        if obj.modified & alias_ip_perms2 != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.perms2)
                if err != nil {
                        return nil, err
                }
                msg["perms2"] = &value
        }

        if obj.modified & alias_ip_display_name != 0 {
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

func (obj *AliasIp) UnmarshalJSON(body []byte) error {
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
                case "alias_ip_address":
                        err = json.Unmarshal(value, &obj.alias_ip_address)
                        if err == nil {
                                obj.valid |= alias_ip_alias_ip_address
                        }
                        break
                case "alias_ip_address_family":
                        err = json.Unmarshal(value, &obj.alias_ip_address_family)
                        if err == nil {
                                obj.valid |= alias_ip_alias_ip_address_family
                        }
                        break
                case "id_perms":
                        err = json.Unmarshal(value, &obj.id_perms)
                        if err == nil {
                                obj.valid |= alias_ip_id_perms
                        }
                        break
                case "perms2":
                        err = json.Unmarshal(value, &obj.perms2)
                        if err == nil {
                                obj.valid |= alias_ip_perms2
                        }
                        break
                case "display_name":
                        err = json.Unmarshal(value, &obj.display_name)
                        if err == nil {
                                obj.valid |= alias_ip_display_name
                        }
                        break
                case "project_refs":
                        err = json.Unmarshal(value, &obj.project_refs)
                        if err == nil {
                                obj.valid |= alias_ip_project_refs
                        }
                        break
                case "virtual_machine_interface_refs":
                        err = json.Unmarshal(value, &obj.virtual_machine_interface_refs)
                        if err == nil {
                                obj.valid |= alias_ip_virtual_machine_interface_refs
                        }
                        break
                }
                if err != nil {
                        return err
                }
        }
        return nil
}

func (obj *AliasIp) UpdateObject() ([]byte, error) {
        msg := map[string]*json.RawMessage {
        }
        err := obj.MarshalId(msg)
        if err != nil {
                return nil, err
        }

        if obj.modified & alias_ip_alias_ip_address != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.alias_ip_address)
                if err != nil {
                        return nil, err
                }
                msg["alias_ip_address"] = &value
        }

        if obj.modified & alias_ip_alias_ip_address_family != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.alias_ip_address_family)
                if err != nil {
                        return nil, err
                }
                msg["alias_ip_address_family"] = &value
        }

        if obj.modified & alias_ip_id_perms != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.id_perms)
                if err != nil {
                        return nil, err
                }
                msg["id_perms"] = &value
        }

        if obj.modified & alias_ip_perms2 != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.perms2)
                if err != nil {
                        return nil, err
                }
                msg["perms2"] = &value
        }

        if obj.modified & alias_ip_display_name != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.display_name)
                if err != nil {
                        return nil, err
                }
                msg["display_name"] = &value
        }

        if obj.modified & alias_ip_project_refs != 0 {
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


        if obj.modified & alias_ip_virtual_machine_interface_refs != 0 {
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

func (obj *AliasIp) UpdateReferences() error {

        if (obj.modified & alias_ip_project_refs != 0) &&
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

        if (obj.modified & alias_ip_virtual_machine_interface_refs != 0) &&
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

func AliasIpByName(c contrail.ApiClient, fqn string) (*AliasIp, error) {
    obj, err := c.FindByName("alias-ip", fqn)
    if err != nil {
        return nil, err
    }
    return obj.(*AliasIp), nil
}

func AliasIpByUuid(c contrail.ApiClient, uuid string) (*AliasIp, error) {
    obj, err := c.FindByUuid("alias-ip", uuid)
    if err != nil {
        return nil, err
    }
    return obj.(*AliasIp), nil
}
