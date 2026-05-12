/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue May 12 13:32:45 KST 2026 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type ReportingOptions struct {
	SamplingRatio   *int             `json:"samplingRatio,omitempty"`
	GuardTime       *int             `json:"guardTime,omitempty"`
	ReportPeriod    *int             `json:"reportPeriod,omitempty"`
	NotifFlag       NotificationFlag `json:"notifFlag,omitempty"`
	ReportMode      EventReportMode  `json:"reportMode,omitempty"`
	MaxNumOfReports *int             `json:"maxNumOfReports,omitempty"`
	Expiry          string           `json:"expiry,omitempty"`
}
