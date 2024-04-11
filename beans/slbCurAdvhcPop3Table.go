package beans

import (
	"fmt"
	"reflect"
)

// SlbCurAdvhcPop3Table Note:This mib is not supported on VX instance of Virtualization.
type SlbCurAdvhcPop3Table struct {
	// POP3 Health check id.
	SlbCurAdvhcPop3ID string
	Params            *SlbCurAdvhcPop3TableParams
}

func NewSlbCurAdvhcPop3TableList() *SlbCurAdvhcPop3Table {
	return &SlbCurAdvhcPop3Table{}
}

func NewSlbCurAdvhcPop3Table(
	slbCurAdvhcPop3ID string,
	params *SlbCurAdvhcPop3TableParams,
) *SlbCurAdvhcPop3Table {
	return &SlbCurAdvhcPop3Table{
		SlbCurAdvhcPop3ID: slbCurAdvhcPop3ID,
		Params:            params,
	}
}

func (c *SlbCurAdvhcPop3Table) Name() string {
	return "SlbCurAdvhcPop3Table"
}

func (c *SlbCurAdvhcPop3Table) GetParams() BeanType {
	return c.Params
}

func (c *SlbCurAdvhcPop3Table) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *SlbCurAdvhcPop3Table) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.SlbCurAdvhcPop3ID).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.SlbCurAdvhcPop3ID)
}

type SlbCurAdvhcPop3TableIPVer int32

const (
	SlbCurAdvhcPop3TableIPVer_Ipv4 SlbCurAdvhcPop3TableIPVer = 1
	SlbCurAdvhcPop3TableIPVer_Ipv6 SlbCurAdvhcPop3TableIPVer = 2
	SlbCurAdvhcPop3TableIPVer_None SlbCurAdvhcPop3TableIPVer = 3
)

type SlbCurAdvhcPop3TableTransparent int32

const (
	SlbCurAdvhcPop3TableTransparent_Enabled     SlbCurAdvhcPop3TableTransparent = 1
	SlbCurAdvhcPop3TableTransparent_Disabled    SlbCurAdvhcPop3TableTransparent = 2
	SlbCurAdvhcPop3TableTransparent_Unsupported SlbCurAdvhcPop3TableTransparent = 2147483647
)

type SlbCurAdvhcPop3TableInvert int32

const (
	SlbCurAdvhcPop3TableInvert_Enabled     SlbCurAdvhcPop3TableInvert = 1
	SlbCurAdvhcPop3TableInvert_Disabled    SlbCurAdvhcPop3TableInvert = 2
	SlbCurAdvhcPop3TableInvert_Unsupported SlbCurAdvhcPop3TableInvert = 2147483647
)

type SlbCurAdvhcPop3TableParams struct {
	// POP3 Health check id.
	ID string `json:"ID,omitempty"`
	// POP3 Health check name.
	Name string `json:"Name,omitempty"`
	// POP3 Health check destination port.
	DPort uint64 `json:"DPort,omitempty"`
	// POP3 Health check destination IP version.
	IPVer SlbCurAdvhcPop3TableIPVer `json:"IPVer,omitempty"`
	// POP3 Health check destination hostname.
	HostName string `json:"HostName,omitempty"`
	// POP3 Health check trasnsparent flag.
	Transparent SlbCurAdvhcPop3TableTransparent `json:"Transparent,omitempty"`
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
	Invert SlbCurAdvhcPop3TableInvert `json:"Invert,omitempty"`
	// POP3 Health check user name.
	UserName string `json:"UserName,omitempty"`
	// POP3 Health check password.
	Password string `json:"Password,omitempty"`
}

func (p SlbCurAdvhcPop3TableParams) iMABean() {}
