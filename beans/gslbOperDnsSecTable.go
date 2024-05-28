package beans

import (
	"fmt"
	"reflect"
)

// GslbOperDnsSecTable Note:This mib is not supported for VX instance of Virtualization.
type GslbOperDnsSecTable struct {
	// DNS Sec Table Key.
	GslbOperDnsSecKeyID string
	Params              *GslbOperDnsSecTableParams
}

func NewGslbOperDnsSecTableList() *GslbOperDnsSecTable {
	return &GslbOperDnsSecTable{}
}

func NewGslbOperDnsSecTable(
	gslbOperDnsSecKeyID string,
	params *GslbOperDnsSecTableParams,
) *GslbOperDnsSecTable {
	return &GslbOperDnsSecTable{
		GslbOperDnsSecKeyID: gslbOperDnsSecKeyID,
		Params:              params,
	}
}

func (c *GslbOperDnsSecTable) Name() string {
	return "GslbOperDnsSecTable"
}

func (c *GslbOperDnsSecTable) GetParams() BeanType {
	return c.Params
}

func (c *GslbOperDnsSecTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *GslbOperDnsSecTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.GslbOperDnsSecKeyID).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.GslbOperDnsSecKeyID)
}

type GslbOperDnsSecTableEmergencyRollover int32

const (
	GslbOperDnsSecTableEmergencyRollover_Other       GslbOperDnsSecTableEmergencyRollover = 1
	GslbOperDnsSecTableEmergencyRollover_Rollover    GslbOperDnsSecTableEmergencyRollover = 2
	GslbOperDnsSecTableEmergencyRollover_Unsupported GslbOperDnsSecTableEmergencyRollover = 2147483647
)

type GslbOperDnsSecTableImmediateRollover int32

const (
	GslbOperDnsSecTableImmediateRollover_Other       GslbOperDnsSecTableImmediateRollover = 1
	GslbOperDnsSecTableImmediateRollover_Rollover    GslbOperDnsSecTableImmediateRollover = 2
	GslbOperDnsSecTableImmediateRollover_Unsupported GslbOperDnsSecTableImmediateRollover = 2147483647
)

type GslbOperDnsSecTableParams struct {
	// DNS Sec Table Key.
	KeyID string `json:"KeyID,omitempty"`
	// By setting the value to rollover(2), DNSSEC emergency rollover procedure of a key will be Performed.
	EmergencyRollover GslbOperDnsSecTableEmergencyRollover `json:"EmergencyRollover,omitempty"`
	// By setting the value to rollover(2), DNSSEC immediate rollover procedure of a key will be Performed.
	ImmediateRollover GslbOperDnsSecTableImmediateRollover `json:"ImmediateRollover,omitempty"`
}

func (p GslbOperDnsSecTableParams) iMABean() {}
