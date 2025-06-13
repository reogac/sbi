/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jun 13 11:41:36 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type QosFlowReleaseRequestItem struct {
	Qfi                int    `json:"qfi"`
	QosRules           string `json:"qosRules,omitempty"`
	QosFlowDescription string `json:"qosFlowDescription,omitempty"`
}
