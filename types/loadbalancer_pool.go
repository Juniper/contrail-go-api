//
// Automatically generated. DO NOT EDIT.
//

package types

import (
        "encoding/json"

        "github.com/Juniper/contrail-go-api"
)

const (
	loadbalancer_pool_loadbalancer_pool_properties uint64 = 1 << iota
	loadbalancer_pool_loadbalancer_pool_provider
	loadbalancer_pool_loadbalancer_pool_custom_attributes
	loadbalancer_pool_id_perms
	loadbalancer_pool_perms2
	loadbalancer_pool_display_name
	loadbalancer_pool_service_instance_refs
	loadbalancer_pool_virtual_machine_interface_refs
	loadbalancer_pool_loadbalancer_listener_refs
	loadbalancer_pool_service_appliance_set_refs
	loadbalancer_pool_loadbalancer_members
	loadbalancer_pool_loadbalancer_healthmonitor_refs
	loadbalancer_pool_virtual_ip_back_refs
)

type LoadbalancerPool struct {
        contrail.ObjectBase
	loadbalancer_pool_properties LoadbalancerPoolType
	loadbalancer_pool_provider string
	loadbalancer_pool_custom_attributes KeyValuePairs
	id_perms IdPermsType
	perms2 PermType2
	display_name string
	service_instance_refs contrail.ReferenceList
	virtual_machine_interface_refs contrail.ReferenceList
	loadbalancer_listener_refs contrail.ReferenceList
	service_appliance_set_refs contrail.ReferenceList
	loadbalancer_members contrail.ReferenceList
	loadbalancer_healthmonitor_refs contrail.ReferenceList
	virtual_ip_back_refs contrail.ReferenceList
        valid uint64
        modified uint64
        baseMap map[string]contrail.ReferenceList
}

func (obj *LoadbalancerPool) GetType() string {
        return "loadbalancer-pool"
}

func (obj *LoadbalancerPool) GetDefaultParent() []string {
        name := []string{"default-domain", "default-project"}
        return name
}

func (obj *LoadbalancerPool) GetDefaultParentType() string {
        return "project"
}

func (obj *LoadbalancerPool) SetName(name string) {
        obj.VSetName(obj, name)
}

func (obj *LoadbalancerPool) SetParent(parent contrail.IObject) {
        obj.VSetParent(obj, parent)
}

func (obj *LoadbalancerPool) storeReferenceBase(
        name string, refList contrail.ReferenceList) {
        if obj.baseMap == nil {
                obj.baseMap = make(map[string]contrail.ReferenceList)
        }
        refCopy := make(contrail.ReferenceList, len(refList))
        copy(refCopy, refList)
        obj.baseMap[name] = refCopy
}

func (obj *LoadbalancerPool) hasReferenceBase(name string) bool {
        if obj.baseMap == nil {
                return false
        }
        _, exists := obj.baseMap[name]
        return exists
}

func (obj *LoadbalancerPool) UpdateDone() {
        obj.modified = 0
        obj.baseMap = nil
}


func (obj *LoadbalancerPool) GetLoadbalancerPoolProperties() LoadbalancerPoolType {
        return obj.loadbalancer_pool_properties
}

func (obj *LoadbalancerPool) SetLoadbalancerPoolProperties(value *LoadbalancerPoolType) {
        obj.loadbalancer_pool_properties = *value
        obj.modified |= loadbalancer_pool_loadbalancer_pool_properties
}

func (obj *LoadbalancerPool) GetLoadbalancerPoolProvider() string {
        return obj.loadbalancer_pool_provider
}

func (obj *LoadbalancerPool) SetLoadbalancerPoolProvider(value string) {
        obj.loadbalancer_pool_provider = value
        obj.modified |= loadbalancer_pool_loadbalancer_pool_provider
}

func (obj *LoadbalancerPool) GetLoadbalancerPoolCustomAttributes() KeyValuePairs {
        return obj.loadbalancer_pool_custom_attributes
}

