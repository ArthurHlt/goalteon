package beans

import (
	"fmt"
	"reflect"
)

// GslbNewDnsSecKeyTable Note:This mib is not supported for VX instance of Virtualization.
type GslbNewDnsSecKeyTable struct {
	// DNS Sec Table Key.
	GslbNewDnsSecKeyID string
	Params             *GslbNewDnsSecKeyTableParams
}

func NewGslbNewDnsSecKeyTableList() *GslbNewDnsSecKeyTable {
	return &GslbNewDnsSecKeyTable{}
}

func NewGslbNewDnsSecKeyTable(
	gslbNewDnsSecKeyID string,
	params *GslbNewDnsSecKeyTableParams,
) *GslbNewDnsSecKeyTable {
	return &GslbNewDnsSecKeyTable{
		GslbNewDnsSecKeyID: gslbNewDnsSecKeyID,
		Params:             params,
	}
}

func (c *GslbNewDnsSecKeyTable) Name() string {
	return "GslbNewDnsSecKeyTable"
}

func (c *GslbNewDnsSecKeyTable) GetParams() BeanType {
	return c.Params
}

func (c *GslbNewDnsSecKeyTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *GslbNewDnsSecKeyTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.GslbNewDnsSecKeyID).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.GslbNewDnsSecKeyID)
}

type GslbNewDnsSecKeyTableType int32

const (
	GslbNewDnsSecKeyTableType_KeyTypeKSK     GslbNewDnsSecKeyTableType = 1
	GslbNewDnsSecKeyTableType_KeyTypeZSK     GslbNewDnsSecKeyTableType = 2
	GslbNewDnsSecKeyTableType_KeyTypeInvalid GslbNewDnsSecKeyTableType = 3
	GslbNewDnsSecKeyTableType_Unsupported    GslbNewDnsSecKeyTableType = 2147483647
)

type GslbNewDnsSecKeyTableStatus int32

const (
	GslbNewDnsSecKeyTableStatus_Enable      GslbNewDnsSecKeyTableStatus = 1
	GslbNewDnsSecKeyTableStatus_Disable     GslbNewDnsSecKeyTableStatus = 2
	GslbNewDnsSecKeyTableStatus_Unsupported GslbNewDnsSecKeyTableStatus = 2147483647
)

type GslbNewDnsSecKeyTableSize int32

const (
	GslbNewDnsSecKeyTableSize_KeySize1024    GslbNewDnsSecKeyTableSize = 1
	GslbNewDnsSecKeyTableSize_KeySize2048    GslbNewDnsSecKeyTableSize = 2
	GslbNewDnsSecKeyTableSize_KeySize4096    GslbNewDnsSecKeyTableSize = 3
	GslbNewDnsSecKeyTableSize_KeySizeInvalid GslbNewDnsSecKeyTableSize = 4
	GslbNewDnsSecKeyTableSize_Unsupported    GslbNewDnsSecKeyTableSize = 2147483647
)

type GslbNewDnsSecKeyTableAlgo int32

const (
	GslbNewDnsSecKeyTableAlgo_KeyAlgoRsaSha1   GslbNewDnsSecKeyTableAlgo = 1
	GslbNewDnsSecKeyTableAlgo_KeyAlgoRsaSha256 GslbNewDnsSecKeyTableAlgo = 2
	GslbNewDnsSecKeyTableAlgo_KeyAlgoRsaSha512 GslbNewDnsSecKeyTableAlgo = 3
	GslbNewDnsSecKeyTableAlgo_Unsupported      GslbNewDnsSecKeyTableAlgo = 2147483647
)

type GslbNewDnsSecKeyTableDelete int32

const (
	GslbNewDnsSecKeyTableDelete_Other       GslbNewDnsSecKeyTableDelete = 1
	GslbNewDnsSecKeyTableDelete_Delete      GslbNewDnsSecKeyTableDelete = 2
	GslbNewDnsSecKeyTableDelete_Unsupported GslbNewDnsSecKeyTableDelete = 2147483647
)

type GslbNewDnsSecKeyTableGenerate int32

const (
	GslbNewDnsSecKeyTableGenerate_Other       GslbNewDnsSecKeyTableGenerate = 1
	GslbNewDnsSecKeyTableGenerate_Generate    GslbNewDnsSecKeyTableGenerate = 2
	GslbNewDnsSecKeyTableGenerate_Unsupported GslbNewDnsSecKeyTableGenerate = 2147483647
)

type GslbNewDnsSecKeyTableGenerateStatus int32

const (
	GslbNewDnsSecKeyTableGenerateStatus_NotGenerated GslbNewDnsSecKeyTableGenerateStatus = 1
	GslbNewDnsSecKeyTableGenerateStatus_Generated    GslbNewDnsSecKeyTableGenerateStatus = 2
	GslbNewDnsSecKeyTableGenerateStatus_InProgress   GslbNewDnsSecKeyTableGenerateStatus = 3
	GslbNewDnsSecKeyTableGenerateStatus_Unsupported  GslbNewDnsSecKeyTableGenerateStatus = 2147483647
)

type GslbNewDnsSecKeyTableParams struct {
	// DNS Sec Table Key.
	ID string `json:"ID,omitempty"`
	// Usage count.
	UseCount uint64 `json:"UseCount,omitempty"`
	// KSK or ZSK.
	Type GslbNewDnsSecKeyTableType `json:"Type,omitempty"`
	// Key state.
	Status GslbNewDnsSecKeyTableStatus `json:"Status,omitempty"`
	// Size of key (bits).
	Size GslbNewDnsSecKeyTableSize `json:"Size,omitempty"`
	// Signing algorithm.
	Algo GslbNewDnsSecKeyTableAlgo `json:"Algo,omitempty"`
	// DNSSEC - TTL.
	TTL uint64 `json:"TTL,omitempty"`
	// Expiration period in Seconds.
	ExpPeriod uint32 `json:"ExpPeriod,omitempty"`
	// RollOver period in Seconds.
	RollOverPeriod uint32 `json:"RollOverPeriod,omitempty"`
	// Signature validity period in Seconds.
	ValidityPeriod uint32 `json:"ValidityPeriod,omitempty"`
	// Signature publication period in Seconds.
	PublicationPeriod uint32 `json:"PublicationPeriod,omitempty"`
	// By setting the value to delete(2), the entire row is deleted.
	Delete GslbNewDnsSecKeyTableDelete `json:"Delete,omitempty"`
	// By setting the value to generate(2), the key will be generated.
	Generate GslbNewDnsSecKeyTableGenerate `json:"Generate,omitempty"`
	// Returns the generate status.
	GenerateStatus GslbNewDnsSecKeyTableGenerateStatus `json:"GenerateStatus,omitempty"`
}

func (p GslbNewDnsSecKeyTableParams) iMABean() {}
