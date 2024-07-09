package beans

import (
	"fmt"
	"reflect"
)

// SlbNewSslCfgSSLPolTable The table for configuring ssl  policy.
// Note:This mib is not supported for VX instance of Virtualization.
type SlbNewSslCfgSSLPolTable struct {
	// The SSL policy name(key id) as an index.
	SlbNewSslCfgSSLPolNameIdIndex string
	Params                        *SlbNewSslCfgSSLPolTableParams
}

func NewSlbNewSslCfgSSLPolTableList() *SlbNewSslCfgSSLPolTable {
	return &SlbNewSslCfgSSLPolTable{}
}

func NewSlbNewSslCfgSSLPolTable(
	slbNewSslCfgSSLPolNameIdIndex string,
	params *SlbNewSslCfgSSLPolTableParams,
) *SlbNewSslCfgSSLPolTable {
	return &SlbNewSslCfgSSLPolTable{
		SlbNewSslCfgSSLPolNameIdIndex: slbNewSslCfgSSLPolNameIdIndex,
		Params:                        params,
	}
}

func (c *SlbNewSslCfgSSLPolTable) Name() string {
	return "SlbNewSslCfgSSLPolTable"
}

func (c *SlbNewSslCfgSSLPolTable) GetParams() BeanType {
	return c.Params
}

func (c *SlbNewSslCfgSSLPolTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *SlbNewSslCfgSSLPolTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.SlbNewSslCfgSSLPolNameIdIndex).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.SlbNewSslCfgSSLPolNameIdIndex)
}

type SlbNewSslCfgSSLPolTablePassInfoCipherFlag int32

const (
	SlbNewSslCfgSSLPolTablePassInfoCipherFlag_Enabled     SlbNewSslCfgSSLPolTablePassInfoCipherFlag = 1
	SlbNewSslCfgSSLPolTablePassInfoCipherFlag_Disabled    SlbNewSslCfgSSLPolTablePassInfoCipherFlag = 2
	SlbNewSslCfgSSLPolTablePassInfoCipherFlag_Unsupported SlbNewSslCfgSSLPolTablePassInfoCipherFlag = 2147483647
)

type SlbNewSslCfgSSLPolTablePassInfoVersionFlag int32

const (
	SlbNewSslCfgSSLPolTablePassInfoVersionFlag_Enabled     SlbNewSslCfgSSLPolTablePassInfoVersionFlag = 1
	SlbNewSslCfgSSLPolTablePassInfoVersionFlag_Disabled    SlbNewSslCfgSSLPolTablePassInfoVersionFlag = 2
	SlbNewSslCfgSSLPolTablePassInfoVersionFlag_Unsupported SlbNewSslCfgSSLPolTablePassInfoVersionFlag = 2147483647
)

type SlbNewSslCfgSSLPolTablePassInfoHeadBitsFlag int32

const (
	SlbNewSslCfgSSLPolTablePassInfoHeadBitsFlag_Enabled     SlbNewSslCfgSSLPolTablePassInfoHeadBitsFlag = 1
	SlbNewSslCfgSSLPolTablePassInfoHeadBitsFlag_Disabled    SlbNewSslCfgSSLPolTablePassInfoHeadBitsFlag = 2
	SlbNewSslCfgSSLPolTablePassInfoHeadBitsFlag_Unsupported SlbNewSslCfgSSLPolTablePassInfoHeadBitsFlag = 2147483647
)

type SlbNewSslCfgSSLPolTablePassInfoFrontend int32

const (
	SlbNewSslCfgSSLPolTablePassInfoFrontend_Enabled     SlbNewSslCfgSSLPolTablePassInfoFrontend = 1
	SlbNewSslCfgSSLPolTablePassInfoFrontend_Disabled    SlbNewSslCfgSSLPolTablePassInfoFrontend = 2
	SlbNewSslCfgSSLPolTablePassInfoFrontend_Unsupported SlbNewSslCfgSSLPolTablePassInfoFrontend = 2147483647
)

type SlbNewSslCfgSSLPolTableCipherName int32

