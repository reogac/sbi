/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jul 18 15:09:50 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type UsageThreshold struct {
	TotalVolume    *int64 `json:"totalVolume,omitempty"`
	DownlinkVolume *int64 `json:"downlinkVolume,omitempty"`
	UplinkVolume   *int64 `json:"uplinkVolume,omitempty"`
	Duration       *int   `json:"duration,omitempty"`
}
