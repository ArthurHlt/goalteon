package beans

import (
	"fmt"
	"reflect"
)

type SlbCurSslCfgCertsTableSlbCurSslCfgCertsType int32

const (
	SlbCurSslCfgCertsTableSlbCurSslCfgCertsType_Key                     SlbCurSslCfgCertsTableSlbCurSslCfgCertsType = 1
	SlbCurSslCfgCertsTableSlbCurSslCfgCertsType_CertificateRequest      SlbCurSslCfgCertsTableSlbCurSslCfgCertsType = 2
	SlbCurSslCfgCertsTableSlbCurSslCfgCertsType_ServerCertificate       SlbCurSslCfgCertsTableSlbCurSslCfgCertsType = 3
	SlbCurSslCfgCertsTableSlbCurSslCfgCertsType_TrustedCertificate      SlbCurSslCfgCertsTableSlbCurSslCfgCertsType = 4
	SlbCurSslCfgCertsTableSlbCurSslCfgCertsType_IntermediateCertificate SlbCurSslCfgCertsTableSlbCurSslCfgCertsType = 5
	SlbCurSslCfgCertsTableSlbCurSslCfgCertsType_Crl                     SlbCurSslCfgCertsTableSlbCurSslCfgCertsType = 6
	SlbCurSslCfgCertsTableSlbCurSslCfgCertsType_Unsupported             SlbCurSslCfgCertsTableSlbCurSslCfgCertsType = 2147483647
)

// SlbCurSslCfgCertsTable Note:This mib is not supported for VX instance of Virtualization.
type SlbCurSslCfgCertsTable struct {
	// Certificate id.
	SlbCurSslCfgCertsID string
	// Certificate type.
	SlbCurSslCfgCertsType SlbCurSslCfgCertsTableSlbCurSslCfgCertsType
	Params                *SlbCurSslCfgCertsTableParams
}

func NewSlbCurSslCfgCertsTableList() *SlbCurSslCfgCertsTable {
	return &SlbCurSslCfgCertsTable{}
}

func NewSlbCurSslCfgCertsTable(
	slbCurSslCfgCertsID string,
	slbCurSslCfgCertsType SlbCurSslCfgCertsTableSlbCurSslCfgCertsType,
	params *SlbCurSslCfgCertsTableParams,
) *SlbCurSslCfgCertsTable {
	return &SlbCurSslCfgCertsTable{
		SlbCurSslCfgCertsID:   slbCurSslCfgCertsID,
		SlbCurSslCfgCertsType: slbCurSslCfgCertsType,
		Params:                params,
	}
}

func (c *SlbCurSslCfgCertsTable) Name() string {
	return "SlbCurSslCfgCertsTable"
}

func (c *SlbCurSslCfgCertsTable) GetParams() BeanType {
	return c.Params
}

func (c *SlbCurSslCfgCertsTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *SlbCurSslCfgCertsTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.SlbCurSslCfgCertsID).IsZero() &&
		reflect.ValueOf(c.SlbCurSslCfgCertsType).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.SlbCurSslCfgCertsID) + "/" + fmt.Sprint(c.SlbCurSslCfgCertsType)
}

type SlbCurSslCfgCertsTableType int32

const (
	SlbCurSslCfgCertsTableType_Key                     SlbCurSslCfgCertsTableType = 1
	SlbCurSslCfgCertsTableType_CertificateRequest      SlbCurSslCfgCertsTableType = 2
	SlbCurSslCfgCertsTableType_ServerCertificate       SlbCurSslCfgCertsTableType = 3
	SlbCurSslCfgCertsTableType_TrustedCertificate      SlbCurSslCfgCertsTableType = 4
	SlbCurSslCfgCertsTableType_IntermediateCertificate SlbCurSslCfgCertsTableType = 5
	SlbCurSslCfgCertsTableType_Crl                     SlbCurSslCfgCertsTableType = 6
	SlbCurSslCfgCertsTableType_Unsupported             SlbCurSslCfgCertsTableType = 2147483647
)

type SlbCurSslCfgCertsTableKeySize int32

