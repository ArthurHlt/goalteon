package beans

import (
	"fmt"
	"reflect"
)

// SlbStatSbBeSslCiphUsageStatsTable Statistics table for Sideband backend SSL cipher usage.
// Note:This mib is not supported for VX instance of Virtualization.
type SlbStatSbBeSslCiphUsageStatsTable struct {
	// The Sideband name that identifies the Sideband policy.
	SlbStatSbBeSslCiphUsageSidebandID string
	// Sideband cipher index.
	SlbStatSbBeSslCiphUsageCiphIndex int32
	Params                           *SlbStatSbBeSslCiphUsageStatsTableParams
}

func NewSlbStatSbBeSslCiphUsageStatsTableList() *SlbStatSbBeSslCiphUsageStatsTable {
	return &SlbStatSbBeSslCiphUsageStatsTable{}
}

func NewSlbStatSbBeSslCiphUsageStatsTable(
	slbStatSbBeSslCiphUsageSidebandID string,
	slbStatSbBeSslCiphUsageCiphIndex int32,
	params *SlbStatSbBeSslCiphUsageStatsTableParams,
) *SlbStatSbBeSslCiphUsageStatsTable {
	return &SlbStatSbBeSslCiphUsageStatsTable{
		SlbStatSbBeSslCiphUsageSidebandID: slbStatSbBeSslCiphUsageSidebandID,
		SlbStatSbBeSslCiphUsageCiphIndex:  slbStatSbBeSslCiphUsageCiphIndex,
		Params:                            params,
	}
}

func (c *SlbStatSbBeSslCiphUsageStatsTable) Name() string {
	return "SlbStatSbBeSslCiphUsageStatsTable"
}

func (c *SlbStatSbBeSslCiphUsageStatsTable) GetParams() BeanType {
	return c.Params
}

func (c *SlbStatSbBeSslCiphUsageStatsTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *SlbStatSbBeSslCiphUsageStatsTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.SlbStatSbBeSslCiphUsageSidebandID).IsZero() &&
		reflect.ValueOf(c.SlbStatSbBeSslCiphUsageCiphIndex).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.SlbStatSbBeSslCiphUsageSidebandID) + "/" + fmt.Sprint(c.SlbStatSbBeSslCiphUsageCiphIndex)
}

type SlbStatSbBeSslCiphUsageStatsTableParams struct {
	// The Sideband name that identifies the Sideband policy.
	SidebandID string `json:"SidebandID,omitempty"`
	// Sideband cipher index.
	CiphIndex int32 `json:"CiphIndex,omitempty"`
	// Sideband cipher name.
	CiphName string `json:"CiphName,omitempty"`
	// Number of New SSL handshakes between AAS and server per second for this Sideband.
	NewhandShake uint32 `json:"NewhandShake,omitempty"`
	// Number of total SSL handshakes between AAS and server per second for this Sideband.
	NewhandShakeTotal uint32 `json:"NewhandShakeTotal,omitempty"`
}

func (p SlbStatSbBeSslCiphUsageStatsTableParams) iMABean() {}