func (obj *LoadbalancerPool) SetLoadbalancerPoolCustomAttributes(value *KeyValuePairs) {
        obj.loadbalancer_pool_custom_attributes = *value
        obj.modified |= loadbalancer_pool_loadbalancer_pool_custom_attributes
}

func (obj *LoadbalancerPool) GetIdPerms() IdPermsType {
        return obj.id_perms
}

func (obj *LoadbalancerPool) SetIdPerms(value *IdPermsType) {
        obj.id_perms = *value
        obj.modified |= loadbalancer_pool_id_perms
}

func (obj *LoadbalancerPool) GetPerms2() PermType2 {
        return obj.perms2
}

func (obj *LoadbalancerPool) SetPerms2(value *PermType2) {
        obj.perms2 = *value
        obj.modified |= loadbalancer_pool_perms2
}

func (obj *LoadbalancerPool) GetDisplayName() string {
        return obj.display_name
}

func (obj *LoadbalancerPool) SetDisplayName(value string) {
        obj.display_name = value
        obj.modified |= loadbalancer_pool_display_name
}

func (obj *LoadbalancerPool) readLoadbalancerMembers() error {
        if !obj.IsTransient() &&
                (obj.valid & loadbalancer_pool_loadbalancer_members == 0) {
                err := obj.GetField(obj, "loadbalancer_members")
                if err != nil {
                        return err
                }
        }
        return nil
}

func (obj *LoadbalancerPool) GetLoadbalancerMembers() (
        contrail.ReferenceList, error) {
        err := obj.readLoadbalancerMembers()
        if err != nil {
                return nil, err
        }
        return obj.loadbalancer_members, nil
}

func (obj *LoadbalancerPool) readServiceInstanceRefs() error {
        if !obj.IsTransient() &&
                (obj.valid & loadbalancer_pool_service_instance_refs == 0) {
                err := obj.GetField(obj, "service_instance_refs")
                if err != nil {
                        return err
                }
        }
        return nil
}

func (obj *LoadbalancerPool) GetServiceInstanceRefs() (
        contrail.ReferenceList, error) {
        err := obj.readServiceInstanceRefs()
        if err != nil {
                return nil, err
        }
        return obj.service_instance_refs, nil
}

func (obj *LoadbalancerPool) AddServiceInstance(
        rhs *ServiceInstance) error {
        err := obj.readServiceInstanceRefs()
        if err != nil {
                return err
        }

        if obj.modified & loadbalancer_pool_service_instance_refs == 0 {
                obj.storeReferenceBase("service-instance", obj.service_instance_refs)
        }

        ref := contrail.Reference {
                rhs.GetFQName(), rhs.GetUuid(), rhs.GetHref(), nil}
        obj.service_instance_refs = append(obj.service_instance_refs, ref)
        obj.modified |= loadbalancer_pool_service_instance_refs
        return nil
}

