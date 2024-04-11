package beans

import (
	"fmt"
	"reflect"
)

// SlbCurSslCfgGroupsTable Note:This mib is not supported for VX instance of Virtualization.
type SlbCurSslCfgGroupsTable struct {
	// Certificate group ID.
	SlbCurSslCfgGroupsID string
	Params               *SlbCurSslCfgGroupsTableParams
}

func NewSlbCurSslCfgGroupsTableList() *SlbCurSslCfgGroupsTable {
	return &SlbCurSslCfgGroupsTable{}
}

func NewSlbCurSslCfgGroupsTable(
	slbCurSslCfgGroupsID string,
	params *SlbCurSslCfgGroupsTableParams,
) *SlbCurSslCfgGroupsTable {
	return &SlbCurSslCfgGroupsTable{
		SlbCurSslCfgGroupsID: slbCurSslCfgGroupsID,
		Params:               params,
	}
}

func (c *SlbCurSslCfgGroupsTable) Name() string {
	return "SlbCurSslCfgGroupsTable"
}

func (c *SlbCurSslCfgGroupsTable) GetParams() BeanType {
	return c.Params
}

func (c *SlbCurSslCfgGroupsTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *SlbCurSslCfgGroupsTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.SlbCurSslCfgGroupsID).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.SlbCurSslCfgGroupsID)
}

type SlbCurSslCfgGroupsTableType int32

const (
	SlbCurSslCfgGroupsTableType_ServerCertificate       SlbCurSslCfgGroupsTableType = 3
	SlbCurSslCfgGroupsTableType_TrustedCertificate      SlbCurSslCfgGroupsTableType = 4
	SlbCurSslCfgGroupsTableType_IntermediateCertificate SlbCurSslCfgGroupsTableType = 5
	SlbCurSslCfgGroupsTableType_Unsupported             SlbCurSslCfgGroupsTableType = 2147483647
)

type SlbCurSslCfgGroupsTableConfigType int32

const (
	SlbCurSslCfgGroupsTableConfigType_Regular     SlbCurSslCfgGroupsTableConfigType = 1
	SlbCurSslCfgGroupsTableConfigType_ReadOnly    SlbCurSslCfgGroupsTableConfigType = 2
	SlbCurSslCfgGroupsTableConfigType_Unsupported SlbCurSslCfgGroupsTableConfigType = 2147483647
)

type SlbCurSslCfgGroupsTableChainingMode int32

const (
	SlbCurSslCfgGroupsTableChainingMode_BySubjectIssuer SlbCurSslCfgGroupsTableChainingMode = 0
	SlbCurSslCfgGroupsTableChainingMode_BySkidAkid      SlbCurSslCfgGroupsTableChainingMode = 1
	SlbCurSslCfgGroupsTableChainingMode_Unsupported     SlbCurSslCfgGroupsTableChainingMode = 2147483647
)

type SlbCurSslCfgGroupsTableParams struct {
	// Certificate group ID.
	ID string `json:"ID,omitempty"`
	// Certificate group name.
	Name string `json:"Name,omitempty"`
	// Certificate group type.
	Type SlbCurSslCfgGroupsTableType `json:"Type,omitempty"`
	// Default certificate.
	DefaultCert string `json:"DefaultCert,omitempty"`
	// List of Certificates the group holds - Selected certificates
	// are presented in a bitmap format.
	// Receiving order is as:
	// 	 OCTET 1  OCTET 2  .....
	// 	 xxxxxxxx xxxxxxxx .....
	// 	 ||    || |_ Cert 9
	// 	 ||    ||
	// 	 ||    ||___ Cert 4
	// 	 ||    |____ Cert 3
	// 	 ||      .    .   .
	// 	 ||_________ Cert 2
	// 	 |__________ Cert 1
	// 	 where x : 1 - The represented certficate is selected
	// 	 0 - The represented certificate is not selected.
	CertBmap string `json:"CertBmap,omitempty"`
	// Group config type either regular or read-only.
	ConfigType SlbCurSslCfgGroupsTableConfigType `json:"ConfigType,omitempty"`
	// Chain Mode. By key - 0 or By Group - 1.
	ChainingMode SlbCurSslCfgGroupsTableChainingMode `json:"ChainingMode,omitempty"`
}

func (p SlbCurSslCfgGroupsTableParams) iMABean() {}
