/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Sat Dec  7 16:57:14 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type Pc5QoSPara struct {
	Pc5QosFlowList []Pc5QosFlowItem `json:"pc5QosFlowList"`
	Pc5LinkAmbr    string           `json:"pc5LinkAmbr,omitempty"`
}
