//
// Automatically generated. DO NOT EDIT.
//

package types

import (
        "encoding/json"

        "github.com/michaelhenkel/contrail-go-api"
)

const (
	qos_config_qos_config_type uint64 = 1 << iota
	qos_config_dscp_entries
	qos_config_vlan_priority_entries
	qos_config_mpls_exp_entries
	qos_config_trusted
	qos_config_id_perms
	qos_config_perms2
	qos_config_display_name
	qos_config_virtual_network_back_refs
	qos_config_virtual_machine_interface_back_refs
)

type QosConfig struct {
        contrail.ObjectBase
	qos_config_type string
	dscp_entries QosIdForwardingClassPairs
	vlan_priority_entries QosIdForwardingClassPairs
	mpls_exp_entries QosIdForwardingClassPairs
	trusted bool
	id_perms IdPermsType
	perms2 PermType2
	display_name string
	virtual_network_back_refs contrail.ReferenceList
	virtual_machine_interface_back_refs contrail.ReferenceList
        valid uint64
        modified uint64
        baseMap map[string]contrail.ReferenceList
}

func (obj *QosConfig) GetType() string {
        return "qos-config"
}

func (obj *QosConfig) GetDefaultParent() []string {
        name := []string{"default-global-system-config", "default-global-qos-config"}
        return name
}

func (obj *QosConfig) GetDefaultParentType() string {
        return "global-qos-config"
}

func (obj *QosConfig) SetName(name string) {
        obj.VSetName(obj, name)
}

func (obj *QosConfig) SetParent(parent contrail.IObject) {
        obj.VSetParent(obj, parent)
}

func (obj *QosConfig) storeReferenceBase(
        name string, refList contrail.ReferenceList) {
        if obj.baseMap == nil {
                obj.baseMap = make(map[string]contrail.ReferenceList)
        }
        refCopy := make(contrail.ReferenceList, len(refList))
        copy(refCopy, refList)
        obj.baseMap[name] = refCopy
}

func (obj *QosConfig) hasReferenceBase(name string) bool {
        if obj.baseMap == nil {
                return false
        }
        _, exists := obj.baseMap[name]
        return exists
}

func (obj *QosConfig) UpdateDone() {
        obj.modified = 0
        obj.baseMap = nil
}


func (obj *QosConfig) GetQosConfigType() string {
        return obj.qos_config_type
}

func (obj *QosConfig) SetQosConfigType(value string) {
        obj.qos_config_type = value
        obj.modified |= qos_config_qos_config_type
}

func (obj *QosConfig) GetDscpEntries() QosIdForwardingClassPairs {
        return obj.dscp_entries
}

func (obj *QosConfig) SetDscpEntries(value *QosIdForwardingClassPairs) {
        obj.dscp_entries = *value
        obj.modified |= qos_config_dscp_entries
}

func (obj *QosConfig) GetVlanPriorityEntries() QosIdForwardingClassPairs {
        return obj.vlan_priority_entries
}

func (obj *QosConfig) SetVlanPriorityEntries(value *QosIdForwardingClassPairs) {
        obj.vlan_priority_entries = *value
        obj.modified |= qos_config_vlan_priority_entries
}

func (obj *QosConfig) GetMplsExpEntries() QosIdForwardingClassPairs {
        return obj.mpls_exp_entries
}

func (obj *QosConfig) SetMplsExpEntries(value *QosIdForwardingClassPairs) {
        obj.mpls_exp_entries = *value
        obj.modified |= qos_config_mpls_exp_entries
}

func (obj *QosConfig) GetTrusted() bool {
        return obj.trusted
}

func (obj *QosConfig) SetTrusted(value bool) {
        obj.trusted = value
        obj.modified |= qos_config_trusted
}

func (obj *QosConfig) GetIdPerms() IdPermsType {
        return obj.id_perms
}

func (obj *QosConfig) SetIdPerms(value *IdPermsType) {
        obj.id_perms = *value
        obj.modified |= qos_config_id_perms
}

func (obj *QosConfig) GetPerms2() PermType2 {
        return obj.perms2
}

func (obj *QosConfig) SetPerms2(value *PermType2) {
        obj.perms2 = *value
        obj.modified |= qos_config_perms2
}

func (obj *QosConfig) GetDisplayName() string {
        return obj.display_name
}

func (obj *QosConfig) SetDisplayName(value string) {
        obj.display_name = value
        obj.modified |= qos_config_display_name
}

func (obj *QosConfig) readVirtualNetworkBackRefs() error {
        if !obj.IsTransient() &&
                (obj.valid & qos_config_virtual_network_back_refs == 0) {
                err := obj.GetField(obj, "virtual_network_back_refs")
                if err != nil {
                        return err
                }
        }
        return nil
}

func (obj *QosConfig) GetVirtualNetworkBackRefs() (
        contrail.ReferenceList, error) {
        err := obj.readVirtualNetworkBackRefs()
        if err != nil {
                return nil, err
        }
        return obj.virtual_network_back_refs, nil
}

func (obj *QosConfig) readVirtualMachineInterfaceBackRefs() error {
        if !obj.IsTransient() &&
                (obj.valid & qos_config_virtual_machine_interface_back_refs == 0) {
                err := obj.GetField(obj, "virtual_machine_interface_back_refs")
                if err != nil {
                        return err
                }
        }
        return nil
}

func (obj *QosConfig) GetVirtualMachineInterfaceBackRefs() (
        contrail.ReferenceList, error) {
        err := obj.readVirtualMachineInterfaceBackRefs()
        if err != nil {
                return nil, err
        }
        return obj.virtual_machine_interface_back_refs, nil
}

