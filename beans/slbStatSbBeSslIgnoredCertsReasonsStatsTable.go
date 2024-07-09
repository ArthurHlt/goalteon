package beans

import (
	"fmt"
	"reflect"
)

// SlbStatSbBeSslIgnoredCertsReasonsStatsTable Statistics table for Sideband backend SSL ingnored certificate reason.
// Note:This mib is not supported for VX instance of Virtualization.
type SlbStatSbBeSslIgnoredCertsReasonsStatsTable struct {
	// The Sideband name that identifies the Sideband policy.
	SlbStatSbBeSslIgnoredCertsReasonsSidebandID string
	// Sideband  index.
	SlbStatSbBeSslIgnoredCertsReasonsReasonIndex int32
	Params                                       *SlbStatSbBeSslIgnoredCertsReasonsStatsTableParams
}

func NewSlbStatSbBeSslIgnoredCertsReasonsStatsTableList() *SlbStatSbBeSslIgnoredCertsReasonsStatsTable {
	return &SlbStatSbBeSslIgnoredCertsReasonsStatsTable{}
}

func NewSlbStatSbBeSslIgnoredCertsReasonsStatsTable(
	slbStatSbBeSslIgnoredCertsReasonsSidebandID string,
	slbStatSbBeSslIgnoredCertsReasonsReasonIndex int32,
	params *SlbStatSbBeSslIgnoredCertsReasonsStatsTableParams,
) *SlbStatSbBeSslIgnoredCertsReasonsStatsTable {
	return &SlbStatSbBeSslIgnoredCertsReasonsStatsTable{
		SlbStatSbBeSslIgnoredCertsReasonsSidebandID:  slbStatSbBeSslIgnoredCertsReasonsSidebandID,
		SlbStatSbBeSslIgnoredCertsReasonsReasonIndex: slbStatSbBeSslIgnoredCertsReasonsReasonIndex,
		Params: params,
	}
}

func (c *SlbStatSbBeSslIgnoredCertsReasonsStatsTable) Name() string {
	return "SlbStatSbBeSslIgnoredCertsReasonsStatsTable"
}

func (c *SlbStatSbBeSslIgnoredCertsReasonsStatsTable) GetParams() BeanType {
	return c.Params
}

func (c *SlbStatSbBeSslIgnoredCertsReasonsStatsTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *SlbStatSbBeSslIgnoredCertsReasonsStatsTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.SlbStatSbBeSslIgnoredCertsReasonsSidebandID).IsZero() &&
		reflect.ValueOf(c.SlbStatSbBeSslIgnoredCertsReasonsReasonIndex).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.SlbStatSbBeSslIgnoredCertsReasonsSidebandID) + "/" + fmt.Sprint(c.SlbStatSbBeSslIgnoredCertsReasonsReasonIndex)
}

type SlbStatSbBeSslIgnoredCertsReasonsStatsTableParams struct {
	// The Sideband name that identifies the Sideband policy.
	SidebandID string `json:"SidebandID,omitempty"`
	// Sideband  index.
	ReasonIndex int32 `json:"ReasonIndex,omitempty"`
	// Sideband ignored certificate name.
	ReasonName string `json:"ReasonName,omitempty"`
	// Number of New SSL handshakes between AAS and server per second for this Sideband.
	NewhandShake uint32 `json:"NewhandShake,omitempty"`
	// Number of total SSL handshakes between AAS and server per second for this Sideband.
	NewhandShakeTotal uint32 `json:"NewhandShakeTotal,omitempty"`
}

func (p SlbStatSbBeSslIgnoredCertsReasonsStatsTableParams) iMABean() {}
