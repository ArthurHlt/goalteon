package beans

import (
	"fmt"
	"reflect"
)

// SlbNewAdvhcScriptTable Note:This mib is not supported on VX instance of Virtualization.
type SlbNewAdvhcScriptTable struct {
	// Script Health check id.
	SlbNewAdvhcScriptID string
	Params              *SlbNewAdvhcScriptTableParams
}

func NewSlbNewAdvhcScriptTableList() *SlbNewAdvhcScriptTable {
	return &SlbNewAdvhcScriptTable{}
}

func NewSlbNewAdvhcScriptTable(
	slbNewAdvhcScriptID string,
	params *SlbNewAdvhcScriptTableParams,
) *SlbNewAdvhcScriptTable {
	return &SlbNewAdvhcScriptTable{
		SlbNewAdvhcScriptID: slbNewAdvhcScriptID,
		Params:              params,
	}
}

func (c *SlbNewAdvhcScriptTable) Name() string {
	return "SlbNewAdvhcScriptTable"
}

func (c *SlbNewAdvhcScriptTable) GetParams() BeanType {
	return c.Params
}

func (c *SlbNewAdvhcScriptTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *SlbNewAdvhcScriptTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.SlbNewAdvhcScriptID).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.SlbNewAdvhcScriptID)
}

type SlbNewAdvhcScriptTableAddCloseCmd int32

const (
	SlbNewAdvhcScriptTableAddCloseCmd_Other       SlbNewAdvhcScriptTableAddCloseCmd = 1
	SlbNewAdvhcScriptTableAddCloseCmd_Close       SlbNewAdvhcScriptTableAddCloseCmd = 2
	SlbNewAdvhcScriptTableAddCloseCmd_Unsupported SlbNewAdvhcScriptTableAddCloseCmd = 2147483647
)

type SlbNewAdvhcScriptTableRemoveLastCmd int32

const (
	SlbNewAdvhcScriptTableRemoveLastCmd_Other       SlbNewAdvhcScriptTableRemoveLastCmd = 1
	SlbNewAdvhcScriptTableRemoveLastCmd_Remove      SlbNewAdvhcScriptTableRemoveLastCmd = 2
	SlbNewAdvhcScriptTableRemoveLastCmd_Unsupported SlbNewAdvhcScriptTableRemoveLastCmd = 2147483647
)

type SlbNewAdvhcScriptTableDelete int32

const (
	SlbNewAdvhcScriptTableDelete_Other       SlbNewAdvhcScriptTableDelete = 1
	SlbNewAdvhcScriptTableDelete_Delete      SlbNewAdvhcScriptTableDelete = 2
	SlbNewAdvhcScriptTableDelete_Unsupported SlbNewAdvhcScriptTableDelete = 2147483647
)

type SlbNewAdvhcScriptTableIPVer int32

const (
	SlbNewAdvhcScriptTableIPVer_Ipv4 SlbNewAdvhcScriptTableIPVer = 1
	SlbNewAdvhcScriptTableIPVer_Ipv6 SlbNewAdvhcScriptTableIPVer = 2
	SlbNewAdvhcScriptTableIPVer_None SlbNewAdvhcScriptTableIPVer = 3
)

type SlbNewAdvhcScriptTableInvert int32

const (
	SlbNewAdvhcScriptTableInvert_Enabled     SlbNewAdvhcScriptTableInvert = 1
	SlbNewAdvhcScriptTableInvert_Disabled    SlbNewAdvhcScriptTableInvert = 2
	SlbNewAdvhcScriptTableInvert_Unsupported SlbNewAdvhcScriptTableInvert = 2147483647
)

type SlbNewAdvhcScriptTableTransparent int32

const (
	SlbNewAdvhcScriptTableTransparent_Enabled     SlbNewAdvhcScriptTableTransparent = 1
	SlbNewAdvhcScriptTableTransparent_Disabled    SlbNewAdvhcScriptTableTransparent = 2
	SlbNewAdvhcScriptTableTransparent_Unsupported SlbNewAdvhcScriptTableTransparent = 2147483647
)

type SlbNewAdvhcScriptTableAlways int32

