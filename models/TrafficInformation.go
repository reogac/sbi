/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jun 13 13:39:10 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type TrafficInformation struct {
	UplinkRate     string `json:"uplinkRate,omitempty"`
	DownlinkRate   string `json:"downlinkRate,omitempty"`
	UplinkVolume   *int64 `json:"uplinkVolume,omitempty"`
	DownlinkVolume *int64 `json:"downlinkVolume,omitempty"`
	TotalVolume    *int64 `json:"totalVolume,omitempty"`
}
