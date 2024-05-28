package beans

import (
	"fmt"
	"reflect"
)

// SlbStatVServerTable The virtual server statistics table.
// Note:This mib is not supported for VX instance of Virtualization.
type SlbStatVServerTable struct {
	// The virtual server number that identifies the server.
	SlbStatVServerIndex int32
	Params              *SlbStatVServerTableParams
}

func NewSlbStatVServerTableList() *SlbStatVServerTable {
	return &SlbStatVServerTable{}
}

func NewSlbStatVServerTable(
	slbStatVServerIndex int32,
	params *SlbStatVServerTableParams,
) *SlbStatVServerTable {
	return &SlbStatVServerTable{
		SlbStatVServerIndex: slbStatVServerIndex,
		Params:              params,
	}
}

func (c *SlbStatVServerTable) Name() string {
	return "SlbStatVServerTable"
}

func (c *SlbStatVServerTable) GetParams() BeanType {
	return c.Params
}

func (c *SlbStatVServerTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *SlbStatVServerTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.SlbStatVServerIndex).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.SlbStatVServerIndex)
}

type SlbStatVServerTableParams struct {
	// The virtual server number that identifies the server.
	Index int32 `json:"Index,omitempty"`
	// The number of sessions that are currently handled by the
	// virtual server.
	CurrSessions uint32 `json:"CurrSessions,omitempty"`
	// The total number of sessions that are handled by the virtual server.
	TotalSessions uint32 `json:"TotalSessions,omitempty"`
	// The highest sessions that have been handled by the virtual server.
	HighestSessions uint32 `json:"HighestSessions,omitempty"`
	// The lower 32 bit value of octets received and transmitted out
	// of the virtual server.
	HCOctetsLow32 uint32 `json:"HCOctetsLow32,omitempty"`
	// The higher 32 bit value of octets received and transmitted out
	// of the virtual server.
	HCOctetsHigh32 uint32 `json:"HCOctetsHigh32,omitempty"`
	// The current HTTP header hits.
	HeaderHits uint32 `json:"HeaderHits,omitempty"`
	// The current HTTP header misses.
	HeaderMisses uint32 `json:"HeaderMisses,omitempty"`
	// The total HTTP header sessions.
	HeaderTotalSessions uint32 `json:"HeaderTotalSessions,omitempty"`
	// The total Cookie Rewrites.
	CookieRewrites uint32 `json:"CookieRewrites,omitempty"`
	// The total Cookie Inserts.
	CookieInserts uint32 `json:"CookieInserts,omitempty"`
	// The total octets received and transmitted by the virtual server.
	HCOctets uint64 `json:"HCOctets,omitempty"`
	// IP address of the virtual server.
	IpAddress string `json:"IpAddress,omitempty"`
}

func (p SlbStatVServerTableParams) iMABean() {}
