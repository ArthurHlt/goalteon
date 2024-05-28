package beans

import (
	"fmt"
	"reflect"
)

// CompStatPerServTable A table for compression statistics per virtual service.
// Note:This mib is not supported for VX instance of Virtualization.
type CompStatPerServTable struct {
	// Virtual server Index.
	CompStatPerServVirtServIndex string
	// Virtual server service index.
	CompStatPerServVirtServiceIndex int32
	Params                          *CompStatPerServTableParams
}

func NewCompStatPerServTableList() *CompStatPerServTable {
	return &CompStatPerServTable{}
}

func NewCompStatPerServTable(
	compStatPerServVirtServIndex string,
	compStatPerServVirtServiceIndex int32,
	params *CompStatPerServTableParams,
) *CompStatPerServTable {
	return &CompStatPerServTable{
		CompStatPerServVirtServIndex:    compStatPerServVirtServIndex,
		CompStatPerServVirtServiceIndex: compStatPerServVirtServiceIndex,
		Params:                          params,
	}
}

func (c *CompStatPerServTable) Name() string {
	return "CompStatPerServTable"
}

func (c *CompStatPerServTable) GetParams() BeanType {
	return c.Params
}

func (c *CompStatPerServTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *CompStatPerServTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.CompStatPerServVirtServIndex).IsZero() &&
		reflect.ValueOf(c.CompStatPerServVirtServiceIndex).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.CompStatPerServVirtServIndex) + "/" + fmt.Sprint(c.CompStatPerServVirtServiceIndex)
}

type CompStatPerServTableParams struct {
	// Virtual server Index.
	VirtServIndex string `json:"VirtServIndex,omitempty"`
	// Virtual server service index.
	VirtServiceIndex int32 `json:"VirtServiceIndex,omitempty"`
	// Virtual server service port number.
	VirtServPort int32 `json:"VirtServPort,omitempty"`
	// Compression policy identifier associated with the virtual service.
	CompPolId string `json:"CompPolId,omitempty"`
	// Total throughput of compressible object before compression counted during measuring period for virtual service.
	UnComprTputKb int32 `json:"UnComprTputKb,omitempty"`
	// Total throughput of compressible objects after compression counted during measuring period for virtual service.
	ComprTputKb int32 `json:"ComprTputKb,omitempty"`
	// Average object size before compression for virtual service.
	AvgSizeBefComp int32 `json:"AvgSizeBefComp,omitempty"`
	// Average object size after compression for virtual service.
	AvgSizeAftComp int32 `json:"AvgSizeAftComp,omitempty"`
	// Average compression ratio during measuring period for virtual service.
	AvgCompRatio int32 `json:"AvgCompRatio,omitempty"`
	// Effective compression ratio out of total traffic for virtual service.
	ThrputCompRatio int32 `json:"ThrputCompRatio,omitempty"`
	// bytes saved for virtual service.
	BytesSaved uint64 `json:"BytesSaved,omitempty"`
	// Peak bytes saved for virtual service.
	BytesSavedPeak uint64 `json:"BytesSavedPeak,omitempty"`
	// Total bytes saved for virtual service.
	BytesSavedTot uint64 `json:"BytesSavedTot,omitempty"`
}

func (p CompStatPerServTableParams) iMABean() {}
