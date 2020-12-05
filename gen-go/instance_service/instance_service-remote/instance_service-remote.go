// Autogenerated by Thrift Compiler (0.9.3)
// DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING

package main

import (
	"flag"
	"fmt"
	"git.apache.org/thrift.git/lib/go/thrift"
	"instance_service"
	"math"
	"net"
	"net/url"
	"os"
	"strconv"
	"strings"
)

func Usage() {
	fmt.Fprintln(os.Stderr, "Usage of ", os.Args[0], " [-h host:port] [-u url] [-f[ramed]] function [arg1 [arg2...]]:")
	flag.PrintDefaults()
	fmt.Fprintln(os.Stderr, "\nFunctions:")
	fmt.Fprintln(os.Stderr, "  bool AddPort(PortList port_list)")
	fmt.Fprintln(os.Stderr, "  bool KeepAliveCheck()")
	fmt.Fprintln(os.Stderr, "  bool Connect()")
	fmt.Fprintln(os.Stderr, "  bool DeletePort(tuuid port_id)")
	fmt.Fprintln(os.Stderr, "  bool AddVirtualGateway(VirtualGatewayRequestList vgw_list)")
	fmt.Fprintln(os.Stderr, "  bool DeleteVirtualGateway( vgw_list)")
	fmt.Fprintln(os.Stderr, "  bool ConnectForVirtualGateway()")
	fmt.Fprintln(os.Stderr, "  bool AuditTimerForVirtualGateway(i32 timeout)")
	fmt.Fprintln(os.Stderr, "  bool TunnelNHEntryAdd(string src_ip, string dst_ip, string vrf_name)")
	fmt.Fprintln(os.Stderr, "  bool TunnelNHEntryDelete(string src_ip, string dst_ip, string vrf_name)")
	fmt.Fprintln(os.Stderr, "  bool RouteEntryAdd(string ip_address, string gw_ip, string vrf_name, string label)")
	fmt.Fprintln(os.Stderr, "  bool RouteEntryDelete(string ip_address, string vrf_name)")
	fmt.Fprintln(os.Stderr, "  bool AddHostRoute(string ip_address, string vrf_name)")
	fmt.Fprintln(os.Stderr, "  bool AddLocalVmRoute(string ip_address, string intf_uuid, string vrf_name, string label)")
	fmt.Fprintln(os.Stderr, "  bool AddRemoteVmRoute(string ip_address, string gw_ip, string vrf_name, string label)")
	fmt.Fprintln(os.Stderr, "  bool CreateVrf(string vrf_name)")
	fmt.Fprintln(os.Stderr)
	os.Exit(0)
}

