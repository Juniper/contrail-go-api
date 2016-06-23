//
// Automatically generated. DO NOT EDIT.
//

package types

import (
        "encoding/json"

        "github.com/michaelhenkel/contrail-go-api"
)

const (
	service_appliance_set_service_appliance_set_properties uint64 = 1 << iota
	service_appliance_set_service_appliance_driver
	service_appliance_set_service_appliance_ha_mode
	service_appliance_set_id_perms
	service_appliance_set_perms2
	service_appliance_set_display_name
	service_appliance_set_service_appliances
	service_appliance_set_service_template_back_refs
	service_appliance_set_loadbalancer_pool_back_refs
)

type ServiceApplianceSet struct {
        contrail.ObjectBase
	service_appliance_set_properties KeyValuePairs
	service_appliance_driver string
	service_appliance_ha_mode string
	id_perms IdPermsType
	perms2 PermType2
	display_name string
	service_appliances contrail.ReferenceList
	service_template_back_refs contrail.ReferenceList
	loadbalancer_pool_back_refs contrail.ReferenceList
        valid uint64
        modified uint64
        baseMap map[string]contrail.ReferenceList
}

func (obj *ServiceApplianceSet) GetType() string {
        return "service-appliance-set"
}

func (obj *ServiceApplianceSet) GetDefaultParent() []string {
        name := []string{"default-global-system-config"}
        return name
}

func (obj *ServiceApplianceSet) GetDefaultParentType() string {
        return "global-system-config"
}

func (obj *ServiceApplianceSet) SetName(name string) {
        obj.VSetName(obj, name)
}

func (obj *ServiceApplianceSet) SetParent(parent contrail.IObject) {
        obj.VSetParent(obj, parent)
}

func (obj *ServiceApplianceSet) storeReferenceBase(
        name string, refList contrail.ReferenceList) {
        if obj.baseMap == nil {
                obj.baseMap = make(map[string]contrail.ReferenceList)
        }
        refCopy := make(contrail.ReferenceList, len(refList))
        copy(refCopy, refList)
        obj.baseMap[name] = refCopy
}

func (obj *ServiceApplianceSet) hasReferenceBase(name string) bool {
        if obj.baseMap == nil {
                return false
        }
        _, exists := obj.baseMap[name]
        return exists
}

func (obj *ServiceApplianceSet) UpdateDone() {
        obj.modified = 0
        obj.baseMap = nil
}


func (obj *ServiceApplianceSet) GetServiceApplianceSetProperties() KeyValuePairs {
        return obj.service_appliance_set_properties
}

func (obj *ServiceApplianceSet) SetServiceApplianceSetProperties(value *KeyValuePairs) {
        obj.service_appliance_set_properties = *value
        obj.modified |= service_appliance_set_service_appliance_set_properties
}

func (obj *ServiceApplianceSet) GetServiceApplianceDriver() string {
        return obj.service_appliance_driver
}

func (obj *ServiceApplianceSet) SetServiceApplianceDriver(value string) {
        obj.service_appliance_driver = value
        obj.modified |= service_appliance_set_service_appliance_driver
}

func (obj *ServiceApplianceSet) GetServiceApplianceHaMode() string {
        return obj.service_appliance_ha_mode
}

func (obj *ServiceApplianceSet) SetServiceApplianceHaMode(value string) {
        obj.service_appliance_ha_mode = value
        obj.modified |= service_appliance_set_service_appliance_ha_mode
}

func (obj *ServiceApplianceSet) GetIdPerms() IdPermsType {
        return obj.id_perms
}

func (obj *ServiceApplianceSet) SetIdPerms(value *IdPermsType) {
        obj.id_perms = *value
        obj.modified |= service_appliance_set_id_perms
}

func (obj *ServiceApplianceSet) GetPerms2() PermType2 {
        return obj.perms2
}

func (obj *ServiceApplianceSet) SetPerms2(value *PermType2) {
        obj.perms2 = *value
        obj.modified |= service_appliance_set_perms2
}

func (obj *ServiceApplianceSet) GetDisplayName() string {
        return obj.display_name
}

func (obj *ServiceApplianceSet) SetDisplayName(value string) {
        obj.display_name = value
        obj.modified |= service_appliance_set_display_name
}

func (obj *ServiceApplianceSet) readServiceAppliances() error {
        if !obj.IsTransient() &&
                (obj.valid & service_appliance_set_service_appliances == 0) {
                err := obj.GetField(obj, "service_appliances")
                if err != nil {
                        return err
                }
        }
        return nil
}

func (obj *ServiceApplianceSet) GetServiceAppliances() (
        contrail.ReferenceList, error) {
        err := obj.readServiceAppliances()
        if err != nil {
                return nil, err
        }
        return obj.service_appliances, nil
}

func (obj *ServiceApplianceSet) readServiceTemplateBackRefs() error {
        if !obj.IsTransient() &&
                (obj.valid & service_appliance_set_service_template_back_refs == 0) {
                err := obj.GetField(obj, "service_template_back_refs")
                if err != nil {
                        return err
                }
        }
        return nil
}

func (obj *ServiceApplianceSet) GetServiceTemplateBackRefs() (
        contrail.ReferenceList, error) {
        err := obj.readServiceTemplateBackRefs()
        if err != nil {
                return nil, err
        }
        return obj.service_template_back_refs, nil
}

func (obj *ServiceApplianceSet) readLoadbalancerPoolBackRefs() error {
        if !obj.IsTransient() &&
                (obj.valid & service_appliance_set_loadbalancer_pool_back_refs == 0) {
                err := obj.GetField(obj, "loadbalancer_pool_back_refs")
                if err != nil {
                        return err
                }
        }
        return nil
}

