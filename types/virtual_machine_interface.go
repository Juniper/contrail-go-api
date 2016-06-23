//
// Automatically generated. DO NOT EDIT.
//

package types

import (
        "encoding/json"

        "github.com/michaelhenkel/contrail-go-api"
)

const (
	virtual_machine_interface_ecmp_hashing_include_fields uint64 = 1 << iota
	virtual_machine_interface_virtual_machine_interface_mac_addresses
	virtual_machine_interface_virtual_machine_interface_dhcp_option_list
	virtual_machine_interface_virtual_machine_interface_host_routes
	virtual_machine_interface_virtual_machine_interface_allowed_address_pairs
	virtual_machine_interface_vrf_assign_table
	virtual_machine_interface_virtual_machine_interface_device_owner
	virtual_machine_interface_virtual_machine_interface_disable_policy
	virtual_machine_interface_virtual_machine_interface_properties
	virtual_machine_interface_virtual_machine_interface_bindings
	virtual_machine_interface_virtual_machine_interface_fat_flow_protocols
	virtual_machine_interface_id_perms
	virtual_machine_interface_perms2
	virtual_machine_interface_display_name
	virtual_machine_interface_qos_config_refs
	virtual_machine_interface_security_group_refs
	virtual_machine_interface_virtual_machine_interface_refs
	virtual_machine_interface_virtual_machine_refs
	virtual_machine_interface_virtual_network_refs
	virtual_machine_interface_routing_instance_refs
	virtual_machine_interface_port_tuple_refs
	virtual_machine_interface_service_health_check_refs
	virtual_machine_interface_interface_route_table_refs
	virtual_machine_interface_physical_interface_refs
	virtual_machine_interface_virtual_machine_interface_back_refs
	virtual_machine_interface_instance_ip_back_refs
	virtual_machine_interface_subnet_back_refs
	virtual_machine_interface_floating_ip_back_refs
	virtual_machine_interface_alias_ip_back_refs
	virtual_machine_interface_logical_interface_back_refs
	virtual_machine_interface_bgp_as_a_service_back_refs
	virtual_machine_interface_customer_attachment_back_refs
	virtual_machine_interface_logical_router_back_refs
	virtual_machine_interface_loadbalancer_pool_back_refs
	virtual_machine_interface_virtual_ip_back_refs
	virtual_machine_interface_loadbalancer_back_refs
)

type VirtualMachineInterface struct {
        contrail.ObjectBase
	ecmp_hashing_include_fields EcmpHashingIncludeFields
	virtual_machine_interface_mac_addresses MacAddressesType
	virtual_machine_interface_dhcp_option_list DhcpOptionsListType
	virtual_machine_interface_host_routes RouteTableType
	virtual_machine_interface_allowed_address_pairs AllowedAddressPairs
	vrf_assign_table VrfAssignTableType
	virtual_machine_interface_device_owner string
	virtual_machine_interface_disable_policy bool
	virtual_machine_interface_properties VirtualMachineInterfacePropertiesType
	virtual_machine_interface_bindings KeyValuePairs
	virtual_machine_interface_fat_flow_protocols FatFlowProtocols
	id_perms IdPermsType
	perms2 PermType2
	display_name string
	qos_config_refs contrail.ReferenceList
	security_group_refs contrail.ReferenceList
	virtual_machine_interface_refs contrail.ReferenceList
	virtual_machine_refs contrail.ReferenceList
	virtual_network_refs contrail.ReferenceList
	routing_instance_refs contrail.ReferenceList
	port_tuple_refs contrail.ReferenceList
	service_health_check_refs contrail.ReferenceList
	interface_route_table_refs contrail.ReferenceList
	physical_interface_refs contrail.ReferenceList
	virtual_machine_interface_back_refs contrail.ReferenceList
	instance_ip_back_refs contrail.ReferenceList
	subnet_back_refs contrail.ReferenceList
	floating_ip_back_refs contrail.ReferenceList
	alias_ip_back_refs contrail.ReferenceList
	logical_interface_back_refs contrail.ReferenceList
	bgp_as_a_service_back_refs contrail.ReferenceList
	customer_attachment_back_refs contrail.ReferenceList
	logical_router_back_refs contrail.ReferenceList
	loadbalancer_pool_back_refs contrail.ReferenceList
	virtual_ip_back_refs contrail.ReferenceList
	loadbalancer_back_refs contrail.ReferenceList
        valid uint64
        modified uint64
        baseMap map[string]contrail.ReferenceList
}

func (obj *VirtualMachineInterface) GetType() string {
        return "virtual-machine-interface"
}

func (obj *VirtualMachineInterface) GetDefaultParent() []string {
        name := []string{"default-virtual-machine"}
        return name
}

func (obj *VirtualMachineInterface) GetDefaultParentType() string {
        return "virtual-machine"
}

func (obj *VirtualMachineInterface) SetName(name string) {
        obj.VSetName(obj, name)
}

func (obj *VirtualMachineInterface) SetParent(parent contrail.IObject) {
        obj.VSetParent(obj, parent)
}

func (obj *VirtualMachineInterface) storeReferenceBase(
        name string, refList contrail.ReferenceList) {
        if obj.baseMap == nil {
                obj.baseMap = make(map[string]contrail.ReferenceList)
        }
        refCopy := make(contrail.ReferenceList, len(refList))
        copy(refCopy, refList)
        obj.baseMap[name] = refCopy
}

func (obj *VirtualMachineInterface) hasReferenceBase(name string) bool {
        if obj.baseMap == nil {
                return false
        }
        _, exists := obj.baseMap[name]
        return exists
}

func (obj *VirtualMachineInterface) UpdateDone() {
        obj.modified = 0
        obj.baseMap = nil
}


func (obj *VirtualMachineInterface) GetEcmpHashingIncludeFields() EcmpHashingIncludeFields {
        return obj.ecmp_hashing_include_fields
}

func (obj *VirtualMachineInterface) SetEcmpHashingIncludeFields(value *EcmpHashingIncludeFields) {
        obj.ecmp_hashing_include_fields = *value
        obj.modified |= virtual_machine_interface_ecmp_hashing_include_fields
}

func (obj *VirtualMachineInterface) GetVirtualMachineInterfaceMacAddresses() MacAddressesType {
        return obj.virtual_machine_interface_mac_addresses
}

func (obj *VirtualMachineInterface) SetVirtualMachineInterfaceMacAddresses(value *MacAddressesType) {
        obj.virtual_machine_interface_mac_addresses = *value
        obj.modified |= virtual_machine_interface_virtual_machine_interface_mac_addresses
}

func (obj *VirtualMachineInterface) GetVirtualMachineInterfaceDhcpOptionList() DhcpOptionsListType {
        return obj.virtual_machine_interface_dhcp_option_list
}

func (obj *VirtualMachineInterface) SetVirtualMachineInterfaceDhcpOptionList(value *DhcpOptionsListType) {
        obj.virtual_machine_interface_dhcp_option_list = *value
        obj.modified |= virtual_machine_interface_virtual_machine_interface_dhcp_option_list
}

func (obj *VirtualMachineInterface) GetVirtualMachineInterfaceHostRoutes() RouteTableType {
        return obj.virtual_machine_interface_host_routes
}

func (obj *VirtualMachineInterface) SetVirtualMachineInterfaceHostRoutes(value *RouteTableType) {
        obj.virtual_machine_interface_host_routes = *value
        obj.modified |= virtual_machine_interface_virtual_machine_interface_host_routes
}

func (obj *VirtualMachineInterface) GetVirtualMachineInterfaceAllowedAddressPairs() AllowedAddressPairs {
        return obj.virtual_machine_interface_allowed_address_pairs
}

func (obj *VirtualMachineInterface) SetVirtualMachineInterfaceAllowedAddressPairs(value *AllowedAddressPairs) {
        obj.virtual_machine_interface_allowed_address_pairs = *value
        obj.modified |= virtual_machine_interface_virtual_machine_interface_allowed_address_pairs
}

func (obj *VirtualMachineInterface) GetVrfAssignTable() VrfAssignTableType {
        return obj.vrf_assign_table
}

func (obj *VirtualMachineInterface) SetVrfAssignTable(value *VrfAssignTableType) {
        obj.vrf_assign_table = *value
        obj.modified |= virtual_machine_interface_vrf_assign_table
}

func (obj *VirtualMachineInterface) GetVirtualMachineInterfaceDeviceOwner() string {
        return obj.virtual_machine_interface_device_owner
}

func (obj *VirtualMachineInterface) SetVirtualMachineInterfaceDeviceOwner(value string) {
        obj.virtual_machine_interface_device_owner = value
        obj.modified |= virtual_machine_interface_virtual_machine_interface_device_owner
}

