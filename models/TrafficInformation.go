/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue May 12 13:32:26 KST 2026 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type TrafficInformation struct {
	TotalVolume    *int64 `json:"totalVolume,omitempty"`
	UplinkRate     string `json:"uplinkRate,omitempty"`
	DownlinkRate   string `json:"downlinkRate,omitempty"`
	UplinkVolume   *int64 `json:"uplinkVolume,omitempty"`
	DownlinkVolume *int64 `json:"downlinkVolume,omitempty"`
}
