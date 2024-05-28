package beans

import (
	"fmt"
	"reflect"
)

// SlbAppShapeStatsTable A table for AppShape statistics.
// Note:This mib is not supported on VX instance of Virtualization.
type SlbAppShapeStatsTable struct {
	// AppShape script identifier.
	SlbAppShapeScriptId string
	Params              *SlbAppShapeStatsTableParams
}

func NewSlbAppShapeStatsTableList() *SlbAppShapeStatsTable {
	return &SlbAppShapeStatsTable{}
}

func NewSlbAppShapeStatsTable(
	slbAppShapeScriptId string,
	params *SlbAppShapeStatsTableParams,
) *SlbAppShapeStatsTable {
	return &SlbAppShapeStatsTable{
		SlbAppShapeScriptId: slbAppShapeScriptId,
		Params:              params,
	}
}

func (c *SlbAppShapeStatsTable) Name() string {
	return "SlbAppShapeStatsTable"
}

func (c *SlbAppShapeStatsTable) GetParams() BeanType {
	return c.Params
}

func (c *SlbAppShapeStatsTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *SlbAppShapeStatsTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.SlbAppShapeScriptId).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.SlbAppShapeScriptId)
}

type SlbAppShapeStatsTableParams struct {
	// AppShape script identifier.
	ScriptId string `json:"ScriptId,omitempty"`
	// AppShape script identifier.
	// 0 BIT is for CLIENT_ACCEPTED,
	// 1 BIT is for CLIENT_DATA,
	// 2 BIT is for CLIENT_CLOSED,
	// 3 BIT is for SERVER_CONNECTED,
	// 4 BIT is for SERVER_DATA,
	// 5 BIT is for SERVER_CLOSED,
	// 6 BIT is for HTTP_REQUEST,
	// 7 BIT is for HTTP_REQUEST_DATA,
	// 8 BIT is for HTTP_CRULE_MATCH,
	// 9 BIT is for HTTP_CRULE_NOMATCH,
	// 10 BIT is for HTTP_RESPONSE,
	// 11 BIT is for HTTP_CACHE_RESPONSE,
	// 12 BIT is for HTTP_RESPONSE_CONTINUE,
	// 13 BIT is for HTTP_RESPONSE_DATA,
	// 14 BIT is for SIP_REQUEST,
	// 15 BIT is for SIP_RESPONSE,
	// 16 BIT is for PERSIST_DOWN,
	// 	17 BIT is for LB_SELECTED,
	// 	18 BIT is for LB_FAILED,
	// 	19 BIT is for NAME_RESOLVED,
	// 	20 BIT is for SESSION_CLOSE.
	// If the bit is SET than the event has occurred.
	Event int32 `json:"Event,omitempty"`
	// Number of activations.
	Activations int32 `json:"Activations,omitempty"`
	// Number of failures.
	Failures int32 `json:"Failures,omitempty"`
	// Number of aborts.
	Aborts int32 `json:"Aborts,omitempty"`
	// The events that occurred - displayed in a string
	EventAsString string `json:"EventAsString,omitempty"`
}

func (p SlbAppShapeStatsTableParams) iMABean() {}
