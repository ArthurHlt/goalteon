package beans

import (
	"fmt"
	"reflect"
)

// SlbNewSslCfgAuthPolTable The table for configuring Client Authentication  policy.
// Note:This mib is not supported for VX instance of Virtualization.
type SlbNewSslCfgAuthPolTable struct {
	// The Auth policy name(key id) as an index.
	SlbNewSslCfgAuthPolNameIdIndex string
	Params                         *SlbNewSslCfgAuthPolTableParams
}

func NewSlbNewSslCfgAuthPolTableList() *SlbNewSslCfgAuthPolTable {
	return &SlbNewSslCfgAuthPolTable{}
}

func NewSlbNewSslCfgAuthPolTable(
	slbNewSslCfgAuthPolNameIdIndex string,
	params *SlbNewSslCfgAuthPolTableParams,
) *SlbNewSslCfgAuthPolTable {
	return &SlbNewSslCfgAuthPolTable{
		SlbNewSslCfgAuthPolNameIdIndex: slbNewSslCfgAuthPolNameIdIndex,
		Params:                         params,
	}
}

func (c *SlbNewSslCfgAuthPolTable) Name() string {
	return "SlbNewSslCfgAuthPolTable"
}

func (c *SlbNewSslCfgAuthPolTable) GetParams() BeanType {
	return c.Params
}

func (c *SlbNewSslCfgAuthPolTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *SlbNewSslCfgAuthPolTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.SlbNewSslCfgAuthPolNameIdIndex).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.SlbNewSslCfgAuthPolNameIdIndex)
}

type SlbNewSslCfgAuthPolTableValidityUriprior int32

const (
	SlbNewSslCfgAuthPolTableValidityUriprior_Clientcert  SlbNewSslCfgAuthPolTableValidityUriprior = 1
	SlbNewSslCfgAuthPolTableValidityUriprior_Staticuri   SlbNewSslCfgAuthPolTableValidityUriprior = 2
	SlbNewSslCfgAuthPolTableValidityUriprior_Unsupported SlbNewSslCfgAuthPolTableValidityUriprior = 2147483647
)

type SlbNewSslCfgAuthPolTableValidityVchain int32

const (
	SlbNewSslCfgAuthPolTableValidityVchain_Enabled     SlbNewSslCfgAuthPolTableValidityVchain = 1
	SlbNewSslCfgAuthPolTableValidityVchain_Disabled    SlbNewSslCfgAuthPolTableValidityVchain = 2
	SlbNewSslCfgAuthPolTableValidityVchain_Unsupported SlbNewSslCfgAuthPolTableValidityVchain = 2147483647
)

type SlbNewSslCfgAuthPolTableValiditySecure int32

const (
	SlbNewSslCfgAuthPolTableValiditySecure_Enabled     SlbNewSslCfgAuthPolTableValiditySecure = 1
	SlbNewSslCfgAuthPolTableValiditySecure_Disabled    SlbNewSslCfgAuthPolTableValiditySecure = 2
	SlbNewSslCfgAuthPolTableValiditySecure_Unsupported SlbNewSslCfgAuthPolTableValiditySecure = 2147483647
)

type SlbNewSslCfgAuthPolTablePassinfoVersionFlag int32

const (
	SlbNewSslCfgAuthPolTablePassinfoVersionFlag_Enabled     SlbNewSslCfgAuthPolTablePassinfoVersionFlag = 1
	SlbNewSslCfgAuthPolTablePassinfoVersionFlag_Disabled    SlbNewSslCfgAuthPolTablePassinfoVersionFlag = 2
	SlbNewSslCfgAuthPolTablePassinfoVersionFlag_Unsupported SlbNewSslCfgAuthPolTablePassinfoVersionFlag = 2147483647
)

type SlbNewSslCfgAuthPolTablePassinfoSerialFlag int32

const (
	SlbNewSslCfgAuthPolTablePassinfoSerialFlag_Enabled     SlbNewSslCfgAuthPolTablePassinfoSerialFlag = 1
	SlbNewSslCfgAuthPolTablePassinfoSerialFlag_Disabled    SlbNewSslCfgAuthPolTablePassinfoSerialFlag = 2
	SlbNewSslCfgAuthPolTablePassinfoSerialFlag_Unsupported SlbNewSslCfgAuthPolTablePassinfoSerialFlag = 2147483647
)

