package beans

import (
	"fmt"
	"reflect"
)

// SlbCurSslCfgAuthPolTable The table for configuring Client Authentication  policy.
// Note:This mib is not supported for VX instance of Virtualization.
type SlbCurSslCfgAuthPolTable struct {
	// The Auth policy name(key id) as an index.
	SlbCurSslCfgAuthPolNameIdIndex string
	Params                         *SlbCurSslCfgAuthPolTableParams
}

func NewSlbCurSslCfgAuthPolTableList() *SlbCurSslCfgAuthPolTable {
	return &SlbCurSslCfgAuthPolTable{}
}

func NewSlbCurSslCfgAuthPolTable(
	slbCurSslCfgAuthPolNameIdIndex string,
	params *SlbCurSslCfgAuthPolTableParams,
) *SlbCurSslCfgAuthPolTable {
	return &SlbCurSslCfgAuthPolTable{
		SlbCurSslCfgAuthPolNameIdIndex: slbCurSslCfgAuthPolNameIdIndex,
		Params:                         params,
	}
}

func (c *SlbCurSslCfgAuthPolTable) Name() string {
	return "SlbCurSslCfgAuthPolTable"
}

func (c *SlbCurSslCfgAuthPolTable) GetParams() BeanType {
	return c.Params
}

func (c *SlbCurSslCfgAuthPolTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *SlbCurSslCfgAuthPolTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.SlbCurSslCfgAuthPolNameIdIndex).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.SlbCurSslCfgAuthPolNameIdIndex)
}

type SlbCurSslCfgAuthPolTableValidityUriprior int32

const (
	SlbCurSslCfgAuthPolTableValidityUriprior_Clientcert  SlbCurSslCfgAuthPolTableValidityUriprior = 1
	SlbCurSslCfgAuthPolTableValidityUriprior_Staticuri   SlbCurSslCfgAuthPolTableValidityUriprior = 2
	SlbCurSslCfgAuthPolTableValidityUriprior_Unsupported SlbCurSslCfgAuthPolTableValidityUriprior = 2147483647
)

type SlbCurSslCfgAuthPolTableValidityVchain int32

const (
	SlbCurSslCfgAuthPolTableValidityVchain_Enabled     SlbCurSslCfgAuthPolTableValidityVchain = 1
	SlbCurSslCfgAuthPolTableValidityVchain_Disabled    SlbCurSslCfgAuthPolTableValidityVchain = 2
	SlbCurSslCfgAuthPolTableValidityVchain_Unsupported SlbCurSslCfgAuthPolTableValidityVchain = 2147483647
)

type SlbCurSslCfgAuthPolTableValiditySecure int32

const (
	SlbCurSslCfgAuthPolTableValiditySecure_Enabled     SlbCurSslCfgAuthPolTableValiditySecure = 1
	SlbCurSslCfgAuthPolTableValiditySecure_Disabled    SlbCurSslCfgAuthPolTableValiditySecure = 2
	SlbCurSslCfgAuthPolTableValiditySecure_Unsupported SlbCurSslCfgAuthPolTableValiditySecure = 2147483647
)

type SlbCurSslCfgAuthPolTableAdminStatus int32

const (
	SlbCurSslCfgAuthPolTableAdminStatus_Enabled     SlbCurSslCfgAuthPolTableAdminStatus = 1
	SlbCurSslCfgAuthPolTableAdminStatus_Disabled    SlbCurSslCfgAuthPolTableAdminStatus = 2
	SlbCurSslCfgAuthPolTableAdminStatus_Unsupported SlbCurSslCfgAuthPolTableAdminStatus = 2147483647
)

type SlbCurSslCfgAuthPolTablePassinfoComp2424 int32

const (
	SlbCurSslCfgAuthPolTablePassinfoComp2424_Enabled     SlbCurSslCfgAuthPolTablePassinfoComp2424 = 1
	SlbCurSslCfgAuthPolTablePassinfoComp2424_Disabled    SlbCurSslCfgAuthPolTablePassinfoComp2424 = 2
	SlbCurSslCfgAuthPolTablePassinfoComp2424_Unsupported SlbCurSslCfgAuthPolTablePassinfoComp2424 = 2147483647
)

type SlbCurSslCfgAuthPolTableValidityCdpPreference int32

const (
	SlbCurSslCfgAuthPolTableValidityCdpPreference_Embedded    SlbCurSslCfgAuthPolTableValidityCdpPreference = 1
	SlbCurSslCfgAuthPolTableValidityCdpPreference_Userdefined SlbCurSslCfgAuthPolTableValidityCdpPreference = 2
	SlbCurSslCfgAuthPolTableValidityCdpPreference_Unsupported SlbCurSslCfgAuthPolTableValidityCdpPreference = 2147483647
)

