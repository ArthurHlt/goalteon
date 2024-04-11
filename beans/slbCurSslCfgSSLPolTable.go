package beans

import (
	"fmt"
	"reflect"
)

// SlbCurSslCfgSSLPolTable The table for configuring ssl  policy.
// Note:This mib is not supported for VX instance of Virtualization.
type SlbCurSslCfgSSLPolTable struct {
	// The SSL policy name(key id) as an index.
	SlbCurSslCfgSSLPolNameIdIndex string
	Params                        *SlbCurSslCfgSSLPolTableParams
}

func NewSlbCurSslCfgSSLPolTableList() *SlbCurSslCfgSSLPolTable {
	return &SlbCurSslCfgSSLPolTable{}
}

func NewSlbCurSslCfgSSLPolTable(
	slbCurSslCfgSSLPolNameIdIndex string,
	params *SlbCurSslCfgSSLPolTableParams,
) *SlbCurSslCfgSSLPolTable {
	return &SlbCurSslCfgSSLPolTable{
		SlbCurSslCfgSSLPolNameIdIndex: slbCurSslCfgSSLPolNameIdIndex,
		Params:                        params,
	}
}

func (c *SlbCurSslCfgSSLPolTable) Name() string {
	return "SlbCurSslCfgSSLPolTable"
}

func (c *SlbCurSslCfgSSLPolTable) GetParams() BeanType {
	return c.Params
}

func (c *SlbCurSslCfgSSLPolTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *SlbCurSslCfgSSLPolTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.SlbCurSslCfgSSLPolNameIdIndex).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.SlbCurSslCfgSSLPolNameIdIndex)
}

type SlbCurSslCfgSSLPolTableBecipher int32

const (
	SlbCurSslCfgSSLPolTableBecipher_Low               SlbCurSslCfgSSLPolTableBecipher = 0
	SlbCurSslCfgSSLPolTableBecipher_Medium            SlbCurSslCfgSSLPolTableBecipher = 1
	SlbCurSslCfgSSLPolTableBecipher_High              SlbCurSslCfgSSLPolTableBecipher = 2
	SlbCurSslCfgSSLPolTableBecipher_UserDefined       SlbCurSslCfgSSLPolTableBecipher = 3
	SlbCurSslCfgSSLPolTableBecipher_UserDefinedExpert SlbCurSslCfgSSLPolTableBecipher = 4
	SlbCurSslCfgSSLPolTableBecipher_Unsupported       SlbCurSslCfgSSLPolTableBecipher = 2147483647
)

type SlbCurSslCfgSSLPolTableBessl int32

const (
	SlbCurSslCfgSSLPolTableBessl_Enabled     SlbCurSslCfgSSLPolTableBessl = 1
	SlbCurSslCfgSSLPolTableBessl_Disabled    SlbCurSslCfgSSLPolTableBessl = 2
	SlbCurSslCfgSSLPolTableBessl_Request     SlbCurSslCfgSSLPolTableBessl = 3
	SlbCurSslCfgSSLPolTableBessl_Handshake   SlbCurSslCfgSSLPolTableBessl = 4
	SlbCurSslCfgSSLPolTableBessl_Unsupported SlbCurSslCfgSSLPolTableBessl = 2147483647
)

type SlbCurSslCfgSSLPolTableConvert int32

const (
	SlbCurSslCfgSSLPolTableConvert_Enabled     SlbCurSslCfgSSLPolTableConvert = 1
	SlbCurSslCfgSSLPolTableConvert_Disabled    SlbCurSslCfgSSLPolTableConvert = 2
	SlbCurSslCfgSSLPolTableConvert_Unsupported SlbCurSslCfgSSLPolTableConvert = 2147483647
)

type SlbCurSslCfgSSLPolTableAdminStatus int32