const (
	SlbCurSslCfgCertsTableKeySize_Ks512       SlbCurSslCfgCertsTableKeySize = 1
	SlbCurSslCfgCertsTableKeySize_Ks1024      SlbCurSslCfgCertsTableKeySize = 2
	SlbCurSslCfgCertsTableKeySize_Ks2048      SlbCurSslCfgCertsTableKeySize = 3
	SlbCurSslCfgCertsTableKeySize_Ks4096      SlbCurSslCfgCertsTableKeySize = 4
	SlbCurSslCfgCertsTableKeySize_Unknown     SlbCurSslCfgCertsTableKeySize = 6
	SlbCurSslCfgCertsTableKeySize_Unsupported SlbCurSslCfgCertsTableKeySize = 2147483647
)

type SlbCurSslCfgCertsTableHashAlgo int32

const (
	SlbCurSslCfgCertsTableHashAlgo_Md5         SlbCurSslCfgCertsTableHashAlgo = 1
	SlbCurSslCfgCertsTableHashAlgo_Sha1        SlbCurSslCfgCertsTableHashAlgo = 2
	SlbCurSslCfgCertsTableHashAlgo_Sha256      SlbCurSslCfgCertsTableHashAlgo = 3
	SlbCurSslCfgCertsTableHashAlgo_Sha384      SlbCurSslCfgCertsTableHashAlgo = 4
	SlbCurSslCfgCertsTableHashAlgo_Sha512      SlbCurSslCfgCertsTableHashAlgo = 5
	SlbCurSslCfgCertsTableHashAlgo_Unknown     SlbCurSslCfgCertsTableHashAlgo = 6
	SlbCurSslCfgCertsTableHashAlgo_Unsupported SlbCurSslCfgCertsTableHashAlgo = 2147483647
)

type SlbCurSslCfgCertsTableStatus int32

const (
	SlbCurSslCfgCertsTableStatus_Generated    SlbCurSslCfgCertsTableStatus = 1
	SlbCurSslCfgCertsTableStatus_NotGenerated SlbCurSslCfgCertsTableStatus = 2
	SlbCurSslCfgCertsTableStatus_Unsupported  SlbCurSslCfgCertsTableStatus = 2147483647
)

type SlbCurSslCfgCertsTableGenerate int32

const (
	SlbCurSslCfgCertsTableGenerate_Other        SlbCurSslCfgCertsTableGenerate = 1
	SlbCurSslCfgCertsTableGenerate_Generate     SlbCurSslCfgCertsTableGenerate = 2
	SlbCurSslCfgCertsTableGenerate_Idle         SlbCurSslCfgCertsTableGenerate = 3
	SlbCurSslCfgCertsTableGenerate_Inprogress   SlbCurSslCfgCertsTableGenerate = 4
	SlbCurSslCfgCertsTableGenerate_Generated    SlbCurSslCfgCertsTableGenerate = 5
	SlbCurSslCfgCertsTableGenerate_NotGenerated SlbCurSslCfgCertsTableGenerate = 6
	SlbCurSslCfgCertsTableGenerate_Unsupported  SlbCurSslCfgCertsTableGenerate = 2147483647
)

type SlbCurSslCfgCertsTableKeyType int32

const (
	SlbCurSslCfgCertsTableKeyType_Rsa         SlbCurSslCfgCertsTableKeyType = 1
	SlbCurSslCfgCertsTableKeyType_Ec          SlbCurSslCfgCertsTableKeyType = 2
	SlbCurSslCfgCertsTableKeyType_Unknown     SlbCurSslCfgCertsTableKeyType = 6
	SlbCurSslCfgCertsTableKeyType_Unsupported SlbCurSslCfgCertsTableKeyType = 2147483647
)

type SlbCurSslCfgCertsTableKeySizeEc int32

const (
	SlbCurSslCfgCertsTableKeySizeEc_Ks0         SlbCurSslCfgCertsTableKeySizeEc = 0
	SlbCurSslCfgCertsTableKeySizeEc_Ks192       SlbCurSslCfgCertsTableKeySizeEc = 1
	SlbCurSslCfgCertsTableKeySizeEc_Ks224       SlbCurSslCfgCertsTableKeySizeEc = 2
	SlbCurSslCfgCertsTableKeySizeEc_Ks256       SlbCurSslCfgCertsTableKeySizeEc = 3
	SlbCurSslCfgCertsTableKeySizeEc_Ks384       SlbCurSslCfgCertsTableKeySizeEc = 4
	SlbCurSslCfgCertsTableKeySizeEc_Ks521       SlbCurSslCfgCertsTableKeySizeEc = 5
	SlbCurSslCfgCertsTableKeySizeEc_Unknown     SlbCurSslCfgCertsTableKeySizeEc = 6
	SlbCurSslCfgCertsTableKeySizeEc_Unsupported SlbCurSslCfgCertsTableKeySizeEc = 2147483647
)

