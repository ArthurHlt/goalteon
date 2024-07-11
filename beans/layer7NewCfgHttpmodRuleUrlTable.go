package beans

import (
	"fmt"
	"reflect"
)

// Layer7NewCfgHttpmodRuleUrlTable The table for configuring HTTP content modification rules.
// Note:This mib is not supported for VX instance of virtualization.
type Layer7NewCfgHttpmodRuleUrlTable struct {
	// The HTTP Modification Rule List (key id) as an index.
	Layer7NewCfgHttpmodRuleUrlListIdIndex string
	// The HTTP Modification Rule number as an index.
	Layer7NewCfgHttpmodRuleUrlIndex int32
	Params                          *Layer7NewCfgHttpmodRuleUrlTableParams
}

func NewLayer7NewCfgHttpmodRuleUrlTableList() *Layer7NewCfgHttpmodRuleUrlTable {
	return &Layer7NewCfgHttpmodRuleUrlTable{}
}

func NewLayer7NewCfgHttpmodRuleUrlTable(
	layer7NewCfgHttpmodRuleUrlListIdIndex string,
	layer7NewCfgHttpmodRuleUrlIndex int32,
	params *Layer7NewCfgHttpmodRuleUrlTableParams,
) *Layer7NewCfgHttpmodRuleUrlTable {
	return &Layer7NewCfgHttpmodRuleUrlTable{
		Layer7NewCfgHttpmodRuleUrlListIdIndex: layer7NewCfgHttpmodRuleUrlListIdIndex,
		Layer7NewCfgHttpmodRuleUrlIndex:       layer7NewCfgHttpmodRuleUrlIndex,
		Params:                                params,
	}
}

func (c *Layer7NewCfgHttpmodRuleUrlTable) Name() string {
	return "Layer7NewCfgHttpmodRuleUrlTable"
}

func (c *Layer7NewCfgHttpmodRuleUrlTable) GetParams() BeanType {
	return c.Params
}

func (c *Layer7NewCfgHttpmodRuleUrlTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *Layer7NewCfgHttpmodRuleUrlTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.Layer7NewCfgHttpmodRuleUrlListIdIndex).IsZero() &&
		reflect.ValueOf(c.Layer7NewCfgHttpmodRuleUrlIndex).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.Layer7NewCfgHttpmodRuleUrlListIdIndex) + "/" + fmt.Sprint(c.Layer7NewCfgHttpmodRuleUrlIndex)
}

type Layer7NewCfgHttpmodRuleUrlTableMtchProtcol int32

const (
	Layer7NewCfgHttpmodRuleUrlTableMtchProtcol_Http        Layer7NewCfgHttpmodRuleUrlTableMtchProtcol = 1
	Layer7NewCfgHttpmodRuleUrlTableMtchProtcol_Https       Layer7NewCfgHttpmodRuleUrlTableMtchProtcol = 2
	Layer7NewCfgHttpmodRuleUrlTableMtchProtcol_Unsupported Layer7NewCfgHttpmodRuleUrlTableMtchProtcol = 2147483647
)

type Layer7NewCfgHttpmodRuleUrlTableMtchHostTyp int32

const (
	Layer7NewCfgHttpmodRuleUrlTableMtchHostTyp_Suffix      Layer7NewCfgHttpmodRuleUrlTableMtchHostTyp = 1
	Layer7NewCfgHttpmodRuleUrlTableMtchHostTyp_Prefix      Layer7NewCfgHttpmodRuleUrlTableMtchHostTyp = 2
	Layer7NewCfgHttpmodRuleUrlTableMtchHostTyp_Equal       Layer7NewCfgHttpmodRuleUrlTableMtchHostTyp = 3
	Layer7NewCfgHttpmodRuleUrlTableMtchHostTyp_Include     Layer7NewCfgHttpmodRuleUrlTableMtchHostTyp = 4
	Layer7NewCfgHttpmodRuleUrlTableMtchHostTyp_Any         Layer7NewCfgHttpmodRuleUrlTableMtchHostTyp = 5
	Layer7NewCfgHttpmodRuleUrlTableMtchHostTyp_Unsupported Layer7NewCfgHttpmodRuleUrlTableMtchHostTyp = 2147483647
)

