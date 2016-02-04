//
// Copyright (c) 2014 Juniper Networks, Inc. All rights reserved.
//

package contrail

import (
	"fmt"
	"testing"
)

type MockObject struct {
	ObjectBase
}

func (*MockObject) SetName(name string) {
}
func (*MockObject) GetDefaultParent() []string {
	return []string{"root"}
}
func (*MockObject) GetDefaultParentType() string {
	return "none"
}
func (*MockObject) GetType() string {
	return "mock"
}
func (*MockObject) UpdateObject() ([]byte, error) {
	return nil, nil
}
func (*MockObject) UpdateReferences() error {
	return nil
}
func (*MockObject) UpdateDone() {
}

type MockClient struct {
	messages []ReferenceUpdateMsg
}

func (*MockClient) GetField(obj IObject, field string) error {
	return nil
}
func (m *MockClient) UpdateReference(msg *ReferenceUpdateMsg) error {
	m.messages = append(m.messages, *msg)
	return nil
}

func TestUpdateAddBefore(t *testing.T) {
	obj := MockObject{}
	client := MockClient{}

	obj.SetClient(&client)
	list1 := ReferenceList{
		Reference{[]string{"y"}, "2", "", nil},
		Reference{[]string{"x"}, "1", "", nil},
	}
	list2 := ReferenceList{
		Reference{[]string{"y"}, "2", "", nil},
	}
	obj.UpdateReference(&obj, "foo", list1, list2)
	if len(client.messages) != 1 {
		t.Error("No messages generated")
	}
	msg := client.messages[0]
	if msg.RefUuid != "1" || msg.Operation != "ADD" {
		t.Errorf("op: %s, id=%s", msg.Operation, msg.RefUuid)
	}
}

func TestUpdateAddAfter(t *testing.T) {
	obj := MockObject{}
	client := MockClient{}

	obj.SetClient(&client)
	list1 := ReferenceList{
		Reference{[]string{"z"}, "3", "", nil},
		Reference{[]string{"x"}, "9", "", nil},
		Reference{[]string{"y"}, "2", "", nil},
	}
	list2 := ReferenceList{
		Reference{[]string{"y"}, "2", "", nil},
		Reference{[]string{"z"}, "3", "", nil},
	}
	obj.UpdateReference(&obj, "foo", list1, list2)
	if len(client.messages) != 1 {
		t.Error("No messages generated")
	}
	msg := client.messages[0]
	if msg.RefUuid != "9" || msg.Operation != "ADD" {
		t.Errorf("op: %s, id=%s", msg.Operation, msg.RefUuid)
	}
}

func TestUpdateDeleteMid(t *testing.T) {
	obj := MockObject{}
	client := MockClient{}

	obj.SetClient(&client)
	list1 := ReferenceList{
		Reference{[]string{"z"}, "3", "", nil},
		Reference{[]string{"x"}, "9", "", nil},
		Reference{[]string{"y"}, "2", "", nil},
	}
	list2 := ReferenceList{
		Reference{[]string{"y"}, "2", "", nil},
		Reference{[]string{"z"}, "3", "", nil},
		Reference{[]string{"x"}, "9", "", nil},
		Reference{[]string{"w"}, "5", "", nil},
	}
	obj.UpdateReference(&obj, "foo", list1, list2)
	if len(client.messages) != 1 {
		t.Error("No messages generated")
	}
	msg := client.messages[0]
	if msg.RefUuid != "5" || msg.Operation != "DELETE" {
		t.Errorf("op: %s, id=%s", msg.Operation, msg.RefUuid)
	}
}

func TestUpdateAttrSimpleEqual(t *testing.T) {
	obj := MockObject{}
	client := MockClient{}

	obj.SetClient(&client)

	type TestAttr string
	var a1, a2 TestAttr
	a1 = "foo"
	a2 = "foo"
	list1 := ReferenceList{
		Reference{[]string{"y"}, "2", "", &a1},
	}
	list2 := ReferenceList{
		Reference{[]string{"y"}, "2", "", &a2},
	}
	obj.UpdateReference(&obj, "foo", list1, list2)
	if len(client.messages) != 0 {
		t.Error("Expected empty change list")
	}
}