func (obj *QosConfig) MarshalJSON() ([]byte, error) {
        msg := map[string]*json.RawMessage {
        }
        err := obj.MarshalCommon(msg)
        if err != nil {
                return nil, err
        }

        if obj.modified & qos_config_qos_config_type != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.qos_config_type)
                if err != nil {
                        return nil, err
                }
                msg["qos_config_type"] = &value
        }

        if obj.modified & qos_config_dscp_entries != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.dscp_entries)
                if err != nil {
                        return nil, err
                }
                msg["dscp_entries"] = &value
        }

        if obj.modified & qos_config_vlan_priority_entries != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.vlan_priority_entries)
                if err != nil {
                        return nil, err
                }
                msg["vlan_priority_entries"] = &value
        }

        if obj.modified & qos_config_mpls_exp_entries != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.mpls_exp_entries)
                if err != nil {
                        return nil, err
                }
                msg["mpls_exp_entries"] = &value
        }

        if obj.modified & qos_config_trusted != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.trusted)
                if err != nil {
                        return nil, err
                }
                msg["trusted"] = &value
        }

        if obj.modified & qos_config_id_perms != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.id_perms)
                if err != nil {
                        return nil, err
                }
                msg["id_perms"] = &value
        }

        if obj.modified & qos_config_perms2 != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.perms2)
                if err != nil {
                        return nil, err
                }
                msg["perms2"] = &value
        }

        if obj.modified & qos_config_display_name != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.display_name)
                if err != nil {
                        return nil, err
                }
                msg["display_name"] = &value
        }

        return json.Marshal(msg)
}

func (obj *QosConfig) UnmarshalJSON(body []byte) error {
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
                case "qos_config_type":
                        err = json.Unmarshal(value, &obj.qos_config_type)
                        if err == nil {
                                obj.valid |= qos_config_qos_config_type
                        }
                        break
                case "dscp_entries":
                        err = json.Unmarshal(value, &obj.dscp_entries)
                        if err == nil {
                                obj.valid |= qos_config_dscp_entries
                        }
                        break
                case "vlan_priority_entries":
                        err = json.Unmarshal(value, &obj.vlan_priority_entries)
                        if err == nil {
                                obj.valid |= qos_config_vlan_priority_entries
                        }
                        break
                case "mpls_exp_entries":
                        err = json.Unmarshal(value, &obj.mpls_exp_entries)
                        if err == nil {
                                obj.valid |= qos_config_mpls_exp_entries
                        }
                        break
                case "trusted":
                        err = json.Unmarshal(value, &obj.trusted)
                        if err == nil {
                                obj.valid |= qos_config_trusted
                        }
                        break
                case "id_perms":
                        err = json.Unmarshal(value, &obj.id_perms)
                        if err == nil {
                                obj.valid |= qos_config_id_perms
                        }
                        break
                case "perms2":
                        err = json.Unmarshal(value, &obj.perms2)
                        if err == nil {
                                obj.valid |= qos_config_perms2
                        }
                        break
                case "display_name":
                        err = json.Unmarshal(value, &obj.display_name)
                        if err == nil {
                                obj.valid |= qos_config_display_name
                        }
                        break
                case "virtual_network_back_refs":
                        err = json.Unmarshal(value, &obj.virtual_network_back_refs)
                        if err == nil {
                                obj.valid |= qos_config_virtual_network_back_refs
                        }
                        break
                case "virtual_machine_interface_back_refs":
                        err = json.Unmarshal(value, &obj.virtual_machine_interface_back_refs)
                        if err == nil {
                                obj.valid |= qos_config_virtual_machine_interface_back_refs
                        }
                        break
                }
                if err != nil {
                        return err
                }
        }
        return nil
}

func (obj *QosConfig) UpdateObject() ([]byte, error) {
        msg := map[string]*json.RawMessage {
        }
        err := obj.MarshalId(msg)
        if err != nil {
                return nil, err
        }

        if obj.modified & qos_config_qos_config_type != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.qos_config_type)
                if err != nil {
                        return nil, err
                }
                msg["qos_config_type"] = &value
        }

        if obj.modified & qos_config_dscp_entries != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.dscp_entries)
                if err != nil {
                        return nil, err
                }
                msg["dscp_entries"] = &value
        }

        if obj.modified & qos_config_vlan_priority_entries != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.vlan_priority_entries)
                if err != nil {
                        return nil, err
                }
                msg["vlan_priority_entries"] = &value
        }

        if obj.modified & qos_config_mpls_exp_entries != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.mpls_exp_entries)
                if err != nil {
                        return nil, err
                }
                msg["mpls_exp_entries"] = &value
        }

        if obj.modified & qos_config_trusted != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.trusted)
                if err != nil {
                        return nil, err
                }
                msg["trusted"] = &value
        }

        if obj.modified & qos_config_id_perms != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.id_perms)
                if err != nil {
                        return nil, err
                }
                msg["id_perms"] = &value
        }

        if obj.modified & qos_config_perms2 != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.perms2)
                if err != nil {
                        return nil, err
                }
                msg["perms2"] = &value
        }

        if obj.modified & qos_config_display_name != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.display_name)
                if err != nil {
                        return nil, err
                }
                msg["display_name"] = &value
        }

        return json.Marshal(msg)
}

func (obj *QosConfig) UpdateReferences() error {

        return nil
}

func QosConfigByName(c contrail.ApiClient, fqn string) (*QosConfig, error) {
    obj, err := c.FindByName("qos-config", fqn)
    if err != nil {
        return nil, err
    }
    return obj.(*QosConfig), nil
}

func QosConfigByUuid(c contrail.ApiClient, uuid string) (*QosConfig, error) {
    obj, err := c.FindByUuid("qos-config", uuid)
    if err != nil {
        return nil, err
    }
    return obj.(*QosConfig), nil
}
