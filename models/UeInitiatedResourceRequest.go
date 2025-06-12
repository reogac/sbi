/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Thu Jun 12 16:32:32 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type UeInitiatedResourceRequest struct {
	PccRuleId    string             `json:"pccRuleId,omitempty"`
	RuleOp       RuleOperation      `json:"ruleOp"`
	Precedence   *int               `json:"precedence,omitempty"`
	PackFiltInfo []PacketFilterInfo `json:"packFiltInfo"`
	ReqQos       *RequestedQos      `json:"reqQos,omitempty"`
}
