package contrail

import (
    "fmt"
    "git.apache.org/thrift.git/lib/go/thrift"
    "./gen-go/instance_service"
    "encoding/hex"
    "strings"
    "strconv"
)

func rpc_client_instance() (error, *instance_service.InstanceServiceClient ) {
        addr := "localhost:9090"
	socket, err := thrift.NewTSocket(addr)
	if err != nil {
            fmt.Println("Error opening socket:", err)
        }
	transportFactory := thrift.NewTTransportFactory()
    	framedTransportFactory := thrift.NewTFramedTransportFactory(transportFactory)
	transport := framedTransportFactory.GetTransport(socket)
	transport.Open()
	protocolFactory := thrift.NewTBinaryProtocolFactoryDefault()
	client := instance_service.NewInstanceServiceClientFactory(transport, protocolFactory)
	client.Connect()
    	return nil, client	
}

func StringToTuuid(Uuid string) (instance_service.Tuuid) {
    uuidSlice := strings.Split(Uuid, "-")
    var intArr []int16
    var tArr instance_service.Tuuid
    for i := 0; i < len(uuidSlice); i++{
        b2, _ := hex.DecodeString(uuidSlice[i])
        for x := 0; x < len(b2); x++{
            tmpS := fmt.Sprintf("%v",b2[x])
            i, _ := strconv.Atoi(tmpS)
            i16 := int16(i)
            intArr = append(intArr,i16)
            tArr = append(tArr,i16)
        }
    }
    return tArr

}

func vrouterDelPort(vmiUuid string) {
        err, client := rpc_client_instance()
	if err != nil {
		fmt.Println("error: ", err)
	}
	vmiUuidT := StringToTuuid(vmiUuid)
	client.DeletePort(vmiUuidT)
}

func vrouterAddPort(vmiUuid string, vmUuid string, hostVethName string, mac string, vnName string, projectUuid string, portTypeString string) {
        err, client := rpc_client_instance()
	if err != nil {
		fmt.Println("error: ", err)
	}
	ipAddress := "0.0.0.0"
        var portList instance_service.PortList
        vmiUuidT := StringToTuuid(vmiUuid)
        vmUuidT := StringToTuuid(vmUuid)
	var vnIdT instance_service.Tuuid
	var portType int16
	portType = 0
	for i := 0; i < 16; i++{
		vnIdT = append(vnIdT, 0)
	} 
	port := &instance_service.Port{
		PortID : vmiUuidT,
		InstanceID : vmUuidT,
		TapName : hostVethName,
		IPAddress : ipAddress,
		VnID : vnIdT,
		MacAddress : mac,
		//VMProjectID : projectUuidT,
		//DisplayName : &vnName,
		PortType : &portType,
	}
	portList = append(portList, port)
	client.AddPort(portList)
}
