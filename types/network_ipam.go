//
// Automatically generated. DO NOT EDIT.
//

package types

import (
        "encoding/json"

        "github.com/Juniper/contrail-go-api"
)

const (
	network_ipam_network_ipam_mgmt uint64 = 1 << iota
	network_ipam_id_perms
	network_ipam_perms2
	network_ipam_display_name
	network_ipam_virtual_DNS_refs
	network_ipam_virtual_network_back_refs
)

type NetworkIpam struct {
        contrail.ObjectBase
	network_ipam_mgmt IpamType
	id_perms IdPermsType
	perms2 PermType2
	display_name string
	virtual_DNS_refs contrail.ReferenceList
	virtual_network_back_refs contrail.ReferenceList
        valid uint64
        modified uint64
        baseMap map[string]contrail.ReferenceList
}

func (obj *NetworkIpam) GetType() string {
        return "network-ipam"
}

func (obj *NetworkIpam) GetDefaultParent() []string {
        name := []string{"default-domain", "default-project"}
        return name
}

func (obj *NetworkIpam) GetDefaultParentType() string {
        return "project"
}

func (obj *NetworkIpam) SetName(name string) {
        obj.VSetName(obj, name)
}

func (obj *NetworkIpam) SetParent(parent contrail.IObject) {
        obj.VSetParent(obj, parent)
}

func (obj *NetworkIpam) storeReferenceBase(
        name string, refList contrail.ReferenceList) {
        if obj.baseMap == nil {
                obj.baseMap = make(map[string]contrail.ReferenceList)
        }
        refCopy := make(contrail.ReferenceList, len(refList))
        copy(refCopy, refList)
        obj.baseMap[name] = refCopy
}

func (obj *NetworkIpam) hasReferenceBase(name string) bool {
        if obj.baseMap == nil {
                return false
        }
        _, exists := obj.baseMap[name]
        return exists
}

func (obj *NetworkIpam) UpdateDone() {
        obj.modified = 0
        obj.baseMap = nil
}


func (obj *NetworkIpam) GetNetworkIpamMgmt() IpamType {
        return obj.network_ipam_mgmt
}

func (obj *NetworkIpam) SetNetworkIpamMgmt(value *IpamType) {
        obj.network_ipam_mgmt = *value
        obj.modified |= network_ipam_network_ipam_mgmt
}

func (obj *NetworkIpam) GetIdPerms() IdPermsType {
        return obj.id_perms
}

func (obj *NetworkIpam) SetIdPerms(value *IdPermsType) {
        obj.id_perms = *value
        obj.modified |= network_ipam_id_perms
}

func (obj *NetworkIpam) GetPerms2() PermType2 {
        return obj.perms2
}

func (obj *NetworkIpam) SetPerms2(value *PermType2) {
        obj.perms2 = *value
        obj.modified |= network_ipam_perms2
}

func (obj *NetworkIpam) GetDisplayName() string {
        return obj.display_name
}

func (obj *NetworkIpam) SetDisplayName(value string) {
        obj.display_name = value
        obj.modified |= network_ipam_display_name
}

func (obj *NetworkIpam) readVirtualDnsRefs() error {
        if !obj.IsTransient() &&
                (obj.valid & network_ipam_virtual_DNS_refs == 0) {
                err := obj.GetField(obj, "virtual_DNS_refs")
                if err != nil {
                        return err
                }
        }
        return nil
}

func (obj *NetworkIpam) GetVirtualDnsRefs() (
        contrail.ReferenceList, error) {
        err := obj.readVirtualDnsRefs()
        if err != nil {
                return nil, err
        }
        return obj.virtual_DNS_refs, nil
}

func (obj *NetworkIpam) AddVirtualDns(
        rhs *VirtualDns) error {
        err := obj.readVirtualDnsRefs()
        if err != nil {
                return err
        }

        if obj.modified & network_ipam_virtual_DNS_refs == 0 {
                obj.storeReferenceBase("virtual-DNS", obj.virtual_DNS_refs)
        }

        ref := contrail.Reference {
                rhs.GetFQName(), rhs.GetUuid(), rhs.GetHref(), nil}
        obj.virtual_DNS_refs = append(obj.virtual_DNS_refs, ref)
        obj.modified |= network_ipam_virtual_DNS_refs
        return nil
}

