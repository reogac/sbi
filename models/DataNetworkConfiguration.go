/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jul 18 15:09:29 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type DataNetworkConfiguration struct {
	Pcscf         *IpAddr `json:"pcscf,omitempty"`
	DhcpServer    string  `json:"dhcpServer,omitempty"`
	Cidr          string  `json:"cidr,omitempty"`
	IpRange       *int64  `json:"ipRange,omitempty"`
	PoolIndexList []int16 `json:"poolIndexList,omitempty"`
	Name          string  `json:"name"`
	Dns           *IpAddr `json:"dns,omitempty"`
}
