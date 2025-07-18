/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jul 18 15:09:35 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type Dynamic5Qi struct {
	PacketErrRate         string          `json:"packetErrRate"`
	ExtMaxDataBurstVol    *int            `json:"extMaxDataBurstVol,omitempty"`
	ExtPacketDelBudget    *int            `json:"extPacketDelBudget,omitempty"`
	CnPacketDelayBudgetDl *int            `json:"cnPacketDelayBudgetDl,omitempty"`
	PriorityLevel         int             `json:"priorityLevel"`
	PacketDelayBudget     int             `json:"packetDelayBudget"`
	MaxDataBurstVol       *int            `json:"maxDataBurstVol,omitempty"`
	CnPacketDelayBudgetUl *int            `json:"cnPacketDelayBudgetUl,omitempty"`
	ResourceType          QosResourceType `json:"resourceType"`
	AverWindow            *int            `json:"averWindow,omitempty"`
}
