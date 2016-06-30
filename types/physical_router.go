//
// Automatically generated. DO NOT EDIT.
//

package types

import (
        "encoding/json"

        "github.com/Juniper/contrail-go-api"
)

const (
	physical_router_physical_router_management_ip uint64 = 1 << iota
	physical_router_physical_router_dataplane_ip
	physical_router_physical_router_vendor_name
	physical_router_physical_router_product_name
	physical_router_physical_router_vnc_managed
	physical_router_physical_router_user_credentials
	physical_router_physical_router_snmp_credentials
	physical_router_physical_router_junos_service_ports
	physical_router_id_perms
	physical_router_perms2
	physical_router_display_name
	physical_router_virtual_router_refs
	physical_router_bgp_router_refs
	physical_router_virtual_network_refs
	physical_router_physical_interfaces
	physical_router_logical_interfaces
	physical_router_instance_ip_back_refs
)

type PhysicalRouter struct {
        contrail.ObjectBase
	physical_router_management_ip string
	physical_router_dataplane_ip string
	physical_router_vendor_name string
	physical_router_product_name string
	physical_router_vnc_managed bool
	physical_router_user_credentials UserCredentials
	physical_router_snmp_credentials SNMPCredentials
	physical_router_junos_service_ports JunosServicePorts
	id_perms IdPermsType
	perms2 PermType2
	display_name string
	virtual_router_refs contrail.ReferenceList
	bgp_router_refs contrail.ReferenceList
	virtual_network_refs contrail.ReferenceList
	physical_interfaces contrail.ReferenceList
	logical_interfaces contrail.ReferenceList
	instance_ip_back_refs contrail.ReferenceList
        valid uint64
        modified uint64
        baseMap map[string]contrail.ReferenceList
}

func (obj *PhysicalRouter) GetType() string {
        return "physical-router"
}

func (obj *PhysicalRouter) GetDefaultParent() []string {
        name := []string{"default-global-system-config"}
        return name
}

func (obj *PhysicalRouter) GetDefaultParentType() string {
        return "global-system-config"
}

func (obj *PhysicalRouter) SetName(name string) {
        obj.VSetName(obj, name)
}

func (obj *PhysicalRouter) SetParent(parent contrail.IObject) {
        obj.VSetParent(obj, parent)
}

func (obj *PhysicalRouter) storeReferenceBase(
        name string, refList contrail.ReferenceList) {
        if obj.baseMap == nil {
                obj.baseMap = make(map[string]contrail.ReferenceList)
        }
        refCopy := make(contrail.ReferenceList, len(refList))
        copy(refCopy, refList)
        obj.baseMap[name] = refCopy
}

func (obj *PhysicalRouter) hasReferenceBase(name string) bool {
        if obj.baseMap == nil {
                return false
        }
        _, exists := obj.baseMap[name]
        return exists
}

func (obj *PhysicalRouter) UpdateDone() {
        obj.modified = 0
        obj.baseMap = nil
}


func (obj *PhysicalRouter) GetPhysicalRouterManagementIp() string {
        return obj.physical_router_management_ip
}

func (obj *PhysicalRouter) SetPhysicalRouterManagementIp(value string) {
        obj.physical_router_management_ip = value
        obj.modified |= physical_router_physical_router_management_ip
}

func (obj *PhysicalRouter) GetPhysicalRouterDataplaneIp() string {
        return obj.physical_router_dataplane_ip
}

func (obj *PhysicalRouter) SetPhysicalRouterDataplaneIp(value string) {
        obj.physical_router_dataplane_ip = value
        obj.modified |= physical_router_physical_router_dataplane_ip
}

func (obj *PhysicalRouter) GetPhysicalRouterVendorName() string {
        return obj.physical_router_vendor_name
}

func (obj *PhysicalRouter) SetPhysicalRouterVendorName(value string) {
        obj.physical_router_vendor_name = value
        obj.modified |= physical_router_physical_router_vendor_name
}

func (obj *PhysicalRouter) GetPhysicalRouterProductName() string {
        return obj.physical_router_product_name
}

func (obj *PhysicalRouter) SetPhysicalRouterProductName(value string) {
        obj.physical_router_product_name = value
        obj.modified |= physical_router_physical_router_product_name
}

