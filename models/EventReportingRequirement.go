/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jul 18 16:49:20 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type EventReportingRequirement struct {
	Accuracy          Accuracy                     `json:"accuracy,omitempty"`
	EndTs             string                       `json:"endTs,omitempty"`
	SampRatio         *int                         `json:"sampRatio,omitempty"`
	MaxSupiNbr        *int                         `json:"maxSupiNbr,omitempty"`
	TimeAnaNeeded     string                       `json:"timeAnaNeeded,omitempty"`
	AnaMetaInd        *AnalyticsMetadataIndication `json:"anaMetaInd,omitempty"`
	HistAnaTimePeriod *TimeWindow                  `json:"histAnaTimePeriod,omitempty"`
	AccPerSubset      []string                     `json:"accPerSubset,omitempty"`
	StartTs           string                       `json:"startTs,omitempty"`
	OffsetPeriod      *int                         `json:"offsetPeriod,omitempty"`
	MaxObjectNbr      *int                         `json:"maxObjectNbr,omitempty"`
	AnaMeta           []string                     `json:"anaMeta,omitempty"`
}