const (
	SlbCurSslCfgSSLPolTableAdminStatus_Enabled     SlbCurSslCfgSSLPolTableAdminStatus = 1
	SlbCurSslCfgSSLPolTableAdminStatus_Disabled    SlbCurSslCfgSSLPolTableAdminStatus = 2
	SlbCurSslCfgSSLPolTableAdminStatus_Unsupported SlbCurSslCfgSSLPolTableAdminStatus = 2147483647
)

type SlbCurSslCfgSSLPolTableFessl int32

const (
	SlbCurSslCfgSSLPolTableFessl_Enabled     SlbCurSslCfgSSLPolTableFessl = 1
	SlbCurSslCfgSSLPolTableFessl_Disabled    SlbCurSslCfgSSLPolTableFessl = 2
	SlbCurSslCfgSSLPolTableFessl_Request     SlbCurSslCfgSSLPolTableFessl = 3
	SlbCurSslCfgSSLPolTableFessl_Handshake   SlbCurSslCfgSSLPolTableFessl = 4
	SlbCurSslCfgSSLPolTableFessl_Unsupported SlbCurSslCfgSSLPolTableFessl = 2147483647
)

type SlbCurSslCfgSSLPolTableFESslv3Version int32

const (
	SlbCurSslCfgSSLPolTableFESslv3Version_Enabled     SlbCurSslCfgSSLPolTableFESslv3Version = 1
	SlbCurSslCfgSSLPolTableFESslv3Version_Disabled    SlbCurSslCfgSSLPolTableFESslv3Version = 2
	SlbCurSslCfgSSLPolTableFESslv3Version_Unsupported SlbCurSslCfgSSLPolTableFESslv3Version = 2147483647
)

type SlbCurSslCfgSSLPolTableFETls10Version int32

const (
	SlbCurSslCfgSSLPolTableFETls10Version_Enabled     SlbCurSslCfgSSLPolTableFETls10Version = 1
	SlbCurSslCfgSSLPolTableFETls10Version_Disabled    SlbCurSslCfgSSLPolTableFETls10Version = 2
	SlbCurSslCfgSSLPolTableFETls10Version_Unsupported SlbCurSslCfgSSLPolTableFETls10Version = 2147483647
)

type SlbCurSslCfgSSLPolTableFETls11Version int32

const (
	SlbCurSslCfgSSLPolTableFETls11Version_Enabled     SlbCurSslCfgSSLPolTableFETls11Version = 1
	SlbCurSslCfgSSLPolTableFETls11Version_Disabled    SlbCurSslCfgSSLPolTableFETls11Version = 2
	SlbCurSslCfgSSLPolTableFETls11Version_Unsupported SlbCurSslCfgSSLPolTableFETls11Version = 2147483647
)

type SlbCurSslCfgSSLPolTableBESslv3Version int32

const (
	SlbCurSslCfgSSLPolTableBESslv3Version_Enabled     SlbCurSslCfgSSLPolTableBESslv3Version = 1
	SlbCurSslCfgSSLPolTableBESslv3Version_Disabled    SlbCurSslCfgSSLPolTableBESslv3Version = 2
	SlbCurSslCfgSSLPolTableBESslv3Version_Unsupported SlbCurSslCfgSSLPolTableBESslv3Version = 2147483647
)

type SlbCurSslCfgSSLPolTableBETls10Version int32

const (
	SlbCurSslCfgSSLPolTableBETls10Version_Enabled     SlbCurSslCfgSSLPolTableBETls10Version = 1
	SlbCurSslCfgSSLPolTableBETls10Version_Disabled    SlbCurSslCfgSSLPolTableBETls10Version = 2
	SlbCurSslCfgSSLPolTableBETls10Version_Unsupported SlbCurSslCfgSSLPolTableBETls10Version = 2147483647
)

type SlbCurSslCfgSSLPolTableBETls11Version int32

