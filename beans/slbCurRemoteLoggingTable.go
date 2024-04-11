package beans

import (
	"fmt"
	"reflect"
)

// SlbCurRemoteLoggingTable Remote logging current configuration table.
// Note:This mib is not supported for VX instance of Virtualization.
type SlbCurRemoteLoggingTable struct {
	// Remote logging Index.
	SlbCurRemoteLoggingID string
	Params                *SlbCurRemoteLoggingTableParams
}

func NewSlbCurRemoteLoggingTableList() *SlbCurRemoteLoggingTable {
	return &SlbCurRemoteLoggingTable{}
}

func NewSlbCurRemoteLoggingTable(
	slbCurRemoteLoggingID string,
	params *SlbCurRemoteLoggingTableParams,
) *SlbCurRemoteLoggingTable {
	return &SlbCurRemoteLoggingTable{
		SlbCurRemoteLoggingID: slbCurRemoteLoggingID,
		Params:                params,
	}
}

func (c *SlbCurRemoteLoggingTable) Name() string {
	return "SlbCurRemoteLoggingTable"
}

func (c *SlbCurRemoteLoggingTable) GetParams() BeanType {
	return c.Params
}

func (c *SlbCurRemoteLoggingTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *SlbCurRemoteLoggingTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.SlbCurRemoteLoggingID).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.SlbCurRemoteLoggingID)
}

type SlbCurRemoteLoggingTableProtocol int32

const (
	SlbCurRemoteLoggingTableProtocol_Udp         SlbCurRemoteLoggingTableProtocol = 0
	SlbCurRemoteLoggingTableProtocol_Tcp         SlbCurRemoteLoggingTableProtocol = 1
	SlbCurRemoteLoggingTableProtocol_Unsupported SlbCurRemoteLoggingTableProtocol = 2147483647
)

type SlbCurRemoteLoggingTableEnaDis int32

const (
	SlbCurRemoteLoggingTableEnaDis_Enabled     SlbCurRemoteLoggingTableEnaDis = 1
	SlbCurRemoteLoggingTableEnaDis_Disabled    SlbCurRemoteLoggingTableEnaDis = 2
	SlbCurRemoteLoggingTableEnaDis_Unsupported SlbCurRemoteLoggingTableEnaDis = 2147483647
)

type SlbCurRemoteLoggingTablePath int32

const (
	SlbCurRemoteLoggingTablePath_Data       SlbCurRemoteLoggingTablePath = 1
	SlbCurRemoteLoggingTablePath_Management SlbCurRemoteLoggingTablePath = 2
)

type SlbCurRemoteLoggingTableParams struct {
	// Remote logging Index.
	ID string `json:"ID,omitempty"`
	// Remote logging protocol type.
	Protocol SlbCurRemoteLoggingTableProtocol `json:"Protocol,omitempty"`
	// Remote logging server port.
	Port uint64 `json:"Port,omitempty"`
	// Remote logging group.
	Group string `json:"Group,omitempty"`
	// Remote logging SSL policy.
	SslPol string `json:"SslPol,omitempty"`
	// Enable/Disable remote logging server.
	EnaDis SlbCurRemoteLoggingTableEnaDis `json:"EnaDis,omitempty"`
	// Remote logging path through data or management port.
	Path SlbCurRemoteLoggingTablePath `json:"Path,omitempty"`
}

func (p SlbCurRemoteLoggingTableParams) iMABean() {}
