package beans

import (
	"fmt"
	"reflect"
)

// SlbNewAdvhcPop3Table Note:This mib is not supported on VX instance of Virtualization.
type SlbNewAdvhcPop3Table struct {
	// POP3 Health check id.
	SlbNewAdvhcPop3ID string
	Params            *SlbNewAdvhcPop3TableParams
}

func NewSlbNewAdvhcPop3TableList() *SlbNewAdvhcPop3Table {
	return &SlbNewAdvhcPop3Table{}
}

func NewSlbNewAdvhcPop3Table(
	slbNewAdvhcPop3ID string,
	params *SlbNewAdvhcPop3TableParams,
) *SlbNewAdvhcPop3Table {
	return &SlbNewAdvhcPop3Table{
		SlbNewAdvhcPop3ID: slbNewAdvhcPop3ID,
		Params:            params,
	}
}

func (c *SlbNewAdvhcPop3Table) Name() string {
	return "SlbNewAdvhcPop3Table"
}

func (c *SlbNewAdvhcPop3Table) GetParams() BeanType {
	return c.Params
}

func (c *SlbNewAdvhcPop3Table) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *SlbNewAdvhcPop3Table) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.SlbNewAdvhcPop3ID).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.SlbNewAdvhcPop3ID)
}

type SlbNewAdvhcPop3TableIPVer int32

const (
	SlbNewAdvhcPop3TableIPVer_Ipv4 SlbNewAdvhcPop3TableIPVer = 1
	SlbNewAdvhcPop3TableIPVer_Ipv6 SlbNewAdvhcPop3TableIPVer = 2
	SlbNewAdvhcPop3TableIPVer_None SlbNewAdvhcPop3TableIPVer = 3
)

type SlbNewAdvhcPop3TableTransparent int32

const (
	SlbNewAdvhcPop3TableTransparent_Enabled     SlbNewAdvhcPop3TableTransparent = 1
	SlbNewAdvhcPop3TableTransparent_Disabled    SlbNewAdvhcPop3TableTransparent = 2
	SlbNewAdvhcPop3TableTransparent_Unsupported SlbNewAdvhcPop3TableTransparent = 2147483647
)

type SlbNewAdvhcPop3TableInvert int32

const (
	SlbNewAdvhcPop3TableInvert_Enabled     SlbNewAdvhcPop3TableInvert = 1
	SlbNewAdvhcPop3TableInvert_Disabled    SlbNewAdvhcPop3TableInvert = 2
	SlbNewAdvhcPop3TableInvert_Unsupported SlbNewAdvhcPop3TableInvert = 2147483647
)

type SlbNewAdvhcPop3TableDelete int32

const (
	SlbNewAdvhcPop3TableDelete_Other       SlbNewAdvhcPop3TableDelete = 1
	SlbNewAdvhcPop3TableDelete_Delete      SlbNewAdvhcPop3TableDelete = 2
	SlbNewAdvhcPop3TableDelete_Unsupported SlbNewAdvhcPop3TableDelete = 2147483647
)

type SlbNewAdvhcPop3TableSnat int32

const (
	SlbNewAdvhcPop3TableSnat_Enabled     SlbNewAdvhcPop3TableSnat = 1
	SlbNewAdvhcPop3TableSnat_Disabled    SlbNewAdvhcPop3TableSnat = 2
	SlbNewAdvhcPop3TableSnat_Unsupported SlbNewAdvhcPop3TableSnat = 2147483647
)

type SlbNewAdvhcPop3TableParams struct {
	// POP3 Health check id.
	ID string `json:"ID,omitempty"`
	// POP3 Health check name.
	Name string `json:"Name,omitempty"`
	// POP3 Health check destination port.
	DPort uint64 `json:"DPort,omitempty"`
	// POP3 Health check destination IP version.
	IPVer SlbNewAdvhcPop3TableIPVer `json:"IPVer,omitempty"`
	// POP3 Health check destination hostname.
	HostName string `json:"HostName,omitempty"`
	// POP3 Health check trasnsparent flag.
	Transparent SlbNewAdvhcPop3TableTransparent `json:"Transparent,omitempty"`
	// POP3 Health check interval.
	Interval uint64 `json:"Interval,omitempty"`
	// POP3 Health check retries counter.
	Retries uint64 `json:"Retries,omitempty"`
	// POP3 Health check retries counter when in the down state.
	RestoreRetries uint64 `json:"RestoreRetries,omitempty"`
	// POP3 Health check timeout.
	Timeout uint64 `json:"Timeout,omitempty"`
	// POP3 Health check overload flag.
	Overflow uint64 `json:"Overflow,omitempty"`
	// POP3 Health check interval when in the down state.
	DownInterval uint64 `json:"DownInterval,omitempty"`
	// POP3 Health check invert flag.
	Invert SlbNewAdvhcPop3TableInvert `json:"Invert,omitempty"`
	// POP3 Health check user name.
	UserName string `json:"UserName,omitempty"`
	// POP3 Health check password.
	Password string `json:"Password,omitempty"`
	// POP3 Health check copy indicator.
	Copy string `json:"Copy,omitempty"`
	// When set to the value of 2 (delete), the entire row is deleted.
	// When read, other(1) is returned. Setting the value to anything
	// other than 2(delete) has no effect on the state of the row.
	Delete SlbNewAdvhcPop3TableDelete `json:"Delete,omitempty"`
	// POP3 Health Check src NAT (PIP) flag.
	Snat SlbNewAdvhcPop3TableSnat `json:"Snat,omitempty"`
}

func (p SlbNewAdvhcPop3TableParams) iMABean() {}
