/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue Jun 17 13:35:41 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type AnalyticsMetadataInfo struct {
	DataStatProps []string       `json:"dataStatProps,omitempty"`
	Strategy      OutputStrategy `json:"strategy,omitempty"`
	Accuracy      Accuracy       `json:"accuracy,omitempty"`
	NumSamples    *int           `json:"numSamples,omitempty"`
	DataWindow    *TimeWindow    `json:"dataWindow,omitempty"`
}
