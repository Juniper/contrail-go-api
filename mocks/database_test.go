package mocks

import (
	"fmt"
	"strings"
	"testing"

	"github.com/pborman/uuid"
	"github.com/stretchr/testify/assert"

	"github.com/Juniper/contrail-go-api/types"
)

func TestReference(t *testing.T) {
	db := NewInMemDatabase()

	ipam := new(types.NetworkIpam)
	ipam.SetUuid(uuid.New())
	ipam.SetFQName("project", strings.Split("default-domain:p1:ipam", ":"))
	assert.NoError(t, db.Put(ipam, nil, UIDList{}))

	net := new(types.VirtualNetwork)
	net.SetUuid(uuid.New())
	net.SetFQName("project", strings.Split("default-domain:p1:network", ":"))
	subnets := types.VnSubnetsType{}
	subnets.AddIpamSubnets(
		&types.IpamSubnetType{
			Subnet: &types.SubnetType{"10.0.0.0", 8}})

	net.AddNetworkIpam(ipam, subnets)
	refs := GetReferenceList(net)
	assert.NoError(t, db.Put(net, nil, refs))

	result, err := db.GetBackReferences(parseUID(ipam.GetUuid()), "virtual_network")
	assert.NoError(t, err)
	assert.Contains(t, result, parseUID(net.GetUuid()))
}

func TestUpdateRefs(t *testing.T) {
	db := NewInMemDatabase()

	var instances [8]*types.VirtualMachine
	for i := 0; i < 8; i++ {
		instance := new(types.VirtualMachine)
		instance.SetUuid(uuid.New())
		instance.SetName(fmt.Sprintf("instance-%d", i))
		assert.NoError(t, db.Put(instance, nil, UIDList{}))
		instances[i] = instance
	}

	vmi1 := new(types.VirtualMachineInterface)
	vmi1.SetUuid(uuid.New())
	vmi1.SetName("port1")
	vmi1.AddVirtualMachine(instances[0])
	vmi1.AddVirtualMachine(instances[1])
	vmi1.AddVirtualMachine(instances[2])
	assert.NoError(t, db.Put(vmi1, nil, GetReferenceList(vmi1)))

	vmi2 := new(types.VirtualMachineInterface)
	vmi2.SetUuid(uuid.New())
	vmi2.SetName("port2")
	vmi2.AddVirtualMachine(instances[2])
	vmi2.AddVirtualMachine(instances[4])
	vmi2.AddVirtualMachine(instances[3])
	assert.NoError(t, db.Put(vmi2, nil, GetReferenceList(vmi2)))

	r2, err := db.GetBackReferences(parseUID(instances[2].GetUuid()), "virtual_machine_interface")
	assert.NoError(t, err)
	assert.Len(t, r2, 2)
	assert.Contains(t, r2, parseUID(vmi1.GetUuid()))
	assert.Contains(t, r2, parseUID(vmi2.GetUuid()))

	r3, err := db.GetBackReferences(parseUID(instances[3].GetUuid()), "virtual_machine_interface")
	assert.NoError(t, err)
	assert.Len(t, r3, 1)

	vmi2.ClearVirtualMachine()
	vmi2.AddVirtualMachine(instances[4])
	vmi2.AddVirtualMachine(instances[6])
	assert.NoError(t, db.Update(vmi2, GetReferenceList(vmi2)))

	r2, err = db.GetBackReferences(parseUID(instances[2].GetUuid()), "virtual_machine_interface")
	assert.NoError(t, err)
	assert.Len(t, r2, 1)
	assert.Contains(t, r2, parseUID(vmi1.GetUuid()))

	r3, err = db.GetBackReferences(parseUID(instances[3].GetUuid()), "virtual_machine_interface")
	assert.NoError(t, err)
	assert.Len(t, r3, 0)

	r4, err := db.GetBackReferences(parseUID(instances[4].GetUuid()), "virtual_machine_interface")
	assert.NoError(t, err)
	assert.Len(t, r4, 1)
	assert.Contains(t, r4, parseUID(vmi2.GetUuid()))

	vmi1.ClearVirtualMachine()
	vmi1.AddVirtualMachine(instances[0])
	vmi1.AddVirtualMachine(instances[1])
	vmi1.AddVirtualMachine(instances[5])
	vmi1.AddVirtualMachine(instances[4])
	assert.NoError(t, db.Update(vmi1, GetReferenceList(vmi1)))

	r4, err = db.GetBackReferences(parseUID(instances[4].GetUuid()), "virtual_machine_interface")
	assert.NoError(t, err)
	assert.Len(t, r4, 2)
	assert.Contains(t, r4, parseUID(vmi1.GetUuid()))
	assert.Contains(t, r4, parseUID(vmi2.GetUuid()))
}