const (
	SlbNewSslCfgSSLPolTableCipherName_Rsa               SlbNewSslCfgSSLPolTableCipherName = 0
	SlbNewSslCfgSSLPolTableCipherName_All               SlbNewSslCfgSSLPolTableCipherName = 1
	SlbNewSslCfgSSLPolTableCipherName_AllNonNullCiphers SlbNewSslCfgSSLPolTableCipherName = 2
	SlbNewSslCfgSSLPolTableCipherName_Sslv3             SlbNewSslCfgSSLPolTableCipherName = 3
	SlbNewSslCfgSSLPolTableCipherName_Tlsv1             SlbNewSslCfgSSLPolTableCipherName = 4
	SlbNewSslCfgSSLPolTableCipherName_Tlsv12            SlbNewSslCfgSSLPolTableCipherName = 5
	SlbNewSslCfgSSLPolTableCipherName_Export            SlbNewSslCfgSSLPolTableCipherName = 6
	SlbNewSslCfgSSLPolTableCipherName_Low               SlbNewSslCfgSSLPolTableCipherName = 7
	SlbNewSslCfgSSLPolTableCipherName_Medium            SlbNewSslCfgSSLPolTableCipherName = 8
	SlbNewSslCfgSSLPolTableCipherName_High              SlbNewSslCfgSSLPolTableCipherName = 9
	SlbNewSslCfgSSLPolTableCipherName_RsaRc4128Md5      SlbNewSslCfgSSLPolTableCipherName = 10
	SlbNewSslCfgSSLPolTableCipherName_RsaRc4128Sha1     SlbNewSslCfgSSLPolTableCipherName = 11
	SlbNewSslCfgSSLPolTableCipherName_RsaDesSha1        SlbNewSslCfgSSLPolTableCipherName = 12
	SlbNewSslCfgSSLPolTableCipherName_Rsa3desSha1       SlbNewSslCfgSSLPolTableCipherName = 13
	SlbNewSslCfgSSLPolTableCipherName_RsaAes128Sha1     SlbNewSslCfgSSLPolTableCipherName = 14
	SlbNewSslCfgSSLPolTableCipherName_RsaAes256Sha1     SlbNewSslCfgSSLPolTableCipherName = 15
	SlbNewSslCfgSSLPolTableCipherName_PciDssCompliance  SlbNewSslCfgSSLPolTableCipherName = 16
	SlbNewSslCfgSSLPolTableCipherName_UserDefined       SlbNewSslCfgSSLPolTableCipherName = 17
	SlbNewSslCfgSSLPolTableCipherName_UserDefinedExpert SlbNewSslCfgSSLPolTableCipherName = 18
	SlbNewSslCfgSSLPolTableCipherName_Main              SlbNewSslCfgSSLPolTableCipherName = 19
	SlbNewSslCfgSSLPolTableCipherName_Http2             SlbNewSslCfgSSLPolTableCipherName = 20
	SlbNewSslCfgSSLPolTableCipherName_Unsupported       SlbNewSslCfgSSLPolTableCipherName = 2147483647
)

type SlbNewSslCfgSSLPolTableBecipher int32

const (
	SlbNewSslCfgSSLPolTableBecipher_Low               SlbNewSslCfgSSLPolTableBecipher = 0
	SlbNewSslCfgSSLPolTableBecipher_Medium            SlbNewSslCfgSSLPolTableBecipher = 1
	SlbNewSslCfgSSLPolTableBecipher_High              SlbNewSslCfgSSLPolTableBecipher = 2
	SlbNewSslCfgSSLPolTableBecipher_UserDefined       SlbNewSslCfgSSLPolTableBecipher = 3
	SlbNewSslCfgSSLPolTableBecipher_UserDefinedExpert SlbNewSslCfgSSLPolTableBecipher = 4
	SlbNewSslCfgSSLPolTableBecipher_Main              SlbNewSslCfgSSLPolTableBecipher = 5
	SlbNewSslCfgSSLPolTableBecipher_Unsupported       SlbNewSslCfgSSLPolTableBecipher = 2147483647
)

type SlbNewSslCfgSSLPolTableBessl int32

const (
	SlbNewSslCfgSSLPolTableBessl_Enabled     SlbNewSslCfgSSLPolTableBessl = 1
	SlbNewSslCfgSSLPolTableBessl_Disabled    SlbNewSslCfgSSLPolTableBessl = 2
	SlbNewSslCfgSSLPolTableBessl_Request     SlbNewSslCfgSSLPolTableBessl = 3
	SlbNewSslCfgSSLPolTableBessl_Handshake   SlbNewSslCfgSSLPolTableBessl = 4
	SlbNewSslCfgSSLPolTableBessl_Unsupported SlbNewSslCfgSSLPolTableBessl = 2147483647
)

type SlbNewSslCfgSSLPolTableConvert int32

const (
	SlbNewSslCfgSSLPolTableConvert_Enabled     SlbNewSslCfgSSLPolTableConvert = 1
	SlbNewSslCfgSSLPolTableConvert_Disabled    SlbNewSslCfgSSLPolTableConvert = 2
	SlbNewSslCfgSSLPolTableConvert_Unsupported SlbNewSslCfgSSLPolTableConvert = 2147483647
)

type SlbNewSslCfgSSLPolTableAdminStatus int32

const (
	SlbNewSslCfgSSLPolTableAdminStatus_Enabled     SlbNewSslCfgSSLPolTableAdminStatus = 1
	SlbNewSslCfgSSLPolTableAdminStatus_Disabled    SlbNewSslCfgSSLPolTableAdminStatus = 2
	SlbNewSslCfgSSLPolTableAdminStatus_Unsupported SlbNewSslCfgSSLPolTableAdminStatus = 2147483647
)

type SlbNewSslCfgSSLPolTableDel int32

