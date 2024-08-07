package beans

import (
	"fmt"
	"reflect"
)

// AgLastSyncInfoTable The table of Last Sync Information.
type AgLastSyncInfoTable struct {
	// Last sync table peer index.
	LastSyncInfoIdx int32
	Params          *AgLastSyncInfoTableParams
}

func NewAgLastSyncInfoTableList() *AgLastSyncInfoTable {
	return &AgLastSyncInfoTable{}
}

func NewAgLastSyncInfoTable(
	lastSyncInfoIdx int32,
	params *AgLastSyncInfoTableParams,
) *AgLastSyncInfoTable {
	return &AgLastSyncInfoTable{
		LastSyncInfoIdx: lastSyncInfoIdx,
		Params:          params,
	}
}

func (c *AgLastSyncInfoTable) Name() string {
	return "AgLastSyncInfoTable"
}

func (c *AgLastSyncInfoTable) GetParams() BeanType {
	return c.Params
}

func (c *AgLastSyncInfoTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *AgLastSyncInfoTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.LastSyncInfoIdx).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.LastSyncInfoIdx)
}

type AgLastSyncInfoTableLastSyncType int32

const (
	AgLastSyncInfoTableLastSyncType_None        AgLastSyncInfoTableLastSyncType = 1
	AgLastSyncInfoTableLastSyncType_Manual      AgLastSyncInfoTableLastSyncType = 2
	AgLastSyncInfoTableLastSyncType_Auto        AgLastSyncInfoTableLastSyncType = 3
	AgLastSyncInfoTableLastSyncType_Unsupported AgLastSyncInfoTableLastSyncType = 2147483647
)

type AgLastSyncInfoTableLastSyncStatus int32

const (
	AgLastSyncInfoTableLastSyncStatus_Idle        AgLastSyncInfoTableLastSyncStatus = 1
	AgLastSyncInfoTableLastSyncStatus_InProgress  AgLastSyncInfoTableLastSyncStatus = 2
	AgLastSyncInfoTableLastSyncStatus_Success     AgLastSyncInfoTableLastSyncStatus = 3
	AgLastSyncInfoTableLastSyncStatus_Failure     AgLastSyncInfoTableLastSyncStatus = 4
	AgLastSyncInfoTableLastSyncStatus_Unsupported AgLastSyncInfoTableLastSyncStatus = 2147483647
)

type AgLastSyncInfoTableLastSuccessfulSyncType int32

const (
	AgLastSyncInfoTableLastSuccessfulSyncType_None        AgLastSyncInfoTableLastSuccessfulSyncType = 1
	AgLastSyncInfoTableLastSuccessfulSyncType_Manual      AgLastSyncInfoTableLastSuccessfulSyncType = 2
	AgLastSyncInfoTableLastSuccessfulSyncType_Auto        AgLastSyncInfoTableLastSuccessfulSyncType = 3
	AgLastSyncInfoTableLastSuccessfulSyncType_Unsupported AgLastSyncInfoTableLastSuccessfulSyncType = 2147483647
)

type AgLastSyncInfoTableParams struct {
	// Last sync table peer index.
	LastSyncInfoIdx int32 `json:"lastSyncInfoIdx,omitempty"`
	// Last sync peer ip address.
	LastSyncPeerIp string `json:"lastSyncPeerIp,omitempty"`
	// Last sync type - none(1),manual(2) and auto(3).
	LastSyncType AgLastSyncInfoTableLastSyncType `json:"lastSyncType,omitempty"`
	// Status of last sync sent- idle(1), in-progress(2),success(3) and failure(4).
	LastSyncStatus AgLastSyncInfoTableLastSyncStatus `json:"lastSyncStatus,omitempty"`
	// Timestamp for last sync sent.
	LastSyncTime string `json:"lastSyncTime,omitempty"`
	// Failure reason for last sync.
	LastSyncFailReason string `json:"lastSyncFailReason,omitempty"`
	// Type of the last successful sync sent - none(1), manual(2), auto(3).
	LastSuccessfulSyncType AgLastSyncInfoTableLastSuccessfulSyncType `json:"lastSuccessfulSyncType,omitempty"`
	// Timestamp for last successful sync sent.
	LastSuccessfulSyncTime string `json:"lastSuccessfulSyncTime,omitempty"`
}

func (p AgLastSyncInfoTableParams) iMABean() {}
