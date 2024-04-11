package beans

import (
	"fmt"
	"reflect"
)

// HcsNewCfgTable The scriptable health check table in the new configuration block.
// Note:This mib is not supported for VX instance of Virtualization.
type HcsNewCfgTable struct {
	// The index the scriptable health check.
	HcsNewCfgScriptIndex int32
	Params               *HcsNewCfgTableParams
}

func NewHcsNewCfgTableList() *HcsNewCfgTable {
	return &HcsNewCfgTable{}
}

func NewHcsNewCfgTable(
	hcsNewCfgScriptIndex int32,
	params *HcsNewCfgTableParams,
) *HcsNewCfgTable {
	return &HcsNewCfgTable{
		HcsNewCfgScriptIndex: hcsNewCfgScriptIndex,
		Params:               params,
	}
}

func (c *HcsNewCfgTable) Name() string {
	return "HcsNewCfgTable"
}

func (c *HcsNewCfgTable) GetParams() BeanType {
	return c.Params
}

func (c *HcsNewCfgTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *HcsNewCfgTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.HcsNewCfgScriptIndex).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.HcsNewCfgScriptIndex)
}

type HcsNewCfgTableAddCloseCmd int32

const (
	HcsNewCfgTableAddCloseCmd_Other       HcsNewCfgTableAddCloseCmd = 1
	HcsNewCfgTableAddCloseCmd_Close       HcsNewCfgTableAddCloseCmd = 2
	HcsNewCfgTableAddCloseCmd_Unsupported HcsNewCfgTableAddCloseCmd = 2147483647
)

type HcsNewCfgTableRemLastCmd int32

const (
	HcsNewCfgTableRemLastCmd_Other       HcsNewCfgTableRemLastCmd = 1
	HcsNewCfgTableRemLastCmd_Remove      HcsNewCfgTableRemLastCmd = 2
	HcsNewCfgTableRemLastCmd_Unsupported HcsNewCfgTableRemLastCmd = 2147483647
)

type HcsNewCfgTableDeleteScript int32

const (
	HcsNewCfgTableDeleteScript_Other       HcsNewCfgTableDeleteScript = 1
	HcsNewCfgTableDeleteScript_Delete      HcsNewCfgTableDeleteScript = 2
	HcsNewCfgTableDeleteScript_Unsupported HcsNewCfgTableDeleteScript = 2147483647
)

