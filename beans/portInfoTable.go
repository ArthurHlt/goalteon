package beans

import (
	"fmt"
	"reflect"
)

// PortInfoTable The table of port information.
type PortInfoTable struct {
	// The port index.
	PortInfoIndx int32
	Params       *PortInfoTableParams
}

func NewPortInfoTableList() *PortInfoTable {
	return &PortInfoTable{}
}

func NewPortInfoTable(
	portInfoIndx int32,
	params *PortInfoTableParams,
) *PortInfoTable {
	return &PortInfoTable{
		PortInfoIndx: portInfoIndx,
		Params:       params,
	}
}

func (c *PortInfoTable) Name() string {
	return "PortInfoTable"
}

func (c *PortInfoTable) GetParams() BeanType {
	return c.Params
}

func (c *PortInfoTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *PortInfoTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.PortInfoIndx).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.PortInfoIndx)
}

type PortInfoTableSpeed int32

const (
	PortInfoTableSpeed_Mbs10       PortInfoTableSpeed = 2
	PortInfoTableSpeed_Mbs100      PortInfoTableSpeed = 3
	PortInfoTableSpeed_Mbs1000     PortInfoTableSpeed = 4
	PortInfoTableSpeed_Any         PortInfoTableSpeed = 5
	PortInfoTableSpeed_Mbs10000    PortInfoTableSpeed = 6
	PortInfoTableSpeed_Mbs40000    PortInfoTableSpeed = 7
	PortInfoTableSpeed_Auto        PortInfoTableSpeed = 8
	PortInfoTableSpeed_Mbs100000   PortInfoTableSpeed = 9
	PortInfoTableSpeed_Mbs25000    PortInfoTableSpeed = 10
	PortInfoTableSpeed_Unsupported PortInfoTableSpeed = 2147483647
)

type PortInfoTableMode int32

const (
	PortInfoTableMode_FullDuplex  PortInfoTableMode = 2
	PortInfoTableMode_HalfDuplex  PortInfoTableMode = 3
	PortInfoTableMode_Any         PortInfoTableMode = 4
	PortInfoTableMode_Unsupported PortInfoTableMode = 2147483647
)

type PortInfoTableFlowCtrl int32

const (
	PortInfoTableFlowCtrl_Transmit    PortInfoTableFlowCtrl = 2
	PortInfoTableFlowCtrl_Receive     PortInfoTableFlowCtrl = 3
	PortInfoTableFlowCtrl_Both        PortInfoTableFlowCtrl = 4
	PortInfoTableFlowCtrl_None        PortInfoTableFlowCtrl = 5
	PortInfoTableFlowCtrl_Unsupported PortInfoTableFlowCtrl = 2147483647
)

type PortInfoTableLink int32

const (
	PortInfoTableLink_Up          PortInfoTableLink = 1
	PortInfoTableLink_Down        PortInfoTableLink = 2
	PortInfoTableLink_Disabled    PortInfoTableLink = 3
	PortInfoTableLink_Inoperative PortInfoTableLink = 4
	PortInfoTableLink_Unsupported PortInfoTableLink = 2147483647
)

type PortInfoTablePhyIfType int32

