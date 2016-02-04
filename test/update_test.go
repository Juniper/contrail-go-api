//
// Test how the generated types package handles updates of references.
//
package contrail_test

import (
	"bytes"
	"encoding/json"
	"testing"
	"text/template"

	"github.com/pborman/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/Juniper/contrail-go-api"
	"github.com/Juniper/contrail-go-api/types"
)

type getArguments struct {
	obj   contrail.IObject
	field string
}

type mockObjectUpdater struct {
	t       *testing.T
	readMap map[getArguments]json.RawMessage
	updates []*contrail.ReferenceUpdateMsg
}

func newMockObjectUpdater(t *testing.T) *mockObjectUpdater {
	mock := new(mockObjectUpdater)
	mock.t = t
	mock.readMap = make(map[getArguments]json.RawMessage, 0)
	mock.updates = make([]*contrail.ReferenceUpdateMsg, 0)
	return mock
}

func (m *mockObjectUpdater) Expect(obj contrail.IObject, field string, result json.RawMessage) {
	m.readMap[getArguments{obj, field}] = result
}

func (m *mockObjectUpdater) GetField(obj contrail.IObject, field string) error {
	data, exists := m.readMap[getArguments{obj, field}]
	require.True(m.t, exists, "%s %s", obj.GetName(), field)
	err := json.Unmarshal(data, obj)
	assert.NoError(m.t, err)
	return err
}

func (m *mockObjectUpdater) UpdateReference(msg *contrail.ReferenceUpdateMsg) error {
	m.updates = append(m.updates, msg)
	return nil
}

type testUpdateFixture struct {
	updater  *mockObjectUpdater
	network  *types.VirtualNetwork
	policies []*types.NetworkPolicy
}

var (
	policyNames = [...]string{"p1", "p2", "p3"}
)

func (f *testUpdateFixture) SetUp(t *testing.T) {
	f.updater = newMockObjectUpdater(t)

	network := new(types.VirtualNetwork)
	network.SetFQName("project", []string{"default-domain", "default-project", "update-test"})
	network.SetUuid(uuid.New())
	network.SetClient(f.updater)
	f.network = network

	f.policies = make([]*types.NetworkPolicy, len(policyNames))
	for i, poName := range policyNames {
		policy := new(types.NetworkPolicy)
		policy.SetFQName("project", []string{"default-domain", "default-project", poName})
		policy.SetUuid(uuid.New())
		policy.SetClient(f.updater)
		f.policies[i] = policy
	}
}

type PolicyRefsData struct {
	Name       string
	Uuid       string
	References []contrail.Reference
}

// Build a json encoded message that represents the response from the contrail
// api-server to a request to read the network-policy refs.
func BuildNetworkPolicyRefsResponse(t *testing.T, info *PolicyRefsData) []byte {
	functions := template.FuncMap{
		"last": func(x, y int) bool {
			return x == y-1
		},
		"quoteAndJoin": func(s []string) string {
			result := ""
			for i, piece := range s {
				result = result + "\"" + piece + "\""
				if i < len(s)-1 {
					result = result + ","
				}
			}
			return result
		},
	}
	format := `
	{
		"fq_name": ["default-domain", "default-project", "{{.Name}}"],
		"uuid": "{{.Uuid}}",
		"name": "{{.Name}}"{{ if .References }},{{ $length := (len .References) }}
		"network_policy_refs": [{{ range $index, $item := .References }}
			{
				"href": "http://localhost:8082/network-policy/{{$item.Uuid}}",
				"to": [{{quoteAndJoin .To}}],
				"uuid": "{{$item.Uuid}}"
			}{{ if not (last $index $length) }},{{ end }}{{ end }}
		]{{ end }}
	}
	`
	var buffer bytes.Buffer
	tmpl := template.Must(template.New("json").Funcs(functions).Parse(format))
	err := tmpl.Execute(&buffer, info)
	require.NoError(t, err)
	return buffer.Bytes()
}

func updateDoesNotInclude(t *testing.T, objJson []byte, field string) {
	var m map[string]*json.RawMessage
	err := json.Unmarshal(objJson, &m)
	require.NoError(t, err)
	_, present := m[field]
	assert.False(t, present)
}

func jsonMessageAdd(content []byte, key, value string) ([]byte, error) {
	var m map[string]*json.RawMessage
	err := json.Unmarshal(content, &m)
	if err != nil {
		return nil, err
	}
	var jsonValue json.RawMessage
	jsonValue, err = json.Marshal(value)
	if err != nil {
		return nil, err
	}
	m[key] = &jsonValue
	return json.Marshal(m)
}

//
// Sequence of events:
// a) Add<Reference> reads the object
// b) stores an empty list in the base reference list
// c) generate incremental update
//
func TestUpdateAddFirstReference(t *testing.T) {
	fixture := new(testUpdateFixture)
	fixture.SetUp(t)

	network := fixture.network
	jsonData := BuildNetworkPolicyRefsResponse(
		t, &PolicyRefsData{network.GetName(), network.GetUuid(), []contrail.Reference{}})
	fixture.updater.Expect(network, "network_policy_refs", jsonData)

	network.AddNetworkPolicy(fixture.policies[0], types.VirtualNetworkPolicyType{
		Sequence: &types.SequenceType{10, 0}})

	objJson, err := network.UpdateObject()
	require.NoError(t, err)
	updateDoesNotInclude(t, objJson, "network_policy_refs")

	err = network.UpdateReferences()
	require.NoError(t, err)
	assert.Len(t, fixture.updater.updates, 1)
	if len(fixture.updater.updates) > 0 {
		update := fixture.updater.updates[0]
		assert.Equal(t, "ADD", update.Operation)
		assert.Equal(t, fixture.policies[0].GetUuid(), update.RefUuid)
	}
}