func (obj *LoadbalancerPool) DeleteServiceInstance(uuid string) error {
        err := obj.readServiceInstanceRefs()
        if err != nil {
                return err
        }

        if obj.modified & loadbalancer_pool_service_instance_refs == 0 {
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
        obj.modified |= loadbalancer_pool_service_instance_refs
        return nil
}

func (obj *LoadbalancerPool) ClearServiceInstance() {
        if (obj.valid & loadbalancer_pool_service_instance_refs != 0) &&
           (obj.modified & loadbalancer_pool_service_instance_refs == 0) {
                obj.storeReferenceBase("service-instance", obj.service_instance_refs)
        }
        obj.service_instance_refs = make([]contrail.Reference, 0)
        obj.valid |= loadbalancer_pool_service_instance_refs
        obj.modified |= loadbalancer_pool_service_instance_refs
}

func (obj *LoadbalancerPool) SetServiceInstanceList(
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


func (obj *LoadbalancerPool) readVirtualMachineInterfaceRefs() error {
        if !obj.IsTransient() &&
                (obj.valid & loadbalancer_pool_virtual_machine_interface_refs == 0) {
                err := obj.GetField(obj, "virtual_machine_interface_refs")
                if err != nil {
                        return err
                }
        }
        return nil
}

func (obj *LoadbalancerPool) GetVirtualMachineInterfaceRefs() (
        contrail.ReferenceList, error) {
        err := obj.readVirtualMachineInterfaceRefs()
        if err != nil {
                return nil, err
        }
        return obj.virtual_machine_interface_refs, nil
}

func (obj *LoadbalancerPool) AddVirtualMachineInterface(
        rhs *VirtualMachineInterface) error {
        err := obj.readVirtualMachineInterfaceRefs()
        if err != nil {
                return err
        }

        if obj.modified & loadbalancer_pool_virtual_machine_interface_refs == 0 {
                obj.storeReferenceBase("virtual-machine-interface", obj.virtual_machine_interface_refs)
        }

        ref := contrail.Reference {
                rhs.GetFQName(), rhs.GetUuid(), rhs.GetHref(), nil}
        obj.virtual_machine_interface_refs = append(obj.virtual_machine_interface_refs, ref)
        obj.modified |= loadbalancer_pool_virtual_machine_interface_refs
        return nil
}

func (obj *LoadbalancerPool) DeleteVirtualMachineInterface(uuid string) error {
        err := obj.readVirtualMachineInterfaceRefs()
        if err != nil {
                return err
        }

        if obj.modified & loadbalancer_pool_virtual_machine_interface_refs == 0 {
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
        obj.modified |= loadbalancer_pool_virtual_machine_interface_refs
        return nil
}

func (obj *LoadbalancerPool) ClearVirtualMachineInterface() {
        if (obj.valid & loadbalancer_pool_virtual_machine_interface_refs != 0) &&
           (obj.modified & loadbalancer_pool_virtual_machine_interface_refs == 0) {
                obj.storeReferenceBase("virtual-machine-interface", obj.virtual_machine_interface_refs)
        }
        obj.virtual_machine_interface_refs = make([]contrail.Reference, 0)
        obj.valid |= loadbalancer_pool_virtual_machine_interface_refs
        obj.modified |= loadbalancer_pool_virtual_machine_interface_refs
}

func (obj *LoadbalancerPool) SetVirtualMachineInterfaceList(
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


func (obj *LoadbalancerPool) readLoadbalancerListenerRefs() error {
        if !obj.IsTransient() &&
                (obj.valid & loadbalancer_pool_loadbalancer_listener_refs == 0) {
                err := obj.GetField(obj, "loadbalancer_listener_refs")
                if err != nil {
                        return err
                }
        }
        return nil
}

func (obj *LoadbalancerPool) GetLoadbalancerListenerRefs() (
        contrail.ReferenceList, error) {
        err := obj.readLoadbalancerListenerRefs()
        if err != nil {
                return nil, err
        }
        return obj.loadbalancer_listener_refs, nil
}

func (obj *LoadbalancerPool) AddLoadbalancerListener(
        rhs *LoadbalancerListener) error {
        err := obj.readLoadbalancerListenerRefs()
        if err != nil {
                return err
        }

        if obj.modified & loadbalancer_pool_loadbalancer_listener_refs == 0 {
                obj.storeReferenceBase("loadbalancer-listener", obj.loadbalancer_listener_refs)
        }

        ref := contrail.Reference {
                rhs.GetFQName(), rhs.GetUuid(), rhs.GetHref(), nil}
        obj.loadbalancer_listener_refs = append(obj.loadbalancer_listener_refs, ref)
        obj.modified |= loadbalancer_pool_loadbalancer_listener_refs
        return nil
}

func (obj *LoadbalancerPool) DeleteLoadbalancerListener(uuid string) error {
        err := obj.readLoadbalancerListenerRefs()
        if err != nil {
                return err
        }

        if obj.modified & loadbalancer_pool_loadbalancer_listener_refs == 0 {
                obj.storeReferenceBase("loadbalancer-listener", obj.loadbalancer_listener_refs)
        }

        for i, ref := range obj.loadbalancer_listener_refs {
                if ref.Uuid == uuid {
                        obj.loadbalancer_listener_refs = append(
                                obj.loadbalancer_listener_refs[:i],
                                obj.loadbalancer_listener_refs[i+1:]...)
                        break
                }
        }
        obj.modified |= loadbalancer_pool_loadbalancer_listener_refs
        return nil
}

func (obj *LoadbalancerPool) ClearLoadbalancerListener() {
        if (obj.valid & loadbalancer_pool_loadbalancer_listener_refs != 0) &&
           (obj.modified & loadbalancer_pool_loadbalancer_listener_refs == 0) {
                obj.storeReferenceBase("loadbalancer-listener", obj.loadbalancer_listener_refs)
        }
        obj.loadbalancer_listener_refs = make([]contrail.Reference, 0)
        obj.valid |= loadbalancer_pool_loadbalancer_listener_refs
        obj.modified |= loadbalancer_pool_loadbalancer_listener_refs
}

func (obj *LoadbalancerPool) SetLoadbalancerListenerList(
        refList []contrail.ReferencePair) {
        obj.ClearLoadbalancerListener()
        obj.loadbalancer_listener_refs = make([]contrail.Reference, len(refList))
        for i, pair := range refList {
                obj.loadbalancer_listener_refs[i] = contrail.Reference {
                        pair.Object.GetFQName(),
                        pair.Object.GetUuid(),
                        pair.Object.GetHref(),
                        pair.Attribute,
                }
        }
}


func (obj *LoadbalancerPool) readServiceApplianceSetRefs() error {
        if !obj.IsTransient() &&
                (obj.valid & loadbalancer_pool_service_appliance_set_refs == 0) {
                err := obj.GetField(obj, "service_appliance_set_refs")
                if err != nil {
                        return err
                }
        }
        return nil
}

func (obj *LoadbalancerPool) GetServiceApplianceSetRefs() (
        contrail.ReferenceList, error) {
        err := obj.readServiceApplianceSetRefs()
        if err != nil {
                return nil, err
        }
        return obj.service_appliance_set_refs, nil
}

func (obj *LoadbalancerPool) AddServiceApplianceSet(
        rhs *ServiceApplianceSet) error {
        err := obj.readServiceApplianceSetRefs()
        if err != nil {
                return err
        }

        if obj.modified & loadbalancer_pool_service_appliance_set_refs == 0 {
                obj.storeReferenceBase("service-appliance-set", obj.service_appliance_set_refs)
        }

        ref := contrail.Reference {
                rhs.GetFQName(), rhs.GetUuid(), rhs.GetHref(), nil}
        obj.service_appliance_set_refs = append(obj.service_appliance_set_refs, ref)
        obj.modified |= loadbalancer_pool_service_appliance_set_refs
        return nil
}

func (obj *LoadbalancerPool) DeleteServiceApplianceSet(uuid string) error {
        err := obj.readServiceApplianceSetRefs()
        if err != nil {
                return err
        }

        if obj.modified & loadbalancer_pool_service_appliance_set_refs == 0 {
                obj.storeReferenceBase("service-appliance-set", obj.service_appliance_set_refs)
        }

        for i, ref := range obj.service_appliance_set_refs {
                if ref.Uuid == uuid {
                        obj.service_appliance_set_refs = append(
                                obj.service_appliance_set_refs[:i],
                                obj.service_appliance_set_refs[i+1:]...)
                        break
                }
        }
        obj.modified |= loadbalancer_pool_service_appliance_set_refs
        return nil
}

func (obj *LoadbalancerPool) ClearServiceApplianceSet() {
        if (obj.valid & loadbalancer_pool_service_appliance_set_refs != 0) &&
           (obj.modified & loadbalancer_pool_service_appliance_set_refs == 0) {
                obj.storeReferenceBase("service-appliance-set", obj.service_appliance_set_refs)
        }
        obj.service_appliance_set_refs = make([]contrail.Reference, 0)
        obj.valid |= loadbalancer_pool_service_appliance_set_refs
        obj.modified |= loadbalancer_pool_service_appliance_set_refs
}

func (obj *LoadbalancerPool) SetServiceApplianceSetList(
        refList []contrail.ReferencePair) {
        obj.ClearServiceApplianceSet()
        obj.service_appliance_set_refs = make([]contrail.Reference, len(refList))
        for i, pair := range refList {
                obj.service_appliance_set_refs[i] = contrail.Reference {
                        pair.Object.GetFQName(),
                        pair.Object.GetUuid(),
                        pair.Object.GetHref(),
                        pair.Attribute,
                }
        }
}


func (obj *LoadbalancerPool) readLoadbalancerHealthmonitorRefs() error {
        if !obj.IsTransient() &&
                (obj.valid & loadbalancer_pool_loadbalancer_healthmonitor_refs == 0) {
                err := obj.GetField(obj, "loadbalancer_healthmonitor_refs")
                if err != nil {
                        return err
                }
        }
        return nil
}

func (obj *LoadbalancerPool) GetLoadbalancerHealthmonitorRefs() (
        contrail.ReferenceList, error) {
        err := obj.readLoadbalancerHealthmonitorRefs()
        if err != nil {
                return nil, err
        }
        return obj.loadbalancer_healthmonitor_refs, nil
}

func (obj *LoadbalancerPool) AddLoadbalancerHealthmonitor(
        rhs *LoadbalancerHealthmonitor) error {
        err := obj.readLoadbalancerHealthmonitorRefs()
        if err != nil {
                return err
        }

        if obj.modified & loadbalancer_pool_loadbalancer_healthmonitor_refs == 0 {
                obj.storeReferenceBase("loadbalancer-healthmonitor", obj.loadbalancer_healthmonitor_refs)
        }

        ref := contrail.Reference {
                rhs.GetFQName(), rhs.GetUuid(), rhs.GetHref(), nil}
        obj.loadbalancer_healthmonitor_refs = append(obj.loadbalancer_healthmonitor_refs, ref)
        obj.modified |= loadbalancer_pool_loadbalancer_healthmonitor_refs
        return nil
}

func (obj *LoadbalancerPool) DeleteLoadbalancerHealthmonitor(uuid string) error {
        err := obj.readLoadbalancerHealthmonitorRefs()
        if err != nil {
                return err
        }

        if obj.modified & loadbalancer_pool_loadbalancer_healthmonitor_refs == 0 {
                obj.storeReferenceBase("loadbalancer-healthmonitor", obj.loadbalancer_healthmonitor_refs)
        }

        for i, ref := range obj.loadbalancer_healthmonitor_refs {
                if ref.Uuid == uuid {
                        obj.loadbalancer_healthmonitor_refs = append(
                                obj.loadbalancer_healthmonitor_refs[:i],
                                obj.loadbalancer_healthmonitor_refs[i+1:]...)
                        break
                }
        }
        obj.modified |= loadbalancer_pool_loadbalancer_healthmonitor_refs
        return nil
}

func (obj *LoadbalancerPool) ClearLoadbalancerHealthmonitor() {
        if (obj.valid & loadbalancer_pool_loadbalancer_healthmonitor_refs != 0) &&
           (obj.modified & loadbalancer_pool_loadbalancer_healthmonitor_refs == 0) {
                obj.storeReferenceBase("loadbalancer-healthmonitor", obj.loadbalancer_healthmonitor_refs)
        }
        obj.loadbalancer_healthmonitor_refs = make([]contrail.Reference, 0)
        obj.valid |= loadbalancer_pool_loadbalancer_healthmonitor_refs
        obj.modified |= loadbalancer_pool_loadbalancer_healthmonitor_refs
}

func (obj *LoadbalancerPool) SetLoadbalancerHealthmonitorList(
        refList []contrail.ReferencePair) {
        obj.ClearLoadbalancerHealthmonitor()
        obj.loadbalancer_healthmonitor_refs = make([]contrail.Reference, len(refList))
        for i, pair := range refList {
                obj.loadbalancer_healthmonitor_refs[i] = contrail.Reference {
                        pair.Object.GetFQName(),
                        pair.Object.GetUuid(),
                        pair.Object.GetHref(),
                        pair.Attribute,
                }
        }
}


func (obj *LoadbalancerPool) readVirtualIpBackRefs() error {
        if !obj.IsTransient() &&
                (obj.valid & loadbalancer_pool_virtual_ip_back_refs == 0) {
                err := obj.GetField(obj, "virtual_ip_back_refs")
                if err != nil {
                        return err
                }
        }
        return nil
}

func (obj *LoadbalancerPool) GetVirtualIpBackRefs() (
        contrail.ReferenceList, error) {
        err := obj.readVirtualIpBackRefs()
        if err != nil {
                return nil, err
        }
        return obj.virtual_ip_back_refs, nil
}

func (obj *LoadbalancerPool) MarshalJSON() ([]byte, error) {
        msg := map[string]*json.RawMessage {
        }
        err := obj.MarshalCommon(msg)
        if err != nil {
                return nil, err
        }

        if obj.modified & loadbalancer_pool_loadbalancer_pool_properties != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.loadbalancer_pool_properties)
                if err != nil {
                        return nil, err
                }
                msg["loadbalancer_pool_properties"] = &value
        }

        if obj.modified & loadbalancer_pool_loadbalancer_pool_provider != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.loadbalancer_pool_provider)
                if err != nil {
                        return nil, err
                }
                msg["loadbalancer_pool_provider"] = &value
        }

        if obj.modified & loadbalancer_pool_loadbalancer_pool_custom_attributes != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.loadbalancer_pool_custom_attributes)
                if err != nil {
                        return nil, err
                }
                msg["loadbalancer_pool_custom_attributes"] = &value
        }

        if obj.modified & loadbalancer_pool_id_perms != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.id_perms)
                if err != nil {
                        return nil, err
                }
                msg["id_perms"] = &value
        }

        if obj.modified & loadbalancer_pool_perms2 != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.perms2)
                if err != nil {
                        return nil, err
                }
                msg["perms2"] = &value
        }

        if obj.modified & loadbalancer_pool_display_name != 0 {
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

        if len(obj.virtual_machine_interface_refs) > 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.virtual_machine_interface_refs)
                if err != nil {
                        return nil, err
                }
                msg["virtual_machine_interface_refs"] = &value
        }

        if len(obj.loadbalancer_listener_refs) > 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.loadbalancer_listener_refs)
                if err != nil {
                        return nil, err
                }
                msg["loadbalancer_listener_refs"] = &value
        }

        if len(obj.service_appliance_set_refs) > 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.service_appliance_set_refs)
                if err != nil {
                        return nil, err
                }
                msg["service_appliance_set_refs"] = &value
        }

        if len(obj.loadbalancer_healthmonitor_refs) > 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.loadbalancer_healthmonitor_refs)
                if err != nil {
                        return nil, err
                }
                msg["loadbalancer_healthmonitor_refs"] = &value
        }

        return json.Marshal(msg)
}