const (
	SlbNewAdvhcScriptTableAlways_Enabled     SlbNewAdvhcScriptTableAlways = 1
	SlbNewAdvhcScriptTableAlways_Disabled    SlbNewAdvhcScriptTableAlways = 2
	SlbNewAdvhcScriptTableAlways_Unsupported SlbNewAdvhcScriptTableAlways = 2147483647
)

type SlbNewAdvhcScriptTableSnat int32

const (
	SlbNewAdvhcScriptTableSnat_Enabled     SlbNewAdvhcScriptTableSnat = 1
	SlbNewAdvhcScriptTableSnat_Disabled    SlbNewAdvhcScriptTableSnat = 2
	SlbNewAdvhcScriptTableSnat_Unsupported SlbNewAdvhcScriptTableSnat = 2147483647
)

type SlbNewAdvhcScriptTableParams struct {
	// Script Health check id.
	ID string `json:"ID,omitempty"`
	// Script Health check name.
	Name string `json:"Name,omitempty"`
	// Script Health check string.
	StringVal string `json:"StringVal,omitempty"`
	// Append an 'open' command to the script. The value of this object
	// should be real server port number or name. Only the port number or
	// name should be set. The 'open ' characters will be automatically
	// appended. Also a 'null terminator'(\0) will be automatically appended.
	// The null terminator and the 'open ' characters will be counted as
	// characters in the script.
	AddOpenCmd string `json:"AddOpenCmd,omitempty"`
	// Append an 'send' command to the script. The set value of this object
	// should be the exact characters to be sent on the port opened with the
	// 'open' command. The 'send ' characters will be automatically appended.
	// Also a 'null terminator' (\0) will be automatically appended. The null
	// terminator and the 'send ' characters will be counted as characters in
	// the script.
	AddSendCmd string `json:"AddSendCmd,omitempty"`
	// Append a 'bsend' command to the script. The set value of this
	// object should be the exact binary data in hex format to be sent
	// on the port opened with the 'open' command. The 'bsend ' characters
	// will be automatically appended. Also a 'null terminator' (\0)
	// will be automatically appended. The null terminator and the 'bsend '
	// characters will be counted as characters in the script.
	AddBSendCmd string `json:"AddBSendCmd,omitempty"`
	// Append an 'nsend' command to the script. The 'nsend' command is
	// used to append additional content to the packet generated by the
	// 'bsend' command when the desired 'bsend' content is more than 255
	// characters.  The set value of this object should be the exact
	// characters to be appended.  The 'nsend ' characters will be
	// automatically appended.  Also a 'null terminator' (\0) will be
	// automatically appended. The null terminator and the 'nsend '
	// characters will be counted as characters in the script.
	AddNSendCmd string `json:"AddNSendCmd,omitempty"`
	// Append an 'expect' command to the script. The set value of this
	// object should be the exact characters expected to be received on
	// the port opened with the 'open' command.  The value could also be
	// a single wildcard character '*' which means any received content
	// will be accepted.  The 'expect ' characters will be automatically
	// appended.  Also a 'null terminator'(\0) will be automatically
	// appended.  The null terminator and the 'expect ' characters will
	// be counted as characters in the script.
	AddExpectCmd string `json:"AddExpectCmd,omitempty"`
	// Append a 'bexpect' command to the script. The set value of this
	// object should be the exact binary data in hex format expected to
	// be received on the port opened with the 'open' command. The value
	// could also be a single wildcard character '*' which means any
	// received content will be accepted.  The 'bexpect ' characters
	// will be automatically appended.  Also a 'null terminator'(\0)
	// will be automatically appended.  The null terminator and the
	// 'bexpect ' characters will be counted as characters in the script.
	AddBEexpectCmd string `json:"AddBEexpectCmd,omitempty"`
	// Append an 'nexpect' command to the script.  The 'nexpect' command
	// is used to append additional characters to the 'bexpect' string
	// when the desired 'bexpect' string is more than 255 characters.  The
	// set value of this object should be the exact characters to be
	// appended.  The 'nexpect ' characters will be automatically appended.
	// Also a 'null terminator'(\0) will be automatically appended.
	// The null terminator and the 'nexpect ' characters will be counted
	// as characters in the script.
	AddNExpectCmd string `json:"AddNExpectCmd,omitempty"`
	// Append an 'offset' command to the script. The set value of this
	// object should be the offset from the beginning of the data area
	// of the TCP/UDP packet to start matching the string configured
	// in the last bexpect command. The 'offset ' characters will be
	// automatically appended. Also a 'null terminator'(\0) will be
	// automatically appended. The null terminator and the 'offset '
	// characters will be counted as characters in the script.  When
	// read, the return value has no significant meaning, but a 1 is
	// always returned.
	AddOffsetCmd uint32 `json:"AddOffsetCmd,omitempty"`
	// Append a 'depth' command to the script. The set value of this
	// 	 object should be the depth (search window) within the data
	// 	 area of the TCP/UDP packet for matching the string configured
	// 	 in the last bexpect command. The 'depth ' characters will be
	// automatically appended. Also a 'null terminator'(\0) will be
	// automatically appended. The null terminator and the 'depth '
	// characters will be counted as characters in the script. When
	// 	 read, the return value has no significant meaning, but a 1 is
	// 	 always returned.
	AddDepthCmd uint32 `json:"AddDepthCmd,omitempty"`
	// Append a 'wait' command to the script. The set value of this
	// object should be the maximum waiting time for the packet
	// containing the content configured in the last expect command.
	// Wait time is in units of milliseconds, and the maximum
	// supported is 65535 ms.  The 'wait ' characters will be
	// automatically appended. Also a 'null terminator'(\0) will be
	// automatically appended. The null terminator and the 'wait '
	// characters will be counted as characters in the script. When
	// read, the return value has no significant meaning, but a 1 is
	// always returned.
	AddWaitCmd uint64 `json:"AddWaitCmd,omitempty"`
	// Add close command to end of the script. When set to the value of 2
	// (add), 'close' will be appended to the script. Also, a
	// 'null terminator'(\0) will be automatically appended. The null
	// terminator and the 'close' characters will be counted as characters
	// in the script.  This command is only valid for TCP.
	AddCloseCmd SlbNewAdvhcScriptTableAddCloseCmd `json:"AddCloseCmd,omitempty"`
	// Remove the last command from the script. When set to the value of 2
	// (remove), last command will be removed from the script. When read,
	// other(1) is returned.
	RemoveLastCmd SlbNewAdvhcScriptTableRemoveLastCmd `json:"RemoveLastCmd,omitempty"`
	// Script Health check copy indicator.
	Copy DisplayString `json:"Copy,omitempty"`
	// When set to the value of 2 (delete), the entire row is deleted.
	// When read, other(1) is returned. Setting the value to anything
	// other than 2(delete) has no effect on the state of the row.
	Delete SlbNewAdvhcScriptTableDelete `json:"Delete,omitempty"`
	// Script Health check Ip Version.
	IPVer SlbNewAdvhcScriptTableIPVer `json:"IPVer,omitempty"`
	// Script Healtch check ip address / Hostname.
	HostName string `json:"HostName,omitempty"`
	// Invert Result
	Invert SlbNewAdvhcScriptTableInvert `json:"Invert,omitempty"`
	// Script Health check retries counter.
	Retries uint64 `json:"Retries,omitempty"`
	// Script Health check retries in the down state counter.
	RestoreRetries uint64 `json:"RestoreRetries,omitempty"`
	// Script Health check timeout.
	Timeout uint64 `json:"Timeout,omitempty"`
	// Script Health check overflow flag.
	Overflow uint64 `json:"Overflow,omitempty"`
	// Script Health check interval in the down state.
	DownInterval uint64 `json:"DownInterval,omitempty"`
	// Script Health check transparent flag.
	Transparent SlbNewAdvhcScriptTableTransparent `json:"Transparent,omitempty"`
	// Script Health check interval.
	Interval uint64 `json:"Interval,omitempty"`
	// This flag determines whether HC is allowed for standalone real.
	Always SlbNewAdvhcScriptTableAlways `json:"Always,omitempty"`
	// Script Health Check src NAT (PIP) flag.
	Snat SlbNewAdvhcScriptTableSnat `json:"Snat,omitempty"`
}

func (p SlbNewAdvhcScriptTableParams) iMABean() {}
