/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue May 12 13:32:26 KST 2026 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type ReportingInformation struct {
	ImmRep            *bool              `json:"immRep,omitempty"`
	MonDur            string             `json:"monDur,omitempty"`
	RepPeriod         *int               `json:"repPeriod,omitempty"`
	GrpRepTime        *int               `json:"grpRepTime,omitempty"`
	NotifFlag         NotificationFlag   `json:"notifFlag,omitempty"`
	NotifMethod       NotificationMethod `json:"notifMethod,omitempty"`
	MaxReportNbr      *int               `json:"maxReportNbr,omitempty"`
	SampRatio         *int               `json:"sampRatio,omitempty"`
	PartitionCriteria []string           `json:"partitionCriteria,omitempty"`
}
