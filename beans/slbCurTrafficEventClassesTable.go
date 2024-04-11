package beans

import (
	"fmt"
	"reflect"
)

// SlbCurTrafficEventClassesTable Current traffic event table.
// Note:This mib is not supported on VX instance of Virtualization.
type SlbCurTrafficEventClassesTable struct {
	// Traffic event class id.
	SlbCurTrafficEventId string
	Params               *SlbCurTrafficEventClassesTableParams
}

func NewSlbCurTrafficEventClassesTableList() *SlbCurTrafficEventClassesTable {
	return &SlbCurTrafficEventClassesTable{}
}

func NewSlbCurTrafficEventClassesTable(
	slbCurTrafficEventId string,
	params *SlbCurTrafficEventClassesTableParams,
) *SlbCurTrafficEventClassesTable {
	return &SlbCurTrafficEventClassesTable{
		SlbCurTrafficEventId: slbCurTrafficEventId,
		Params:               params,
	}
}

func (c *SlbCurTrafficEventClassesTable) Name() string {
	return "SlbCurTrafficEventClassesTable"
}

func (c *SlbCurTrafficEventClassesTable) GetParams() BeanType {
	return c.Params
}

func (c *SlbCurTrafficEventClassesTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *SlbCurTrafficEventClassesTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.SlbCurTrafficEventId).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.SlbCurTrafficEventId)
}

type SlbCurTrafficEventClassesTableEnaDis int32

const (
	SlbCurTrafficEventClassesTableEnaDis_Enabled     SlbCurTrafficEventClassesTableEnaDis = 1
	SlbCurTrafficEventClassesTableEnaDis_Disabled    SlbCurTrafficEventClassesTableEnaDis = 2
	SlbCurTrafficEventClassesTableEnaDis_Unsupported SlbCurTrafficEventClassesTableEnaDis = 2147483647
)

type SlbCurTrafficEventClassesTableHttpTran int32

const (
	SlbCurTrafficEventClassesTableHttpTran_Enabled     SlbCurTrafficEventClassesTableHttpTran = 1
	SlbCurTrafficEventClassesTableHttpTran_Disabled    SlbCurTrafficEventClassesTableHttpTran = 2
	SlbCurTrafficEventClassesTableHttpTran_Unsupported SlbCurTrafficEventClassesTableHttpTran = 2147483647
)

type SlbCurTrafficEventClassesTablePathCorl int32

const (
	SlbCurTrafficEventClassesTablePathCorl_Disabled    SlbCurTrafficEventClassesTablePathCorl = 1
	SlbCurTrafficEventClassesTablePathCorl_Entry       SlbCurTrafficEventClassesTablePathCorl = 2
	SlbCurTrafficEventClassesTablePathCorl_Exit        SlbCurTrafficEventClassesTablePathCorl = 3
	SlbCurTrafficEventClassesTablePathCorl_Unsupported SlbCurTrafficEventClassesTablePathCorl = 2147483647
)

type SlbCurTrafficEventClassesTableSslConn int32

const (
	SlbCurTrafficEventClassesTableSslConn_Disabled    SlbCurTrafficEventClassesTableSslConn = 1
	SlbCurTrafficEventClassesTableSslConn_Frontend    SlbCurTrafficEventClassesTableSslConn = 2
	SlbCurTrafficEventClassesTableSslConn_Backend     SlbCurTrafficEventClassesTableSslConn = 3
	SlbCurTrafficEventClassesTableSslConn_Both        SlbCurTrafficEventClassesTableSslConn = 4
	SlbCurTrafficEventClassesTableSslConn_Unsupported SlbCurTrafficEventClassesTableSslConn = 2147483647
)

type SlbCurTrafficEventClassesTableSslFail int32

const (
	SlbCurTrafficEventClassesTableSslFail_Disabled    SlbCurTrafficEventClassesTableSslFail = 1
	SlbCurTrafficEventClassesTableSslFail_Frontend    SlbCurTrafficEventClassesTableSslFail = 2
	SlbCurTrafficEventClassesTableSslFail_Backend     SlbCurTrafficEventClassesTableSslFail = 3
	SlbCurTrafficEventClassesTableSslFail_Both        SlbCurTrafficEventClassesTableSslFail = 4
	SlbCurTrafficEventClassesTableSslFail_Unsupported SlbCurTrafficEventClassesTableSslFail = 2147483647
)