func (obj *PhysicalRouter) GetPhysicalRouterVncManaged() bool {
        return obj.physical_router_vnc_managed
}

func (obj *PhysicalRouter) SetPhysicalRouterVncManaged(value bool) {
        obj.physical_router_vnc_managed = value
        obj.modified |= physical_router_physical_router_vnc_managed
}

func (obj *PhysicalRouter) GetPhysicalRouterUserCredentials() UserCredentials {
        return obj.physical_router_user_credentials
}

func (obj *PhysicalRouter) SetPhysicalRouterUserCredentials(value *UserCredentials) {
        obj.physical_router_user_credentials = *value
        obj.modified |= physical_router_physical_router_user_credentials
}

func (obj *PhysicalRouter) GetPhysicalRouterSnmpCredentials() SNMPCredentials {
        return obj.physical_router_snmp_credentials
}

func (obj *PhysicalRouter) SetPhysicalRouterSnmpCredentials(value *SNMPCredentials) {
        obj.physical_router_snmp_credentials = *value
        obj.modified |= physical_router_physical_router_snmp_credentials
}

func (obj *PhysicalRouter) GetPhysicalRouterJunosServicePorts() JunosServicePorts {
        return obj.physical_router_junos_service_ports
}

func (obj *PhysicalRouter) SetPhysicalRouterJunosServicePorts(value *JunosServicePorts) {
        obj.physical_router_junos_service_ports = *value
        obj.modified |= physical_router_physical_router_junos_service_ports
}

func (obj *PhysicalRouter) GetIdPerms() IdPermsType {
        return obj.id_perms
}

func (obj *PhysicalRouter) SetIdPerms(value *IdPermsType) {
        obj.id_perms = *value
        obj.modified |= physical_router_id_perms
}

func (obj *PhysicalRouter) GetPerms2() PermType2 {
        return obj.perms2
}

func (obj *PhysicalRouter) SetPerms2(value *PermType2) {
        obj.perms2 = *value
        obj.modified |= physical_router_perms2
}

func (obj *PhysicalRouter) GetDisplayName() string {
        return obj.display_name
}

func (obj *PhysicalRouter) SetDisplayName(value string) {
        obj.display_name = value
        obj.modified |= physical_router_display_name
}

func (obj *PhysicalRouter) readPhysicalInterfaces() error {
        if !obj.IsTransient() &&
                (obj.valid & physical_router_physical_interfaces == 0) {
                err := obj.GetField(obj, "physical_interfaces")
                if err != nil {
                        return err
                }
        }
        return nil
}

func (obj *PhysicalRouter) GetPhysicalInterfaces() (
        contrail.ReferenceList, error) {
        err := obj.readPhysicalInterfaces()
        if err != nil {
                return nil, err
        }
        return obj.physical_interfaces, nil
}

func (obj *PhysicalRouter) readLogicalInterfaces() error {
        if !obj.IsTransient() &&
                (obj.valid & physical_router_logical_interfaces == 0) {
                err := obj.GetField(obj, "logical_interfaces")
                if err != nil {
                        return err
                }
        }
        return nil
}

func (obj *PhysicalRouter) GetLogicalInterfaces() (
        contrail.ReferenceList, error) {
        err := obj.readLogicalInterfaces()
        if err != nil {
                return nil, err
        }
        return obj.logical_interfaces, nil
}

func (obj *PhysicalRouter) readVirtualRouterRefs() error {
        if !obj.IsTransient() &&
                (obj.valid & physical_router_virtual_router_refs == 0) {
                err := obj.GetField(obj, "virtual_router_refs")
                if err != nil {
                        return err
                }
        }
        return nil
}

func (obj *PhysicalRouter) GetVirtualRouterRefs() (
        contrail.ReferenceList, error) {
        err := obj.readVirtualRouterRefs()
        if err != nil {
                return nil, err
        }
        return obj.virtual_router_refs, nil
}

