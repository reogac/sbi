/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue May 12 13:32:26 KST 2026 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type AnalyticsMetadataIndication struct {
	DataWindow    *TimeWindow    `json:"dataWindow,omitempty"`
	DataStatProps []string       `json:"dataStatProps,omitempty"`
	Strategy      OutputStrategy `json:"strategy,omitempty"`
	AggrNwdafIds  []string       `json:"aggrNwdafIds,omitempty"`
}