type SlbCurSslCfgCertsTableCurveNameEc int32

const (
	SlbCurSslCfgCertsTableCurveNameEc_Unknown     SlbCurSslCfgCertsTableCurveNameEc = 0
	SlbCurSslCfgCertsTableCurveNameEc_Secp112r1   SlbCurSslCfgCertsTableCurveNameEc = 1
	SlbCurSslCfgCertsTableCurveNameEc_Secp112r2   SlbCurSslCfgCertsTableCurveNameEc = 2
	SlbCurSslCfgCertsTableCurveNameEc_Secp128r1   SlbCurSslCfgCertsTableCurveNameEc = 3
	SlbCurSslCfgCertsTableCurveNameEc_Secp128r2   SlbCurSslCfgCertsTableCurveNameEc = 4
	SlbCurSslCfgCertsTableCurveNameEc_Secp160k1   SlbCurSslCfgCertsTableCurveNameEc = 5
	SlbCurSslCfgCertsTableCurveNameEc_Secp160r1   SlbCurSslCfgCertsTableCurveNameEc = 6
	SlbCurSslCfgCertsTableCurveNameEc_Secp160r2   SlbCurSslCfgCertsTableCurveNameEc = 7
	SlbCurSslCfgCertsTableCurveNameEc_Secp192k1   SlbCurSslCfgCertsTableCurveNameEc = 8
	SlbCurSslCfgCertsTableCurveNameEc_Secp224k1   SlbCurSslCfgCertsTableCurveNameEc = 9
	SlbCurSslCfgCertsTableCurveNameEc_Secp224r1   SlbCurSslCfgCertsTableCurveNameEc = 10
	SlbCurSslCfgCertsTableCurveNameEc_Secp256k1   SlbCurSslCfgCertsTableCurveNameEc = 11
	SlbCurSslCfgCertsTableCurveNameEc_Secp384r1   SlbCurSslCfgCertsTableCurveNameEc = 12
	SlbCurSslCfgCertsTableCurveNameEc_Secp521r1   SlbCurSslCfgCertsTableCurveNameEc = 13
	SlbCurSslCfgCertsTableCurveNameEc_Prime192v1  SlbCurSslCfgCertsTableCurveNameEc = 14
	SlbCurSslCfgCertsTableCurveNameEc_Prime192v2  SlbCurSslCfgCertsTableCurveNameEc = 15
	SlbCurSslCfgCertsTableCurveNameEc_Prime192v3  SlbCurSslCfgCertsTableCurveNameEc = 16
	SlbCurSslCfgCertsTableCurveNameEc_Prime239v1  SlbCurSslCfgCertsTableCurveNameEc = 17
	SlbCurSslCfgCertsTableCurveNameEc_Prime239v2  SlbCurSslCfgCertsTableCurveNameEc = 18
	SlbCurSslCfgCertsTableCurveNameEc_Prime239v3  SlbCurSslCfgCertsTableCurveNameEc = 19
	SlbCurSslCfgCertsTableCurveNameEc_Prime256v1  SlbCurSslCfgCertsTableCurveNameEc = 20
	SlbCurSslCfgCertsTableCurveNameEc_Sect113r1   SlbCurSslCfgCertsTableCurveNameEc = 21
	SlbCurSslCfgCertsTableCurveNameEc_Sect113r2   SlbCurSslCfgCertsTableCurveNameEc = 22
	SlbCurSslCfgCertsTableCurveNameEc_Sect131r1   SlbCurSslCfgCertsTableCurveNameEc = 23
	SlbCurSslCfgCertsTableCurveNameEc_Sect131r2   SlbCurSslCfgCertsTableCurveNameEc = 24
	SlbCurSslCfgCertsTableCurveNameEc_Sect163k1   SlbCurSslCfgCertsTableCurveNameEc = 25
	SlbCurSslCfgCertsTableCurveNameEc_Sect163r1   SlbCurSslCfgCertsTableCurveNameEc = 26
	SlbCurSslCfgCertsTableCurveNameEc_Sect163r2   SlbCurSslCfgCertsTableCurveNameEc = 27
	SlbCurSslCfgCertsTableCurveNameEc_Sect193r1   SlbCurSslCfgCertsTableCurveNameEc = 28
	SlbCurSslCfgCertsTableCurveNameEc_Sect193r2   SlbCurSslCfgCertsTableCurveNameEc = 29
	SlbCurSslCfgCertsTableCurveNameEc_Sect233k1   SlbCurSslCfgCertsTableCurveNameEc = 30
	SlbCurSslCfgCertsTableCurveNameEc_Sect233r1   SlbCurSslCfgCertsTableCurveNameEc = 31
	SlbCurSslCfgCertsTableCurveNameEc_Sect239k1   SlbCurSslCfgCertsTableCurveNameEc = 32
	SlbCurSslCfgCertsTableCurveNameEc_Sect283k1   SlbCurSslCfgCertsTableCurveNameEc = 33
	SlbCurSslCfgCertsTableCurveNameEc_Sect283r1   SlbCurSslCfgCertsTableCurveNameEc = 34
	SlbCurSslCfgCertsTableCurveNameEc_Sect409k1   SlbCurSslCfgCertsTableCurveNameEc = 35
	SlbCurSslCfgCertsTableCurveNameEc_Sect409r1   SlbCurSslCfgCertsTableCurveNameEc = 36
	SlbCurSslCfgCertsTableCurveNameEc_Sect571k1   SlbCurSslCfgCertsTableCurveNameEc = 37
	SlbCurSslCfgCertsTableCurveNameEc_Sect571r1   SlbCurSslCfgCertsTableCurveNameEc = 38
	SlbCurSslCfgCertsTableCurveNameEc_C2pnb163v1  SlbCurSslCfgCertsTableCurveNameEc = 39
	SlbCurSslCfgCertsTableCurveNameEc_C2pnb163v2  SlbCurSslCfgCertsTableCurveNameEc = 40
	SlbCurSslCfgCertsTableCurveNameEc_C2pnb163v3  SlbCurSslCfgCertsTableCurveNameEc = 41
	SlbCurSslCfgCertsTableCurveNameEc_C2pnb176v1  SlbCurSslCfgCertsTableCurveNameEc = 42
	SlbCurSslCfgCertsTableCurveNameEc_C2tnb191v1  SlbCurSslCfgCertsTableCurveNameEc = 43
	SlbCurSslCfgCertsTableCurveNameEc_C2tnb191v2  SlbCurSslCfgCertsTableCurveNameEc = 44
	SlbCurSslCfgCertsTableCurveNameEc_C2tnb191v3  SlbCurSslCfgCertsTableCurveNameEc = 45
	SlbCurSslCfgCertsTableCurveNameEc_C2pnb208w1  SlbCurSslCfgCertsTableCurveNameEc = 46
	SlbCurSslCfgCertsTableCurveNameEc_C2tnb239v1  SlbCurSslCfgCertsTableCurveNameEc = 47
	SlbCurSslCfgCertsTableCurveNameEc_C2tnb239v2  SlbCurSslCfgCertsTableCurveNameEc = 48
	SlbCurSslCfgCertsTableCurveNameEc_C2tnb239v3  SlbCurSslCfgCertsTableCurveNameEc = 49
	SlbCurSslCfgCertsTableCurveNameEc_C2pnb272w1  SlbCurSslCfgCertsTableCurveNameEc = 50
	SlbCurSslCfgCertsTableCurveNameEc_C2pnb304w1  SlbCurSslCfgCertsTableCurveNameEc = 51
	SlbCurSslCfgCertsTableCurveNameEc_C2tnb359v1  SlbCurSslCfgCertsTableCurveNameEc = 52
	SlbCurSslCfgCertsTableCurveNameEc_C2pnb368w1  SlbCurSslCfgCertsTableCurveNameEc = 53
	SlbCurSslCfgCertsTableCurveNameEc_C2tnb431r1  SlbCurSslCfgCertsTableCurveNameEc = 54
	SlbCurSslCfgCertsTableCurveNameEc_Wtls1       SlbCurSslCfgCertsTableCurveNameEc = 55
	SlbCurSslCfgCertsTableCurveNameEc_Wtls3       SlbCurSslCfgCertsTableCurveNameEc = 56
	SlbCurSslCfgCertsTableCurveNameEc_Wtls4       SlbCurSslCfgCertsTableCurveNameEc = 57
	SlbCurSslCfgCertsTableCurveNameEc_Wtls5       SlbCurSslCfgCertsTableCurveNameEc = 58
	SlbCurSslCfgCertsTableCurveNameEc_Wtls6       SlbCurSslCfgCertsTableCurveNameEc = 59
	SlbCurSslCfgCertsTableCurveNameEc_Wtls7       SlbCurSslCfgCertsTableCurveNameEc = 60
	SlbCurSslCfgCertsTableCurveNameEc_Wtls8       SlbCurSslCfgCertsTableCurveNameEc = 61
	SlbCurSslCfgCertsTableCurveNameEc_Wtls9       SlbCurSslCfgCertsTableCurveNameEc = 62
	SlbCurSslCfgCertsTableCurveNameEc_Wtls10      SlbCurSslCfgCertsTableCurveNameEc = 63
	SlbCurSslCfgCertsTableCurveNameEc_Wtls11      SlbCurSslCfgCertsTableCurveNameEc = 64
	SlbCurSslCfgCertsTableCurveNameEc_Wtls12      SlbCurSslCfgCertsTableCurveNameEc = 65
	SlbCurSslCfgCertsTableCurveNameEc_Unsupported SlbCurSslCfgCertsTableCurveNameEc = 2147483647
)

