/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue May 12 13:32:26 KST 2026 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type ThresholdLevel struct {
	NfLoadLevel       *int     `json:"nfLoadLevel,omitempty"`
	NfMemoryUsage     *int     `json:"nfMemoryUsage,omitempty"`
	NfStorageUsage    *int     `json:"nfStorageUsage,omitempty"`
	AvgTrafficRate    string   `json:"avgTrafficRate,omitempty"`
	MaxTrafficRate    string   `json:"maxTrafficRate,omitempty"`
	AvgPacketLossRate *int     `json:"avgPacketLossRate,omitempty"`
	SvcExpLevel       *float64 `json:"svcExpLevel,omitempty"`
	CongLevel         *int     `json:"congLevel,omitempty"`
	NfCpuUsage        *int     `json:"nfCpuUsage,omitempty"`
	AvgPacketDelay    *int     `json:"avgPacketDelay,omitempty"`
	MaxPacketDelay    *int     `json:"maxPacketDelay,omitempty"`
}
