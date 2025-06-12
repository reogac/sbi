/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Thu Jun 12 16:32:14 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type DataNetworkConfiguration struct {
	DhcpServer    string  `json:"dhcpServer,omitempty"`
	Cidr          string  `json:"cidr,omitempty"`
	IpRange       *int64  `json:"ipRange,omitempty"`
	PoolIndexList []int16 `json:"poolIndexList,omitempty"`
	Name          string  `json:"name"`
	Dns           *IpAddr `json:"dns,omitempty"`
	Pcscf         *IpAddr `json:"pcscf,omitempty"`
}
