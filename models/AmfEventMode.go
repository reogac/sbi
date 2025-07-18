/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jul 18 16:49:20 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type AmfEventMode struct {
	Expiry               string           `json:"expiry,omitempty"`
	RepPeriod            *int             `json:"repPeriod,omitempty"`
	SampRatio            *int             `json:"sampRatio,omitempty"`
	PartitioningCriteria []string         `json:"partitioningCriteria,omitempty"`
	NotifFlag            NotificationFlag `json:"notifFlag,omitempty"`
	Trigger              AmfEventTrigger  `json:"trigger"`
	MaxReports           *int             `json:"maxReports,omitempty"`
}