const (
	SlbNewSslCfgSSLPolTableDel_Other       SlbNewSslCfgSSLPolTableDel = 1
	SlbNewSslCfgSSLPolTableDel_Delete      SlbNewSslCfgSSLPolTableDel = 2
	SlbNewSslCfgSSLPolTableDel_Unsupported SlbNewSslCfgSSLPolTableDel = 2147483647
)

type SlbNewSslCfgSSLPolTablePassInfoComply int32

const (
	SlbNewSslCfgSSLPolTablePassInfoComply_Enabled     SlbNewSslCfgSSLPolTablePassInfoComply = 1
	SlbNewSslCfgSSLPolTablePassInfoComply_Disabled    SlbNewSslCfgSSLPolTablePassInfoComply = 2
	SlbNewSslCfgSSLPolTablePassInfoComply_Unsupported SlbNewSslCfgSSLPolTablePassInfoComply = 2147483647
)

type SlbNewSslCfgSSLPolTableFessl int32

const (
	SlbNewSslCfgSSLPolTableFessl_Enabled     SlbNewSslCfgSSLPolTableFessl = 1
	SlbNewSslCfgSSLPolTableFessl_Disabled    SlbNewSslCfgSSLPolTableFessl = 2
	SlbNewSslCfgSSLPolTableFessl_Request     SlbNewSslCfgSSLPolTableFessl = 3
	SlbNewSslCfgSSLPolTableFessl_Handshake   SlbNewSslCfgSSLPolTableFessl = 4
	SlbNewSslCfgSSLPolTableFessl_Unsupported SlbNewSslCfgSSLPolTableFessl = 2147483647
)

type SlbNewSslCfgSSLPolTableFESslv3Version int32

const (
	SlbNewSslCfgSSLPolTableFESslv3Version_Enabled     SlbNewSslCfgSSLPolTableFESslv3Version = 1
	SlbNewSslCfgSSLPolTableFESslv3Version_Disabled    SlbNewSslCfgSSLPolTableFESslv3Version = 2
	SlbNewSslCfgSSLPolTableFESslv3Version_Unsupported SlbNewSslCfgSSLPolTableFESslv3Version = 2147483647
)

type SlbNewSslCfgSSLPolTableFETls10Version int32

const (
	SlbNewSslCfgSSLPolTableFETls10Version_Enabled     SlbNewSslCfgSSLPolTableFETls10Version = 1
	SlbNewSslCfgSSLPolTableFETls10Version_Disabled    SlbNewSslCfgSSLPolTableFETls10Version = 2
	SlbNewSslCfgSSLPolTableFETls10Version_Unsupported SlbNewSslCfgSSLPolTableFETls10Version = 2147483647
)

type SlbNewSslCfgSSLPolTableFETls11Version int32

const (
	SlbNewSslCfgSSLPolTableFETls11Version_Enabled     SlbNewSslCfgSSLPolTableFETls11Version = 1
	SlbNewSslCfgSSLPolTableFETls11Version_Disabled    SlbNewSslCfgSSLPolTableFETls11Version = 2
	SlbNewSslCfgSSLPolTableFETls11Version_Unsupported SlbNewSslCfgSSLPolTableFETls11Version = 2147483647
)

type SlbNewSslCfgSSLPolTableBESslv3Version int32

const (
	SlbNewSslCfgSSLPolTableBESslv3Version_Enabled     SlbNewSslCfgSSLPolTableBESslv3Version = 1
	SlbNewSslCfgSSLPolTableBESslv3Version_Disabled    SlbNewSslCfgSSLPolTableBESslv3Version = 2
	SlbNewSslCfgSSLPolTableBESslv3Version_Unsupported SlbNewSslCfgSSLPolTableBESslv3Version = 2147483647
)

type SlbNewSslCfgSSLPolTableBETls10Version int32

const (
	SlbNewSslCfgSSLPolTableBETls10Version_Enabled     SlbNewSslCfgSSLPolTableBETls10Version = 1
	SlbNewSslCfgSSLPolTableBETls10Version_Disabled    SlbNewSslCfgSSLPolTableBETls10Version = 2
	SlbNewSslCfgSSLPolTableBETls10Version_Unsupported SlbNewSslCfgSSLPolTableBETls10Version = 2147483647
)

type SlbNewSslCfgSSLPolTableBETls11Version int32

const (
	SlbNewSslCfgSSLPolTableBETls11Version_Enabled     SlbNewSslCfgSSLPolTableBETls11Version = 1
	SlbNewSslCfgSSLPolTableBETls11Version_Disabled    SlbNewSslCfgSSLPolTableBETls11Version = 2
	SlbNewSslCfgSSLPolTableBETls11Version_Unsupported SlbNewSslCfgSSLPolTableBETls11Version = 2147483647
)

type SlbNewSslCfgSSLPolTableFETls12Version int32

