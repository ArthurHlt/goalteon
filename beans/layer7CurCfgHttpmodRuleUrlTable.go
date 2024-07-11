package beans

import (
	"fmt"
	"reflect"
)

// Layer7CurCfgHttpmodRuleUrlTable The table for configuring HTTP content modification rules.
// Note:This mib is not supported for VX instance of virtualization.
type Layer7CurCfgHttpmodRuleUrlTable struct {
	// The HTTP Modification Rule List (key id) as an index.
	Layer7CurCfgHttpmodRuleUrlListIdIndex string
	// The HTTP Modification Rule number as an index.
	Layer7CurCfgHttpmodRuleUrlIndex int32
	Params                          *Layer7CurCfgHttpmodRuleUrlTableParams
}

func NewLayer7CurCfgHttpmodRuleUrlTableList() *Layer7CurCfgHttpmodRuleUrlTable {
	return &Layer7CurCfgHttpmodRuleUrlTable{}
}

func NewLayer7CurCfgHttpmodRuleUrlTable(
	layer7CurCfgHttpmodRuleUrlListIdIndex string,
	layer7CurCfgHttpmodRuleUrlIndex int32,
	params *Layer7CurCfgHttpmodRuleUrlTableParams,
) *Layer7CurCfgHttpmodRuleUrlTable {
	return &Layer7CurCfgHttpmodRuleUrlTable{
		Layer7CurCfgHttpmodRuleUrlListIdIndex: layer7CurCfgHttpmodRuleUrlListIdIndex,
		Layer7CurCfgHttpmodRuleUrlIndex:       layer7CurCfgHttpmodRuleUrlIndex,
		Params:                                params,
	}
}

func (c *Layer7CurCfgHttpmodRuleUrlTable) Name() string {
	return "Layer7CurCfgHttpmodRuleUrlTable"
}

func (c *Layer7CurCfgHttpmodRuleUrlTable) GetParams() BeanType {
	return c.Params
}

func (c *Layer7CurCfgHttpmodRuleUrlTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *Layer7CurCfgHttpmodRuleUrlTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.Layer7CurCfgHttpmodRuleUrlListIdIndex).IsZero() &&
		reflect.ValueOf(c.Layer7CurCfgHttpmodRuleUrlIndex).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.Layer7CurCfgHttpmodRuleUrlListIdIndex) + "/" + fmt.Sprint(c.Layer7CurCfgHttpmodRuleUrlIndex)
}

type Layer7CurCfgHttpmodRuleUrlTableMtchProtcol int32

const (
	Layer7CurCfgHttpmodRuleUrlTableMtchProtcol_Http        Layer7CurCfgHttpmodRuleUrlTableMtchProtcol = 1
	Layer7CurCfgHttpmodRuleUrlTableMtchProtcol_Https       Layer7CurCfgHttpmodRuleUrlTableMtchProtcol = 2
	Layer7CurCfgHttpmodRuleUrlTableMtchProtcol_Unsupported Layer7CurCfgHttpmodRuleUrlTableMtchProtcol = 2147483647
)

type Layer7CurCfgHttpmodRuleUrlTableMtchHostTyp int32

const (
	Layer7CurCfgHttpmodRuleUrlTableMtchHostTyp_Suffix      Layer7CurCfgHttpmodRuleUrlTableMtchHostTyp = 1
	Layer7CurCfgHttpmodRuleUrlTableMtchHostTyp_Prefix      Layer7CurCfgHttpmodRuleUrlTableMtchHostTyp = 2
	Layer7CurCfgHttpmodRuleUrlTableMtchHostTyp_Equal       Layer7CurCfgHttpmodRuleUrlTableMtchHostTyp = 3
	Layer7CurCfgHttpmodRuleUrlTableMtchHostTyp_Include     Layer7CurCfgHttpmodRuleUrlTableMtchHostTyp = 4
	Layer7CurCfgHttpmodRuleUrlTableMtchHostTyp_Any         Layer7CurCfgHttpmodRuleUrlTableMtchHostTyp = 5
	Layer7CurCfgHttpmodRuleUrlTableMtchHostTyp_Unsupported Layer7CurCfgHttpmodRuleUrlTableMtchHostTyp = 2147483647
)

