//
// Automatically generated. DO NOT EDIT.
//

package types

import (
        "encoding/json"

        "github.com/michaelhenkel/contrail-go-api"
)

const (
	bgp_as_a_service_autonomous_system uint64 = 1 << iota
	bgp_as_a_service_bgpaas_ip_address
	bgp_as_a_service_bgpaas_session_attributes
	bgp_as_a_service_bgpaas_ipv4_mapped_ipv6_nexthop
	bgp_as_a_service_bgpaas_suppress_route_advertisement
	bgp_as_a_service_id_perms
	bgp_as_a_service_perms2
	bgp_as_a_service_display_name
	bgp_as_a_service_virtual_machine_interface_refs
)

type BgpAsAService struct {
        contrail.ObjectBase
	autonomous_system int
	bgpaas_ip_address string
	bgpaas_session_attributes string
	bgpaas_ipv4_mapped_ipv6_nexthop bool
	bgpaas_suppress_route_advertisement bool
	id_perms IdPermsType
	perms2 PermType2
	display_name string
	virtual_machine_interface_refs contrail.ReferenceList
        valid uint64
        modified uint64
        baseMap map[string]contrail.ReferenceList
}

func (obj *BgpAsAService) GetType() string {
        return "bgp-as-a-service"
}

func (obj *BgpAsAService) GetDefaultParent() []string {
        name := []string{"default-domain", "default-project"}
        return name
}

func (obj *BgpAsAService) GetDefaultParentType() string {
        return "project"
}

func (obj *BgpAsAService) SetName(name string) {
        obj.VSetName(obj, name)
}

func (obj *BgpAsAService) SetParent(parent contrail.IObject) {
        obj.VSetParent(obj, parent)
}

func (obj *BgpAsAService) storeReferenceBase(
        name string, refList contrail.ReferenceList) {
        if obj.baseMap == nil {
                obj.baseMap = make(map[string]contrail.ReferenceList)
        }
        refCopy := make(contrail.ReferenceList, len(refList))
        copy(refCopy, refList)
        obj.baseMap[name] = refCopy
}

func (obj *BgpAsAService) hasReferenceBase(name string) bool {
        if obj.baseMap == nil {
                return false
        }
        _, exists := obj.baseMap[name]
        return exists
}

func (obj *BgpAsAService) UpdateDone() {
        obj.modified = 0
        obj.baseMap = nil
}


func (obj *BgpAsAService) GetAutonomousSystem() int {
        return obj.autonomous_system
}

func (obj *BgpAsAService) SetAutonomousSystem(value int) {
        obj.autonomous_system = value
        obj.modified |= bgp_as_a_service_autonomous_system
}

func (obj *BgpAsAService) GetBgpaasIpAddress() string {
        return obj.bgpaas_ip_address
}

func (obj *BgpAsAService) SetBgpaasIpAddress(value string) {
        obj.bgpaas_ip_address = value
        obj.modified |= bgp_as_a_service_bgpaas_ip_address
}

func (obj *BgpAsAService) GetBgpaasSessionAttributes() string {
        return obj.bgpaas_session_attributes
}

func (obj *BgpAsAService) SetBgpaasSessionAttributes(value string) {
        obj.bgpaas_session_attributes = value
        obj.modified |= bgp_as_a_service_bgpaas_session_attributes
}

func (obj *BgpAsAService) GetBgpaasIpv4MappedIpv6Nexthop() bool {
        return obj.bgpaas_ipv4_mapped_ipv6_nexthop
}

func (obj *BgpAsAService) SetBgpaasIpv4MappedIpv6Nexthop(value bool) {
        obj.bgpaas_ipv4_mapped_ipv6_nexthop = value
        obj.modified |= bgp_as_a_service_bgpaas_ipv4_mapped_ipv6_nexthop
}

func (obj *BgpAsAService) GetBgpaasSuppressRouteAdvertisement() bool {
        return obj.bgpaas_suppress_route_advertisement
}

func (obj *BgpAsAService) SetBgpaasSuppressRouteAdvertisement(value bool) {
        obj.bgpaas_suppress_route_advertisement = value
        obj.modified |= bgp_as_a_service_bgpaas_suppress_route_advertisement
}

func (obj *BgpAsAService) GetIdPerms() IdPermsType {
        return obj.id_perms
}

func (obj *BgpAsAService) SetIdPerms(value *IdPermsType) {
        obj.id_perms = *value
        obj.modified |= bgp_as_a_service_id_perms
}

func (obj *BgpAsAService) GetPerms2() PermType2 {
        return obj.perms2
}

func (obj *BgpAsAService) SetPerms2(value *PermType2) {
        obj.perms2 = *value
        obj.modified |= bgp_as_a_service_perms2
}

func (obj *BgpAsAService) GetDisplayName() string {
        return obj.display_name
}

