package beans

import (
	"fmt"
	"reflect"
)

type SlbNewSslCfgCertsTableSlbNewSslCfgCertsType int32

const (
	SlbNewSslCfgCertsTableSlbNewSslCfgCertsType_Key                     SlbNewSslCfgCertsTableSlbNewSslCfgCertsType = 1
	SlbNewSslCfgCertsTableSlbNewSslCfgCertsType_CertificateRequest      SlbNewSslCfgCertsTableSlbNewSslCfgCertsType = 2
	SlbNewSslCfgCertsTableSlbNewSslCfgCertsType_ServerCertificate       SlbNewSslCfgCertsTableSlbNewSslCfgCertsType = 3
	SlbNewSslCfgCertsTableSlbNewSslCfgCertsType_TrustedCertificate      SlbNewSslCfgCertsTableSlbNewSslCfgCertsType = 4
	SlbNewSslCfgCertsTableSlbNewSslCfgCertsType_IntermediateCertificate SlbNewSslCfgCertsTableSlbNewSslCfgCertsType = 5
	SlbNewSslCfgCertsTableSlbNewSslCfgCertsType_Crl                     SlbNewSslCfgCertsTableSlbNewSslCfgCertsType = 6
	SlbNewSslCfgCertsTableSlbNewSslCfgCertsType_Unsupported             SlbNewSslCfgCertsTableSlbNewSslCfgCertsType = 2147483647
)

// SlbNewSslCfgCertsTable Note:This mib is not supported for VX instance of Virtualization.
type SlbNewSslCfgCertsTable struct {
	// Certificate id.
	SlbNewSslCfgCertsID string
	// Certificate type.
	SlbNewSslCfgCertsType SlbNewSslCfgCertsTableSlbNewSslCfgCertsType
	Params                *SlbNewSslCfgCertsTableParams
}

func NewSlbNewSslCfgCertsTableList() *SlbNewSslCfgCertsTable {
	return &SlbNewSslCfgCertsTable{}
}

func NewSlbNewSslCfgCertsTable(
	slbNewSslCfgCertsID string,
	slbNewSslCfgCertsType SlbNewSslCfgCertsTableSlbNewSslCfgCertsType,
	params *SlbNewSslCfgCertsTableParams,
) *SlbNewSslCfgCertsTable {
	return &SlbNewSslCfgCertsTable{
		SlbNewSslCfgCertsID:   slbNewSslCfgCertsID,
		SlbNewSslCfgCertsType: slbNewSslCfgCertsType,
		Params:                params,
	}
}

func (c *SlbNewSslCfgCertsTable) Name() string {
	return "SlbNewSslCfgCertsTable"
}

func (c *SlbNewSslCfgCertsTable) GetParams() BeanType {
	return c.Params
}

func (c *SlbNewSslCfgCertsTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *SlbNewSslCfgCertsTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.SlbNewSslCfgCertsID).IsZero() &&
		reflect.ValueOf(c.SlbNewSslCfgCertsType).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.SlbNewSslCfgCertsID) + "/" + fmt.Sprint(c.SlbNewSslCfgCertsType)
}

type SlbNewSslCfgCertsTableType int32

const (
	SlbNewSslCfgCertsTableType_Key                     SlbNewSslCfgCertsTableType = 1
	SlbNewSslCfgCertsTableType_CertificateRequest      SlbNewSslCfgCertsTableType = 2
	SlbNewSslCfgCertsTableType_ServerCertificate       SlbNewSslCfgCertsTableType = 3
	SlbNewSslCfgCertsTableType_TrustedCertificate      SlbNewSslCfgCertsTableType = 4
	SlbNewSslCfgCertsTableType_IntermediateCertificate SlbNewSslCfgCertsTableType = 5
	SlbNewSslCfgCertsTableType_Crl                     SlbNewSslCfgCertsTableType = 6
	SlbNewSslCfgCertsTableType_Unsupported             SlbNewSslCfgCertsTableType = 2147483647
)

type SlbNewSslCfgCertsTableKeySize int32

