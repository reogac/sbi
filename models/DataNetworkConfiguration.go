/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Mon Mar 24 10:34:58 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type DataNetworkConfiguration struct {
	Cidr          string  `json:"cidr,omitempty"`
	IpRange       *int64  `json:"ipRange,omitempty"`
	PoolIndexList []int16 `json:"poolIndexList,omitempty"`
	Name          string  `json:"name"`
	DhcpServer    string  `json:"dhcpServer,omitempty"`
}
