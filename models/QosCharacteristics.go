/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue May 12 13:32:42 KST 2026 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type QosCharacteristics struct {
	PriorityLevel      int             `json:"priorityLevel"`
	PacketDelayBudget  int             `json:"packetDelayBudget"`
	PacketErrorRate    string          `json:"packetErrorRate"`
	AveragingWindow    *int            `json:"averagingWindow,omitempty"`
	MaxDataBurstVol    *int            `json:"maxDataBurstVol,omitempty"`
	ExtMaxDataBurstVol *int            `json:"extMaxDataBurstVol,omitempty"`
	FiveQi             int             `json:"5qi"`
	ResourceType       QosResourceType `json:"resourceType"`
}
