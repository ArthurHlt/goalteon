package beans

import (
	"fmt"
	"reflect"
)

// SlbNewAdvhcSnmpTable Note:This mib is not supported on VX instance of Virtualization.
type SlbNewAdvhcSnmpTable struct {
	// SNMP Health check id.
	SlbNewAdvhcSnmpID string
	Params            *SlbNewAdvhcSnmpTableParams
}

func NewSlbNewAdvhcSnmpTableList() *SlbNewAdvhcSnmpTable {
	return &SlbNewAdvhcSnmpTable{}
}

func NewSlbNewAdvhcSnmpTable(
	slbNewAdvhcSnmpID string,
	params *SlbNewAdvhcSnmpTableParams,
) *SlbNewAdvhcSnmpTable {
	return &SlbNewAdvhcSnmpTable{
		SlbNewAdvhcSnmpID: slbNewAdvhcSnmpID,
		Params:            params,
	}
}

func (c *SlbNewAdvhcSnmpTable) Name() string {
	return "SlbNewAdvhcSnmpTable"
}

func (c *SlbNewAdvhcSnmpTable) GetParams() BeanType {
	return c.Params
}

func (c *SlbNewAdvhcSnmpTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *SlbNewAdvhcSnmpTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.SlbNewAdvhcSnmpID).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.SlbNewAdvhcSnmpID)
}

type SlbNewAdvhcSnmpTableIPVer int32

const (
	SlbNewAdvhcSnmpTableIPVer_Ipv4 SlbNewAdvhcSnmpTableIPVer = 1
	SlbNewAdvhcSnmpTableIPVer_Ipv6 SlbNewAdvhcSnmpTableIPVer = 2
	SlbNewAdvhcSnmpTableIPVer_None SlbNewAdvhcSnmpTableIPVer = 3
)

type SlbNewAdvhcSnmpTableTransparent int32

const (
	SlbNewAdvhcSnmpTableTransparent_Enabled     SlbNewAdvhcSnmpTableTransparent = 1
	SlbNewAdvhcSnmpTableTransparent_Disabled    SlbNewAdvhcSnmpTableTransparent = 2
	SlbNewAdvhcSnmpTableTransparent_Unsupported SlbNewAdvhcSnmpTableTransparent = 2147483647
)

type SlbNewAdvhcSnmpTableInvert int32

const (
	SlbNewAdvhcSnmpTableInvert_Enabled     SlbNewAdvhcSnmpTableInvert = 1
	SlbNewAdvhcSnmpTableInvert_Disabled    SlbNewAdvhcSnmpTableInvert = 2
	SlbNewAdvhcSnmpTableInvert_Unsupported SlbNewAdvhcSnmpTableInvert = 2147483647
)

type SlbNewAdvhcSnmpTableType int32

const (
	SlbNewAdvhcSnmpTableType_Integer     SlbNewAdvhcSnmpTableType = 1
	SlbNewAdvhcSnmpTableType_String      SlbNewAdvhcSnmpTableType = 2
	SlbNewAdvhcSnmpTableType_Unsupported SlbNewAdvhcSnmpTableType = 2147483647
)

type SlbNewAdvhcSnmpTableReadjustWeight int32

const (
	SlbNewAdvhcSnmpTableReadjustWeight_Enabled     SlbNewAdvhcSnmpTableReadjustWeight = 1
	SlbNewAdvhcSnmpTableReadjustWeight_Disabled    SlbNewAdvhcSnmpTableReadjustWeight = 2
	SlbNewAdvhcSnmpTableReadjustWeight_Inverted    SlbNewAdvhcSnmpTableReadjustWeight = 3
	SlbNewAdvhcSnmpTableReadjustWeight_Unsupported SlbNewAdvhcSnmpTableReadjustWeight = 2147483647
)

type SlbNewAdvhcSnmpTableDelete int32

const (
	SlbNewAdvhcSnmpTableDelete_Other       SlbNewAdvhcSnmpTableDelete = 1
	SlbNewAdvhcSnmpTableDelete_Delete      SlbNewAdvhcSnmpTableDelete = 2
	SlbNewAdvhcSnmpTableDelete_Unsupported SlbNewAdvhcSnmpTableDelete = 2147483647
)

type SlbNewAdvhcSnmpTableParams struct {
	// SNMP Health check id.
	ID string `json:"ID,omitempty"`
	// SNMP Health check name.
	Name string `json:"Name,omitempty"`
	// SNMP Health check destination port.
	DPort uint64 `json:"DPort,omitempty"`
	// SNMP Health check destination IP version.
	IPVer SlbNewAdvhcSnmpTableIPVer `json:"IPVer,omitempty"`
	// SNMP Health check destination hostname.
	HostName string `json:"HostName,omitempty"`
	// SNMP Health check transparent flag.
	Transparent SlbNewAdvhcSnmpTableTransparent `json:"Transparent,omitempty"`
	// SNMP Health check interval.
	Interval uint64 `json:"Interval,omitempty"`
	// SNMP Health check retries counter.
	Retries uint64 `json:"Retries,omitempty"`
	// SNMP Health check retries counter in the down state.
	RestoreRetries uint64 `json:"RestoreRetries,omitempty"`
	// SNMP Health check timeout.
	Timeout uint64 `json:"Timeout,omitempty"`
	// SNMP Health check overflow flag.
	Overflow uint64 `json:"Overflow,omitempty"`
	// SNMP Health check interval when in the down state.
	DownInterval uint64 `json:"DownInterval,omitempty"`
	// SNMP Health check invert flag.
	Invert SlbNewAdvhcSnmpTableInvert `json:"Invert,omitempty"`
	// SNMP Health check OID.
	Oid string `json:"Oid,omitempty"`
	// SNMP Health check community string.
	Community string `json:"Community,omitempty"`
	// SNMP Health check OID type.
	Type SlbNewAdvhcSnmpTableType `json:"Type,omitempty"`
	// SNMP Health check OID variable minimum value.
	MinValue int32 `json:"MinValue,omitempty"`
	// SNMP Health check OID maximum value.
	MaxValue int32 `json:"MaxValue,omitempty"`
	// SNMP Health check receive string.
	ReceiveString string `json:"ReceiveString,omitempty"`
	// Enabled: value received in SNMP health check response is used as real server weight.
	// Disabled: value received in SNMP health check response is not used as real server weight.
	// Inverted: value received in SNMP health check response is subtracted from 100 (100 - value) is used as real server weight
	ReadjustWeight SlbNewAdvhcSnmpTableReadjustWeight `json:"ReadjustWeight,omitempty"`
	// SNMP Health check copy indicator.
	Copy string `json:"Copy,omitempty"`
	// When set to the value of 2 (delete), the entire row is deleted.
	// When read, other(1) is returned. Setting the value to anything
	// other than 2(delete) has no effect on the state of the row.
	Delete SlbNewAdvhcSnmpTableDelete `json:"Delete,omitempty"`
	// Oveload minimum value.
	OverloadMinValue int32 `json:"OverloadMinValue,omitempty"`
	// Oveload maximum value.
	OverloadMaxValue int32 `json:"OverloadMaxValue,omitempty"`
	// expected response for server overload.
	OverloadString string `json:"OverloadString,omitempty"`
}

func (p SlbNewAdvhcSnmpTableParams) iMABean() {}
