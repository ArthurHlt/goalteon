package beans

import (
	"fmt"
	"reflect"
)

// SlbAdvhcGeneralTable Note:This mib is not supported on VX instance of Virtualization.
type SlbAdvhcGeneralTable struct {
	// Health check ID.
	SlbAdvhcGeneralID string
	Params            *SlbAdvhcGeneralTableParams
}

func NewSlbAdvhcGeneralTableList() *SlbAdvhcGeneralTable {
	return &SlbAdvhcGeneralTable{}
}

func NewSlbAdvhcGeneralTable(
	slbAdvhcGeneralID string,
	params *SlbAdvhcGeneralTableParams,
) *SlbAdvhcGeneralTable {
	return &SlbAdvhcGeneralTable{
		SlbAdvhcGeneralID: slbAdvhcGeneralID,
		Params:            params,
	}
}

func (c *SlbAdvhcGeneralTable) Name() string {
	return "SlbAdvhcGeneralTable"
}

func (c *SlbAdvhcGeneralTable) GetParams() BeanType {
	return c.Params
}

func (c *SlbAdvhcGeneralTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *SlbAdvhcGeneralTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.SlbAdvhcGeneralID).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.SlbAdvhcGeneralID)
}

type SlbAdvhcGeneralTableType int32

const (
	SlbAdvhcGeneralTableType_Empty          SlbAdvhcGeneralTableType = 0
	SlbAdvhcGeneralTableType_Arp            SlbAdvhcGeneralTableType = 1
	SlbAdvhcGeneralTableType_Dhcp           SlbAdvhcGeneralTableType = 2
	SlbAdvhcGeneralTableType_Dns            SlbAdvhcGeneralTableType = 3
	SlbAdvhcGeneralTableType_Ftp            SlbAdvhcGeneralTableType = 4
	SlbAdvhcGeneralTableType_Http           SlbAdvhcGeneralTableType = 5
	SlbAdvhcGeneralTableType_Imap           SlbAdvhcGeneralTableType = 6
	SlbAdvhcGeneralTableType_Ldap           SlbAdvhcGeneralTableType = 7
	SlbAdvhcGeneralTableType_Db             SlbAdvhcGeneralTableType = 8
	SlbAdvhcGeneralTableType_Nntp           SlbAdvhcGeneralTableType = 9
	SlbAdvhcGeneralTableType_Icmp           SlbAdvhcGeneralTableType = 10
	SlbAdvhcGeneralTableType_Pop3           SlbAdvhcGeneralTableType = 11
	SlbAdvhcGeneralTableType_Radius         SlbAdvhcGeneralTableType = 12
	SlbAdvhcGeneralTableType_Rtsp           SlbAdvhcGeneralTableType = 13
	SlbAdvhcGeneralTableType_Sip            SlbAdvhcGeneralTableType = 14
	SlbAdvhcGeneralTableType_Smtp           SlbAdvhcGeneralTableType = 15
	SlbAdvhcGeneralTableType_Snmp           SlbAdvhcGeneralTableType = 16
	SlbAdvhcGeneralTableType_Ssl            SlbAdvhcGeneralTableType = 17
	SlbAdvhcGeneralTableType_Tcp            SlbAdvhcGeneralTableType = 18
	SlbAdvhcGeneralTableType_Tftp           SlbAdvhcGeneralTableType = 19
	SlbAdvhcGeneralTableType_Udp            SlbAdvhcGeneralTableType = 20
	SlbAdvhcGeneralTableType_Wap            SlbAdvhcGeneralTableType = 21
	SlbAdvhcGeneralTableType_Wts            SlbAdvhcGeneralTableType = 22
	SlbAdvhcGeneralTableType_Script         SlbAdvhcGeneralTableType = 23
	SlbAdvhcGeneralTableType_Link           SlbAdvhcGeneralTableType = 24
	SlbAdvhcGeneralTableType_Logexp         SlbAdvhcGeneralTableType = 25
	SlbAdvhcGeneralTableType_Virtualwire    SlbAdvhcGeneralTableType = 26
	SlbAdvhcGeneralTableType_Dssp           SlbAdvhcGeneralTableType = 27
	SlbAdvhcGeneralTableType_Clusthcfr      SlbAdvhcGeneralTableType = 28
	SlbAdvhcGeneralTableType_Clusthcme      SlbAdvhcGeneralTableType = 29
	SlbAdvhcGeneralTableType_Advvirtualwire SlbAdvhcGeneralTableType = 30
	SlbAdvhcGeneralTableType_Ocsp           SlbAdvhcGeneralTableType = 32
	SlbAdvhcGeneralTableType_Unsupported    SlbAdvhcGeneralTableType = 2147483647
)

type SlbAdvhcGeneralTableIPVer int32

const (
	SlbAdvhcGeneralTableIPVer_Ipv4 SlbAdvhcGeneralTableIPVer = 1
	SlbAdvhcGeneralTableIPVer_Ipv6 SlbAdvhcGeneralTableIPVer = 2
	SlbAdvhcGeneralTableIPVer_None SlbAdvhcGeneralTableIPVer = 3
)

type SlbAdvhcGeneralTableParams struct {
	// Health check ID.
	ID string `json:"ID,omitempty"`
	// Health check type.
	Type SlbAdvhcGeneralTableType `json:"Type,omitempty"`
	// Health check name.
	Name string `json:"Name,omitempty"`
	// Health check destination port.
	DPort uint64 `json:"DPort,omitempty"`
	// Health check IP version.
	IPVer SlbAdvhcGeneralTableIPVer `json:"IPVer,omitempty"`
	// Health check hostname.
	HostName string `json:"HostName,omitempty"`
}

func (p SlbAdvhcGeneralTableParams) iMABean() {}