type Layer7NewCfgHttpmodRuleUrlTableMtchPathTyp int32

const (
	Layer7NewCfgHttpmodRuleUrlTableMtchPathTyp_Suffix      Layer7NewCfgHttpmodRuleUrlTableMtchPathTyp = 1
	Layer7NewCfgHttpmodRuleUrlTableMtchPathTyp_Prefix      Layer7NewCfgHttpmodRuleUrlTableMtchPathTyp = 2
	Layer7NewCfgHttpmodRuleUrlTableMtchPathTyp_Equal       Layer7NewCfgHttpmodRuleUrlTableMtchPathTyp = 3
	Layer7NewCfgHttpmodRuleUrlTableMtchPathTyp_Include     Layer7NewCfgHttpmodRuleUrlTableMtchPathTyp = 4
	Layer7NewCfgHttpmodRuleUrlTableMtchPathTyp_Any         Layer7NewCfgHttpmodRuleUrlTableMtchPathTyp = 5
	Layer7NewCfgHttpmodRuleUrlTableMtchPathTyp_Unsupported Layer7NewCfgHttpmodRuleUrlTableMtchPathTyp = 2147483647
)

type Layer7NewCfgHttpmodRuleUrlTableActnProtcl int32

const (
	Layer7NewCfgHttpmodRuleUrlTableActnProtcl_Http        Layer7NewCfgHttpmodRuleUrlTableActnProtcl = 1
	Layer7NewCfgHttpmodRuleUrlTableActnProtcl_Https       Layer7NewCfgHttpmodRuleUrlTableActnProtcl = 2
	Layer7NewCfgHttpmodRuleUrlTableActnProtcl_Unsupported Layer7NewCfgHttpmodRuleUrlTableActnProtcl = 2147483647
)

type Layer7NewCfgHttpmodRuleUrlTableActnHostTyp int32

const (
	Layer7NewCfgHttpmodRuleUrlTableActnHostTyp_Insert      Layer7NewCfgHttpmodRuleUrlTableActnHostTyp = 1
	Layer7NewCfgHttpmodRuleUrlTableActnHostTyp_Replace     Layer7NewCfgHttpmodRuleUrlTableActnHostTyp = 2
	Layer7NewCfgHttpmodRuleUrlTableActnHostTyp_Remove      Layer7NewCfgHttpmodRuleUrlTableActnHostTyp = 3
	Layer7NewCfgHttpmodRuleUrlTableActnHostTyp_None        Layer7NewCfgHttpmodRuleUrlTableActnHostTyp = 4
	Layer7NewCfgHttpmodRuleUrlTableActnHostTyp_Unsupported Layer7NewCfgHttpmodRuleUrlTableActnHostTyp = 2147483647
)

type Layer7NewCfgHttpmodRuleUrlTableActnHstSec int32

const (
	Layer7NewCfgHttpmodRuleUrlTableActnHstSec_Before      Layer7NewCfgHttpmodRuleUrlTableActnHstSec = 1
	Layer7NewCfgHttpmodRuleUrlTableActnHstSec_After       Layer7NewCfgHttpmodRuleUrlTableActnHstSec = 2
	Layer7NewCfgHttpmodRuleUrlTableActnHstSec_Unsupported Layer7NewCfgHttpmodRuleUrlTableActnHstSec = 2147483647
)

type Layer7NewCfgHttpmodRuleUrlTableActnPathTyp int32

