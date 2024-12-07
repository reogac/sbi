/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Sat Dec  7 16:57:19 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type QosFlowSetupItem struct {
	DefaultQosRuleInd  *bool             `json:"defaultQosRuleInd,omitempty"`
	Qfi                int               `json:"qfi"`
	QosRules           string            `json:"qosRules"`
	Ebi                *int              `json:"ebi,omitempty"`
	QosFlowDescription string            `json:"qosFlowDescription,omitempty"`
	QosFlowProfile     *QosFlowProfile   `json:"qosFlowProfile,omitempty"`
	AssociatedAnType   QosFlowAccessType `json:"associatedAnType,omitempty"`
}
