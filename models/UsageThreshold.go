/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jun 13 11:28:33 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type UsageThreshold struct {
	UplinkVolume   *int64 `json:"uplinkVolume,omitempty"`
	Duration       *int   `json:"duration,omitempty"`
	TotalVolume    *int64 `json:"totalVolume,omitempty"`
	DownlinkVolume *int64 `json:"downlinkVolume,omitempty"`
}
