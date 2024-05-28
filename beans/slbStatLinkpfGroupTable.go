package beans

import (
	"fmt"
	"reflect"
)

// SlbStatLinkpfGroupTable The link proof real server group statistics table.
// Note:This mib is not supported for VX instance of Virtualization.
type SlbStatLinkpfGroupTable struct {
	// The WAN Link real server group index that identifies the group.
	SlbStatLinkpfGroupIndex string
	Params                  *SlbStatLinkpfGroupTableParams
}

func NewSlbStatLinkpfGroupTableList() *SlbStatLinkpfGroupTable {
	return &SlbStatLinkpfGroupTable{}
}

func NewSlbStatLinkpfGroupTable(
	slbStatLinkpfGroupIndex string,
	params *SlbStatLinkpfGroupTableParams,
) *SlbStatLinkpfGroupTable {
	return &SlbStatLinkpfGroupTable{
		SlbStatLinkpfGroupIndex: slbStatLinkpfGroupIndex,
		Params:                  params,
	}
}

func (c *SlbStatLinkpfGroupTable) Name() string {
	return "SlbStatLinkpfGroupTable"
}

func (c *SlbStatLinkpfGroupTable) GetParams() BeanType {
	return c.Params
}

func (c *SlbStatLinkpfGroupTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *SlbStatLinkpfGroupTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.SlbStatLinkpfGroupIndex).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.SlbStatLinkpfGroupIndex)
}

type SlbStatLinkpfGroupTableParams struct {
	// The WAN Link real server group index that identifies the group.
	Index string `json:"Index,omitempty"`
	// The total current upload bandwidth by WAN Link group for all WAN Link real server.
	CurrUpBw string `json:"CurrUpBw,omitempty"`
	// The total current download bandwidth by WAN Link group for all WAN Link real server.
	CurrDnBw string `json:"CurrDnBw,omitempty"`
	// The total upload and download bandwidth by WAN Link group for all WAN Link real server.
	CurrTotBw string `json:"CurrTotBw,omitempty"`
	// The number of sessions that are currently handled by the WAN Link group.
	CurrSess uint32 `json:"CurrSess,omitempty"`
	// Last reset or clear time stamp for this group.
	LastClearTmSt string `json:"LastClearTmSt,omitempty"`
	// The total upload bandwidth that are handled by WAN Link group for all WAN Link real server.
	TotUpBw string `json:"TotUpBw,omitempty"`
	// The total download bandwidth that are handled by WAN Link group for all WAN Link real server.
	TotDnBw string `json:"TotDnBw,omitempty"`
	// The total upload and download bandwidth that are handled by WAN Link group for all WAN Link real server.
	TotBw string `json:"TotBw,omitempty"`
}

func (p SlbStatLinkpfGroupTableParams) iMABean() {}