const (
	SlbCurSslCfgSSLPolTableBETls11Version_Enabled     SlbCurSslCfgSSLPolTableBETls11Version = 1
	SlbCurSslCfgSSLPolTableBETls11Version_Disabled    SlbCurSslCfgSSLPolTableBETls11Version = 2
	SlbCurSslCfgSSLPolTableBETls11Version_Unsupported SlbCurSslCfgSSLPolTableBETls11Version = 2147483647
)

type SlbCurSslCfgSSLPolTableFETls12Version int32

const (
	SlbCurSslCfgSSLPolTableFETls12Version_Enabled     SlbCurSslCfgSSLPolTableFETls12Version = 1
	SlbCurSslCfgSSLPolTableFETls12Version_Disabled    SlbCurSslCfgSSLPolTableFETls12Version = 2
	SlbCurSslCfgSSLPolTableFETls12Version_Unsupported SlbCurSslCfgSSLPolTableFETls12Version = 2147483647
)

type SlbCurSslCfgSSLPolTableBETls12Version int32

const (
	SlbCurSslCfgSSLPolTableBETls12Version_Enabled     SlbCurSslCfgSSLPolTableBETls12Version = 1
	SlbCurSslCfgSSLPolTableBETls12Version_Disabled    SlbCurSslCfgSSLPolTableBETls12Version = 2
	SlbCurSslCfgSSLPolTableBETls12Version_Unsupported SlbCurSslCfgSSLPolTableBETls12Version = 2147483647
)

type SlbCurSslCfgSSLPolTableDHkey int32

const (
	SlbCurSslCfgSSLPolTableDHkey_KeySize1024 SlbCurSslCfgSSLPolTableDHkey = 1
	SlbCurSslCfgSSLPolTableDHkey_KeySize2048 SlbCurSslCfgSSLPolTableDHkey = 2
	SlbCurSslCfgSSLPolTableDHkey_Unsupported SlbCurSslCfgSSLPolTableDHkey = 2147483647
)

type SlbCurSslCfgSSLPolTableHwoffldFeRsa int32

const (
	SlbCurSslCfgSSLPolTableHwoffldFeRsa_Enabled     SlbCurSslCfgSSLPolTableHwoffldFeRsa = 1
	SlbCurSslCfgSSLPolTableHwoffldFeRsa_Disabled    SlbCurSslCfgSSLPolTableHwoffldFeRsa = 2
	SlbCurSslCfgSSLPolTableHwoffldFeRsa_Unsupported SlbCurSslCfgSSLPolTableHwoffldFeRsa = 2147483647
)

type SlbCurSslCfgSSLPolTableHwoffldFeDh int32

const (
	SlbCurSslCfgSSLPolTableHwoffldFeDh_Enabled     SlbCurSslCfgSSLPolTableHwoffldFeDh = 1
	SlbCurSslCfgSSLPolTableHwoffldFeDh_Disabled    SlbCurSslCfgSSLPolTableHwoffldFeDh = 2
	SlbCurSslCfgSSLPolTableHwoffldFeDh_Unsupported SlbCurSslCfgSSLPolTableHwoffldFeDh = 2147483647
)

type SlbCurSslCfgSSLPolTableHwoffldFeEc int32

const (
	SlbCurSslCfgSSLPolTableHwoffldFeEc_Enabled     SlbCurSslCfgSSLPolTableHwoffldFeEc = 1
	SlbCurSslCfgSSLPolTableHwoffldFeEc_Disabled    SlbCurSslCfgSSLPolTableHwoffldFeEc = 2
	SlbCurSslCfgSSLPolTableHwoffldFeEc_Unsupported SlbCurSslCfgSSLPolTableHwoffldFeEc = 2147483647
)

type SlbCurSslCfgSSLPolTableHwoffldFeBulk int32

