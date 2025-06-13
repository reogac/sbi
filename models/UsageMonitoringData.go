/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jun 13 13:39:26 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type UsageMonitoringData struct {
	NextVolThresholdDownlink *int64   `json:"nextVolThresholdDownlink,omitempty"`
	InactivityTime           *int     `json:"inactivityTime,omitempty"`
	ExUsagePccRuleIds        []string `json:"exUsagePccRuleIds,omitempty"`
	UmId                     string   `json:"umId"`
	VolumeThresholdUplink    *int64   `json:"volumeThresholdUplink,omitempty"`
	VolumeThresholdDownlink  *int64   `json:"volumeThresholdDownlink,omitempty"`
	NextVolThreshold         *int64   `json:"nextVolThreshold,omitempty"`
	NextVolThresholdUplink   *int64   `json:"nextVolThresholdUplink,omitempty"`
	VolumeThreshold          *int64   `json:"volumeThreshold,omitempty"`
	TimeThreshold            *int     `json:"timeThreshold,omitempty"`
	MonitoringTime           string   `json:"monitoringTime,omitempty"`
	NextTimeThreshold        *int     `json:"nextTimeThreshold,omitempty"`
}