func (obj *LoadbalancerPool) UnmarshalJSON(body []byte) error {
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
                case "loadbalancer_pool_properties":
                        err = json.Unmarshal(value, &obj.loadbalancer_pool_properties)
                        if err == nil {
                                obj.valid |= loadbalancer_pool_loadbalancer_pool_properties
                        }
                        break
                case "loadbalancer_pool_provider":
                        err = json.Unmarshal(value, &obj.loadbalancer_pool_provider)
                        if err == nil {
                                obj.valid |= loadbalancer_pool_loadbalancer_pool_provider
                        }
                        break
                case "loadbalancer_pool_custom_attributes":
                        err = json.Unmarshal(value, &obj.loadbalancer_pool_custom_attributes)
                        if err == nil {
                                obj.valid |= loadbalancer_pool_loadbalancer_pool_custom_attributes
                        }
                        break
                case "id_perms":
                        err = json.Unmarshal(value, &obj.id_perms)
                        if err == nil {
                                obj.valid |= loadbalancer_pool_id_perms
                        }
                        break
                case "perms2":
                        err = json.Unmarshal(value, &obj.perms2)
                        if err == nil {
                                obj.valid |= loadbalancer_pool_perms2
                        }
                        break
                case "display_name":
                        err = json.Unmarshal(value, &obj.display_name)
                        if err == nil {
                                obj.valid |= loadbalancer_pool_display_name
                        }
                        break
                case "service_instance_refs":
                        err = json.Unmarshal(value, &obj.service_instance_refs)
                        if err == nil {
                                obj.valid |= loadbalancer_pool_service_instance_refs
                        }
                        break
                case "virtual_machine_interface_refs":
                        err = json.Unmarshal(value, &obj.virtual_machine_interface_refs)
                        if err == nil {
                                obj.valid |= loadbalancer_pool_virtual_machine_interface_refs
                        }
                        break
                case "loadbalancer_listener_refs":
                        err = json.Unmarshal(value, &obj.loadbalancer_listener_refs)
                        if err == nil {
                                obj.valid |= loadbalancer_pool_loadbalancer_listener_refs
                        }
                        break
                case "service_appliance_set_refs":
                        err = json.Unmarshal(value, &obj.service_appliance_set_refs)
                        if err == nil {
                                obj.valid |= loadbalancer_pool_service_appliance_set_refs
                        }
                        break
                case "loadbalancer_members":
                        err = json.Unmarshal(value, &obj.loadbalancer_members)
                        if err == nil {
                                obj.valid |= loadbalancer_pool_loadbalancer_members
                        }
                        break
                case "loadbalancer_healthmonitor_refs":
                        err = json.Unmarshal(value, &obj.loadbalancer_healthmonitor_refs)
                        if err == nil {
                                obj.valid |= loadbalancer_pool_loadbalancer_healthmonitor_refs
                        }
                        break
                case "virtual_ip_back_refs":
                        err = json.Unmarshal(value, &obj.virtual_ip_back_refs)
                        if err == nil {
                                obj.valid |= loadbalancer_pool_virtual_ip_back_refs
                        }
                        break
                }
                if err != nil {
                        return err
                }
        }
        return nil
}