type SlbNewSslCfgAuthPolTablePassinfoAlgoFlag int32

const (
	SlbNewSslCfgAuthPolTablePassinfoAlgoFlag_Enabled     SlbNewSslCfgAuthPolTablePassinfoAlgoFlag = 1
	SlbNewSslCfgAuthPolTablePassinfoAlgoFlag_Disabled    SlbNewSslCfgAuthPolTablePassinfoAlgoFlag = 2
	SlbNewSslCfgAuthPolTablePassinfoAlgoFlag_Unsupported SlbNewSslCfgAuthPolTablePassinfoAlgoFlag = 2147483647
)

type SlbNewSslCfgAuthPolTablePassinfoIssuerFlag int32

const (
	SlbNewSslCfgAuthPolTablePassinfoIssuerFlag_Enabled     SlbNewSslCfgAuthPolTablePassinfoIssuerFlag = 1
	SlbNewSslCfgAuthPolTablePassinfoIssuerFlag_Disabled    SlbNewSslCfgAuthPolTablePassinfoIssuerFlag = 2
	SlbNewSslCfgAuthPolTablePassinfoIssuerFlag_Unsupported SlbNewSslCfgAuthPolTablePassinfoIssuerFlag = 2147483647
)

type SlbNewSslCfgAuthPolTablePassinfoNbeforeFlag int32

const (
	SlbNewSslCfgAuthPolTablePassinfoNbeforeFlag_Enabled     SlbNewSslCfgAuthPolTablePassinfoNbeforeFlag = 1
	SlbNewSslCfgAuthPolTablePassinfoNbeforeFlag_Disabled    SlbNewSslCfgAuthPolTablePassinfoNbeforeFlag = 2
	SlbNewSslCfgAuthPolTablePassinfoNbeforeFlag_Unsupported SlbNewSslCfgAuthPolTablePassinfoNbeforeFlag = 2147483647
)

type SlbNewSslCfgAuthPolTablePassinfoNafterFlag int32

const (
	SlbNewSslCfgAuthPolTablePassinfoNafterFlag_Enabled     SlbNewSslCfgAuthPolTablePassinfoNafterFlag = 1
	SlbNewSslCfgAuthPolTablePassinfoNafterFlag_Disabled    SlbNewSslCfgAuthPolTablePassinfoNafterFlag = 2
	SlbNewSslCfgAuthPolTablePassinfoNafterFlag_Unsupported SlbNewSslCfgAuthPolTablePassinfoNafterFlag = 2147483647
)

type SlbNewSslCfgAuthPolTablePassinfoSubjectFlag int32

const (
	SlbNewSslCfgAuthPolTablePassinfoSubjectFlag_Enabled     SlbNewSslCfgAuthPolTablePassinfoSubjectFlag = 1
	SlbNewSslCfgAuthPolTablePassinfoSubjectFlag_Disabled    SlbNewSslCfgAuthPolTablePassinfoSubjectFlag = 2
	SlbNewSslCfgAuthPolTablePassinfoSubjectFlag_Unsupported SlbNewSslCfgAuthPolTablePassinfoSubjectFlag = 2147483647
)

type SlbNewSslCfgAuthPolTablePassinfoKeytypeFlag int32

const (
	SlbNewSslCfgAuthPolTablePassinfoKeytypeFlag_Enabled     SlbNewSslCfgAuthPolTablePassinfoKeytypeFlag = 1
	SlbNewSslCfgAuthPolTablePassinfoKeytypeFlag_Disabled    SlbNewSslCfgAuthPolTablePassinfoKeytypeFlag = 2
	SlbNewSslCfgAuthPolTablePassinfoKeytypeFlag_Unsupported SlbNewSslCfgAuthPolTablePassinfoKeytypeFlag = 2147483647
)

type SlbNewSslCfgAuthPolTablePassinfoMd5Flag int32

