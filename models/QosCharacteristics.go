/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Sat Dec  7 16:57:30 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type QosCharacteristics struct {
	PacketErrorRate    string          `json:"packetErrorRate"`
	AveragingWindow    *int            `json:"averagingWindow,omitempty"`
	MaxDataBurstVol    *int            `json:"maxDataBurstVol,omitempty"`
	ExtMaxDataBurstVol *int            `json:"extMaxDataBurstVol,omitempty"`
	FiveQi             int             `json:"5qi"`
	ResourceType       QosResourceType `json:"resourceType"`
	PriorityLevel      int             `json:"priorityLevel"`
	PacketDelayBudget  int             `json:"packetDelayBudget"`
}
