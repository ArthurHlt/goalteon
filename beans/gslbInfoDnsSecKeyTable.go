package beans

import (
	"fmt"
	"reflect"
)

// GslbInfoDnsSecKeyTable Note:This mib is not supported for VX instance of Virtualization.
type GslbInfoDnsSecKeyTable struct {
	// DNS Sec Table Key.
	GslbInfoDnsSecKeyID string
	Params              *GslbInfoDnsSecKeyTableParams
}

func NewGslbInfoDnsSecKeyTableList() *GslbInfoDnsSecKeyTable {
	return &GslbInfoDnsSecKeyTable{}
}

func NewGslbInfoDnsSecKeyTable(
	gslbInfoDnsSecKeyID string,
	params *GslbInfoDnsSecKeyTableParams,
) *GslbInfoDnsSecKeyTable {
	return &GslbInfoDnsSecKeyTable{
		GslbInfoDnsSecKeyID: gslbInfoDnsSecKeyID,
		Params:              params,
	}
}

func (c *GslbInfoDnsSecKeyTable) Name() string {
	return "GslbInfoDnsSecKeyTable"
}

func (c *GslbInfoDnsSecKeyTable) GetParams() BeanType {
	return c.Params
}

func (c *GslbInfoDnsSecKeyTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *GslbInfoDnsSecKeyTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.GslbInfoDnsSecKeyID).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.GslbInfoDnsSecKeyID)
}

type GslbInfoDnsSecKeyTableStatus int32

const (
	GslbInfoDnsSecKeyTableStatus_InitRollProess       GslbInfoDnsSecKeyTableStatus = 1
	GslbInfoDnsSecKeyTableStatus_NewKeyCreated        GslbInfoDnsSecKeyTableStatus = 2
	GslbInfoDnsSecKeyTableStatus_NewZskKeyDeployed    GslbInfoDnsSecKeyTableStatus = 3
	GslbInfoDnsSecKeyTableStatus_OldKeyRemoval        GslbInfoDnsSecKeyTableStatus = 4
	GslbInfoDnsSecKeyTableStatus_RetrDsFromParent     GslbInfoDnsSecKeyTableStatus = 5
	GslbInfoDnsSecKeyTableStatus_NewKskKeyDeployed    GslbInfoDnsSecKeyTableStatus = 6
	GslbInfoDnsSecKeyTableStatus_WaitDsChangeOnParent GslbInfoDnsSecKeyTableStatus = 7
	GslbInfoDnsSecKeyTableStatus_RolloverNotRunning   GslbInfoDnsSecKeyTableStatus = 8
	GslbInfoDnsSecKeyTableStatus_Invalid              GslbInfoDnsSecKeyTableStatus = 9
	GslbInfoDnsSecKeyTableStatus_Expired              GslbInfoDnsSecKeyTableStatus = 10
	GslbInfoDnsSecKeyTableStatus_Unsupported          GslbInfoDnsSecKeyTableStatus = 2147483647
)

type GslbInfoDnsSecKeyTableParams struct {
	// DNS Sec Table Key.
	ID string `json:"ID,omitempty"`
	// Key status.
	Status GslbInfoDnsSecKeyTableStatus `json:"Status,omitempty"`
}

func (p GslbInfoDnsSecKeyTableParams) iMABean() {}
