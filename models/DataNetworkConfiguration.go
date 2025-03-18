/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue Mar 18 17:35:08 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type DataNetworkConfiguration struct {
	Cidr          string  `json:"cidr,omitempty"`
	NumPools      *int16  `json:"numPools,omitempty"`
	PoolIndexList []int16 `json:"poolIndexList,omitempty"`
	Name          string  `json:"name"`
	DhcpServer    string  `json:"dhcpServer,omitempty"`
}
