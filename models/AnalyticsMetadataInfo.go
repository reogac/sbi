/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jul 18 16:49:20 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type AnalyticsMetadataInfo struct {
	Strategy      OutputStrategy `json:"strategy,omitempty"`
	Accuracy      Accuracy       `json:"accuracy,omitempty"`
	NumSamples    *int           `json:"numSamples,omitempty"`
	DataWindow    *TimeWindow    `json:"dataWindow,omitempty"`
	DataStatProps []string       `json:"dataStatProps,omitempty"`
}
