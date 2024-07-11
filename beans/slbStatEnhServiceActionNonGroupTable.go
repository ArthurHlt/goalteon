package beans

import (
	"fmt"
	"reflect"
)

// SlbStatEnhServiceActionNonGroupTable The statistics table of instances for service action non group.
// Note:This mib is not supported for VX instance of virtualization.
type SlbStatEnhServiceActionNonGroupTable struct {
	// The virtual server id.
	SlbStatEnhServiceActionNonGroupVirtServIndex string
	// Virtual service index.
	SlbStatEnhServiceActionNonGroupVirtServiceIndex int32
	Params                                          *SlbStatEnhServiceActionNonGroupTableParams
}

func NewSlbStatEnhServiceActionNonGroupTableList() *SlbStatEnhServiceActionNonGroupTable {
	return &SlbStatEnhServiceActionNonGroupTable{}
}

func NewSlbStatEnhServiceActionNonGroupTable(
	slbStatEnhServiceActionNonGroupVirtServIndex string,
	slbStatEnhServiceActionNonGroupVirtServiceIndex int32,
	params *SlbStatEnhServiceActionNonGroupTableParams,
) *SlbStatEnhServiceActionNonGroupTable {
	return &SlbStatEnhServiceActionNonGroupTable{
		SlbStatEnhServiceActionNonGroupVirtServIndex:    slbStatEnhServiceActionNonGroupVirtServIndex,
		SlbStatEnhServiceActionNonGroupVirtServiceIndex: slbStatEnhServiceActionNonGroupVirtServiceIndex,
		Params: params,
	}
}

func (c *SlbStatEnhServiceActionNonGroupTable) Name() string {
	return "SlbStatEnhServiceActionNonGroupTable"
}

func (c *SlbStatEnhServiceActionNonGroupTable) GetParams() BeanType {
	return c.Params
}

func (c *SlbStatEnhServiceActionNonGroupTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *SlbStatEnhServiceActionNonGroupTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.SlbStatEnhServiceActionNonGroupVirtServIndex).IsZero() &&
		reflect.ValueOf(c.SlbStatEnhServiceActionNonGroupVirtServiceIndex).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.SlbStatEnhServiceActionNonGroupVirtServIndex) + "/" + fmt.Sprint(c.SlbStatEnhServiceActionNonGroupVirtServiceIndex)
}

type SlbStatEnhServiceActionNonGroupTableParams struct {
	// The virtual server id.
	VirtServIndex string `json:"VirtServIndex,omitempty"`
	// Virtual service index.
	VirtServiceIndex int32 `json:"VirtServiceIndex,omitempty"`
	// The service non group current sessions.
	CurrSess int32 `json:"CurrSess,omitempty"`
	// The service non group total sessions.
	TotSess int32 `json:"TotSess,omitempty"`
	// The service non group highest sessions.
	HighSess int32 `json:"HighSess,omitempty"`
	// The service non group total octets.
	TotOcts int32 `json:"TotOcts,omitempty"`
}

func (p SlbStatEnhServiceActionNonGroupTableParams) iMABean() {}
