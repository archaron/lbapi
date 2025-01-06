package types

type AgentType int

const (
	AgentTypeEthernetPCAP  = 1
	AgentTypeEthernetULOG  = 2
	AgentTypeEthernetTEE   = 3
	AgentTypeNetflow       = 4
	AgentTypeSFlow         = 5
	AgentTypeRadius        = 6
	AgentTypeLBPhone       = 7
	AgentTypePABXRS232     = 8
	AgentTypePABXFIFO      = 9
	AgentTypePABXTCPClient = 10
	AgentTypePABXTCPServer = 11
	AgentTypeVOIP          = 12
	AgentTypeServices      = 13
	AgentTypeSNMP          = 14
)

var AgentTypesInternet = []AgentType{
	AgentTypeEthernetPCAP,
	AgentTypeEthernetULOG,
	AgentTypeEthernetTEE,
	AgentTypeNetflow,
	AgentTypeSFlow,
	AgentTypeRadius,
}

var _ = AgentTypesInternet

func (a AgentType) String() string {
	switch a {
	case AgentTypeEthernetPCAP:
		return "Ethernet/PCAP"
	case AgentTypeEthernetULOG:
		return "Ethernet/ULOG"
	case AgentTypeEthernetTEE:
		return "Ethernet/tee"
	case AgentTypeNetflow:
		return "Netflow"
	case AgentTypeSFlow:
		return "SFlow"
	case AgentTypeRadius:
		return "RADIUS/DialUP"
	case AgentTypeLBPhone:
		return "PCDR"
	case AgentTypePABXRS232:
		return "PABX/RS232"
	case AgentTypePABXFIFO:
		return "PABX/FIFO"
	case AgentTypePABXTCPClient:
		return "PABX/TCPclient"
	case AgentTypePABXTCPServer:
		return "PABX/TCPserver"
	case AgentTypeVOIP:
		return "RADIUS/VoIP"
	case AgentTypeServices:
		return "USBox"
	case AgentTypeSNMP:
		return "SNMP"
	default:
		return "Unknown"
	}
}
