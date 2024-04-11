package beans

import (
	"fmt"
	"reflect"
)

// NewWafSyslogServerCfgTable New Syslog servers table.
// Note:This mib is not supported on VX instance of Virtualization.
type NewWafSyslogServerCfgTable struct {
	// Syslog server ID
	NewWafSyslogServerId string
	Params               *NewWafSyslogServerCfgTableParams
}

func NewNewWafSyslogServerCfgTableList() *NewWafSyslogServerCfgTable {
	return &NewWafSyslogServerCfgTable{}
}

func NewNewWafSyslogServerCfgTable(
	newWafSyslogServerId string,
	params *NewWafSyslogServerCfgTableParams,
) *NewWafSyslogServerCfgTable {
	return &NewWafSyslogServerCfgTable{
		NewWafSyslogServerId: newWafSyslogServerId,
		Params:               params,
	}
}

func (c *NewWafSyslogServerCfgTable) Name() string {
	return "NewWafSyslogServerCfgTable"
}

func (c *NewWafSyslogServerCfgTable) GetParams() BeanType {
	return c.Params
}

func (c *NewWafSyslogServerCfgTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *NewWafSyslogServerCfgTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.NewWafSyslogServerId).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.NewWafSyslogServerId)
}

type NewWafSyslogServerCfgTableDel int32

const (
	NewWafSyslogServerCfgTableDel_Other       NewWafSyslogServerCfgTableDel = 1
	NewWafSyslogServerCfgTableDel_Delete      NewWafSyslogServerCfgTableDel = 2
	NewWafSyslogServerCfgTableDel_Unsupported NewWafSyslogServerCfgTableDel = 2147483647
)

type NewWafSyslogServerCfgTableParams struct {
	// Syslog server ID
	Id string `json:"Id,omitempty"`
	// Syslog server IP
	IpAddress string `json:"IpAddress,omitempty"`
	// Syslog server Port
	Port uint64 `json:"Port,omitempty"`
	// By setting the value to delete(2), the entire row is deleted.On GET always other(1)
	Del NewWafSyslogServerCfgTableDel `json:"Del,omitempty"`
}

func (p NewWafSyslogServerCfgTableParams) iMABean() {}
