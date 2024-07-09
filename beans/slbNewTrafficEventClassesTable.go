package beans

import (
	"fmt"
	"reflect"
)

// SlbNewTrafficEventClassesTable Current traffic event table.
// Note:This mib is not supported on VX instance of Virtualization.
type SlbNewTrafficEventClassesTable struct {
	// Set traffic event class id.
	SlbNewTrafficEventId string
	Params               *SlbNewTrafficEventClassesTableParams
}

func NewSlbNewTrafficEventClassesTableList() *SlbNewTrafficEventClassesTable {
	return &SlbNewTrafficEventClassesTable{}
}

func NewSlbNewTrafficEventClassesTable(
	slbNewTrafficEventId string,
	params *SlbNewTrafficEventClassesTableParams,
) *SlbNewTrafficEventClassesTable {
	return &SlbNewTrafficEventClassesTable{
		SlbNewTrafficEventId: slbNewTrafficEventId,
		Params:               params,
	}
}

func (c *SlbNewTrafficEventClassesTable) Name() string {
	return "SlbNewTrafficEventClassesTable"
}

func (c *SlbNewTrafficEventClassesTable) GetParams() BeanType {
	return c.Params
}

func (c *SlbNewTrafficEventClassesTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *SlbNewTrafficEventClassesTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.SlbNewTrafficEventId).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.SlbNewTrafficEventId)
}

type SlbNewTrafficEventClassesTableEnaDis int32

const (
	SlbNewTrafficEventClassesTableEnaDis_Enabled     SlbNewTrafficEventClassesTableEnaDis = 1
	SlbNewTrafficEventClassesTableEnaDis_Disabled    SlbNewTrafficEventClassesTableEnaDis = 2
	SlbNewTrafficEventClassesTableEnaDis_Unsupported SlbNewTrafficEventClassesTableEnaDis = 2147483647
)

type SlbNewTrafficEventClassesTableHttpTran int32

const (
	SlbNewTrafficEventClassesTableHttpTran_Enabled     SlbNewTrafficEventClassesTableHttpTran = 1
	SlbNewTrafficEventClassesTableHttpTran_Disabled    SlbNewTrafficEventClassesTableHttpTran = 2
	SlbNewTrafficEventClassesTableHttpTran_Unsupported SlbNewTrafficEventClassesTableHttpTran = 2147483647
)

type SlbNewTrafficEventClassesTablePathCorl int32

const (
	SlbNewTrafficEventClassesTablePathCorl_Disabled    SlbNewTrafficEventClassesTablePathCorl = 1
	SlbNewTrafficEventClassesTablePathCorl_Entry       SlbNewTrafficEventClassesTablePathCorl = 2
	SlbNewTrafficEventClassesTablePathCorl_Exit        SlbNewTrafficEventClassesTablePathCorl = 3
	SlbNewTrafficEventClassesTablePathCorl_Unsupported SlbNewTrafficEventClassesTablePathCorl = 2147483647
)

type SlbNewTrafficEventClassesTableSslConn int32

const (
	SlbNewTrafficEventClassesTableSslConn_Disabled    SlbNewTrafficEventClassesTableSslConn = 1
	SlbNewTrafficEventClassesTableSslConn_Frontend    SlbNewTrafficEventClassesTableSslConn = 2
	SlbNewTrafficEventClassesTableSslConn_Backend     SlbNewTrafficEventClassesTableSslConn = 3
	SlbNewTrafficEventClassesTableSslConn_Both        SlbNewTrafficEventClassesTableSslConn = 4
	SlbNewTrafficEventClassesTableSslConn_Unsupported SlbNewTrafficEventClassesTableSslConn = 2147483647
)

type SlbNewTrafficEventClassesTableSslFail int32

const (
	SlbNewTrafficEventClassesTableSslFail_Disabled    SlbNewTrafficEventClassesTableSslFail = 1
	SlbNewTrafficEventClassesTableSslFail_Frontend    SlbNewTrafficEventClassesTableSslFail = 2
	SlbNewTrafficEventClassesTableSslFail_Backend     SlbNewTrafficEventClassesTableSslFail = 3
	SlbNewTrafficEventClassesTableSslFail_Both        SlbNewTrafficEventClassesTableSslFail = 4
	SlbNewTrafficEventClassesTableSslFail_Unsupported SlbNewTrafficEventClassesTableSslFail = 2147483647
)

type SlbNewTrafficEventClassesTableHostByPass int32

const (
	SlbNewTrafficEventClassesTableHostByPass_Enabled     SlbNewTrafficEventClassesTableHostByPass = 1
	SlbNewTrafficEventClassesTableHostByPass_Disabled    SlbNewTrafficEventClassesTableHostByPass = 2
	SlbNewTrafficEventClassesTableHostByPass_Unsupported SlbNewTrafficEventClassesTableHostByPass = 2147483647
)

type SlbNewTrafficEventClassesTableL4Conn int32

