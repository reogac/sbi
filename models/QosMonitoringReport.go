/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jul 18 15:09:46 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type QosMonitoringReport struct {
	DlDelays      []int    `json:"dlDelays,omitempty"`
	RtDelays      []int    `json:"rtDelays,omitempty"`
	Pdmf          *bool    `json:"pdmf,omitempty"`
	RefPccRuleIds []string `json:"refPccRuleIds"`
	UlDelays      []int    `json:"ulDelays,omitempty"`
}