const (
	SlbCurSslCfgSSLPolTableHwoffldFeBulk_Enabled     SlbCurSslCfgSSLPolTableHwoffldFeBulk = 1
	SlbCurSslCfgSSLPolTableHwoffldFeBulk_Disabled    SlbCurSslCfgSSLPolTableHwoffldFeBulk = 2
	SlbCurSslCfgSSLPolTableHwoffldFeBulk_Unsupported SlbCurSslCfgSSLPolTableHwoffldFeBulk = 2147483647
)

type SlbCurSslCfgSSLPolTableHwoffldFePkey int32

const (
	SlbCurSslCfgSSLPolTableHwoffldFePkey_Enabled     SlbCurSslCfgSSLPolTableHwoffldFePkey = 1
	SlbCurSslCfgSSLPolTableHwoffldFePkey_Disabled    SlbCurSslCfgSSLPolTableHwoffldFePkey = 2
	SlbCurSslCfgSSLPolTableHwoffldFePkey_Unsupported SlbCurSslCfgSSLPolTableHwoffldFePkey = 2147483647
)

type SlbCurSslCfgSSLPolTableHwoffldFeFeatures int32

const (
	SlbCurSslCfgSSLPolTableHwoffldFeFeatures_Enabled     SlbCurSslCfgSSLPolTableHwoffldFeFeatures = 1
	SlbCurSslCfgSSLPolTableHwoffldFeFeatures_Disabled    SlbCurSslCfgSSLPolTableHwoffldFeFeatures = 2
	SlbCurSslCfgSSLPolTableHwoffldFeFeatures_Unsupported SlbCurSslCfgSSLPolTableHwoffldFeFeatures = 2147483647
)

type SlbCurSslCfgSSLPolTableHwoffldBeRsa int32

const (
	SlbCurSslCfgSSLPolTableHwoffldBeRsa_Enabled     SlbCurSslCfgSSLPolTableHwoffldBeRsa = 1
	SlbCurSslCfgSSLPolTableHwoffldBeRsa_Disabled    SlbCurSslCfgSSLPolTableHwoffldBeRsa = 2
	SlbCurSslCfgSSLPolTableHwoffldBeRsa_Unsupported SlbCurSslCfgSSLPolTableHwoffldBeRsa = 2147483647
)

type SlbCurSslCfgSSLPolTableHwoffldBeDh int32

const (
	SlbCurSslCfgSSLPolTableHwoffldBeDh_Enabled     SlbCurSslCfgSSLPolTableHwoffldBeDh = 1
	SlbCurSslCfgSSLPolTableHwoffldBeDh_Disabled    SlbCurSslCfgSSLPolTableHwoffldBeDh = 2
	SlbCurSslCfgSSLPolTableHwoffldBeDh_Unsupported SlbCurSslCfgSSLPolTableHwoffldBeDh = 2147483647
)

type SlbCurSslCfgSSLPolTableHwoffldBeEc int32

const (
	SlbCurSslCfgSSLPolTableHwoffldBeEc_Enabled     SlbCurSslCfgSSLPolTableHwoffldBeEc = 1
	SlbCurSslCfgSSLPolTableHwoffldBeEc_Disabled    SlbCurSslCfgSSLPolTableHwoffldBeEc = 2
	SlbCurSslCfgSSLPolTableHwoffldBeEc_Unsupported SlbCurSslCfgSSLPolTableHwoffldBeEc = 2147483647
)

type SlbCurSslCfgSSLPolTableHwoffldBeBulk int32

const (
	SlbCurSslCfgSSLPolTableHwoffldBeBulk_Enabled     SlbCurSslCfgSSLPolTableHwoffldBeBulk = 1
	SlbCurSslCfgSSLPolTableHwoffldBeBulk_Disabled    SlbCurSslCfgSSLPolTableHwoffldBeBulk = 2
	SlbCurSslCfgSSLPolTableHwoffldBeBulk_Unsupported SlbCurSslCfgSSLPolTableHwoffldBeBulk = 2147483647
)

type SlbCurSslCfgSSLPolTableHwoffldBePkey int32

