/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jun 13 11:41:51 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type ReportingOptions struct {
	ReportPeriod    *int             `json:"reportPeriod,omitempty"`
	NotifFlag       NotificationFlag `json:"notifFlag,omitempty"`
	ReportMode      EventReportMode  `json:"reportMode,omitempty"`
	MaxNumOfReports *int             `json:"maxNumOfReports,omitempty"`
	Expiry          string           `json:"expiry,omitempty"`
	SamplingRatio   *int             `json:"samplingRatio,omitempty"`
	GuardTime       *int             `json:"guardTime,omitempty"`
}
