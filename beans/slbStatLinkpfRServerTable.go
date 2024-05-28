package beans

import (
	"fmt"
	"reflect"
)

// SlbStatLinkpfRServerTable The link proof real server statistics table.
// Note:This mib is not supported for VX instance of Virtualization.
type SlbStatLinkpfRServerTable struct {
	// The WAN Link real server index that identifies the server.
	SlbStatLinkpfRServerIndex string
	Params                    *SlbStatLinkpfRServerTableParams
}

func NewSlbStatLinkpfRServerTableList() *SlbStatLinkpfRServerTable {
	return &SlbStatLinkpfRServerTable{}
}

func NewSlbStatLinkpfRServerTable(
	slbStatLinkpfRServerIndex string,
	params *SlbStatLinkpfRServerTableParams,
) *SlbStatLinkpfRServerTable {
	return &SlbStatLinkpfRServerTable{
		SlbStatLinkpfRServerIndex: slbStatLinkpfRServerIndex,
		Params:                    params,
	}
}

func (c *SlbStatLinkpfRServerTable) Name() string {
	return "SlbStatLinkpfRServerTable"
}

func (c *SlbStatLinkpfRServerTable) GetParams() BeanType {
	return c.Params
}

func (c *SlbStatLinkpfRServerTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *SlbStatLinkpfRServerTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.SlbStatLinkpfRServerIndex).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.SlbStatLinkpfRServerIndex)
}

type SlbStatLinkpfRServerTableState int32

const (
	SlbStatLinkpfRServerTableState_Running     SlbStatLinkpfRServerTableState = 1
	SlbStatLinkpfRServerTableState_Failed      SlbStatLinkpfRServerTableState = 2
	SlbStatLinkpfRServerTableState_Disabled    SlbStatLinkpfRServerTableState = 3
	SlbStatLinkpfRServerTableState_Unsupported SlbStatLinkpfRServerTableState = 2147483647
)

type SlbStatLinkpfRServerTableParams struct {
	// The WAN Link real server index that identifies the server.
	Index string `json:"Index,omitempty"`
	// IP address of the WAN Link real server identified by the instance of the
	// slbStatLinkpfRServerIndex.
	IpAddr string `json:"IpAddr,omitempty"`
	// The total number of current upload bandwidth by WAN Link real server.
	UpBwCurr string `json:"UpBwCurr,omitempty"`
	// The Usage of upload bandwidth by WAN Link real server.
	UpBwUsage string `json:"UpBwUsage,omitempty"`
	// The total number of current download bandwidth by WAN Link real server.
	DwBwCurr string `json:"DwBwCurr,omitempty"`
	// The Usage of download bandwidth by WAN Link real server.
	DwBwUSage string `json:"DwBwUSage,omitempty"`
	// The total number of upload and download bandwidth by WAN Link real server.
	TotCurrbw string `json:"TotCurrbw,omitempty"`
	// The Total Usage of upload and Download bandwidth by WAN Link real server.
	TotCurrUsage string `json:"TotCurrUsage,omitempty"`
	// The current number of sessions by WAN Link real server.
	CurrSess uint32 `json:"CurrSess,omitempty"`
	// Upload peak bandwidth by WAN Link real server.
	UpBwPeak string `json:"UpBwPeak,omitempty"`
	// Upload per peak bandwidth by WAN Link real server.
	UpBwPeakPer string `json:"UpBwPeakPer,omitempty"`
	// Upload peak bandwidth time stamp by WAN Link real server.
	UpBwPeakTmSt string `json:"UpBwPeakTmSt,omitempty"`
	// Download peak bandwidth by WAN Link real server.
	DnBwPeak string `json:"DnBwPeak,omitempty"`
	// Download per peak bandwidth by WAN Link real server.
	DnBwPeakPer string `json:"DnBwPeakPer,omitempty"`
	// Download peak bandwidth time stamp by WAN Link real server.
	DnBwPeakTmSt string `json:"DnBwPeakTmSt,omitempty"`
	// Total peak bandwidth by WAN Link real server.
	TotBwPeak string `json:"TotBwPeak,omitempty"`
	// Total per peak bandwidth by WAN Link real server.
	TotBwPeakPer string `json:"TotBwPeakPer,omitempty"`
	// Total bandwidth peak time stamp of the WAN link real server.
	TotBwPeakTmSt string `json:"TotBwPeakTmSt,omitempty"`
	// Last clear time stamp of the WAN Link real server .
	LastTranfetTmSt string `json:"LastTranfetTmSt,omitempty"`
	// The total upload bandwidth by WAN Link real server.
	UpBwTot string `json:"UpBwTot,omitempty"`
	// The total download bandwidth by WAN Link real server.
	DnBwTot string `json:"DnBwTot,omitempty"`
	// The total upload and download Bandwidth by WAN Link real server.
	UpDnBwTot string `json:"UpDnBwTot,omitempty"`
	// The state of the wan link real server.
	State SlbStatLinkpfRServerTableState `json:"State,omitempty"`
}

func (p SlbStatLinkpfRServerTableParams) iMABean() {}
