package beans

import (
	"fmt"
	"reflect"
)

// SlbNewSslCfgGroupsTable Note:This mib is not supported for VX instance of Virtualization.
type SlbNewSslCfgGroupsTable struct {
	// Certificate group ID.
	SlbNewSslCfgGroupsID string
	Params               *SlbNewSslCfgGroupsTableParams
}

func NewSlbNewSslCfgGroupsTableList() *SlbNewSslCfgGroupsTable {
	return &SlbNewSslCfgGroupsTable{}
}

func NewSlbNewSslCfgGroupsTable(
	slbNewSslCfgGroupsID string,
	params *SlbNewSslCfgGroupsTableParams,
) *SlbNewSslCfgGroupsTable {
	return &SlbNewSslCfgGroupsTable{
		SlbNewSslCfgGroupsID: slbNewSslCfgGroupsID,
		Params:               params,
	}
}

func (c *SlbNewSslCfgGroupsTable) Name() string {
	return "SlbNewSslCfgGroupsTable"
}

func (c *SlbNewSslCfgGroupsTable) GetParams() BeanType {
	return c.Params
}

func (c *SlbNewSslCfgGroupsTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *SlbNewSslCfgGroupsTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.SlbNewSslCfgGroupsID).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.SlbNewSslCfgGroupsID)
}

type SlbNewSslCfgGroupsTableType int32

const (
	SlbNewSslCfgGroupsTableType_ServerCertificate       SlbNewSslCfgGroupsTableType = 3
	SlbNewSslCfgGroupsTableType_TrustedCertificate      SlbNewSslCfgGroupsTableType = 4
	SlbNewSslCfgGroupsTableType_IntermediateCertificate SlbNewSslCfgGroupsTableType = 5
	SlbNewSslCfgGroupsTableType_Unsupported             SlbNewSslCfgGroupsTableType = 2147483647
)

type SlbNewSslCfgGroupsTableDelete int32

const (
	SlbNewSslCfgGroupsTableDelete_Other       SlbNewSslCfgGroupsTableDelete = 1
	SlbNewSslCfgGroupsTableDelete_Delete      SlbNewSslCfgGroupsTableDelete = 2
	SlbNewSslCfgGroupsTableDelete_Unsupported SlbNewSslCfgGroupsTableDelete = 2147483647
)

type SlbNewSslCfgGroupsTableConfigType int32

const (
	SlbNewSslCfgGroupsTableConfigType_Regular     SlbNewSslCfgGroupsTableConfigType = 1
	SlbNewSslCfgGroupsTableConfigType_ReadOnly    SlbNewSslCfgGroupsTableConfigType = 2
	SlbNewSslCfgGroupsTableConfigType_Unsupported SlbNewSslCfgGroupsTableConfigType = 2147483647
)

type SlbNewSslCfgGroupsTableChainingMode int32

const (
	SlbNewSslCfgGroupsTableChainingMode_BySubjectIssuer SlbNewSslCfgGroupsTableChainingMode = 0
	SlbNewSslCfgGroupsTableChainingMode_BySkidAkid      SlbNewSslCfgGroupsTableChainingMode = 1
	SlbNewSslCfgGroupsTableChainingMode_Unsupported     SlbNewSslCfgGroupsTableChainingMode = 2147483647
)

type SlbNewSslCfgGroupsTableParams struct {
	// Certificate group ID.
	ID string `json:"ID,omitempty"`
	// Certificate group name.
	Name string `json:"Name,omitempty"`
	// Certificate group type.
	Type SlbNewSslCfgGroupsTableType `json:"Type,omitempty"`
	// Delete certificate group.
	Delete SlbNewSslCfgGroupsTableDelete `json:"Delete,omitempty"`
	// Default certificate.
	DefaultCert string `json:"DefaultCert,omitempty"`
	// Add certificate to the group
	AddCert string `json:"AddCert,omitempty"`
	// Remove certificate from the group
	RemCert string `json:"RemCert,omitempty"`
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
	ConfigType SlbNewSslCfgGroupsTableConfigType `json:"ConfigType,omitempty"`
	// Chain Mode. By key - 0 or By Group - 1.
	ChainingMode SlbNewSslCfgGroupsTableChainingMode `json:"ChainingMode,omitempty"`
}

func (p SlbNewSslCfgGroupsTableParams) iMABean() {}
