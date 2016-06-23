//
// Automatically generated. DO NOT EDIT.
//

package types

import (
        "encoding/json"

        "github.com/michaelhenkel/contrail-go-api"
)

const (
	forwarding_class_forwarding_class_id uint64 = 1 << iota
	forwarding_class_forwarding_class_dscp
	forwarding_class_forwarding_class_vlan_priority
	forwarding_class_forwarding_class_mpls_exp
	forwarding_class_id_perms
	forwarding_class_perms2
	forwarding_class_display_name
	forwarding_class_qos_queue_refs
)

type ForwardingClass struct {
        contrail.ObjectBase
	forwarding_class_id int
	forwarding_class_dscp int
	forwarding_class_vlan_priority int
	forwarding_class_mpls_exp int
	id_perms IdPermsType
	perms2 PermType2
	display_name string
	qos_queue_refs contrail.ReferenceList
        valid uint64
        modified uint64
        baseMap map[string]contrail.ReferenceList
}

func (obj *ForwardingClass) GetType() string {
        return "forwarding-class"
}

func (obj *ForwardingClass) GetDefaultParent() []string {
        name := []string{"default-global-system-config", "default-global-qos-config"}
        return name
}

func (obj *ForwardingClass) GetDefaultParentType() string {
        return "global-qos-config"
}

func (obj *ForwardingClass) SetName(name string) {
        obj.VSetName(obj, name)
}

func (obj *ForwardingClass) SetParent(parent contrail.IObject) {
        obj.VSetParent(obj, parent)
}

func (obj *ForwardingClass) storeReferenceBase(
        name string, refList contrail.ReferenceList) {
        if obj.baseMap == nil {
                obj.baseMap = make(map[string]contrail.ReferenceList)
        }
        refCopy := make(contrail.ReferenceList, len(refList))
        copy(refCopy, refList)
        obj.baseMap[name] = refCopy
}

func (obj *ForwardingClass) hasReferenceBase(name string) bool {
        if obj.baseMap == nil {
                return false
        }
        _, exists := obj.baseMap[name]
        return exists
}

func (obj *ForwardingClass) UpdateDone() {
        obj.modified = 0
        obj.baseMap = nil
}


func (obj *ForwardingClass) GetForwardingClassId() int {
        return obj.forwarding_class_id
}

func (obj *ForwardingClass) SetForwardingClassId(value int) {
        obj.forwarding_class_id = value
        obj.modified |= forwarding_class_forwarding_class_id
}

func (obj *ForwardingClass) GetForwardingClassDscp() int {
        return obj.forwarding_class_dscp
}

func (obj *ForwardingClass) SetForwardingClassDscp(value int) {
        obj.forwarding_class_dscp = value
        obj.modified |= forwarding_class_forwarding_class_dscp
}

func (obj *ForwardingClass) GetForwardingClassVlanPriority() int {
        return obj.forwarding_class_vlan_priority
}

func (obj *ForwardingClass) SetForwardingClassVlanPriority(value int) {
        obj.forwarding_class_vlan_priority = value
        obj.modified |= forwarding_class_forwarding_class_vlan_priority
}

func (obj *ForwardingClass) GetForwardingClassMplsExp() int {
        return obj.forwarding_class_mpls_exp
}

func (obj *ForwardingClass) SetForwardingClassMplsExp(value int) {
        obj.forwarding_class_mpls_exp = value
        obj.modified |= forwarding_class_forwarding_class_mpls_exp
}

func (obj *ForwardingClass) GetIdPerms() IdPermsType {
        return obj.id_perms
}

func (obj *ForwardingClass) SetIdPerms(value *IdPermsType) {
        obj.id_perms = *value
        obj.modified |= forwarding_class_id_perms
}

func (obj *ForwardingClass) GetPerms2() PermType2 {
        return obj.perms2
}

func (obj *ForwardingClass) SetPerms2(value *PermType2) {
        obj.perms2 = *value
        obj.modified |= forwarding_class_perms2
}

func (obj *ForwardingClass) GetDisplayName() string {
        return obj.display_name
}

func (obj *ForwardingClass) SetDisplayName(value string) {
        obj.display_name = value
        obj.modified |= forwarding_class_display_name
}

func (obj *ForwardingClass) readQosQueueRefs() error {
        if !obj.IsTransient() &&
                (obj.valid & forwarding_class_qos_queue_refs == 0) {
                err := obj.GetField(obj, "qos_queue_refs")
                if err != nil {
                        return err
                }
        }
        return nil
}

