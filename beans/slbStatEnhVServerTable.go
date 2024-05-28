package beans

import (
	"fmt"
	"reflect"
)

// SlbStatEnhVServerTable The virtual server statistics table.
// Note:This mib is not supported for VX instance of Virtualization.
type SlbStatEnhVServerTable struct {
	// The virtual server number that identifies the server.
	SlbStatEnhVServerIndex string
	Params                 *SlbStatEnhVServerTableParams
}

func NewSlbStatEnhVServerTableList() *SlbStatEnhVServerTable {
	return &SlbStatEnhVServerTable{}
}

func NewSlbStatEnhVServerTable(
	slbStatEnhVServerIndex string,
	params *SlbStatEnhVServerTableParams,
) *SlbStatEnhVServerTable {
	return &SlbStatEnhVServerTable{
		SlbStatEnhVServerIndex: slbStatEnhVServerIndex,
		Params:                 params,
	}
}

func (c *SlbStatEnhVServerTable) Name() string {
	return "SlbStatEnhVServerTable"
}

func (c *SlbStatEnhVServerTable) GetParams() BeanType {
	return c.Params
}

func (c *SlbStatEnhVServerTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *SlbStatEnhVServerTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.SlbStatEnhVServerIndex).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.SlbStatEnhVServerIndex)
}

type SlbStatEnhVServerTableParams struct {
	// The virtual server number that identifies the server.
	Index string `json:"Index,omitempty"`
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
	// Number of Sessions handled by the virtual server per second.
	SessionsPerSec uint32 `json:"SessionsPerSec,omitempty"`
	// Number of Mbps received and transmitted by the virtual server per second.
	OctetsPerSec string `json:"OctetsPerSec,omitempty"`
	// The number of persistent sessions that are currently handled by the
	// virtual server.
	CurrPSessions uint64 `json:"CurrPSessions,omitempty"`
	// The total number of persistent sessions that are handled by the virtual server.
	TotalPSessions uint64 `json:"TotalPSessions,omitempty"`
	// The highest persistent sessions that have been handled by the virtual server.
	HighestPSessions uint64 `json:"HighestPSessions,omitempty"`
	// Number of Packets handled by the virtual server per second.
	PacketsPerSec uint32 `json:"PacketsPerSec,omitempty"`
}

func (p SlbStatEnhVServerTableParams) iMABean() {}
