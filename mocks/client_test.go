package mocks

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/Juniper/contrail-go-api/types"
)

func TestUpdateClearVM(t *testing.T) {
	client := new(ApiClient)
	client.Init()

	project := new(types.Project)
	project.SetFQName("domain", []string{"default-domain", "tenant"})
	assert.NoError(t, client.Create(project))

	vm := new(types.VirtualMachine)
	vm.SetFQName("project", []string{"default-domain", "tenant", "instance"})
	assert.NoError(t, client.Create(vm))

	vmi := new(types.VirtualMachineInterface)
	vmi.SetFQName("project", []string{"default-domain", "tenant", "instance"})
	vmi.AddVirtualMachine(vm)
	assert.NoError(t, client.Create(vmi))

	vmi.ClearVirtualMachine()
	assert.NoError(t, client.Update(vmi))

	assert.NoError(t, client.Delete(vm))
}

func TestListByParent(t *testing.T) {
	client := new(ApiClient)
	client.Init()

	projectNames := []string{"p1", "p2", "p3"}
	projects := make([]*types.Project, 3)
	vmNames := []string{"a", "b", "c", "d"}

	for i, projectName := range projectNames {
		project := new(types.Project)
		project.SetFQName("domain", []string{"default-domain", projectName})
		assert.NoError(t, client.Create(project))
		projects[i] = project

		for _, vmName := range vmNames {
			vm := new(types.VirtualMachine)
			vm.SetFQName("project", []string{"default-domain", projectName, vmName})
			assert.NoError(t, client.Create(vm))
		}
	}

	elements, err := client.ListByParent("virtual-machine", projects[1].GetUuid())
	assert.NoError(t, err)
	for _, element := range elements {
		assert.Equal(t, "p2", element.Fq_name[1])
	}
}

func TestListAny(t *testing.T) {
	client := new(ApiClient)
	client.Init()

	projectNames := []string{"p1", "p2", "p3"}
	projects := make([]*types.Project, 3)
	vmNames := []string{"a", "b", "c", "d"}

	for i, projectName := range projectNames {
		project := new(types.Project)
		project.SetFQName("domain", []string{"default-domain", projectName})
		assert.NoError(t, client.Create(project))
		projects[i] = project

		for _, vmName := range vmNames {
			vm := new(types.VirtualMachine)
			vm.SetFQName("project", []string{"default-domain", projectName, vmName})
			assert.NoError(t, client.Create(vm))
		}
	}

	elements, err := client.List("virtual-machine")
	require.NoError(t, err)
	assert.Len(t, elements, len(projectNames)*len(vmNames))
}
