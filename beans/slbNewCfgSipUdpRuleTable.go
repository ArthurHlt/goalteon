package beans

import (
	"fmt"
	"reflect"
)

// SlbNewCfgSipUdpRuleTable The table of SIP Udp rule configuration in the current_config.
// Note:This mib is not supported for VX instance of virtualization.
type SlbNewCfgSipUdpRuleTable struct {
	// The UDP Rule number for which the SIP UDP rule table is related.
	SlbNewCfgSipUdpRuleIndex int32
	Params                   *SlbNewCfgSipUdpRuleTableParams
}

func NewSlbNewCfgSipUdpRuleTableList() *SlbNewCfgSipUdpRuleTable {
	return &SlbNewCfgSipUdpRuleTable{}
}

func NewSlbNewCfgSipUdpRuleTable(
	slbNewCfgSipUdpRuleIndex int32,
	params *SlbNewCfgSipUdpRuleTableParams,
) *SlbNewCfgSipUdpRuleTable {
	return &SlbNewCfgSipUdpRuleTable{
		SlbNewCfgSipUdpRuleIndex: slbNewCfgSipUdpRuleIndex,
		Params:                   params,
	}
}

func (c *SlbNewCfgSipUdpRuleTable) Name() string {
	return "SlbNewCfgSipUdpRuleTable"
}

func (c *SlbNewCfgSipUdpRuleTable) GetParams() BeanType {
	return c.Params
}

func (c *SlbNewCfgSipUdpRuleTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *SlbNewCfgSipUdpRuleTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.SlbNewCfgSipUdpRuleIndex).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.SlbNewCfgSipUdpRuleIndex)
}

type SlbNewCfgSipUdpRuleTableHdrFld int32

const (
	SlbNewCfgSipUdpRuleTableHdrFld_None        SlbNewCfgSipUdpRuleTableHdrFld = 0
	SlbNewCfgSipUdpRuleTableHdrFld_Callid      SlbNewCfgSipUdpRuleTableHdrFld = 1
	SlbNewCfgSipUdpRuleTableHdrFld_Contact     SlbNewCfgSipUdpRuleTableHdrFld = 2
	SlbNewCfgSipUdpRuleTableHdrFld_Contentlen  SlbNewCfgSipUdpRuleTableHdrFld = 3
	SlbNewCfgSipUdpRuleTableHdrFld_Cseq        SlbNewCfgSipUdpRuleTableHdrFld = 4
	SlbNewCfgSipUdpRuleTableHdrFld_Expires     SlbNewCfgSipUdpRuleTableHdrFld = 5
	SlbNewCfgSipUdpRuleTableHdrFld_From        SlbNewCfgSipUdpRuleTableHdrFld = 6
	SlbNewCfgSipUdpRuleTableHdrFld_Replyto     SlbNewCfgSipUdpRuleTableHdrFld = 7
	SlbNewCfgSipUdpRuleTableHdrFld_To          SlbNewCfgSipUdpRuleTableHdrFld = 8
	SlbNewCfgSipUdpRuleTableHdrFld_Via         SlbNewCfgSipUdpRuleTableHdrFld = 9
	SlbNewCfgSipUdpRuleTableHdrFld_Reqline     SlbNewCfgSipUdpRuleTableHdrFld = 10
	SlbNewCfgSipUdpRuleTableHdrFld_Method      SlbNewCfgSipUdpRuleTableHdrFld = 11
	SlbNewCfgSipUdpRuleTableHdrFld_Sdpcontent  SlbNewCfgSipUdpRuleTableHdrFld = 12
	SlbNewCfgSipUdpRuleTableHdrFld_Unsupported SlbNewCfgSipUdpRuleTableHdrFld = 2147483647
)

type SlbNewCfgSipUdpRuleTableState int32

const (
	SlbNewCfgSipUdpRuleTableState_Enabled     SlbNewCfgSipUdpRuleTableState = 1
	SlbNewCfgSipUdpRuleTableState_Disabled    SlbNewCfgSipUdpRuleTableState = 2
	SlbNewCfgSipUdpRuleTableState_Unsupported SlbNewCfgSipUdpRuleTableState = 2147483647
)

type SlbNewCfgSipUdpRuleTableDelete int32

const (
	SlbNewCfgSipUdpRuleTableDelete_Other       SlbNewCfgSipUdpRuleTableDelete = 1
	SlbNewCfgSipUdpRuleTableDelete_Delete      SlbNewCfgSipUdpRuleTableDelete = 2
	SlbNewCfgSipUdpRuleTableDelete_Unsupported SlbNewCfgSipUdpRuleTableDelete = 2147483647
)

type SlbNewCfgSipUdpRuleTableParams struct {
	// The UDP Rule number for which the SIP UDP rule table is related.
	Index int32 `json:"Index,omitempty"`
	// The name of the header field.
	HdrFld SlbNewCfgSipUdpRuleTableHdrFld `json:"HdrFld,omitempty"`
	// The content length of the UDP header field.
	Content string `json:"Content,omitempty"`
	// BWM contract for this rule.
	Contract int32 `json:"Contract,omitempty"`
	// The Alert message for this rule.
	Msg string `json:"Msg,omitempty"`
	// The severity of this rule.
	Severity int32 `json:"Severity,omitempty"`
	// Add dependent rules to this rule. A '0' value is returned when read.
	Add int32 `json:"Add,omitempty"`
	// Remove dependent rules from this rule.
	Rem int32 `json:"Rem,omitempty"`
	// Enable or disable this rule.
	State SlbNewCfgSipUdpRuleTableState `json:"State,omitempty"`
	// When set to the value of 2 (delete), the entire row is
	// deleted. When read, other(1) is returned. Setting the value
	// to anything other than 2(delete) has no effect on the state of the row.
	Delete SlbNewCfgSipUdpRuleTableDelete `json:"Delete,omitempty"`
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

func (p SlbNewCfgSipUdpRuleTableParams) iMABean() {}