const (
	SlbNewTrafficEventClassesTableL4Conn_Enabled     SlbNewTrafficEventClassesTableL4Conn = 1
	SlbNewTrafficEventClassesTableL4Conn_Disabled    SlbNewTrafficEventClassesTableL4Conn = 2
	SlbNewTrafficEventClassesTableL4Conn_Unsupported SlbNewTrafficEventClassesTableL4Conn = 2147483647
)

type SlbNewTrafficEventClassesTableDel int32

const (
	SlbNewTrafficEventClassesTableDel_Other       SlbNewTrafficEventClassesTableDel = 1
	SlbNewTrafficEventClassesTableDel_Delete      SlbNewTrafficEventClassesTableDel = 2
	SlbNewTrafficEventClassesTableDel_Unsupported SlbNewTrafficEventClassesTableDel = 2147483647
)

type SlbNewTrafficEventClassesTableUnified int32

const (
	SlbNewTrafficEventClassesTableUnified_Enabled     SlbNewTrafficEventClassesTableUnified = 1
	SlbNewTrafficEventClassesTableUnified_Disabled    SlbNewTrafficEventClassesTableUnified = 2
	SlbNewTrafficEventClassesTableUnified_Unsupported SlbNewTrafficEventClassesTableUnified = 2147483647
)

type SlbNewTrafficEventClassesTableSecurity int32

const (
	SlbNewTrafficEventClassesTableSecurity_Enabled     SlbNewTrafficEventClassesTableSecurity = 1
	SlbNewTrafficEventClassesTableSecurity_Disabled    SlbNewTrafficEventClassesTableSecurity = 2
	SlbNewTrafficEventClassesTableSecurity_Unsupported SlbNewTrafficEventClassesTableSecurity = 2147483647
)

type SlbNewTrafficEventClassesTableIPREP int32

const (
	SlbNewTrafficEventClassesTableIPREP_Enabled     SlbNewTrafficEventClassesTableIPREP = 1
	SlbNewTrafficEventClassesTableIPREP_Disabled    SlbNewTrafficEventClassesTableIPREP = 2
	SlbNewTrafficEventClassesTableIPREP_Unsupported SlbNewTrafficEventClassesTableIPREP = 2147483647
)

type SlbNewTrafficEventClassesTableParams struct {
	// Set traffic event class id.
	Id string `json:"Id,omitempty"`
	// Set traffic event class name.
	Name string `json:"Name,omitempty"`
	// Set traffic event group name.
	GroupName string `json:"GroupName,omitempty"`
	// Set traffic event log policy syslog server port.
	SysPort uint64 `json:"SysPort,omitempty"`
	// Set traffic event log policy sampling rate.
	SampleRate uint32 `json:"SampleRate,omitempty"`
	// Set traffic event policy status either ena or dis.
	EnaDis SlbNewTrafficEventClassesTableEnaDis `json:"EnaDis,omitempty"`
	// Set http transaction status for the traffic event policy.
	HttpTran SlbNewTrafficEventClassesTableHttpTran `json:"HttpTran,omitempty"`
	// Set http transaction path correlation for the traffic event policy.
	PathCorl SlbNewTrafficEventClassesTablePathCorl `json:"PathCorl,omitempty"`
	// Set ssl connection status for traffic event policy.
	SslConn SlbNewTrafficEventClassesTableSslConn `json:"SslConn,omitempty"`
	// Set ssl connection fail status for traffic event policy.
	SslFail SlbNewTrafficEventClassesTableSslFail `json:"SslFail,omitempty"`
	// set hotbypass status either ena or dis for the traffic event policy.
	HostByPass SlbNewTrafficEventClassesTableHostByPass `json:"HostByPass,omitempty"`
	// Set l4 connection status either ena or dis for the traffic event policy.
	L4Conn SlbNewTrafficEventClassesTableL4Conn `json:"L4Conn,omitempty"`
	// By setting the value to delete(2), the entire row is deleted.
	Del SlbNewTrafficEventClassesTableDel `json:"Del,omitempty"`
	// Set Remote logging object
	RemoteLogId string `json:"RemoteLogId,omitempty"`
	// Set unified event status either ena or dis for the traffic event policy.
	Unified SlbNewTrafficEventClassesTableUnified `json:"Unified,omitempty"`
	// Set traffic event log policy Normal EPS limit: 1-999999 Events per second limit, disable or 0, Unlimited or 1000000
	NormalEps uint32 `json:"NormalEps,omitempty"`
	// Set traffic event log policy Exception EPS limit: 1-999999 Events per second limit, disable or 0, Unlimited or 1000000
	ExceptionEps uint32 `json:"ExceptionEps,omitempty"`
	// Set security event status either ena or dis for the traffic event policy.
	Security SlbNewTrafficEventClassesTableSecurity `json:"Security,omitempty"`
	// Traffic event policy serialization format CEF or JSON.
	Format string `json:"Format,omitempty"`
	// Set EAAF status either ena or dis for the traffic event policy.
	IPREP SlbNewTrafficEventClassesTableIPREP `json:"IPREP,omitempty"`
}

func (p SlbNewTrafficEventClassesTableParams) iMABean() {}
