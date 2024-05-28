package beans

import (
	"fmt"
	"reflect"
)

// SlbSessionInfoTable The table of SLB session entries.
// Note:This mib is not supported for VX instance of Virtualization.
type SlbSessionInfoTable struct {
	// The SP which created the session.
	SlbSessionInfoSpIndex int32
	// The index to the slbSessionInfoTable.
	SlbSessionInfoIndex int32
	Params              *SlbSessionInfoTableParams
}

func NewSlbSessionInfoTableList() *SlbSessionInfoTable {
	return &SlbSessionInfoTable{}
}

func NewSlbSessionInfoTable(
	slbSessionInfoSpIndex int32,
	slbSessionInfoIndex int32,
	params *SlbSessionInfoTableParams,
) *SlbSessionInfoTable {
	return &SlbSessionInfoTable{
		SlbSessionInfoSpIndex: slbSessionInfoSpIndex,
		SlbSessionInfoIndex:   slbSessionInfoIndex,
		Params:                params,
	}
}

func (c *SlbSessionInfoTable) Name() string {
	return "SlbSessionInfoTable"
}

func (c *SlbSessionInfoTable) GetParams() BeanType {
	return c.Params
}

func (c *SlbSessionInfoTable) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}

func (c *SlbSessionInfoTable) Path() string {
	path := "/config/" + c.Name()
	if reflect.ValueOf(c.SlbSessionInfoSpIndex).IsZero() &&
		reflect.ValueOf(c.SlbSessionInfoIndex).IsZero() {
		return path
	}
	return path + "/" + fmt.Sprint(c.SlbSessionInfoSpIndex) + "/" + fmt.Sprint(c.SlbSessionInfoIndex)
}

type SlbSessionInfoTableParams struct {
	// The SP which created the session.
	SpIndex int32 `json:"SpIndex,omitempty"`
	// The index to the slbSessionInfoTable.
	Index int32 `json:"Index,omitempty"`
	// The string representing the SLB session entries in the formatted case:
	// 	 The fields, (1)-(13), associated with a session, as identified in the
	// 	 example below are described in the following.
	// 3, 01: 1.1.1.1 4586, 2.2.2.1 http -> 3567 3.3.3.1 http age 6 f:10 ELNPSRtUW c:#
	// (1) (2) (3)     (4)  (5)      (6)     (7)   (8)    (9)  (10)  (11)  (12) (13)
	// 3, 01: 1.1.1.1 4586, 2.2.2.1 http ->
	// (1) (2) (3)     (4)  (5)      (6)
	// 1.1.1.2 3567 3.3.3.1 http age 6 f:10 ELNPSRtUW c:#
	// (7a)     (7)   (8)    (9) (10)  (11)  (12) (13)
	// ------------------------------------------------------------------
	// (1) SP number:
	// This field indicates which SP created the session.
	// (2) Ingress port:
	// This field shows the physical port# of the client traffic that
	// entered to the switch.
	// (3) Source IP address:
	// This field contains the source IP address from client IP packet.
	// (4) Source port:
	// This field identifies the TCP/UDP source port from client packet.
	// (5) Destination IP address:
	// This is the destination IP address from client TCP/UDP packet.
	// For load balancing, this address is the virtual IP address.
	// For filtering redirect, this address is the destination server's address.
	// (6) Destination port:
	// This field identifies the TCP/UDP destination port from client packet.
	// (7a)Proxy IP address:
	// This field contains the proxy IP address substituted by the switch.
	// (7) Proxy port:
	// This field identifies the TCP/UDP source port substituted by the switch.
	// The switch does this substitution when Direct Access Mode is enabled or
	// Proxy is enabled.
	// (8) Real server IP address:
	// For load balancing, this field contains the IP address of the real
	// server that the switch selects to forward client packet to.
	// If the switch does not find live server, this field is the
	// same as destination IP address(5).
	// e.g. 3,01: 1.1.1.1 1040, 2.2.2.1 http -> 3.3.3.1 http age 10
	// 3,01: 1.1.1.1 1040, 2.2.2.1 http -> 2.2.2.1 http age 10
	// For filtering, this field also shows the real server IP address. No
	// address is shown if the filter action is Allow, Deny or NAT.
	// It will show ALLOW, DENY or NAT instead.
	// e.g. 3,01: 1.1.1.1 1040, 2.2.2.1 http -> 3.3.3.1 http age 10 f:11
	// 3,01: 1.1.1.1 1040, 2.2.2.1 http ALLOW age 10 f:22
	// (9) Server port:
	// For load balancing, this field is the same as destination port(6)
	// except for the RTSP UDP session. For RTSP UDP session, this server
	// port is obtained from client-server negotiation.
	// e.g. 3,01: 1.1.1.1 1040, 2.2.2.1 http -> 3.3.3.1 http age 10
	// 3,01: 1.1.1.1 6970, 2.2.2.1 rtsp -> 3.3.3.1 21220 age 10 P
	// For filtering, this field is the filtering application port.
	// It is for internal use only. This field can be urlwcr, wcr,
	// idslb, linkslb or nonat.
	// e.g. 3,01: 1.1.1.1 1040, 2.2.2.1 http -> 3.3.3.1 urlwcr age 4 f:123
	// 2,07: 1.1.1.1 1706, 2.2.2.1 http -> 192.168.4.10 linklb age 8 f:10 E
	// (10) Age:
	// This is the session timeout value. If no packet is received within
	// the value specified, the session is freed.
	// e.g. age 10    - The session is aged out in 10 minutes.
	// age < 160 - The session is aged out in 160 minutes. This <
	// indicates that slowage is used(user configures
	// /cfg/slb/adv/slowage)
	// (11) Filter number:
	// This field indicates that the session is created by filtering
	// code as a result of the IP header keys matching the filtering
	// criteria.
	// (12) Flag:
	// E: Indicates the session is in use and will be aged out if no
	// traffic is received within session timeout value.
	// L: Indicates the session is a link load balance session.
	// N: Indicates the session only translates the destination
	// MAC when forwarding client traffic to the real server.
	// P: Indicates the session is a persistent session and is not to
	// be aged out. Fields (6), (7) and (8) are not shown for
	// persistent session.
	// S: Indicates the session is a persistent session and the
	// application is SSL or Cookie Pbind.
	// Rt: Indicates the session is TCP rate limiting per-client entry.
	// Ru: Indicates the session is UDP rate limiting per-client entry.
	// Ri: Indicates the session is ICMP rate limiting per-client entry.
	// U: Indicates the session is L7 delay binding and switch is
	// trying to open TCP connection to real server.
	// W: Indicates the session only translates the destination
	// MAC when forwarding layer 7 WCR traffic to the real server.
	// (13) Persistent session user count:
	// This counter indicates the number of client sessions created
	// associated with this persistent session.
	StringVal string `json:"StringVal,omitempty"`
}

func (p SlbSessionInfoTableParams) iMABean() {}