const (
	SlbNewSslCfgCertsTableKeySize_Ks512       SlbNewSslCfgCertsTableKeySize = 1
	SlbNewSslCfgCertsTableKeySize_Ks1024      SlbNewSslCfgCertsTableKeySize = 2
	SlbNewSslCfgCertsTableKeySize_Ks2048      SlbNewSslCfgCertsTableKeySize = 3
	SlbNewSslCfgCertsTableKeySize_Ks4096      SlbNewSslCfgCertsTableKeySize = 4
	SlbNewSslCfgCertsTableKeySize_Unknown     SlbNewSslCfgCertsTableKeySize = 6
	SlbNewSslCfgCertsTableKeySize_Unsupported SlbNewSslCfgCertsTableKeySize = 2147483647
)

type SlbNewSslCfgCertsTableHashAlgo int32

const (
	SlbNewSslCfgCertsTableHashAlgo_Md5         SlbNewSslCfgCertsTableHashAlgo = 1
	SlbNewSslCfgCertsTableHashAlgo_Sha1        SlbNewSslCfgCertsTableHashAlgo = 2
	SlbNewSslCfgCertsTableHashAlgo_Sha256      SlbNewSslCfgCertsTableHashAlgo = 3
	SlbNewSslCfgCertsTableHashAlgo_Sha384      SlbNewSslCfgCertsTableHashAlgo = 4
	SlbNewSslCfgCertsTableHashAlgo_Sha512      SlbNewSslCfgCertsTableHashAlgo = 5
	SlbNewSslCfgCertsTableHashAlgo_Unknown     SlbNewSslCfgCertsTableHashAlgo = 6
	SlbNewSslCfgCertsTableHashAlgo_Sm3         SlbNewSslCfgCertsTableHashAlgo = 7
	SlbNewSslCfgCertsTableHashAlgo_Unsupported SlbNewSslCfgCertsTableHashAlgo = 2147483647
)

type SlbNewSslCfgCertsTableDelete int32

const (
	SlbNewSslCfgCertsTableDelete_Other       SlbNewSslCfgCertsTableDelete = 1
	SlbNewSslCfgCertsTableDelete_Delete      SlbNewSslCfgCertsTableDelete = 2
	SlbNewSslCfgCertsTableDelete_Unsupported SlbNewSslCfgCertsTableDelete = 2147483647
)

type SlbNewSslCfgCertsTableGenerate int32

const (
	SlbNewSslCfgCertsTableGenerate_Other        SlbNewSslCfgCertsTableGenerate = 1
	SlbNewSslCfgCertsTableGenerate_Generate     SlbNewSslCfgCertsTableGenerate = 2
	SlbNewSslCfgCertsTableGenerate_Idle         SlbNewSslCfgCertsTableGenerate = 3
	SlbNewSslCfgCertsTableGenerate_Inprogress   SlbNewSslCfgCertsTableGenerate = 4
	SlbNewSslCfgCertsTableGenerate_Generated    SlbNewSslCfgCertsTableGenerate = 5
	SlbNewSslCfgCertsTableGenerate_NotGenerated SlbNewSslCfgCertsTableGenerate = 6
	SlbNewSslCfgCertsTableGenerate_Unsupported  SlbNewSslCfgCertsTableGenerate = 2147483647
)

type SlbNewSslCfgCertsTableStatus int32

const (
	SlbNewSslCfgCertsTableStatus_Generated    SlbNewSslCfgCertsTableStatus = 1
	SlbNewSslCfgCertsTableStatus_NotGenerated SlbNewSslCfgCertsTableStatus = 2
	SlbNewSslCfgCertsTableStatus_InProgress   SlbNewSslCfgCertsTableStatus = 3
	SlbNewSslCfgCertsTableStatus_Unsupported  SlbNewSslCfgCertsTableStatus = 2147483647
)

type SlbNewSslCfgCertsTableKeyType int32

const (
	SlbNewSslCfgCertsTableKeyType_Rsa         SlbNewSslCfgCertsTableKeyType = 1
	SlbNewSslCfgCertsTableKeyType_Ec          SlbNewSslCfgCertsTableKeyType = 2
	SlbNewSslCfgCertsTableKeyType_Unknown     SlbNewSslCfgCertsTableKeyType = 6
	SlbNewSslCfgCertsTableKeyType_Sm2         SlbNewSslCfgCertsTableKeyType = 7
	SlbNewSslCfgCertsTableKeyType_Unsupported SlbNewSslCfgCertsTableKeyType = 2147483647
)

type SlbNewSslCfgCertsTableKeySizeEc int32