const (
	Layer7NewCfgHttpmodRuleUrlTableActnPathTyp_Insert      Layer7NewCfgHttpmodRuleUrlTableActnPathTyp = 1
	Layer7NewCfgHttpmodRuleUrlTableActnPathTyp_Replace     Layer7NewCfgHttpmodRuleUrlTableActnPathTyp = 2
	Layer7NewCfgHttpmodRuleUrlTableActnPathTyp_Remove      Layer7NewCfgHttpmodRuleUrlTableActnPathTyp = 3
	Layer7NewCfgHttpmodRuleUrlTableActnPathTyp_None        Layer7NewCfgHttpmodRuleUrlTableActnPathTyp = 4
	Layer7NewCfgHttpmodRuleUrlTableActnPathTyp_Unsupported Layer7NewCfgHttpmodRuleUrlTableActnPathTyp = 2147483647
)

type Layer7NewCfgHttpmodRuleUrlTableActnPthSctn int32

const (
	Layer7NewCfgHttpmodRuleUrlTableActnPthSctn_Before      Layer7NewCfgHttpmodRuleUrlTableActnPthSctn = 1
	Layer7NewCfgHttpmodRuleUrlTableActnPthSctn_After       Layer7NewCfgHttpmodRuleUrlTableActnPthSctn = 2
	Layer7NewCfgHttpmodRuleUrlTableActnPthSctn_Unsupported Layer7NewCfgHttpmodRuleUrlTableActnPthSctn = 2147483647
)

type Layer7NewCfgHttpmodRuleUrlTableParams struct {
	// The HTTP Modification Rule List (key id) as an index.
	ListIdIndex string `json:"ListIdIndex,omitempty"`
	// The HTTP Modification Rule number as an index.
	Index int32 `json:"Index,omitempty"`
	// Match protocol.
	MtchProtcol Layer7NewCfgHttpmodRuleUrlTableMtchProtcol `json:"MtchProtcol,omitempty"`
	// Match port.
	MtchPort int32 `json:"MtchPort,omitempty"`
	// Host match parameters.
	MtchHostTyp Layer7NewCfgHttpmodRuleUrlTableMtchHostTyp `json:"MtchHostTyp,omitempty"`
	// Host to Match.
	MtchHost string `json:"MtchHost,omitempty"`
	// Path match parameters.
	MtchPathTyp Layer7NewCfgHttpmodRuleUrlTableMtchPathTyp `json:"MtchPathTyp,omitempty"`
	// Path to match.
	MtchPath string `json:"MtchPath,omitempty"`
	// Match page.
	MtchPgName string `json:"MtchPgName,omitempty"`
	// Match page type.
	MtchPgTyp string `json:"MtchPgTyp,omitempty"`
	// Action protocol.
	ActnProtcl Layer7NewCfgHttpmodRuleUrlTableActnProtcl `json:"ActnProtcl,omitempty"`
	// Action port.
	ActnPort int32 `json:"ActnPort,omitempty"`
	// Host action parameters.
	ActnHostTyp Layer7NewCfgHttpmodRuleUrlTableActnHostTyp `json:"ActnHostTyp,omitempty"`
	// Host to Insert.
	ActnHost string `json:"ActnHost,omitempty"`
	// Insert Matched Section.
	ActnHstSec Layer7NewCfgHttpmodRuleUrlTableActnHstSec `json:"ActnHstSec,omitempty"`
	// New Host to Replace.
	ActnHstRplc string `json:"ActnHstRplc,omitempty"`
	// Path action parameters.
	ActnPathTyp Layer7NewCfgHttpmodRuleUrlTableActnPathTyp `json:"ActnPathTyp,omitempty"`
	// Path to Insert.
	ActnPath string `json:"ActnPath,omitempty"`
	// Insert Matched Section.
	ActnPthSctn Layer7NewCfgHttpmodRuleUrlTableActnPthSctn `json:"ActnPthSctn,omitempty"`
	// New Path to Replace.
	ActnPthRplc string `json:"ActnPthRplc,omitempty"`
	// Action page.
	ActnPgName string `json:"ActnPgName,omitempty"`
	// Action page type.
	ActnPgTyp string `json:"ActnPgTyp,omitempty"`
}

func (p Layer7NewCfgHttpmodRuleUrlTableParams) iMABean() {}
