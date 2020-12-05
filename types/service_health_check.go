//
// Automatically generated. DO NOT EDIT.
//

package types

import (
        "encoding/json"

        "github.com/Juniper/contrail-go-api"
)

const (
	service_health_check_service_health_check_properties uint64 = 1 << iota
	service_health_check_id_perms
	service_health_check_perms2
	service_health_check_display_name
	service_health_check_service_instance_refs
	service_health_check_virtual_machine_interface_back_refs
)

type ServiceHealthCheck struct {
        contrail.ObjectBase
	service_health_check_properties ServiceHealthCheckType
	id_perms IdPermsType
	perms2 PermType2
	display_name string
	service_instance_refs contrail.ReferenceList
	virtual_machine_interface_back_refs contrail.ReferenceList
        valid uint64
        modified uint64
        baseMap map[string]contrail.ReferenceList
}

func (obj *ServiceHealthCheck) GetType() string {
        return "service-health-check"
}

func (obj *ServiceHealthCheck) GetDefaultParent() []string {
        name := []string{"default-domain", "default-project"}
        return name
}

func (obj *ServiceHealthCheck) GetDefaultParentType() string {
        return "project"
}

func (obj *ServiceHealthCheck) SetName(name string) {
        obj.VSetName(obj, name)
}

func (obj *ServiceHealthCheck) SetParent(parent contrail.IObject) {
        obj.VSetParent(obj, parent)
}

func (obj *ServiceHealthCheck) storeReferenceBase(
        name string, refList contrail.ReferenceList) {
        if obj.baseMap == nil {
                obj.baseMap = make(map[string]contrail.ReferenceList)
        }
        refCopy := make(contrail.ReferenceList, len(refList))
        copy(refCopy, refList)
        obj.baseMap[name] = refCopy
}

func (obj *ServiceHealthCheck) hasReferenceBase(name string) bool {
        if obj.baseMap == nil {
                return false
        }
        _, exists := obj.baseMap[name]
        return exists
}

func (obj *ServiceHealthCheck) UpdateDone() {
        obj.modified = 0
        obj.baseMap = nil
}


func (obj *ServiceHealthCheck) GetServiceHealthCheckProperties() ServiceHealthCheckType {
        return obj.service_health_check_properties
}

func (obj *ServiceHealthCheck) SetServiceHealthCheckProperties(value *ServiceHealthCheckType) {
        obj.service_health_check_properties = *value
        obj.modified |= service_health_check_service_health_check_properties
}

func (obj *ServiceHealthCheck) GetIdPerms() IdPermsType {
        return obj.id_perms
}

func (obj *ServiceHealthCheck) SetIdPerms(value *IdPermsType) {
        obj.id_perms = *value
        obj.modified |= service_health_check_id_perms
}

func (obj *ServiceHealthCheck) GetPerms2() PermType2 {
        return obj.perms2
}

func (obj *ServiceHealthCheck) SetPerms2(value *PermType2) {
        obj.perms2 = *value
        obj.modified |= service_health_check_perms2
}

func (obj *ServiceHealthCheck) GetDisplayName() string {
        return obj.display_name
}

func (obj *ServiceHealthCheck) SetDisplayName(value string) {
        obj.display_name = value
        obj.modified |= service_health_check_display_name
}

func (obj *ServiceHealthCheck) readServiceInstanceRefs() error {
        if !obj.IsTransient() &&
                (obj.valid & service_health_check_service_instance_refs == 0) {
                err := obj.GetField(obj, "service_instance_refs")
                if err != nil {
                        return err
                }
        }
        return nil
}

func (obj *ServiceHealthCheck) GetServiceInstanceRefs() (
        contrail.ReferenceList, error) {
        err := obj.readServiceInstanceRefs()
        if err != nil {
                return nil, err
        }
        return obj.service_instance_refs, nil
}

func (obj *ServiceHealthCheck) AddServiceInstance(
        rhs *ServiceInstance, data ServiceInterfaceTag) error {
        err := obj.readServiceInstanceRefs()
        if err != nil {
                return err
        }

        if obj.modified & service_health_check_service_instance_refs == 0 {
                obj.storeReferenceBase("service-instance", obj.service_instance_refs)
        }

        ref := contrail.Reference {
                rhs.GetFQName(), rhs.GetUuid(), rhs.GetHref(), data}
        obj.service_instance_refs = append(obj.service_instance_refs, ref)
        obj.modified |= service_health_check_service_instance_refs
        return nil
}