const (
	SlbNewSslCfgCertsTableKeySizeEc_Ks0         SlbNewSslCfgCertsTableKeySizeEc = 0
	SlbNewSslCfgCertsTableKeySizeEc_Ks192       SlbNewSslCfgCertsTableKeySizeEc = 1
	SlbNewSslCfgCertsTableKeySizeEc_Ks224       SlbNewSslCfgCertsTableKeySizeEc = 2
	SlbNewSslCfgCertsTableKeySizeEc_Ks256       SlbNewSslCfgCertsTableKeySizeEc = 3
	SlbNewSslCfgCertsTableKeySizeEc_Ks384       SlbNewSslCfgCertsTableKeySizeEc = 4
	SlbNewSslCfgCertsTableKeySizeEc_Ks521       SlbNewSslCfgCertsTableKeySizeEc = 5
	SlbNewSslCfgCertsTableKeySizeEc_Unknown     SlbNewSslCfgCertsTableKeySizeEc = 6
	SlbNewSslCfgCertsTableKeySizeEc_Unsupported SlbNewSslCfgCertsTableKeySizeEc = 2147483647
)

type SlbNewSslCfgCertsTableCurveNameEc int32

const (
	SlbNewSslCfgCertsTableCurveNameEc_Unknown     SlbNewSslCfgCertsTableCurveNameEc = 0
	SlbNewSslCfgCertsTableCurveNameEc_Secp112r1   SlbNewSslCfgCertsTableCurveNameEc = 1
	SlbNewSslCfgCertsTableCurveNameEc_Secp112r2   SlbNewSslCfgCertsTableCurveNameEc = 2
	SlbNewSslCfgCertsTableCurveNameEc_Secp128r1   SlbNewSslCfgCertsTableCurveNameEc = 3
	SlbNewSslCfgCertsTableCurveNameEc_Secp128r2   SlbNewSslCfgCertsTableCurveNameEc = 4
	SlbNewSslCfgCertsTableCurveNameEc_Secp160k1   SlbNewSslCfgCertsTableCurveNameEc = 5
	SlbNewSslCfgCertsTableCurveNameEc_Secp160r1   SlbNewSslCfgCertsTableCurveNameEc = 6
	SlbNewSslCfgCertsTableCurveNameEc_Secp160r2   SlbNewSslCfgCertsTableCurveNameEc = 7
	SlbNewSslCfgCertsTableCurveNameEc_Secp192k1   SlbNewSslCfgCertsTableCurveNameEc = 8
	SlbNewSslCfgCertsTableCurveNameEc_Secp224k1   SlbNewSslCfgCertsTableCurveNameEc = 9
	SlbNewSslCfgCertsTableCurveNameEc_Secp224r1   SlbNewSslCfgCertsTableCurveNameEc = 10
	SlbNewSslCfgCertsTableCurveNameEc_Secp256k1   SlbNewSslCfgCertsTableCurveNameEc = 11
	SlbNewSslCfgCertsTableCurveNameEc_Secp384r1   SlbNewSslCfgCertsTableCurveNameEc = 12
	SlbNewSslCfgCertsTableCurveNameEc_Secp521r1   SlbNewSslCfgCertsTableCurveNameEc = 13
	SlbNewSslCfgCertsTableCurveNameEc_Prime192v1  SlbNewSslCfgCertsTableCurveNameEc = 14
	SlbNewSslCfgCertsTableCurveNameEc_Prime192v2  SlbNewSslCfgCertsTableCurveNameEc = 15
	SlbNewSslCfgCertsTableCurveNameEc_Prime192v3  SlbNewSslCfgCertsTableCurveNameEc = 16
	SlbNewSslCfgCertsTableCurveNameEc_Prime239v1  SlbNewSslCfgCertsTableCurveNameEc = 17
	SlbNewSslCfgCertsTableCurveNameEc_Prime239v2  SlbNewSslCfgCertsTableCurveNameEc = 18
	SlbNewSslCfgCertsTableCurveNameEc_Prime239v3  SlbNewSslCfgCertsTableCurveNameEc = 19
	SlbNewSslCfgCertsTableCurveNameEc_Prime256v1  SlbNewSslCfgCertsTableCurveNameEc = 20
	SlbNewSslCfgCertsTableCurveNameEc_Sect113r1   SlbNewSslCfgCertsTableCurveNameEc = 21
	SlbNewSslCfgCertsTableCurveNameEc_Sect113r2   SlbNewSslCfgCertsTableCurveNameEc = 22
	SlbNewSslCfgCertsTableCurveNameEc_Sect131r1   SlbNewSslCfgCertsTableCurveNameEc = 23
	SlbNewSslCfgCertsTableCurveNameEc_Sect131r2   SlbNewSslCfgCertsTableCurveNameEc = 24
	SlbNewSslCfgCertsTableCurveNameEc_Sect163k1   SlbNewSslCfgCertsTableCurveNameEc = 25
	SlbNewSslCfgCertsTableCurveNameEc_Sect163r1   SlbNewSslCfgCertsTableCurveNameEc = 26
	SlbNewSslCfgCertsTableCurveNameEc_Sect163r2   SlbNewSslCfgCertsTableCurveNameEc = 27
	SlbNewSslCfgCertsTableCurveNameEc_Sect193r1   SlbNewSslCfgCertsTableCurveNameEc = 28
	SlbNewSslCfgCertsTableCurveNameEc_Sect193r2   SlbNewSslCfgCertsTableCurveNameEc = 29
	SlbNewSslCfgCertsTableCurveNameEc_Sect233k1   SlbNewSslCfgCertsTableCurveNameEc = 30
	SlbNewSslCfgCertsTableCurveNameEc_Sect233r1   SlbNewSslCfgCertsTableCurveNameEc = 31
	SlbNewSslCfgCertsTableCurveNameEc_Sect239k1   SlbNewSslCfgCertsTableCurveNameEc = 32
	SlbNewSslCfgCertsTableCurveNameEc_Sect283k1   SlbNewSslCfgCertsTableCurveNameEc = 33
	SlbNewSslCfgCertsTableCurveNameEc_Sect283r1   SlbNewSslCfgCertsTableCurveNameEc = 34
	SlbNewSslCfgCertsTableCurveNameEc_Sect409k1   SlbNewSslCfgCertsTableCurveNameEc = 35
	SlbNewSslCfgCertsTableCurveNameEc_Sect409r1   SlbNewSslCfgCertsTableCurveNameEc = 36
	SlbNewSslCfgCertsTableCurveNameEc_Sect571k1   SlbNewSslCfgCertsTableCurveNameEc = 37
	SlbNewSslCfgCertsTableCurveNameEc_Sect571r1   SlbNewSslCfgCertsTableCurveNameEc = 38
	SlbNewSslCfgCertsTableCurveNameEc_C2pnb163v1  SlbNewSslCfgCertsTableCurveNameEc = 39
	SlbNewSslCfgCertsTableCurveNameEc_C2pnb163v2  SlbNewSslCfgCertsTableCurveNameEc = 40
	SlbNewSslCfgCertsTableCurveNameEc_C2pnb163v3  SlbNewSslCfgCertsTableCurveNameEc = 41
	SlbNewSslCfgCertsTableCurveNameEc_C2pnb176v1  SlbNewSslCfgCertsTableCurveNameEc = 42
	SlbNewSslCfgCertsTableCurveNameEc_C2tnb191v1  SlbNewSslCfgCertsTableCurveNameEc = 43
	SlbNewSslCfgCertsTableCurveNameEc_C2tnb191v2  SlbNewSslCfgCertsTableCurveNameEc = 44
	SlbNewSslCfgCertsTableCurveNameEc_C2tnb191v3  SlbNewSslCfgCertsTableCurveNameEc = 45
	SlbNewSslCfgCertsTableCurveNameEc_C2pnb208w1  SlbNewSslCfgCertsTableCurveNameEc = 46
	SlbNewSslCfgCertsTableCurveNameEc_C2tnb239v1  SlbNewSslCfgCertsTableCurveNameEc = 47
	SlbNewSslCfgCertsTableCurveNameEc_C2tnb239v2  SlbNewSslCfgCertsTableCurveNameEc = 48
	SlbNewSslCfgCertsTableCurveNameEc_C2tnb239v3  SlbNewSslCfgCertsTableCurveNameEc = 49
	SlbNewSslCfgCertsTableCurveNameEc_C2pnb272w1  SlbNewSslCfgCertsTableCurveNameEc = 50
	SlbNewSslCfgCertsTableCurveNameEc_C2pnb304w1  SlbNewSslCfgCertsTableCurveNameEc = 51
	SlbNewSslCfgCertsTableCurveNameEc_C2tnb359v1  SlbNewSslCfgCertsTableCurveNameEc = 52
	SlbNewSslCfgCertsTableCurveNameEc_C2pnb368w1  SlbNewSslCfgCertsTableCurveNameEc = 53
	SlbNewSslCfgCertsTableCurveNameEc_C2tnb431r1  SlbNewSslCfgCertsTableCurveNameEc = 54
	SlbNewSslCfgCertsTableCurveNameEc_Wtls1       SlbNewSslCfgCertsTableCurveNameEc = 55
	SlbNewSslCfgCertsTableCurveNameEc_Wtls3       SlbNewSslCfgCertsTableCurveNameEc = 56
	SlbNewSslCfgCertsTableCurveNameEc_Wtls4       SlbNewSslCfgCertsTableCurveNameEc = 57
	SlbNewSslCfgCertsTableCurveNameEc_Wtls5       SlbNewSslCfgCertsTableCurveNameEc = 58
	SlbNewSslCfgCertsTableCurveNameEc_Wtls6       SlbNewSslCfgCertsTableCurveNameEc = 59
	SlbNewSslCfgCertsTableCurveNameEc_Wtls7       SlbNewSslCfgCertsTableCurveNameEc = 60
	SlbNewSslCfgCertsTableCurveNameEc_Wtls8       SlbNewSslCfgCertsTableCurveNameEc = 61
	SlbNewSslCfgCertsTableCurveNameEc_Wtls9       SlbNewSslCfgCertsTableCurveNameEc = 62
	SlbNewSslCfgCertsTableCurveNameEc_Wtls10      SlbNewSslCfgCertsTableCurveNameEc = 63
	SlbNewSslCfgCertsTableCurveNameEc_Wtls11      SlbNewSslCfgCertsTableCurveNameEc = 64
	SlbNewSslCfgCertsTableCurveNameEc_Wtls12      SlbNewSslCfgCertsTableCurveNameEc = 65
	SlbNewSslCfgCertsTableCurveNameEc_Sm          SlbNewSslCfgCertsTableCurveNameEc = 66
	SlbNewSslCfgCertsTableCurveNameEc_Unsupported SlbNewSslCfgCertsTableCurveNameEc = 2147483647
)

