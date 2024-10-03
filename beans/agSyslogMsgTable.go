package beans

import (
	"fmt"
	"reflect"
)

// AgSyslogMsgTable The table of syslog messages.
type AgSyslogMsgTable struct {
	// The syslog message table index.
	AgSyslogMsgIndex int32
	Params           *AgSyslogMsgTableParams
}

func NewAgSyslogMsgTableList() *AgSyslogMsgTable {
	return &AgSyslogMsgTable{}
}

func NewAgSyslogMsgTable(
	agSyslogMsgIndex int32,
	params *AgSyslogMsgTableParams,
) *AgSyslogMsgTable {
	return &AgSyslogMsgTable{
		AgSyslogMsgIndex: agSyslogMsgIndex,
		Params:           params,
	}
}

func (c *AgSyslogMsgTable) Name() string {
	return "AgSyslogMsgTable"
}

func (c *AgSyslogMsgTable) GetParams() BeanType {
	return c.Params
}

func (c *AgSyslogMsgTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *AgSyslogMsgTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.AgSyslogMsgIndex).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.AgSyslogMsgIndex)
}

type AgSyslogMsgTableParams struct {
	// The syslog message table index.
	Index int32 `json:"Index,omitempty"`
	// The syslog message.
	Message string `json:"Message,omitempty"`
	// The syslog message date.
	MessageDate string `json:"MessageDate,omitempty"`
	// The syslog message severity.
	MessageSeverity string `json:"MessageSeverity,omitempty"`
	// The syslog message plain message.
	MessageOnly string `json:"MessageOnly,omitempty"`
}

func (p AgSyslogMsgTableParams) iMABean() {}
