/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue Jul 22 12:00:17 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type AmfEventMode struct {
	PartitioningCriteria []string         `json:"partitioningCriteria,omitempty"`
	NotifFlag            NotificationFlag `json:"notifFlag,omitempty"`
	Trigger              AmfEventTrigger  `json:"trigger"`
	MaxReports           *int             `json:"maxReports,omitempty"`
	Expiry               string           `json:"expiry,omitempty"`
	RepPeriod            *int             `json:"repPeriod,omitempty"`
	SampRatio            *int             `json:"sampRatio,omitempty"`
}
