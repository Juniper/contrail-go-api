//
// Automatically generated. DO NOT EDIT.
//

package types

import (
        "encoding/json"

        "github.com/Juniper/contrail-go-api"
)

const (
	loadbalancer_listener_loadbalancer_listener_properties uint64 = 1 << iota
	loadbalancer_listener_id_perms
	loadbalancer_listener_perms2
	loadbalancer_listener_display_name
	loadbalancer_listener_loadbalancer_refs
	loadbalancer_listener_loadbalancer_pool_back_refs
)

type LoadbalancerListener struct {
        contrail.ObjectBase
	loadbalancer_listener_properties LoadbalancerListenerType
	id_perms IdPermsType
	perms2 PermType2
	display_name string
	loadbalancer_refs contrail.ReferenceList
	loadbalancer_pool_back_refs contrail.ReferenceList
        valid uint64
        modified uint64
        baseMap map[string]contrail.ReferenceList
}

func (obj *LoadbalancerListener) GetType() string {
        return "loadbalancer-listener"
}

func (obj *LoadbalancerListener) GetDefaultParent() []string {
        name := []string{"default-domain", "default-project"}
        return name
}

func (obj *LoadbalancerListener) GetDefaultParentType() string {
        return "project"
}

func (obj *LoadbalancerListener) SetName(name string) {
        obj.VSetName(obj, name)
}

func (obj *LoadbalancerListener) SetParent(parent contrail.IObject) {
        obj.VSetParent(obj, parent)
}

func (obj *LoadbalancerListener) storeReferenceBase(
        name string, refList contrail.ReferenceList) {
        if obj.baseMap == nil {
                obj.baseMap = make(map[string]contrail.ReferenceList)
        }
        refCopy := make(contrail.ReferenceList, len(refList))
        copy(refCopy, refList)
        obj.baseMap[name] = refCopy
}

func (obj *LoadbalancerListener) hasReferenceBase(name string) bool {
        if obj.baseMap == nil {
                return false
        }
        _, exists := obj.baseMap[name]
        return exists
}

func (obj *LoadbalancerListener) UpdateDone() {
        obj.modified = 0
        obj.baseMap = nil
}


func (obj *LoadbalancerListener) GetLoadbalancerListenerProperties() LoadbalancerListenerType {
        return obj.loadbalancer_listener_properties
}

func (obj *LoadbalancerListener) SetLoadbalancerListenerProperties(value *LoadbalancerListenerType) {
        obj.loadbalancer_listener_properties = *value
        obj.modified |= loadbalancer_listener_loadbalancer_listener_properties
}

func (obj *LoadbalancerListener) GetIdPerms() IdPermsType {
        return obj.id_perms
}

func (obj *LoadbalancerListener) SetIdPerms(value *IdPermsType) {
        obj.id_perms = *value
        obj.modified |= loadbalancer_listener_id_perms
}

func (obj *LoadbalancerListener) GetPerms2() PermType2 {
        return obj.perms2
}

func (obj *LoadbalancerListener) SetPerms2(value *PermType2) {
        obj.perms2 = *value
        obj.modified |= loadbalancer_listener_perms2
}

func (obj *LoadbalancerListener) GetDisplayName() string {
        return obj.display_name
}

func (obj *LoadbalancerListener) SetDisplayName(value string) {
        obj.display_name = value
        obj.modified |= loadbalancer_listener_display_name
}

func (obj *LoadbalancerListener) readLoadbalancerRefs() error {
        if !obj.IsTransient() &&
                (obj.valid & loadbalancer_listener_loadbalancer_refs == 0) {
                err := obj.GetField(obj, "loadbalancer_refs")
                if err != nil {
                        return err
                }
        }
        return nil
}

func (obj *LoadbalancerListener) GetLoadbalancerRefs() (
        contrail.ReferenceList, error) {
        err := obj.readLoadbalancerRefs()
        if err != nil {
                return nil, err
        }
        return obj.loadbalancer_refs, nil
}