func (obj *PhysicalRouter) AddVirtualRouter(
        rhs *VirtualRouter) error {
        err := obj.readVirtualRouterRefs()
        if err != nil {
                return err
        }

        if obj.modified & physical_router_virtual_router_refs == 0 {
                obj.storeReferenceBase("virtual-router", obj.virtual_router_refs)
        }

        ref := contrail.Reference {
                rhs.GetFQName(), rhs.GetUuid(), rhs.GetHref(), nil}
        obj.virtual_router_refs = append(obj.virtual_router_refs, ref)
        obj.modified |= physical_router_virtual_router_refs
        return nil
}

func (obj *PhysicalRouter) DeleteVirtualRouter(uuid string) error {
        err := obj.readVirtualRouterRefs()
        if err != nil {
                return err
        }

        if obj.modified & physical_router_virtual_router_refs == 0 {
                obj.storeReferenceBase("virtual-router", obj.virtual_router_refs)
        }

        for i, ref := range obj.virtual_router_refs {
                if ref.Uuid == uuid {
                        obj.virtual_router_refs = append(
                                obj.virtual_router_refs[:i],
                                obj.virtual_router_refs[i+1:]...)
                        break
                }
        }
        obj.modified |= physical_router_virtual_router_refs
        return nil
}

func (obj *PhysicalRouter) ClearVirtualRouter() {
        if (obj.valid & physical_router_virtual_router_refs != 0) &&
           (obj.modified & physical_router_virtual_router_refs == 0) {
                obj.storeReferenceBase("virtual-router", obj.virtual_router_refs)
        }
        obj.virtual_router_refs = make([]contrail.Reference, 0)
        obj.valid |= physical_router_virtual_router_refs
        obj.modified |= physical_router_virtual_router_refs
}

func (obj *PhysicalRouter) SetVirtualRouterList(
        refList []contrail.ReferencePair) {
        obj.ClearVirtualRouter()
        obj.virtual_router_refs = make([]contrail.Reference, len(refList))
        for i, pair := range refList {
                obj.virtual_router_refs[i] = contrail.Reference {
                        pair.Object.GetFQName(),
                        pair.Object.GetUuid(),
                        pair.Object.GetHref(),
                        pair.Attribute,
                }
        }
}


func (obj *PhysicalRouter) readBgpRouterRefs() error {
        if !obj.IsTransient() &&
                (obj.valid & physical_router_bgp_router_refs == 0) {
                err := obj.GetField(obj, "bgp_router_refs")
                if err != nil {
                        return err
                }
        }
        return nil
}

func (obj *PhysicalRouter) GetBgpRouterRefs() (
        contrail.ReferenceList, error) {
        err := obj.readBgpRouterRefs()
        if err != nil {
                return nil, err
        }
        return obj.bgp_router_refs, nil
}

func (obj *PhysicalRouter) AddBgpRouter(
        rhs *BgpRouter) error {
        err := obj.readBgpRouterRefs()
        if err != nil {
                return err
        }

        if obj.modified & physical_router_bgp_router_refs == 0 {
                obj.storeReferenceBase("bgp-router", obj.bgp_router_refs)
        }

        ref := contrail.Reference {
                rhs.GetFQName(), rhs.GetUuid(), rhs.GetHref(), nil}
        obj.bgp_router_refs = append(obj.bgp_router_refs, ref)
        obj.modified |= physical_router_bgp_router_refs
        return nil
}

func (obj *PhysicalRouter) DeleteBgpRouter(uuid string) error {
        err := obj.readBgpRouterRefs()
        if err != nil {
                return err
        }

        if obj.modified & physical_router_bgp_router_refs == 0 {
                obj.storeReferenceBase("bgp-router", obj.bgp_router_refs)
        }

        for i, ref := range obj.bgp_router_refs {
                if ref.Uuid == uuid {
                        obj.bgp_router_refs = append(
                                obj.bgp_router_refs[:i],
                                obj.bgp_router_refs[i+1:]...)
                        break
                }
        }
        obj.modified |= physical_router_bgp_router_refs
        return nil
}

func (obj *PhysicalRouter) ClearBgpRouter() {
        if (obj.valid & physical_router_bgp_router_refs != 0) &&
           (obj.modified & physical_router_bgp_router_refs == 0) {
                obj.storeReferenceBase("bgp-router", obj.bgp_router_refs)
        }
        obj.bgp_router_refs = make([]contrail.Reference, 0)
        obj.valid |= physical_router_bgp_router_refs
        obj.modified |= physical_router_bgp_router_refs
}