type SlbCurSslCfgAuthPolTableValidityOcspMode int32

const (
	SlbCurSslCfgAuthPolTableValidityOcspMode_Ocsp        SlbCurSslCfgAuthPolTableValidityOcspMode = 0
	SlbCurSslCfgAuthPolTableValidityOcspMode_Stapling    SlbCurSslCfgAuthPolTableValidityOcspMode = 1
	SlbCurSslCfgAuthPolTableValidityOcspMode_Both        SlbCurSslCfgAuthPolTableValidityOcspMode = 2
	SlbCurSslCfgAuthPolTableValidityOcspMode_Unsupported SlbCurSslCfgAuthPolTableValidityOcspMode = 2147483647
)

type SlbCurSslCfgAuthPolTableValidityStaplingUriprior int32

const (
	SlbCurSslCfgAuthPolTableValidityStaplingUriprior_Clientcert  SlbCurSslCfgAuthPolTableValidityStaplingUriprior = 1
	SlbCurSslCfgAuthPolTableValidityStaplingUriprior_Staticuri   SlbCurSslCfgAuthPolTableValidityStaplingUriprior = 2
	SlbCurSslCfgAuthPolTableValidityStaplingUriprior_Unsupported SlbCurSslCfgAuthPolTableValidityStaplingUriprior = 2147483647
)

