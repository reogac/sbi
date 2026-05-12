/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue May 12 13:32:25 KST 2026 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type DataNetworkConfiguration struct {
	Cidr          string  `json:"cidr,omitempty"`
	IpRange       *int64  `json:"ipRange,omitempty"`
	PoolIndexList []int16 `json:"poolIndexList,omitempty"`
	Name          string  `json:"name"`
	Dns           *IpAddr `json:"dns,omitempty"`
	Pcscf         *IpAddr `json:"pcscf,omitempty"`
	DhcpServer    string  `json:"dhcpServer,omitempty"`
}
