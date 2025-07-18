/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jul 18 16:49:20 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type DnPerformanceReq struct {
	DnPerfOrderCriter DnPerfOrderingCriterion `json:"dnPerfOrderCriter,omitempty"`
	Order             MatchingDirection       `json:"order,omitempty"`
	ReportThresholds  []ThresholdLevel        `json:"reportThresholds,omitempty"`
}
