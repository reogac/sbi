/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jun 13 11:41:31 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type WlanPerTsPerformanceInfo struct {
	Rtt         *int                `json:"rtt,omitempty"`
	TrafficInfo *TrafficInformation `json:"trafficInfo,omitempty"`
	NumberOfUes *int                `json:"numberOfUes,omitempty"`
	Confidence  *int                `json:"confidence,omitempty"`
	TsStart     string              `json:"tsStart"`
	TsDuration  int                 `json:"tsDuration"`
	Rssi        *int                `json:"rssi,omitempty"`
}
