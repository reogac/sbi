/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Thu Jun 12 16:32:32 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type QosData struct {
	Qnc                  *bool  `json:"qnc,omitempty"`
	PriorityLevel        *int   `json:"priorityLevel,omitempty"`
	ReflectiveQos        *bool  `json:"reflectiveQos,omitempty"`
	MaxPacketLossRateDl  *int   `json:"maxPacketLossRateDl,omitempty"`
	DefQosFlowIndication *bool  `json:"defQosFlowIndication,omitempty"`
	ExtMaxDataBurstVol   *int   `json:"extMaxDataBurstVol,omitempty"`
	FiveQi               *int   `json:"5qi,omitempty"`
	GbrDl                string `json:"gbrDl,omitempty"`
	MaxDataBurstVol      *int   `json:"maxDataBurstVol,omitempty"`
	MaxPacketLossRateUl  *int   `json:"maxPacketLossRateUl,omitempty"`
	MaxbrDl              string `json:"maxbrDl,omitempty"`
	GbrUl                string `json:"gbrUl,omitempty"`
	AverWindow           *int   `json:"averWindow,omitempty"`
	SharingKeyDl         string `json:"sharingKeyDl,omitempty"`
	SharingKeyUl         string `json:"sharingKeyUl,omitempty"`
	PacketDelayBudget    *int   `json:"packetDelayBudget,omitempty"`
	QosId                string `json:"qosId"`
	MaxbrUl              string `json:"maxbrUl,omitempty"`
	Arp                  *Arp   `json:"arp,omitempty"`
	PacketErrorRate      string `json:"packetErrorRate,omitempty"`
}