func (obj *LoadbalancerListener) AddLoadbalancer(
        rhs *Loadbalancer) error {
        err := obj.readLoadbalancerRefs()
        if err != nil {
                return err
        }

        if obj.modified & loadbalancer_listener_loadbalancer_refs == 0 {
                obj.storeReferenceBase("loadbalancer", obj.loadbalancer_refs)
        }

        ref := contrail.Reference {
                rhs.GetFQName(), rhs.GetUuid(), rhs.GetHref(), nil}
        obj.loadbalancer_refs = append(obj.loadbalancer_refs, ref)
        obj.modified |= loadbalancer_listener_loadbalancer_refs
        return nil
}

func (obj *LoadbalancerListener) DeleteLoadbalancer(uuid string) error {
        err := obj.readLoadbalancerRefs()
        if err != nil {
                return err
        }

        if obj.modified & loadbalancer_listener_loadbalancer_refs == 0 {
                obj.storeReferenceBase("loadbalancer", obj.loadbalancer_refs)
        }

        for i, ref := range obj.loadbalancer_refs {
                if ref.Uuid == uuid {
                        obj.loadbalancer_refs = append(
                                obj.loadbalancer_refs[:i],
                                obj.loadbalancer_refs[i+1:]...)
                        break
                }
        }
        obj.modified |= loadbalancer_listener_loadbalancer_refs
        return nil
}

func (obj *LoadbalancerListener) ClearLoadbalancer() {
        if (obj.valid & loadbalancer_listener_loadbalancer_refs != 0) &&
           (obj.modified & loadbalancer_listener_loadbalancer_refs == 0) {
                obj.storeReferenceBase("loadbalancer", obj.loadbalancer_refs)
        }
        obj.loadbalancer_refs = make([]contrail.Reference, 0)
        obj.valid |= loadbalancer_listener_loadbalancer_refs
        obj.modified |= loadbalancer_listener_loadbalancer_refs
}

func (obj *LoadbalancerListener) SetLoadbalancerList(
        refList []contrail.ReferencePair) {
        obj.ClearLoadbalancer()
        obj.loadbalancer_refs = make([]contrail.Reference, len(refList))
        for i, pair := range refList {
                obj.loadbalancer_refs[i] = contrail.Reference {
                        pair.Object.GetFQName(),
                        pair.Object.GetUuid(),
                        pair.Object.GetHref(),
                        pair.Attribute,
                }
        }
}


func (obj *LoadbalancerListener) readLoadbalancerPoolBackRefs() error {
        if !obj.IsTransient() &&
                (obj.valid & loadbalancer_listener_loadbalancer_pool_back_refs == 0) {
                err := obj.GetField(obj, "loadbalancer_pool_back_refs")
                if err != nil {
                        return err
                }
        }
        return nil
}

func (obj *LoadbalancerListener) GetLoadbalancerPoolBackRefs() (
        contrail.ReferenceList, error) {
        err := obj.readLoadbalancerPoolBackRefs()
        if err != nil {
                return nil, err
        }
        return obj.loadbalancer_pool_back_refs, nil
}

func (obj *LoadbalancerListener) MarshalJSON() ([]byte, error) {
        msg := map[string]*json.RawMessage {
        }
        err := obj.MarshalCommon(msg)
        if err != nil {
                return nil, err
        }

        if obj.modified & loadbalancer_listener_loadbalancer_listener_properties != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.loadbalancer_listener_properties)
                if err != nil {
                        return nil, err
                }
                msg["loadbalancer_listener_properties"] = &value
        }

        if obj.modified & loadbalancer_listener_id_perms != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.id_perms)
                if err != nil {
                        return nil, err
                }
                msg["id_perms"] = &value
        }

        if obj.modified & loadbalancer_listener_perms2 != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.perms2)
                if err != nil {
                        return nil, err
                }
                msg["perms2"] = &value
        }

        if obj.modified & loadbalancer_listener_display_name != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.display_name)
                if err != nil {
                        return nil, err
                }
                msg["display_name"] = &value
        }

        if len(obj.loadbalancer_refs) > 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.loadbalancer_refs)
                if err != nil {
                        return nil, err
                }
                msg["loadbalancer_refs"] = &value
        }

        return json.Marshal(msg)
}

