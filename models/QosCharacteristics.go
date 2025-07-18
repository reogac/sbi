/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jul 18 16:49:36 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type QosCharacteristics struct {
	AveragingWindow    *int            `json:"averagingWindow,omitempty"`
	MaxDataBurstVol    *int            `json:"maxDataBurstVol,omitempty"`
	ExtMaxDataBurstVol *int            `json:"extMaxDataBurstVol,omitempty"`
	FiveQi             int             `json:"5qi"`
	ResourceType       QosResourceType `json:"resourceType"`
	PriorityLevel      int             `json:"priorityLevel"`
	PacketDelayBudget  int             `json:"packetDelayBudget"`
	PacketErrorRate    string          `json:"packetErrorRate"`
}
