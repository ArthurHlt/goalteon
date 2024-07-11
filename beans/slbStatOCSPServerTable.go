package beans

import (
	"fmt"
	"reflect"
)

// SlbStatOCSPServerTable The statistics table of OCSP servers.
// Note:This mib is not supported for VX instance of virtualization.
type SlbStatOCSPServerTable struct {
	// OCSP Servers Address.
	SlbStatOCSPServerIndex int32
	Params                 *SlbStatOCSPServerTableParams
}

func NewSlbStatOCSPServerTableList() *SlbStatOCSPServerTable {
	return &SlbStatOCSPServerTable{}
}

func NewSlbStatOCSPServerTable(
	slbStatOCSPServerIndex int32,
	params *SlbStatOCSPServerTableParams,
) *SlbStatOCSPServerTable {
	return &SlbStatOCSPServerTable{
		SlbStatOCSPServerIndex: slbStatOCSPServerIndex,
		Params:                 params,
	}
}

func (c *SlbStatOCSPServerTable) Name() string {
	return "SlbStatOCSPServerTable"
}

func (c *SlbStatOCSPServerTable) GetParams() BeanType {
	return c.Params
}

func (c *SlbStatOCSPServerTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *SlbStatOCSPServerTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.SlbStatOCSPServerIndex).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.SlbStatOCSPServerIndex)
}

type SlbStatOCSPServerTableParams struct {
	// OCSP Servers Address.
	Index int32 `json:"Index,omitempty"`
	// OCSP Servers statistic.
	Address string `json:"Address,omitempty"`
	// OCSP Servers statistic total sent req.
	SentReq uint64 `json:"SentReq,omitempty"`
	// OCSP Servers statistic total failed retiries.
	FailedRetries uint64 `json:"FailedRetries,omitempty"`
	// OCSP Servers statistic total failed requests.
	FailedReq uint64 `json:"FailedReq,omitempty"`
	// OCSP Servers statistic total bad response.
	BadResponse uint64 `json:"BadResponse,omitempty"`
	// OCSP Servers statistic total revoked responses.
	RevokedResponse uint64 `json:"RevokedResponse,omitempty"`
	// OCSP Servers statistic total unknown responses.
	UnknownResponse uint64 `json:"UnknownResponse,omitempty"`
	// OCSP Servers statistic total invalid algorithm failed.
	InvalidAlgorithm uint64 `json:"InvalidAlgorithm,omitempty"`
	// OCSP Servers statistic total invalid time failed.
	InvalidTime uint64 `json:"InvalidTime,omitempty"`
	// OCSP Servers statistic total invalid nonce failed.
	InvalidNonce uint64 `json:"InvalidNonce,omitempty"`
	// OCSP Servers statistic total irrelevant responses.
	IrrelevantResponse uint64 `json:"IrrelevantResponse,omitempty"`
	// OCSP Servers statistic total invalid signutre responses.
	InvalidSignature uint64 `json:"InvalidSignature,omitempty"`
	// OCSP Servers statistic total success.
	Success uint64 `json:"Success,omitempty"`
	// OCSP Servers statistic total general failure responses.
	GeneralFailure uint64 `json:"GeneralFailure,omitempty"`
	// OCSP Servers statistic 1- POST 2 - GET.
	RequestMethod int32 `json:"RequestMethod,omitempty"`
}

func (p SlbStatOCSPServerTableParams) iMABean() {}