func (obj *NetworkIpam) DeleteVirtualDns(uuid string) error {
        err := obj.readVirtualDnsRefs()
        if err != nil {
                return err
        }

        if obj.modified & network_ipam_virtual_DNS_refs == 0 {
                obj.storeReferenceBase("virtual-DNS", obj.virtual_DNS_refs)
        }

        for i, ref := range obj.virtual_DNS_refs {
                if ref.Uuid == uuid {
                        obj.virtual_DNS_refs = append(
                                obj.virtual_DNS_refs[:i],
                                obj.virtual_DNS_refs[i+1:]...)
                        break
                }
        }
        obj.modified |= network_ipam_virtual_DNS_refs
        return nil
}

func (obj *NetworkIpam) ClearVirtualDns() {
        if (obj.valid & network_ipam_virtual_DNS_refs != 0) &&
           (obj.modified & network_ipam_virtual_DNS_refs == 0) {
                obj.storeReferenceBase("virtual-DNS", obj.virtual_DNS_refs)
        }
        obj.virtual_DNS_refs = make([]contrail.Reference, 0)
        obj.valid |= network_ipam_virtual_DNS_refs
        obj.modified |= network_ipam_virtual_DNS_refs
}

func (obj *NetworkIpam) SetVirtualDnsList(
        refList []contrail.ReferencePair) {
        obj.ClearVirtualDns()
        obj.virtual_DNS_refs = make([]contrail.Reference, len(refList))
        for i, pair := range refList {
                obj.virtual_DNS_refs[i] = contrail.Reference {
                        pair.Object.GetFQName(),
                        pair.Object.GetUuid(),
                        pair.Object.GetHref(),
                        pair.Attribute,
                }
        }
}


func (obj *NetworkIpam) readVirtualNetworkBackRefs() error {
        if !obj.IsTransient() &&
                (obj.valid & network_ipam_virtual_network_back_refs == 0) {
                err := obj.GetField(obj, "virtual_network_back_refs")
                if err != nil {
                        return err
                }
        }
        return nil
}

func (obj *NetworkIpam) GetVirtualNetworkBackRefs() (
        contrail.ReferenceList, error) {
        err := obj.readVirtualNetworkBackRefs()
        if err != nil {
                return nil, err
        }
        return obj.virtual_network_back_refs, nil
}

func (obj *NetworkIpam) MarshalJSON() ([]byte, error) {
        msg := map[string]*json.RawMessage {
        }
        err := obj.MarshalCommon(msg)
        if err != nil {
                return nil, err
        }

        if obj.modified & network_ipam_network_ipam_mgmt != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.network_ipam_mgmt)
                if err != nil {
                        return nil, err
                }
                msg["network_ipam_mgmt"] = &value
        }

        if obj.modified & network_ipam_id_perms != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.id_perms)
                if err != nil {
                        return nil, err
                }
                msg["id_perms"] = &value
        }

        if obj.modified & network_ipam_perms2 != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.perms2)
                if err != nil {
                        return nil, err
                }
                msg["perms2"] = &value
        }

        if obj.modified & network_ipam_display_name != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.display_name)
                if err != nil {
                        return nil, err
                }
                msg["display_name"] = &value
        }

        if len(obj.virtual_DNS_refs) > 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.virtual_DNS_refs)
                if err != nil {
                        return nil, err
                }
                msg["virtual_DNS_refs"] = &value
        }

        return json.Marshal(msg)
}