func (obj *BgpAsAService) SetDisplayName(value string) {
        obj.display_name = value
        obj.modified |= bgp_as_a_service_display_name
}

func (obj *BgpAsAService) readVirtualMachineInterfaceRefs() error {
        if !obj.IsTransient() &&
                (obj.valid & bgp_as_a_service_virtual_machine_interface_refs == 0) {
                err := obj.GetField(obj, "virtual_machine_interface_refs")
                if err != nil {
                        return err
                }
        }
        return nil
}

func (obj *BgpAsAService) GetVirtualMachineInterfaceRefs() (
        contrail.ReferenceList, error) {
        err := obj.readVirtualMachineInterfaceRefs()
        if err != nil {
                return nil, err
        }
        return obj.virtual_machine_interface_refs, nil
}

func (obj *BgpAsAService) AddVirtualMachineInterface(
        rhs *VirtualMachineInterface) error {
        err := obj.readVirtualMachineInterfaceRefs()
        if err != nil {
                return err
        }

        if obj.modified & bgp_as_a_service_virtual_machine_interface_refs == 0 {
                obj.storeReferenceBase("virtual-machine-interface", obj.virtual_machine_interface_refs)
        }

        ref := contrail.Reference {
                rhs.GetFQName(), rhs.GetUuid(), rhs.GetHref(), nil}
        obj.virtual_machine_interface_refs = append(obj.virtual_machine_interface_refs, ref)
        obj.modified |= bgp_as_a_service_virtual_machine_interface_refs
        return nil
}

func (obj *BgpAsAService) DeleteVirtualMachineInterface(uuid string) error {
        err := obj.readVirtualMachineInterfaceRefs()
        if err != nil {
                return err
        }

        if obj.modified & bgp_as_a_service_virtual_machine_interface_refs == 0 {
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
        obj.modified |= bgp_as_a_service_virtual_machine_interface_refs
        return nil
}

func (obj *BgpAsAService) ClearVirtualMachineInterface() {
        if (obj.valid & bgp_as_a_service_virtual_machine_interface_refs != 0) &&
           (obj.modified & bgp_as_a_service_virtual_machine_interface_refs == 0) {
                obj.storeReferenceBase("virtual-machine-interface", obj.virtual_machine_interface_refs)
        }
        obj.virtual_machine_interface_refs = make([]contrail.Reference, 0)
        obj.valid |= bgp_as_a_service_virtual_machine_interface_refs
        obj.modified |= bgp_as_a_service_virtual_machine_interface_refs
}

func (obj *BgpAsAService) SetVirtualMachineInterfaceList(
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


func (obj *BgpAsAService) MarshalJSON() ([]byte, error) {
        msg := map[string]*json.RawMessage {
        }
        err := obj.MarshalCommon(msg)
        if err != nil {
                return nil, err
        }

        if obj.modified & bgp_as_a_service_autonomous_system != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.autonomous_system)
                if err != nil {
                        return nil, err
                }
                msg["autonomous_system"] = &value
        }

        if obj.modified & bgp_as_a_service_bgpaas_ip_address != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.bgpaas_ip_address)
                if err != nil {
                        return nil, err
                }
                msg["bgpaas_ip_address"] = &value
        }

        if obj.modified & bgp_as_a_service_bgpaas_session_attributes != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.bgpaas_session_attributes)
                if err != nil {
                        return nil, err
                }
                msg["bgpaas_session_attributes"] = &value
        }

        if obj.modified & bgp_as_a_service_bgpaas_ipv4_mapped_ipv6_nexthop != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.bgpaas_ipv4_mapped_ipv6_nexthop)
                if err != nil {
                        return nil, err
                }
                msg["bgpaas_ipv4_mapped_ipv6_nexthop"] = &value
        }

        if obj.modified & bgp_as_a_service_bgpaas_suppress_route_advertisement != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.bgpaas_suppress_route_advertisement)
                if err != nil {
                        return nil, err
                }
                msg["bgpaas_suppress_route_advertisement"] = &value
        }

        if obj.modified & bgp_as_a_service_id_perms != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.id_perms)
                if err != nil {
                        return nil, err
                }
                msg["id_perms"] = &value
        }

        if obj.modified & bgp_as_a_service_perms2 != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.perms2)
                if err != nil {
                        return nil, err
                }
                msg["perms2"] = &value
        }

        if obj.modified & bgp_as_a_service_display_name != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.display_name)
                if err != nil {
                        return nil, err
                }
                msg["display_name"] = &value
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