const (
	SlbNewSslCfgAuthPolTablePassinfoMd5Flag_Enabled     SlbNewSslCfgAuthPolTablePassinfoMd5Flag = 1
	SlbNewSslCfgAuthPolTablePassinfoMd5Flag_Disabled    SlbNewSslCfgAuthPolTablePassinfoMd5Flag = 2
	SlbNewSslCfgAuthPolTablePassinfoMd5Flag_Unsupported SlbNewSslCfgAuthPolTablePassinfoMd5Flag = 2147483647
)

type SlbNewSslCfgAuthPolTablePassinfoCertFlag int32

const (
	SlbNewSslCfgAuthPolTablePassinfoCertFlag_Enabled     SlbNewSslCfgAuthPolTablePassinfoCertFlag = 1
	SlbNewSslCfgAuthPolTablePassinfoCertFlag_Disabled    SlbNewSslCfgAuthPolTablePassinfoCertFlag = 2
	SlbNewSslCfgAuthPolTablePassinfoCertFlag_Unsupported SlbNewSslCfgAuthPolTablePassinfoCertFlag = 2147483647
)

type SlbNewSslCfgAuthPolTableAdminStatus int32

const (
	SlbNewSslCfgAuthPolTableAdminStatus_Enabled     SlbNewSslCfgAuthPolTableAdminStatus = 1
	SlbNewSslCfgAuthPolTableAdminStatus_Disabled    SlbNewSslCfgAuthPolTableAdminStatus = 2
	SlbNewSslCfgAuthPolTableAdminStatus_Unsupported SlbNewSslCfgAuthPolTableAdminStatus = 2147483647
)

type SlbNewSslCfgAuthPolTableDelete int32

const (
	SlbNewSslCfgAuthPolTableDelete_Other       SlbNewSslCfgAuthPolTableDelete = 1
	SlbNewSslCfgAuthPolTableDelete_Delete      SlbNewSslCfgAuthPolTableDelete = 2
	SlbNewSslCfgAuthPolTableDelete_Unsupported SlbNewSslCfgAuthPolTableDelete = 2147483647
)

type SlbNewSslCfgAuthPolTablePassinfoComp2424 int32

const (
	SlbNewSslCfgAuthPolTablePassinfoComp2424_Enabled     SlbNewSslCfgAuthPolTablePassinfoComp2424 = 1
	SlbNewSslCfgAuthPolTablePassinfoComp2424_Disabled    SlbNewSslCfgAuthPolTablePassinfoComp2424 = 2
	SlbNewSslCfgAuthPolTablePassinfoComp2424_Unsupported SlbNewSslCfgAuthPolTablePassinfoComp2424 = 2147483647
)

type SlbNewSslCfgAuthPolTableValidityCdpPreference int32

const (
	SlbNewSslCfgAuthPolTableValidityCdpPreference_Embedded    SlbNewSslCfgAuthPolTableValidityCdpPreference = 1
	SlbNewSslCfgAuthPolTableValidityCdpPreference_Userdefined SlbNewSslCfgAuthPolTableValidityCdpPreference = 2
	SlbNewSslCfgAuthPolTableValidityCdpPreference_Unsupported SlbNewSslCfgAuthPolTableValidityCdpPreference = 2147483647
)

type SlbNewSslCfgAuthPolTableType int32

const (
	SlbNewSslCfgAuthPolTableType_Frontend    SlbNewSslCfgAuthPolTableType = 1
	SlbNewSslCfgAuthPolTableType_Backend     SlbNewSslCfgAuthPolTableType = 2
	SlbNewSslCfgAuthPolTableType_Unsupported SlbNewSslCfgAuthPolTableType = 2147483647
)

type SlbNewSslCfgAuthPolTableSerCertExp int32

const (
	SlbNewSslCfgAuthPolTableSerCertExp_Ignore      SlbNewSslCfgAuthPolTableSerCertExp = 1
	SlbNewSslCfgAuthPolTableSerCertExp_Reject      SlbNewSslCfgAuthPolTableSerCertExp = 2
	SlbNewSslCfgAuthPolTableSerCertExp_Unsupported SlbNewSslCfgAuthPolTableSerCertExp = 2147483647
)

type SlbNewSslCfgAuthPolTableSerCertMis int32

