package beans

import (
	"fmt"
	"reflect"
)

// HcsCurCfgTable The scriptable health check table in the current configuration block.
// Note:This mib is not supported for VX instance of Virtualization.
type HcsCurCfgTable struct {
	// The index in the scriptable health check table.
	HcsCurCfgScriptIndex int32
	Params               *HcsCurCfgTableParams
}

func NewHcsCurCfgTableList() *HcsCurCfgTable {
	return &HcsCurCfgTable{}
}

func NewHcsCurCfgTable(
	hcsCurCfgScriptIndex int32,
	params *HcsCurCfgTableParams,
) *HcsCurCfgTable {
	return &HcsCurCfgTable{
		HcsCurCfgScriptIndex: hcsCurCfgScriptIndex,
		Params:               params,
	}
}

func (c *HcsCurCfgTable) Name() string {
	return "HcsCurCfgTable"
}

func (c *HcsCurCfgTable) GetParams() BeanType {
	return c.Params
}

func (c *HcsCurCfgTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *HcsCurCfgTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.HcsCurCfgScriptIndex).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.HcsCurCfgScriptIndex)
}

type HcsCurCfgTableParams struct {
	// The index in the scriptable health check table.
	ScriptIndex int32 `json:"ScriptIndex,omitempty"`
	// The scriptable health check string.
	ScriptString string `json:"ScriptString,omitempty"`
}

func (p HcsCurCfgTableParams) iMABean() {}
