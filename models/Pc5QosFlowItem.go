/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Sat Dec  7 16:57:14 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type Pc5QosFlowItem struct {
	Pqi             int              `json:"pqi"`
	Pc5FlowBitRates *Pc5FlowBitRates `json:"pc5FlowBitRates,omitempty"`
	Range           *int             `json:"range,omitempty"`
}
