package beans

import (
	"fmt"
	"reflect"
)

// SlbCDPGrpStatsTable A table for CDP group statistics.
// Note:This mib is not supported on VX instance of Virtualization.
type SlbCDPGrpStatsTable struct {
	// CDP group index in alphanumeric.
	SlbCDPGrpIndex string
	Params         *SlbCDPGrpStatsTableParams
}

func NewSlbCDPGrpStatsTableList() *SlbCDPGrpStatsTable {
	return &SlbCDPGrpStatsTable{}
}

func NewSlbCDPGrpStatsTable(
	slbCDPGrpIndex string,
	params *SlbCDPGrpStatsTableParams,
) *SlbCDPGrpStatsTable {
	return &SlbCDPGrpStatsTable{
		SlbCDPGrpIndex: slbCDPGrpIndex,
		Params:         params,
	}
}

func (c *SlbCDPGrpStatsTable) Name() string {
	return "SlbCDPGrpStatsTable"
}

func (c *SlbCDPGrpStatsTable) GetParams() BeanType {
	return c.Params
}

func (c *SlbCDPGrpStatsTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *SlbCDPGrpStatsTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.SlbCDPGrpIndex).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.SlbCDPGrpIndex)
}

type SlbCDPGrpStatsTableParams struct {
	// CDP group index in alphanumeric.
	Index string `json:"Index,omitempty"`
	// Last Successful Download.
	LastSuccess string `json:"LastSuccess,omitempty"`
	// Last Failed Download.
	LastFailed string `json:"LastFailed,omitempty"`
}

func (p SlbCDPGrpStatsTableParams) iMABean() {}
