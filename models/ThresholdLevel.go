/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jul 18 15:09:30 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type ThresholdLevel struct {
	SvcExpLevel       *float64 `json:"svcExpLevel,omitempty"`
	CongLevel         *int     `json:"congLevel,omitempty"`
	NfLoadLevel       *int     `json:"nfLoadLevel,omitempty"`
	AvgTrafficRate    string   `json:"avgTrafficRate,omitempty"`
	MaxPacketDelay    *int     `json:"maxPacketDelay,omitempty"`
	AvgPacketDelay    *int     `json:"avgPacketDelay,omitempty"`
	AvgPacketLossRate *int     `json:"avgPacketLossRate,omitempty"`
	NfCpuUsage        *int     `json:"nfCpuUsage,omitempty"`
	NfMemoryUsage     *int     `json:"nfMemoryUsage,omitempty"`
	NfStorageUsage    *int     `json:"nfStorageUsage,omitempty"`
	MaxTrafficRate    string   `json:"maxTrafficRate,omitempty"`
}