func main() {
	flag.Usage = Usage
	var host string
	var port int
	var protocol string
	var urlString string
	var framed bool
	var useHttp bool
	var parsedUrl url.URL
	var trans thrift.TTransport
	_ = strconv.Atoi
	_ = math.Abs
	flag.Usage = Usage
	flag.StringVar(&host, "h", "localhost", "Specify host and port")
	flag.IntVar(&port, "p", 9090, "Specify port")
	flag.StringVar(&protocol, "P", "binary", "Specify the protocol (binary, compact, simplejson, json)")
	flag.StringVar(&urlString, "u", "", "Specify the url")
	flag.BoolVar(&framed, "framed", false, "Use framed transport")
	flag.BoolVar(&useHttp, "http", false, "Use http")
	flag.Parse()

	if len(urlString) > 0 {
		parsedUrl, err := url.Parse(urlString)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error parsing URL: ", err)
			flag.Usage()
		}
		host = parsedUrl.Host
		useHttp = len(parsedUrl.Scheme) <= 0 || parsedUrl.Scheme == "http"
	} else if useHttp {
		_, err := url.Parse(fmt.Sprint("http://", host, ":", port))
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error parsing URL: ", err)
			flag.Usage()
		}
	}

	cmd := flag.Arg(0)
	var err error
	if useHttp {
		trans, err = thrift.NewTHttpClient(parsedUrl.String())
	} else {
		portStr := fmt.Sprint(port)
		if strings.Contains(host, ":") {
			host, portStr, err = net.SplitHostPort(host)
			if err != nil {
				fmt.Fprintln(os.Stderr, "error with host:", err)
				os.Exit(1)
			}
		}
		trans, err = thrift.NewTSocket(net.JoinHostPort(host, portStr))
		if err != nil {
			fmt.Fprintln(os.Stderr, "error resolving address:", err)
			os.Exit(1)
		}
		if framed {
			trans = thrift.NewTFramedTransport(trans)
		}
	}
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error creating transport", err)
		os.Exit(1)
	}
	defer trans.Close()
	var protocolFactory thrift.TProtocolFactory
	switch protocol {
	case "compact":
		protocolFactory = thrift.NewTCompactProtocolFactory()
		break
	case "simplejson":
		protocolFactory = thrift.NewTSimpleJSONProtocolFactory()
		break
	case "json":
		protocolFactory = thrift.NewTJSONProtocolFactory()
		break
	case "binary", "":
		protocolFactory = thrift.NewTBinaryProtocolFactoryDefault()
		break
	default:
		fmt.Fprintln(os.Stderr, "Invalid protocol specified: ", protocol)
		Usage()
		os.Exit(1)
	}
	client := instance_service.NewInstanceServiceClientFactory(trans, protocolFactory)
	if err := trans.Open(); err != nil {
		fmt.Fprintln(os.Stderr, "Error opening socket to ", host, ":", port, " ", err)
		os.Exit(1)
	}

	switch cmd {
	case "AddPort":
		if flag.NArg()-1 != 1 {
			fmt.Fprintln(os.Stderr, "AddPort requires 1 args")
			flag.Usage()
		}
		arg44 := flag.Arg(1)
		mbTrans45 := thrift.NewTMemoryBufferLen(len(arg44))
		defer mbTrans45.Close()
		_, err46 := mbTrans45.WriteString(arg44)
		if err46 != nil {
			Usage()
			return
		}
		factory47 := thrift.NewTSimpleJSONProtocolFactory()
		jsProt48 := factory47.GetProtocol(mbTrans45)
		containerStruct0 := instance_service.NewInstanceServiceAddPortArgs()
		err49 := containerStruct0.ReadField1(jsProt48)
		if err49 != nil {
			Usage()
			return
		}
		argvalue0 := containerStruct0.PortList
		value0 := instance_service.PortList(argvalue0)
		fmt.Print(client.AddPort(value0))
		fmt.Print("\n")
		break
	case "KeepAliveCheck":
		if flag.NArg()-1 != 0 {
			fmt.Fprintln(os.Stderr, "KeepAliveCheck requires 0 args")
			flag.Usage()
		}
		fmt.Print(client.KeepAliveCheck())
		fmt.Print("\n")
		break
	case "Connect":
		if flag.NArg()-1 != 0 {
			fmt.Fprintln(os.Stderr, "Connect requires 0 args")
			flag.Usage()
		}
		fmt.Print(client.Connect())
		fmt.Print("\n")
		break
	case "DeletePort":
		if flag.NArg()-1 != 1 {
			fmt.Fprintln(os.Stderr, "DeletePort requires 1 args")
			flag.Usage()
		}
		arg50 := flag.Arg(1)
		mbTrans51 := thrift.NewTMemoryBufferLen(len(arg50))
		defer mbTrans51.Close()
		_, err52 := mbTrans51.WriteString(arg50)
		if err52 != nil {
			Usage()
			return
		}
		factory53 := thrift.NewTSimpleJSONProtocolFactory()
		jsProt54 := factory53.GetProtocol(mbTrans51)
		containerStruct0 := instance_service.NewInstanceServiceDeletePortArgs()
		err55 := containerStruct0.ReadField1(jsProt54)
		if err55 != nil {
			Usage()
			return
		}
		argvalue0 := containerStruct0.PortID
		value0 := instance_service.Tuuid(argvalue0)
		fmt.Print(client.DeletePort(value0))
		fmt.Print("\n")
		break
	case "AddVirtualGateway":
		if flag.NArg()-1 != 1 {
			fmt.Fprintln(os.Stderr, "AddVirtualGateway requires 1 args")
			flag.Usage()
		}
		arg56 := flag.Arg(1)
		mbTrans57 := thrift.NewTMemoryBufferLen(len(arg56))
		defer mbTrans57.Close()
		_, err58 := mbTrans57.WriteString(arg56)
		if err58 != nil {
			Usage()
			return
		}
		factory59 := thrift.NewTSimpleJSONProtocolFactory()
		jsProt60 := factory59.GetProtocol(mbTrans57)
		containerStruct0 := instance_service.NewInstanceServiceAddVirtualGatewayArgs()
		err61 := containerStruct0.ReadField1(jsProt60)
		if err61 != nil {
			Usage()
			return
		}
		argvalue0 := containerStruct0.VgwList
		value0 := instance_service.VirtualGatewayRequestList(argvalue0)
		fmt.Print(client.AddVirtualGateway(value0))
		fmt.Print("\n")
		break
	case "DeleteVirtualGateway":
		if flag.NArg()-1 != 1 {
			fmt.Fprintln(os.Stderr, "DeleteVirtualGateway requires 1 args")
			flag.Usage()
		}
		arg62 := flag.Arg(1)
		mbTrans63 := thrift.NewTMemoryBufferLen(len(arg62))
		defer mbTrans63.Close()
		_, err64 := mbTrans63.WriteString(arg62)
		if err64 != nil {
			Usage()
			return
		}
		factory65 := thrift.NewTSimpleJSONProtocolFactory()
		jsProt66 := factory65.GetProtocol(mbTrans63)
		containerStruct0 := instance_service.NewInstanceServiceDeleteVirtualGatewayArgs()
		err67 := containerStruct0.ReadField1(jsProt66)
		if err67 != nil {
			Usage()
			return
		}
		argvalue0 := containerStruct0.VgwList
		value0 := argvalue0
		fmt.Print(client.DeleteVirtualGateway(value0))
		fmt.Print("\n")
		break
	case "ConnectForVirtualGateway":
		if flag.NArg()-1 != 0 {
			fmt.Fprintln(os.Stderr, "ConnectForVirtualGateway requires 0 args")
			flag.Usage()
		}
		fmt.Print(client.ConnectForVirtualGateway())
		fmt.Print("\n")
		break
	case "AuditTimerForVirtualGateway":
		if flag.NArg()-1 != 1 {
			fmt.Fprintln(os.Stderr, "AuditTimerForVirtualGateway requires 1 args")
			flag.Usage()
		}
		tmp0, err68 := (strconv.Atoi(flag.Arg(1)))
		if err68 != nil {
			Usage()
			return
		}
		argvalue0 := int32(tmp0)
		value0 := argvalue0
		fmt.Print(client.AuditTimerForVirtualGateway(value0))
		fmt.Print("\n")
		break
	case "TunnelNHEntryAdd":
		if flag.NArg()-1 != 3 {
			fmt.Fprintln(os.Stderr, "TunnelNHEntryAdd requires 3 args")
			flag.Usage()
		}
		argvalue0 := flag.Arg(1)
		value0 := argvalue0
		argvalue1 := flag.Arg(2)
		value1 := argvalue1
		argvalue2 := flag.Arg(3)
		value2 := argvalue2
		fmt.Print(client.TunnelNHEntryAdd(value0, value1, value2))
		fmt.Print("\n")
		break
	case "TunnelNHEntryDelete":
		if flag.NArg()-1 != 3 {
			fmt.Fprintln(os.Stderr, "TunnelNHEntryDelete requires 3 args")
			flag.Usage()
		}
		argvalue0 := flag.Arg(1)
		value0 := argvalue0
		argvalue1 := flag.Arg(2)
		value1 := argvalue1
		argvalue2 := flag.Arg(3)
		value2 := argvalue2
		fmt.Print(client.TunnelNHEntryDelete(value0, value1, value2))
		fmt.Print("\n")
		break
	case "RouteEntryAdd":
		if flag.NArg()-1 != 4 {
			fmt.Fprintln(os.Stderr, "RouteEntryAdd requires 4 args")
			flag.Usage()
		}
		argvalue0 := flag.Arg(1)
		value0 := argvalue0
		argvalue1 := flag.Arg(2)
		value1 := argvalue1
		argvalue2 := flag.Arg(3)
		value2 := argvalue2
		argvalue3 := flag.Arg(4)
		value3 := argvalue3
		fmt.Print(client.RouteEntryAdd(value0, value1, value2, value3))
		fmt.Print("\n")
		break
	case "RouteEntryDelete":
		if flag.NArg()-1 != 2 {
			fmt.Fprintln(os.Stderr, "RouteEntryDelete requires 2 args")
			flag.Usage()
		}
		argvalue0 := flag.Arg(1)
		value0 := argvalue0
		argvalue1 := flag.Arg(2)
		value1 := argvalue1
		fmt.Print(client.RouteEntryDelete(value0, value1))
		fmt.Print("\n")
		break
	case "AddHostRoute":
		if flag.NArg()-1 != 2 {
			fmt.Fprintln(os.Stderr, "AddHostRoute requires 2 args")
			flag.Usage()
		}
		argvalue0 := flag.Arg(1)
		value0 := argvalue0
		argvalue1 := flag.Arg(2)
		value1 := argvalue1
		fmt.Print(client.AddHostRoute(value0, value1))
		fmt.Print("\n")
		break
	case "AddLocalVmRoute":
		if flag.NArg()-1 != 4 {
			fmt.Fprintln(os.Stderr, "AddLocalVmRoute requires 4 args")
			flag.Usage()
		}
		argvalue0 := flag.Arg(1)
		value0 := argvalue0
		argvalue1 := flag.Arg(2)
		value1 := argvalue1
		argvalue2 := flag.Arg(3)
		value2 := argvalue2
		argvalue3 := flag.Arg(4)
		value3 := argvalue3
		fmt.Print(client.AddLocalVmRoute(value0, value1, value2, value3))
		fmt.Print("\n")
		break
	case "AddRemoteVmRoute":
		if flag.NArg()-1 != 4 {
			fmt.Fprintln(os.Stderr, "AddRemoteVmRoute requires 4 args")
			flag.Usage()
		}
		argvalue0 := flag.Arg(1)
		value0 := argvalue0
		argvalue1 := flag.Arg(2)
		value1 := argvalue1
		argvalue2 := flag.Arg(3)
		value2 := argvalue2
		argvalue3 := flag.Arg(4)
		value3 := argvalue3
		fmt.Print(client.AddRemoteVmRoute(value0, value1, value2, value3))
		fmt.Print("\n")
		break
	case "CreateVrf":
		if flag.NArg()-1 != 1 {
			fmt.Fprintln(os.Stderr, "CreateVrf requires 1 args")
			flag.Usage()
		}
		argvalue0 := flag.Arg(1)
		value0 := argvalue0
		fmt.Print(client.CreateVrf(value0))
		fmt.Print("\n")
		break
	case "":
		Usage()
		break
	default:
		fmt.Fprintln(os.Stderr, "Invalid function ", cmd)
	}
}