const (
	SlbNewSslCfgSSLPolTableFETls12Version_Enabled     SlbNewSslCfgSSLPolTableFETls12Version = 1
	SlbNewSslCfgSSLPolTableFETls12Version_Disabled    SlbNewSslCfgSSLPolTableFETls12Version = 2
	SlbNewSslCfgSSLPolTableFETls12Version_Unsupported SlbNewSslCfgSSLPolTableFETls12Version = 2147483647
)

type SlbNewSslCfgSSLPolTableBETls12Version int32

const (
	SlbNewSslCfgSSLPolTableBETls12Version_Enabled     SlbNewSslCfgSSLPolTableBETls12Version = 1
	SlbNewSslCfgSSLPolTableBETls12Version_Disabled    SlbNewSslCfgSSLPolTableBETls12Version = 2
	SlbNewSslCfgSSLPolTableBETls12Version_Unsupported SlbNewSslCfgSSLPolTableBETls12Version = 2147483647
)

type SlbNewSslCfgSSLPolTableDHkey int32

const (
	SlbNewSslCfgSSLPolTableDHkey_KeySize1024 SlbNewSslCfgSSLPolTableDHkey = 1
	SlbNewSslCfgSSLPolTableDHkey_KeySize2048 SlbNewSslCfgSSLPolTableDHkey = 2
	SlbNewSslCfgSSLPolTableDHkey_Unsupported SlbNewSslCfgSSLPolTableDHkey = 2147483647
)

type SlbNewSslCfgSSLPolTableHwoffldFeRsa int32

const (
	SlbNewSslCfgSSLPolTableHwoffldFeRsa_Enabled     SlbNewSslCfgSSLPolTableHwoffldFeRsa = 1
	SlbNewSslCfgSSLPolTableHwoffldFeRsa_Disabled    SlbNewSslCfgSSLPolTableHwoffldFeRsa = 2
	SlbNewSslCfgSSLPolTableHwoffldFeRsa_Unsupported SlbNewSslCfgSSLPolTableHwoffldFeRsa = 2147483647
)

type SlbNewSslCfgSSLPolTableHwoffldFeDh int32

const (
	SlbNewSslCfgSSLPolTableHwoffldFeDh_Enabled     SlbNewSslCfgSSLPolTableHwoffldFeDh = 1
	SlbNewSslCfgSSLPolTableHwoffldFeDh_Disabled    SlbNewSslCfgSSLPolTableHwoffldFeDh = 2
	SlbNewSslCfgSSLPolTableHwoffldFeDh_Unsupported SlbNewSslCfgSSLPolTableHwoffldFeDh = 2147483647
)

type SlbNewSslCfgSSLPolTableHwoffldFeEc int32

const (
	SlbNewSslCfgSSLPolTableHwoffldFeEc_Enabled     SlbNewSslCfgSSLPolTableHwoffldFeEc = 1
	SlbNewSslCfgSSLPolTableHwoffldFeEc_Disabled    SlbNewSslCfgSSLPolTableHwoffldFeEc = 2
	SlbNewSslCfgSSLPolTableHwoffldFeEc_Unsupported SlbNewSslCfgSSLPolTableHwoffldFeEc = 2147483647
)

type SlbNewSslCfgSSLPolTableHwoffldFeBulk int32

const (
	SlbNewSslCfgSSLPolTableHwoffldFeBulk_Enabled     SlbNewSslCfgSSLPolTableHwoffldFeBulk = 1
	SlbNewSslCfgSSLPolTableHwoffldFeBulk_Disabled    SlbNewSslCfgSSLPolTableHwoffldFeBulk = 2
	SlbNewSslCfgSSLPolTableHwoffldFeBulk_Unsupported SlbNewSslCfgSSLPolTableHwoffldFeBulk = 2147483647
)

type SlbNewSslCfgSSLPolTableHwoffldFePkey int32

const (
	SlbNewSslCfgSSLPolTableHwoffldFePkey_Enabled     SlbNewSslCfgSSLPolTableHwoffldFePkey = 1
	SlbNewSslCfgSSLPolTableHwoffldFePkey_Disabled    SlbNewSslCfgSSLPolTableHwoffldFePkey = 2
	SlbNewSslCfgSSLPolTableHwoffldFePkey_Unsupported SlbNewSslCfgSSLPolTableHwoffldFePkey = 2147483647
)

type SlbNewSslCfgSSLPolTableHwoffldFeFeatures int32

const (
	SlbNewSslCfgSSLPolTableHwoffldFeFeatures_Enabled     SlbNewSslCfgSSLPolTableHwoffldFeFeatures = 1
	SlbNewSslCfgSSLPolTableHwoffldFeFeatures_Disabled    SlbNewSslCfgSSLPolTableHwoffldFeFeatures = 2
	SlbNewSslCfgSSLPolTableHwoffldFeFeatures_Unsupported SlbNewSslCfgSSLPolTableHwoffldFeFeatures = 2147483647
)

