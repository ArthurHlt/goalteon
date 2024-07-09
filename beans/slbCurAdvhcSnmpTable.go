package beans

import (
	"fmt"
	"reflect"
)

// SlbCurAdvhcSnmpTable Note:This mib is not supported on VX instance of Virtualization.
type SlbCurAdvhcSnmpTable struct {
	// SNMP Health check id.
	SlbCurAdvhcSnmpID string
	Params            *SlbCurAdvhcSnmpTableParams
}

func NewSlbCurAdvhcSnmpTableList() *SlbCurAdvhcSnmpTable {
	return &SlbCurAdvhcSnmpTable{}
}

func NewSlbCurAdvhcSnmpTable(
	slbCurAdvhcSnmpID string,
	params *SlbCurAdvhcSnmpTableParams,
) *SlbCurAdvhcSnmpTable {
	return &SlbCurAdvhcSnmpTable{
		SlbCurAdvhcSnmpID: slbCurAdvhcSnmpID,
		Params:            params,
	}
}

func (c *SlbCurAdvhcSnmpTable) Name() string {
	return "SlbCurAdvhcSnmpTable"
}

func (c *SlbCurAdvhcSnmpTable) GetParams() BeanType {
	return c.Params
}

func (c *SlbCurAdvhcSnmpTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *SlbCurAdvhcSnmpTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.SlbCurAdvhcSnmpID).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.SlbCurAdvhcSnmpID)
}

type SlbCurAdvhcSnmpTableIPVer int32

const (
	SlbCurAdvhcSnmpTableIPVer_Ipv4 SlbCurAdvhcSnmpTableIPVer = 1
	SlbCurAdvhcSnmpTableIPVer_Ipv6 SlbCurAdvhcSnmpTableIPVer = 2
	SlbCurAdvhcSnmpTableIPVer_None SlbCurAdvhcSnmpTableIPVer = 3
)

type SlbCurAdvhcSnmpTableTransparent int32

const (
	SlbCurAdvhcSnmpTableTransparent_Enabled     SlbCurAdvhcSnmpTableTransparent = 1
	SlbCurAdvhcSnmpTableTransparent_Disabled    SlbCurAdvhcSnmpTableTransparent = 2
	SlbCurAdvhcSnmpTableTransparent_Unsupported SlbCurAdvhcSnmpTableTransparent = 2147483647
)

type SlbCurAdvhcSnmpTableInvert int32

const (
	SlbCurAdvhcSnmpTableInvert_Enabled     SlbCurAdvhcSnmpTableInvert = 1
	SlbCurAdvhcSnmpTableInvert_Disabled    SlbCurAdvhcSnmpTableInvert = 2
	SlbCurAdvhcSnmpTableInvert_Unsupported SlbCurAdvhcSnmpTableInvert = 2147483647
)

type SlbCurAdvhcSnmpTableType int32

const (
	SlbCurAdvhcSnmpTableType_Integer     SlbCurAdvhcSnmpTableType = 1
	SlbCurAdvhcSnmpTableType_String      SlbCurAdvhcSnmpTableType = 2
	SlbCurAdvhcSnmpTableType_Unsupported SlbCurAdvhcSnmpTableType = 2147483647
)

type SlbCurAdvhcSnmpTableReadjustWeight int32

const (
	SlbCurAdvhcSnmpTableReadjustWeight_Enabled     SlbCurAdvhcSnmpTableReadjustWeight = 1
	SlbCurAdvhcSnmpTableReadjustWeight_Disabled    SlbCurAdvhcSnmpTableReadjustWeight = 2
	SlbCurAdvhcSnmpTableReadjustWeight_Inverted    SlbCurAdvhcSnmpTableReadjustWeight = 3
	SlbCurAdvhcSnmpTableReadjustWeight_Unsupported SlbCurAdvhcSnmpTableReadjustWeight = 2147483647
)

type SlbCurAdvhcSnmpTableSnat int32

const (
	SlbCurAdvhcSnmpTableSnat_Enabled     SlbCurAdvhcSnmpTableSnat = 1
	SlbCurAdvhcSnmpTableSnat_Disabled    SlbCurAdvhcSnmpTableSnat = 2
	SlbCurAdvhcSnmpTableSnat_Unsupported SlbCurAdvhcSnmpTableSnat = 2147483647
)

type SlbCurAdvhcSnmpTableParams struct {
	// SNMP Health check id.
	ID string `json:"ID,omitempty"`
	// SNMP Health check name.
	Name string `json:"Name,omitempty"`
	// SNMP Health check destination port.
	DPort uint64 `json:"DPort,omitempty"`
	// SNMP Health check destination IP version.
	IPVer SlbCurAdvhcSnmpTableIPVer `json:"IPVer,omitempty"`
	// SNMP Health check destination hostname.
	HostName string `json:"HostName,omitempty"`
	// SNMP Health check transparent flag.
	Transparent SlbCurAdvhcSnmpTableTransparent `json:"Transparent,omitempty"`
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
	Invert SlbCurAdvhcSnmpTableInvert `json:"Invert,omitempty"`
	// SNMP Health check OID.
	Oid string `json:"Oid,omitempty"`
	// SNMP Health check community string.
	Community string `json:"Community,omitempty"`
	// SNMP Health check OID type.
	Type SlbCurAdvhcSnmpTableType `json:"Type,omitempty"`
	// SNMP Health check OID variable minimum value.
	MinValue int32 `json:"MinValue,omitempty"`
	// SNMP Health check OID maximum value.
	MaxValue int32 `json:"MaxValue,omitempty"`
	// SNMP Health check receive string.
	ReceiveString string `json:"ReceiveString,omitempty"`
	// Enabled: value received in SNMP health check response is used as real server weight.
	// Disabled: value received in SNMP health check response is not used as real server weight.
	// Inverted: value received in SNMP health check response is subtracted from 100 (100 - value) is used as real server weight
	ReadjustWeight SlbCurAdvhcSnmpTableReadjustWeight `json:"ReadjustWeight,omitempty"`
	// .
	OverloadMinValue int32 `json:"OverloadMinValue,omitempty"`
	// .
	OverloadMaxValue int32 `json:"OverloadMaxValue,omitempty"`
	// .
	OverloadString string `json:"OverloadString,omitempty"`
	// SNMP Health Check src NAT (PIP) flag.
	Snat SlbCurAdvhcSnmpTableSnat `json:"Snat,omitempty"`
}

func (p SlbCurAdvhcSnmpTableParams) iMABean() {}
