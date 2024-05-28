package beans

import (
	"fmt"
	"reflect"
)

// SlbStatLinkpfIpTable The link proof real server statistics table.
// Note:This mib is not supported for VX instance of Virtualization.
type SlbStatLinkpfIpTable struct {
	// The WAN Link real server IP that identifies the server.
	SlbStatLinkpfIpIndex string
	Params               *SlbStatLinkpfIpTableParams
}

func NewSlbStatLinkpfIpTableList() *SlbStatLinkpfIpTable {
	return &SlbStatLinkpfIpTable{}
}

func NewSlbStatLinkpfIpTable(
	slbStatLinkpfIpIndex string,
	params *SlbStatLinkpfIpTableParams,
) *SlbStatLinkpfIpTable {
	return &SlbStatLinkpfIpTable{
		SlbStatLinkpfIpIndex: slbStatLinkpfIpIndex,
		Params:               params,
	}
}

func (c *SlbStatLinkpfIpTable) Name() string {
	return "SlbStatLinkpfIpTable"
}

func (c *SlbStatLinkpfIpTable) GetParams() BeanType {
	return c.Params
}

func (c *SlbStatLinkpfIpTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *SlbStatLinkpfIpTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.SlbStatLinkpfIpIndex).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.SlbStatLinkpfIpIndex)
}

type SlbStatLinkpfIpTableParams struct {
	// The WAN Link real server IP that identifies the server.
	Index string `json:"Index,omitempty"`
	// The total number of current upload bandwidth by WAN Link real servers with same IP.
	UpBwCurr string `json:"UpBwCurr,omitempty"`
	// The total number of current upload bandwidth Usages by WAN Link real servers with same IP.
	UpBwCurrUsage string `json:"UpBwCurrUsage,omitempty"`
	// The total number of current download bandwidth by WAN Link real servers with same IP.
	DnBwCurr string `json:"DnBwCurr,omitempty"`
	// The total number of current upload bandwidth Usages by WAN Link real servers with same IP.
	DnBwCurrUsage string `json:"DnBwCurrUsage,omitempty"`
	// The total number of current upload and download bandwidth by WAN Link real servers with same IP.
	TotBwCurr string `json:"TotBwCurr,omitempty"`
	// The total number of current upload bandwidth Usages by WAN Link real servers with same IP.
	TotBwCurrUsage string `json:"TotBwCurrUsage,omitempty"`
	// The total number of current sessions that are handled by WAN Link real servers with same IP.
	CurrSessions uint32 `json:"CurrSessions,omitempty"`
	// The total number of current upload bandwidth by WAN Link real servers with same IP.
	UpBwPeak string `json:"UpBwPeak,omitempty"`
	// The total number of current upload bandwidth Usages by WAN Link real servers with same IP.
	UpBwPeakPer string `json:"UpBwPeakPer,omitempty"`
	// The total number of current upload bandwidth peak time stamp by WAN Link real server with same IP.
	UpBwPeakTmSt string `json:"UpBwPeakTmSt,omitempty"`
	// The total number of current download bandwidth by WAN Link real servers with same IP.
	DnBwPeak string `json:"DnBwPeak,omitempty"`
	// The total number of current upload bandwidth Usages by WAN Link real servers with same IP.
	DnBwPeakPer string `json:"DnBwPeakPer,omitempty"`
	// The total number of current upload bandwidth peak time stamp by WAN Link real server with same IP.
	DnBwPeakTmSt string `json:"DnBwPeakTmSt,omitempty"`
	// The total number of current download and download bandwidth by WAN Link real servers with same IP.
	TotBwPeak string `json:"TotBwPeak,omitempty"`
	// The total number of current upload and download bandwidth Usages by WAN Link real servers with same IP.
	TotBwPeakPer string `json:"TotBwPeakPer,omitempty"`
	// The total number of current upload and download bandwidth peak time stamp by WAN Link real server with same IP.
	TotBwPeakTmSt string `json:"TotBwPeakTmSt,omitempty"`
	// Last clear time stamp by WAN Link real server with same IP.
	LastClearTmSt string `json:"LastClearTmSt,omitempty"`
	// The total number of upload bandwidth by WAN Link real servers with same IP.
	UpBwTot string `json:"UpBwTot,omitempty"`
	// The total number of download bandwidth by WAN Link real servers with same IP.
	DnBwTot string `json:"DnBwTot,omitempty"`
	// The total number of upload and download bandwidth by WAN Link real servers with same IP.
	UpDnBwTot string `json:"UpDnBwTot,omitempty"`
}

func (p SlbStatLinkpfIpTableParams) iMABean() {}
