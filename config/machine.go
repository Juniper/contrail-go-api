package config

import (
	"github.com/Juniper/contrail-go-api"
    "github.com/Juniper/contrail-go-api/types"
)

type MachineInfo struct {
	DisplayName         string
    Uuid                string
    InstanceIp          string
    FloatingIp          string
}

func buildMachineInfo(client contrail.ApiClient, machine *types.VirtualMachineInterface, detail bool) (*MachineInfo, error) {
    var machineDisplayName   string
    var machineUuid          string
    var machineInstanceIp    string
    var machineFloatingIp    string

    virtualMachineRefs, err := machine.GetVirtualMachineRefs()
    if err != nil {
        return nil, err
    }

    virtualMachine, err := client.FindByUuid("virtual-machine", virtualMachineRefs[0].Uuid)
    if err != nil {
        return nil, err
    }

    machineDisplayName = virtualMachine.(*types.VirtualMachine).GetDisplayName()
    machineUuid = virtualMachine.(*types.VirtualMachine).GetUuid()

    //get instance ip
    instanceIpBackRefs, err := machine.GetInstanceIpBackRefs()
    if err != nil {
        return nil, err
    }
    instanceIp, err := client.FindByUuid("instance-ip", instanceIpBackRefs[0].Uuid)
    if err != nil { 
        return nil, err
    }   
    machineInstanceIp = instanceIp.(*types.InstanceIp).GetInstanceIpAddress()    

    //get floating ip
    floatingIpBackRefs, err := machine.GetFloatingIpBackRefs()
    if err != nil {
        return nil, err
    }
    if len(floatingIpBackRefs) == 1 {
        floatingIp, err := client.FindByUuid("floating-ip", floatingIpBackRefs[0].Uuid)
        if err != nil {
            return nil, err
        }
        machineFloatingIp = floatingIp.(*types.FloatingIp).GetFloatingIpAddress()
    }

    info := &MachineInfo{
        machineDisplayName,
        machineUuid,
        machineInstanceIp,
        machineFloatingIp,
    }

	return info, nil
}

func MachineShow(client contrail.ApiClient, uuid string, detail bool) (*MachineInfo, error) {
	machine, err := client.FindByUuid("virtual-machine", uuid)
	if err != nil {
		return nil, err
	}

    virtualMachineInterfaceRefs, err := machine.(*types.VirtualMachine).GetVirtualMachineInterfaceBackRefs()
    if err != nil {
        return nil, err
    }

    virtualMachineInterface, err := client.FindByUuid("virtual-machine-interface", virtualMachineInterfaceRefs[0].Uuid)
    if err != nil {
        return nil, err
    }

	return buildMachineInfo(client, virtualMachineInterface.(*types.VirtualMachineInterface), detail)   
}

func MachineList(client contrail.ApiClient, project_id string, detail bool) ([]*MachineInfo, error) {
    var machineList []*MachineInfo
    var fields []string

    machineInterfaces, err := client.ListDetailByParent("virtual-machine-interface", project_id, fields)
    if err != nil {
        return nil, err
    }

    for _, reference := range machineInterfaces {
        info, err := buildMachineInfo(client, reference.(*types.VirtualMachineInterface), detail)
        if err != nil {
            return nil, err
        }
        machineList = append(machineList, info)
    }

    return machineList, nil 
}


