/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue Jul 22 12:00:22 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type VolumeTimedReport struct {
	UplinkVolume   int64  `json:"uplinkVolume"`
	StartTimeStamp string `json:"startTimeStamp"`
	EndTimeStamp   string `json:"endTimeStamp"`
	DownlinkVolume int64  `json:"downlinkVolume"`
}