func (obj *LoadbalancerListener) UnmarshalJSON(body []byte) error {
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
                case "loadbalancer_listener_properties":
                        err = json.Unmarshal(value, &obj.loadbalancer_listener_properties)
                        if err == nil {
                                obj.valid |= loadbalancer_listener_loadbalancer_listener_properties
                        }
                        break
                case "id_perms":
                        err = json.Unmarshal(value, &obj.id_perms)
                        if err == nil {
                                obj.valid |= loadbalancer_listener_id_perms
                        }
                        break
                case "perms2":
                        err = json.Unmarshal(value, &obj.perms2)
                        if err == nil {
                                obj.valid |= loadbalancer_listener_perms2
                        }
                        break
                case "display_name":
                        err = json.Unmarshal(value, &obj.display_name)
                        if err == nil {
                                obj.valid |= loadbalancer_listener_display_name
                        }
                        break
                case "loadbalancer_refs":
                        err = json.Unmarshal(value, &obj.loadbalancer_refs)
                        if err == nil {
                                obj.valid |= loadbalancer_listener_loadbalancer_refs
                        }
                        break
                case "loadbalancer_pool_back_refs":
                        err = json.Unmarshal(value, &obj.loadbalancer_pool_back_refs)
                        if err == nil {
                                obj.valid |= loadbalancer_listener_loadbalancer_pool_back_refs
                        }
                        break
                }
                if err != nil {
                        return err
                }
        }
        return nil
}

func (obj *LoadbalancerListener) UpdateObject() ([]byte, error) {
        msg := map[string]*json.RawMessage {
        }
        err := obj.MarshalId(msg)
        if err != nil {
                return nil, err
        }

        if obj.modified & loadbalancer_listener_loadbalancer_listener_properties != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.loadbalancer_listener_properties)
                if err != nil {
                        return nil, err
                }
                msg["loadbalancer_listener_properties"] = &value
        }

        if obj.modified & loadbalancer_listener_id_perms != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.id_perms)
                if err != nil {
                        return nil, err
                }
                msg["id_perms"] = &value
        }

        if obj.modified & loadbalancer_listener_perms2 != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.perms2)
                if err != nil {
                        return nil, err
                }
                msg["perms2"] = &value
        }

        if obj.modified & loadbalancer_listener_display_name != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.display_name)
                if err != nil {
                        return nil, err
                }
                msg["display_name"] = &value
        }

        if obj.modified & loadbalancer_listener_loadbalancer_refs != 0 {
                if len(obj.loadbalancer_refs) == 0 {
                        var value json.RawMessage
                        value, err := json.Marshal(
                                          make([]contrail.Reference, 0))
                        if err != nil {
                                return nil, err
                        }
                        msg["loadbalancer_refs"] = &value
                } else if !obj.hasReferenceBase("loadbalancer") {
                        var value json.RawMessage
                        value, err := json.Marshal(&obj.loadbalancer_refs)
                        if err != nil {
                                return nil, err
                        }
                        msg["loadbalancer_refs"] = &value
                }
        }


        return json.Marshal(msg)
}

func (obj *LoadbalancerListener) UpdateReferences() error {

        if (obj.modified & loadbalancer_listener_loadbalancer_refs != 0) &&
           len(obj.loadbalancer_refs) > 0 &&
           obj.hasReferenceBase("loadbalancer") {
                err := obj.UpdateReference(
                        obj, "loadbalancer",
                        obj.loadbalancer_refs,
                        obj.baseMap["loadbalancer"])
                if err != nil {
                        return err
                }
        }

        return nil
}

func LoadbalancerListenerByName(c contrail.ApiClient, fqn string) (*LoadbalancerListener, error) {
    obj, err := c.FindByName("loadbalancer-listener", fqn)
    if err != nil {
        return nil, err
    }
    return obj.(*LoadbalancerListener), nil
}

func LoadbalancerListenerByUuid(c contrail.ApiClient, uuid string) (*LoadbalancerListener, error) {
    obj, err := c.FindByUuid("loadbalancer-listener", uuid)
    if err != nil {
        return nil, err
    }
    return obj.(*LoadbalancerListener), nil
}