type SlbNewSslCfgSSLPolTableHwoffldBeRsa int32

const (
	SlbNewSslCfgSSLPolTableHwoffldBeRsa_Enabled     SlbNewSslCfgSSLPolTableHwoffldBeRsa = 1
	SlbNewSslCfgSSLPolTableHwoffldBeRsa_Disabled    SlbNewSslCfgSSLPolTableHwoffldBeRsa = 2
	SlbNewSslCfgSSLPolTableHwoffldBeRsa_Unsupported SlbNewSslCfgSSLPolTableHwoffldBeRsa = 2147483647
)

type SlbNewSslCfgSSLPolTableHwoffldBeDh int32

const (
	SlbNewSslCfgSSLPolTableHwoffldBeDh_Enabled     SlbNewSslCfgSSLPolTableHwoffldBeDh = 1
	SlbNewSslCfgSSLPolTableHwoffldBeDh_Disabled    SlbNewSslCfgSSLPolTableHwoffldBeDh = 2
	SlbNewSslCfgSSLPolTableHwoffldBeDh_Unsupported SlbNewSslCfgSSLPolTableHwoffldBeDh = 2147483647
)

type SlbNewSslCfgSSLPolTableHwoffldBeEc int32

const (
	SlbNewSslCfgSSLPolTableHwoffldBeEc_Enabled     SlbNewSslCfgSSLPolTableHwoffldBeEc = 1
	SlbNewSslCfgSSLPolTableHwoffldBeEc_Disabled    SlbNewSslCfgSSLPolTableHwoffldBeEc = 2
	SlbNewSslCfgSSLPolTableHwoffldBeEc_Unsupported SlbNewSslCfgSSLPolTableHwoffldBeEc = 2147483647
)

type SlbNewSslCfgSSLPolTableHwoffldBeBulk int32

const (
	SlbNewSslCfgSSLPolTableHwoffldBeBulk_Enabled     SlbNewSslCfgSSLPolTableHwoffldBeBulk = 1
	SlbNewSslCfgSSLPolTableHwoffldBeBulk_Disabled    SlbNewSslCfgSSLPolTableHwoffldBeBulk = 2
	SlbNewSslCfgSSLPolTableHwoffldBeBulk_Unsupported SlbNewSslCfgSSLPolTableHwoffldBeBulk = 2147483647
)

type SlbNewSslCfgSSLPolTableHwoffldBePkey int32

const (
	SlbNewSslCfgSSLPolTableHwoffldBePkey_Enabled     SlbNewSslCfgSSLPolTableHwoffldBePkey = 1
	SlbNewSslCfgSSLPolTableHwoffldBePkey_Disabled    SlbNewSslCfgSSLPolTableHwoffldBePkey = 2
	SlbNewSslCfgSSLPolTableHwoffldBePkey_Unsupported SlbNewSslCfgSSLPolTableHwoffldBePkey = 2147483647
)

type SlbNewSslCfgSSLPolTableHwoffldBeFeatures int32

const (
	SlbNewSslCfgSSLPolTableHwoffldBeFeatures_Enabled     SlbNewSslCfgSSLPolTableHwoffldBeFeatures = 1
	SlbNewSslCfgSSLPolTableHwoffldBeFeatures_Disabled    SlbNewSslCfgSSLPolTableHwoffldBeFeatures = 2
	SlbNewSslCfgSSLPolTableHwoffldBeFeatures_Unsupported SlbNewSslCfgSSLPolTableHwoffldBeFeatures = 2147483647
)

type SlbNewSslCfgSSLPolTableBesni int32

const (
	SlbNewSslCfgSSLPolTableBesni_Enabled     SlbNewSslCfgSSLPolTableBesni = 1
	SlbNewSslCfgSSLPolTableBesni_Disabled    SlbNewSslCfgSSLPolTableBesni = 2
	SlbNewSslCfgSSLPolTableBesni_Unsupported SlbNewSslCfgSSLPolTableBesni = 2147483647
)

type SlbNewSslCfgSSLPolTableFETls13Version int32

const (
	SlbNewSslCfgSSLPolTableFETls13Version_Enabled     SlbNewSslCfgSSLPolTableFETls13Version = 1
	SlbNewSslCfgSSLPolTableFETls13Version_Disabled    SlbNewSslCfgSSLPolTableFETls13Version = 2
	SlbNewSslCfgSSLPolTableFETls13Version_Unsupported SlbNewSslCfgSSLPolTableFETls13Version = 2147483647
)

type SlbNewSslCfgSSLPolTableBETls13Version int32

const (
	SlbNewSslCfgSSLPolTableBETls13Version_Enabled     SlbNewSslCfgSSLPolTableBETls13Version = 1
	SlbNewSslCfgSSLPolTableBETls13Version_Disabled    SlbNewSslCfgSSLPolTableBETls13Version = 2
	SlbNewSslCfgSSLPolTableBETls13Version_Unsupported SlbNewSslCfgSSLPolTableBETls13Version = 2147483647
)

