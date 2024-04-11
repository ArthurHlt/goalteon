package beans

import (
	"fmt"
	"reflect"
)

// SlbAdvhcDsspTable Note:This mib is not supported on VX instance of Virtualization.
type SlbAdvhcDsspTable struct {
	// .
	SlbAdvhcDsspID string
	Params         *SlbAdvhcDsspTableParams
}

func NewSlbAdvhcDsspTableList() *SlbAdvhcDsspTable {
	return &SlbAdvhcDsspTable{}
}

func NewSlbAdvhcDsspTable(
	slbAdvhcDsspID string,
	params *SlbAdvhcDsspTableParams,
) *SlbAdvhcDsspTable {
	return &SlbAdvhcDsspTable{
		SlbAdvhcDsspID: slbAdvhcDsspID,
		Params:         params,
	}
}

func (c *SlbAdvhcDsspTable) Name() string {
	return "SlbAdvhcDsspTable"
}

func (c *SlbAdvhcDsspTable) GetParams() BeanType {
	return c.Params
}

func (c *SlbAdvhcDsspTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *SlbAdvhcDsspTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.SlbAdvhcDsspID).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.SlbAdvhcDsspID)
}

type SlbAdvhcDsspTableIPVer int32

const (
	SlbAdvhcDsspTableIPVer_Ipv4 SlbAdvhcDsspTableIPVer = 1
	SlbAdvhcDsspTableIPVer_Ipv6 SlbAdvhcDsspTableIPVer = 2
	SlbAdvhcDsspTableIPVer_None SlbAdvhcDsspTableIPVer = 3
)

type SlbAdvhcDsspTableTransparent int32

const (
	SlbAdvhcDsspTableTransparent_Enabled     SlbAdvhcDsspTableTransparent = 1
	SlbAdvhcDsspTableTransparent_Disabled    SlbAdvhcDsspTableTransparent = 2
	SlbAdvhcDsspTableTransparent_Unsupported SlbAdvhcDsspTableTransparent = 2147483647
)

type SlbAdvhcDsspTableParams struct {
	// .
	ID string `json:"ID,omitempty"`
	// .
	Name string `json:"Name,omitempty"`
	// .
	DPort uint64 `json:"DPort,omitempty"`
	// .
	IPVer SlbAdvhcDsspTableIPVer `json:"IPVer,omitempty"`
	// .
	HostName string `json:"HostName,omitempty"`
	// .
	Transparent SlbAdvhcDsspTableTransparent `json:"Transparent,omitempty"`
	// .
	Interval uint64 `json:"Interval,omitempty"`
	// .
	Retries uint64 `json:"Retries,omitempty"`
	// .
	RestoreRetries uint64 `json:"RestoreRetries,omitempty"`
	// .
	Timeout uint64 `json:"Timeout,omitempty"`
}

func (p SlbAdvhcDsspTableParams) iMABean() {}
