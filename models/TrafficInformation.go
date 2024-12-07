/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Sat Dec  7 16:57:14 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type TrafficInformation struct {
	UplinkVolume   *int64 `json:"uplinkVolume,omitempty"`
	DownlinkVolume *int64 `json:"downlinkVolume,omitempty"`
	TotalVolume    *int64 `json:"totalVolume,omitempty"`
	UplinkRate     string `json:"uplinkRate,omitempty"`
	DownlinkRate   string `json:"downlinkRate,omitempty"`
}