type SlbNewSslCfgSSLPolTableFeReuseState int32

const (
	SlbNewSslCfgSSLPolTableFeReuseState_Enabled     SlbNewSslCfgSSLPolTableFeReuseState = 1
	SlbNewSslCfgSSLPolTableFeReuseState_Disabled    SlbNewSslCfgSSLPolTableFeReuseState = 2
	SlbNewSslCfgSSLPolTableFeReuseState_Inherit     SlbNewSslCfgSSLPolTableFeReuseState = 3
	SlbNewSslCfgSSLPolTableFeReuseState_Unsupported SlbNewSslCfgSSLPolTableFeReuseState = 2147483647
)

type SlbNewSslCfgSSLPolTableBeReuseState int32

const (
	SlbNewSslCfgSSLPolTableBeReuseState_Enabled     SlbNewSslCfgSSLPolTableBeReuseState = 1
	SlbNewSslCfgSSLPolTableBeReuseState_Disabled    SlbNewSslCfgSSLPolTableBeReuseState = 2
	SlbNewSslCfgSSLPolTableBeReuseState_Inherit     SlbNewSslCfgSSLPolTableBeReuseState = 3
	SlbNewSslCfgSSLPolTableBeReuseState_Unsupported SlbNewSslCfgSSLPolTableBeReuseState = 2147483647
)

type SlbNewSslCfgSSLPolTableBeReuseSrcMatch int32

const (
	SlbNewSslCfgSSLPolTableBeReuseSrcMatch_Enabled     SlbNewSslCfgSSLPolTableBeReuseSrcMatch = 1
	SlbNewSslCfgSSLPolTableBeReuseSrcMatch_Disabled    SlbNewSslCfgSSLPolTableBeReuseSrcMatch = 2
	SlbNewSslCfgSSLPolTableBeReuseSrcMatch_Unsupported SlbNewSslCfgSSLPolTableBeReuseSrcMatch = 2147483647
)

type SlbNewSslCfgSSLPolTableBeReuseTicket int32

const (
	SlbNewSslCfgSSLPolTableBeReuseTicket_Enabled     SlbNewSslCfgSSLPolTableBeReuseTicket = 1
	SlbNewSslCfgSSLPolTableBeReuseTicket_Disabled    SlbNewSslCfgSSLPolTableBeReuseTicket = 2
	SlbNewSslCfgSSLPolTableBeReuseTicket_Unsupported SlbNewSslCfgSSLPolTableBeReuseTicket = 2147483647
)

type SlbNewSslCfgSSLPolTableFeReuseTicket int32

const (
	SlbNewSslCfgSSLPolTableFeReuseTicket_Enabled     SlbNewSslCfgSSLPolTableFeReuseTicket = 1
	SlbNewSslCfgSSLPolTableFeReuseTicket_Disabled    SlbNewSslCfgSSLPolTableFeReuseTicket = 2
	SlbNewSslCfgSSLPolTableFeReuseTicket_Unsupported SlbNewSslCfgSSLPolTableFeReuseTicket = 2147483647
)

type SlbNewSslCfgSSLPolTableFEGmSslVersion int32

const (
	SlbNewSslCfgSSLPolTableFEGmSslVersion_Enabled     SlbNewSslCfgSSLPolTableFEGmSslVersion = 1
	SlbNewSslCfgSSLPolTableFEGmSslVersion_Disabled    SlbNewSslCfgSSLPolTableFEGmSslVersion = 2
	SlbNewSslCfgSSLPolTableFEGmSslVersion_Unsupported SlbNewSslCfgSSLPolTableFEGmSslVersion = 2147483647
)

type SlbNewSslCfgSSLPolTableBEGmSslVersion int32

const (
	SlbNewSslCfgSSLPolTableBEGmSslVersion_Enabled     SlbNewSslCfgSSLPolTableBEGmSslVersion = 1
	SlbNewSslCfgSSLPolTableBEGmSslVersion_Disabled    SlbNewSslCfgSSLPolTableBEGmSslVersion = 2
	SlbNewSslCfgSSLPolTableBEGmSslVersion_Unsupported SlbNewSslCfgSSLPolTableBEGmSslVersion = 2147483647
)

type SlbNewSslCfgSSLPolTableFEGmSslPriority int32

const (
	SlbNewSslCfgSSLPolTableFEGmSslPriority_Enabled     SlbNewSslCfgSSLPolTableFEGmSslPriority = 1
	SlbNewSslCfgSSLPolTableFEGmSslPriority_Disabled    SlbNewSslCfgSSLPolTableFEGmSslPriority = 2
	SlbNewSslCfgSSLPolTableFEGmSslPriority_Unsupported SlbNewSslCfgSSLPolTableFEGmSslPriority = 2147483647
)