const (
	SlbCurSslCfgSSLPolTableHwoffldBePkey_Enabled     SlbCurSslCfgSSLPolTableHwoffldBePkey = 1
	SlbCurSslCfgSSLPolTableHwoffldBePkey_Disabled    SlbCurSslCfgSSLPolTableHwoffldBePkey = 2
	SlbCurSslCfgSSLPolTableHwoffldBePkey_Unsupported SlbCurSslCfgSSLPolTableHwoffldBePkey = 2147483647
)

type SlbCurSslCfgSSLPolTableHwoffldBeFeatures int32

const (
	SlbCurSslCfgSSLPolTableHwoffldBeFeatures_Enabled     SlbCurSslCfgSSLPolTableHwoffldBeFeatures = 1
	SlbCurSslCfgSSLPolTableHwoffldBeFeatures_Disabled    SlbCurSslCfgSSLPolTableHwoffldBeFeatures = 2
	SlbCurSslCfgSSLPolTableHwoffldBeFeatures_Unsupported SlbCurSslCfgSSLPolTableHwoffldBeFeatures = 2147483647
)

type SlbCurSslCfgSSLPolTableBesni int32

const (
	SlbCurSslCfgSSLPolTableBesni_Enabled     SlbCurSslCfgSSLPolTableBesni = 1
	SlbCurSslCfgSSLPolTableBesni_Disabled    SlbCurSslCfgSSLPolTableBesni = 2
	SlbCurSslCfgSSLPolTableBesni_Unsupported SlbCurSslCfgSSLPolTableBesni = 2147483647
)

type SlbCurSslCfgSSLPolTableFETls13Version int32

const (
	SlbCurSslCfgSSLPolTableFETls13Version_Enabled     SlbCurSslCfgSSLPolTableFETls13Version = 1
	SlbCurSslCfgSSLPolTableFETls13Version_Disabled    SlbCurSslCfgSSLPolTableFETls13Version = 2
	SlbCurSslCfgSSLPolTableFETls13Version_Unsupported SlbCurSslCfgSSLPolTableFETls13Version = 2147483647
)

type SlbCurSslCfgSSLPolTableBETls13Version int32

const (
	SlbCurSslCfgSSLPolTableBETls13Version_Enabled     SlbCurSslCfgSSLPolTableBETls13Version = 1
	SlbCurSslCfgSSLPolTableBETls13Version_Disabled    SlbCurSslCfgSSLPolTableBETls13Version = 2
	SlbCurSslCfgSSLPolTableBETls13Version_Unsupported SlbCurSslCfgSSLPolTableBETls13Version = 2147483647
)

type SlbCurSslCfgSSLPolTableFeReuseState int32

const (
	SlbCurSslCfgSSLPolTableFeReuseState_Enabled     SlbCurSslCfgSSLPolTableFeReuseState = 1
	SlbCurSslCfgSSLPolTableFeReuseState_Disabled    SlbCurSslCfgSSLPolTableFeReuseState = 2
	SlbCurSslCfgSSLPolTableFeReuseState_Inherit     SlbCurSslCfgSSLPolTableFeReuseState = 3
	SlbCurSslCfgSSLPolTableFeReuseState_Unsupported SlbCurSslCfgSSLPolTableFeReuseState = 2147483647
)

type SlbCurSslCfgSSLPolTableBeReuseState int32

const (
	SlbCurSslCfgSSLPolTableBeReuseState_Enabled     SlbCurSslCfgSSLPolTableBeReuseState = 1
	SlbCurSslCfgSSLPolTableBeReuseState_Disabled    SlbCurSslCfgSSLPolTableBeReuseState = 2
	SlbCurSslCfgSSLPolTableBeReuseState_Inherit     SlbCurSslCfgSSLPolTableBeReuseState = 3
	SlbCurSslCfgSSLPolTableBeReuseState_Unsupported SlbCurSslCfgSSLPolTableBeReuseState = 2147483647
)