func (obj *LoadbalancerPool) UpdateObject() ([]byte, error) {
        msg := map[string]*json.RawMessage {
        }
        err := obj.MarshalId(msg)
        if err != nil {
                return nil, err
        }

        if obj.modified & loadbalancer_pool_loadbalancer_pool_properties != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.loadbalancer_pool_properties)
                if err != nil {
                        return nil, err
                }
                msg["loadbalancer_pool_properties"] = &value
        }

        if obj.modified & loadbalancer_pool_loadbalancer_pool_provider != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.loadbalancer_pool_provider)
                if err != nil {
                        return nil, err
                }
                msg["loadbalancer_pool_provider"] = &value
        }

        if obj.modified & loadbalancer_pool_loadbalancer_pool_custom_attributes != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.loadbalancer_pool_custom_attributes)
                if err != nil {
                        return nil, err
                }
                msg["loadbalancer_pool_custom_attributes"] = &value
        }

        if obj.modified & loadbalancer_pool_id_perms != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.id_perms)
                if err != nil {
                        return nil, err
                }
                msg["id_perms"] = &value
        }

        if obj.modified & loadbalancer_pool_perms2 != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.perms2)
                if err != nil {
                        return nil, err
                }
                msg["perms2"] = &value
        }

        if obj.modified & loadbalancer_pool_display_name != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.display_name)
                if err != nil {
                        return nil, err
                }
                msg["display_name"] = &value
        }

        if obj.modified & loadbalancer_pool_service_instance_refs != 0 {
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


        if obj.modified & loadbalancer_pool_virtual_machine_interface_refs != 0 {
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


        if obj.modified & loadbalancer_pool_loadbalancer_listener_refs != 0 {
                if len(obj.loadbalancer_listener_refs) == 0 {
                        var value json.RawMessage
                        value, err := json.Marshal(
                                          make([]contrail.Reference, 0))
                        if err != nil {
                                return nil, err
                        }
                        msg["loadbalancer_listener_refs"] = &value
                } else if !obj.hasReferenceBase("loadbalancer-listener") {
                        var value json.RawMessage
                        value, err := json.Marshal(&obj.loadbalancer_listener_refs)
                        if err != nil {
                                return nil, err
                        }
                        msg["loadbalancer_listener_refs"] = &value
                }
        }


        if obj.modified & loadbalancer_pool_service_appliance_set_refs != 0 {
                if len(obj.service_appliance_set_refs) == 0 {
                        var value json.RawMessage
                        value, err := json.Marshal(
                                          make([]contrail.Reference, 0))
                        if err != nil {
                                return nil, err
                        }
                        msg["service_appliance_set_refs"] = &value
                } else if !obj.hasReferenceBase("service-appliance-set") {
                        var value json.RawMessage
                        value, err := json.Marshal(&obj.service_appliance_set_refs)
                        if err != nil {
                                return nil, err
                        }
                        msg["service_appliance_set_refs"] = &value
                }
        }


        if obj.modified & loadbalancer_pool_loadbalancer_healthmonitor_refs != 0 {
                if len(obj.loadbalancer_healthmonitor_refs) == 0 {
                        var value json.RawMessage
                        value, err := json.Marshal(
                                          make([]contrail.Reference, 0))
                        if err != nil {
                                return nil, err
                        }
                        msg["loadbalancer_healthmonitor_refs"] = &value
                } else if !obj.hasReferenceBase("loadbalancer-healthmonitor") {
                        var value json.RawMessage
                        value, err := json.Marshal(&obj.loadbalancer_healthmonitor_refs)
                        if err != nil {
                                return nil, err
                        }
                        msg["loadbalancer_healthmonitor_refs"] = &value
                }
        }


        return json.Marshal(msg)
}

func (obj *LoadbalancerPool) UpdateReferences() error {

        if (obj.modified & loadbalancer_pool_service_instance_refs != 0) &&
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

        if (obj.modified & loadbalancer_pool_virtual_machine_interface_refs != 0) &&
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

        if (obj.modified & loadbalancer_pool_loadbalancer_listener_refs != 0) &&
           len(obj.loadbalancer_listener_refs) > 0 &&
           obj.hasReferenceBase("loadbalancer-listener") {
                err := obj.UpdateReference(
                        obj, "loadbalancer-listener",
                        obj.loadbalancer_listener_refs,
                        obj.baseMap["loadbalancer-listener"])
                if err != nil {
                        return err
                }
        }

        if (obj.modified & loadbalancer_pool_service_appliance_set_refs != 0) &&
           len(obj.service_appliance_set_refs) > 0 &&
           obj.hasReferenceBase("service-appliance-set") {
                err := obj.UpdateReference(
                        obj, "service-appliance-set",
                        obj.service_appliance_set_refs,
                        obj.baseMap["service-appliance-set"])
                if err != nil {
                        return err
                }
        }

        if (obj.modified & loadbalancer_pool_loadbalancer_healthmonitor_refs != 0) &&
           len(obj.loadbalancer_healthmonitor_refs) > 0 &&
           obj.hasReferenceBase("loadbalancer-healthmonitor") {
                err := obj.UpdateReference(
                        obj, "loadbalancer-healthmonitor",
                        obj.loadbalancer_healthmonitor_refs,
                        obj.baseMap["loadbalancer-healthmonitor"])
                if err != nil {
                        return err
                }
        }

        return nil
}

func LoadbalancerPoolByName(c contrail.ApiClient, fqn string) (*LoadbalancerPool, error) {
    obj, err := c.FindByName("loadbalancer-pool", fqn)
    if err != nil {
        return nil, err
    }
    return obj.(*LoadbalancerPool), nil
}

func LoadbalancerPoolByUuid(c contrail.ApiClient, uuid string) (*LoadbalancerPool, error) {
    obj, err := c.FindByUuid("loadbalancer-pool", uuid)
    if err != nil {
        return nil, err
    }
    return obj.(*LoadbalancerPool), nil
}
