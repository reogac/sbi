/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jun 13 13:39:10 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type EventReportingRequirement struct {
	AccPerSubset      []string                     `json:"accPerSubset,omitempty"`
	MaxSupiNbr        *int                         `json:"maxSupiNbr,omitempty"`
	TimeAnaNeeded     string                       `json:"timeAnaNeeded,omitempty"`
	AnaMetaInd        *AnalyticsMetadataIndication `json:"anaMetaInd,omitempty"`
	MaxObjectNbr      *int                         `json:"maxObjectNbr,omitempty"`
	AnaMeta           []string                     `json:"anaMeta,omitempty"`
	HistAnaTimePeriod *TimeWindow                  `json:"histAnaTimePeriod,omitempty"`
	Accuracy          Accuracy                     `json:"accuracy,omitempty"`
	StartTs           string                       `json:"startTs,omitempty"`
	EndTs             string                       `json:"endTs,omitempty"`
	OffsetPeriod      *int                         `json:"offsetPeriod,omitempty"`
	SampRatio         *int                         `json:"sampRatio,omitempty"`
}
