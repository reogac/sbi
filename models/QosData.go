/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Sat Dec  7 16:57:30 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type QosData struct {
	Qnc                  *bool  `json:"qnc,omitempty"`
	PriorityLevel        *int   `json:"priorityLevel,omitempty"`
	AverWindow           *int   `json:"averWindow,omitempty"`
	SharingKeyUl         string `json:"sharingKeyUl,omitempty"`
	MaxPacketLossRateUl  *int   `json:"maxPacketLossRateUl,omitempty"`
	DefQosFlowIndication *bool  `json:"defQosFlowIndication,omitempty"`
	Arp                  *Arp   `json:"arp,omitempty"`
	MaxbrUl              string `json:"maxbrUl,omitempty"`
	MaxPacketLossRateDl  *int   `json:"maxPacketLossRateDl,omitempty"`
	ExtMaxDataBurstVol   *int   `json:"extMaxDataBurstVol,omitempty"`
	PacketErrorRate      string `json:"packetErrorRate,omitempty"`
	QosId                string `json:"qosId"`
	GbrUl                string `json:"gbrUl,omitempty"`
	GbrDl                string `json:"gbrDl,omitempty"`
	ReflectiveQos        *bool  `json:"reflectiveQos,omitempty"`
	SharingKeyDl         string `json:"sharingKeyDl,omitempty"`
	PacketDelayBudget    *int   `json:"packetDelayBudget,omitempty"`
	FiveQi               *int   `json:"5qi,omitempty"`
	MaxDataBurstVol      *int   `json:"maxDataBurstVol,omitempty"`
	MaxbrDl              string `json:"maxbrDl,omitempty"`
}