type SlbCurSslCfgSSLPolTableBeReuseSrcMatch int32

const (
	SlbCurSslCfgSSLPolTableBeReuseSrcMatch_Enabled     SlbCurSslCfgSSLPolTableBeReuseSrcMatch = 1
	SlbCurSslCfgSSLPolTableBeReuseSrcMatch_Disabled    SlbCurSslCfgSSLPolTableBeReuseSrcMatch = 2
	SlbCurSslCfgSSLPolTableBeReuseSrcMatch_Unsupported SlbCurSslCfgSSLPolTableBeReuseSrcMatch = 2147483647
)

type SlbCurSslCfgSSLPolTableBeReuseTicket int32

const (
	SlbCurSslCfgSSLPolTableBeReuseTicket_Enabled     SlbCurSslCfgSSLPolTableBeReuseTicket = 1
	SlbCurSslCfgSSLPolTableBeReuseTicket_Disabled    SlbCurSslCfgSSLPolTableBeReuseTicket = 2
	SlbCurSslCfgSSLPolTableBeReuseTicket_Unsupported SlbCurSslCfgSSLPolTableBeReuseTicket = 2147483647
)

type SlbCurSslCfgSSLPolTableFeReuseTicket int32

const (
	SlbCurSslCfgSSLPolTableFeReuseTicket_Enabled     SlbCurSslCfgSSLPolTableFeReuseTicket = 1
	SlbCurSslCfgSSLPolTableFeReuseTicket_Disabled    SlbCurSslCfgSSLPolTableFeReuseTicket = 2
	SlbCurSslCfgSSLPolTableFeReuseTicket_Unsupported SlbCurSslCfgSSLPolTableFeReuseTicket = 2147483647
)