type SlbNewSslCfgCertsTableParams struct {
	// Certificate id.
	ID string `json:"ID,omitempty"`
	// Certificate type.
	Type SlbNewSslCfgCertsTableType `json:"Type,omitempty"`
	// Certificate Name.
	Name string `json:"Name,omitempty"`
	// Certificate key size.
	KeySize SlbNewSslCfgCertsTableKeySize `json:"KeySize,omitempty"`
	// Certificate expiry.
	Expirty string `json:"Expirty,omitempty"`
	// Certificate common name.
	CommonName string `json:"CommonName,omitempty"`
	// Certificate hash algorithm.
	HashAlgo SlbNewSslCfgCertsTableHashAlgo `json:"HashAlgo,omitempty"`
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
	// Certificate validity period.
	ValidityPeriod uint32 `json:"ValidityPeriod,omitempty"`
	// Delete certificate.
	Delete SlbNewSslCfgCertsTableDelete `json:"Delete,omitempty"`
	// When this object is read the current state is returned.
	// Setting the value to generate(2),the Certificate will be generated.
	// idle(3) indicates that there is no Certificate generate is in progess.
	// generated(5) indicates that the last certificate generate operation is completed.
	// If the generate is successful this variable will return the state
	// 'generated' else it will return 'notGenerated'. In case of failure('notGenerated')
	// we should set this variable back to 'idle' state so others can issue the generate
	// command via SNMP.
	Generate SlbNewSslCfgCertsTableGenerate `json:"Generate,omitempty"`
	// Shows if the current certificate is generated.
	Status SlbNewSslCfgCertsTableStatus `json:"Status,omitempty"`
	// Algorithm to generate the key.
	KeyType SlbNewSslCfgCertsTableKeyType `json:"KeyType,omitempty"`
	// EC key size.
	KeySizeEc SlbNewSslCfgCertsTableKeySizeEc `json:"KeySizeEc,omitempty"`
	// EC curve name.
	CurveNameEc SlbNewSslCfgCertsTableCurveNameEc `json:"CurveNameEc,omitempty"`
	// New config Read-Only Key size common for both RSA and EC
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

func (p SlbNewSslCfgCertsTableParams) iMABean() {}