func TestUpdateAddAndDeleteReference(t *testing.T) {
	fixture := new(testUpdateFixture)
	fixture.SetUp(t)

	network := fixture.network
	jsonData := BuildNetworkPolicyRefsResponse(
		t, &PolicyRefsData{
			network.GetName(),
			network.GetUuid(),
			[]contrail.Reference{
				{To: fixture.policies[0].GetFQName(), Uuid: fixture.policies[0].GetUuid()},
			},
		})

	fixture.updater.Expect(network, "network_policy_refs", jsonData)

	network.DeleteNetworkPolicy(fixture.policies[0].GetUuid())
	network.AddNetworkPolicy(fixture.policies[1], types.VirtualNetworkPolicyType{
		Sequence: &types.SequenceType{10, 0}})

	refs, err := network.GetNetworkPolicyRefs()
	require.NoError(t, err)
	assert.Len(t, refs, 1)

	objJson, err := network.UpdateObject()
	require.NoError(t, err)
	updateDoesNotInclude(t, objJson, "network_policy_refs")

	err = network.UpdateReferences()
	require.NoError(t, err)
	assert.Len(t, fixture.updater.updates, 2)
	if len(fixture.updater.updates) > 0 {
		update := fixture.updater.updates[0]
		assert.Equal(t, "DELETE", update.Operation)
		assert.Equal(t, fixture.policies[0].GetUuid(), update.RefUuid)
	}
	if len(fixture.updater.updates) > 1 {
		update := fixture.updater.updates[1]
		assert.Equal(t, "ADD", update.Operation)
		assert.Equal(t, fixture.policies[1].GetUuid(), update.RefUuid)
	}

}

func TestUpdateRemoveOneReference(t *testing.T) {
	fixture := new(testUpdateFixture)
	fixture.SetUp(t)

	network := fixture.network
	jsonData := BuildNetworkPolicyRefsResponse(
		t, &PolicyRefsData{
			network.GetName(),
			network.GetUuid(),
			[]contrail.Reference{
				{To: fixture.policies[0].GetFQName(), Uuid: fixture.policies[0].GetUuid()},
				{To: fixture.policies[1].GetFQName(), Uuid: fixture.policies[1].GetUuid()},
				{To: fixture.policies[2].GetFQName(), Uuid: fixture.policies[2].GetUuid()},
			},
		})
	fixture.updater.Expect(network, "network_policy_refs", jsonData)

	network.DeleteNetworkPolicy(fixture.policies[0].GetUuid())

	refs, err := network.GetNetworkPolicyRefs()
	require.NoError(t, err)
	assert.Len(t, refs, 2)

	objJson, err := network.UpdateObject()
	require.NoError(t, err)
	updateDoesNotInclude(t, objJson, "network_policy_refs")

	err = network.UpdateReferences()
	require.NoError(t, err)
	assert.Len(t, fixture.updater.updates, 1)
	if len(fixture.updater.updates) > 0 {
		update := fixture.updater.updates[0]
		assert.Equal(t, "DELETE", update.Operation)
		assert.Equal(t, fixture.policies[0].GetUuid(), update.RefUuid)
	}

}

func TestUpdateRemoveLastReference(t *testing.T) {
	fixture := new(testUpdateFixture)
	fixture.SetUp(t)

	network := fixture.network
	jsonData := BuildNetworkPolicyRefsResponse(
		t, &PolicyRefsData{
			network.GetName(),
			network.GetUuid(),
			[]contrail.Reference{
				{To: fixture.policies[0].GetFQName(), Uuid: fixture.policies[0].GetUuid()},
			},
		})
	fixture.updater.Expect(network, "network_policy_refs", jsonData)

	network.DeleteNetworkPolicy(fixture.policies[0].GetUuid())

	refs, err := network.GetNetworkPolicyRefs()
	require.NoError(t, err)
	assert.Len(t, refs, 0)

	objJson, err := network.UpdateObject()
	var m map[string]*json.RawMessage
	err = json.Unmarshal(objJson, &m)
	require.NoError(t, err)
	v, present := m["network_policy_refs"]
	assert.True(t, present)
	assert.Equal(t, "[]", string(*v))

	err = network.UpdateReferences()
	require.NoError(t, err)
	assert.Len(t, fixture.updater.updates, 0)
}

func TestUpdateClearAndAdd(t *testing.T) {
	fixture := new(testUpdateFixture)
	fixture.SetUp(t)

	network := fixture.network

	network.ClearNetworkPolicy()
	network.AddNetworkPolicy(fixture.policies[0], types.VirtualNetworkPolicyType{
		Sequence: &types.SequenceType{10, 0}})

	objJson, err := network.UpdateObject()
	require.NoError(t, err, "UpdateObject")
	objJson, err = jsonMessageAdd(objJson, "name", network.GetName())
	require.NoError(t, err)

	putObject := new(types.VirtualNetwork)
	err = json.Unmarshal(objJson, putObject)
	require.NoError(t, err, "Decode %s", string(objJson))

	refs, err := putObject.GetNetworkPolicyRefs()
	assert.Len(t, refs, 1)

	err = network.UpdateReferences()
	require.NoError(t, err)
	assert.Len(t, fixture.updater.updates, 0)
}
