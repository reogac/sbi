/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Wed Mar 19 09:38:39 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type DataNetworkConfiguration struct {
	Name          string  `json:"name"`
	DhcpServer    string  `json:"dhcpServer,omitempty"`
	Cidr          string  `json:"cidr,omitempty"`
	NumPools      *int16  `json:"numPools,omitempty"`
	PoolIndexList []int16 `json:"poolIndexList,omitempty"`
}
