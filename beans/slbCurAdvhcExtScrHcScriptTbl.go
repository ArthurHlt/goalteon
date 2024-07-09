package beans

import (
	"fmt"
	"reflect"
)

// SlbCurAdvhcExtScrHcScriptTbl The table of script entries which are used to do external script health check.
// Note: This MIB is not supported for VX instance of Virtualization.
type SlbCurAdvhcExtScrHcScriptTbl struct {
	// The alphanumeric ID of the script table entry. It is used to
	// associate the imported script file in this entry to an external
	// script health check object.
	// Note: This MIB object is not supported for VX instance of
	// virtualization.
	SlbCurAdvhcExtScrHcScriptId string
	Params                      *SlbCurAdvhcExtScrHcScriptTblParams
}

func NewSlbCurAdvhcExtScrHcScriptTblList() *SlbCurAdvhcExtScrHcScriptTbl {
	return &SlbCurAdvhcExtScrHcScriptTbl{}
}

func NewSlbCurAdvhcExtScrHcScriptTbl(
	slbCurAdvhcExtScrHcScriptId string,
	params *SlbCurAdvhcExtScrHcScriptTblParams,
) *SlbCurAdvhcExtScrHcScriptTbl {
	return &SlbCurAdvhcExtScrHcScriptTbl{
		SlbCurAdvhcExtScrHcScriptId: slbCurAdvhcExtScrHcScriptId,
		Params:                      params,
	}
}

func (c *SlbCurAdvhcExtScrHcScriptTbl) Name() string {
	return "SlbCurAdvhcExtScrHcScriptTbl"
}

func (c *SlbCurAdvhcExtScrHcScriptTbl) GetParams() BeanType {
	return c.Params
}

func (c *SlbCurAdvhcExtScrHcScriptTbl) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *SlbCurAdvhcExtScrHcScriptTbl) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.SlbCurAdvhcExtScrHcScriptId).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.SlbCurAdvhcExtScrHcScriptId)
}

type SlbCurAdvhcExtScrHcScriptTblParams struct {
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
}

func (p SlbCurAdvhcExtScrHcScriptTblParams) iMABean() {}
