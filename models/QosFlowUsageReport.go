/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Thu Jun 12 16:32:20 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type QosFlowUsageReport struct {
	UplinkVolume   int64  `json:"uplinkVolume"`
	Qfi            int    `json:"qfi"`
	StartTimeStamp string `json:"startTimeStamp"`
	EndTimeStamp   string `json:"endTimeStamp"`
	DownlinkVolume int64  `json:"downlinkVolume"`
}
