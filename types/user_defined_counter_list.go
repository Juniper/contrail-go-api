//
// Automatically generated. DO NOT EDIT.
//

package types

type UserDefinedCounter struct {
	Name string `json:"name,omitempty"`
	Pattern string `json:"pattern,omitempty"`
}

type UserDefinedCounterList struct {
	Counter []UserDefinedCounter `json:"counter,omitempty"`
}

func (obj *UserDefinedCounterList) AddCounter(value *UserDefinedCounter) {
        obj.Counter = append(obj.Counter, *value)
}