func TestUpdateOne(t *testing.T) {
	db := NewInMemDatabase()

	instance := new(types.VirtualMachine)
	instance.SetUuid(uuid.New())
	instance.SetName("instance")
	assert.NoError(t, db.Put(instance, nil, UIDList{}))

	vmi1 := new(types.VirtualMachineInterface)
	vmi1.SetUuid(uuid.New())
	vmi1.SetName("port1")
	assert.NoError(t, db.Put(vmi1, nil, GetReferenceList(vmi1)))

	vmi1.AddVirtualMachine(instance)
	assert.NoError(t, db.Update(vmi1, GetReferenceList(vmi1)))

	result, err := db.GetBackReferences(parseUID(instance.GetUuid()), "virtual_machine_interface")
	assert.NoError(t, err)
	assert.Contains(t, result, parseUID(vmi1.GetUuid()))

	vmi1.ClearVirtualMachine()
	assert.NoError(t, db.Update(vmi1, GetReferenceList(vmi1)))
	result, err = db.GetBackReferences(parseUID(instance.GetUuid()), "virtual_machine_interface")
	assert.NoError(t, err)
	assert.Len(t, result, 0)
}

// Create floating-ip with 2 vmi references, delete it and verify that
// the back_refs are updated as expected.
func TestDeleteRefs(t *testing.T) {
	db := NewInMemDatabase()

	vmi1 := new(types.VirtualMachineInterface)
	vmi1.SetUuid(uuid.New())
	vmi1.SetName("port1")
	assert.NoError(t, db.Put(vmi1, nil, GetReferenceList(vmi1)))

	vmi2 := new(types.VirtualMachineInterface)
	vmi2.SetUuid(uuid.New())
	vmi2.SetName("port2")
	assert.NoError(t, db.Put(vmi2, nil, GetReferenceList(vmi2)))

	fip := new(types.FloatingIp)
	fip.SetUuid(uuid.New())
	fip.SetName("fip")
	fip.AddVirtualMachineInterface(vmi1)
	fip.AddVirtualMachineInterface(vmi2)
	assert.NoError(t, db.Put(fip, nil, GetReferenceList(fip)))

	assert.Error(t, db.Delete(vmi1))

	result, err := db.GetBackReferences(parseUID(vmi2.GetUuid()), "floating_ip")
	assert.NoError(t, err)
	assert.Len(t, result, 1)

	assert.NoError(t, db.Delete(fip))

	result, err = db.GetBackReferences(parseUID(vmi2.GetUuid()), "floating_ip")
	assert.NoError(t, err)
	assert.Len(t, result, 0)

	assert.NoError(t, db.Delete(vmi1))
	assert.NoError(t, db.Delete(vmi2))
}

func TestChildrenRefs(t *testing.T) {
	db := NewInMemDatabase()
	project := new(types.Project)
	project.SetUuid(uuid.New())
	project.SetName("p1")
	assert.NoError(t, db.Put(project, nil, GetReferenceList(project)))

	net := new(types.VirtualNetwork)
	net.SetUuid(uuid.New())
	net.SetFQName("project", []string{"p1", "n1"})
	assert.NoError(t, db.Put(net, project, UIDList{}))

	vmi1 := new(types.VirtualMachineInterface)
	vmi1.SetUuid(uuid.New())
	vmi1.SetName("port1")
	vmi1.AddVirtualNetwork(net)
	assert.NoError(t, db.Put(vmi1, project, GetReferenceList(vmi1)))

	vmi2 := new(types.VirtualMachineInterface)
	vmi2.SetUuid(uuid.New())
	vmi2.SetName("port2")
	vmi2.AddVirtualNetwork(net)
	assert.NoError(t, db.Put(vmi2, project, GetReferenceList(vmi2)))

	result, err := db.GetChildren(parseUID(project.GetUuid()), "virtual_network")
	assert.NoError(t, err)
	assert.Len(t, result, 1)

	result, err = db.GetChildren(parseUID(project.GetUuid()), "virtual_machine_interface")
	assert.NoError(t, err)
	assert.Len(t, result, 2)
	assert.Contains(t, result, parseUID(vmi1.GetUuid()))
	assert.Contains(t, result, parseUID(vmi2.GetUuid()))

	assert.Error(t, db.Delete(project))

	assert.NoError(t, db.Delete(vmi1))
	assert.Error(t, db.Delete(net))
	assert.NoError(t, db.Delete(vmi2))

	assert.NoError(t, db.Delete(net))
	assert.NoError(t, db.Delete(project))
}
