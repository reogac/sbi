/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jul 18 15:09:30 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type Pc5QoSPara struct {
	Pc5QosFlowList []Pc5QosFlowItem `json:"pc5QosFlowList"`
	Pc5LinkAmbr    string           `json:"pc5LinkAmbr,omitempty"`
}