func (obj *PhysicalRouter) SetBgpRouterList(
        refList []contrail.ReferencePair) {
        obj.ClearBgpRouter()
        obj.bgp_router_refs = make([]contrail.Reference, len(refList))
        for i, pair := range refList {
                obj.bgp_router_refs[i] = contrail.Reference {
                        pair.Object.GetFQName(),
                        pair.Object.GetUuid(),
                        pair.Object.GetHref(),
                        pair.Attribute,
                }
        }
}


func (obj *PhysicalRouter) readVirtualNetworkRefs() error {
        if !obj.IsTransient() &&
                (obj.valid & physical_router_virtual_network_refs == 0) {
                err := obj.GetField(obj, "virtual_network_refs")
                if err != nil {
                        return err
                }
        }
        return nil
}

func (obj *PhysicalRouter) GetVirtualNetworkRefs() (
        contrail.ReferenceList, error) {
        err := obj.readVirtualNetworkRefs()
        if err != nil {
                return nil, err
        }
        return obj.virtual_network_refs, nil
}

func (obj *PhysicalRouter) AddVirtualNetwork(
        rhs *VirtualNetwork) error {
        err := obj.readVirtualNetworkRefs()
        if err != nil {
                return err
        }

        if obj.modified & physical_router_virtual_network_refs == 0 {
                obj.storeReferenceBase("virtual-network", obj.virtual_network_refs)
        }

        ref := contrail.Reference {
                rhs.GetFQName(), rhs.GetUuid(), rhs.GetHref(), nil}
        obj.virtual_network_refs = append(obj.virtual_network_refs, ref)
        obj.modified |= physical_router_virtual_network_refs
        return nil
}

func (obj *PhysicalRouter) DeleteVirtualNetwork(uuid string) error {
        err := obj.readVirtualNetworkRefs()
        if err != nil {
                return err
        }

        if obj.modified & physical_router_virtual_network_refs == 0 {
                obj.storeReferenceBase("virtual-network", obj.virtual_network_refs)
        }

        for i, ref := range obj.virtual_network_refs {
                if ref.Uuid == uuid {
                        obj.virtual_network_refs = append(
                                obj.virtual_network_refs[:i],
                                obj.virtual_network_refs[i+1:]...)
                        break
                }
        }
        obj.modified |= physical_router_virtual_network_refs
        return nil
}

func (obj *PhysicalRouter) ClearVirtualNetwork() {
        if (obj.valid & physical_router_virtual_network_refs != 0) &&
           (obj.modified & physical_router_virtual_network_refs == 0) {
                obj.storeReferenceBase("virtual-network", obj.virtual_network_refs)
        }
        obj.virtual_network_refs = make([]contrail.Reference, 0)
        obj.valid |= physical_router_virtual_network_refs
        obj.modified |= physical_router_virtual_network_refs
}

func (obj *PhysicalRouter) SetVirtualNetworkList(
        refList []contrail.ReferencePair) {
        obj.ClearVirtualNetwork()
        obj.virtual_network_refs = make([]contrail.Reference, len(refList))
        for i, pair := range refList {
                obj.virtual_network_refs[i] = contrail.Reference {
                        pair.Object.GetFQName(),
                        pair.Object.GetUuid(),
                        pair.Object.GetHref(),
                        pair.Attribute,
                }
        }
}


func (obj *PhysicalRouter) readInstanceIpBackRefs() error {
        if !obj.IsTransient() &&
                (obj.valid & physical_router_instance_ip_back_refs == 0) {
                err := obj.GetField(obj, "instance_ip_back_refs")
                if err != nil {
                        return err
                }
        }
        return nil
}

func (obj *PhysicalRouter) GetInstanceIpBackRefs() (
        contrail.ReferenceList, error) {
        err := obj.readInstanceIpBackRefs()
        if err != nil {
                return nil, err
        }
        return obj.instance_ip_back_refs, nil
}