type SlbCurSslCfgSSLPolTableParams struct {
	// The SSL policy name(key id) as an index.
	NameIdIndex string `json:"NameIdIndex,omitempty"`
	// SSL policy name.
	Name string `json:"Name,omitempty"`
	// The SSL cipher-suite header name.
	PassInfoCipherName string `json:"PassInfoCipherName,omitempty"`
	// SSL cipher-suite information to backend servers enabled.
	PassInfoCipherFlag int32 `json:"PassInfoCipherFlag,omitempty"`
	// SSL version header name.
	PassInfoVersionName string `json:"PassInfoVersionName,omitempty"`
	// SSL version information to backend servers enabled.
	PassInfoVersionFlag int32 `json:"PassInfoVersionFlag,omitempty"`
	// The passive cipher bits information to backend server.
	PassInfoHeadBitsName string `json:"PassInfoHeadBitsName,omitempty"`
	// Cipher bits information to backend servers enabled.
	PassInfoHeadBitsFlag int32 `json:"PassInfoHeadBitsFlag,omitempty"`
	// Enable/Disable add Front-End-Https: on header.
	PassInfoFrontend int32 `json:"PassInfoFrontend,omitempty"`
	// Cipher name for SSL.
	CipherName int32 `json:"CipherName,omitempty"`
	// Cipher-suite allowed for SSL.
	CipherUserdef string `json:"CipherUserdef,omitempty"`
	// Intermediate CA certificate chain name.
	IntermcaChainName string `json:"IntermcaChainName,omitempty"`
	// Intermediate CA certificate chain type certificate=cert,Group=group,None=empty string.
	IntermcaChainType string `json:"IntermcaChainType,omitempty"`
	// Allowed cipher-suites in backend SSL [0-low, 1-midium, 2-high, 3-user-defined, 4-user-defined-expert] .
	Becipher SlbCurSslCfgSSLPolTableBecipher `json:"Becipher,omitempty"`
	// Client authentication policy.
	Authpol string `json:"Authpol,omitempty"`
	// Host regex for HTTP redirection conversion.
	Convuri string `json:"Convuri,omitempty"`
	// Enable/Disable backend SSL encryption.
	Bessl SlbCurSslCfgSSLPolTableBessl `json:"Bessl,omitempty"`
	// Enable/Disable HTTP redirection conversion.
	Convert SlbCurSslCfgSSLPolTableConvert `json:"Convert,omitempty"`
	// Enable or disable ssl policy.
	AdminStatus SlbCurSslCfgSSLPolTableAdminStatus `json:"AdminStatus,omitempty"`
	// Enable/Disable X-SSL header compatible with 2424SSL headers.
	PassInfoComply int32 `json:"PassInfoComply,omitempty"`
	// Set frontend SSL encryption mode.
	Fessl SlbCurSslCfgSSLPolTableFessl `json:"Fessl,omitempty"`
	// Enable or disable frontend sslv3.
	FESslv3Version SlbCurSslCfgSSLPolTableFESslv3Version `json:"FESslv3Version,omitempty"`
	// Enable or disable frontend tls1_0.
	FETls10Version SlbCurSslCfgSSLPolTableFETls10Version `json:"FETls10Version,omitempty"`
	// Enable or disable frontend tls1_1.
	FETls11Version SlbCurSslCfgSSLPolTableFETls11Version `json:"FETls11Version,omitempty"`
	// Enable or disable backend sslv3.
	BESslv3Version SlbCurSslCfgSSLPolTableBESslv3Version `json:"BESslv3Version,omitempty"`
	// Enable or disable backend tls10.
	BETls10Version SlbCurSslCfgSSLPolTableBETls10Version `json:"BETls10Version,omitempty"`
	// Enable or disable backend tls11.
	BETls11Version SlbCurSslCfgSSLPolTableBETls11Version `json:"BETls11Version,omitempty"`
	// Enable or disable frontend tls1_2.
	FETls12Version SlbCurSslCfgSSLPolTableFETls12Version `json:"FETls12Version,omitempty"`
	// Enable or disable backend tls1_2.
	BETls12Version SlbCurSslCfgSSLPolTableBETls12Version `json:"BETls12Version,omitempty"`
	// Expert-Cipher-suite allowed for SSL.
	CipherExpertUserdef string `json:"CipherExpertUserdef,omitempty"`
	// BeCipher-suite allowed for SSL.
	BeCipherUserdef string `json:"BeCipherUserdef,omitempty"`
	// Expert-BeCipher-suite allowed for SSL.
	BeCipherExpertUserdef string `json:"BeCipherExpertUserdef,omitempty"`
	// Backend client certificate name.
	BEClientCertName string `json:"BEClientCertName,omitempty"`
	// Backend trusted CA certificate chain name.
	BETrustedCAcertName string `json:"BETrustedCAcertName,omitempty"`
	// Backend trusted CA certificate chain type certificate=cert,Group=group,None=empty string.
	BETrustedCAcertType string `json:"BETrustedCAcertType,omitempty"`
	// Secure Renegotiation Frontend and Backend SSL.
	Secreneg string `json:"Secreneg,omitempty"`
	// num of bits in Diffie Helman key.
	DHkey SlbCurSslCfgSSLPolTableDHkey `json:"DHkey,omitempty"`
	// backend authentication policy.
	BEAuthpol string `json:"BEAuthpol,omitempty"`
	// Disable/Enable HW offload for RSA algorithm.
	HwoffldFeRsa SlbCurSslCfgSSLPolTableHwoffldFeRsa `json:"HwoffldFeRsa,omitempty"`
	// Disable/Enable HW offload for DH algorithm.
	HwoffldFeDh SlbCurSslCfgSSLPolTableHwoffldFeDh `json:"HwoffldFeDh,omitempty"`
	// Disable/Enable HW offload for EC algorithm.
	HwoffldFeEc SlbCurSslCfgSSLPolTableHwoffldFeEc `json:"HwoffldFeEc,omitempty"`
	// Disable/Enable HW offload for bulk encryption ciphers.
	HwoffldFeBulk SlbCurSslCfgSSLPolTableHwoffldFeBulk `json:"HwoffldFeBulk,omitempty"`
	// Disable/Enable HW offload for PKEY functionality.
	HwoffldFePkey SlbCurSslCfgSSLPolTableHwoffldFePkey `json:"HwoffldFePkey,omitempty"`
	// Disable/Enable HW offload Features.
	HwoffldFeFeatures SlbCurSslCfgSSLPolTableHwoffldFeFeatures `json:"HwoffldFeFeatures,omitempty"`
	// Disable/Enable HW offload for RSA algorithm.
	HwoffldBeRsa SlbCurSslCfgSSLPolTableHwoffldBeRsa `json:"HwoffldBeRsa,omitempty"`
	// Disable/Enable HW offload for DH algorithm.
	HwoffldBeDh SlbCurSslCfgSSLPolTableHwoffldBeDh `json:"HwoffldBeDh,omitempty"`
	// Disable/Enable HW offload for EC algorithm.
	HwoffldBeEc SlbCurSslCfgSSLPolTableHwoffldBeEc `json:"HwoffldBeEc,omitempty"`
	// Disable/Enable HW offload for bulk encryption ciphers.
	HwoffldBeBulk SlbCurSslCfgSSLPolTableHwoffldBeBulk `json:"HwoffldBeBulk,omitempty"`
	// Disable/Enable HW offload for PKEY functionality.
	HwoffldBePkey SlbCurSslCfgSSLPolTableHwoffldBePkey `json:"HwoffldBePkey,omitempty"`
	// Disable/Enable HW offload Features.
	HwoffldBeFeatures SlbCurSslCfgSSLPolTableHwoffldBeFeatures `json:"HwoffldBeFeatures,omitempty"`
	// Enable/Disable to include SNI.
	Besni SlbCurSslCfgSSLPolTableBesni `json:"Besni,omitempty"`
	// Enable or disable frontend tls1_3.
	FETls13Version SlbCurSslCfgSSLPolTableFETls13Version `json:"FETls13Version,omitempty"`
	// Enable or disable backend tls1_3.
	BETls13Version SlbCurSslCfgSSLPolTableBETls13Version `json:"BETls13Version,omitempty"`
	// Set maximum allowed early data on frontend connection
	SSLPol0RTTFEData uint32 `json:"SSLPol0RTTFEData,omitempty"`
	// Set Frontend SSL reuse state
	FeReuseState SlbCurSslCfgSSLPolTableFeReuseState `json:"FeReuseState,omitempty"`
	// Set Backend SSL reuse state
	BeReuseState SlbCurSslCfgSSLPolTableBeReuseState `json:"BeReuseState,omitempty"`
	// Enable/disable reuse for same client only
	BeReuseSrcMatch SlbCurSslCfgSSLPolTableBeReuseSrcMatch `json:"BeReuseSrcMatch,omitempty"`
	// Enable/disable TLS 1.2 session ticket
	BeReuseTicket SlbCurSslCfgSSLPolTableBeReuseTicket `json:"BeReuseTicket,omitempty"`
	// Enable/disable TLS 1.2 session ticket
	FeReuseTicket SlbCurSslCfgSSLPolTableFeReuseTicket `json:"FeReuseTicket,omitempty"`
	// Allowed signature algorithms.
	FESslsigs string `json:"FESslsigs,omitempty"`
	// Allowed signature algorithms in backend SSL.
	BESslsigs string `json:"BESslsigs,omitempty"`
	// Allowed groups.
	FESslgroups string `json:"FESslgroups,omitempty"`
	// Set allowed groups in backend SSL.
	BESslgroups string `json:"BESslgroups,omitempty"`
}

func (p SlbCurSslCfgSSLPolTableParams) iMABean() {}
