/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jul 18 15:09:35 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type QosFlowAddModifyRequestItem struct {
	Ebi                *int              `json:"ebi,omitempty"`
	QosRules           string            `json:"qosRules,omitempty"`
	QosFlowDescription string            `json:"qosFlowDescription,omitempty"`
	QosFlowProfile     *QosFlowProfile   `json:"qosFlowProfile,omitempty"`
	AssociatedAnType   QosFlowAccessType `json:"associatedAnType,omitempty"`
	Qfi                int               `json:"qfi"`
}
