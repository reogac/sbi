/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue May 12 13:32:30 KST 2026 by TungTQ<tqtung@etri.re.kr>
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