func (obj *VirtualMachineInterface) GetVirtualMachineInterfaceDisablePolicy() bool {
        return obj.virtual_machine_interface_disable_policy
}

func (obj *VirtualMachineInterface) SetVirtualMachineInterfaceDisablePolicy(value bool) {
        obj.virtual_machine_interface_disable_policy = value
        obj.modified |= virtual_machine_interface_virtual_machine_interface_disable_policy
}

func (obj *VirtualMachineInterface) GetVirtualMachineInterfaceProperties() VirtualMachineInterfacePropertiesType {
        return obj.virtual_machine_interface_properties
}

func (obj *VirtualMachineInterface) SetVirtualMachineInterfaceProperties(value *VirtualMachineInterfacePropertiesType) {
        obj.virtual_machine_interface_properties = *value
        obj.modified |= virtual_machine_interface_virtual_machine_interface_properties
}

func (obj *VirtualMachineInterface) GetVirtualMachineInterfaceBindings() KeyValuePairs {
        return obj.virtual_machine_interface_bindings
}

func (obj *VirtualMachineInterface) SetVirtualMachineInterfaceBindings(value *KeyValuePairs) {
        obj.virtual_machine_interface_bindings = *value
        obj.modified |= virtual_machine_interface_virtual_machine_interface_bindings
}

func (obj *VirtualMachineInterface) GetVirtualMachineInterfaceFatFlowProtocols() FatFlowProtocols {
        return obj.virtual_machine_interface_fat_flow_protocols
}

func (obj *VirtualMachineInterface) SetVirtualMachineInterfaceFatFlowProtocols(value *FatFlowProtocols) {
        obj.virtual_machine_interface_fat_flow_protocols = *value
        obj.modified |= virtual_machine_interface_virtual_machine_interface_fat_flow_protocols
}

func (obj *VirtualMachineInterface) GetIdPerms() IdPermsType {
        return obj.id_perms
}

func (obj *VirtualMachineInterface) SetIdPerms(value *IdPermsType) {
        obj.id_perms = *value
        obj.modified |= virtual_machine_interface_id_perms
}

func (obj *VirtualMachineInterface) GetPerms2() PermType2 {
        return obj.perms2
}

func (obj *VirtualMachineInterface) SetPerms2(value *PermType2) {
        obj.perms2 = *value
        obj.modified |= virtual_machine_interface_perms2
}

func (obj *VirtualMachineInterface) GetDisplayName() string {
        return obj.display_name
}

func (obj *VirtualMachineInterface) SetDisplayName(value string) {
        obj.display_name = value
        obj.modified |= virtual_machine_interface_display_name
}

func (obj *VirtualMachineInterface) readQosConfigRefs() error {
        if !obj.IsTransient() &&
                (obj.valid & virtual_machine_interface_qos_config_refs == 0) {
                err := obj.GetField(obj, "qos_config_refs")
                if err != nil {
                        return err
                }
        }
        return nil
}

func (obj *VirtualMachineInterface) GetQosConfigRefs() (
        contrail.ReferenceList, error) {
        err := obj.readQosConfigRefs()
        if err != nil {
                return nil, err
        }
        return obj.qos_config_refs, nil
}

func (obj *VirtualMachineInterface) AddQosConfig(
        rhs *QosConfig) error {
        err := obj.readQosConfigRefs()
        if err != nil {
                return err
        }

        if obj.modified & virtual_machine_interface_qos_config_refs == 0 {
                obj.storeReferenceBase("qos-config", obj.qos_config_refs)
        }

        ref := contrail.Reference {
                rhs.GetFQName(), rhs.GetUuid(), rhs.GetHref(), nil}
        obj.qos_config_refs = append(obj.qos_config_refs, ref)
        obj.modified |= virtual_machine_interface_qos_config_refs
        return nil
}

func (obj *VirtualMachineInterface) DeleteQosConfig(uuid string) error {
        err := obj.readQosConfigRefs()
        if err != nil {
                return err
        }

        if obj.modified & virtual_machine_interface_qos_config_refs == 0 {
                obj.storeReferenceBase("qos-config", obj.qos_config_refs)
        }

        for i, ref := range obj.qos_config_refs {
                if ref.Uuid == uuid {
                        obj.qos_config_refs = append(
                                obj.qos_config_refs[:i],
                                obj.qos_config_refs[i+1:]...)
                        break
                }
        }
        obj.modified |= virtual_machine_interface_qos_config_refs
        return nil
}

func (obj *VirtualMachineInterface) ClearQosConfig() {
        if (obj.valid & virtual_machine_interface_qos_config_refs != 0) &&
           (obj.modified & virtual_machine_interface_qos_config_refs == 0) {
                obj.storeReferenceBase("qos-config", obj.qos_config_refs)
        }
        obj.qos_config_refs = make([]contrail.Reference, 0)
        obj.valid |= virtual_machine_interface_qos_config_refs
        obj.modified |= virtual_machine_interface_qos_config_refs
}

func (obj *VirtualMachineInterface) SetQosConfigList(
        refList []contrail.ReferencePair) {
        obj.ClearQosConfig()
        obj.qos_config_refs = make([]contrail.Reference, len(refList))
        for i, pair := range refList {
                obj.qos_config_refs[i] = contrail.Reference {
                        pair.Object.GetFQName(),
                        pair.Object.GetUuid(),
                        pair.Object.GetHref(),
                        pair.Attribute,
                }
        }
}


func (obj *VirtualMachineInterface) readSecurityGroupRefs() error {
        if !obj.IsTransient() &&
                (obj.valid & virtual_machine_interface_security_group_refs == 0) {
                err := obj.GetField(obj, "security_group_refs")
                if err != nil {
                        return err
                }
        }
        return nil
}

func (obj *VirtualMachineInterface) GetSecurityGroupRefs() (
        contrail.ReferenceList, error) {
        err := obj.readSecurityGroupRefs()
        if err != nil {
                return nil, err
        }
        return obj.security_group_refs, nil
}

func (obj *VirtualMachineInterface) AddSecurityGroup(
        rhs *SecurityGroup) error {
        err := obj.readSecurityGroupRefs()
        if err != nil {
                return err
        }

        if obj.modified & virtual_machine_interface_security_group_refs == 0 {
                obj.storeReferenceBase("security-group", obj.security_group_refs)
        }

        ref := contrail.Reference {
                rhs.GetFQName(), rhs.GetUuid(), rhs.GetHref(), nil}
        obj.security_group_refs = append(obj.security_group_refs, ref)
        obj.modified |= virtual_machine_interface_security_group_refs
        return nil
}

func (obj *VirtualMachineInterface) DeleteSecurityGroup(uuid string) error {
        err := obj.readSecurityGroupRefs()
        if err != nil {
                return err
        }

        if obj.modified & virtual_machine_interface_security_group_refs == 0 {
                obj.storeReferenceBase("security-group", obj.security_group_refs)
        }

        for i, ref := range obj.security_group_refs {
                if ref.Uuid == uuid {
                        obj.security_group_refs = append(
                                obj.security_group_refs[:i],
                                obj.security_group_refs[i+1:]...)
                        break
                }
        }
        obj.modified |= virtual_machine_interface_security_group_refs
        return nil
}

func (obj *VirtualMachineInterface) ClearSecurityGroup() {
        if (obj.valid & virtual_machine_interface_security_group_refs != 0) &&
           (obj.modified & virtual_machine_interface_security_group_refs == 0) {
                obj.storeReferenceBase("security-group", obj.security_group_refs)
        }
        obj.security_group_refs = make([]contrail.Reference, 0)
        obj.valid |= virtual_machine_interface_security_group_refs
        obj.modified |= virtual_machine_interface_security_group_refs
}

func (obj *VirtualMachineInterface) SetSecurityGroupList(
        refList []contrail.ReferencePair) {
        obj.ClearSecurityGroup()
        obj.security_group_refs = make([]contrail.Reference, len(refList))
        for i, pair := range refList {
                obj.security_group_refs[i] = contrail.Reference {
                        pair.Object.GetFQName(),
                        pair.Object.GetUuid(),
                        pair.Object.GetHref(),
                        pair.Attribute,
                }
        }
}


func (obj *VirtualMachineInterface) readVirtualMachineInterfaceRefs() error {
        if !obj.IsTransient() &&
                (obj.valid & virtual_machine_interface_virtual_machine_interface_refs == 0) {
                err := obj.GetField(obj, "virtual_machine_interface_refs")
                if err != nil {
                        return err
                }
        }
        return nil
}

func (obj *VirtualMachineInterface) GetVirtualMachineInterfaceRefs() (
        contrail.ReferenceList, error) {
        err := obj.readVirtualMachineInterfaceRefs()
        if err != nil {
                return nil, err
        }
        return obj.virtual_machine_interface_refs, nil
}

