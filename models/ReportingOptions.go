/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue Jul 22 12:00:37 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type ReportingOptions struct {
	ReportMode      EventReportMode  `json:"reportMode,omitempty"`
	MaxNumOfReports *int             `json:"maxNumOfReports,omitempty"`
	Expiry          string           `json:"expiry,omitempty"`
	SamplingRatio   *int             `json:"samplingRatio,omitempty"`
	GuardTime       *int             `json:"guardTime,omitempty"`
	ReportPeriod    *int             `json:"reportPeriod,omitempty"`
	NotifFlag       NotificationFlag `json:"notifFlag,omitempty"`
}
