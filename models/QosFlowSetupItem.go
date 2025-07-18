/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jul 18 15:09:35 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type QosFlowSetupItem struct {
	Ebi                *int              `json:"ebi,omitempty"`
	QosFlowDescription string            `json:"qosFlowDescription,omitempty"`
	QosFlowProfile     *QosFlowProfile   `json:"qosFlowProfile,omitempty"`
	AssociatedAnType   QosFlowAccessType `json:"associatedAnType,omitempty"`
	DefaultQosRuleInd  *bool             `json:"defaultQosRuleInd,omitempty"`
	Qfi                int               `json:"qfi"`
	QosRules           string            `json:"qosRules"`
}
