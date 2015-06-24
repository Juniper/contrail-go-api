package mocks

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/Juniper/contrail-go-api/types"
)

func TestReadBackRefs(t *testing.T) {
	client := new(ApiClient)
	client.Init()

	project := new(types.Project)
	project.SetFQName("domain", []string{"default-domain", "p1"})
	assert.NoError(t, client.Create(project))

	net := new(types.VirtualNetwork)
	net.SetFQName("project", []string{"default-domain", "p1", "n1"})
	props := new(types.VirtualNetworkType)
	props.ForwardingMode = "turbo"
	net.SetVirtualNetworkProperties(props)
	assert.NoError(t, client.Create(net))

	vmi1 := new(types.VirtualMachineInterface)
	vmi1.SetFQName("project", []string{"default-domain", "p1", "port1"})
	vmi1.AddVirtualNetwork(net)
	assert.NoError(t, client.Create(vmi1))

	vmi2 := new(types.VirtualMachineInterface)
	vmi2.SetFQName("project", []string{"default-domain", "p1", "port2"})
	vmi2.AddVirtualNetwork(net)
	assert.NoError(t, client.Create(vmi2))

	refs, err := net.GetVirtualMachineInterfaceBackRefs()
	assert.NoError(t, err)
	assert.Len(t, refs, 2)
	idList := make([]string, len(refs))
	for i, ref := range refs {
		idList[i] = ref.Uuid
	}

	assert.Contains(t, idList, vmi1.GetUuid())
	assert.Contains(t, idList, vmi2.GetUuid())

	rProps := net.GetVirtualNetworkProperties()
	assert.Equal(t, "turbo", rProps.ForwardingMode)
}

func TestReadChildren(t *testing.T) {
	client := new(ApiClient)
	client.Init()

	project := new(types.Project)
	project.SetFQName("domain", []string{"default-domain", "p1"})
	assert.NoError(t, client.Create(project))

	net := new(types.VirtualNetwork)
	net.SetFQName("project", []string{"default-domain", "p1", "n1"})
	assert.NoError(t, client.Create(net))

	vmi1 := new(types.VirtualMachineInterface)
	vmi1.SetFQName("project", []string{"default-domain", "p1", "port1"})
	vmi1.AddVirtualNetwork(net)
	assert.NoError(t, client.Create(vmi1))

	vmi2 := new(types.VirtualMachineInterface)
	vmi2.SetFQName("project", []string{"default-domain", "p1", "port2"})
	vmi2.AddVirtualNetwork(net)
	assert.NoError(t, client.Create(vmi2))

	refs, err := project.GetVirtualMachineInterfaces()
	assert.NoError(t, err)
	assert.Len(t, refs, 2)
}

func TestReadRefs(t *testing.T) {
	client := new(ApiClient)
	client.Init()

	project := new(types.Project)
	project.SetFQName("domain", []string{"default-domain", "p1"})
	assert.NoError(t, client.Create(project))

	net := new(types.VirtualNetwork)
	net.SetFQName("project", []string{"default-domain", "p1", "n1"})
	assert.NoError(t, client.Create(net))

	vmi1 := new(types.VirtualMachineInterface)
	vmi1.SetFQName("project", []string{"default-domain", "p1", "port1"})
	vmi1.AddVirtualNetwork(net)
	assert.NoError(t, client.Create(vmi1))

	refs, err := vmi1.GetVirtualNetworkRefs()
	assert.NoError(t, err)
	assert.Len(t, refs, 1)
	assert.Equal(t, net.GetUuid(), refs[0].Uuid)
}

func TestReadModifiedRefs(t *testing.T) {
	client := new(ApiClient)
	client.Init()

	project := new(types.Project)
	project.SetFQName("domain", []string{"default-domain", "p1"})
	assert.NoError(t, client.Create(project))

	net := new(types.VirtualNetwork)
	net.SetFQName("project", []string{"default-domain", "p1", "n1"})
	assert.NoError(t, client.Create(net))

	vmi1 := new(types.VirtualMachineInterface)
	vmi1.SetFQName("project", []string{"default-domain", "p1", "port1"})
	vmi1.AddVirtualNetwork(net)
	assert.NoError(t, client.Create(vmi1))

	vmi2 := new(types.VirtualMachineInterface)
	vmi2.SetFQName("project", []string{"default-domain", "p1", "port2"})
	vmi2.AddVirtualNetwork(net)
	assert.NoError(t, client.Create(vmi2))

	refs, err := net.GetVirtualMachineInterfaceBackRefs()
	assert.NoError(t, err)
	assert.Len(t, refs, 2)

	assert.NoError(t, client.Delete(vmi1))
	refs, err = net.GetVirtualMachineInterfaceBackRefs()
	assert.NoError(t, err)
	assert.Len(t, refs, 1)
}
