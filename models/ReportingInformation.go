/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue Jul 22 12:00:17 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type ReportingInformation struct {
	PartitionCriteria []string           `json:"partitionCriteria,omitempty"`
	NotifFlag         NotificationFlag   `json:"notifFlag,omitempty"`
	ImmRep            *bool              `json:"immRep,omitempty"`
	NotifMethod       NotificationMethod `json:"notifMethod,omitempty"`
	RepPeriod         *int               `json:"repPeriod,omitempty"`
	GrpRepTime        *int               `json:"grpRepTime,omitempty"`
	MaxReportNbr      *int               `json:"maxReportNbr,omitempty"`
	MonDur            string             `json:"monDur,omitempty"`
	SampRatio         *int               `json:"sampRatio,omitempty"`
}