type Layer7CurCfgHttpmodRuleUrlTableMtchPathTyp int32

const (
	Layer7CurCfgHttpmodRuleUrlTableMtchPathTyp_Suffix      Layer7CurCfgHttpmodRuleUrlTableMtchPathTyp = 1
	Layer7CurCfgHttpmodRuleUrlTableMtchPathTyp_Prefix      Layer7CurCfgHttpmodRuleUrlTableMtchPathTyp = 2
	Layer7CurCfgHttpmodRuleUrlTableMtchPathTyp_Equal       Layer7CurCfgHttpmodRuleUrlTableMtchPathTyp = 3
	Layer7CurCfgHttpmodRuleUrlTableMtchPathTyp_Include     Layer7CurCfgHttpmodRuleUrlTableMtchPathTyp = 4
	Layer7CurCfgHttpmodRuleUrlTableMtchPathTyp_Any         Layer7CurCfgHttpmodRuleUrlTableMtchPathTyp = 5
	Layer7CurCfgHttpmodRuleUrlTableMtchPathTyp_Unsupported Layer7CurCfgHttpmodRuleUrlTableMtchPathTyp = 2147483647
)

type Layer7CurCfgHttpmodRuleUrlTableActnProtcl int32

const (
	Layer7CurCfgHttpmodRuleUrlTableActnProtcl_Http        Layer7CurCfgHttpmodRuleUrlTableActnProtcl = 1
	Layer7CurCfgHttpmodRuleUrlTableActnProtcl_Https       Layer7CurCfgHttpmodRuleUrlTableActnProtcl = 2
	Layer7CurCfgHttpmodRuleUrlTableActnProtcl_Unsupported Layer7CurCfgHttpmodRuleUrlTableActnProtcl = 2147483647
)

type Layer7CurCfgHttpmodRuleUrlTableActnHostTyp int32

const (
	Layer7CurCfgHttpmodRuleUrlTableActnHostTyp_Insert      Layer7CurCfgHttpmodRuleUrlTableActnHostTyp = 1
	Layer7CurCfgHttpmodRuleUrlTableActnHostTyp_Replace     Layer7CurCfgHttpmodRuleUrlTableActnHostTyp = 2
	Layer7CurCfgHttpmodRuleUrlTableActnHostTyp_Remove      Layer7CurCfgHttpmodRuleUrlTableActnHostTyp = 3
	Layer7CurCfgHttpmodRuleUrlTableActnHostTyp_None        Layer7CurCfgHttpmodRuleUrlTableActnHostTyp = 4
	Layer7CurCfgHttpmodRuleUrlTableActnHostTyp_Unsupported Layer7CurCfgHttpmodRuleUrlTableActnHostTyp = 2147483647
)

type Layer7CurCfgHttpmodRuleUrlTableActnHstSec int32

const (
	Layer7CurCfgHttpmodRuleUrlTableActnHstSec_Before      Layer7CurCfgHttpmodRuleUrlTableActnHstSec = 1
	Layer7CurCfgHttpmodRuleUrlTableActnHstSec_After       Layer7CurCfgHttpmodRuleUrlTableActnHstSec = 2
	Layer7CurCfgHttpmodRuleUrlTableActnHstSec_Unsupported Layer7CurCfgHttpmodRuleUrlTableActnHstSec = 2147483647
)

type Layer7CurCfgHttpmodRuleUrlTableActnPathTyp int32

