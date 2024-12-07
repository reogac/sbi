/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Sat Dec  7 16:57:30 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type UeInitiatedResourceRequest struct {
	PackFiltInfo []PacketFilterInfo `json:"packFiltInfo"`
	ReqQos       *RequestedQos      `json:"reqQos,omitempty"`
	PccRuleId    string             `json:"pccRuleId,omitempty"`
	RuleOp       RuleOperation      `json:"ruleOp"`
	Precedence   *int               `json:"precedence,omitempty"`
}