func (obj *VirtualMachineInterface) AddVirtualMachineInterface(
        rhs *VirtualMachineInterface) error {
        err := obj.readVirtualMachineInterfaceRefs()
        if err != nil {
                return err
        }

        if obj.modified & virtual_machine_interface_virtual_machine_interface_refs == 0 {
                obj.storeReferenceBase("virtual-machine-interface", obj.virtual_machine_interface_refs)
        }

        ref := contrail.Reference {
                rhs.GetFQName(), rhs.GetUuid(), rhs.GetHref(), nil}
        obj.virtual_machine_interface_refs = append(obj.virtual_machine_interface_refs, ref)
        obj.modified |= virtual_machine_interface_virtual_machine_interface_refs
        return nil
}

func (obj *VirtualMachineInterface) DeleteVirtualMachineInterface(uuid string) error {
        err := obj.readVirtualMachineInterfaceRefs()
        if err != nil {
                return err
        }

        if obj.modified & virtual_machine_interface_virtual_machine_interface_refs == 0 {
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
        obj.modified |= virtual_machine_interface_virtual_machine_interface_refs
        return nil
}

func (obj *VirtualMachineInterface) ClearVirtualMachineInterface() {
        if (obj.valid & virtual_machine_interface_virtual_machine_interface_refs != 0) &&
           (obj.modified & virtual_machine_interface_virtual_machine_interface_refs == 0) {
                obj.storeReferenceBase("virtual-machine-interface", obj.virtual_machine_interface_refs)
        }
        obj.virtual_machine_interface_refs = make([]contrail.Reference, 0)
        obj.valid |= virtual_machine_interface_virtual_machine_interface_refs
        obj.modified |= virtual_machine_interface_virtual_machine_interface_refs
}

func (obj *VirtualMachineInterface) SetVirtualMachineInterfaceList(
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


func (obj *VirtualMachineInterface) readVirtualMachineRefs() error {
        if !obj.IsTransient() &&
                (obj.valid & virtual_machine_interface_virtual_machine_refs == 0) {
                err := obj.GetField(obj, "virtual_machine_refs")
                if err != nil {
                        return err
                }
        }
        return nil
}

func (obj *VirtualMachineInterface) GetVirtualMachineRefs() (
        contrail.ReferenceList, error) {
        err := obj.readVirtualMachineRefs()
        if err != nil {
                return nil, err
        }
        return obj.virtual_machine_refs, nil
}

func (obj *VirtualMachineInterface) AddVirtualMachine(
        rhs *VirtualMachine) error {
        err := obj.readVirtualMachineRefs()
        if err != nil {
                return err
        }

        if obj.modified & virtual_machine_interface_virtual_machine_refs == 0 {
                obj.storeReferenceBase("virtual-machine", obj.virtual_machine_refs)
        }

        ref := contrail.Reference {
                rhs.GetFQName(), rhs.GetUuid(), rhs.GetHref(), nil}
        obj.virtual_machine_refs = append(obj.virtual_machine_refs, ref)
        obj.modified |= virtual_machine_interface_virtual_machine_refs
        return nil
}

func (obj *VirtualMachineInterface) DeleteVirtualMachine(uuid string) error {
        err := obj.readVirtualMachineRefs()
        if err != nil {
                return err
        }

        if obj.modified & virtual_machine_interface_virtual_machine_refs == 0 {
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
        obj.modified |= virtual_machine_interface_virtual_machine_refs
        return nil
}

func (obj *VirtualMachineInterface) ClearVirtualMachine() {
        if (obj.valid & virtual_machine_interface_virtual_machine_refs != 0) &&
           (obj.modified & virtual_machine_interface_virtual_machine_refs == 0) {
                obj.storeReferenceBase("virtual-machine", obj.virtual_machine_refs)
        }
        obj.virtual_machine_refs = make([]contrail.Reference, 0)
        obj.valid |= virtual_machine_interface_virtual_machine_refs
        obj.modified |= virtual_machine_interface_virtual_machine_refs
}

func (obj *VirtualMachineInterface) SetVirtualMachineList(
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


func (obj *VirtualMachineInterface) readVirtualNetworkRefs() error {
        if !obj.IsTransient() &&
                (obj.valid & virtual_machine_interface_virtual_network_refs == 0) {
                err := obj.GetField(obj, "virtual_network_refs")
                if err != nil {
                        return err
                }
        }
        return nil
}

func (obj *VirtualMachineInterface) GetVirtualNetworkRefs() (
        contrail.ReferenceList, error) {
        err := obj.readVirtualNetworkRefs()
        if err != nil {
                return nil, err
        }
        return obj.virtual_network_refs, nil
}

func (obj *VirtualMachineInterface) AddVirtualNetwork(
        rhs *VirtualNetwork) error {
        err := obj.readVirtualNetworkRefs()
        if err != nil {
                return err
        }

        if obj.modified & virtual_machine_interface_virtual_network_refs == 0 {
                obj.storeReferenceBase("virtual-network", obj.virtual_network_refs)
        }

        ref := contrail.Reference {
                rhs.GetFQName(), rhs.GetUuid(), rhs.GetHref(), nil}
        obj.virtual_network_refs = append(obj.virtual_network_refs, ref)
        obj.modified |= virtual_machine_interface_virtual_network_refs
        return nil
}

func (obj *VirtualMachineInterface) DeleteVirtualNetwork(uuid string) error {
        err := obj.readVirtualNetworkRefs()
        if err != nil {
                return err
        }

        if obj.modified & virtual_machine_interface_virtual_network_refs == 0 {
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
        obj.modified |= virtual_machine_interface_virtual_network_refs
        return nil
}

func (obj *VirtualMachineInterface) ClearVirtualNetwork() {
        if (obj.valid & virtual_machine_interface_virtual_network_refs != 0) &&
           (obj.modified & virtual_machine_interface_virtual_network_refs == 0) {
                obj.storeReferenceBase("virtual-network", obj.virtual_network_refs)
        }
        obj.virtual_network_refs = make([]contrail.Reference, 0)
        obj.valid |= virtual_machine_interface_virtual_network_refs
        obj.modified |= virtual_machine_interface_virtual_network_refs
}

func (obj *VirtualMachineInterface) SetVirtualNetworkList(
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


func (obj *VirtualMachineInterface) readRoutingInstanceRefs() error {
        if !obj.IsTransient() &&
                (obj.valid & virtual_machine_interface_routing_instance_refs == 0) {
                err := obj.GetField(obj, "routing_instance_refs")
                if err != nil {
                        return err
                }
        }
        return nil
}

func (obj *VirtualMachineInterface) GetRoutingInstanceRefs() (
        contrail.ReferenceList, error) {
        err := obj.readRoutingInstanceRefs()
        if err != nil {
                return nil, err
        }
        return obj.routing_instance_refs, nil
}

func (obj *VirtualMachineInterface) AddRoutingInstance(
        rhs *RoutingInstance, data PolicyBasedForwardingRuleType) error {
        err := obj.readRoutingInstanceRefs()
        if err != nil {
                return err
        }

        if obj.modified & virtual_machine_interface_routing_instance_refs == 0 {
                obj.storeReferenceBase("routing-instance", obj.routing_instance_refs)
        }

        ref := contrail.Reference {
                rhs.GetFQName(), rhs.GetUuid(), rhs.GetHref(), data}
        obj.routing_instance_refs = append(obj.routing_instance_refs, ref)
        obj.modified |= virtual_machine_interface_routing_instance_refs
        return nil
}

func (obj *VirtualMachineInterface) DeleteRoutingInstance(uuid string) error {
        err := obj.readRoutingInstanceRefs()
        if err != nil {
                return err
        }

        if obj.modified & virtual_machine_interface_routing_instance_refs == 0 {
                obj.storeReferenceBase("routing-instance", obj.routing_instance_refs)
        }

        for i, ref := range obj.routing_instance_refs {
                if ref.Uuid == uuid {
                        obj.routing_instance_refs = append(
                                obj.routing_instance_refs[:i],
                                obj.routing_instance_refs[i+1:]...)
                        break
                }
        }
        obj.modified |= virtual_machine_interface_routing_instance_refs
        return nil
}

func (obj *VirtualMachineInterface) ClearRoutingInstance() {
        if (obj.valid & virtual_machine_interface_routing_instance_refs != 0) &&
           (obj.modified & virtual_machine_interface_routing_instance_refs == 0) {
                obj.storeReferenceBase("routing-instance", obj.routing_instance_refs)
        }
        obj.routing_instance_refs = make([]contrail.Reference, 0)
        obj.valid |= virtual_machine_interface_routing_instance_refs
        obj.modified |= virtual_machine_interface_routing_instance_refs
}

func (obj *VirtualMachineInterface) SetRoutingInstanceList(
        refList []contrail.ReferencePair) {
        obj.ClearRoutingInstance()
        obj.routing_instance_refs = make([]contrail.Reference, len(refList))
        for i, pair := range refList {
                obj.routing_instance_refs[i] = contrail.Reference {
                        pair.Object.GetFQName(),
                        pair.Object.GetUuid(),
                        pair.Object.GetHref(),
                        pair.Attribute,
                }
        }
}


func (obj *VirtualMachineInterface) readPortTupleRefs() error {
        if !obj.IsTransient() &&
                (obj.valid & virtual_machine_interface_port_tuple_refs == 0) {
                err := obj.GetField(obj, "port_tuple_refs")
                if err != nil {
                        return err
                }
        }
        return nil
}

func (obj *VirtualMachineInterface) GetPortTupleRefs() (
        contrail.ReferenceList, error) {
        err := obj.readPortTupleRefs()
        if err != nil {
                return nil, err
        }
        return obj.port_tuple_refs, nil
}

func (obj *VirtualMachineInterface) AddPortTuple(
        rhs *PortTuple) error {
        err := obj.readPortTupleRefs()
        if err != nil {
                return err
        }

        if obj.modified & virtual_machine_interface_port_tuple_refs == 0 {
                obj.storeReferenceBase("port-tuple", obj.port_tuple_refs)
        }

        ref := contrail.Reference {
                rhs.GetFQName(), rhs.GetUuid(), rhs.GetHref(), nil}
        obj.port_tuple_refs = append(obj.port_tuple_refs, ref)
        obj.modified |= virtual_machine_interface_port_tuple_refs
        return nil
}

func (obj *VirtualMachineInterface) DeletePortTuple(uuid string) error {
        err := obj.readPortTupleRefs()
        if err != nil {
                return err
        }

        if obj.modified & virtual_machine_interface_port_tuple_refs == 0 {
                obj.storeReferenceBase("port-tuple", obj.port_tuple_refs)
        }

        for i, ref := range obj.port_tuple_refs {
                if ref.Uuid == uuid {
                        obj.port_tuple_refs = append(
                                obj.port_tuple_refs[:i],
                                obj.port_tuple_refs[i+1:]...)
                        break
                }
        }
        obj.modified |= virtual_machine_interface_port_tuple_refs
        return nil
}

func (obj *VirtualMachineInterface) ClearPortTuple() {
        if (obj.valid & virtual_machine_interface_port_tuple_refs != 0) &&
           (obj.modified & virtual_machine_interface_port_tuple_refs == 0) {
                obj.storeReferenceBase("port-tuple", obj.port_tuple_refs)
        }
        obj.port_tuple_refs = make([]contrail.Reference, 0)
        obj.valid |= virtual_machine_interface_port_tuple_refs
        obj.modified |= virtual_machine_interface_port_tuple_refs
}

func (obj *VirtualMachineInterface) SetPortTupleList(
        refList []contrail.ReferencePair) {
        obj.ClearPortTuple()
        obj.port_tuple_refs = make([]contrail.Reference, len(refList))
        for i, pair := range refList {
                obj.port_tuple_refs[i] = contrail.Reference {
                        pair.Object.GetFQName(),
                        pair.Object.GetUuid(),
                        pair.Object.GetHref(),
                        pair.Attribute,
                }
        }
}


func (obj *VirtualMachineInterface) readServiceHealthCheckRefs() error {
        if !obj.IsTransient() &&
                (obj.valid & virtual_machine_interface_service_health_check_refs == 0) {
                err := obj.GetField(obj, "service_health_check_refs")
                if err != nil {
                        return err
                }
        }
        return nil
}

func (obj *VirtualMachineInterface) GetServiceHealthCheckRefs() (
        contrail.ReferenceList, error) {
        err := obj.readServiceHealthCheckRefs()
        if err != nil {
                return nil, err
        }
        return obj.service_health_check_refs, nil
}

func (obj *VirtualMachineInterface) AddServiceHealthCheck(
        rhs *ServiceHealthCheck) error {
        err := obj.readServiceHealthCheckRefs()
        if err != nil {
                return err
        }

        if obj.modified & virtual_machine_interface_service_health_check_refs == 0 {
                obj.storeReferenceBase("service-health-check", obj.service_health_check_refs)
        }

        ref := contrail.Reference {
                rhs.GetFQName(), rhs.GetUuid(), rhs.GetHref(), nil}
        obj.service_health_check_refs = append(obj.service_health_check_refs, ref)
        obj.modified |= virtual_machine_interface_service_health_check_refs
        return nil
}

func (obj *VirtualMachineInterface) DeleteServiceHealthCheck(uuid string) error {
        err := obj.readServiceHealthCheckRefs()
        if err != nil {
                return err
        }

        if obj.modified & virtual_machine_interface_service_health_check_refs == 0 {
                obj.storeReferenceBase("service-health-check", obj.service_health_check_refs)
        }

        for i, ref := range obj.service_health_check_refs {
                if ref.Uuid == uuid {
                        obj.service_health_check_refs = append(
                                obj.service_health_check_refs[:i],
                                obj.service_health_check_refs[i+1:]...)
                        break
                }
        }
        obj.modified |= virtual_machine_interface_service_health_check_refs
        return nil
}

func (obj *VirtualMachineInterface) ClearServiceHealthCheck() {
        if (obj.valid & virtual_machine_interface_service_health_check_refs != 0) &&
           (obj.modified & virtual_machine_interface_service_health_check_refs == 0) {
                obj.storeReferenceBase("service-health-check", obj.service_health_check_refs)
        }
        obj.service_health_check_refs = make([]contrail.Reference, 0)
        obj.valid |= virtual_machine_interface_service_health_check_refs
        obj.modified |= virtual_machine_interface_service_health_check_refs
}

func (obj *VirtualMachineInterface) SetServiceHealthCheckList(
        refList []contrail.ReferencePair) {
        obj.ClearServiceHealthCheck()
        obj.service_health_check_refs = make([]contrail.Reference, len(refList))
        for i, pair := range refList {
                obj.service_health_check_refs[i] = contrail.Reference {
                        pair.Object.GetFQName(),
                        pair.Object.GetUuid(),
                        pair.Object.GetHref(),
                        pair.Attribute,
                }
        }
}


func (obj *VirtualMachineInterface) readInterfaceRouteTableRefs() error {
        if !obj.IsTransient() &&
                (obj.valid & virtual_machine_interface_interface_route_table_refs == 0) {
                err := obj.GetField(obj, "interface_route_table_refs")
                if err != nil {
                        return err
                }
        }
        return nil
}

func (obj *VirtualMachineInterface) GetInterfaceRouteTableRefs() (
        contrail.ReferenceList, error) {
        err := obj.readInterfaceRouteTableRefs()
        if err != nil {
                return nil, err
        }
        return obj.interface_route_table_refs, nil
}

func (obj *VirtualMachineInterface) AddInterfaceRouteTable(
        rhs *InterfaceRouteTable) error {
        err := obj.readInterfaceRouteTableRefs()
        if err != nil {
                return err
        }

        if obj.modified & virtual_machine_interface_interface_route_table_refs == 0 {
                obj.storeReferenceBase("interface-route-table", obj.interface_route_table_refs)
        }

        ref := contrail.Reference {
                rhs.GetFQName(), rhs.GetUuid(), rhs.GetHref(), nil}
        obj.interface_route_table_refs = append(obj.interface_route_table_refs, ref)
        obj.modified |= virtual_machine_interface_interface_route_table_refs
        return nil
}

func (obj *VirtualMachineInterface) DeleteInterfaceRouteTable(uuid string) error {
        err := obj.readInterfaceRouteTableRefs()
        if err != nil {
                return err
        }

        if obj.modified & virtual_machine_interface_interface_route_table_refs == 0 {
                obj.storeReferenceBase("interface-route-table", obj.interface_route_table_refs)
        }

        for i, ref := range obj.interface_route_table_refs {
                if ref.Uuid == uuid {
                        obj.interface_route_table_refs = append(
                                obj.interface_route_table_refs[:i],
                                obj.interface_route_table_refs[i+1:]...)
                        break
                }
        }
        obj.modified |= virtual_machine_interface_interface_route_table_refs
        return nil
}

func (obj *VirtualMachineInterface) ClearInterfaceRouteTable() {
        if (obj.valid & virtual_machine_interface_interface_route_table_refs != 0) &&
           (obj.modified & virtual_machine_interface_interface_route_table_refs == 0) {
                obj.storeReferenceBase("interface-route-table", obj.interface_route_table_refs)
        }
        obj.interface_route_table_refs = make([]contrail.Reference, 0)
        obj.valid |= virtual_machine_interface_interface_route_table_refs
        obj.modified |= virtual_machine_interface_interface_route_table_refs
}

func (obj *VirtualMachineInterface) SetInterfaceRouteTableList(
        refList []contrail.ReferencePair) {
        obj.ClearInterfaceRouteTable()
        obj.interface_route_table_refs = make([]contrail.Reference, len(refList))
        for i, pair := range refList {
                obj.interface_route_table_refs[i] = contrail.Reference {
                        pair.Object.GetFQName(),
                        pair.Object.GetUuid(),
                        pair.Object.GetHref(),
                        pair.Attribute,
                }
        }
}


func (obj *VirtualMachineInterface) readPhysicalInterfaceRefs() error {
        if !obj.IsTransient() &&
                (obj.valid & virtual_machine_interface_physical_interface_refs == 0) {
                err := obj.GetField(obj, "physical_interface_refs")
                if err != nil {
                        return err
                }
        }
        return nil
}

func (obj *VirtualMachineInterface) GetPhysicalInterfaceRefs() (
        contrail.ReferenceList, error) {
        err := obj.readPhysicalInterfaceRefs()
        if err != nil {
                return nil, err
        }
        return obj.physical_interface_refs, nil
}

func (obj *VirtualMachineInterface) AddPhysicalInterface(
        rhs *PhysicalInterface) error {
        err := obj.readPhysicalInterfaceRefs()
        if err != nil {
                return err
        }

        if obj.modified & virtual_machine_interface_physical_interface_refs == 0 {
                obj.storeReferenceBase("physical-interface", obj.physical_interface_refs)
        }

        ref := contrail.Reference {
                rhs.GetFQName(), rhs.GetUuid(), rhs.GetHref(), nil}
        obj.physical_interface_refs = append(obj.physical_interface_refs, ref)
        obj.modified |= virtual_machine_interface_physical_interface_refs
        return nil
}

func (obj *VirtualMachineInterface) DeletePhysicalInterface(uuid string) error {
        err := obj.readPhysicalInterfaceRefs()
        if err != nil {
                return err
        }

        if obj.modified & virtual_machine_interface_physical_interface_refs == 0 {
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
        obj.modified |= virtual_machine_interface_physical_interface_refs
        return nil
}

func (obj *VirtualMachineInterface) ClearPhysicalInterface() {
        if (obj.valid & virtual_machine_interface_physical_interface_refs != 0) &&
           (obj.modified & virtual_machine_interface_physical_interface_refs == 0) {
                obj.storeReferenceBase("physical-interface", obj.physical_interface_refs)
        }
        obj.physical_interface_refs = make([]contrail.Reference, 0)
        obj.valid |= virtual_machine_interface_physical_interface_refs
        obj.modified |= virtual_machine_interface_physical_interface_refs
}

func (obj *VirtualMachineInterface) SetPhysicalInterfaceList(
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


func (obj *VirtualMachineInterface) readVirtualMachineInterfaceBackRefs() error {
        if !obj.IsTransient() &&
                (obj.valid & virtual_machine_interface_virtual_machine_interface_back_refs == 0) {
                err := obj.GetField(obj, "virtual_machine_interface_back_refs")
                if err != nil {
                        return err
                }
        }
        return nil
}

func (obj *VirtualMachineInterface) GetVirtualMachineInterfaceBackRefs() (
        contrail.ReferenceList, error) {
        err := obj.readVirtualMachineInterfaceBackRefs()
        if err != nil {
                return nil, err
        }
        return obj.virtual_machine_interface_back_refs, nil
}

func (obj *VirtualMachineInterface) readInstanceIpBackRefs() error {
        if !obj.IsTransient() &&
                (obj.valid & virtual_machine_interface_instance_ip_back_refs == 0) {
                err := obj.GetField(obj, "instance_ip_back_refs")
                if err != nil {
                        return err
                }
        }
        return nil
}

func (obj *VirtualMachineInterface) GetInstanceIpBackRefs() (
        contrail.ReferenceList, error) {
        err := obj.readInstanceIpBackRefs()
        if err != nil {
                return nil, err
        }
        return obj.instance_ip_back_refs, nil
}

func (obj *VirtualMachineInterface) readSubnetBackRefs() error {
        if !obj.IsTransient() &&
                (obj.valid & virtual_machine_interface_subnet_back_refs == 0) {
                err := obj.GetField(obj, "subnet_back_refs")
                if err != nil {
                        return err
                }
        }
        return nil
}

func (obj *VirtualMachineInterface) GetSubnetBackRefs() (
        contrail.ReferenceList, error) {
        err := obj.readSubnetBackRefs()
        if err != nil {
                return nil, err
        }
        return obj.subnet_back_refs, nil
}

func (obj *VirtualMachineInterface) readFloatingIpBackRefs() error {
        if !obj.IsTransient() &&
                (obj.valid & virtual_machine_interface_floating_ip_back_refs == 0) {
                err := obj.GetField(obj, "floating_ip_back_refs")
                if err != nil {
                        return err
                }
        }
        return nil
}

func (obj *VirtualMachineInterface) GetFloatingIpBackRefs() (
        contrail.ReferenceList, error) {
        err := obj.readFloatingIpBackRefs()
        if err != nil {
                return nil, err
        }
        return obj.floating_ip_back_refs, nil
}

func (obj *VirtualMachineInterface) readAliasIpBackRefs() error {
        if !obj.IsTransient() &&
                (obj.valid & virtual_machine_interface_alias_ip_back_refs == 0) {
                err := obj.GetField(obj, "alias_ip_back_refs")
                if err != nil {
                        return err
                }
        }
        return nil
}

func (obj *VirtualMachineInterface) GetAliasIpBackRefs() (
        contrail.ReferenceList, error) {
        err := obj.readAliasIpBackRefs()
        if err != nil {
                return nil, err
        }
        return obj.alias_ip_back_refs, nil
}

func (obj *VirtualMachineInterface) readLogicalInterfaceBackRefs() error {
        if !obj.IsTransient() &&
                (obj.valid & virtual_machine_interface_logical_interface_back_refs == 0) {
                err := obj.GetField(obj, "logical_interface_back_refs")
                if err != nil {
                        return err
                }
        }
        return nil
}

func (obj *VirtualMachineInterface) GetLogicalInterfaceBackRefs() (
        contrail.ReferenceList, error) {
        err := obj.readLogicalInterfaceBackRefs()
        if err != nil {
                return nil, err
        }
        return obj.logical_interface_back_refs, nil
}

func (obj *VirtualMachineInterface) readBgpAsAServiceBackRefs() error {
        if !obj.IsTransient() &&
                (obj.valid & virtual_machine_interface_bgp_as_a_service_back_refs == 0) {
                err := obj.GetField(obj, "bgp_as_a_service_back_refs")
                if err != nil {
                        return err
                }
        }
        return nil
}

func (obj *VirtualMachineInterface) GetBgpAsAServiceBackRefs() (
        contrail.ReferenceList, error) {
        err := obj.readBgpAsAServiceBackRefs()
        if err != nil {
                return nil, err
        }
        return obj.bgp_as_a_service_back_refs, nil
}

func (obj *VirtualMachineInterface) readCustomerAttachmentBackRefs() error {
        if !obj.IsTransient() &&
                (obj.valid & virtual_machine_interface_customer_attachment_back_refs == 0) {
                err := obj.GetField(obj, "customer_attachment_back_refs")
                if err != nil {
                        return err
                }
        }
        return nil
}

func (obj *VirtualMachineInterface) GetCustomerAttachmentBackRefs() (
        contrail.ReferenceList, error) {
        err := obj.readCustomerAttachmentBackRefs()
        if err != nil {
                return nil, err
        }
        return obj.customer_attachment_back_refs, nil
}

func (obj *VirtualMachineInterface) readLogicalRouterBackRefs() error {
        if !obj.IsTransient() &&
                (obj.valid & virtual_machine_interface_logical_router_back_refs == 0) {
                err := obj.GetField(obj, "logical_router_back_refs")
                if err != nil {
                        return err
                }
        }
        return nil
}

func (obj *VirtualMachineInterface) GetLogicalRouterBackRefs() (
        contrail.ReferenceList, error) {
        err := obj.readLogicalRouterBackRefs()
        if err != nil {
                return nil, err
        }
        return obj.logical_router_back_refs, nil
}

func (obj *VirtualMachineInterface) readLoadbalancerPoolBackRefs() error {
        if !obj.IsTransient() &&
                (obj.valid & virtual_machine_interface_loadbalancer_pool_back_refs == 0) {
                err := obj.GetField(obj, "loadbalancer_pool_back_refs")
                if err != nil {
                        return err
                }
        }
        return nil
}

func (obj *VirtualMachineInterface) GetLoadbalancerPoolBackRefs() (
        contrail.ReferenceList, error) {
        err := obj.readLoadbalancerPoolBackRefs()
        if err != nil {
                return nil, err
        }
        return obj.loadbalancer_pool_back_refs, nil
}

func (obj *VirtualMachineInterface) readVirtualIpBackRefs() error {
        if !obj.IsTransient() &&
                (obj.valid & virtual_machine_interface_virtual_ip_back_refs == 0) {
                err := obj.GetField(obj, "virtual_ip_back_refs")
                if err != nil {
                        return err
                }
        }
        return nil
}

func (obj *VirtualMachineInterface) GetVirtualIpBackRefs() (
        contrail.ReferenceList, error) {
        err := obj.readVirtualIpBackRefs()
        if err != nil {
                return nil, err
        }
        return obj.virtual_ip_back_refs, nil
}

func (obj *VirtualMachineInterface) readLoadbalancerBackRefs() error {
        if !obj.IsTransient() &&
                (obj.valid & virtual_machine_interface_loadbalancer_back_refs == 0) {
                err := obj.GetField(obj, "loadbalancer_back_refs")
                if err != nil {
                        return err
                }
        }
        return nil
}

func (obj *VirtualMachineInterface) GetLoadbalancerBackRefs() (
        contrail.ReferenceList, error) {
        err := obj.readLoadbalancerBackRefs()
        if err != nil {
                return nil, err
        }
        return obj.loadbalancer_back_refs, nil
}

func (obj *VirtualMachineInterface) MarshalJSON() ([]byte, error) {
        msg := map[string]*json.RawMessage {
        }
        err := obj.MarshalCommon(msg)
        if err != nil {
                return nil, err
        }

        if obj.modified & virtual_machine_interface_ecmp_hashing_include_fields != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.ecmp_hashing_include_fields)
                if err != nil {
                        return nil, err
                }
                msg["ecmp_hashing_include_fields"] = &value
        }

        if obj.modified & virtual_machine_interface_virtual_machine_interface_mac_addresses != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.virtual_machine_interface_mac_addresses)
                if err != nil {
                        return nil, err
                }
                msg["virtual_machine_interface_mac_addresses"] = &value
        }

        if obj.modified & virtual_machine_interface_virtual_machine_interface_dhcp_option_list != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.virtual_machine_interface_dhcp_option_list)
                if err != nil {
                        return nil, err
                }
                msg["virtual_machine_interface_dhcp_option_list"] = &value
        }

        if obj.modified & virtual_machine_interface_virtual_machine_interface_host_routes != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.virtual_machine_interface_host_routes)
                if err != nil {
                        return nil, err
                }
                msg["virtual_machine_interface_host_routes"] = &value
        }

        if obj.modified & virtual_machine_interface_virtual_machine_interface_allowed_address_pairs != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.virtual_machine_interface_allowed_address_pairs)
                if err != nil {
                        return nil, err
                }
                msg["virtual_machine_interface_allowed_address_pairs"] = &value
        }

        if obj.modified & virtual_machine_interface_vrf_assign_table != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.vrf_assign_table)
                if err != nil {
                        return nil, err
                }
                msg["vrf_assign_table"] = &value
        }

        if obj.modified & virtual_machine_interface_virtual_machine_interface_device_owner != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.virtual_machine_interface_device_owner)
                if err != nil {
                        return nil, err
                }
                msg["virtual_machine_interface_device_owner"] = &value
        }

        if obj.modified & virtual_machine_interface_virtual_machine_interface_disable_policy != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.virtual_machine_interface_disable_policy)
                if err != nil {
                        return nil, err
                }
                msg["virtual_machine_interface_disable_policy"] = &value
        }

        if obj.modified & virtual_machine_interface_virtual_machine_interface_properties != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.virtual_machine_interface_properties)
                if err != nil {
                        return nil, err
                }
                msg["virtual_machine_interface_properties"] = &value
        }

        if obj.modified & virtual_machine_interface_virtual_machine_interface_bindings != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.virtual_machine_interface_bindings)
                if err != nil {
                        return nil, err
                }
                msg["virtual_machine_interface_bindings"] = &value
        }

        if obj.modified & virtual_machine_interface_virtual_machine_interface_fat_flow_protocols != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.virtual_machine_interface_fat_flow_protocols)
                if err != nil {
                        return nil, err
                }
                msg["virtual_machine_interface_fat_flow_protocols"] = &value
        }

        if obj.modified & virtual_machine_interface_id_perms != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.id_perms)
                if err != nil {
                        return nil, err
                }
                msg["id_perms"] = &value
        }

        if obj.modified & virtual_machine_interface_perms2 != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.perms2)
                if err != nil {
                        return nil, err
                }
                msg["perms2"] = &value
        }

        if obj.modified & virtual_machine_interface_display_name != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.display_name)
                if err != nil {
                        return nil, err
                }
                msg["display_name"] = &value
        }

        if len(obj.qos_config_refs) > 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.qos_config_refs)
                if err != nil {
                        return nil, err
                }
                msg["qos_config_refs"] = &value
        }

        if len(obj.security_group_refs) > 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.security_group_refs)
                if err != nil {
                        return nil, err
                }
                msg["security_group_refs"] = &value
        }

        if len(obj.virtual_machine_interface_refs) > 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.virtual_machine_interface_refs)
                if err != nil {
                        return nil, err
                }
                msg["virtual_machine_interface_refs"] = &value
        }

        if len(obj.virtual_machine_refs) > 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.virtual_machine_refs)
                if err != nil {
                        return nil, err
                }
                msg["virtual_machine_refs"] = &value
        }

        if len(obj.virtual_network_refs) > 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.virtual_network_refs)
                if err != nil {
                        return nil, err
                }
                msg["virtual_network_refs"] = &value
        }

        if len(obj.routing_instance_refs) > 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.routing_instance_refs)
                if err != nil {
                        return nil, err
                }
                msg["routing_instance_refs"] = &value
        }

        if len(obj.port_tuple_refs) > 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.port_tuple_refs)
                if err != nil {
                        return nil, err
                }
                msg["port_tuple_refs"] = &value
        }

        if len(obj.service_health_check_refs) > 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.service_health_check_refs)
                if err != nil {
                        return nil, err
                }
                msg["service_health_check_refs"] = &value
        }

        if len(obj.interface_route_table_refs) > 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.interface_route_table_refs)
                if err != nil {
                        return nil, err
                }
                msg["interface_route_table_refs"] = &value
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

