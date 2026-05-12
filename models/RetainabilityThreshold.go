/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue May 12 13:32:26 KST 2026 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type RetainabilityThreshold struct {
	RelFlowNum   *int     `json:"relFlowNum,omitempty"`
	RelTimeUnit  TimeUnit `json:"relTimeUnit,omitempty"`
	RelFlowRatio *int     `json:"relFlowRatio,omitempty"`
}
