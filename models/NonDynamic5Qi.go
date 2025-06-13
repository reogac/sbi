/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jun 13 11:41:36 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type NonDynamic5Qi struct {
	AverWindow            *int `json:"averWindow,omitempty"`
	MaxDataBurstVol       *int `json:"maxDataBurstVol,omitempty"`
	ExtMaxDataBurstVol    *int `json:"extMaxDataBurstVol,omitempty"`
	CnPacketDelayBudgetDl *int `json:"cnPacketDelayBudgetDl,omitempty"`
	CnPacketDelayBudgetUl *int `json:"cnPacketDelayBudgetUl,omitempty"`
	PriorityLevel         *int `json:"priorityLevel,omitempty"`
}
