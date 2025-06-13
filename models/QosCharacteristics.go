/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jun 13 11:41:48 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type QosCharacteristics struct {
	ExtMaxDataBurstVol *int            `json:"extMaxDataBurstVol,omitempty"`
	FiveQi             int             `json:"5qi"`
	ResourceType       QosResourceType `json:"resourceType"`
	PriorityLevel      int             `json:"priorityLevel"`
	PacketDelayBudget  int             `json:"packetDelayBudget"`
	PacketErrorRate    string          `json:"packetErrorRate"`
	AveragingWindow    *int            `json:"averagingWindow,omitempty"`
	MaxDataBurstVol    *int            `json:"maxDataBurstVol,omitempty"`
}