type SlbNewSslCfgSSLPolTableParams struct {
	// The SSL policy name(key id) as an index.
	NameIdIndex string `json:"NameIdIndex,omitempty"`
	// SSL policy name,length of the string should be 32 characters.
	Name string `json:"Name,omitempty"`
	// The SSL cipher-suite header name.
	PassInfoCipherName string `json:"PassInfoCipherName,omitempty"`
	// Enable/Disable cipher-suite information to backend servers.
	PassInfoCipherFlag SlbNewSslCfgSSLPolTablePassInfoCipherFlag `json:"PassInfoCipherFlag,omitempty"`
	// SSL version header name.
	PassInfoVersionName string `json:"PassInfoVersionName,omitempty"`
	// Enable/Disable  SSL version information to backend servers.
	PassInfoVersionFlag SlbNewSslCfgSSLPolTablePassInfoVersionFlag `json:"PassInfoVersionFlag,omitempty"`
	// The passive cipher bits information to backend server.
	PassInfoHeadBitsName string `json:"PassInfoHeadBitsName,omitempty"`
	// Enable/Disable Cipher bits information to backend servers.
	PassInfoHeadBitsFlag SlbNewSslCfgSSLPolTablePassInfoHeadBitsFlag `json:"PassInfoHeadBitsFlag,omitempty"`
	// Enable/Disable add Front-End-Https: on header.
	PassInfoFrontend SlbNewSslCfgSSLPolTablePassInfoFrontend `json:"PassInfoFrontend,omitempty"`
	// Allowed cipher-suites in frontend SSL.
	CipherName SlbNewSslCfgSSLPolTableCipherName `json:"CipherName,omitempty"`
	// Cipher-suite allowed for SSL.
	CipherUserdef string `json:"CipherUserdef,omitempty"`
	// Intermediate CA certificate chain name.
	IntermcaChainName string `json:"IntermcaChainName,omitempty"`
	// Intermediate CA certificate chain type certificate=cert,Group=group,None=empty string.
	IntermcaChainType string `json:"IntermcaChainType,omitempty"`
	// Allowed cipher-suites in backend SSL.
	Becipher SlbNewSslCfgSSLPolTableBecipher `json:"Becipher,omitempty"`
	// Client authentication policy.
	Authpol string `json:"Authpol,omitempty"`
	// Host regex for HTTP redirection conversion.
	Convuri string `json:"Convuri,omitempty"`
	// Enable/Disable backend SSL encryption.
	Bessl SlbNewSslCfgSSLPolTableBessl `json:"Bessl,omitempty"`
	// Enable/Disable HTTP redirection conversion.
	Convert SlbNewSslCfgSSLPolTableConvert `json:"Convert,omitempty"`
	// Enable or disable ssl policy.
	AdminStatus SlbNewSslCfgSSLPolTableAdminStatus `json:"AdminStatus,omitempty"`
	// Delete SSL policy.
	Del SlbNewSslCfgSSLPolTableDel `json:"Del,omitempty"`
	// Enable/Disable X-SSL header compatible with 2424SSL headers.
	PassInfoComply SlbNewSslCfgSSLPolTablePassInfoComply `json:"PassInfoComply,omitempty"`
	// Set frontend SSL encryption mode, default value is enabled.
	Fessl SlbNewSslCfgSSLPolTableFessl `json:"Fessl,omitempty"`
	// Enable or disable frontend sslv3.
	FESslv3Version SlbNewSslCfgSSLPolTableFESslv3Version `json:"FESslv3Version,omitempty"`
	// Enable or disable frontend tls1_0.
	FETls10Version SlbNewSslCfgSSLPolTableFETls10Version `json:"FETls10Version,omitempty"`
	// Enable or disable frontend tls1_1.
	FETls11Version SlbNewSslCfgSSLPolTableFETls11Version `json:"FETls11Version,omitempty"`
	// Enable or disable backend sslv3.
	BESslv3Version SlbNewSslCfgSSLPolTableBESslv3Version `json:"BESslv3Version,omitempty"`
	// Enable or disable backend tls1_0.
	BETls10Version SlbNewSslCfgSSLPolTableBETls10Version `json:"BETls10Version,omitempty"`
	// Enable or disable backend tls1_1.
	BETls11Version SlbNewSslCfgSSLPolTableBETls11Version `json:"BETls11Version,omitempty"`
	// Enable or disable frontend tls1_2.
	FETls12Version SlbNewSslCfgSSLPolTableFETls12Version `json:"FETls12Version,omitempty"`
	// Enable or disable backend tls1_2.
	BETls12Version SlbNewSslCfgSSLPolTableBETls12Version `json:"BETls12Version,omitempty"`
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
	DHkey SlbNewSslCfgSSLPolTableDHkey `json:"DHkey,omitempty"`
	// Backend authentication policy.
	BEAuthpol string `json:"BEAuthpol,omitempty"`
	// Disable/Enable HW offload for RSA algorithm.
	HwoffldFeRsa SlbNewSslCfgSSLPolTableHwoffldFeRsa `json:"HwoffldFeRsa,omitempty"`
	// Disable/Enable HW offload for DH algorithm.
	HwoffldFeDh SlbNewSslCfgSSLPolTableHwoffldFeDh `json:"HwoffldFeDh,omitempty"`
	// Disable/Enable HW offload for EC algorithm.
	HwoffldFeEc SlbNewSslCfgSSLPolTableHwoffldFeEc `json:"HwoffldFeEc,omitempty"`
	// Disable/Enable HW offload for bulk encryption ciphers.
	HwoffldFeBulk SlbNewSslCfgSSLPolTableHwoffldFeBulk `json:"HwoffldFeBulk,omitempty"`
	// Disable/Enable HW offload for PKEY functionality.
	HwoffldFePkey SlbNewSslCfgSSLPolTableHwoffldFePkey `json:"HwoffldFePkey,omitempty"`
	// Disable/Enable HW offload Features.
	HwoffldFeFeatures SlbNewSslCfgSSLPolTableHwoffldFeFeatures `json:"HwoffldFeFeatures,omitempty"`
	// Disable/Enable HW offload for RSA algorithm.
	HwoffldBeRsa SlbNewSslCfgSSLPolTableHwoffldBeRsa `json:"HwoffldBeRsa,omitempty"`
	// Disable/Enable HW offload for DH algorithm.
	HwoffldBeDh SlbNewSslCfgSSLPolTableHwoffldBeDh `json:"HwoffldBeDh,omitempty"`
	// Disable/Enable HW offload for EC algorithm.
	HwoffldBeEc SlbNewSslCfgSSLPolTableHwoffldBeEc `json:"HwoffldBeEc,omitempty"`
	// Disable/Enable HW offload for bulk encryption ciphers.
	HwoffldBeBulk SlbNewSslCfgSSLPolTableHwoffldBeBulk `json:"HwoffldBeBulk,omitempty"`
	// Disable/Enable HW offload for PKEY functionality.
	HwoffldBePkey SlbNewSslCfgSSLPolTableHwoffldBePkey `json:"HwoffldBePkey,omitempty"`
	// Disable/Enable HW offload Features.
	HwoffldBeFeatures SlbNewSslCfgSSLPolTableHwoffldBeFeatures `json:"HwoffldBeFeatures,omitempty"`
	// Enable/Disable to include SNI.
	Besni SlbNewSslCfgSSLPolTableBesni `json:"Besni,omitempty"`
	// Enable or disable frontend tls1_3.
	FETls13Version SlbNewSslCfgSSLPolTableFETls13Version `json:"FETls13Version,omitempty"`
	// Enable or disable backend tls1_3.
	BETls13Version SlbNewSslCfgSSLPolTableBETls13Version `json:"BETls13Version,omitempty"`
	// Set maximum allowed early data on frontend connection
	SSLPol0RTTFEData uint32 `json:"SSLPol0RTTFEData,omitempty"`
	// Set Frontend SSL reuse state
	FeReuseState SlbNewSslCfgSSLPolTableFeReuseState `json:"FeReuseState,omitempty"`
	// Set Backend SSL reuse state
	BeReuseState SlbNewSslCfgSSLPolTableBeReuseState `json:"BeReuseState,omitempty"`
	// Enable/disable reuse for same client only
	BeReuseSrcMatch SlbNewSslCfgSSLPolTableBeReuseSrcMatch `json:"BeReuseSrcMatch,omitempty"`
	// Enable/disable TLS 1.2 session ticket
	BeReuseTicket SlbNewSslCfgSSLPolTableBeReuseTicket `json:"BeReuseTicket,omitempty"`
	// Enable/disable TLS 1.2 session ticket
	FeReuseTicket SlbNewSslCfgSSLPolTableFeReuseTicket `json:"FeReuseTicket,omitempty"`
	// Enable or disable frontend GM SSL.
	FEGmSslVersion SlbNewSslCfgSSLPolTableFEGmSslVersion `json:"FEGmSslVersion,omitempty"`
	// Enable or disable backend GM SSL.
	BEGmSslVersion SlbNewSslCfgSSLPolTableBEGmSslVersion `json:"BEGmSslVersion,omitempty"`
	// Enable or disable frontend GM SSL.
	FEGmSslPriority SlbNewSslCfgSSLPolTableFEGmSslPriority `json:"FEGmSslPriority,omitempty"`
	// Allowed signature algorithms.
	FESslsigs string `json:"FESslsigs,omitempty"`
	// Allowed signature algorithms in backend SSL.
	BESslsigs string `json:"BESslsigs,omitempty"`
	// Allowed groups.
	FESslgroups string `json:"FESslgroups,omitempty"`
	// Set allowed groups in backend SSL.
	BESslgroups string `json:"BESslgroups,omitempty"`
}

func (p SlbNewSslCfgSSLPolTableParams) iMABean() {}
