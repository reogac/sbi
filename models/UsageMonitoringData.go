/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jul 18 15:09:46 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type UsageMonitoringData struct {
	VolumeThreshold          *int64   `json:"volumeThreshold,omitempty"`
	VolumeThresholdUplink    *int64   `json:"volumeThresholdUplink,omitempty"`
	TimeThreshold            *int     `json:"timeThreshold,omitempty"`
	MonitoringTime           string   `json:"monitoringTime,omitempty"`
	NextVolThresholdUplink   *int64   `json:"nextVolThresholdUplink,omitempty"`
	NextVolThresholdDownlink *int64   `json:"nextVolThresholdDownlink,omitempty"`
	NextTimeThreshold        *int     `json:"nextTimeThreshold,omitempty"`
	InactivityTime           *int     `json:"inactivityTime,omitempty"`
	UmId                     string   `json:"umId"`
	VolumeThresholdDownlink  *int64   `json:"volumeThresholdDownlink,omitempty"`
	NextVolThreshold         *int64   `json:"nextVolThreshold,omitempty"`
	ExUsagePccRuleIds        []string `json:"exUsagePccRuleIds,omitempty"`
}
