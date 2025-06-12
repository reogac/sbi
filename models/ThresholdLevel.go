/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Thu Jun 12 16:32:15 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type ThresholdLevel struct {
	NfLoadLevel       *int     `json:"nfLoadLevel,omitempty"`
	NfCpuUsage        *int     `json:"nfCpuUsage,omitempty"`
	NfStorageUsage    *int     `json:"nfStorageUsage,omitempty"`
	MaxTrafficRate    string   `json:"maxTrafficRate,omitempty"`
	AvgPacketDelay    *int     `json:"avgPacketDelay,omitempty"`
	MaxPacketDelay    *int     `json:"maxPacketDelay,omitempty"`
	SvcExpLevel       *float64 `json:"svcExpLevel,omitempty"`
	CongLevel         *int     `json:"congLevel,omitempty"`
	AvgTrafficRate    string   `json:"avgTrafficRate,omitempty"`
	AvgPacketLossRate *int     `json:"avgPacketLossRate,omitempty"`
	NfMemoryUsage     *int     `json:"nfMemoryUsage,omitempty"`
}
