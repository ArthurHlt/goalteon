package beans

import (
	"fmt"
	"reflect"
)

// SlbStatWlmTable The WLM statistics table.
// Note:This mib is not supported for VX instance of Virtualization.
type SlbStatWlmTable struct {
	// The work load manager number.
	SlbStatWlmIndex int32
	Params          *SlbStatWlmTableParams
}

func NewSlbStatWlmTableList() *SlbStatWlmTable {
	return &SlbStatWlmTable{}
}

func NewSlbStatWlmTable(
	slbStatWlmIndex int32,
	params *SlbStatWlmTableParams,
) *SlbStatWlmTable {
	return &SlbStatWlmTable{
		SlbStatWlmIndex: slbStatWlmIndex,
		Params:          params,
	}
}

func (c *SlbStatWlmTable) Name() string {
	return "SlbStatWlmTable"
}

func (c *SlbStatWlmTable) GetParams() BeanType {
	return c.Params
}

func (c *SlbStatWlmTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *SlbStatWlmTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.SlbStatWlmIndex).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.SlbStatWlmIndex)
}

type SlbStatWlmTableParams struct {
	// The work load manager number.
	Index int32 `json:"Index,omitempty"`
	// Total number of registration request messages sent by the switch.
	RegReq uint32 `json:"RegReq,omitempty"`
	// Total number of registration replies received from the GWM.
	RegRep uint32 `json:"RegRep,omitempty"`
	// Total number of registration replies received where the return
	// code indicates an error.
	RegRepErr uint32 `json:"RegRepErr,omitempty"`
	// Total number of deregistration request messages sent by the switch.
	DeregReq uint32 `json:"DeregReq,omitempty"`
	// Total number of deregistration replies received from the GWM.
	DeregRep uint32 `json:"DeregRep,omitempty"`
	// Total number of deregistration replies received where the return code flag
	// set indicating an error in processing of the deregistration request.
	DeregRepErr uint32 `json:"DeregRepErr,omitempty"`
	// Total number of set LB state  request messages sent by the switch.
	LbStateReq uint32 `json:"LbStateReq,omitempty"`
	// Total number of set LB state replies received from the GWM.
	LbStateRep uint32 `json:"LbStateRep,omitempty"`
	// Total number of set LB state replies received where the return code code
	// flag set indicating an error in processing of the set LB state request.
	LbStateRepErr uint32 `json:"LbStateRepErr,omitempty"`
	// Total number of set member state  request messages sent by the switch.
	MembStateReq uint32 `json:"MembStateReq,omitempty"`
	// Total number of set member state  replies received from the GWM.
	MembStateRep uint32 `json:"MembStateRep,omitempty"`
	// Total number of set member state replies received where the return
	// code code flag set indicating an error in processing of the set
	// 	 member state request.
	MembStateRepErr uint32 `json:"MembStateRepErr,omitempty"`
	// Total number of weight updates received from the GWM.
	WtMsgRecv uint32 `json:"WtMsgRecv,omitempty"`
	// Total number of errorsencountered during parsing of the send weights message.
	WtMsgParErr uint32 `json:"WtMsgParErr,omitempty"`
	// This counter is incremented when the LB name in send weights message
	// does not match the switch LB name.
	TotInvalidLb uint32 `json:"TotInvalidLb,omitempty"`
	// This counter is incremented when the group name in send weights message
	// does not match with any of the groups configured on the switch.
	TotInvalidGrp uint32 `json:"TotInvalidGrp,omitempty"`
	// This counter is incremented when the real server name in send weights message
	// does not match with any of the reals configured on the switch.
	TotInvalidRealSer uint32 `json:"TotInvalidRealSer,omitempty"`
	// Total number of messages received from the GWM with invalid headers.
	MsgInvalidSASPHeader uint32 `json:"MsgInvalidSASPHeader,omitempty"`
	// Total number of parse errors.
	MsgParseErr uint32 `json:"MsgParseErr,omitempty"`
	// Total messages received where the message type is not supported by the switch.
	MsgUnsupMsgType uint32 `json:"MsgUnsupMsgType,omitempty"`
}

func (p SlbStatWlmTableParams) iMABean() {}
