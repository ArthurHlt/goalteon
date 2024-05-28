package beans

import (
	"fmt"
	"reflect"
)

// SlbStatLinkpfSmartNATTable The link proof SmartNAT statistics table.
// Note:This mib is not supported for VX instance of Virtualization.
type SlbStatLinkpfSmartNATTable struct {
	// The SmartNAT index that identifies the SmartNAT entry.
	SlbStatLinkpfSmartNATIndex string
	Params                     *SlbStatLinkpfSmartNATTableParams
}

func NewSlbStatLinkpfSmartNATTableList() *SlbStatLinkpfSmartNATTable {
	return &SlbStatLinkpfSmartNATTable{}
}

func NewSlbStatLinkpfSmartNATTable(
	slbStatLinkpfSmartNATIndex string,
	params *SlbStatLinkpfSmartNATTableParams,
) *SlbStatLinkpfSmartNATTable {
	return &SlbStatLinkpfSmartNATTable{
		SlbStatLinkpfSmartNATIndex: slbStatLinkpfSmartNATIndex,
		Params:                     params,
	}
}

func (c *SlbStatLinkpfSmartNATTable) Name() string {
	return "SlbStatLinkpfSmartNATTable"
}

func (c *SlbStatLinkpfSmartNATTable) GetParams() BeanType {
	return c.Params
}

func (c *SlbStatLinkpfSmartNATTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *SlbStatLinkpfSmartNATTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.SlbStatLinkpfSmartNATIndex).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.SlbStatLinkpfSmartNATIndex)
}

type SlbStatLinkpfSmartNATTableNATType int32

const (
	SlbStatLinkpfSmartNATTableNATType_Nonat       SlbStatLinkpfSmartNATTableNATType = 0
	SlbStatLinkpfSmartNATTableNATType_Static      SlbStatLinkpfSmartNATTableNATType = 1
	SlbStatLinkpfSmartNATTableNATType_Dynamic     SlbStatLinkpfSmartNATTableNATType = 2
	SlbStatLinkpfSmartNATTableNATType_Unsupported SlbStatLinkpfSmartNATTableNATType = 2147483647
)

type SlbStatLinkpfSmartNATTableParams struct {
	// The SmartNAT index that identifies the SmartNAT entry.
	NATIndex string `json:"NATIndex,omitempty"`
	// The number of sessions that are currently handled by this SmartNAT entry.
	NATCurrSess uint32 `json:"NATCurrSess,omitempty"`
	// Total number of sessions that are handled by this SmartNAT entry.
	NATTotSess uint32 `json:"NATTotSess,omitempty"`
	// NAT type of sessions that are handled by this SmartNAT entry.
	NATType SlbStatLinkpfSmartNATTableNATType `json:"NATType,omitempty"`
}

func (p SlbStatLinkpfSmartNATTableParams) iMABean() {}
