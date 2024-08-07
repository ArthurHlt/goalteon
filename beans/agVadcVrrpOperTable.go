package beans

import (
	"fmt"
	"reflect"
)

// AgVadcVrrpOperTable The table of vADC operations. Note:This mib is supported only in VX instance on Virtualization.
type AgVadcVrrpOperTable struct {
	// The vadc index. Note:This mib is supported only in VX instance on Virtualization.
	VadcVrrpOperIdx int32
	Params          *AgVadcVrrpOperTableParams
}

func NewAgVadcVrrpOperTableList() *AgVadcVrrpOperTable {
	return &AgVadcVrrpOperTable{}
}

func NewAgVadcVrrpOperTable(
	vadcVrrpOperIdx int32,
	params *AgVadcVrrpOperTableParams,
) *AgVadcVrrpOperTable {
	return &AgVadcVrrpOperTable{
		VadcVrrpOperIdx: vadcVrrpOperIdx,
		Params:          params,
	}
}

func (c *AgVadcVrrpOperTable) Name() string {
	return "AgVadcVrrpOperTable"
}

func (c *AgVadcVrrpOperTable) GetParams() BeanType {
	return c.Params
}

func (c *AgVadcVrrpOperTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *AgVadcVrrpOperTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.VadcVrrpOperIdx).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.VadcVrrpOperIdx)
}

type AgVadcVrrpOperTableVadcVrrpOperBackup int32

const (
	AgVadcVrrpOperTableVadcVrrpOperBackup_Ok          AgVadcVrrpOperTableVadcVrrpOperBackup = 1
	AgVadcVrrpOperTableVadcVrrpOperBackup_Backup      AgVadcVrrpOperTableVadcVrrpOperBackup = 2
	AgVadcVrrpOperTableVadcVrrpOperBackup_Unsupported AgVadcVrrpOperTableVadcVrrpOperBackup = 2147483647
)

type AgVadcVrrpOperTableParams struct {
	// The vadc index. Note:This mib is supported only in VX instance on Virtualization.
	VadcVrrpOperIdx int32 `json:"vadcVrrpOperIdx,omitempty"`
	// When set to a value of 'backup(2)' it forces the specified
	// vADC master virtual router into backup mode.
	// 'ok(1)' is returned when the object os read.
	// Note:This mib is supported only in VX instance on Virtualization.
	VadcVrrpOperBackup AgVadcVrrpOperTableVadcVrrpOperBackup `json:"vadcVrrpOperBackup,omitempty"`
}

func (p AgVadcVrrpOperTableParams) iMABean() {}