func (obj *NetworkIpam) UnmarshalJSON(body []byte) error {
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
                case "network_ipam_mgmt":
                        err = json.Unmarshal(value, &obj.network_ipam_mgmt)
                        if err == nil {
                                obj.valid |= network_ipam_network_ipam_mgmt
                        }
                        break
                case "id_perms":
                        err = json.Unmarshal(value, &obj.id_perms)
                        if err == nil {
                                obj.valid |= network_ipam_id_perms
                        }
                        break
                case "perms2":
                        err = json.Unmarshal(value, &obj.perms2)
                        if err == nil {
                                obj.valid |= network_ipam_perms2
                        }
                        break
                case "display_name":
                        err = json.Unmarshal(value, &obj.display_name)
                        if err == nil {
                                obj.valid |= network_ipam_display_name
                        }
                        break
                case "virtual_DNS_refs":
                        err = json.Unmarshal(value, &obj.virtual_DNS_refs)
                        if err == nil {
                                obj.valid |= network_ipam_virtual_DNS_refs
                        }
                        break
                case "virtual_network_back_refs": {
                        type ReferenceElement struct {
                                To []string
                                Uuid string
                                Href string
                                Attr VnSubnetsType
                        }
                        var array []ReferenceElement
                        err = json.Unmarshal(value, &array)
                        if err != nil {
                            break
                        }
                        obj.valid |= network_ipam_virtual_network_back_refs
                        obj.virtual_network_back_refs = make(contrail.ReferenceList, 0)
                        for _, element := range array {
                                ref := contrail.Reference {
                                        element.To,
                                        element.Uuid,
                                        element.Href,
                                        element.Attr,
                                }
                                obj.virtual_network_back_refs = append(obj.virtual_network_back_refs, ref)
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

func (obj *NetworkIpam) UpdateObject() ([]byte, error) {
        msg := map[string]*json.RawMessage {
        }
        err := obj.MarshalId(msg)
        if err != nil {
                return nil, err
        }

        if obj.modified & network_ipam_network_ipam_mgmt != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.network_ipam_mgmt)
                if err != nil {
                        return nil, err
                }
                msg["network_ipam_mgmt"] = &value
        }

        if obj.modified & network_ipam_id_perms != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.id_perms)
                if err != nil {
                        return nil, err
                }
                msg["id_perms"] = &value
        }

        if obj.modified & network_ipam_perms2 != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.perms2)
                if err != nil {
                        return nil, err
                }
                msg["perms2"] = &value
        }

        if obj.modified & network_ipam_display_name != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.display_name)
                if err != nil {
                        return nil, err
                }
                msg["display_name"] = &value
        }

        if obj.modified & network_ipam_virtual_DNS_refs != 0 {
                if len(obj.virtual_DNS_refs) == 0 {
                        var value json.RawMessage
                        value, err := json.Marshal(
                                          make([]contrail.Reference, 0))
                        if err != nil {
                                return nil, err
                        }
                        msg["virtual_DNS_refs"] = &value
                } else if !obj.hasReferenceBase("virtual-DNS") {
                        var value json.RawMessage
                        value, err := json.Marshal(&obj.virtual_DNS_refs)
                        if err != nil {
                                return nil, err
                        }
                        msg["virtual_DNS_refs"] = &value
                }
        }


        return json.Marshal(msg)
}

func (obj *NetworkIpam) UpdateReferences() error {

        if (obj.modified & network_ipam_virtual_DNS_refs != 0) &&
           len(obj.virtual_DNS_refs) > 0 &&
           obj.hasReferenceBase("virtual-DNS") {
                err := obj.UpdateReference(
                        obj, "virtual-DNS",
                        obj.virtual_DNS_refs,
                        obj.baseMap["virtual-DNS"])
                if err != nil {
                        return err
                }
        }

        return nil
}

func NetworkIpamByName(c contrail.ApiClient, fqn string) (*NetworkIpam, error) {
    obj, err := c.FindByName("network-ipam", fqn)
    if err != nil {
        return nil, err
    }
    return obj.(*NetworkIpam), nil
}

func NetworkIpamByUuid(c contrail.ApiClient, uuid string) (*NetworkIpam, error) {
    obj, err := c.FindByUuid("network-ipam", uuid)
    if err != nil {
        return nil, err
    }
    return obj.(*NetworkIpam), nil
}