type SlbCurSslCfgAuthPolTableParams struct {
	// The Auth policy name(key id) as an index.
	NameIdIndex string `json:"NameIdIndex,omitempty"`
	// Auth policy name.
	Name string `json:"Name,omitempty"`
	// Certificate validation check method.
	ValidityMethod int32 `json:"ValidityMethod,omitempty"`
	// Static URI for OCSP validation requests.
	ValidityStaturi string `json:"ValidityStaturi,omitempty"`
	// OCSP URI priority.
	ValidityUriprior SlbCurSslCfgAuthPolTableValidityUriprior `json:"ValidityUriprior,omitempty"`
	// OCSP response cache time.
	ValidityCachtime int32 `json:"ValidityCachtime,omitempty"`
	// OCSP response time deviation.
	ValidityTimedev int32 `json:"ValidityTimedev,omitempty"`
	// Allowed signing algorithm for the OCSP response[all, md5, sha1, sha256, sha384, sha512, sm2_with_sm3].
	ValidityAlgorthmName string `json:"ValidityAlgorthmName,omitempty"`
	// Enable/Disable validating every CA certificate in the CA chain.
	ValidityVchain SlbCurSslCfgAuthPolTableValidityVchain `json:"ValidityVchain,omitempty"`
	// Enable/Disable secure OCSP response by sending random nonce.
	ValiditySecure SlbCurSslCfgAuthPolTableValiditySecure `json:"ValiditySecure,omitempty"`
	// Certificate version header.
	PassinfoVersionName string `json:"PassinfoVersionName,omitempty"`
	// Pass certificate version information to backend server.
	PassinfoVersionFlag int32 `json:"PassinfoVersionFlag,omitempty"`
	// Certificate serial-number header .
	PassinfoSerialName string `json:"PassinfoSerialName,omitempty"`
	// Pass certificate serial-number to backend server.
	PassinfoSerialFlag int32 `json:"PassinfoSerialFlag,omitempty"`
	// Certificate signature algorithm header name.
	PassinfoAlgoName string `json:"PassinfoAlgoName,omitempty"`
	// Pass certificate Signature Algorithm to backend server
	PassinfoAlgoFlag int32 `json:"PassinfoAlgoFlag,omitempty"`
	// Certificate issuer header.
	PassinfoIssuerName string `json:"PassinfoIssuerName,omitempty"`
	// Pass certificate issuer information to backend server
	PassinfoIssuerFlag int32 `json:"PassinfoIssuerFlag,omitempty"`
	// Certificate 'Not Before Validity Dates' header.
	PassinfoNbeforeName string `json:"PassinfoNbeforeName,omitempty"`
	// Pass certificate 'Not Before' Validity Date to backend server
	PassinfoNbeforeFlag int32 `json:"PassinfoNbeforeFlag,omitempty"`
	// Certificate 'Not After Validity Dates' header.
	PassinfoNafterName string `json:"PassinfoNafterName,omitempty"`
	// Pass certificate 'Not After' Validity Date to backend server
	PassinfoNafterFlag int32 `json:"PassinfoNafterFlag,omitempty"`
	// Certificate subject header name.
	PassinfoSubjectName string `json:"PassinfoSubjectName,omitempty"`
	// Pass certificate subject to backend server
	PassinfoSubjectFlag int32 `json:"PassinfoSubjectFlag,omitempty"`
	// Certificate Public Key Type header.
	PassinfoKeytypeName string `json:"PassinfoKeytypeName,omitempty"`
	// Pass certificate Public Key Type information to backend servers
	PassinfoKeytypeFlag int32 `json:"PassinfoKeytypeFlag,omitempty"`
	// Certificate MD5 Hash header.
	PassinfoMd5Name string `json:"PassinfoMd5Name,omitempty"`
	// Pass certificate MD5 Hash information to backend servers
	PassinfoMd5Flag int32 `json:"PassinfoMd5Flag,omitempty"`
	// Certificate header.
	PassinfoCertName string `json:"PassinfoCertName,omitempty"`
	// Certificate Header Lines Format.
	PassinfoCertFormat int32 `json:"PassinfoCertFormat,omitempty"`
	// Pass certificate information to backend servers
	PassinfoCertFlag int32 `json:"PassinfoCertFlag,omitempty"`
	// Set the character set to be used for information.
	PassinfoCharset int32 `json:"PassinfoCharset,omitempty"`
	// Trusted client's CA certificate.
	TrustcaChainName string `json:"TrustcaChainName,omitempty"`
	// Tusted client's CA certificate type certificate=cert,Group=group,None=empty string.
	TrustcaChainType string `json:"TrustcaChainType,omitempty"`
	// Maximum depth to search the trusted CA in the CA certificate.
	Cadepth int32 `json:"Cadepth,omitempty"`
	// Certificate's CA verification level.
	Caverify int32 `json:"Caverify,omitempty"`
	// URL for redirection when client authentication fails
	Failurl string `json:"Failurl,omitempty"`
	// Enable or disable Auth policy.
	AdminStatus SlbCurSslCfgAuthPolTableAdminStatus `json:"AdminStatus,omitempty"`
	// Enable/Disable the 2424SSL Headers Compliance Mode.
	PassinfoComp2424 SlbCurSslCfgAuthPolTablePassinfoComp2424 `json:"PassinfoComp2424,omitempty"`
	// File Name for CRL validation requests.
	ValidityCrlFile string `json:"ValidityCrlFile,omitempty"`
	// Group for CDP validation requests.
	ValidityCdpGroup string `json:"ValidityCdpGroup,omitempty"`
	// CDP General grace time.
	ValidityCdpGracetime uint32 `json:"ValidityCdpGracetime,omitempty"`
	// CDP update interval.
	ValidityCdpInterval uint64 `json:"ValidityCdpInterval,omitempty"`
	// CDP preference.
	ValidityCdpPreference SlbCurSslCfgAuthPolTableValidityCdpPreference `json:"ValidityCdpPreference,omitempty"`
	// Client CA for Request certificate or group name.
	ClientcaReqChainName string `json:"ClientcaReqChainName,omitempty"`
	// Client CA for Request type Certificate=cert,Group=group,None=none,Default=default.
	ClientcaReqChainType string `json:"ClientcaReqChainType,omitempty"`
	// Authentication policy type (frontend/backend).
	Type int32 `json:"Type,omitempty"`
	// Authentication policy action on expired cert (ignore/reject).
	SerCertExp int32 `json:"SerCertExp,omitempty"`
	// Authentication policy action on mismatch (ignore/reject).
	SerCertMis int32 `json:"SerCertMis,omitempty"`
	// Authentication policy action on untrusted (ignore/reject).
	SerCertUntrust int32 `json:"SerCertUntrust,omitempty"`
	// Set the order in which the issuer names will be passed
	IssuerNamesOrder int32 `json:"IssuerNamesOrder,omitempty"`
	// Set the order in which the subject names will be passed
	SubjectNamesOrder int32 `json:"SubjectNamesOrder,omitempty"`
	// OCSP mode.
	ValidityOcspMode SlbCurSslCfgAuthPolTableValidityOcspMode `json:"ValidityOcspMode,omitempty"`
	// Static URI for OCSP Stapling validation requests.
	ValidityStaplinguri string `json:"ValidityStaplinguri,omitempty"`
	// OCSP Stapling URI priority.
	ValidityStaplingUriprior SlbCurSslCfgAuthPolTableValidityStaplingUriprior `json:"ValidityStaplingUriprior,omitempty"`
	// Static Secondary URI for OCSP validation requests.
	ValidityStatSecUri string `json:"ValidityStatSecUri,omitempty"`
	// Static Secondary URI for OCSP Stapling validation requests.
	ValidityStaplingSecUri string `json:"ValidityStaplingSecUri,omitempty"`
	// Number of retries reaching the OCSP server.
	ValidityOcspRetries uint64 `json:"ValidityOcspRetries,omitempty"`
}

func (p SlbCurSslCfgAuthPolTableParams) iMABean() {}