const (
	Layer7CurCfgHttpmodRuleUrlTableActnPathTyp_Insert      Layer7CurCfgHttpmodRuleUrlTableActnPathTyp = 1
	Layer7CurCfgHttpmodRuleUrlTableActnPathTyp_Replace     Layer7CurCfgHttpmodRuleUrlTableActnPathTyp = 2
	Layer7CurCfgHttpmodRuleUrlTableActnPathTyp_Remove      Layer7CurCfgHttpmodRuleUrlTableActnPathTyp = 3
	Layer7CurCfgHttpmodRuleUrlTableActnPathTyp_None        Layer7CurCfgHttpmodRuleUrlTableActnPathTyp = 4
	Layer7CurCfgHttpmodRuleUrlTableActnPathTyp_Unsupported Layer7CurCfgHttpmodRuleUrlTableActnPathTyp = 2147483647
)

type Layer7CurCfgHttpmodRuleUrlTableActnPthSctn int32

const (
	Layer7CurCfgHttpmodRuleUrlTableActnPthSctn_Before      Layer7CurCfgHttpmodRuleUrlTableActnPthSctn = 1
	Layer7CurCfgHttpmodRuleUrlTableActnPthSctn_After       Layer7CurCfgHttpmodRuleUrlTableActnPthSctn = 2
	Layer7CurCfgHttpmodRuleUrlTableActnPthSctn_Unsupported Layer7CurCfgHttpmodRuleUrlTableActnPthSctn = 2147483647
)

type Layer7CurCfgHttpmodRuleUrlTableParams struct {
	// The HTTP Modification Rule List (key id) as an index.
	ListIdIndex string `json:"ListIdIndex,omitempty"`
	// The HTTP Modification Rule number as an index.
	Index int32 `json:"Index,omitempty"`
	// Match protocol.
	MtchProtcol Layer7CurCfgHttpmodRuleUrlTableMtchProtcol `json:"MtchProtcol,omitempty"`
	// Match port.
	MtchPort int32 `json:"MtchPort,omitempty"`
	// Host match parameters.
	MtchHostTyp Layer7CurCfgHttpmodRuleUrlTableMtchHostTyp `json:"MtchHostTyp,omitempty"`
	// Host to Match.
	MtchHost string `json:"MtchHost,omitempty"`
	// Path match parameters.
	MtchPathTyp Layer7CurCfgHttpmodRuleUrlTableMtchPathTyp `json:"MtchPathTyp,omitempty"`
	// Path to match.
	MtchPath string `json:"MtchPath,omitempty"`
	// Match page.
	MtchPgName string `json:"MtchPgName,omitempty"`
	// Match page type.
	MtchPgTyp string `json:"MtchPgTyp,omitempty"`
	// Action protocol.
	ActnProtcl Layer7CurCfgHttpmodRuleUrlTableActnProtcl `json:"ActnProtcl,omitempty"`
	// Action port.
	ActnPort int32 `json:"ActnPort,omitempty"`
	// Host action parameters.
	ActnHostTyp Layer7CurCfgHttpmodRuleUrlTableActnHostTyp `json:"ActnHostTyp,omitempty"`
	// Host to Insert.
	ActnHost string `json:"ActnHost,omitempty"`
	// Insert Matched Section.
	ActnHstSec Layer7CurCfgHttpmodRuleUrlTableActnHstSec `json:"ActnHstSec,omitempty"`
	// New Host to Replace.
	ActnHstRplc string `json:"ActnHstRplc,omitempty"`
	// Path action parameters.
	ActnPathTyp Layer7CurCfgHttpmodRuleUrlTableActnPathTyp `json:"ActnPathTyp,omitempty"`
	// Path to Insert.
	ActnPath string `json:"ActnPath,omitempty"`
	// Insert Matched Section.
	ActnPthSctn Layer7CurCfgHttpmodRuleUrlTableActnPthSctn `json:"ActnPthSctn,omitempty"`
	// New Path to Replace.
	ActnPthRplc string `json:"ActnPthRplc,omitempty"`
	// Action page.
	ActnPgName string `json:"ActnPgName,omitempty"`
	// Action page type.
	ActnPgTyp string `json:"ActnPgTyp,omitempty"`
}

func (p Layer7CurCfgHttpmodRuleUrlTableParams) iMABean() {}