func (obj *PhysicalRouter) MarshalJSON() ([]byte, error) {
        msg := map[string]*json.RawMessage {
        }
        err := obj.MarshalCommon(msg)
        if err != nil {
                return nil, err
        }

        if obj.modified & physical_router_physical_router_management_ip != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.physical_router_management_ip)
                if err != nil {
                        return nil, err
                }
                msg["physical_router_management_ip"] = &value
        }

        if obj.modified & physical_router_physical_router_dataplane_ip != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.physical_router_dataplane_ip)
                if err != nil {
                        return nil, err
                }
                msg["physical_router_dataplane_ip"] = &value
        }

        if obj.modified & physical_router_physical_router_vendor_name != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.physical_router_vendor_name)
                if err != nil {
                        return nil, err
                }
                msg["physical_router_vendor_name"] = &value
        }

        if obj.modified & physical_router_physical_router_product_name != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.physical_router_product_name)
                if err != nil {
                        return nil, err
                }
                msg["physical_router_product_name"] = &value
        }

        if obj.modified & physical_router_physical_router_vnc_managed != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.physical_router_vnc_managed)
                if err != nil {
                        return nil, err
                }
                msg["physical_router_vnc_managed"] = &value
        }

        if obj.modified & physical_router_physical_router_user_credentials != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.physical_router_user_credentials)
                if err != nil {
                        return nil, err
                }
                msg["physical_router_user_credentials"] = &value
        }

        if obj.modified & physical_router_physical_router_snmp_credentials != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.physical_router_snmp_credentials)
                if err != nil {
                        return nil, err
                }
                msg["physical_router_snmp_credentials"] = &value
        }

        if obj.modified & physical_router_physical_router_junos_service_ports != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.physical_router_junos_service_ports)
                if err != nil {
                        return nil, err
                }
                msg["physical_router_junos_service_ports"] = &value
        }

        if obj.modified & physical_router_id_perms != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.id_perms)
                if err != nil {
                        return nil, err
                }
                msg["id_perms"] = &value
        }

        if obj.modified & physical_router_perms2 != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.perms2)
                if err != nil {
                        return nil, err
                }
                msg["perms2"] = &value
        }

        if obj.modified & physical_router_display_name != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.display_name)
                if err != nil {
                        return nil, err
                }
                msg["display_name"] = &value
        }

        if len(obj.virtual_router_refs) > 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.virtual_router_refs)
                if err != nil {
                        return nil, err
                }
                msg["virtual_router_refs"] = &value
        }

        if len(obj.bgp_router_refs) > 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.bgp_router_refs)
                if err != nil {
                        return nil, err
                }
                msg["bgp_router_refs"] = &value
        }

        if len(obj.virtual_network_refs) > 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.virtual_network_refs)
                if err != nil {
                        return nil, err
                }
                msg["virtual_network_refs"] = &value
        }

        return json.Marshal(msg)
}