const (
	SlbNewSslCfgAuthPolTableSerCertMis_Ignore      SlbNewSslCfgAuthPolTableSerCertMis = 1
	SlbNewSslCfgAuthPolTableSerCertMis_Reject      SlbNewSslCfgAuthPolTableSerCertMis = 2
	SlbNewSslCfgAuthPolTableSerCertMis_Unsupported SlbNewSslCfgAuthPolTableSerCertMis = 2147483647
)

type SlbNewSslCfgAuthPolTableSerCertUntrust int32

const (
	SlbNewSslCfgAuthPolTableSerCertUntrust_Ignore      SlbNewSslCfgAuthPolTableSerCertUntrust = 1
	SlbNewSslCfgAuthPolTableSerCertUntrust_Reject      SlbNewSslCfgAuthPolTableSerCertUntrust = 2
	SlbNewSslCfgAuthPolTableSerCertUntrust_Unsupported SlbNewSslCfgAuthPolTableSerCertUntrust = 2147483647
)

type SlbNewSslCfgAuthPolTableIssuerNamesOrder int32

const (
	SlbNewSslCfgAuthPolTableIssuerNamesOrder_Regular     SlbNewSslCfgAuthPolTableIssuerNamesOrder = 0
	SlbNewSslCfgAuthPolTableIssuerNamesOrder_Reverse     SlbNewSslCfgAuthPolTableIssuerNamesOrder = 1
	SlbNewSslCfgAuthPolTableIssuerNamesOrder_Unsupported SlbNewSslCfgAuthPolTableIssuerNamesOrder = 2147483647
)

type SlbNewSslCfgAuthPolTableSubjectNamesOrder int32

const (
	SlbNewSslCfgAuthPolTableSubjectNamesOrder_Regular     SlbNewSslCfgAuthPolTableSubjectNamesOrder = 0
	SlbNewSslCfgAuthPolTableSubjectNamesOrder_Reverse     SlbNewSslCfgAuthPolTableSubjectNamesOrder = 1
	SlbNewSslCfgAuthPolTableSubjectNamesOrder_Unsupported SlbNewSslCfgAuthPolTableSubjectNamesOrder = 2147483647
)

type SlbNewSslCfgAuthPolTableValidityOcspMode int32

const (
	SlbNewSslCfgAuthPolTableValidityOcspMode_Ocsp        SlbNewSslCfgAuthPolTableValidityOcspMode = 0
	SlbNewSslCfgAuthPolTableValidityOcspMode_Stapling    SlbNewSslCfgAuthPolTableValidityOcspMode = 1
	SlbNewSslCfgAuthPolTableValidityOcspMode_Both        SlbNewSslCfgAuthPolTableValidityOcspMode = 2
	SlbNewSslCfgAuthPolTableValidityOcspMode_Unsupported SlbNewSslCfgAuthPolTableValidityOcspMode = 2147483647
)

type SlbNewSslCfgAuthPolTableValidityStaplingUriprior int32

const (
	SlbNewSslCfgAuthPolTableValidityStaplingUriprior_Clientcert  SlbNewSslCfgAuthPolTableValidityStaplingUriprior = 1
	SlbNewSslCfgAuthPolTableValidityStaplingUriprior_Staticuri   SlbNewSslCfgAuthPolTableValidityStaplingUriprior = 2
	SlbNewSslCfgAuthPolTableValidityStaplingUriprior_Unsupported SlbNewSslCfgAuthPolTableValidityStaplingUriprior = 2147483647
)

