/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jul 18 16:49:25 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type VolumeTimedReport struct {
	DownlinkVolume int64  `json:"downlinkVolume"`
	UplinkVolume   int64  `json:"uplinkVolume"`
	StartTimeStamp string `json:"startTimeStamp"`
	EndTimeStamp   string `json:"endTimeStamp"`
}