func (obj *PhysicalRouter) UnmarshalJSON(body []byte) error {
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
                case "physical_router_management_ip":
                        err = json.Unmarshal(value, &obj.physical_router_management_ip)
                        if err == nil {
                                obj.valid |= physical_router_physical_router_management_ip
                        }
                        break
                case "physical_router_dataplane_ip":
                        err = json.Unmarshal(value, &obj.physical_router_dataplane_ip)
                        if err == nil {
                                obj.valid |= physical_router_physical_router_dataplane_ip
                        }
                        break
                case "physical_router_vendor_name":
                        err = json.Unmarshal(value, &obj.physical_router_vendor_name)
                        if err == nil {
                                obj.valid |= physical_router_physical_router_vendor_name
                        }
                        break
                case "physical_router_product_name":
                        err = json.Unmarshal(value, &obj.physical_router_product_name)
                        if err == nil {
                                obj.valid |= physical_router_physical_router_product_name
                        }
                        break
                case "physical_router_vnc_managed":
                        err = json.Unmarshal(value, &obj.physical_router_vnc_managed)
                        if err == nil {
                                obj.valid |= physical_router_physical_router_vnc_managed
                        }
                        break
                case "physical_router_user_credentials":
                        err = json.Unmarshal(value, &obj.physical_router_user_credentials)
                        if err == nil {
                                obj.valid |= physical_router_physical_router_user_credentials
                        }
                        break
                case "physical_router_snmp_credentials":
                        err = json.Unmarshal(value, &obj.physical_router_snmp_credentials)
                        if err == nil {
                                obj.valid |= physical_router_physical_router_snmp_credentials
                        }
                        break
                case "physical_router_junos_service_ports":
                        err = json.Unmarshal(value, &obj.physical_router_junos_service_ports)
                        if err == nil {
                                obj.valid |= physical_router_physical_router_junos_service_ports
                        }
                        break
                case "id_perms":
                        err = json.Unmarshal(value, &obj.id_perms)
                        if err == nil {
                                obj.valid |= physical_router_id_perms
                        }
                        break
                case "perms2":
                        err = json.Unmarshal(value, &obj.perms2)
                        if err == nil {
                                obj.valid |= physical_router_perms2
                        }
                        break
                case "display_name":
                        err = json.Unmarshal(value, &obj.display_name)
                        if err == nil {
                                obj.valid |= physical_router_display_name
                        }
                        break
                case "virtual_router_refs":
                        err = json.Unmarshal(value, &obj.virtual_router_refs)
                        if err == nil {
                                obj.valid |= physical_router_virtual_router_refs
                        }
                        break
                case "bgp_router_refs":
                        err = json.Unmarshal(value, &obj.bgp_router_refs)
                        if err == nil {
                                obj.valid |= physical_router_bgp_router_refs
                        }
                        break
                case "virtual_network_refs":
                        err = json.Unmarshal(value, &obj.virtual_network_refs)
                        if err == nil {
                                obj.valid |= physical_router_virtual_network_refs
                        }
                        break
                case "physical_interfaces":
                        err = json.Unmarshal(value, &obj.physical_interfaces)
                        if err == nil {
                                obj.valid |= physical_router_physical_interfaces
                        }
                        break
                case "logical_interfaces":
                        err = json.Unmarshal(value, &obj.logical_interfaces)
                        if err == nil {
                                obj.valid |= physical_router_logical_interfaces
                        }
                        break
                case "instance_ip_back_refs":
                        err = json.Unmarshal(value, &obj.instance_ip_back_refs)
                        if err == nil {
                                obj.valid |= physical_router_instance_ip_back_refs
                        }
                        break
                }
                if err != nil {
                        return err
                }
        }
        return nil
}

