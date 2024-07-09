package beans

import (
	"fmt"
	"reflect"
)

// SlbNewAdvhcExtScrHcScriptTbl The table of script entries which are used to do external script health check.
// Note: This MIB is not supported for VX instance of Virtualization.
type SlbNewAdvhcExtScrHcScriptTbl struct {
	// The alphanumeric ID of the script table entry. It is used to
	// associate the imported script file in this entry to an external
	// script health check object.
	// Note: This MIB object is not supported for VX instance of
	// virtualization.
	SlbNewAdvhcExtScrHcScriptId string
	Params                      *SlbNewAdvhcExtScrHcScriptTblParams
}

func NewSlbNewAdvhcExtScrHcScriptTblList() *SlbNewAdvhcExtScrHcScriptTbl {
	return &SlbNewAdvhcExtScrHcScriptTbl{}
}

func NewSlbNewAdvhcExtScrHcScriptTbl(
	slbNewAdvhcExtScrHcScriptId string,
	params *SlbNewAdvhcExtScrHcScriptTblParams,
) *SlbNewAdvhcExtScrHcScriptTbl {
	return &SlbNewAdvhcExtScrHcScriptTbl{
		SlbNewAdvhcExtScrHcScriptId: slbNewAdvhcExtScrHcScriptId,
		Params:                      params,
	}
}

func (c *SlbNewAdvhcExtScrHcScriptTbl) Name() string {
	return "SlbNewAdvhcExtScrHcScriptTbl"
}

func (c *SlbNewAdvhcExtScrHcScriptTbl) GetParams() BeanType {
	return c.Params
}

func (c *SlbNewAdvhcExtScrHcScriptTbl) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *SlbNewAdvhcExtScrHcScriptTbl) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.SlbNewAdvhcExtScrHcScriptId).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.SlbNewAdvhcExtScrHcScriptId)
}

type SlbNewAdvhcExtScrHcScriptTblDel int32

const (
	SlbNewAdvhcExtScrHcScriptTblDel_Other       SlbNewAdvhcExtScrHcScriptTblDel = 1
	SlbNewAdvhcExtScrHcScriptTblDel_Delete      SlbNewAdvhcExtScrHcScriptTblDel = 2
	SlbNewAdvhcExtScrHcScriptTblDel_Unsupported SlbNewAdvhcExtScrHcScriptTblDel = 2147483647
)

type SlbNewAdvhcExtScrHcScriptTblParams struct {
	// The alphanumeric ID of the script table entry. It is used to
	// associate the imported script file in this entry to an external
	// script health check object.
	// Note: This MIB object is not supported for VX instance of
	// virtualization.
	Id string `json:"Id,omitempty"`
	// The script file name which is going to be imported for external
	// script health check operation. The imported script file should
	// match the name configured in this object. Otherwise error will
	// be thrown.
	// Note: This MIB object is not supported for the VX instance of
	// Virtualization.
	Name string `json:"Name,omitempty"`
	// When set to the value of 2 (delete), the entire row is deleted.
	// When read, other(1) is returned. Setting the value to anything
	// other than 2(delete) has no effect on the state of the row.
	// This MIB object is used to delete the imported script which is
	// used in the external script health check operation. When this
	// entry is deleted, it removes its association from the external
	// script HC object to which it is associated. Subsequently, it
	// also removes the external script HC object association from the
	// reals and groups to which that HC object is associated.
	// Note: This MIB object is not supported for the VX instance of
	// Virtualization.
	Del SlbNewAdvhcExtScrHcScriptTblDel `json:"Del,omitempty"`
}

func (p SlbNewAdvhcExtScrHcScriptTblParams) iMABean() {}