const (
	PortInfoTablePhyIfType_Other                  PortInfoTablePhyIfType = 1
	PortInfoTablePhyIfType_Regular1822            PortInfoTablePhyIfType = 2
	PortInfoTablePhyIfType_Hdh1822                PortInfoTablePhyIfType = 3
	PortInfoTablePhyIfType_DdnX25                 PortInfoTablePhyIfType = 4
	PortInfoTablePhyIfType_Rfc877X25              PortInfoTablePhyIfType = 5
	PortInfoTablePhyIfType_EthernetCsmacd         PortInfoTablePhyIfType = 6
	PortInfoTablePhyIfType_Iso88023Csmacd         PortInfoTablePhyIfType = 7
	PortInfoTablePhyIfType_Iso88024TokenBus       PortInfoTablePhyIfType = 8
	PortInfoTablePhyIfType_Iso88025TokenRing      PortInfoTablePhyIfType = 9
	PortInfoTablePhyIfType_Iso88026Man            PortInfoTablePhyIfType = 10
	PortInfoTablePhyIfType_StarLan                PortInfoTablePhyIfType = 11
	PortInfoTablePhyIfType_Proteon10Mbit          PortInfoTablePhyIfType = 12
	PortInfoTablePhyIfType_Proteon80Mbit          PortInfoTablePhyIfType = 13
	PortInfoTablePhyIfType_Hyperchannel           PortInfoTablePhyIfType = 14
	PortInfoTablePhyIfType_Fddi                   PortInfoTablePhyIfType = 15
	PortInfoTablePhyIfType_Lapb                   PortInfoTablePhyIfType = 16
	PortInfoTablePhyIfType_Sdlc                   PortInfoTablePhyIfType = 17
	PortInfoTablePhyIfType_Ds1                    PortInfoTablePhyIfType = 18
	PortInfoTablePhyIfType_E1                     PortInfoTablePhyIfType = 19
	PortInfoTablePhyIfType_BasicISDN              PortInfoTablePhyIfType = 20
	PortInfoTablePhyIfType_PrimaryISDN            PortInfoTablePhyIfType = 21
	PortInfoTablePhyIfType_PropPointToPointSerial PortInfoTablePhyIfType = 22
	PortInfoTablePhyIfType_Ppp                    PortInfoTablePhyIfType = 23
	PortInfoTablePhyIfType_SoftwareLoopback       PortInfoTablePhyIfType = 24
	PortInfoTablePhyIfType_Eon                    PortInfoTablePhyIfType = 25
	PortInfoTablePhyIfType_Ethernet3Mbit          PortInfoTablePhyIfType = 26
	PortInfoTablePhyIfType_Nsip                   PortInfoTablePhyIfType = 27
	PortInfoTablePhyIfType_Slip                   PortInfoTablePhyIfType = 28
	PortInfoTablePhyIfType_Ultra                  PortInfoTablePhyIfType = 29
	PortInfoTablePhyIfType_Ds3                    PortInfoTablePhyIfType = 30
	PortInfoTablePhyIfType_Sip                    PortInfoTablePhyIfType = 31
	PortInfoTablePhyIfType_FrameRelay             PortInfoTablePhyIfType = 32
	PortInfoTablePhyIfType_Unsupported            PortInfoTablePhyIfType = 2147483647
)

type PortInfoTablePhyIfOperStatus int32

const (
	PortInfoTablePhyIfOperStatus_Up          PortInfoTablePhyIfOperStatus = 1
	PortInfoTablePhyIfOperStatus_Down        PortInfoTablePhyIfOperStatus = 2
	PortInfoTablePhyIfOperStatus_Testing     PortInfoTablePhyIfOperStatus = 3
	PortInfoTablePhyIfOperStatus_Unsupported PortInfoTablePhyIfOperStatus = 2147483647
)

type PortInfoTablePhyConnType int32

const (
	PortInfoTablePhyConnType_FeCopper    PortInfoTablePhyConnType = 1
	PortInfoTablePhyConnType_GeCopper    PortInfoTablePhyConnType = 2
	PortInfoTablePhyConnType_GeSFP       PortInfoTablePhyConnType = 3
	PortInfoTablePhyConnType_Unknown     PortInfoTablePhyConnType = 4
	PortInfoTablePhyConnType_XGeSFP      PortInfoTablePhyConnType = 5
	PortInfoTablePhyConnType_XGeQSFP     PortInfoTablePhyConnType = 6
	PortInfoTablePhyConnType_XGeQSFP28   PortInfoTablePhyConnType = 7
	PortInfoTablePhyConnType_XGeSFP28    PortInfoTablePhyConnType = 8
	PortInfoTablePhyConnType_Unsupported PortInfoTablePhyConnType = 2147483647
)

type PortInfoTablePreferred int32

const (
	PortInfoTablePreferred_Invalid     PortInfoTablePreferred = 1
	PortInfoTablePreferred_Copper      PortInfoTablePreferred = 2
	PortInfoTablePreferred_Sfp         PortInfoTablePreferred = 3
	PortInfoTablePreferred_Unsupported PortInfoTablePreferred = 2147483647
)

type PortInfoTableBackup int32

const (
	PortInfoTableBackup_Invalid     PortInfoTableBackup = 1
	PortInfoTableBackup_None        PortInfoTableBackup = 2
	PortInfoTableBackup_Copper      PortInfoTableBackup = 3
	PortInfoTableBackup_Sfp         PortInfoTableBackup = 4
	PortInfoTableBackup_Unsupported PortInfoTableBackup = 2147483647
)