func (obj *BgpAsAService) UnmarshalJSON(body []byte) error {
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
                case "autonomous_system":
                        err = json.Unmarshal(value, &obj.autonomous_system)
                        if err == nil {
                                obj.valid |= bgp_as_a_service_autonomous_system
                        }
                        break
                case "bgpaas_ip_address":
                        err = json.Unmarshal(value, &obj.bgpaas_ip_address)
                        if err == nil {
                                obj.valid |= bgp_as_a_service_bgpaas_ip_address
                        }
                        break
                case "bgpaas_session_attributes":
                        err = json.Unmarshal(value, &obj.bgpaas_session_attributes)
                        if err == nil {
                                obj.valid |= bgp_as_a_service_bgpaas_session_attributes
                        }
                        break
                case "bgpaas_ipv4_mapped_ipv6_nexthop":
                        err = json.Unmarshal(value, &obj.bgpaas_ipv4_mapped_ipv6_nexthop)
                        if err == nil {
                                obj.valid |= bgp_as_a_service_bgpaas_ipv4_mapped_ipv6_nexthop
                        }
                        break
                case "bgpaas_suppress_route_advertisement":
                        err = json.Unmarshal(value, &obj.bgpaas_suppress_route_advertisement)
                        if err == nil {
                                obj.valid |= bgp_as_a_service_bgpaas_suppress_route_advertisement
                        }
                        break
                case "id_perms":
                        err = json.Unmarshal(value, &obj.id_perms)
                        if err == nil {
                                obj.valid |= bgp_as_a_service_id_perms
                        }
                        break
                case "perms2":
                        err = json.Unmarshal(value, &obj.perms2)
                        if err == nil {
                                obj.valid |= bgp_as_a_service_perms2
                        }
                        break
                case "display_name":
                        err = json.Unmarshal(value, &obj.display_name)
                        if err == nil {
                                obj.valid |= bgp_as_a_service_display_name
                        }
                        break
                case "virtual_machine_interface_refs":
                        err = json.Unmarshal(value, &obj.virtual_machine_interface_refs)
                        if err == nil {
                                obj.valid |= bgp_as_a_service_virtual_machine_interface_refs
                        }
                        break
                }
                if err != nil {
                        return err
                }
        }
        return nil
}

func (obj *BgpAsAService) UpdateObject() ([]byte, error) {
        msg := map[string]*json.RawMessage {
        }
        err := obj.MarshalId(msg)
        if err != nil {
                return nil, err
        }

        if obj.modified & bgp_as_a_service_autonomous_system != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.autonomous_system)
                if err != nil {
                        return nil, err
                }
                msg["autonomous_system"] = &value
        }

        if obj.modified & bgp_as_a_service_bgpaas_ip_address != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.bgpaas_ip_address)
                if err != nil {
                        return nil, err
                }
                msg["bgpaas_ip_address"] = &value
        }

        if obj.modified & bgp_as_a_service_bgpaas_session_attributes != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.bgpaas_session_attributes)
                if err != nil {
                        return nil, err
                }
                msg["bgpaas_session_attributes"] = &value
        }

        if obj.modified & bgp_as_a_service_bgpaas_ipv4_mapped_ipv6_nexthop != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.bgpaas_ipv4_mapped_ipv6_nexthop)
                if err != nil {
                        return nil, err
                }
                msg["bgpaas_ipv4_mapped_ipv6_nexthop"] = &value
        }

        if obj.modified & bgp_as_a_service_bgpaas_suppress_route_advertisement != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.bgpaas_suppress_route_advertisement)
                if err != nil {
                        return nil, err
                }
                msg["bgpaas_suppress_route_advertisement"] = &value
        }

        if obj.modified & bgp_as_a_service_id_perms != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.id_perms)
                if err != nil {
                        return nil, err
                }
                msg["id_perms"] = &value
        }

        if obj.modified & bgp_as_a_service_perms2 != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.perms2)
                if err != nil {
                        return nil, err
                }
                msg["perms2"] = &value
        }

        if obj.modified & bgp_as_a_service_display_name != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.display_name)
                if err != nil {
                        return nil, err
                }
                msg["display_name"] = &value
        }

        if obj.modified & bgp_as_a_service_virtual_machine_interface_refs != 0 {
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

func (obj *BgpAsAService) UpdateReferences() error {

        if (obj.modified & bgp_as_a_service_virtual_machine_interface_refs != 0) &&
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

func BgpAsAServiceByName(c contrail.ApiClient, fqn string) (*BgpAsAService, error) {
    obj, err := c.FindByName("bgp-as-a-service", fqn)
    if err != nil {
        return nil, err
    }
    return obj.(*BgpAsAService), nil
}

func BgpAsAServiceByUuid(c contrail.ApiClient, uuid string) (*BgpAsAService, error) {
    obj, err := c.FindByUuid("bgp-as-a-service", uuid)
    if err != nil {
        return nil, err
    }
    return obj.(*BgpAsAService), nil
}
