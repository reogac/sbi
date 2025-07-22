/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue Jul 22 12:00:17 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type EventReportingRequirement struct {
	EndTs             string                       `json:"endTs,omitempty"`
	OffsetPeriod      *int                         `json:"offsetPeriod,omitempty"`
	MaxObjectNbr      *int                         `json:"maxObjectNbr,omitempty"`
	TimeAnaNeeded     string                       `json:"timeAnaNeeded,omitempty"`
	AnaMeta           []string                     `json:"anaMeta,omitempty"`
	AnaMetaInd        *AnalyticsMetadataIndication `json:"anaMetaInd,omitempty"`
	HistAnaTimePeriod *TimeWindow                  `json:"histAnaTimePeriod,omitempty"`
	Accuracy          Accuracy                     `json:"accuracy,omitempty"`
	StartTs           string                       `json:"startTs,omitempty"`
	SampRatio         *int                         `json:"sampRatio,omitempty"`
	MaxSupiNbr        *int                         `json:"maxSupiNbr,omitempty"`
	AccPerSubset      []string                     `json:"accPerSubset,omitempty"`
}