type SlbNewSslCfgAuthPolTableParams struct {
	// The Auth policy name(key id) as an index.
	NameIdIndex string `json:"NameIdIndex,omitempty"`
	// Auth policy name.
	Name string `json:"Name,omitempty"`
	// Certificate validation check method.
	ValidityMethod int32 `json:"ValidityMethod,omitempty"`
	// Static URI for OCSP validation requests.
	ValidityStaturi string `json:"ValidityStaturi,omitempty"`
	// OCSP URI priority.
	ValidityUriprior SlbNewSslCfgAuthPolTableValidityUriprior `json:"ValidityUriprior,omitempty"`
	// OCSP response cache time.
	ValidityCachtime uint32 `json:"ValidityCachtime,omitempty"`
	// OCSP response time deviation.
	ValidityTimedev uint32 `json:"ValidityTimedev,omitempty"`
	// Allowed signing algorithm for the OCSP response [all, md5, sha1, sha256, sha384, sha512].
	ValidityAlgorthmName string `json:"ValidityAlgorthmName,omitempty"`
	// Enable/Disable validating every CA certificate in the CA chain.
	ValidityVchain SlbNewSslCfgAuthPolTableValidityVchain `json:"ValidityVchain,omitempty"`
	// Enable/Disable secure OCSP response by sending random nonce.
	ValiditySecure SlbNewSslCfgAuthPolTableValiditySecure `json:"ValiditySecure,omitempty"`
	// Certificate version header.
	PassinfoVersionName string `json:"PassinfoVersionName,omitempty"`
	// Pass certificate version information to backend server.
	PassinfoVersionFlag SlbNewSslCfgAuthPolTablePassinfoVersionFlag `json:"PassinfoVersionFlag,omitempty"`
	// Certificate serial-number header.
	PassinfoSerialName string `json:"PassinfoSerialName,omitempty"`
	// Pass certificate serial-number to backend server
	PassinfoSerialFlag SlbNewSslCfgAuthPolTablePassinfoSerialFlag `json:"PassinfoSerialFlag,omitempty"`
	// Certificate signature algorithm header name.
	PassinfoAlgoName string `json:"PassinfoAlgoName,omitempty"`
	// Pass certificate Signature Algorithm to backend server
	PassinfoAlgoFlag SlbNewSslCfgAuthPolTablePassinfoAlgoFlag `json:"PassinfoAlgoFlag,omitempty"`
	// Certificate issuer header.
	PassinfoIssuerName string `json:"PassinfoIssuerName,omitempty"`
	// Pass certificate issuer information to backend server
	PassinfoIssuerFlag SlbNewSslCfgAuthPolTablePassinfoIssuerFlag `json:"PassinfoIssuerFlag,omitempty"`
	// Certificate 'Not Before Validity Dates' header.
	PassinfoNbeforeName string `json:"PassinfoNbeforeName,omitempty"`
	// Pass certificate 'Not Before' Validity Date to backend server
	PassinfoNbeforeFlag SlbNewSslCfgAuthPolTablePassinfoNbeforeFlag `json:"PassinfoNbeforeFlag,omitempty"`
	// Certificate 'Not After Validity Dates' header.
	PassinfoNafterName string `json:"PassinfoNafterName,omitempty"`
	// Pass certificate 'Not After' Validity Date to backend server
	PassinfoNafterFlag SlbNewSslCfgAuthPolTablePassinfoNafterFlag `json:"PassinfoNafterFlag,omitempty"`
	// Certificate subject header name.
	PassinfoSubjectName string `json:"PassinfoSubjectName,omitempty"`
	// Pass certificate subject to backend server
	PassinfoSubjectFlag SlbNewSslCfgAuthPolTablePassinfoSubjectFlag `json:"PassinfoSubjectFlag,omitempty"`
	// Certificate Public Key Type header.
	PassinfoKeytypeName string `json:"PassinfoKeytypeName,omitempty"`
	// Pass certificate Public Key Type information to backend servers
	PassinfoKeytypeFlag SlbNewSslCfgAuthPolTablePassinfoKeytypeFlag `json:"PassinfoKeytypeFlag,omitempty"`
	// Certificate MD5 Hash header.
	PassinfoMd5Name string `json:"PassinfoMd5Name,omitempty"`
	// Pass certificate MD5 Hash information to backend servers
	PassinfoMd5Flag SlbNewSslCfgAuthPolTablePassinfoMd5Flag `json:"PassinfoMd5Flag,omitempty"`
	// Certificate header.
	PassinfoCertName string `json:"PassinfoCertName,omitempty"`
	// Certificate Header Lines Format.
	PassinfoCertFormat int32 `json:"PassinfoCertFormat,omitempty"`
	// Pass certificate information to backend servers
	PassinfoCertFlag SlbNewSslCfgAuthPolTablePassinfoCertFlag `json:"PassinfoCertFlag,omitempty"`
	// Set the character set to be used for information.
	PassinfoCharset int32 `json:"PassinfoCharset,omitempty"`
	// Trusted client's CA certificate name.
	TrustcaChainName string `json:"TrustcaChainName,omitempty"`
	// Trusted client's CA certificate type certificate=cert,Group=group,None=empty string.
	TrustcaChainType string `json:"TrustcaChainType,omitempty"`
	// Maximum depth to search the trusted CA in the CA certificate.
	Cadepth int32 `json:"Cadepth,omitempty"`
	// Certificate's CA verification level.
	Caverify int32 `json:"Caverify,omitempty"`
	// URL for redirection when client authentication fails
	Failurl string `json:"Failurl,omitempty"`
	// Enable or disable Auth policy.
	AdminStatus SlbNewSslCfgAuthPolTableAdminStatus `json:"AdminStatus,omitempty"`
	// Delete Client Authentication Policy.
	Delete SlbNewSslCfgAuthPolTableDelete `json:"Delete,omitempty"`
	// Enable/Disable the 2424SSL Headers Compliance Mode.
	PassinfoComp2424 SlbNewSslCfgAuthPolTablePassinfoComp2424 `json:"PassinfoComp2424,omitempty"`
	// File Name for CRL validation requests.
	ValidityCrlFile string `json:"ValidityCrlFile,omitempty"`
	// Group for CDP validation requests.
	ValidityCdpGroup string `json:"ValidityCdpGroup,omitempty"`
	// CDP General grace time.
	ValidityCdpGracetime uint32 `json:"ValidityCdpGracetime,omitempty"`
	// CDP update interval.
	ValidityCdpInterval uint64 `json:"ValidityCdpInterval,omitempty"`
	// CDP preference.
	ValidityCdpPreference SlbNewSslCfgAuthPolTableValidityCdpPreference `json:"ValidityCdpPreference,omitempty"`
	// Client CA for Request certificate or group name.
	ClientcaReqChainName string `json:"ClientcaReqChainName,omitempty"`
	// Client CA for Request type Certificate=cert,Group=group,None=none,Default=default.
	ClientcaReqChainType string `json:"ClientcaReqChainType,omitempty"`
	// Authentication policy type (frontend/backend).
	Type SlbNewSslCfgAuthPolTableType `json:"Type,omitempty"`
	// Authentication policy action on expired cert (ignore/reject).
	SerCertExp SlbNewSslCfgAuthPolTableSerCertExp `json:"SerCertExp,omitempty"`
	// Authentication policy action on mismatch (ignore/reject).
	SerCertMis SlbNewSslCfgAuthPolTableSerCertMis `json:"SerCertMis,omitempty"`
	// Authentication policy action on untrusted (ignore/reject).
	SerCertUntrust SlbNewSslCfgAuthPolTableSerCertUntrust `json:"SerCertUntrust,omitempty"`
	// Set the order in which the issuer names will be passed
	IssuerNamesOrder SlbNewSslCfgAuthPolTableIssuerNamesOrder `json:"IssuerNamesOrder,omitempty"`
	// Set the order in which the subject names will be passed
	SubjectNamesOrder SlbNewSslCfgAuthPolTableSubjectNamesOrder `json:"SubjectNamesOrder,omitempty"`
	// OCSP mode.
	ValidityOcspMode SlbNewSslCfgAuthPolTableValidityOcspMode `json:"ValidityOcspMode,omitempty"`
	// Static URI for OCSP Stapling validation requests.
	ValidityStaplinguri string `json:"ValidityStaplinguri,omitempty"`
	// OCSP Stapling URI priority.
	ValidityStaplingUriprior SlbNewSslCfgAuthPolTableValidityStaplingUriprior `json:"ValidityStaplingUriprior,omitempty"`
	// Static secondary URI for OCSP validation requests.
	ValidityStatSecUri string `json:"ValidityStatSecUri,omitempty"`
	// Static secondary URI for OCSP Stapling validation requests.
	ValidityStaplingSecUri string `json:"ValidityStaplingSecUri,omitempty"`
	// Number of retries reaching the OCSP server.
	ValidityOcspRetries uint64 `json:"ValidityOcspRetries,omitempty"`
}

func (p SlbNewSslCfgAuthPolTableParams) iMABean() {}
