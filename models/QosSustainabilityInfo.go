/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jun 13 13:39:10 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type QosSustainabilityInfo struct {
	Snssai        *Snssai                 `json:"snssai,omitempty"`
	Confidence    *int                    `json:"confidence,omitempty"`
	AreaInfo      *NetworkAreaInfo        `json:"areaInfo,omitempty"`
	StartTs       string                  `json:"startTs,omitempty"`
	EndTs         string                  `json:"endTs,omitempty"`
	QosFlowRetThd *RetainabilityThreshold `json:"qosFlowRetThd,omitempty"`
	RanUeThrouThd string                  `json:"ranUeThrouThd,omitempty"`
}
