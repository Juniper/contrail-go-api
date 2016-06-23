//
// Automatically generated. DO NOT EDIT.
//

package types

type AlarmElement struct {
	Operation string `json:"operation,omitempty"`
	Operand1 string `json:"operand1,omitempty"`
	Operand2 string `json:"operand2,omitempty"`
	Vars []string `json:"vars,omitempty"`
}

func (obj *AlarmElement) AddVars(value string) {
        obj.Vars = append(obj.Vars, value)
}

type AlarmRule struct {
	Rule []AlarmElement `json:"rule,omitempty"`
}

func (obj *AlarmRule) AddRule(value *AlarmElement) {
        obj.Rule = append(obj.Rule, *value)
}