func (obj *ForwardingClass) GetQosQueueRefs() (
        contrail.ReferenceList, error) {
        err := obj.readQosQueueRefs()
        if err != nil {
                return nil, err
        }
        return obj.qos_queue_refs, nil
}

func (obj *ForwardingClass) AddQosQueue(
        rhs *QosQueue) error {
        err := obj.readQosQueueRefs()
        if err != nil {
                return err
        }

        if obj.modified & forwarding_class_qos_queue_refs == 0 {
                obj.storeReferenceBase("qos-queue", obj.qos_queue_refs)
        }

        ref := contrail.Reference {
                rhs.GetFQName(), rhs.GetUuid(), rhs.GetHref(), nil}
        obj.qos_queue_refs = append(obj.qos_queue_refs, ref)
        obj.modified |= forwarding_class_qos_queue_refs
        return nil
}

func (obj *ForwardingClass) DeleteQosQueue(uuid string) error {
        err := obj.readQosQueueRefs()
        if err != nil {
                return err
        }

        if obj.modified & forwarding_class_qos_queue_refs == 0 {
                obj.storeReferenceBase("qos-queue", obj.qos_queue_refs)
        }

        for i, ref := range obj.qos_queue_refs {
                if ref.Uuid == uuid {
                        obj.qos_queue_refs = append(
                                obj.qos_queue_refs[:i],
                                obj.qos_queue_refs[i+1:]...)
                        break
                }
        }
        obj.modified |= forwarding_class_qos_queue_refs
        return nil
}

func (obj *ForwardingClass) ClearQosQueue() {
        if (obj.valid & forwarding_class_qos_queue_refs != 0) &&
           (obj.modified & forwarding_class_qos_queue_refs == 0) {
                obj.storeReferenceBase("qos-queue", obj.qos_queue_refs)
        }
        obj.qos_queue_refs = make([]contrail.Reference, 0)
        obj.valid |= forwarding_class_qos_queue_refs
        obj.modified |= forwarding_class_qos_queue_refs
}

func (obj *ForwardingClass) SetQosQueueList(
        refList []contrail.ReferencePair) {
        obj.ClearQosQueue()
        obj.qos_queue_refs = make([]contrail.Reference, len(refList))
        for i, pair := range refList {
                obj.qos_queue_refs[i] = contrail.Reference {
                        pair.Object.GetFQName(),
                        pair.Object.GetUuid(),
                        pair.Object.GetHref(),
                        pair.Attribute,
                }
        }
}


func (obj *ForwardingClass) MarshalJSON() ([]byte, error) {
        msg := map[string]*json.RawMessage {
        }
        err := obj.MarshalCommon(msg)
        if err != nil {
                return nil, err
        }

        if obj.modified & forwarding_class_forwarding_class_id != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.forwarding_class_id)
                if err != nil {
                        return nil, err
                }
                msg["forwarding_class_id"] = &value
        }

        if obj.modified & forwarding_class_forwarding_class_dscp != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.forwarding_class_dscp)
                if err != nil {
                        return nil, err
                }
                msg["forwarding_class_dscp"] = &value
        }

        if obj.modified & forwarding_class_forwarding_class_vlan_priority != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.forwarding_class_vlan_priority)
                if err != nil {
                        return nil, err
                }
                msg["forwarding_class_vlan_priority"] = &value
        }

        if obj.modified & forwarding_class_forwarding_class_mpls_exp != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.forwarding_class_mpls_exp)
                if err != nil {
                        return nil, err
                }
                msg["forwarding_class_mpls_exp"] = &value
        }

        if obj.modified & forwarding_class_id_perms != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.id_perms)
                if err != nil {
                        return nil, err
                }
                msg["id_perms"] = &value
        }

        if obj.modified & forwarding_class_perms2 != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.perms2)
                if err != nil {
                        return nil, err
                }
                msg["perms2"] = &value
        }

        if obj.modified & forwarding_class_display_name != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.display_name)
                if err != nil {
                        return nil, err
                }
                msg["display_name"] = &value
        }

        if len(obj.qos_queue_refs) > 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.qos_queue_refs)
                if err != nil {
                        return nil, err
                }
                msg["qos_queue_refs"] = &value
        }

        return json.Marshal(msg)
}