func (obj *ServiceApplianceSet) GetLoadbalancerPoolBackRefs() (
        contrail.ReferenceList, error) {
        err := obj.readLoadbalancerPoolBackRefs()
        if err != nil {
                return nil, err
        }
        return obj.loadbalancer_pool_back_refs, nil
}

func (obj *ServiceApplianceSet) MarshalJSON() ([]byte, error) {
        msg := map[string]*json.RawMessage {
        }
        err := obj.MarshalCommon(msg)
        if err != nil {
                return nil, err
        }

        if obj.modified & service_appliance_set_service_appliance_set_properties != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.service_appliance_set_properties)
                if err != nil {
                        return nil, err
                }
                msg["service_appliance_set_properties"] = &value
        }

        if obj.modified & service_appliance_set_service_appliance_driver != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.service_appliance_driver)
                if err != nil {
                        return nil, err
                }
                msg["service_appliance_driver"] = &value
        }

        if obj.modified & service_appliance_set_service_appliance_ha_mode != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.service_appliance_ha_mode)
                if err != nil {
                        return nil, err
                }
                msg["service_appliance_ha_mode"] = &value
        }

        if obj.modified & service_appliance_set_id_perms != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.id_perms)
                if err != nil {
                        return nil, err
                }
                msg["id_perms"] = &value
        }

        if obj.modified & service_appliance_set_perms2 != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.perms2)
                if err != nil {
                        return nil, err
                }
                msg["perms2"] = &value
        }

        if obj.modified & service_appliance_set_display_name != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.display_name)
                if err != nil {
                        return nil, err
                }
                msg["display_name"] = &value
        }

        return json.Marshal(msg)
}

func (obj *ServiceApplianceSet) UnmarshalJSON(body []byte) error {
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
                case "service_appliance_set_properties":
                        err = json.Unmarshal(value, &obj.service_appliance_set_properties)
                        if err == nil {
                                obj.valid |= service_appliance_set_service_appliance_set_properties
                        }
                        break
                case "service_appliance_driver":
                        err = json.Unmarshal(value, &obj.service_appliance_driver)
                        if err == nil {
                                obj.valid |= service_appliance_set_service_appliance_driver
                        }
                        break
                case "service_appliance_ha_mode":
                        err = json.Unmarshal(value, &obj.service_appliance_ha_mode)
                        if err == nil {
                                obj.valid |= service_appliance_set_service_appliance_ha_mode
                        }
                        break
                case "id_perms":
                        err = json.Unmarshal(value, &obj.id_perms)
                        if err == nil {
                                obj.valid |= service_appliance_set_id_perms
                        }
                        break
                case "perms2":
                        err = json.Unmarshal(value, &obj.perms2)
                        if err == nil {
                                obj.valid |= service_appliance_set_perms2
                        }
                        break
                case "display_name":
                        err = json.Unmarshal(value, &obj.display_name)
                        if err == nil {
                                obj.valid |= service_appliance_set_display_name
                        }
                        break
                case "service_appliances":
                        err = json.Unmarshal(value, &obj.service_appliances)
                        if err == nil {
                                obj.valid |= service_appliance_set_service_appliances
                        }
                        break
                case "service_template_back_refs":
                        err = json.Unmarshal(value, &obj.service_template_back_refs)
                        if err == nil {
                                obj.valid |= service_appliance_set_service_template_back_refs
                        }
                        break
                case "loadbalancer_pool_back_refs":
                        err = json.Unmarshal(value, &obj.loadbalancer_pool_back_refs)
                        if err == nil {
                                obj.valid |= service_appliance_set_loadbalancer_pool_back_refs
                        }
                        break
                }
                if err != nil {
                        return err
                }
        }
        return nil
}

func (obj *ServiceApplianceSet) UpdateObject() ([]byte, error) {
        msg := map[string]*json.RawMessage {
        }
        err := obj.MarshalId(msg)
        if err != nil {
                return nil, err
        }

        if obj.modified & service_appliance_set_service_appliance_set_properties != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.service_appliance_set_properties)
                if err != nil {
                        return nil, err
                }
                msg["service_appliance_set_properties"] = &value
        }

        if obj.modified & service_appliance_set_service_appliance_driver != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.service_appliance_driver)
                if err != nil {
                        return nil, err
                }
                msg["service_appliance_driver"] = &value
        }

        if obj.modified & service_appliance_set_service_appliance_ha_mode != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.service_appliance_ha_mode)
                if err != nil {
                        return nil, err
                }
                msg["service_appliance_ha_mode"] = &value
        }

        if obj.modified & service_appliance_set_id_perms != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.id_perms)
                if err != nil {
                        return nil, err
                }
                msg["id_perms"] = &value
        }

        if obj.modified & service_appliance_set_perms2 != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.perms2)
                if err != nil {
                        return nil, err
                }
                msg["perms2"] = &value
        }

        if obj.modified & service_appliance_set_display_name != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.display_name)
                if err != nil {
                        return nil, err
                }
                msg["display_name"] = &value
        }

        return json.Marshal(msg)
}

func (obj *ServiceApplianceSet) UpdateReferences() error {

        return nil
}

func ServiceApplianceSetByName(c contrail.ApiClient, fqn string) (*ServiceApplianceSet, error) {
    obj, err := c.FindByName("service-appliance-set", fqn)
    if err != nil {
        return nil, err
    }
    return obj.(*ServiceApplianceSet), nil
}

func ServiceApplianceSetByUuid(c contrail.ApiClient, uuid string) (*ServiceApplianceSet, error) {
    obj, err := c.FindByUuid("service-appliance-set", uuid)
    if err != nil {
        return nil, err
    }
    return obj.(*ServiceApplianceSet), nil
}
