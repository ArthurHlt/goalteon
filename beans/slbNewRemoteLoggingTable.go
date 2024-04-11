package beans

import (
	"fmt"
	"reflect"
)

// SlbNewRemoteLoggingTable Remote logging new configuration table.
// Note:This mib is not supported for VX instance of Virtualization.
type SlbNewRemoteLoggingTable struct {
	// Remote logging Index.
	SlbNewRemoteLoggingID string
	Params                *SlbNewRemoteLoggingTableParams
}

func NewSlbNewRemoteLoggingTableList() *SlbNewRemoteLoggingTable {
	return &SlbNewRemoteLoggingTable{}
}

func NewSlbNewRemoteLoggingTable(
	slbNewRemoteLoggingID string,
	params *SlbNewRemoteLoggingTableParams,
) *SlbNewRemoteLoggingTable {
	return &SlbNewRemoteLoggingTable{
		SlbNewRemoteLoggingID: slbNewRemoteLoggingID,
		Params:                params,
	}
}

func (c *SlbNewRemoteLoggingTable) Name() string {
	return "SlbNewRemoteLoggingTable"
}

func (c *SlbNewRemoteLoggingTable) GetParams() BeanType {
	return c.Params
}

func (c *SlbNewRemoteLoggingTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *SlbNewRemoteLoggingTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.SlbNewRemoteLoggingID).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.SlbNewRemoteLoggingID)
}

type SlbNewRemoteLoggingTableProtocol int32

const (
	SlbNewRemoteLoggingTableProtocol_Udp         SlbNewRemoteLoggingTableProtocol = 0
	SlbNewRemoteLoggingTableProtocol_Tcp         SlbNewRemoteLoggingTableProtocol = 1
	SlbNewRemoteLoggingTableProtocol_Unsupported SlbNewRemoteLoggingTableProtocol = 2147483647
)

type SlbNewRemoteLoggingTableEnaDis int32

const (
	SlbNewRemoteLoggingTableEnaDis_Enabled     SlbNewRemoteLoggingTableEnaDis = 1
	SlbNewRemoteLoggingTableEnaDis_Disabled    SlbNewRemoteLoggingTableEnaDis = 2
	SlbNewRemoteLoggingTableEnaDis_Unsupported SlbNewRemoteLoggingTableEnaDis = 2147483647
)

type SlbNewRemoteLoggingTableDel int32

const (
	SlbNewRemoteLoggingTableDel_Other       SlbNewRemoteLoggingTableDel = 1
	SlbNewRemoteLoggingTableDel_Delete      SlbNewRemoteLoggingTableDel = 2
	SlbNewRemoteLoggingTableDel_Unsupported SlbNewRemoteLoggingTableDel = 2147483647
)

type SlbNewRemoteLoggingTablePath int32

const (
	SlbNewRemoteLoggingTablePath_Data       SlbNewRemoteLoggingTablePath = 1
	SlbNewRemoteLoggingTablePath_Management SlbNewRemoteLoggingTablePath = 2
)

type SlbNewRemoteLoggingTableParams struct {
	// Remote logging Index.
	ID string `json:"ID,omitempty"`
	// Remote logging protocol type.
	Protocol SlbNewRemoteLoggingTableProtocol `json:"Protocol,omitempty"`
	// Remote logging server port.
	Port uint64 `json:"Port,omitempty"`
	// Remote logging group.
	Group string `json:"Group,omitempty"`
	// Remote logging SSL policy.
	SslPol string `json:"SslPol,omitempty"`
	// Enable/Disable remote logging server.
	EnaDis SlbNewRemoteLoggingTableEnaDis `json:"EnaDis,omitempty"`
	// When set to the value of 2 (delete), the entire row is deleted.
	// When read, other(1) is returned. Setting the value to anything
	// other than 2(delete) has no effect on the state of the row.
	Del SlbNewRemoteLoggingTableDel `json:"Del,omitempty"`
	// Remote logging path through data or management port.
	Path SlbNewRemoteLoggingTablePath `json:"Path,omitempty"`
}

func (p SlbNewRemoteLoggingTableParams) iMABean() {}
