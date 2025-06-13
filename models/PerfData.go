/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jun 13 13:39:10 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type PerfData struct {
	AvgTrafficRate    string `json:"avgTrafficRate,omitempty"`
	MaxTrafficRate    string `json:"maxTrafficRate,omitempty"`
	AvePacketDelay    *int   `json:"avePacketDelay,omitempty"`
	MaxPacketDelay    *int   `json:"maxPacketDelay,omitempty"`
	AvgPacketLossRate *int   `json:"avgPacketLossRate,omitempty"`
}
