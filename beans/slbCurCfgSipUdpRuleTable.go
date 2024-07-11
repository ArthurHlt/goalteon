package beans

import (
	"fmt"
	"reflect"
)

// SlbCurCfgSipUdpRuleTable The table of SIP Udp rule configuration in the current_config.
// Note:This mib is not supported for VX instance of virtualization.
type SlbCurCfgSipUdpRuleTable struct {
	// The UDP Rule number for which the SIP UDP rule table is related.
	SlbCurCfgSipUdpRuleIndex int32
	Params                   *SlbCurCfgSipUdpRuleTableParams
}

func NewSlbCurCfgSipUdpRuleTableList() *SlbCurCfgSipUdpRuleTable {
	return &SlbCurCfgSipUdpRuleTable{}
}

func NewSlbCurCfgSipUdpRuleTable(
	slbCurCfgSipUdpRuleIndex int32,
	params *SlbCurCfgSipUdpRuleTableParams,
) *SlbCurCfgSipUdpRuleTable {
	return &SlbCurCfgSipUdpRuleTable{
		SlbCurCfgSipUdpRuleIndex: slbCurCfgSipUdpRuleIndex,
		Params:                   params,
	}
}

func (c *SlbCurCfgSipUdpRuleTable) Name() string {
	return "SlbCurCfgSipUdpRuleTable"
}

func (c *SlbCurCfgSipUdpRuleTable) GetParams() BeanType {
	return c.Params
}

func (c *SlbCurCfgSipUdpRuleTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *SlbCurCfgSipUdpRuleTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.SlbCurCfgSipUdpRuleIndex).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.SlbCurCfgSipUdpRuleIndex)
}

type SlbCurCfgSipUdpRuleTableHdrFld int32

const (
	SlbCurCfgSipUdpRuleTableHdrFld_None        SlbCurCfgSipUdpRuleTableHdrFld = 0
	SlbCurCfgSipUdpRuleTableHdrFld_Callid      SlbCurCfgSipUdpRuleTableHdrFld = 1
	SlbCurCfgSipUdpRuleTableHdrFld_Contact     SlbCurCfgSipUdpRuleTableHdrFld = 2
	SlbCurCfgSipUdpRuleTableHdrFld_Contentlen  SlbCurCfgSipUdpRuleTableHdrFld = 3
	SlbCurCfgSipUdpRuleTableHdrFld_Cseq        SlbCurCfgSipUdpRuleTableHdrFld = 4
	SlbCurCfgSipUdpRuleTableHdrFld_Expires     SlbCurCfgSipUdpRuleTableHdrFld = 5
	SlbCurCfgSipUdpRuleTableHdrFld_From        SlbCurCfgSipUdpRuleTableHdrFld = 6
	SlbCurCfgSipUdpRuleTableHdrFld_Replyto     SlbCurCfgSipUdpRuleTableHdrFld = 7
	SlbCurCfgSipUdpRuleTableHdrFld_To          SlbCurCfgSipUdpRuleTableHdrFld = 8
	SlbCurCfgSipUdpRuleTableHdrFld_Via         SlbCurCfgSipUdpRuleTableHdrFld = 9
	SlbCurCfgSipUdpRuleTableHdrFld_Reqline     SlbCurCfgSipUdpRuleTableHdrFld = 10
	SlbCurCfgSipUdpRuleTableHdrFld_Method      SlbCurCfgSipUdpRuleTableHdrFld = 11
	SlbCurCfgSipUdpRuleTableHdrFld_Sdpcontent  SlbCurCfgSipUdpRuleTableHdrFld = 12
	SlbCurCfgSipUdpRuleTableHdrFld_Unsupported SlbCurCfgSipUdpRuleTableHdrFld = 2147483647
)

type SlbCurCfgSipUdpRuleTableState int32

const (
	SlbCurCfgSipUdpRuleTableState_Enabled     SlbCurCfgSipUdpRuleTableState = 1
	SlbCurCfgSipUdpRuleTableState_Disabled    SlbCurCfgSipUdpRuleTableState = 2
	SlbCurCfgSipUdpRuleTableState_Unsupported SlbCurCfgSipUdpRuleTableState = 2147483647
)

type SlbCurCfgSipUdpRuleTableParams struct {
	// The UDP Rule number for which the SIP UDP rule table is related.
	Index int32 `json:"Index,omitempty"`
	// The name of the header field.
	HdrFld SlbCurCfgSipUdpRuleTableHdrFld `json:"HdrFld,omitempty"`
	// The content length of the UDP header field.
	Content string `json:"Content,omitempty"`
	// BWM contract for this rule.
	Contract int32 `json:"Contract,omitempty"`
	// The Alert message for this rule.
	Msg string `json:"Msg,omitempty"`
	// The severity of this rule.
	Severity int32 `json:"Severity,omitempty"`
	// The dependent rules of this rule.
	Add string `json:"Add,omitempty"`
	// Enable or disable this rule.
	State SlbCurCfgSipUdpRuleTableState `json:"State,omitempty"`
	// The SIP UDP Dependent Rules.
	// The SIP UDP Dependent Rules are presented in a bitmap format.
	// in receiving order:
	// OCTET 1  OCTET 2  .....
	// xxxxxxxx xxxxxxxx .....
	// ||    || |_ Dependent Rule
	// ||    ||
	// ||    ||___ Dependent Rule 4
	// ||    |____ Dependent Rule 3
	// ||      .    .   .
	// ||_________ Dependent Rule 2
	// |__________ Dependent Rule 1
	// where x : 1 - The represented Dependent Rule is selected
	// 0 - The represented Dependent Rule is not selected
	Bmap string `json:"Bmap,omitempty"`
}

func (p SlbCurCfgSipUdpRuleTableParams) iMABean() {}