func (obj *ServiceHealthCheck) DeleteServiceInstance(uuid string) error {
        err := obj.readServiceInstanceRefs()
        if err != nil {
                return err
        }

        if obj.modified & service_health_check_service_instance_refs == 0 {
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
        obj.modified |= service_health_check_service_instance_refs
        return nil
}

func (obj *ServiceHealthCheck) ClearServiceInstance() {
        if (obj.valid & service_health_check_service_instance_refs != 0) &&
           (obj.modified & service_health_check_service_instance_refs == 0) {
                obj.storeReferenceBase("service-instance", obj.service_instance_refs)
        }
        obj.service_instance_refs = make([]contrail.Reference, 0)
        obj.valid |= service_health_check_service_instance_refs
        obj.modified |= service_health_check_service_instance_refs
}

func (obj *ServiceHealthCheck) SetServiceInstanceList(
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


func (obj *ServiceHealthCheck) readVirtualMachineInterfaceBackRefs() error {
        if !obj.IsTransient() &&
                (obj.valid & service_health_check_virtual_machine_interface_back_refs == 0) {
                err := obj.GetField(obj, "virtual_machine_interface_back_refs")
                if err != nil {
                        return err
                }
        }
        return nil
}

func (obj *ServiceHealthCheck) GetVirtualMachineInterfaceBackRefs() (
        contrail.ReferenceList, error) {
        err := obj.readVirtualMachineInterfaceBackRefs()
        if err != nil {
                return nil, err
        }
        return obj.virtual_machine_interface_back_refs, nil
}

func (obj *ServiceHealthCheck) MarshalJSON() ([]byte, error) {
        msg := map[string]*json.RawMessage {
        }
        err := obj.MarshalCommon(msg)
        if err != nil {
                return nil, err
        }

        if obj.modified & service_health_check_service_health_check_properties != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.service_health_check_properties)
                if err != nil {
                        return nil, err
                }
                msg["service_health_check_properties"] = &value
        }

        if obj.modified & service_health_check_id_perms != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.id_perms)
                if err != nil {
                        return nil, err
                }
                msg["id_perms"] = &value
        }

        if obj.modified & service_health_check_perms2 != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.perms2)
                if err != nil {
                        return nil, err
                }
                msg["perms2"] = &value
        }

        if obj.modified & service_health_check_display_name != 0 {
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

        return json.Marshal(msg)
}

func (obj *ServiceHealthCheck) UnmarshalJSON(body []byte) error {
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
                case "service_health_check_properties":
                        err = json.Unmarshal(value, &obj.service_health_check_properties)
                        if err == nil {
                                obj.valid |= service_health_check_service_health_check_properties
                        }
                        break
                case "id_perms":
                        err = json.Unmarshal(value, &obj.id_perms)
                        if err == nil {
                                obj.valid |= service_health_check_id_perms
                        }
                        break
                case "perms2":
                        err = json.Unmarshal(value, &obj.perms2)
                        if err == nil {
                                obj.valid |= service_health_check_perms2
                        }
                        break
                case "display_name":
                        err = json.Unmarshal(value, &obj.display_name)
                        if err == nil {
                                obj.valid |= service_health_check_display_name
                        }
                        break
                case "virtual_machine_interface_back_refs":
                        err = json.Unmarshal(value, &obj.virtual_machine_interface_back_refs)
                        if err == nil {
                                obj.valid |= service_health_check_virtual_machine_interface_back_refs
                        }
                        break
                case "service_instance_refs": {
                        type ReferenceElement struct {
                                To []string
                                Uuid string
                                Href string
                                Attr ServiceInterfaceTag
                        }
                        var array []ReferenceElement
                        err = json.Unmarshal(value, &array)
                        if err != nil {
                            break
                        }
                        obj.valid |= service_health_check_service_instance_refs
                        obj.service_instance_refs = make(contrail.ReferenceList, 0)
                        for _, element := range array {
                                ref := contrail.Reference {
                                        element.To,
                                        element.Uuid,
                                        element.Href,
                                        element.Attr,
                                }
                                obj.service_instance_refs = append(obj.service_instance_refs, ref)
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

func (obj *ServiceHealthCheck) UpdateObject() ([]byte, error) {
        msg := map[string]*json.RawMessage {
        }
        err := obj.MarshalId(msg)
        if err != nil {
                return nil, err
        }

        if obj.modified & service_health_check_service_health_check_properties != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.service_health_check_properties)
                if err != nil {
                        return nil, err
                }
                msg["service_health_check_properties"] = &value
        }

        if obj.modified & service_health_check_id_perms != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.id_perms)
                if err != nil {
                        return nil, err
                }
                msg["id_perms"] = &value
        }

        if obj.modified & service_health_check_perms2 != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.perms2)
                if err != nil {
                        return nil, err
                }
                msg["perms2"] = &value
        }

        if obj.modified & service_health_check_display_name != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.display_name)
                if err != nil {
                        return nil, err
                }
                msg["display_name"] = &value
        }

        if obj.modified & service_health_check_service_instance_refs != 0 {
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


        return json.Marshal(msg)
}

func (obj *ServiceHealthCheck) UpdateReferences() error {

        if (obj.modified & service_health_check_service_instance_refs != 0) &&
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

        return nil
}

func ServiceHealthCheckByName(c contrail.ApiClient, fqn string) (*ServiceHealthCheck, error) {
    obj, err := c.FindByName("service-health-check", fqn)
    if err != nil {
        return nil, err
    }
    return obj.(*ServiceHealthCheck), nil
}

func ServiceHealthCheckByUuid(c contrail.ApiClient, uuid string) (*ServiceHealthCheck, error) {
    obj, err := c.FindByUuid("service-health-check", uuid)
    if err != nil {
        return nil, err
    }
    return obj.(*ServiceHealthCheck), nil
}
