package beans

import (
	"fmt"
	"reflect"
)

// CurWafSyslogServerCfgTable Cur Syslog servers table.
// Note:This mib is not supported on VX instance of Virtualization.
type CurWafSyslogServerCfgTable struct {
	// Syslog server ID
	CurWafSyslogServerId string
	Params               *CurWafSyslogServerCfgTableParams
}

func NewCurWafSyslogServerCfgTableList() *CurWafSyslogServerCfgTable {
	return &CurWafSyslogServerCfgTable{}
}

func NewCurWafSyslogServerCfgTable(
	curWafSyslogServerId string,
	params *CurWafSyslogServerCfgTableParams,
) *CurWafSyslogServerCfgTable {
	return &CurWafSyslogServerCfgTable{
		CurWafSyslogServerId: curWafSyslogServerId,
		Params:               params,
	}
}

func (c *CurWafSyslogServerCfgTable) Name() string {
	return "CurWafSyslogServerCfgTable"
}

func (c *CurWafSyslogServerCfgTable) GetParams() BeanType {
	return c.Params
}

func (c *CurWafSyslogServerCfgTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *CurWafSyslogServerCfgTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.CurWafSyslogServerId).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.CurWafSyslogServerId)
}

type CurWafSyslogServerCfgTableDel int32

const (
	CurWafSyslogServerCfgTableDel_Other       CurWafSyslogServerCfgTableDel = 1
	CurWafSyslogServerCfgTableDel_Delete      CurWafSyslogServerCfgTableDel = 2
	CurWafSyslogServerCfgTableDel_Unsupported CurWafSyslogServerCfgTableDel = 2147483647
)

type CurWafSyslogServerCfgTableParams struct {
	// Syslog server ID
	Id string `json:"Id,omitempty"`
	// Syslog server IP
	IpAddress string `json:"IpAddress,omitempty"`
	// Syslog server Port
	Port uint64 `json:"Port,omitempty"`
	// By setting the value to delete(2), the entire row is deleted.On GET always other(1)
	Del CurWafSyslogServerCfgTableDel `json:"Del,omitempty"`
}

func (p CurWafSyslogServerCfgTableParams) iMABean() {}
