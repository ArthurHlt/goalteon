package beans

import (
	"fmt"
	"reflect"
)

// GslbCurDnsSecKeyTable Note:This mib is not supported for VX instance of Virtualization.
type GslbCurDnsSecKeyTable struct {
	// DNS Sec Table Key.
	GslbCurDnsSecKeyID string
	Params             *GslbCurDnsSecKeyTableParams
}

func NewGslbCurDnsSecKeyTableList() *GslbCurDnsSecKeyTable {
	return &GslbCurDnsSecKeyTable{}
}

func NewGslbCurDnsSecKeyTable(
	gslbCurDnsSecKeyID string,
	params *GslbCurDnsSecKeyTableParams,
) *GslbCurDnsSecKeyTable {
	return &GslbCurDnsSecKeyTable{
		GslbCurDnsSecKeyID: gslbCurDnsSecKeyID,
		Params:             params,
	}
}

func (c *GslbCurDnsSecKeyTable) Name() string {
	return "GslbCurDnsSecKeyTable"
}

func (c *GslbCurDnsSecKeyTable) GetParams() BeanType {
	return c.Params
}

func (c *GslbCurDnsSecKeyTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *GslbCurDnsSecKeyTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.GslbCurDnsSecKeyID).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.GslbCurDnsSecKeyID)
}

type GslbCurDnsSecKeyTableType int32

const (
	GslbCurDnsSecKeyTableType_KeyTypeKSK     GslbCurDnsSecKeyTableType = 1
	GslbCurDnsSecKeyTableType_KeyTypeZSK     GslbCurDnsSecKeyTableType = 2
	GslbCurDnsSecKeyTableType_KeyTypeInvalid GslbCurDnsSecKeyTableType = 3
	GslbCurDnsSecKeyTableType_Unsupported    GslbCurDnsSecKeyTableType = 2147483647
)

type GslbCurDnsSecKeyTableStatus int32

const (
	GslbCurDnsSecKeyTableStatus_Enable      GslbCurDnsSecKeyTableStatus = 1
	GslbCurDnsSecKeyTableStatus_Disable     GslbCurDnsSecKeyTableStatus = 2
	GslbCurDnsSecKeyTableStatus_Unsupported GslbCurDnsSecKeyTableStatus = 2147483647
)

type GslbCurDnsSecKeyTableSize int32

const (
	GslbCurDnsSecKeyTableSize_KeySize1024    GslbCurDnsSecKeyTableSize = 1
	GslbCurDnsSecKeyTableSize_KeySize2048    GslbCurDnsSecKeyTableSize = 2
	GslbCurDnsSecKeyTableSize_KeySize4096    GslbCurDnsSecKeyTableSize = 3
	GslbCurDnsSecKeyTableSize_KeySizeInvalid GslbCurDnsSecKeyTableSize = 4
	GslbCurDnsSecKeyTableSize_Unsupported    GslbCurDnsSecKeyTableSize = 2147483647
)

type GslbCurDnsSecKeyTableAlgo int32

const (
	GslbCurDnsSecKeyTableAlgo_KeyAlgoRsaSha1   GslbCurDnsSecKeyTableAlgo = 1
	GslbCurDnsSecKeyTableAlgo_KeyAlgoRsaSha256 GslbCurDnsSecKeyTableAlgo = 2
	GslbCurDnsSecKeyTableAlgo_KeyAlgoRsaSha512 GslbCurDnsSecKeyTableAlgo = 3
	GslbCurDnsSecKeyTableAlgo_Unsupported      GslbCurDnsSecKeyTableAlgo = 2147483647
)

type GslbCurDnsSecKeyTableParams struct {
	// DNS Sec Table Key.
	ID string `json:"ID,omitempty"`
	// Usage count.
	UseCount uint64 `json:"UseCount,omitempty"`
	// KSK or ZSK.
	Type GslbCurDnsSecKeyTableType `json:"Type,omitempty"`
	// Key status.
	Status GslbCurDnsSecKeyTableStatus `json:"Status,omitempty"`
	// Size of key (bits).
	Size GslbCurDnsSecKeyTableSize `json:"Size,omitempty"`
	// Signing algorithm.
	Algo GslbCurDnsSecKeyTableAlgo `json:"Algo,omitempty"`
	// DNSSEC - TTL in Seconds.
	TTL uint64 `json:"TTL,omitempty"`
	// Expiration period in Seconds.
	ExpPeriod uint32 `json:"ExpPeriod,omitempty"`
	// RollOver period in Seconds.
	RollOverPeriod uint32 `json:"RollOverPeriod,omitempty"`
	// Signature validity period in Seconds.
	ValidityPeriod uint32 `json:"ValidityPeriod,omitempty"`
	// Signature publication period in Seconds.
	PublicationPeriod uint32 `json:"PublicationPeriod,omitempty"`
}

func (p GslbCurDnsSecKeyTableParams) iMABean() {}