func (obj *ForwardingClass) UnmarshalJSON(body []byte) error {
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
                case "forwarding_class_id":
                        err = json.Unmarshal(value, &obj.forwarding_class_id)
                        if err == nil {
                                obj.valid |= forwarding_class_forwarding_class_id
                        }
                        break
                case "forwarding_class_dscp":
                        err = json.Unmarshal(value, &obj.forwarding_class_dscp)
                        if err == nil {
                                obj.valid |= forwarding_class_forwarding_class_dscp
                        }
                        break
                case "forwarding_class_vlan_priority":
                        err = json.Unmarshal(value, &obj.forwarding_class_vlan_priority)
                        if err == nil {
                                obj.valid |= forwarding_class_forwarding_class_vlan_priority
                        }
                        break
                case "forwarding_class_mpls_exp":
                        err = json.Unmarshal(value, &obj.forwarding_class_mpls_exp)
                        if err == nil {
                                obj.valid |= forwarding_class_forwarding_class_mpls_exp
                        }
                        break
                case "id_perms":
                        err = json.Unmarshal(value, &obj.id_perms)
                        if err == nil {
                                obj.valid |= forwarding_class_id_perms
                        }
                        break
                case "perms2":
                        err = json.Unmarshal(value, &obj.perms2)
                        if err == nil {
                                obj.valid |= forwarding_class_perms2
                        }
                        break
                case "display_name":
                        err = json.Unmarshal(value, &obj.display_name)
                        if err == nil {
                                obj.valid |= forwarding_class_display_name
                        }
                        break
                case "qos_queue_refs":
                        err = json.Unmarshal(value, &obj.qos_queue_refs)
                        if err == nil {
                                obj.valid |= forwarding_class_qos_queue_refs
                        }
                        break
                }
                if err != nil {
                        return err
                }
        }
        return nil
}

func (obj *ForwardingClass) UpdateObject() ([]byte, error) {
        msg := map[string]*json.RawMessage {
        }
        err := obj.MarshalId(msg)
        if err != nil {
                return nil, err
        }

        if obj.modified & forwarding_class_forwarding_class_id != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.forwarding_class_id)
                if err != nil {
                        return nil, err
                }
                msg["forwarding_class_id"] = &value
        }

        if obj.modified & forwarding_class_forwarding_class_dscp != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.forwarding_class_dscp)
                if err != nil {
                        return nil, err
                }
                msg["forwarding_class_dscp"] = &value
        }

        if obj.modified & forwarding_class_forwarding_class_vlan_priority != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.forwarding_class_vlan_priority)
                if err != nil {
                        return nil, err
                }
                msg["forwarding_class_vlan_priority"] = &value
        }

        if obj.modified & forwarding_class_forwarding_class_mpls_exp != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.forwarding_class_mpls_exp)
                if err != nil {
                        return nil, err
                }
                msg["forwarding_class_mpls_exp"] = &value
        }

        if obj.modified & forwarding_class_id_perms != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.id_perms)
                if err != nil {
                        return nil, err
                }
                msg["id_perms"] = &value
        }

        if obj.modified & forwarding_class_perms2 != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.perms2)
                if err != nil {
                        return nil, err
                }
                msg["perms2"] = &value
        }

        if obj.modified & forwarding_class_display_name != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.display_name)
                if err != nil {
                        return nil, err
                }
                msg["display_name"] = &value
        }

        if obj.modified & forwarding_class_qos_queue_refs != 0 {
                if len(obj.qos_queue_refs) == 0 {
                        var value json.RawMessage
                        value, err := json.Marshal(
                                          make([]contrail.Reference, 0))
                        if err != nil {
                                return nil, err
                        }
                        msg["qos_queue_refs"] = &value
                } else if !obj.hasReferenceBase("qos-queue") {
                        var value json.RawMessage
                        value, err := json.Marshal(&obj.qos_queue_refs)
                        if err != nil {
                                return nil, err
                        }
                        msg["qos_queue_refs"] = &value
                }
        }


        return json.Marshal(msg)
}

func (obj *ForwardingClass) UpdateReferences() error {

        if (obj.modified & forwarding_class_qos_queue_refs != 0) &&
           len(obj.qos_queue_refs) > 0 &&
           obj.hasReferenceBase("qos-queue") {
                err := obj.UpdateReference(
                        obj, "qos-queue",
                        obj.qos_queue_refs,
                        obj.baseMap["qos-queue"])
                if err != nil {
                        return err
                }
        }

        return nil
}

func ForwardingClassByName(c contrail.ApiClient, fqn string) (*ForwardingClass, error) {
    obj, err := c.FindByName("forwarding-class", fqn)
    if err != nil {
        return nil, err
    }
    return obj.(*ForwardingClass), nil
}

func ForwardingClassByUuid(c contrail.ApiClient, uuid string) (*ForwardingClass, error) {
    obj, err := c.FindByUuid("forwarding-class", uuid)
    if err != nil {
        return nil, err
    }
    return obj.(*ForwardingClass), nil
}