type PortInfoTableSFPType int32

const (
	PortInfoTableSFPType_Invalid       PortInfoTableSFPType = 1
	PortInfoTableSFPType_SfpTypeSX     PortInfoTableSFPType = 2
	PortInfoTableSFPType_SfpTypeLX     PortInfoTableSFPType = 3
	PortInfoTableSFPType_SfpTypeCX     PortInfoTableSFPType = 4
	PortInfoTableSFPType_SfpTypeCopper PortInfoTableSFPType = 5
	PortInfoTableSFPType_Unsupported   PortInfoTableSFPType = 2147483647
)

type PortInfoTableShared int32

const (
	PortInfoTableShared_Enabled     PortInfoTableShared = 1
	PortInfoTableShared_Disabled    PortInfoTableShared = 2
	PortInfoTableShared_Unsupported PortInfoTableShared = 2147483647
)

type PortInfoTableParams struct {
	// The port index.
	Indx int32 `json:"Indx,omitempty"`
	// The current operational speed of the port.
	Speed PortInfoTableSpeed `json:"Speed,omitempty"`
	// The current operational mode of the port.
	Mode PortInfoTableMode `json:"Mode,omitempty"`
	// The current operational flow control of the port.
	FlowCtrl PortInfoTableFlowCtrl `json:"FlowCtrl,omitempty"`
	// The current operational link status of the port.
	Link PortInfoTableLink `json:"Link,omitempty"`
	// A textual string containing information about the
	// 	    interface.  This string should include the name of
	// 	    the manufacturer, the product name and the version
	// 	    of the hardware interface.
	PhyIfDescr string `json:"PhyIfDescr,omitempty"`
	// The type of interface, distinguished according to
	// 	    the physical/link protocol(s) immediately `below'
	// 	    the network layer in the protocol stack.
	PhyIfType PortInfoTablePhyIfType `json:"PhyIfType,omitempty"`
	// The size of the largest datagram which can be
	// 	    sent/received on the interface, specified in
	// 	    octets.  For interfaces that are used for
	// 	    transmitting network datagrams, this is the size
	// 	    of the largest network datagram that can be sent
	// 	    on the interface.
	PhyIfMtu int32 `json:"PhyIfMtu,omitempty"`
	// The interface's address at the protocol layer
	// 	    immediately `below' the network layer in the
	// 	    protocol stack.  For interfaces which do not have
	// 	    such an address (e.g., a serial line), this object
	// 	    should contain an octet string of zero length.
	PhyIfPhysAddress string `json:"PhyIfPhysAddress,omitempty"`
	// The current operational state of the interface.
	// 	    The testing(3) state indicates that no operational
	// 	    packets can be passed.
	PhyIfOperStatus PortInfoTablePhyIfOperStatus `json:"PhyIfOperStatus,omitempty"`
	// The value of sysUpTime at the time the interface
	// 	    entered its current operational state.  If the
	// 	    current state was entered prior to the last re-
	// 	    initialization of the local network management
	// 	    subsystem, then this object contains a zero
	// 	    value.
	PhyIfLastChange uint32 `json:"PhyIfLastChange,omitempty"`
	// The physical connection type, i.e. copper or SFP port.
	PhyConnType PortInfoTablePhyConnType `json:"PhyConnType,omitempty"`
	// The port preferred connection type. For dual ports only
	Preferred PortInfoTablePreferred `json:"Preferred,omitempty"`
	// The port backup setting. For dual ports only
	Backup PortInfoTableBackup `json:"Backup,omitempty"`
	// A textual string containing information about the
	// 	    SFP manufacturer. For valid ports only.
	SFPName string `json:"SFPName,omitempty"`
	// The inserted SFP type. For valid ports only.
	SFPType PortInfoTableSFPType `json:"SFPType,omitempty"`
	// Indicates if the port is shared or not.
	Shared PortInfoTableShared `json:"Shared,omitempty"`
	// The vADC number that the port belongs to.
	// Note:This mib is supported only in VX and vADC instance
	VADC int32 `json:"VADC,omitempty"`
	// The Range of vADCS that the port belongs to.
	// Note:This mib is supported only in VX and vADC instance
	VADCs string `json:"VADCs,omitempty"`
}

func (p PortInfoTableParams) iMABean() {}
