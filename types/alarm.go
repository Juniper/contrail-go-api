//
// Automatically generated. DO NOT EDIT.
//

package types

import (
        "encoding/json"

        "github.com/michaelhenkel/contrail-go-api"
)

const (
	alarm_uve_keys uint64 = 1 << iota
	alarm_alarm_severity
	alarm_alarm_rules
	alarm_id_perms
	alarm_perms2
	alarm_display_name
)

type Alarm struct {
        contrail.ObjectBase
	uve_keys string
	alarm_severity int
	alarm_rules AlarmRule
	id_perms IdPermsType
	perms2 PermType2
	display_name string
        valid uint64
        modified uint64
        baseMap map[string]contrail.ReferenceList
}

func (obj *Alarm) GetType() string {
        return "alarm"
}

func (obj *Alarm) GetDefaultParent() []string {
        name := []string{"default-global-system-config"}
        return name
}

func (obj *Alarm) GetDefaultParentType() string {
        return "global-system-config"
}

func (obj *Alarm) SetName(name string) {
        obj.VSetName(obj, name)
}

func (obj *Alarm) SetParent(parent contrail.IObject) {
        obj.VSetParent(obj, parent)
}

func (obj *Alarm) storeReferenceBase(
        name string, refList contrail.ReferenceList) {
        if obj.baseMap == nil {
                obj.baseMap = make(map[string]contrail.ReferenceList)
        }
        refCopy := make(contrail.ReferenceList, len(refList))
        copy(refCopy, refList)
        obj.baseMap[name] = refCopy
}

func (obj *Alarm) hasReferenceBase(name string) bool {
        if obj.baseMap == nil {
                return false
        }
        _, exists := obj.baseMap[name]
        return exists
}

func (obj *Alarm) UpdateDone() {
        obj.modified = 0
        obj.baseMap = nil
}


func (obj *Alarm) GetUveKeys() string {
        return obj.uve_keys
}

func (obj *Alarm) SetUveKeys(value string) {
        obj.uve_keys = value
        obj.modified |= alarm_uve_keys
}

func (obj *Alarm) GetAlarmSeverity() int {
        return obj.alarm_severity
}

func (obj *Alarm) SetAlarmSeverity(value int) {
        obj.alarm_severity = value
        obj.modified |= alarm_alarm_severity
}

func (obj *Alarm) GetAlarmRules() AlarmRule {
        return obj.alarm_rules
}

func (obj *Alarm) SetAlarmRules(value *AlarmRule) {
        obj.alarm_rules = *value
        obj.modified |= alarm_alarm_rules
}

func (obj *Alarm) GetIdPerms() IdPermsType {
        return obj.id_perms
}

func (obj *Alarm) SetIdPerms(value *IdPermsType) {
        obj.id_perms = *value
        obj.modified |= alarm_id_perms
}

func (obj *Alarm) GetPerms2() PermType2 {
        return obj.perms2
}

func (obj *Alarm) SetPerms2(value *PermType2) {
        obj.perms2 = *value
        obj.modified |= alarm_perms2
}

func (obj *Alarm) GetDisplayName() string {
        return obj.display_name
}

func (obj *Alarm) SetDisplayName(value string) {
        obj.display_name = value
        obj.modified |= alarm_display_name
}

func (obj *Alarm) MarshalJSON() ([]byte, error) {
        msg := map[string]*json.RawMessage {
        }
        err := obj.MarshalCommon(msg)
        if err != nil {
                return nil, err
        }

        if obj.modified & alarm_uve_keys != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.uve_keys)
                if err != nil {
                        return nil, err
                }
                msg["uve_keys"] = &value
        }

        if obj.modified & alarm_alarm_severity != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.alarm_severity)
                if err != nil {
                        return nil, err
                }
                msg["alarm_severity"] = &value
        }

        if obj.modified & alarm_alarm_rules != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.alarm_rules)
                if err != nil {
                        return nil, err
                }
                msg["alarm_rules"] = &value
        }

        if obj.modified & alarm_id_perms != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.id_perms)
                if err != nil {
                        return nil, err
                }
                msg["id_perms"] = &value
        }

        if obj.modified & alarm_perms2 != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.perms2)
                if err != nil {
                        return nil, err
                }
                msg["perms2"] = &value
        }

        if obj.modified & alarm_display_name != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.display_name)
                if err != nil {
                        return nil, err
                }
                msg["display_name"] = &value
        }

        return json.Marshal(msg)
}

func (obj *Alarm) UnmarshalJSON(body []byte) error {
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
                case "uve_keys":
                        err = json.Unmarshal(value, &obj.uve_keys)
                        if err == nil {
                                obj.valid |= alarm_uve_keys
                        }
                        break
                case "alarm_severity":
                        err = json.Unmarshal(value, &obj.alarm_severity)
                        if err == nil {
                                obj.valid |= alarm_alarm_severity
                        }
                        break
                case "alarm_rules":
                        err = json.Unmarshal(value, &obj.alarm_rules)
                        if err == nil {
                                obj.valid |= alarm_alarm_rules
                        }
                        break
                case "id_perms":
                        err = json.Unmarshal(value, &obj.id_perms)
                        if err == nil {
                                obj.valid |= alarm_id_perms
                        }
                        break
                case "perms2":
                        err = json.Unmarshal(value, &obj.perms2)
                        if err == nil {
                                obj.valid |= alarm_perms2
                        }
                        break
                case "display_name":
                        err = json.Unmarshal(value, &obj.display_name)
                        if err == nil {
                                obj.valid |= alarm_display_name
                        }
                        break
                }
                if err != nil {
                        return err
                }
        }
        return nil
}

func (obj *Alarm) UpdateObject() ([]byte, error) {
        msg := map[string]*json.RawMessage {
        }
        err := obj.MarshalId(msg)
        if err != nil {
                return nil, err
        }

        if obj.modified & alarm_uve_keys != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.uve_keys)
                if err != nil {
                        return nil, err
                }
                msg["uve_keys"] = &value
        }

        if obj.modified & alarm_alarm_severity != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.alarm_severity)
                if err != nil {
                        return nil, err
                }
                msg["alarm_severity"] = &value
        }

        if obj.modified & alarm_alarm_rules != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.alarm_rules)
                if err != nil {
                        return nil, err
                }
                msg["alarm_rules"] = &value
        }

        if obj.modified & alarm_id_perms != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.id_perms)
                if err != nil {
                        return nil, err
                }
                msg["id_perms"] = &value
        }

        if obj.modified & alarm_perms2 != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.perms2)
                if err != nil {
                        return nil, err
                }
                msg["perms2"] = &value
        }

        if obj.modified & alarm_display_name != 0 {
                var value json.RawMessage
                value, err := json.Marshal(&obj.display_name)
                if err != nil {
                        return nil, err
                }
                msg["display_name"] = &value
        }

        return json.Marshal(msg)
}

func (obj *Alarm) UpdateReferences() error {

        return nil
}

func AlarmByName(c contrail.ApiClient, fqn string) (*Alarm, error) {
    obj, err := c.FindByName("alarm", fqn)
    if err != nil {
        return nil, err
    }
    return obj.(*Alarm), nil
}

func AlarmByUuid(c contrail.ApiClient, uuid string) (*Alarm, error) {
    obj, err := c.FindByUuid("alarm", uuid)
    if err != nil {
        return nil, err
    }
    return obj.(*Alarm), nil
}