type HcsNewCfgTableParams struct {
	// The index the scriptable health check.
	ScriptIndex int32 `json:"ScriptIndex,omitempty"`
	// The scriptable health check string.
	ScriptString string `json:"ScriptString,omitempty"`
	// Append an 'send' command to the script. The set value of this object
	// should be the exact characters to be sent on the port opened with the
	// 'open' command. The 'send ' characters will be automatically appended.
	// Also a 'null terminator' (\0) will be automatically appended. The null
	// terminator and the 'send ' characters will be counted as characters in
	// the script.
	AddSendCmd string `json:"AddSendCmd,omitempty"`
	// Append an 'expect' command to the script. The set value of this
	// 	 object should be the exact characters expected to be received on
	// the port opened with the 'open' command.  The value could also be
	// a single wildcard character '*' which means any received content
	// 	 will be accepted.  The 'expect ' characters will be automatically
	// 	 appended.  Also a 'null terminator'(\0) will be automatically
	// appended.  The null terminator and the 'expect ' characters will
	// be counted as characters in the script.
	AddExpectCmd string `json:"AddExpectCmd,omitempty"`
	// Add close command to end of the script. When set to the value of 2
	// (add), 'close' will be appended to the script. Also, a
	// 'null terminator'(\0) will be automatically appended. The null
	// terminator and the 'close' characters will be counted as characters
	// in the script.  This command is only valid for TCP.
	AddCloseCmd HcsNewCfgTableAddCloseCmd `json:"AddCloseCmd,omitempty"`
	// Remove the last command from the script. When set to the value of 2
	// (remove), last command will be removed from the script. When read,
	// other(1) is returned.
	RemLastCmd HcsNewCfgTableRemLastCmd `json:"RemLastCmd,omitempty"`
	// When set to the value of 2 (delete), the entire row is deleted.
	// When read, other(1) is returned. Setting the value to anything
	// other than 2(delete) has no effect on the state of the row.
	DeleteScript HcsNewCfgTableDeleteScript `json:"DeleteScript,omitempty"`
	// Append an 'offset' command to the script. The set value of this
	// 	 object should be the offset from the beginning of the data area
	// 	 of the TCP/UDP packet to start matching the string configured
	// 	 in the last bexpect command. The 'offset ' characters will be
	// automatically appended. Also a 'null terminator'(\0) will be
	// automatically appended. The null terminator and the 'offset '
	// characters will be counted as characters in the script.  When
	// 	 read, the return value has no significant meaning, but a 1 is
	// 	 always returned.
	AddOffsetCmd uint32 `json:"AddOffsetCmd,omitempty"`
	// Append a 'wait' command to the script. The set value of this
	// 	 object should be the maximum waiting time for the packet
	// 	 containing the content configured in the last expect command.
	// 	 Wait time is in units of milliseconds, and the maximum
	// 	 supported is 65535 ms.  The 'wait ' characters will be
	// 	 automatically appended. Also a 'null terminator'(\0) will be
	// automatically appended. The null terminator and the 'wait '
	// characters will be counted as characters in the script. When
	// 	 read, the return value has no significant meaning, but a 1 is
	// 	 always returned.
	AddWaitCmd uint64 `json:"AddWaitCmd,omitempty"`
	// Append an 'open' command to the script. The value of this object
	// should be real server port number or name, followed by a comma,
	// 	 and then the protocol (either tcp or udp), e.g. '80,tcp'.  Only
	// 	 the port number or name, and tcp or udp should be set. The 'open '
	// 	 characters will be automatically appended. Also a 'null terminator'
	// (\0) will be automatically appended. The null terminator and the
	// 'open ' characters will be counted as characters in the script.
	AddOpenProtCmd string `json:"AddOpenProtCmd,omitempty"`
	// Append an 'nsend' command to the script. The 'nsend' command is
	// 	 used to append additional content to the packet generated by the
	// 	 'bsend' command when the desired 'bsend' content is more than 255
	// 	 characters.  The set value of this object should be the exact
	// 	 characters to be appended.  The 'nsend ' characters will be
	// automatically appended.  Also a 'null terminator' (\0) will be
	// automatically appended. The null terminator and the 'nsend '
	// characters will be counted as characters in the script.
	AddNsendCmd string `json:"AddNsendCmd,omitempty"`
	// Append an 'nexpect' command to the script.  The 'nexpect' command
	// 	 is used to append additional characters to the 'bexpect' string
	// 	 when the desired 'bexpect' string is more than 255 characters.  The
	// 	 set value of this object should be the exact characters to be
	// 	 appended.  The 'nexpect ' characters will be automatically appended.
	// Also a 'null terminator'(\0) will be automatically appended.
	// The null terminator and the 'nexpect ' characters will be counted
	// 	 as characters in the script.
	AddNexpectCmd string `json:"AddNexpectCmd,omitempty"`
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
	// Append a 'bsend' command to the script. The set value of this
	// 	 object should be the exact binary data in hex format to be sent
	// on the port opened with the 'open' command. The 'bsend ' characters
	// will be automatically appended. Also a 'null terminator' (\0)
	// 	 will be automatically appended. The null terminator and the 'bsend '
	// 	 characters will be counted as characters in the script.
	AddLongBsendCmd string `json:"AddLongBsendCmd,omitempty"`
	// Append a 'bexpect' command to the script. The set value of this
	// 	 object should be the exact binary data in hex format expected to
	// 	 be received on the port opened with the 'open' command. The value
	// could also be a single wildcard character '*' which means any
	// received content will be accepted.  The 'bexpect ' characters
	// 	 will be automatically appended.  Also a 'null terminator'(\0)
	// 	 will be automatically appended.  The null terminator and the
	// 	 'bexpect ' characters will be counted as characters in the script.
	AddLongBexpectCmd string `json:"AddLongBexpectCmd,omitempty"`
	// Append an 'send' command to the script. The set value of this object
	// should be the exact characters to be sent on the port opened with the
	// 'open' command. The 'send ' characters will be automatically appended.
	// Also a 'null terminator' (\0) will be automatically appended. The null
	// terminator and the 'send ' characters will be counted as characters in
	// the script.
	AddLongSendCmd string `json:"AddLongSendCmd,omitempty"`
	// Append an 'expect' command to the script. The set value of this
	// 	 object should be the exact characters expected to be received on
	// the port opened with the 'open' command.  The value could also be
	// a single wildcard character '*' which means any received content
	// 	 will be accepted.  The 'expect ' characters will be automatically
	// 	 appended.  Also a 'null terminator'(\0) will be automatically
	// appended.  The null terminator and the 'expect ' characters will
	// be counted as characters in the script.
	AddLongExpectCmd string `json:"AddLongExpectCmd,omitempty"`
	// Append an 'nsend' command to the script. The 'nsend' command is
	// 	 used to append additional content to the packet generated by the
	// 	 'bsend' command when the desired 'bsend' content is more than 255
	// 	 characters.  The set value of this object should be the exact
	// 	 characters to be appended.  The 'nsend ' characters will be
	// automatically appended.  Also a 'null terminator' (\0) will be
	// automatically appended. The null terminator and the 'nsend '
	// characters will be counted as characters in the script.
	AddLongNsendCmd string `json:"AddLongNsendCmd,omitempty"`
	// Append an 'nexpect' command to the script.  The 'nexpect' command
	// 	 is used to append additional characters to the 'bexpect' string
	// 	 when the desired 'bexpect' string is more than 255 characters.  The
	// 	 set value of this object should be the exact characters to be
	// 	 appended.  The 'nexpect ' characters will be automatically appended.
	// Also a 'null terminator'(\0) will be automatically appended.
	// The null terminator and the 'nexpect ' characters will be counted
	// 	 as characters in the script.
	AddLongNexpectCmd string `json:"AddLongNexpectCmd,omitempty"`
}

func (p HcsNewCfgTableParams) iMABean() {}
