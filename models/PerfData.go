/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Sat Dec  7 16:57:14 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type PerfData struct {
	MaxTrafficRate    string `json:"maxTrafficRate,omitempty"`
	AvePacketDelay    *int   `json:"avePacketDelay,omitempty"`
	MaxPacketDelay    *int   `json:"maxPacketDelay,omitempty"`
	AvgPacketLossRate *int   `json:"avgPacketLossRate,omitempty"`
	AvgTrafficRate    string `json:"avgTrafficRate,omitempty"`
}