func TestUpdateAttrSimple(t *testing.T) {
	obj := MockObject{}
	client := MockClient{}

	obj.SetClient(&client)

	type TestAttr string
	var a1, a2 TestAttr
	a1 = "foo"
	a2 = "bar"
	list1 := ReferenceList{
		Reference{[]string{"y"}, "2", "", &a1},
	}
	list2 := ReferenceList{
		Reference{[]string{"y"}, "2", "", &a2},
	}
	obj.UpdateReference(&obj, "foo", list1, list2)
	if len(client.messages) != 1 {
		t.Error("No messages generated")
	}
	msg := client.messages[0]
	if msg.RefUuid != "2" || msg.Operation != "ADD" {
		t.Error(fmt.Sprintf(
			"op: %s, id=%s", msg.Operation, msg.RefUuid))
	}
}

func TestUpdateAttrStructEqual(t *testing.T) {
	obj := MockObject{}
	client := MockClient{}

	obj.SetClient(&client)

	type TestAttr struct {
		X string
		Y string
	}
	a1 := TestAttr{
		X: "x",
		Y: "foo",
	}
	a2 := TestAttr{
		X: "x",
		Y: "foo",
	}
	list1 := ReferenceList{
		Reference{[]string{"y"}, "2", "", &a1},
	}
	list2 := ReferenceList{
		Reference{[]string{"y"}, "2", "", &a2},
	}
	obj.UpdateReference(&obj, "foo", list1, list2)
	if len(client.messages) != 0 {
		t.Error("Expected empty change list")
	}
}

func TestUpdateAttrStruct(t *testing.T) {
	obj := MockObject{}
	client := MockClient{}

	obj.SetClient(&client)

	type TestAttr struct {
		X string
		Y string
	}
	a1 := TestAttr{
		X: "x",
		Y: "foo",
	}
	a2 := TestAttr{
		X: "x",
		Y: "bar",
	}
	list1 := ReferenceList{
		Reference{[]string{"y"}, "2", "", &a1},
	}
	list2 := ReferenceList{
		Reference{[]string{"y"}, "2", "", &a2},
	}
	obj.UpdateReference(&obj, "foo", list1, list2)
	if len(client.messages) != 1 {
		t.Error("No messages generated")
	}
	msg := client.messages[0]
	if msg.RefUuid != "2" || msg.Operation != "ADD" {
		t.Error(fmt.Sprintf(
			"op: %s, id=%s", msg.Operation, msg.RefUuid))
	}

}

func TestUpdateAttrArrayEqual(t *testing.T) {
	obj := MockObject{}
	client := MockClient{}

	obj.SetClient(&client)

	type TestAttr struct {
		X string
		Y string
	}
	a1 := []TestAttr{
		TestAttr{X: "x", Y: "foo"},
	}
	a2 := []TestAttr{
		TestAttr{X: "x", Y: "foo"},
	}
	list1 := ReferenceList{
		Reference{[]string{"y"}, "2", "", &a1},
	}
	list2 := ReferenceList{
		Reference{[]string{"y"}, "2", "", &a2},
	}
	obj.UpdateReference(&obj, "foo", list1, list2)
	if len(client.messages) != 0 {
		t.Error("Expected empty change list")
	}
}

func TestUpdateAttrArray(t *testing.T) {
	obj := MockObject{}
	client := MockClient{}

	obj.SetClient(&client)

	type TestAttr struct {
		X string
		Y string
	}
	a1 := []TestAttr{
		TestAttr{X: "x", Y: "foo"},
	}
	a2 := []TestAttr{
		TestAttr{X: "x", Y: "foo"},
		TestAttr{X: "y", Y: "bar"},
	}

	list1 := ReferenceList{
		Reference{[]string{"y"}, "2", "", &a1},
	}
	list2 := ReferenceList{
		Reference{[]string{"y"}, "2", "", &a2},
	}
	obj.UpdateReference(&obj, "foo", list1, list2)
	if len(client.messages) != 1 {
		t.Error("No messages generated")
	}
	msg := client.messages[0]
	if msg.RefUuid != "2" || msg.Operation != "ADD" {
		t.Errorf("op: %s, id=%s", msg.Operation, msg.RefUuid)
	}

}