func (obj *PhysicalRouter) UpdateObject() ([]byte, error) {
        msg := map[string]*json.RawMessage {
        }
        err := obj.MarshalId(msg)
        if err != nil {
                return nil, err
        }

        if obj.modified & physical_router_physical_router_management_ip != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.physical_router_management_ip)
                if err != nil {
                        return nil, err
                }
                msg["physical_router_management_ip"] = &value
        }

        if obj.modified & physical_router_physical_router_dataplane_ip != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.physical_router_dataplane_ip)
                if err != nil {
                        return nil, err
                }
                msg["physical_router_dataplane_ip"] = &value
        }

        if obj.modified & physical_router_physical_router_vendor_name != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.physical_router_vendor_name)
                if err != nil {
                        return nil, err
                }
                msg["physical_router_vendor_name"] = &value
        }

        if obj.modified & physical_router_physical_router_product_name != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.physical_router_product_name)
                if err != nil {
                        return nil, err
                }
                msg["physical_router_product_name"] = &value
        }

        if obj.modified & physical_router_physical_router_vnc_managed != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.physical_router_vnc_managed)
                if err != nil {
                        return nil, err
                }
                msg["physical_router_vnc_managed"] = &value
        }

        if obj.modified & physical_router_physical_router_user_credentials != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.physical_router_user_credentials)
                if err != nil {
                        return nil, err
                }
                msg["physical_router_user_credentials"] = &value
        }

        if obj.modified & physical_router_physical_router_snmp_credentials != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.physical_router_snmp_credentials)
                if err != nil {
                        return nil, err
                }
                msg["physical_router_snmp_credentials"] = &value
        }

        if obj.modified & physical_router_physical_router_junos_service_ports != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.physical_router_junos_service_ports)
                if err != nil {
                        return nil, err
                }
                msg["physical_router_junos_service_ports"] = &value
        }

        if obj.modified & physical_router_id_perms != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.id_perms)
                if err != nil {
                        return nil, err
                }
                msg["id_perms"] = &value
        }

        if obj.modified & physical_router_perms2 != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.perms2)
                if err != nil {
                        return nil, err
                }
                msg["perms2"] = &value
        }

        if obj.modified & physical_router_display_name != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.display_name)
                if err != nil {
                        return nil, err
                }
                msg["display_name"] = &value
        }

        if obj.modified & physical_router_virtual_router_refs != 0 {
                if len(obj.virtual_router_refs) == 0 {
                        var value json.RawMessage
                        value, err := json.Marshal(
                                          make([]contrail.Reference, 0))
                        if err != nil {
                                return nil, err
                        }
                        msg["virtual_router_refs"] = &value
                } else if !obj.hasReferenceBase("virtual-router") {
                        var value json.RawMessage
                        value, err := json.Marshal(&obj.virtual_router_refs)
                        if err != nil {
                                return nil, err
                        }
                        msg["virtual_router_refs"] = &value
                }
        }


        if obj.modified & physical_router_bgp_router_refs != 0 {
                if len(obj.bgp_router_refs) == 0 {
                        var value json.RawMessage
                        value, err := json.Marshal(
                                          make([]contrail.Reference, 0))
                        if err != nil {
                                return nil, err
                        }
                        msg["bgp_router_refs"] = &value
                } else if !obj.hasReferenceBase("bgp-router") {
                        var value json.RawMessage
                        value, err := json.Marshal(&obj.bgp_router_refs)
                        if err != nil {
                                return nil, err
                        }
                        msg["bgp_router_refs"] = &value
                }
        }


        if obj.modified & physical_router_virtual_network_refs != 0 {
                if len(obj.virtual_network_refs) == 0 {
                        var value json.RawMessage
                        value, err := json.Marshal(
                                          make([]contrail.Reference, 0))
                        if err != nil {
                                return nil, err
                        }
                        msg["virtual_network_refs"] = &value
                } else if !obj.hasReferenceBase("virtual-network") {
                        var value json.RawMessage
                        value, err := json.Marshal(&obj.virtual_network_refs)
                        if err != nil {
                                return nil, err
                        }
                        msg["virtual_network_refs"] = &value
                }
        }


        return json.Marshal(msg)
}

func (obj *PhysicalRouter) UpdateReferences() error {

        if (obj.modified & physical_router_virtual_router_refs != 0) &&
           len(obj.virtual_router_refs) > 0 &&
           obj.hasReferenceBase("virtual-router") {
                err := obj.UpdateReference(
                        obj, "virtual-router",
                        obj.virtual_router_refs,
                        obj.baseMap["virtual-router"])
                if err != nil {
                        return err
                }
        }

        if (obj.modified & physical_router_bgp_router_refs != 0) &&
           len(obj.bgp_router_refs) > 0 &&
           obj.hasReferenceBase("bgp-router") {
                err := obj.UpdateReference(
                        obj, "bgp-router",
                        obj.bgp_router_refs,
                        obj.baseMap["bgp-router"])
                if err != nil {
                        return err
                }
        }

        if (obj.modified & physical_router_virtual_network_refs != 0) &&
           len(obj.virtual_network_refs) > 0 &&
           obj.hasReferenceBase("virtual-network") {
                err := obj.UpdateReference(
                        obj, "virtual-network",
                        obj.virtual_network_refs,
                        obj.baseMap["virtual-network"])
                if err != nil {
                        return err
                }
        }

        return nil
}

func PhysicalRouterByName(c contrail.ApiClient, fqn string) (*PhysicalRouter, error) {
    obj, err := c.FindByName("physical-router", fqn)
    if err != nil {
        return nil, err
    }
    return obj.(*PhysicalRouter), nil
}

func PhysicalRouterByUuid(c contrail.ApiClient, uuid string) (*PhysicalRouter, error) {
    obj, err := c.FindByUuid("physical-router", uuid)
    if err != nil {
        return nil, err
    }
    return obj.(*PhysicalRouter), nil
}
