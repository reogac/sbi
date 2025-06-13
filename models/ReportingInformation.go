/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jun 13 13:39:10 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type ReportingInformation struct {
	SampRatio         *int               `json:"sampRatio,omitempty"`
	PartitionCriteria []string           `json:"partitionCriteria,omitempty"`
	ImmRep            *bool              `json:"immRep,omitempty"`
	NotifMethod       NotificationMethod `json:"notifMethod,omitempty"`
	MaxReportNbr      *int               `json:"maxReportNbr,omitempty"`
	MonDur            string             `json:"monDur,omitempty"`
	RepPeriod         *int               `json:"repPeriod,omitempty"`
	GrpRepTime        *int               `json:"grpRepTime,omitempty"`
	NotifFlag         NotificationFlag   `json:"notifFlag,omitempty"`
}
