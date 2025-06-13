/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jun 13 13:39:10 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type ThresholdLevel struct {
	MaxPacketDelay    *int     `json:"maxPacketDelay,omitempty"`
	AvgPacketLossRate *int     `json:"avgPacketLossRate,omitempty"`
	CongLevel         *int     `json:"congLevel,omitempty"`
	NfLoadLevel       *int     `json:"nfLoadLevel,omitempty"`
	NfCpuUsage        *int     `json:"nfCpuUsage,omitempty"`
	NfMemoryUsage     *int     `json:"nfMemoryUsage,omitempty"`
	NfStorageUsage    *int     `json:"nfStorageUsage,omitempty"`
	AvgPacketDelay    *int     `json:"avgPacketDelay,omitempty"`
	SvcExpLevel       *float64 `json:"svcExpLevel,omitempty"`
	AvgTrafficRate    string   `json:"avgTrafficRate,omitempty"`
	MaxTrafficRate    string   `json:"maxTrafficRate,omitempty"`
}
