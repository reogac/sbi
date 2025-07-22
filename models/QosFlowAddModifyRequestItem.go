/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue Jul 22 12:00:22 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type QosFlowAddModifyRequestItem struct {
	AssociatedAnType   QosFlowAccessType `json:"associatedAnType,omitempty"`
	Qfi                int               `json:"qfi"`
	Ebi                *int              `json:"ebi,omitempty"`
	QosRules           string            `json:"qosRules,omitempty"`
	QosFlowDescription string            `json:"qosFlowDescription,omitempty"`
	QosFlowProfile     *QosFlowProfile   `json:"qosFlowProfile,omitempty"`
}
