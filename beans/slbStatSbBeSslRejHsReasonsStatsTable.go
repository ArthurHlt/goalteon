package beans

import (
	"fmt"
	"reflect"
)

// SlbStatSbBeSslRejHsReasonsStatsTable Statistics table for Sideband backend SSL reject handshake reason.
// Note:This mib is not supported for VX instance of Virtualization.
type SlbStatSbBeSslRejHsReasonsStatsTable struct {
	// The Sideband name that identifies the Sideband policy.
	SlbStatSbBeSslRejHsReasonsSidebandID string
	// Sideband  index.
	SlbStatSbBeSslRejHsReasonsReasonIndex int32
	Params                                *SlbStatSbBeSslRejHsReasonsStatsTableParams
}

func NewSlbStatSbBeSslRejHsReasonsStatsTableList() *SlbStatSbBeSslRejHsReasonsStatsTable {
	return &SlbStatSbBeSslRejHsReasonsStatsTable{}
}

func NewSlbStatSbBeSslRejHsReasonsStatsTable(
	slbStatSbBeSslRejHsReasonsSidebandID string,
	slbStatSbBeSslRejHsReasonsReasonIndex int32,
	params *SlbStatSbBeSslRejHsReasonsStatsTableParams,
) *SlbStatSbBeSslRejHsReasonsStatsTable {
	return &SlbStatSbBeSslRejHsReasonsStatsTable{
		SlbStatSbBeSslRejHsReasonsSidebandID:  slbStatSbBeSslRejHsReasonsSidebandID,
		SlbStatSbBeSslRejHsReasonsReasonIndex: slbStatSbBeSslRejHsReasonsReasonIndex,
		Params:                                params,
	}
}

func (c *SlbStatSbBeSslRejHsReasonsStatsTable) Name() string {
	return "SlbStatSbBeSslRejHsReasonsStatsTable"
}

func (c *SlbStatSbBeSslRejHsReasonsStatsTable) GetParams() BeanType {
	return c.Params
}

func (c *SlbStatSbBeSslRejHsReasonsStatsTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *SlbStatSbBeSslRejHsReasonsStatsTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.SlbStatSbBeSslRejHsReasonsSidebandID).IsZero() &&
		reflect.ValueOf(c.SlbStatSbBeSslRejHsReasonsReasonIndex).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.SlbStatSbBeSslRejHsReasonsSidebandID) + "/" + fmt.Sprint(c.SlbStatSbBeSslRejHsReasonsReasonIndex)
}

type SlbStatSbBeSslRejHsReasonsStatsTableParams struct {
	// The Sideband name that identifies the Sideband policy.
	SidebandID string `json:"SidebandID,omitempty"`
	// Sideband  index.
	ReasonIndex int32 `json:"ReasonIndex,omitempty"`
	// Sideband reject handshake name.
	ReasonName string `json:"ReasonName,omitempty"`
	// Number of New SSL handshakes between AAS and server per second for this Sideband.
	NewhandShake uint32 `json:"NewhandShake,omitempty"`
	// Number of total SSL handshakes between AAS and server per second for this Sideband.
	NewhandShakeTotal uint32 `json:"NewhandShakeTotal,omitempty"`
}

func (p SlbStatSbBeSslRejHsReasonsStatsTableParams) iMABean() {}
