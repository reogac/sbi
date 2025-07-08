/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue Jul  8 13:19:32 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type NonDynamic5Qi struct {
	PriorityLevel         *int `json:"priorityLevel,omitempty"`
	AverWindow            *int `json:"averWindow,omitempty"`
	MaxDataBurstVol       *int `json:"maxDataBurstVol,omitempty"`
	ExtMaxDataBurstVol    *int `json:"extMaxDataBurstVol,omitempty"`
	CnPacketDelayBudgetDl *int `json:"cnPacketDelayBudgetDl,omitempty"`
	CnPacketDelayBudgetUl *int `json:"cnPacketDelayBudgetUl,omitempty"`
}