func (obj *VirtualMachineInterface) UnmarshalJSON(body []byte) error {
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
                case "ecmp_hashing_include_fields":
                        err = json.Unmarshal(value, &obj.ecmp_hashing_include_fields)
                        if err == nil {
                                obj.valid |= virtual_machine_interface_ecmp_hashing_include_fields
                        }
                        break
                case "virtual_machine_interface_mac_addresses":
                        err = json.Unmarshal(value, &obj.virtual_machine_interface_mac_addresses)
                        if err == nil {
                                obj.valid |= virtual_machine_interface_virtual_machine_interface_mac_addresses
                        }
                        break
                case "virtual_machine_interface_dhcp_option_list":
                        err = json.Unmarshal(value, &obj.virtual_machine_interface_dhcp_option_list)
                        if err == nil {
                                obj.valid |= virtual_machine_interface_virtual_machine_interface_dhcp_option_list
                        }
                        break
                case "virtual_machine_interface_host_routes":
                        err = json.Unmarshal(value, &obj.virtual_machine_interface_host_routes)
                        if err == nil {
                                obj.valid |= virtual_machine_interface_virtual_machine_interface_host_routes
                        }
                        break
                case "virtual_machine_interface_allowed_address_pairs":
                        err = json.Unmarshal(value, &obj.virtual_machine_interface_allowed_address_pairs)
                        if err == nil {
                                obj.valid |= virtual_machine_interface_virtual_machine_interface_allowed_address_pairs
                        }
                        break
                case "vrf_assign_table":
                        err = json.Unmarshal(value, &obj.vrf_assign_table)
                        if err == nil {
                                obj.valid |= virtual_machine_interface_vrf_assign_table
                        }
                        break
                case "virtual_machine_interface_device_owner":
                        err = json.Unmarshal(value, &obj.virtual_machine_interface_device_owner)
                        if err == nil {
                                obj.valid |= virtual_machine_interface_virtual_machine_interface_device_owner
                        }
                        break
                case "virtual_machine_interface_disable_policy":
                        err = json.Unmarshal(value, &obj.virtual_machine_interface_disable_policy)
                        if err == nil {
                                obj.valid |= virtual_machine_interface_virtual_machine_interface_disable_policy
                        }
                        break
                case "virtual_machine_interface_properties":
                        err = json.Unmarshal(value, &obj.virtual_machine_interface_properties)
                        if err == nil {
                                obj.valid |= virtual_machine_interface_virtual_machine_interface_properties
                        }
                        break
                case "virtual_machine_interface_bindings":
                        err = json.Unmarshal(value, &obj.virtual_machine_interface_bindings)
                        if err == nil {
                                obj.valid |= virtual_machine_interface_virtual_machine_interface_bindings
                        }
                        break
                case "virtual_machine_interface_fat_flow_protocols":
                        err = json.Unmarshal(value, &obj.virtual_machine_interface_fat_flow_protocols)
                        if err == nil {
                                obj.valid |= virtual_machine_interface_virtual_machine_interface_fat_flow_protocols
                        }
                        break
                case "id_perms":
                        err = json.Unmarshal(value, &obj.id_perms)
                        if err == nil {
                                obj.valid |= virtual_machine_interface_id_perms
                        }
                        break
                case "perms2":
                        err = json.Unmarshal(value, &obj.perms2)
                        if err == nil {
                                obj.valid |= virtual_machine_interface_perms2
                        }
                        break
                case "display_name":
                        err = json.Unmarshal(value, &obj.display_name)
                        if err == nil {
                                obj.valid |= virtual_machine_interface_display_name
                        }
                        break
                case "qos_config_refs":
                        err = json.Unmarshal(value, &obj.qos_config_refs)
                        if err == nil {
                                obj.valid |= virtual_machine_interface_qos_config_refs
                        }
                        break
                case "security_group_refs":
                        err = json.Unmarshal(value, &obj.security_group_refs)
                        if err == nil {
                                obj.valid |= virtual_machine_interface_security_group_refs
                        }
                        break
                case "virtual_machine_interface_refs":
                        err = json.Unmarshal(value, &obj.virtual_machine_interface_refs)
                        if err == nil {
                                obj.valid |= virtual_machine_interface_virtual_machine_interface_refs
                        }
                        break
                case "virtual_machine_refs":
                        err = json.Unmarshal(value, &obj.virtual_machine_refs)
                        if err == nil {
                                obj.valid |= virtual_machine_interface_virtual_machine_refs
                        }
                        break
                case "virtual_network_refs":
                        err = json.Unmarshal(value, &obj.virtual_network_refs)
                        if err == nil {
                                obj.valid |= virtual_machine_interface_virtual_network_refs
                        }
                        break
                case "port_tuple_refs":
                        err = json.Unmarshal(value, &obj.port_tuple_refs)
                        if err == nil {
                                obj.valid |= virtual_machine_interface_port_tuple_refs
                        }
                        break
                case "service_health_check_refs":
                        err = json.Unmarshal(value, &obj.service_health_check_refs)
                        if err == nil {
                                obj.valid |= virtual_machine_interface_service_health_check_refs
                        }
                        break
                case "interface_route_table_refs":
                        err = json.Unmarshal(value, &obj.interface_route_table_refs)
                        if err == nil {
                                obj.valid |= virtual_machine_interface_interface_route_table_refs
                        }
                        break
                case "physical_interface_refs":
                        err = json.Unmarshal(value, &obj.physical_interface_refs)
                        if err == nil {
                                obj.valid |= virtual_machine_interface_physical_interface_refs
                        }
                        break
                case "virtual_machine_interface_back_refs":
                        err = json.Unmarshal(value, &obj.virtual_machine_interface_back_refs)
                        if err == nil {
                                obj.valid |= virtual_machine_interface_virtual_machine_interface_back_refs
                        }
                        break
                case "instance_ip_back_refs":
                        err = json.Unmarshal(value, &obj.instance_ip_back_refs)
                        if err == nil {
                                obj.valid |= virtual_machine_interface_instance_ip_back_refs
                        }
                        break
                case "subnet_back_refs":
                        err = json.Unmarshal(value, &obj.subnet_back_refs)
                        if err == nil {
                                obj.valid |= virtual_machine_interface_subnet_back_refs
                        }
                        break
                case "floating_ip_back_refs":
                        err = json.Unmarshal(value, &obj.floating_ip_back_refs)
                        if err == nil {
                                obj.valid |= virtual_machine_interface_floating_ip_back_refs
                        }
                        break
                case "alias_ip_back_refs":
                        err = json.Unmarshal(value, &obj.alias_ip_back_refs)
                        if err == nil {
                                obj.valid |= virtual_machine_interface_alias_ip_back_refs
                        }
                        break
                case "logical_interface_back_refs":
                        err = json.Unmarshal(value, &obj.logical_interface_back_refs)
                        if err == nil {
                                obj.valid |= virtual_machine_interface_logical_interface_back_refs
                        }
                        break
                case "bgp_as_a_service_back_refs":
                        err = json.Unmarshal(value, &obj.bgp_as_a_service_back_refs)
                        if err == nil {
                                obj.valid |= virtual_machine_interface_bgp_as_a_service_back_refs
                        }
                        break
                case "customer_attachment_back_refs":
                        err = json.Unmarshal(value, &obj.customer_attachment_back_refs)
                        if err == nil {
                                obj.valid |= virtual_machine_interface_customer_attachment_back_refs
                        }
                        break
                case "logical_router_back_refs":
                        err = json.Unmarshal(value, &obj.logical_router_back_refs)
                        if err == nil {
                                obj.valid |= virtual_machine_interface_logical_router_back_refs
                        }
                        break
                case "loadbalancer_pool_back_refs":
                        err = json.Unmarshal(value, &obj.loadbalancer_pool_back_refs)
                        if err == nil {
                                obj.valid |= virtual_machine_interface_loadbalancer_pool_back_refs
                        }
                        break
                case "virtual_ip_back_refs":
                        err = json.Unmarshal(value, &obj.virtual_ip_back_refs)
                        if err == nil {
                                obj.valid |= virtual_machine_interface_virtual_ip_back_refs
                        }
                        break
                case "loadbalancer_back_refs":
                        err = json.Unmarshal(value, &obj.loadbalancer_back_refs)
                        if err == nil {
                                obj.valid |= virtual_machine_interface_loadbalancer_back_refs
                        }
                        break
                case "routing_instance_refs": {
                        type ReferenceElement struct {
                                To []string
                                Uuid string
                                Href string
                                Attr PolicyBasedForwardingRuleType
                        }
                        var array []ReferenceElement
                        err = json.Unmarshal(value, &array)
                        if err != nil {
                            break
                        }
                        obj.valid |= virtual_machine_interface_routing_instance_refs
                        obj.routing_instance_refs = make(contrail.ReferenceList, 0)
                        for _, element := range array {
                                ref := contrail.Reference {
                                        element.To,
                                        element.Uuid,
                                        element.Href,
                                        element.Attr,
                                }
                                obj.routing_instance_refs = append(obj.routing_instance_refs, ref)
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

func (obj *VirtualMachineInterface) UpdateObject() ([]byte, error) {
        msg := map[string]*json.RawMessage {
        }
        err := obj.MarshalId(msg)
        if err != nil {
                return nil, err
        }

        if obj.modified & virtual_machine_interface_ecmp_hashing_include_fields != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.ecmp_hashing_include_fields)
                if err != nil {
                        return nil, err
                }
                msg["ecmp_hashing_include_fields"] = &value
        }

        if obj.modified & virtual_machine_interface_virtual_machine_interface_mac_addresses != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.virtual_machine_interface_mac_addresses)
                if err != nil {
                        return nil, err
                }
                msg["virtual_machine_interface_mac_addresses"] = &value
        }

        if obj.modified & virtual_machine_interface_virtual_machine_interface_dhcp_option_list != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.virtual_machine_interface_dhcp_option_list)
                if err != nil {
                        return nil, err
                }
                msg["virtual_machine_interface_dhcp_option_list"] = &value
        }

        if obj.modified & virtual_machine_interface_virtual_machine_interface_host_routes != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.virtual_machine_interface_host_routes)
                if err != nil {
                        return nil, err
                }
                msg["virtual_machine_interface_host_routes"] = &value
        }

        if obj.modified & virtual_machine_interface_virtual_machine_interface_allowed_address_pairs != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.virtual_machine_interface_allowed_address_pairs)
                if err != nil {
                        return nil, err
                }
                msg["virtual_machine_interface_allowed_address_pairs"] = &value
        }

        if obj.modified & virtual_machine_interface_vrf_assign_table != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.vrf_assign_table)
                if err != nil {
                        return nil, err
                }
                msg["vrf_assign_table"] = &value
        }

        if obj.modified & virtual_machine_interface_virtual_machine_interface_device_owner != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.virtual_machine_interface_device_owner)
                if err != nil {
                        return nil, err
                }
                msg["virtual_machine_interface_device_owner"] = &value
        }

        if obj.modified & virtual_machine_interface_virtual_machine_interface_disable_policy != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.virtual_machine_interface_disable_policy)
                if err != nil {
                        return nil, err
                }
                msg["virtual_machine_interface_disable_policy"] = &value
        }

        if obj.modified & virtual_machine_interface_virtual_machine_interface_properties != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.virtual_machine_interface_properties)
                if err != nil {
                        return nil, err
                }
                msg["virtual_machine_interface_properties"] = &value
        }

        if obj.modified & virtual_machine_interface_virtual_machine_interface_bindings != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.virtual_machine_interface_bindings)
                if err != nil {
                        return nil, err
                }
                msg["virtual_machine_interface_bindings"] = &value
        }

        if obj.modified & virtual_machine_interface_virtual_machine_interface_fat_flow_protocols != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.virtual_machine_interface_fat_flow_protocols)
                if err != nil {
                        return nil, err
                }
                msg["virtual_machine_interface_fat_flow_protocols"] = &value
        }

        if obj.modified & virtual_machine_interface_id_perms != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.id_perms)
                if err != nil {
                        return nil, err
                }
                msg["id_perms"] = &value
        }

        if obj.modified & virtual_machine_interface_perms2 != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.perms2)
                if err != nil {
                        return nil, err
                }
                msg["perms2"] = &value
        }

        if obj.modified & virtual_machine_interface_display_name != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.display_name)
                if err != nil {
                        return nil, err
                }
                msg["display_name"] = &value
        }

        if obj.modified & virtual_machine_interface_qos_config_refs != 0 {
                if len(obj.qos_config_refs) == 0 {
                        var value json.RawMessage
                        value, err := json.Marshal(
                                          make([]contrail.Reference, 0))
                        if err != nil {
                                return nil, err
                        }
                        msg["qos_config_refs"] = &value
                } else if !obj.hasReferenceBase("qos-config") {
                        var value json.RawMessage
                        value, err := json.Marshal(&obj.qos_config_refs)
                        if err != nil {
                                return nil, err
                        }
                        msg["qos_config_refs"] = &value
                }
        }


        if obj.modified & virtual_machine_interface_security_group_refs != 0 {
                if len(obj.security_group_refs) == 0 {
                        var value json.RawMessage
                        value, err := json.Marshal(
                                          make([]contrail.Reference, 0))
                        if err != nil {
                                return nil, err
                        }
                        msg["security_group_refs"] = &value
                } else if !obj.hasReferenceBase("security-group") {
                        var value json.RawMessage
                        value, err := json.Marshal(&obj.security_group_refs)
                        if err != nil {
                                return nil, err
                        }
                        msg["security_group_refs"] = &value
                }
        }


        if obj.modified & virtual_machine_interface_virtual_machine_interface_refs != 0 {
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


        if obj.modified & virtual_machine_interface_virtual_machine_refs != 0 {
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


        if obj.modified & virtual_machine_interface_virtual_network_refs != 0 {
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


        if obj.modified & virtual_machine_interface_routing_instance_refs != 0 {
                if len(obj.routing_instance_refs) == 0 {
                        var value json.RawMessage
                        value, err := json.Marshal(
                                          make([]contrail.Reference, 0))
                        if err != nil {
                                return nil, err
                        }
                        msg["routing_instance_refs"] = &value
                } else if !obj.hasReferenceBase("routing-instance") {
                        var value json.RawMessage
                        value, err := json.Marshal(&obj.routing_instance_refs)
                        if err != nil {
                                return nil, err
                        }
                        msg["routing_instance_refs"] = &value
                }
        }


        if obj.modified & virtual_machine_interface_port_tuple_refs != 0 {
                if len(obj.port_tuple_refs) == 0 {
                        var value json.RawMessage
                        value, err := json.Marshal(
                                          make([]contrail.Reference, 0))
                        if err != nil {
                                return nil, err
                        }
                        msg["port_tuple_refs"] = &value
                } else if !obj.hasReferenceBase("port-tuple") {
                        var value json.RawMessage
                        value, err := json.Marshal(&obj.port_tuple_refs)
                        if err != nil {
                                return nil, err
                        }
                        msg["port_tuple_refs"] = &value
                }
        }


        if obj.modified & virtual_machine_interface_service_health_check_refs != 0 {
                if len(obj.service_health_check_refs) == 0 {
                        var value json.RawMessage
                        value, err := json.Marshal(
                                          make([]contrail.Reference, 0))
                        if err != nil {
                                return nil, err
                        }
                        msg["service_health_check_refs"] = &value
                } else if !obj.hasReferenceBase("service-health-check") {
                        var value json.RawMessage
                        value, err := json.Marshal(&obj.service_health_check_refs)
                        if err != nil {
                                return nil, err
                        }
                        msg["service_health_check_refs"] = &value
                }
        }


        if obj.modified & virtual_machine_interface_interface_route_table_refs != 0 {
                if len(obj.interface_route_table_refs) == 0 {
                        var value json.RawMessage
                        value, err := json.Marshal(
                                          make([]contrail.Reference, 0))
                        if err != nil {
                                return nil, err
                        }
                        msg["interface_route_table_refs"] = &value
                } else if !obj.hasReferenceBase("interface-route-table") {
                        var value json.RawMessage
                        value, err := json.Marshal(&obj.interface_route_table_refs)
                        if err != nil {
                                return nil, err
                        }
                        msg["interface_route_table_refs"] = &value
                }
        }


        if obj.modified & virtual_machine_interface_physical_interface_refs != 0 {
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

func (obj *VirtualMachineInterface) UpdateReferences() error {

        if (obj.modified & virtual_machine_interface_qos_config_refs != 0) &&
           len(obj.qos_config_refs) > 0 &&
           obj.hasReferenceBase("qos-config") {
                err := obj.UpdateReference(
                        obj, "qos-config",
                        obj.qos_config_refs,
                        obj.baseMap["qos-config"])
                if err != nil {
                        return err
                }
        }

        if (obj.modified & virtual_machine_interface_security_group_refs != 0) &&
           len(obj.security_group_refs) > 0 &&
           obj.hasReferenceBase("security-group") {
                err := obj.UpdateReference(
                        obj, "security-group",
                        obj.security_group_refs,
                        obj.baseMap["security-group"])
                if err != nil {
                        return err
                }
        }

        if (obj.modified & virtual_machine_interface_virtual_machine_interface_refs != 0) &&
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

        if (obj.modified & virtual_machine_interface_virtual_machine_refs != 0) &&
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

        if (obj.modified & virtual_machine_interface_virtual_network_refs != 0) &&
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

        if (obj.modified & virtual_machine_interface_routing_instance_refs != 0) &&
           len(obj.routing_instance_refs) > 0 &&
           obj.hasReferenceBase("routing-instance") {
                err := obj.UpdateReference(
                        obj, "routing-instance",
                        obj.routing_instance_refs,
                        obj.baseMap["routing-instance"])
                if err != nil {
                        return err
                }
        }

        if (obj.modified & virtual_machine_interface_port_tuple_refs != 0) &&
           len(obj.port_tuple_refs) > 0 &&
           obj.hasReferenceBase("port-tuple") {
                err := obj.UpdateReference(
                        obj, "port-tuple",
                        obj.port_tuple_refs,
                        obj.baseMap["port-tuple"])
                if err != nil {
                        return err
                }
        }

        if (obj.modified & virtual_machine_interface_service_health_check_refs != 0) &&
           len(obj.service_health_check_refs) > 0 &&
           obj.hasReferenceBase("service-health-check") {
                err := obj.UpdateReference(
                        obj, "service-health-check",
                        obj.service_health_check_refs,
                        obj.baseMap["service-health-check"])
                if err != nil {
                        return err
                }
        }

        if (obj.modified & virtual_machine_interface_interface_route_table_refs != 0) &&
           len(obj.interface_route_table_refs) > 0 &&
           obj.hasReferenceBase("interface-route-table") {
                err := obj.UpdateReference(
                        obj, "interface-route-table",
                        obj.interface_route_table_refs,
                        obj.baseMap["interface-route-table"])
                if err != nil {
                        return err
                }
        }

        if (obj.modified & virtual_machine_interface_physical_interface_refs != 0) &&
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

func VirtualMachineInterfaceByName(c contrail.ApiClient, fqn string) (*VirtualMachineInterface, error) {
    obj, err := c.FindByName("virtual-machine-interface", fqn)
    if err != nil {
        return nil, err
    }
    return obj.(*VirtualMachineInterface), nil
}

func VirtualMachineInterfaceByUuid(c contrail.ApiClient, uuid string) (*VirtualMachineInterface, error) {
    obj, err := c.FindByUuid("virtual-machine-interface", uuid)
    if err != nil {
        return nil, err
    }
    return obj.(*VirtualMachineInterface), nil
}