type SlbCurSslCfgCertsTableParams struct {
	// Certificate id.
	ID string `json:"ID,omitempty"`
	// Certificate type.
	Type SlbCurSslCfgCertsTableType `json:"Type,omitempty"`
	// Certificate name.
	Name string `json:"Name,omitempty"`
	// Certificate key size.
	KeySize SlbCurSslCfgCertsTableKeySize `json:"KeySize,omitempty"`
	// Certificate expirty.
	Expirty string `json:"Expirty,omitempty"`
	// Certificate common name.
	CommonName string `json:"CommonName,omitempty"`
	// Certificate hash algorithm.
	HashAlgo SlbCurSslCfgCertsTableHashAlgo `json:"HashAlgo,omitempty"`
	// Certificate country name.
	CountryName string `json:"CountryName,omitempty"`
	// Certificate province name.
	PrpvinceName string `json:"PrpvinceName,omitempty"`
	// Certificate locality name.
	LocalityName string `json:"LocalityName,omitempty"`
	// Certificate organization name.
	OrganizationName string `json:"OrganizationName,omitempty"`
	// Certificate organization unit name.
	OrganizationUnitName string `json:"OrganizationUnitName,omitempty"`
	// Email address.
	EMail string `json:"EMail,omitempty"`
	// Certificate Validity period.
	ValidityPeriod uint32 `json:"ValidityPeriod,omitempty"`
	// Shows if the current certificate is generated.
	Status SlbCurSslCfgCertsTableStatus `json:"Status,omitempty"`
	// When this object is read the current state is returned.
	// Setting the value to generate(2),the Certificate will be generated.
	// idle(3) indicates that there is no Certificate generate is in progess.
	// generated(5) indicates that the last certificate generate operation is completed.
	// If the generate is successful this variable will return the state
	// 'generated' else it will return 'notGenerated'. In case of failure('notGenerated')
	// we should set this variable back to 'idle' state so others can issue the generate
	// command via SNMP.
	Generate SlbCurSslCfgCertsTableGenerate `json:"Generate,omitempty"`
	// Algorithm to generate the key.
	KeyType SlbCurSslCfgCertsTableKeyType `json:"KeyType,omitempty"`
	// EC key size.
	KeySizeEc SlbCurSslCfgCertsTableKeySizeEc `json:"KeySizeEc,omitempty"`
	// EC curve name.
	CurveNameEc SlbCurSslCfgCertsTableCurveNameEc `json:"CurveNameEc,omitempty"`
	// Current config Read-Only Key size common for both RSA and EC
	KeySizeCommon uint32 `json:"KeySizeCommon,omitempty"`
	// Intermediate CA certificate chain name.
	IntermcaChainName string `json:"IntermcaChainName,omitempty"`
	// Intermediate CA certificate chain type certificate=cert,Group=group,None=empty string.
	IntermcaChainType string `json:"IntermcaChainType,omitempty"`
	// SSL certificate serial number
	Serial string `json:"Serial,omitempty"`
	// Subject Alternative Name extension, format allowed (DNS:test.com), (IP:1.1.1.1), (email:test@mail.com)
	SubjectAltName string `json:"SubjectAltName,omitempty"`
	// Certificate province name.
	ProvinceName string `json:"ProvinceName,omitempty"`
	// Certificate expiry.
	Expiry string `json:"Expiry,omitempty"`
}

func (p SlbCurSslCfgCertsTableParams) iMABean() {}
