/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue Jul  8 13:19:44 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type QosMonitoringReport struct {
	RefPccRuleIds []string `json:"refPccRuleIds"`
	UlDelays      []int    `json:"ulDelays,omitempty"`
	DlDelays      []int    `json:"dlDelays,omitempty"`
	RtDelays      []int    `json:"rtDelays,omitempty"`
	Pdmf          *bool    `json:"pdmf,omitempty"`
}
