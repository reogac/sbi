/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue May 12 13:32:26 KST 2026 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type WlanPerTsPerformanceInfo struct {
	TsDuration  int                 `json:"tsDuration"`
	Rssi        *int                `json:"rssi,omitempty"`
	Rtt         *int                `json:"rtt,omitempty"`
	TrafficInfo *TrafficInformation `json:"trafficInfo,omitempty"`
	NumberOfUes *int                `json:"numberOfUes,omitempty"`
	Confidence  *int                `json:"confidence,omitempty"`
	TsStart     string              `json:"tsStart"`
}