type SlbCurTrafficEventClassesTableHostByPass int32

const (
	SlbCurTrafficEventClassesTableHostByPass_Enabled     SlbCurTrafficEventClassesTableHostByPass = 1
	SlbCurTrafficEventClassesTableHostByPass_Disabled    SlbCurTrafficEventClassesTableHostByPass = 2
	SlbCurTrafficEventClassesTableHostByPass_Unsupported SlbCurTrafficEventClassesTableHostByPass = 2147483647
)

type SlbCurTrafficEventClassesTableL4Conn int32

const (
	SlbCurTrafficEventClassesTableL4Conn_Enabled     SlbCurTrafficEventClassesTableL4Conn = 1
	SlbCurTrafficEventClassesTableL4Conn_Disabled    SlbCurTrafficEventClassesTableL4Conn = 2
	SlbCurTrafficEventClassesTableL4Conn_Unsupported SlbCurTrafficEventClassesTableL4Conn = 2147483647
)

type SlbCurTrafficEventClassesTableUnified int32

const (
	SlbCurTrafficEventClassesTableUnified_Enabled     SlbCurTrafficEventClassesTableUnified = 1
	SlbCurTrafficEventClassesTableUnified_Disabled    SlbCurTrafficEventClassesTableUnified = 2
	SlbCurTrafficEventClassesTableUnified_Unsupported SlbCurTrafficEventClassesTableUnified = 2147483647
)

type SlbCurTrafficEventClassesTableSecurity int32

const (
	SlbCurTrafficEventClassesTableSecurity_Enabled     SlbCurTrafficEventClassesTableSecurity = 1
	SlbCurTrafficEventClassesTableSecurity_Disabled    SlbCurTrafficEventClassesTableSecurity = 2
	SlbCurTrafficEventClassesTableSecurity_Unsupported SlbCurTrafficEventClassesTableSecurity = 2147483647
)

type SlbCurTrafficEventClassesTableParams struct {
	// Traffic event class id.
	Id string `json:"Id,omitempty"`
	// Traffic event class name.
	Name string `json:"Name,omitempty"`
	// Traffic event class attach group name.
	GroupName string `json:"GroupName,omitempty"`
	// Traffic event log policy syslog server port.
	SysPort uint64 `json:"SysPort,omitempty"`
	// Traffic event log policy sampling rate.
	SampleRate uint32 `json:"SampleRate,omitempty"`
	// Traffic event policy current status as Ena or Dis.
	EnaDis SlbCurTrafficEventClassesTableEnaDis `json:"EnaDis,omitempty"`
	// Traffic event policy http transaction status.
	HttpTran SlbCurTrafficEventClassesTableHttpTran `json:"HttpTran,omitempty"`
	// Traffic event policy http transaction path correlation.
	PathCorl SlbCurTrafficEventClassesTablePathCorl `json:"PathCorl,omitempty"`
	// Traffic event policy ssl connection status.
	SslConn SlbCurTrafficEventClassesTableSslConn `json:"SslConn,omitempty"`
	// Traffic event policy ssl connection fail status.
	SslFail SlbCurTrafficEventClassesTableSslFail `json:"SslFail,omitempty"`
	// Traffic event policy hotbypass current status as Ena or Dis.
	HostByPass SlbCurTrafficEventClassesTableHostByPass `json:"HostByPass,omitempty"`
	// Traffic event policy current status as Ena or Dis.
	L4Conn SlbCurTrafficEventClassesTableL4Conn `json:"L4Conn,omitempty"`
	// Current Remote logging object.
	RemoteLogId string `json:"RemoteLogId,omitempty"`
	// Traffic event policy unified event current status as Ena or Dis.
	Unified SlbCurTrafficEventClassesTableUnified `json:"Unified,omitempty"`
	// Traffic event log policy normal EPS limit.
	NormalEps uint32 `json:"NormalEps,omitempty"`
	// Traffic event log policy exception EPS limit.
	ExceptionEps uint32 `json:"ExceptionEps,omitempty"`
	// Traffic event policy security event current status as Ena or Dis.
	Security SlbCurTrafficEventClassesTableSecurity `json:"Security,omitempty"`
	// Traffic event policy serialization format CEF or JSON.
	Format string `json:"Format,omitempty"`
}

func (p SlbCurTrafficEventClassesTableParams) iMABean() {}
